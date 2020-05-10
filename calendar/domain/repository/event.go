package repository

import (
	"context"
	"github.com/teploff/otus/calendar/domain/entity"
	"time"
)

// EventRepository interface to work with Event Repository: in memory, db ect.
type EventRepository interface {
	InsertEvent(ctx context.Context, event entity.Event) error
	UpdateEvent(ctx context.Context, event entity.Event) error
	DeleteEvent(ctx context.Context, eventID string, userID string) error
	GetEvents(ctx context.Context, userID string, startDate time.Time, duration time.Duration) ([]entity.Event, error)
}
