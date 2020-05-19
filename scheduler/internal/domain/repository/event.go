package repository

import (
	"context"
	"github.com/teploff/otus/scheduler/internal/domain/entity"
)

// EventRepository provides query to database to get events to notify and clean expired events
type EventRepository interface {
	GetEventsRequiringNotice(context.Context) ([]entity.Event, error)
	ConfirmEvents(ctx context.Context, events []entity.Event) error
	CleanExpiredEvents(context.Context) error
}
