package service

import (
	"context"
	"github.com/teploff/otus/calendar/domain/entity"
	"time"
)

type CalendarService interface {
	CreateEvent(ctx context.Context, event entity.Event) error
	UpdateEvent(ctx context.Context, event entity.Event) error
	DeleteEvent(ctx context.Context, eventID, userID string) error
	GetDailyEvent(ctx context.Context, userID string, date time.Time) ([]entity.Event, error)
	GetWeeklyEvent(ctx context.Context, userID string, date time.Time) ([]entity.Event, error)
	GetMonthlyEvent(ctx context.Context, userID string, date time.Time) ([]entity.Event, error)
}
