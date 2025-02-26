package storage

import (
	"app/internal/domain"
	"app/internal/storage/dto"
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

func (s Storage) CreateTrack(ctx context.Context, track domain.Track) error {
	query, args, err := sq.Insert("tracks").
		Columns(
			"id",
			"source",
			"target").
		Values(
			track.ID,
			track.Source,
			track.Target).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		s.logger.WithContext(ctx).
			WithError(err).
			Error("failed to build insert track query")
		return err
	}

	_, err = s.ext.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.WithContext(ctx).
			WithError(err).
			Error("failed to insert track")
		return err
	}

	return nil
}

func (s Storage) GetTracks(ctx context.Context, track domain.GetTrack) ([]domain.Track, error) {
	q := sq.Select(`
		id AS id,
		source AS source,
		target AS target
	`).From("tracks").
		PlaceholderFormat(sq.Dollar)
	q = generateGetTrack(q, track)

	toSQL, args, err := q.ToSql()
	if err != nil {
		s.logger.WithError(err).Error("Failed to generate SQL query")
		return nil, err
	}

	rows, err := s.ext.QueryxContext(ctx, toSQL, args...)
	if err != nil {
		s.logger.WithError(err).Error("Failed to execute query")
		return nil, err
	}
	defer rows.Close()

	var tracks []dto.TrackStorage
	for rows.Next() {
		var trackStrg dto.TrackStorage
		if err = rows.StructScan(&trackStrg); err != nil {
			s.logger.WithError(err).Error("Failed to scan row")
			return nil, err
		}
		tracks = append(tracks, trackStrg)
	}

	if err := rows.Err(); err != nil {
		s.logger.WithError(err).Error("Error iterating over rows")
		return nil, err
	}

	return dto.TracksStorage(tracks).ToDomain(), nil
}

func (s Storage) UpdateTrack(ctx context.Context, track domain.UpdateTrack) error {
	builder := sq.Update("tracks").
		Where(sq.Eq{"id": track.ID})

	updateCount := 0
	if track.Source != nil {
		builder = builder.Set("source", *track.Source)
		updateCount++
	}
	if track.Target != nil {
		builder = builder.Set("target", *track.Target)
		updateCount++
	}
	if track.IsDeleted != nil {
		builder = builder.Set("is_deleted", *track.IsDeleted)
		updateCount++
	}

	if updateCount == 0 {
		return fmt.Errorf("no fields to update")
	}

	query, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = s.ext.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s Storage) TrackExists(ctx context.Context, trackID int) (bool, error) {
	q := sq.Select("COUNT(*)").
		From("tracks").
		Where(sq.Eq{"id": trackID}).
		PlaceholderFormat(sq.Dollar)

	toSQL, args, err := q.ToSql()
	if err != nil {
		s.logger.WithError(err).Error("Failed to generate SQL query for TrackExists")
		return false, err
	}

	var count int
	err = s.ext.QueryRowxContext(ctx, toSQL, args...).Scan(&count)
	if err != nil {
		s.logger.WithError(err).Error("Failed to execute TrackExists query")
		return false, err
	}

	return count > 0, nil
}

func generateGetTrack(q sq.SelectBuilder, data domain.GetTrack) sq.SelectBuilder {
	q = q.Where(sq.Eq{"is_deleted": false})

	if data.ID != nil {
		q = q.Where(sq.Eq{"id": data.ID})
	}
	if data.Source != nil {
		q = q.Where(sq.Eq{"source": data.Source})
	}
	if data.Target != nil {
		q = q.Where(sq.Eq{"target": data.Target})
	}
	return q
}
