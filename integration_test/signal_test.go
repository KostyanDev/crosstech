package integration_test

import (
	"app/internal/domain"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func (s *TestSuite) TestCreateSignal() {
	s.ClearTestData(context.Background())
	s.ProcessFile(context.Background())

	requestPayload := map[string]interface{}{
		"signal_id":   1,
		"signal_name": "Test Signal",
		"elr":         "ELR123",
		"mileage":     10.5,
		"track_id":    11111111,
	}
	requestBody, err := json.Marshal(requestPayload)
	s.Require().NoError(err)

	req, err := http.NewRequest(http.MethodPost, s.server.URL+"/signal/create", bytes.NewBuffer(requestBody))
	s.Require().NoError(err)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	s.router.ServeHTTP(recorder, req)

	s.Equal(http.StatusOK, recorder.Code, "Expected HTTP 200 response")

	var response map[string]string
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	s.Require().NoError(err)

	signalID := 1
	signalsStorage, err := s.storage.GetSignals(context.Background(), domain.GetSignal{ID: &signalID})

	s.Equal("success", response["status"])
	s.Equal("Signal created successfully", response["message"])
	s.Equal(signalsStorage[0].ID, signalID)
}

func (s *TestSuite) TestUpdateSignal() {
	s.ClearTestData(context.Background())
	s.ProcessFile(context.Background())

	requestPayload := map[string]interface{}{
		"signal_id":   111111111,
		"signal_name": "Updated Signal",
		"elr":         "ELR456",
		"mileage":     12.0,
		"track_id":    11111111,
	}
	requestBody, err := json.Marshal(requestPayload)
	s.Require().NoError(err)

	req, err := http.NewRequest(http.MethodPost, s.server.URL+"/signal/update", bytes.NewBuffer(requestBody))
	s.Require().NoError(err)

	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	s.router.ServeHTTP(recorder, req)

	s.Equal(http.StatusOK, recorder.Code, "Expected HTTP 200 response")

	signalID := 111111111
	signalsStorage, err := s.storage.GetSignals(context.Background(), domain.GetSignal{ID: &signalID})

	var response map[string]string
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	s.Require().NoError(err)
	s.Equal("success", response["status"])
	s.Equal("Signal updated successfully", response["message"])
	s.Equal(signalsStorage[0].ID, signalID)
	s.Equal(signalsStorage[0].Name, "Updated Signal")
}

func (s *TestSuite) TestGetSignals() {
	s.ClearTestData(context.Background())
	s.ProcessFile(context.Background())

	req, err := http.NewRequest(http.MethodGet, s.server.URL+"/signals?signal_id=111111111&track_id=11111111", nil)
	s.Require().NoError(err)

	recorder := httptest.NewRecorder()
	s.router.ServeHTTP(recorder, req)

	s.Equal(http.StatusOK, recorder.Code, "Expected HTTP 200 response")

	var signals []domain.Signal
	err = json.Unmarshal(recorder.Body.Bytes(), &signals)
	s.Require().NoError(err)

	s.Require().Greater(len(signals), 0, "Expected at least one signal in the response")
	s.Equal(111111111, signals[0].ID, "Expected Signal ID to match")
	s.Equal(11111111, signals[0].TrackID, "Expected Track ID to match")
}
