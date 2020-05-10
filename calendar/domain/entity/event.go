package entity

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

// Event entity
type Event struct {
	id         uuid.UUID
	payload    string
	createTime time.Time
}

// NewEvent get Event instance
func NewEvent(payload string) Event {
	return Event{
		id:         uuid.NewV4(),
		payload:    payload,
		createTime: time.Now(),
	}
}

// GetID get event uuid
func (e Event) GetID() uuid.UUID {
	return e.id
}

// GetPayload get event payload
func (e Event) GetPayload() string {
	return e.payload
}

// GetCreateTime get event create time
func (e Event) GetCreateTime() time.Time {
	return e.createTime
}
