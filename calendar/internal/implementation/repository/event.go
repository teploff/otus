package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/teploff/otus/calendar/domain/entity"
	"time"
)

type eventRepository struct {
	pgxPool *pgxpool.Pool
}

func NewEventRepository(pgxPool *pgxpool.Pool) *eventRepository {
	return &eventRepository{pgxPool: pgxPool}
}

func (e eventRepository) InsertEvent(ctx context.Context, event entity.Event) error {
	tx, err := e.pgxPool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	_, err = tx.Exec(ctx, `
	INSERT
		INTO public."Event"(short_description, date, duration, full_description, remind_before, user_id)
	VALUES
		($1, $2, $3, $4, $5, $6, $7)`, event.ShortDescription, event.Date, event.Duration, event.FullDescription,
		event.RemindBefore, event.UserID)
	if err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (e eventRepository) UpdateEvent(ctx context.Context, event entity.Event) error {
	tx, err := e.pgxPool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `
	UPDATE public."Event"
		SET short_description = $3, date = $4, duration = $5, full_description = $6, remind_before = $7
	WHERE id = $1 AND user_id = $2;`, event.ID, event.UserID, event.ShortDescription, event.Date, event.Duration,
		event.FullDescription, event.RemindBefore)
	if err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (e eventRepository) DeleteEvent(ctx context.Context, eventID string, userID string) error {
	tx, err := e.pgxPool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `
	DELETE
	FROM public."Event"
	WHERE id = $1 AND user_id = $2`, eventID, userID)
	if err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (e eventRepository) GetEvents(ctx context.Context, userID string, startDate time.Time, duration time.Duration) ([]entity.Event, error) {
	tx, err := e.pgxPool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	events := make([]entity.Event, 0, 100)
	finishDate := startDate.Add(duration)
	rows, err := tx.Query(ctx, `
	SELECT id, short_description, date, duration, full_description, remind_before, user_id 
	FROM public."Event"
	WHERE user_id = $1 AND date >= $2 AND date < $3`, userID, startDate, finishDate)
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
