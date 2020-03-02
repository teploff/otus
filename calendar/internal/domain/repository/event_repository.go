package repository

import (
	"github.com/otus/calendar/internal/domain/model"
	uuid "github.com/satori/go.uuid"
	"time"
)

// EventRepository interface to work with Event Repository: in memory, db ect.
type EventRepository interface {
	Insert(events ...model.Event) error
	Delete(events ...model.Event) error
	Update(srcEventIDs []uuid.UUID, destEvents []model.Event) error
	GetByID(id uuid.UUID) (model.Event, error)
	GetByPayload(payload string) ([]model.Event, error)
	GetEarlierThanTime(t time.Time) ([]model.Event, error)
	GetLaterThanTime(t time.Time) ([]model.Event, error)
	GetAll() ([]model.Event, error)
}
