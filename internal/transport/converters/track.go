package converters

import (
	"app/internal/domain"
	"fmt"
	"net/url"
	"strconv"
)

type CreatTrackRequest struct {
	ID      int                  `json:"track_id"`
	Source  string               `json:"source"`
	Target  string               `json:"target"`
	Signals []CreatSignalRequest `json:"signal,omitempty"`
}
type UpdateTrackRequest struct {
	ID        int     `json:"track_id"`
	Source    *string `json:"source"`
	Target    *string `json:"target"`
	IsDeleted *bool   `json:"is_deleted"`
}
type GetTrackRequest struct {
	ID     *int    `json:"track_id"`
	Source *string `json:"source"`
	Target *string `json:"target"`
}

type TrackResponse struct {
	ID     int    `json:"track_id"`
	Source string `json:"source"`
	Target string `json:"target"`
}

func ToDomainCreatTrack(req CreatTrackRequest) domain.Track {
	retVal := domain.Track{
		ID:     req.ID,
		Source: req.Source,
		Target: req.Target,
	}
	if req.Signals != nil {
		signals := make([]domain.Signal, len(req.Signals))
		for i := range req.Signals {
			signals = append(signals, ToDomainCreatSignal(req.Signals[i]))
		}
		retVal.Signals = signals
	}

	return retVal
}
func ToDomainUpdateTrack(req UpdateTrackRequest) domain.UpdateTrack {
	retVal := domain.UpdateTrack{
		ID: req.ID,
	}

	if req.Target != nil {
		retVal.Target = req.Target
	}
	if req.Source != nil {
		retVal.Source = req.Source
	}
	if req.IsDeleted != nil {
		retVal.IsDeleted = req.IsDeleted
	}

	return retVal
}
func ToRespTrack(resp domain.Track) TrackResponse {
	return TrackResponse{
		ID:     resp.ID,
		Source: resp.Source,
		Target: resp.Target,
	}
}
func ToRespTracks(resp []domain.Track) []TrackResponse {
	retVal := make([]TrackResponse, len(resp))
	for i, val := range resp {
		retVal[i] = ToRespTrack(val)
	}

	return retVal
}

func ParseGetTrackRequest(query url.Values) (domain.GetTrack, error) {
	var request domain.GetTrack

	if trackIDStr := query.Get("track_id"); trackIDStr != "" {
		trackID, err := strconv.Atoi(trackIDStr)
		if err != nil {
			return request, fmt.Errorf("invalid track_id: %w", err)
		}
		request.ID = &trackID
	}

	if source := query.Get("source"); source != "" {
		request.Source = &source
	}

	if target := query.Get("target"); target != "" {
		request.Target = &target
	}

	return request, nil
}
