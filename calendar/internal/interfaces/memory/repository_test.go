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
func TestEventRepositoryInsertOne(t *testing.T) {
	repo := NewEventRepository()
	firstEvent := model.NewEvent("first event")

	assert.NoError(t, repo.Insert(firstEvent))
	assert.NotEmpty(t, repo.events)
	assert.Equal(t, 1, len(repo.events))
}

// Test Case insert several events into empty repository
func TestEventRepositoryInsertSeveral(t *testing.T) {
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

//
//func TestEventRepositoryDeleteOne(t *testing.T) {
//	repo := NewEventRepository()
//
//	assert.Nil()
//}
//
//func TestEventRepositoryDeleteSeveral(t *testing.T) {
//	repo := NewEventRepository()
//
//	assert.Nil()
//}
//
//func TestEventRepositoryDeleteAbsentEvent(t *testing.T) {
//	repo := NewEventRepository()
//
//	assert.Nil()
//}
