package hw_4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Pushing front

// TestCase
func TestPushFrontInEmptyList(t *testing.T) {
	l := List{}

	l.PushFront(1)

	assert.Equal(t, 1, l.Len())

	assert.NotNil(t, l.First())
	assert.Equal(t, l.First().Value(), 1)
	assert.Nil(t, l.First().prev)
	assert.Nil(t, l.First().next)

	assert.NotNil(t, l.Last())
	assert.Equal(t, l.Last().Value(), 1)
	assert.Nil(t, l.Last().prev)
	assert.Nil(t, l.Last().next)
}

func TestPushFrontInListWithOneBackItem(t *testing.T) {
	l := List{}

	l.PushBack(0)
	l.PushFront(1)

	assert.Equal(t, 2, l.Len())

	assert.NotNil(t, l.First())
	assert.Equal(t, l.First().Value(), 1)
	assert.Equal(t, l.First().prev, l.Last())
	assert.Nil(t, l.First().next)

	assert.NotNil(t, l.Last())
	assert.Equal(t, l.Last().Value(), 0)
	assert.Equal(t, l.Last().next, l.First())
	assert.Nil(t, l.Last().prev)
}

func TestPushFrontInListWithOneFrontItem(t *testing.T) {
	l := List{}

	l.PushFront(0)
	l.PushFront(1)

	assert.Equal(t, 2, l.Len())

	assert.NotNil(t, l.First())
	assert.Equal(t, l.First().Value(), 1)
	assert.Equal(t, l.First().prev, l.Last())
	assert.Nil(t, l.First().next)

	assert.NotNil(t, l.Last())
	assert.Equal(t, l.Last().Value(), 0)
	assert.Equal(t, l.Last().next, l.First())
	assert.Nil(t, l.Last().prev)
}

func TestPushFrontInListWithTwoItems(t *testing.T) {

}

func TestPushFrontInListWithThreeItems(t *testing.T) {

}

// Pushing back

// TestCase when an empty string is passed to the frequency analyzer
func TestPushBackInEmptyList(t *testing.T) {

}

func TestPushBackInListWithOneBackItem(t *testing.T) {

}

func TestPushBackInListWithOneFrontItem(t *testing.T) {

}

func TestPushBackInListWithTwoItems(t *testing.T) {

}

func TestPushBackInListWithThreeItems(t *testing.T) {

}

// Removing Item

//
func TestRemovingItemFromEmptyList(t *testing.T) {

}

func TestRemovingItemFromListWithOneBackItem(t *testing.T) {

}

func TestRemovingItemFromListWithOneFrontItem(t *testing.T) {

}

func TestRemovingItemFromListWithTwoItemsFromHead(t *testing.T) {

}

func TestRemovingItemFromListWithTwoItemsFromTail(t *testing.T) {

}

func TestRemovingItemFromListWithThreeItemsFromCenter(t *testing.T) {

}
