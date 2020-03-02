package memory

import (
	"github.com/otus/calendar/internal/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//testData := []struct {
//	in  string
//	out string
//}{
//	{"abed", "abed"},
//	{"a4bc2d5e", "aaaabccddddde"},
//}
//
//for _, data := range testData {
//	converter := NewStringConverter(data.in)
//	out, err := converter.Do()
//	assert.Nil(t, err)
//	assert.Equal(t, out, data.out)
//}

// Test Case insert one event into empty repository
func TestEventRepositoryInsertOneIntoEmptySliceEvents(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")

	assert.NoError(t, repo.Insert(event1))
	assert.NotEmpty(t, repo.events)
	assert.Equal(t, 1, len(repo.events))
	assert.Equal(t, event1.GetID(), repo.events[0].GetID())
	assert.Equal(t, event1.GetPayload(), repo.events[0].GetPayload())
	assert.Equal(t, event1.GetCreateTime(), repo.events[0].GetCreateTime())
}

// Test Case insert several events into empty repository
func TestEventRepositoryInsertSeveralIntoEmptySliceEvents(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")
	event2 := model.NewEvent("event 2")
	event3 := model.NewEvent("event 3")

	assert.NoError(t, repo.Insert(event1, event2, event3))
	assert.NotEmpty(t, repo.events)
	assert.Equal(t, 3, len(repo.events))
	assert.Equal(t, event1.GetID(), repo.events[0].GetID())
	assert.Equal(t, event1.GetPayload(), repo.events[0].GetPayload())
	assert.Equal(t, event1.GetCreateTime(), repo.events[0].GetCreateTime())
	assert.Equal(t, event2.GetID(), repo.events[1].GetID())
	assert.Equal(t, event2.GetPayload(), repo.events[1].GetPayload())
	assert.Equal(t, event2.GetCreateTime(), repo.events[1].GetCreateTime())
	assert.Equal(t, event3.GetID(), repo.events[2].GetID())
	assert.Equal(t, event3.GetPayload(), repo.events[2].GetPayload())
	assert.Equal(t, event3.GetCreateTime(), repo.events[2].GetCreateTime())
}

// Test Case insert existed event into empty repository
func TestEventRepositoryInsertExistedEventIntoEmptySliceEvents(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")

	assert.Error(t, repo.Insert(event1, event1))
	assert.Empty(t, repo.events)
}

// Test Case insert existed event into not empty repository
func TestEventRepositoryInsertExistedEventIntoNotEmptySliceEvents(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")
	assert.NoError(t, repo.Insert(event1))
	assert.Error(t, repo.Insert(event1))
	assert.NotEmpty(t, repo.events)
	assert.Equal(t, 1, len(repo.events))
	assert.Equal(t, event1.GetID(), repo.events[0].GetID())
	assert.Equal(t, event1.GetPayload(), repo.events[0].GetPayload())
	assert.Equal(t, event1.GetCreateTime(), repo.events[0].GetCreateTime())
}

// Test Case delete one existed event from the repository
func TestEventRepositoryDeleteOneExistedEvent(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")
	event2 := model.NewEvent("event 2")

	_ = repo.Insert(event1, event2)
	assert.NoError(t, repo.Delete(event1))
	assert.NotEmpty(t, repo.events)
	assert.Equal(t, 1, len(repo.events))
	assert.Equal(t, event2.GetID(), repo.events[0].GetID())
	assert.Equal(t, event2.GetPayload(), repo.events[0].GetPayload())
	assert.Equal(t, event2.GetCreateTime(), repo.events[0].GetCreateTime())
}

// Test Case delete several existed events from the repository
func TestEventRepositoryDeleteSeveralExistedEvents(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")
	event2 := model.NewEvent("event 2")
	event3 := model.NewEvent("event 3")

	_ = repo.Insert(event1, event2, event3)
	assert.NoError(t, repo.Delete(event1, event3))
	assert.NotEmpty(t, repo.events)
	assert.Equal(t, 1, len(repo.events))
	assert.Equal(t, event2.GetID(), repo.events[0].GetID())
	assert.Equal(t, event2.GetPayload(), repo.events[0].GetPayload())
	assert.Equal(t, event2.GetCreateTime(), repo.events[0].GetCreateTime())
}

// Test Case delete one event from the empty repository
func TestEventRepositoryDeleteOneFromEmptyRepository(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("first event")

	assert.Error(t, repo.Delete(event1))
	assert.Empty(t, repo.events)
}

// Test Case delete several events from the empty repository
func TestEventRepositoryDeleteSeveralFromEmptyRepository(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")
	event2 := model.NewEvent("event 2")
	event3 := model.NewEvent("event 3")

	assert.Error(t, repo.Delete(event1, event2, event3))
	assert.Empty(t, repo.events)
}

// Test Case delete one absent event from the repository
func TestEventRepositoryDeleteOneAbsentEventFromNotEmptyRepository(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")
	event2 := model.NewEvent("event 2")
	event3 := model.NewEvent("event 3")

	_ = repo.Insert(event1, event2)
	assert.Error(t, repo.Delete(event3))
	assert.NotEmpty(t, repo.events)
	assert.Equal(t, 2, len(repo.events))
	assert.Equal(t, event1.GetID(), repo.events[0].GetID())
	assert.Equal(t, event1.GetPayload(), repo.events[0].GetPayload())
	assert.Equal(t, event1.GetCreateTime(), repo.events[0].GetCreateTime())
	assert.Equal(t, event2.GetID(), repo.events[1].GetID())
	assert.Equal(t, event2.GetPayload(), repo.events[1].GetPayload())
	assert.Equal(t, event2.GetCreateTime(), repo.events[1].GetCreateTime())
}

// Test Case delete several absent events from the repository
func TestEventRepositoryDeleteSeveralAbsentEventFromNotEmptyRepository(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")
	event2 := model.NewEvent("event 2")
	event3 := model.NewEvent("event 3")

	_ = repo.Insert(event2)
	assert.Error(t, repo.Delete(event1, event3))
	assert.NotEmpty(t, repo.events)
	assert.Equal(t, 1, len(repo.events))
	assert.Equal(t, event2.GetID(), repo.events[0].GetID())
	assert.Equal(t, event2.GetPayload(), repo.events[0].GetPayload())
	assert.Equal(t, event2.GetCreateTime(), repo.events[0].GetCreateTime())
}

// Test Case update one existed event from the repository
func TestEventRepositoryUpdateOneExistedEvent(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")
	event2 := model.NewEvent("event 2")
	newEvent := model.NewEvent("new event")

	_ = repo.Insert(event1, event2)
	assert.NoError(t, repo.Update([]uuid.UUID{event1.GetID()}, []model.Event{newEvent}))
	assert.NotEmpty(t, repo.events)
	assert.Equal(t, 2, len(repo.events))
	assert.Equal(t, newEvent.GetID(), repo.events[0].GetID())
	assert.Equal(t, newEvent.GetPayload(), repo.events[0].GetPayload())
	assert.Equal(t, newEvent.GetCreateTime(), repo.events[0].GetCreateTime())
	assert.Equal(t, event2.GetID(), repo.events[1].GetID())
	assert.Equal(t, event2.GetPayload(), repo.events[1].GetPayload())
	assert.Equal(t, event2.GetCreateTime(), repo.events[1].GetCreateTime())
}

// Test Case update several existed events from the repository
func TestEventRepositoryUpdateSeveralExistedEvents(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")
	event2 := model.NewEvent("event 2")
	newEvent1 := model.NewEvent("new event 1")
	newEvent2 := model.NewEvent("new event 2")

	_ = repo.Insert(event1, event2)
	assert.NoError(t, repo.Update([]uuid.UUID{event1.GetID(), event2.GetID()}, []model.Event{newEvent1,
		newEvent2}))
	assert.NotEmpty(t, repo.events)
	assert.Equal(t, 2, len(repo.events))
	assert.Equal(t, newEvent1.GetID(), repo.events[0].GetID())
	assert.Equal(t, newEvent1.GetPayload(), repo.events[0].GetPayload())
	assert.Equal(t, newEvent1.GetCreateTime(), repo.events[0].GetCreateTime())
	assert.Equal(t, newEvent2.GetID(), repo.events[1].GetID())
	assert.Equal(t, newEvent2.GetPayload(), repo.events[1].GetPayload())
	assert.Equal(t, newEvent2.GetCreateTime(), repo.events[1].GetCreateTime())
}

// Test Case update one event from the empty repository
func TestEventRepositoryUpdateOneFromEmptyRepository(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")
	newEvent := model.NewEvent("new event")

	assert.Error(t, repo.Update([]uuid.UUID{event1.GetID()}, []model.Event{newEvent}))
	assert.Empty(t, repo.events)
}

// Test Case update several events from the empty repository
func TestEventRepositoryUpdateSeveralFromEmptyRepository(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")
	event2 := model.NewEvent("event 2")
	newEvent1 := model.NewEvent("new event 1")
	newEvent2 := model.NewEvent("new event 2")

	assert.Error(t, repo.Update([]uuid.UUID{event1.GetID(), event2.GetID()}, []model.Event{newEvent1, newEvent2}))
	assert.Empty(t, repo.events)
}

// Test Case update one absent event from the repository
func TestEventRepositoryUpdateOneAbsentEventFromNotEmptyRepository(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")
	event2 := model.NewEvent("event 2")
	newEvent1 := model.NewEvent("new event 1")

	_ = repo.Insert(event1)
	assert.Error(t, repo.Update([]uuid.UUID{event2.GetID()}, []model.Event{newEvent1}))
	assert.NotEmpty(t, repo.events)
	assert.Equal(t, 1, len(repo.events))
	assert.Equal(t, event1.GetID(), repo.events[0].GetID())
	assert.Equal(t, event1.GetPayload(), repo.events[0].GetPayload())
	assert.Equal(t, event1.GetCreateTime(), repo.events[0].GetCreateTime())

}

// Test Case update several absent events from the repository
func TestEventRepositoryUpdateSeveralAbsentEventFromNotEmptyRepository(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")
	event2 := model.NewEvent("event 2")
	event3 := model.NewEvent("event 3")
	newEvent1 := model.NewEvent("new event 1")
	newEvent2 := model.NewEvent("new event 2")

	_ = repo.Insert(event1)
	assert.Error(t, repo.Update([]uuid.UUID{event2.GetID(), event3.GetID()}, []model.Event{newEvent1, newEvent2}))
	assert.NotEmpty(t, repo.events)
	assert.Equal(t, 1, len(repo.events))
	assert.Equal(t, event1.GetID(), repo.events[0].GetID())
	assert.Equal(t, event1.GetPayload(), repo.events[0].GetPayload())
	assert.Equal(t, event1.GetCreateTime(), repo.events[0].GetCreateTime())
}

// Test Case getting existed event by id from the repository
func TestEventRepositoryGetExistedEventByID(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")
	event2 := model.NewEvent("event 2")
	event3 := model.NewEvent("event 3")

	_ = repo.Insert(event1, event2, event3)
	event, err := repo.GetByID(event3.GetID())
	assert.NoError(t, err)
	assert.Equal(t, event3.GetID(), event.GetID())
	assert.Equal(t, event3.GetPayload(), event.GetPayload())
	assert.Equal(t, event3.GetCreateTime(), event.GetCreateTime())
}

// Test Case getting event by id from the empty repository
func TestEventRepositoryGetEventByIDFromEmptyRepository(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")

	event, err := repo.GetByID(event1.GetID())
	assert.Error(t, err)
	assert.Empty(t, repo.events)
	assert.Empty(t, event)
}

// Test Case getting absent event by id from the repository
func TestEventRepositoryGetAbsentEventByID(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")
	event2 := model.NewEvent("event 2")
	event3 := model.NewEvent("event 3")

	_ = repo.Insert(event1, event3)
	event, err := repo.GetByID(event2.GetID())
	assert.Error(t, err)
	assert.Empty(t, event)
}

// Test Case getting one existed event by payload from the repository
func TestEventRepositoryGetOneExistedEventByPayload(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")
	event2 := model.NewEvent("event 2")
	event3 := model.NewEvent("event 3")

	_ = repo.Insert(event1, event2, event3)
	events, err := repo.GetByPayload(event3.GetPayload())
	assert.NoError(t, err)
	assert.Equal(t, 1, len(events))
	assert.Equal(t, event3.GetID(), events[0].GetID())
	assert.Equal(t, event3.GetPayload(), events[0].GetPayload())
	assert.Equal(t, event3.GetCreateTime(), events[0].GetCreateTime())
}

// Test Case getting several existed event by payload from the repository
func TestEventRepositoryGetSeveralExistedEventByPayload(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")
	event2 := model.NewEvent("event 2")
	event3 := model.NewEvent("event 1")

	_ = repo.Insert(event1, event2, event3)
	events, err := repo.GetByPayload(event3.GetPayload())
	assert.NoError(t, err)
	assert.Equal(t, 2, len(events))
	assert.Equal(t, event1.GetID(), events[0].GetID())
	assert.Equal(t, event1.GetPayload(), events[0].GetPayload())
	assert.Equal(t, event1.GetCreateTime(), events[0].GetCreateTime())
	assert.Equal(t, event3.GetID(), events[1].GetID())
	assert.Equal(t, event3.GetPayload(), events[1].GetPayload())
	assert.Equal(t, event3.GetCreateTime(), events[1].GetCreateTime())
}

// Test Case getting event by payload from the empty repository
func TestEventRepositoryGetEventByPayloadFromEmptyRepository(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")

	events, err := repo.GetByPayload(event1.GetPayload())
	assert.Error(t, err)
	assert.Empty(t, repo.events)
	assert.Empty(t, events)
}

// Test Case getting absent event by payload from the repository
func TestEventRepositoryGetAbsentEventByPayload(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")
	event2 := model.NewEvent("event 2")
	event3 := model.NewEvent("event 3")

	_ = repo.Insert(event1, event3)
	event, err := repo.GetByPayload(event2.GetPayload())
	assert.Error(t, err)
	assert.Empty(t, event)
}

// Test Case getting one event early than time from the repository
func TestEventRepositoryGetOneEventEarlierThanTime(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")

	_ = repo.Insert(event1)
	events, err := repo.GetEarlierThanTime(time.Now())
	assert.NoError(t, err)
	assert.Equal(t, 1, len(events))
	assert.Equal(t, event1.GetID(), events[0].GetID())
	assert.Equal(t, event1.GetPayload(), events[0].GetPayload())
	assert.Equal(t, event1.GetCreateTime(), events[0].GetCreateTime())
}

// Test Case getting several events early than time from the repository
func TestEventRepositoryGetSeveralEventsEarlierThanTime(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")
	event2 := model.NewEvent("event 2")
	timeNow := time.Now()
	event3 := model.NewEvent("event 3")

	_ = repo.Insert(event1, event2, event3)
	events, err := repo.GetEarlierThanTime(timeNow)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(events))
	assert.Equal(t, event1.GetID(), events[0].GetID())
	assert.Equal(t, event1.GetPayload(), events[0].GetPayload())
	assert.Equal(t, event1.GetCreateTime(), events[0].GetCreateTime())
	assert.Equal(t, event2.GetID(), events[1].GetID())
	assert.Equal(t, event2.GetPayload(), events[1].GetPayload())
	assert.Equal(t, event2.GetCreateTime(), events[1].GetCreateTime())
}

// Test Case getting all events early than time from the repository
func TestEventRepositoryGetAllEventsEarlierThanTime(t *testing.T) {
	repo := NewEventRepository()
	event1 := model.NewEvent("event 1")
	event2 := model.NewEvent("event 2")
	event3 := model.NewEvent("event 3")
	timeNow := time.Now()

	_ = repo.Insert(event1, event2, event3)
	events, err := repo.GetEarlierThanTime(timeNow)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(events))
	assert.Equal(t, event1.GetID(), events[0].GetID())
	assert.Equal(t, event1.GetPayload(), events[0].GetPayload())
	assert.Equal(t, event1.GetCreateTime(), events[0].GetCreateTime())
	assert.Equal(t, event2.GetID(), events[1].GetID())
	assert.Equal(t, event2.GetPayload(), events[1].GetPayload())
	assert.Equal(t, event2.GetCreateTime(), events[1].GetCreateTime())
	assert.Equal(t, event3.GetID(), events[2].GetID())
	assert.Equal(t, event3.GetPayload(), events[2].GetPayload())
	assert.Equal(t, event3.GetCreateTime(), events[2].GetCreateTime())
}

// Test Case getting events early than time from the empty repository
func TestEventRepositoryGetEventsEarlierThanTimeFromEmptyRepository(t *testing.T) {
	repo := NewEventRepository()

	events, err := repo.GetEarlierThanTime(time.Now())
	assert.Error(t, err)
	assert.Empty(t, events)
	assert.Empty(t, repo.events)
}

// Test Case getting events early than time from repository which consists all events after time t
func TestEventRepositoryGetEventsEarlierThanTimeFromTheAfterEvents(t *testing.T) {
	repo := NewEventRepository()
	timeNow := time.Now()
	event1 := model.NewEvent("event 1")
	event2 := model.NewEvent("event 2")
	event3 := model.NewEvent("event 3")

	_ = repo.Insert(event1, event2, event3)
	events, err := repo.GetEarlierThanTime(timeNow)
	assert.Error(t, err)
	assert.Empty(t, events)
}
