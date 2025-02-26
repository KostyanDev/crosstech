package integration_test

import (
	"app/internal/domain"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

func (s *TestSuite) TestUploadJsonFile() {
	s.ClearTestData(context.TODO())
	requestPayload := map[string]string{"file_name": "test_data.json"}
	requestBody, _ := json.Marshal(requestPayload)

	req, _ := http.NewRequest(http.MethodPost, s.server.URL+"/upload_file", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	s.Require().NoError(err)
	defer resp.Body.Close()

	s.Equal(http.StatusOK, resp.StatusCode)

	var responseBody map[string]string
	_ = json.NewDecoder(resp.Body).Decode(&responseBody)

	trackID := 11111111
	tracks, err := s.storage.GetTracks(context.Background(), domain.GetTrack{ID: &trackID})
	s.Require().NoError(err)

	signals, err := s.storage.GetSignals(context.Background(), domain.GetSignal{TrackID: &trackID})
	s.Require().NoError(err)

	s.Equal("success", responseBody["status"])
	s.Equal("File processed successfully", responseBody["message"])
	s.Equal(trackID, signals[0].TrackID)
	s.Equal(trackID, tracks[0].ID)
}

func (s *TestSuite) ClearTestData(ctx context.Context) {
	_, err := s.db.ExecContext(ctx, `DELETE FROM signals WHERE track_id = $1`, 11111111)
	s.Require().NoError(err, "Failed to delete signals")

	_, err = s.db.ExecContext(ctx, `DELETE FROM tracks WHERE id = $1`, 11111111)
	s.Require().NoError(err, "Failed to delete tracks")

}

func (s *TestSuite) ProcessFile(ctx context.Context) {
	file := domain.File{Name: "test_data.json"}
	err := s.service.ProcessJSONFile(ctx, file)
	s.Require().NoError(err, "Failed to process JSON file")
}
