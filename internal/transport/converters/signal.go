package converters

import (
	"app/internal/domain"
	"net/url"
	"strconv"
)

type CreatSignalRequest struct {
	ID      int     `json:"signal_id"`
	Name    string  `json:"signal_name" `
	ELR     string  `json:"elr"`
	Mileage float64 `json:"mileage" `
	TrackID int     `json:"track_id"`
}
type UpdateSignalRequest struct {
	ID        int      `json:"signal_id"`
	Name      *string  `json:"signal_name" `
	ELR       *string  `json:"elr"`
	Mileage   *float64 `json:"mileage" `
	TrackID   *int     `json:"track_id"`
	IsDeleted *bool    `json:"is_deleted"`
}
type SignalResponse struct {
	ID      int     `json:"signal_id"`
	Name    string  `json:"signal_name" `
	ELR     string  `json:"elr"`
	Mileage float64 `json:"mileage" `
	TrackID *int    `json:"track_id"`
}

func ToDomainCreatSignal(req CreatSignalRequest) domain.Signal {
	return domain.Signal{
		ID:      req.ID,
		Name:    req.Name,
		ELR:     req.ELR,
		Mileage: req.Mileage,
		TrackID: req.TrackID,
	}
}
func ToDomainUpdateSignal(req UpdateSignalRequest) domain.UpdateSignal {
	retVal := domain.UpdateSignal{
		ID: req.ID,
	}

	if req.Name != nil {
		retVal.Name = req.Name
	}
	if req.ELR != nil {
		retVal.ELR = req.ELR
	}
	if req.Mileage != nil {
		retVal.Mileage = req.Mileage
	}
	if req.TrackID != nil {
		retVal.TrackID = req.TrackID
	}
	if req.IsDeleted != nil {
		retVal.IsDeleted = req.IsDeleted
	}

	return retVal
}
func ToRespSignal(resp domain.Signal) SignalResponse {
	retVal := SignalResponse{
		ID:      resp.ID,
		Name:    resp.Name,
		ELR:     resp.ELR,
		Mileage: resp.Mileage,
	}
	if resp.TrackID != 0 {
		retVal.TrackID = &resp.TrackID
	}

	return retVal
}
func ToRespSignals(resp []domain.Signal) []SignalResponse {
	retVal := make([]SignalResponse, len(resp))
	for i, val := range resp {
		retVal[i] = ToRespSignal(val)
	}

	return retVal
}

func ParseGetSignalRequest(query url.Values) (domain.GetSignal, error) {
	var request domain.GetSignal

	if signalIDStr := query.Get("signal_id"); signalIDStr != "" {
		signalID, err := strconv.Atoi(signalIDStr)
		if err != nil {
			return request, err
		}
		request.ID = &signalID
	}

	if trackIDStr := query.Get("track_id"); trackIDStr != "" {
		trackID, err := strconv.Atoi(trackIDStr)
		if err != nil {
			return request, err
		}
		request.TrackID = &trackID
	}

	return request, nil
}
