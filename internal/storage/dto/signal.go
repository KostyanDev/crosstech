package dto

import "app/internal/domain"

type SignalStorage struct {
	ID         int     `db:"id"`
	SignalName string  `db:"signal_name"`
	ELR        string  `db:"elr"`
	Mileage    float64 `db:"mileage"`
	TrackID    *int    `db:"track_id"`
	IsDeleted  bool    `db:"is_deleted"`
}
type SignalsStorage []SignalStorage

func (s SignalStorage) ToDomain() domain.Signal {
	retVal := domain.Signal{
		ID:        s.ID,
		Name:      s.SignalName,
		ELR:       s.ELR,
		Mileage:   s.Mileage,
		IsDeleted: s.IsDeleted,
	}
	if s.TrackID != nil {
		retVal.TrackID = *s.TrackID
	}
	return retVal
}

func (s SignalsStorage) ToDomain() []domain.Signal {
	signalsArr := make([]domain.Signal, len(s))
	for i := range s {
		signalsArr[i] = s[i].ToDomain()
	}
	return signalsArr
}
