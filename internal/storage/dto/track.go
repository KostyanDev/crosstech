package dto

import "app/internal/domain"

type TrackStorage struct {
	ID     int     `db:"id"`
	Source *string `db:"source"`
	Target *string `db:"target"`
}
type TracksStorage []TrackStorage

func (s TrackStorage) ToDomain() domain.Track {
	retVal := domain.Track{
		ID: s.ID,
	}
	if s.Source != nil {
		retVal.Source = *s.Source
	}
	if s.Target != nil {
		retVal.Target = *s.Target
	}

	return retVal
}

func (s TracksStorage) ToDomain() []domain.Track {
	signalsArr := make([]domain.Track, len(s))
	for i := range s {
		signalsArr[i] = s[i].ToDomain()
	}
	return signalsArr
}
