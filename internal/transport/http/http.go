package http

import (
	"app/internal/domain"
	"context"

	"app/internal/service"

	"github.com/sirupsen/logrus"
)

type Service interface {
	ProcessJSONFile(ctx context.Context, file domain.File) error

	CreateSignal(ctx context.Context, signal domain.Signal) error
	GetSignalByParam(ctx context.Context, request domain.GetSignal) ([]domain.Signal, error)
	UpdateSignalByParam(ctx context.Context, request domain.UpdateSignal) error

	CreateTrack(ctx context.Context, request domain.Track) error
	GetTrackByParam(ctx context.Context, request domain.GetTrack) ([]domain.Track, error)
	UpdateTrackByParam(ctx context.Context, request domain.UpdateTrack) error
}

type Handler struct {
	ctx     context.Context
	log     *logrus.Logger
	service Service
}

func New(ctx context.Context, log *logrus.Logger, service *service.Service) *Handler {
	return &Handler{
		ctx:     ctx,
		log:     log,
		service: service,
	}
}
