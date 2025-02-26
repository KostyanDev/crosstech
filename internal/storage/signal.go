package storage

import (
	"app/internal/domain"
	"app/internal/storage/dto"
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

func (s Storage) CreateSignal(ctx context.Context, signal domain.Signal) error {
	query, args, err := sq.Insert("signals").
		Columns(
			"id",
			"signal_name",
			"elr",
			"mileage",
			"track_id").
		Values(
			signal.ID,
			signal.Name,
			signal.ELR,
			signal.Mileage,
			signal.TrackID).
		Suffix("ON CONFLICT (id) DO NOTHING").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		s.logger.WithContext(ctx).WithError(err).Error("Failed to build insert signal query")
		return fmt.Errorf("failed to build insert signal query: %w", err)
	}

	_, err = s.ext.ExecContext(ctx, query, args...)
	if err != nil {
		s.logger.WithContext(ctx).WithError(err).Error("Failed to insert signal")
		return fmt.Errorf("failed to insert signal: %w", err)
	}

	return nil
}

func (s Storage) GetSignals(ctx context.Context, signal domain.GetSignal) ([]domain.Signal, error) {
	q := sq.Select(`
		id AS id,
		signal_name AS signal_name,
		elr AS elr,
		mileage AS mileage,
		track_id AS track_id,
		is_deleted AS is_deleted
	`).From("signals").
		PlaceholderFormat(sq.Dollar)

	q = generateGetSignal(q, signal)

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

	var signals []dto.SignalStorage
	for rows.Next() {
		var signalStrg dto.SignalStorage
		if err = rows.StructScan(&signalStrg); err != nil {
			s.logger.WithError(err).Error("Failed to scan row")
			return nil, err
		}
		signals = append(signals, signalStrg)
	}

	if err := rows.Err(); err != nil {
		s.logger.WithError(err).Error("Error iterating over rows")
		return nil, err
	}

	return dto.SignalsStorage(signals).ToDomain(), nil
}
func (s Storage) UpdateSignal(ctx context.Context, signal domain.UpdateSignal) error {
	builder := sq.Update("signals").
		Where(sq.Eq{"id": signal.ID})

	updateCount := 0
	if signal.Name != nil {
		builder = builder.Set("signal_name", *signal.Name)
		updateCount++
	}
	if signal.ELR != nil {
		builder = builder.Set("elr", *signal.ELR)
		updateCount++
	}
	if signal.Mileage != nil {
		builder = builder.Set("mileage", *signal.Mileage)
		updateCount++
	}
	if signal.TrackID != nil {
		builder = builder.Set("track_id", *signal.TrackID)
		updateCount++
	}
	if signal.IsDeleted != nil {
		builder = builder.Set("is_deleted", *signal.IsDeleted)
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

func (s Storage) SignalExists(ctx context.Context, signalID int) (bool, error) {
	q := sq.Select("COUNT(*)").
		From("signals").
		Where(sq.Eq{"id": signalID}).
		PlaceholderFormat(sq.Dollar)

	toSQL, args, err := q.ToSql()
	if err != nil {
		s.logger.WithError(err).Error("Failed to generate SQL query for SignalExists")
		return false, err
	}

	var count int
	err = s.ext.QueryRowxContext(ctx, toSQL, args...).Scan(&count)
	if err != nil {
		s.logger.WithError(err).Error("Failed to execute SignalExists query")
		return false, err
	}

	return count > 0, nil
}

func generateGetSignal(q sq.SelectBuilder, data domain.GetSignal) sq.SelectBuilder {
	q = q.Where(sq.Eq{"is_deleted": false})

	if data.ID != nil {
		q = q.Where(sq.Eq{"id": data.ID})
	}
	if data.TrackID != nil {
		q = q.Where(sq.Eq{"track_id": data.TrackID})
	}
	return q
}
