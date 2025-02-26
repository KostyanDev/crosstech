package service

import (
	"app/internal/domain"
	"context"
	"errors"
	"fmt"
)

func (s Service) CreateSignal(ctx context.Context, request domain.Signal) error {
	isExist, err := s.storage.SignalExists(ctx, request.ID)
	if err != nil {
		return err
	}

	if isExist {
		return errors.New(fmt.Sprintf("signal with id %s created", request.ID))
	}

	return s.storage.CreateSignal(ctx, request)
}

func (s Service) GetSignalByParam(ctx context.Context, request domain.GetSignal) ([]domain.Signal, error) {
	return s.storage.GetSignals(ctx, request)
}

func (s Service) UpdateSignalByParam(ctx context.Context, request domain.UpdateSignal) error {
	isExist, err := s.storage.SignalExists(ctx, request.ID)
	if err != nil {
		return err
	}

	if !isExist {
		return errors.New(fmt.Sprintf("signal with id %s not found", request.ID))
	}

	return s.storage.UpdateSignal(ctx, request)
}
