package service

import (
	"app/internal/domain"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sync"
)

func (s Service) ProcessJSONFile(ctx context.Context, file domain.File) error {
	filePath := filepath.Join("raw_data", file.Name)
	fileRawData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read JSON file: %w", err)
	}

	cleanedJSON := bytes.ReplaceAll(fileRawData, []byte("NaN"), []byte("null"))

	var tracks []domain.Track
	if err := json.Unmarshal(cleanedJSON, &tracks); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(tracks))

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, track := range tracks {
			if err := s.storage.CreateTrack(ctx, track); err != nil {
				s.log.Errorf("Error inserting track %d: %v", track.ID, err)
				errCh <- err
				continue
			}

			for _, signal := range track.Signals {
				signal.TrackID = track.ID

				if math.IsNaN(signal.Mileage) {
					signal.Mileage = 0
				}

				if err := s.storage.CreateSignal(ctx, signal); err != nil {
					s.log.Errorf("Error inserting signal %d: %v", signal.ID, err)
					errCh <- err
				}
			}
		}
	}()

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return fmt.Errorf("error processing JSON: %w", err)
		}
	}

	s.log.Info("JSON data processed successfully!")
	return nil
}
