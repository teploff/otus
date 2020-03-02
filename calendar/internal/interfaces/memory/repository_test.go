package memory

import (
	"github.com/otus/calendar/internal/domain/model"
	"github.com/stretchr/testify/assert"
	"testing"
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
	firstEvent := model.NewEvent("first event")

	assert.NoError(t, repo.Insert(firstEvent))
	assert.NotEmpty(t, repo.events)
	assert.Equal(t, 1, len(repo.events))
}

// Test Case insert several events into empty repository
func TestEventRepositoryInsertSeveralIntoEmptySliceEvents(t *testing.T) {
	repo := NewEventRepository()
	firstEvent := model.NewEvent("first event")
	secondEvent := model.NewEvent("second event")
	thirdEvent := model.NewEvent("third event")

	assert.NoError(t, repo.Insert(firstEvent, secondEvent, thirdEvent))
	assert.NotEmpty(t, repo.events)
	assert.Equal(t, 3, len(repo.events))
}

// Test Case insert existed event into empty repository
func TestEventRepositoryInsertExistedEventIntoEmptySliceEvents(t *testing.T) {
	repo := NewEventRepository()
	firstEvent := model.NewEvent("first event")

	assert.Error(t, repo.Insert(firstEvent, firstEvent))
	assert.Empty(t, repo.events)
}

// Test Case insert existed event into not empty repository
func TestEventRepositoryInsertExistedEventIntoNotEmptySliceEvents(t *testing.T) {
	repo := NewEventRepository()
	firstEvent := model.NewEvent("first event")
	assert.NoError(t, repo.Insert(firstEvent))
	assert.NotEmpty(t, repo.events)
	assert.Equal(t, 1, len(repo.events))

	assert.Error(t, repo.Insert(firstEvent))
	assert.NotEmpty(t, repo.events)
	assert.Equal(t, 1, len(repo.events))
}

// Test Case delete one existed event from the repository
func TestEventRepositoryDeleteOneExistedEvent(t *testing.T) {
	repo := NewEventRepository()
	firstEvent := model.NewEvent("first event")
	secondEvent := model.NewEvent("second event")

	_ = repo.Insert(firstEvent, secondEvent)
	assert.NoError(t, repo.Delete(firstEvent))
	assert.NotEmpty(t, repo.events)
	assert.Equal(t, 1, len(repo.events))
	assert.Equal(t, "second event", repo.events[0].GetPayload())
}

// Test Case delete several existed events from the repository
func TestEventRepositoryDeleteSeveralExistedEvents(t *testing.T) {
	repo := NewEventRepository()
	firstEvent := model.NewEvent("first event")
	secondEvent := model.NewEvent("second event")
	thirdEvent := model.NewEvent("third event")

	_ = repo.Insert(firstEvent, secondEvent, thirdEvent)
	assert.NoError(t, repo.Delete(firstEvent, thirdEvent))
	assert.NotEmpty(t, repo.events)
	assert.Equal(t, 1, len(repo.events))
	assert.Equal(t, "second event", repo.events[0].GetPayload())
}

// Test Case delete one event from the empty repository
func TestEventRepositoryDeleteOneFromEmptyRepository(t *testing.T) {
	repo := NewEventRepository()
	firstEvent := model.NewEvent("first event")

	assert.Error(t, repo.Delete(firstEvent))
	assert.Empty(t, repo.events)
}

// Test Case delete several events from the empty repository
func TestEventRepositoryDeleteSeveralFromEmptyRepository(t *testing.T) {
	repo := NewEventRepository()
	firstEvent := model.NewEvent("first event")
	secondEvent := model.NewEvent("second event")
	thirdEvent := model.NewEvent("third event")

	assert.Error(t, repo.Delete(firstEvent, secondEvent, thirdEvent))
	assert.Empty(t, repo.events)
}

// Test Case delete one absent event from the repository
func TestEventRepositoryDeleteOneAbsentEventFromNotEmptyRepository(t *testing.T) {
	repo := NewEventRepository()
	firstEvent := model.NewEvent("first event")
	secondEvent := model.NewEvent("second event")
	thirdEvent := model.NewEvent("third event")

	_ = repo.Insert(firstEvent, secondEvent)
	assert.Error(t, repo.Delete(thirdEvent))
	assert.NotEmpty(t, repo.events)
	assert.Equal(t, 2, len(repo.events))
}

// Test Case delete several absent events from the repository
func TestEventRepositoryDeleteSeveralAbsentEventFromNotEmptyRepository(t *testing.T) {
	repo := NewEventRepository()
	firstEvent := model.NewEvent("first event")
	secondEvent := model.NewEvent("second event")
	thirdEvent := model.NewEvent("third event")

	_ = repo.Insert(secondEvent)
	assert.Error(t, repo.Delete(firstEvent, thirdEvent))
	assert.NotEmpty(t, repo.events)
	assert.Equal(t, 1, len(repo.events))
	assert.Equal(t, "second event", repo.events[0].GetPayload())
}
