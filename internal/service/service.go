package service

import (
	"app/internal/domain"
	"context"

	"github.com/sirupsen/logrus"
)

type Storage interface {
	CreateSignal(ctx context.Context, signal domain.Signal) error
	GetSignals(ctx context.Context, signal domain.GetSignal) ([]domain.Signal, error)
	UpdateSignal(ctx context.Context, signal domain.UpdateSignal) error
	SignalExists(ctx context.Context, signalID int) (bool, error)

	CreateTrack(ctx context.Context, track domain.Track) error
	GetTracks(ctx context.Context, track domain.GetTrack) ([]domain.Track, error)
	UpdateTrack(ctx context.Context, track domain.UpdateTrack) error
	TrackExists(ctx context.Context, trackID int) (bool, error)
}
type Service struct {
	ctx     context.Context
	log     *logrus.Logger
	storage Storage
}

func New(ctx context.Context, log *logrus.Logger, storage Storage) *Service {
	return &Service{
		ctx:     ctx,
		log:     log,
		storage: storage,
	}
}
