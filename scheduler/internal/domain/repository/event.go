package repository

import (
	"context"
	"github.com/teploff/otus/scheduler/internal/domain/entity"
)

type EventRepository interface {
	GetEventsRequiringNotice(context.Context) ([]entity.Event, error)
	CleanExpiredEvents(context.Context) error
}
