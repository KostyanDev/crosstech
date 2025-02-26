package storage

import (
	"github.com/jmoiron/sqlx"

	"github.com/sirupsen/logrus"
)

type Storage struct {
	logger *logrus.Logger
	ext    *sqlx.DB
}

func New(logger *logrus.Logger, ext *sqlx.DB) *Storage {
	return &Storage{
		logger: logger,
		ext:    ext,
	}
}
