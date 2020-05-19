package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/teploff/otus/scheduler/internal/domain/entity"
	"strconv"
	"strings"
)

// eventRepository implements EventRepository interface
type eventRepository struct {
	pgxPool *pgxpool.Pool
}

// NewEventRepository gets Event repository instance
func NewEventRepository(pgxPool *pgxpool.Pool) *eventRepository {
	return &eventRepository{pgxPool: pgxPool}
}

// GetEventsRequiringNotice provides to get notification which users must get from db
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
	WHERE (SELECT EXTRACT(MINUTE FROM (date - now()))) = remind_before AND is_received = FALSE`)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var event entity.Event
		if err = rows.Scan(&event.ID, &event.ShortDescription, &event.Date, &event.Duration, &event.FullDescription,
			&event.RemindBefore, &event.UserID); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return events, nil
}

// ConfirmEvents make confirmation events
func (e eventRepository) ConfirmEvents(ctx context.Context, events []entity.Event) error {
	tx, err := e.pgxPool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	ids := make([]interface{}, 0, len(events))
	placeHolders := make([]string, 0, len(events))
	for index, event := range events {
		ids = append(ids, event.ID)
		placeHolders = append(placeHolders, "$"+strconv.Itoa(index+1))
	}

	query := fmt.Sprintf("UPDATE\n"+
		"public.\"Event\"\n"+
		"SET\n"+
		"is_received = TRUE\n"+
		"WHERE id IN (%s)", strings.Join(placeHolders, ", "))

	_, err = tx.Exec(ctx, query, ids...)
	if err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

// CleanExpiredEvents provides expired events which older than 1 year
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
