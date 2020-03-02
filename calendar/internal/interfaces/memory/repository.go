package memory

import (
	"fmt"
	"github.com/otus/calendar/internal/domain/model"
	uuid "github.com/satori/go.uuid"
	"sort"
	"sync"
	"time"
)

// EventRepository in memory repository
type EventRepository struct {
	mu     *sync.Mutex
	events []model.Event
}

// NewEventRepository get EventRepository instance
func NewEventRepository() *EventRepository {
	return &EventRepository{
		mu:     &sync.Mutex{},
		events: make([]model.Event, 0),
	}
}

// Insert creates new events in the repository
func (r *EventRepository) Insert(events ...model.Event) error {
	id, unique := r.checkForUniqueness(events...)
	if !unique {
		return fmt.Errorf("event with id = %s already exist", id)
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	for _, insertedEvent := range events {
		for _, existedEvent := range r.events {
			if insertedEvent == existedEvent {
				return fmt.Errorf("event with id = %s already exist", insertedEvent.GetID())
			}
		}
	}

	// transaction insert
	r.events = append(r.events, events...)
	return nil
}

// checkForUniqueness finds out not unique event in events slice
func (r *EventRepository) checkForUniqueness(events ...model.Event) (uuid.UUID, bool) {
	for i := 0; i < len(events)-1; i++ {
		for j := i + 1; j < len(events); j++ {
			if events[i].GetID() == events[j].GetID() {
				return events[i].GetID(), false
			}
		}
	}

	return uuid.UUID{}, true
}

// Delete removes events from the repository
func (r *EventRepository) Delete(events ...model.Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	indexesMustDelete := make([]int, 0, len(events))
	for _, event := range events {
		eventExist := false
		for index, existedEvent := range r.events {
			if event == existedEvent {
				indexesMustDelete = append(indexesMustDelete, index)
				eventExist = true
				break
			}
		}
		if !eventExist {
			return fmt.Errorf("event with id = %s, payload = %s and time = %s doesn't exist",
				event.GetID(), event.GetPayload(), event.GetCreateTime())
		}
	}

	// transaction delete
	sort.Ints(indexesMustDelete)
	currEvents := make([]model.Event, 0, len(r.events)-len(indexesMustDelete))
	for index, event := range events {
		for _, value := range indexesMustDelete {
			if index == value {
				break
			}
			currEvents = append(currEvents, event)
		}
	}
	r.events = currEvents
	return nil
}

// Update updates existed events with srcEventIDs on destEvents events
func (r *EventRepository) Update(srcEventIDs []uuid.UUID, destEvents []model.Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	indexesMustUpdate := make([]int, 0, len(destEvents))
	for destIndex, id := range srcEventIDs {
		eventExist := false
		for srcIndex, event := range r.events {
			if event.GetID() == id {
				indexesMustUpdate = append(indexesMustUpdate, srcIndex)
				eventExist = true
				break
			}
		}
		if !eventExist {
			event := destEvents[destIndex]
			return fmt.Errorf("event with id = %s, payload = %s and time = %s doesn't exist",
				event.GetID(), event.GetPayload(), event.GetCreateTime())
		}
	}

	// transaction update
	for index, value := range indexesMustUpdate {
		r.events[value] = destEvents[index]
	}

	return nil
}

// GetByID get event by id
func (r *EventRepository) GetByID(id uuid.UUID) (model.Event, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, event := range r.events {
		if event.GetID() == id {
			return event, nil
		}
	}

	return model.Event{}, fmt.Errorf("event with id = %s doesn't exist", id)
}

// GetByPayload get events by payload
func (r *EventRepository) GetByPayload(payload string) ([]model.Event, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := make([]model.Event, 0, len(r.events))
	for _, event := range r.events {
		if event.GetPayload() == payload {
			result = append(result, event)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("events with payload = %s don't exist", payload)
	}
	return result, nil
}

// GetEarlierThanTime get events which was created before than time t
func (r *EventRepository) GetEarlierThanTime(t time.Time) ([]model.Event, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := make([]model.Event, 0, len(r.events))
	for _, event := range r.events {
		if event.GetCreateTime().Before(t) {
			result = append(result, event)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("events early than time = %s don't exist", t.String())
	}
	return result, nil
}

// GetLaterThanTime get events which was created after than time t
func (r *EventRepository) GetLaterThanTime(t time.Time) ([]model.Event, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := make([]model.Event, 0, len(r.events))
	for _, event := range r.events {
		if event.GetCreateTime().After(t) {
			result = append(result, event)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("events later than time = %s don't exist", t.String())
	}
	return result, nil
}

// GetAll get all events
func (r *EventRepository) GetAll() ([]model.Event, error) {
	if len(r.events) == 0 {
		return nil, fmt.Errorf("events don't exist")
	}
	return r.events, nil
}
