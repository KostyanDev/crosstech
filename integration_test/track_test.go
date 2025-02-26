package integration_test

import (
	"app/internal/domain"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func (s *TestSuite) TestCreateTrack() {
	s.ClearTestData(context.Background())

	requestPayload := map[string]interface{}{
		"track_id": 11111111,
		"source":   "Station A",
		"target":   "Station B",
	}

	requestBody, err := json.Marshal(requestPayload)
	s.Require().NoError(err)

	req, err := http.NewRequest(http.MethodPost, s.server.URL+"/track/create", bytes.NewBuffer(requestBody))
	s.Require().NoError(err)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	s.router.ServeHTTP(recorder, req)

	s.Equal(http.StatusCreated, recorder.Code, "Expected HTTP 201 response")

	var response map[string]string
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	s.Require().NoError(err)

	s.Equal("success", response["status"])
	s.Equal("Track created successfully", response["message"])

	trackID := 11111111
	tracksStorage, err := s.storage.GetTracks(context.Background(), domain.GetTrack{ID: &trackID})
	s.Require().NoError(err)
	s.Require().Greater(len(tracksStorage), 0, "Expected track to be stored in DB")
	s.Equal(tracksStorage[0].ID, trackID)
}

func (s *TestSuite) TestUpdateTrack() {
	s.ClearTestData(context.Background())
	s.ProcessFile(context.Background())

	requestPayload := map[string]interface{}{
		"track_id": 11111111,
		"source":   "Station C",
		"target":   "Station D",
	}

	requestBody, err := json.Marshal(requestPayload)
	s.Require().NoError(err)

	req, err := http.NewRequest(http.MethodPost, s.server.URL+"/track/update", bytes.NewBuffer(requestBody))
	s.Require().NoError(err)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	s.router.ServeHTTP(recorder, req)

	s.Equal(http.StatusOK, recorder.Code, "Expected HTTP 200 response")

	var response map[string]string
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	s.Require().NoError(err)

	s.Equal("success", response["status"])
	s.Equal("Track updated successfully", response["message"])

	trackID := 11111111
	tracksStorage, err := s.storage.GetTracks(context.Background(), domain.GetTrack{ID: &trackID})
	s.Require().NoError(err)
	s.Require().Greater(len(tracksStorage), 0, "Expected track to exist in DB")

	s.Equal(tracksStorage[0].ID, trackID)
	s.Equal(tracksStorage[0].Source, "Station C")
	s.Equal(tracksStorage[0].Target, "Station D")
}

func (s *TestSuite) TestGetTracks() {
	s.ClearTestData(context.Background())
	s.ProcessFile(context.Background())

	req, err := http.NewRequest(http.MethodGet, s.server.URL+"/tracks?track_id=11111111", nil)
	s.Require().NoError(err)

	recorder := httptest.NewRecorder()
	s.router.ServeHTTP(recorder, req)

	s.Equal(http.StatusOK, recorder.Code, "Expected HTTP 200 response")

	var tracks []domain.Track
	err = json.Unmarshal(recorder.Body.Bytes(), &tracks)
	s.Require().NoError(err)

	s.Require().Greater(len(tracks), 0, "Expected at least one track in the response")

	s.Equal(11111111, tracks[0].ID, "Expected Track ID to match")
	s.Equal("Acton Central", tracks[0].Source, "Expected Source to match")
	s.Equal("Willesden Junction", tracks[0].Target, "Expected Target to match")
}
