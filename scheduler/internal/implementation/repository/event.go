package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/teploff/otus/scheduler/internal/domain/entity"
)

type eventRepository struct {
	pgxPool *pgxpool.Pool
}

func NewEventRepository(pgxPool *pgxpool.Pool) *eventRepository {
	return &eventRepository{pgxPool: pgxPool}
}

func (e eventRepository) GetEventsRequiringNotice(ctx context.Context) ([]entity.Event, error) {
	tx, err := e.pgxPool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	events := make([]entity.Event, 0, 100)
	rows, err := tx.Query(ctx, `
	SELECT id, short_description, date, duration, full_description, remind_before, user_id 
	FROM public."Event" 
	WHERE (SELECT EXTRACT(MINUTE FROM (date - now()))) = remind_before`)
	defer rows.Close()

	for rows.Next() {
		var event entity.Event
		if err = rows.Scan(&event.ID, &event.ShortDescription, &event.Date, &event.Duration, &event.FullDescription,
			&event.RemindBefore, &event.UserID); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return events, nil
}

func (e eventRepository) CleanExpiredEvents(ctx context.Context) error {
	tx, err := e.pgxPool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `
	DELETE
	FROM public."Event"
	WHERE (SELECT EXTRACT(YEAR FROM (now()))) - (SELECT EXTRACT(YEAR FROM date)) >= 1`)
	if err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
