package service

import (
	"app/internal/domain"
	"context"
	"errors"
	"fmt"
)

func (s Service) CreateTrack(ctx context.Context, request domain.Track) error {
	isExist, err := s.storage.TrackExists(ctx, request.ID)
	if err != nil {
		return err
	}

	if isExist {
		return errors.New(fmt.Sprintf("track with id %s alredy created", request.ID))
	}

	if len(request.Signals) > 0 {
		for i := range request.Signals {
			if err := s.storage.CreateSignal(ctx, request.Signals[i]); err != nil {
				return err
			}
		}
	}

	return s.storage.CreateTrack(ctx, request)
}

func (s Service) GetTrackByParam(ctx context.Context, request domain.GetTrack) ([]domain.Track, error) {
	return s.storage.GetTracks(ctx, request)
}

func (s Service) UpdateTrackByParam(ctx context.Context, request domain.UpdateTrack) error {
	isTrackExist, err := s.storage.TrackExists(ctx, request.ID)
	if err != nil {
		return err
	}

	if !isTrackExist {
		return errors.New(fmt.Sprintf("track with id %s not found", request.ID))
	}

	return s.storage.UpdateTrack(ctx, request)
}
