package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Pushing front

// TestCase pushing Item to the head of the empty list
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

// TestCase pushing Item to the head of the list which consists only tail Item
func TestPushFrontInListWithOneTailItem(t *testing.T) {
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

// TestCase pushing Item to the head of the list which consists only head Item
func TestPushFrontInListWithOneHeadItem(t *testing.T) {
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

// TestCase pushing Item to the head of the list which consists two Items
func TestPushFrontInListWithTwoItems(t *testing.T) {
	l := List{}

	l.PushFront(0)
	l.PushFront(1)
	l.PushFront(2)

	assert.Equal(t, 3, l.Len())

	assert.NotNil(t, l.First())
	assert.Equal(t, l.First().Value(), 2)
	assert.Equal(t, l.First().prev, l.Last().next)
	assert.Nil(t, l.First().next)

	assert.NotNil(t, l.First().prev)
	assert.Equal(t, l.First().prev.Value(), l.Last().next.Value())

	assert.NotNil(t, l.Last())
	assert.Equal(t, l.Last().Value(), 0)
	assert.Equal(t, l.Last().next, l.First().prev)
	assert.Nil(t, l.Last().prev)
}

// TestCase pushing Item to the head of the list which consists three Items
func TestPushFrontInListWithThreeItems(t *testing.T) {
	l := List{}

	l.PushFront(0)
	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)

	assert.Equal(t, 4, l.Len())

	assert.NotNil(t, l.First())
	assert.Equal(t, l.First().Value(), 3)
	assert.Nil(t, l.First().next)

	assert.NotNil(t, l.First().prev)
	assert.NotNil(t, l.Last().next.next)
	assert.Equal(t, l.First().prev, l.Last().next.next)
	assert.Equal(t, l.First().prev.Value(), l.Last().next.next.Value())
	assert.Equal(t, l.First().prev.Value(), 2)

	assert.NotNil(t, l.First().prev.prev)
	assert.NotNil(t, l.Last().next)
	assert.Equal(t, l.First().prev.prev, l.Last().next)
	assert.Equal(t, l.First().prev.prev.Value(), l.Last().next.Value())
	assert.Equal(t, l.First().prev.prev.Value(), 1)

	assert.NotNil(t, l.Last())
	assert.Equal(t, l.Last().Value(), 0)
	assert.Equal(t, l.Last().next, l.First().prev.prev)
	assert.Nil(t, l.Last().prev)
}

// Pushing back

// TestCase pushing Item to the tail of the empty list
func TestPushBackInEmptyList(t *testing.T) {
	l := List{}

	l.PushBack(1)

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

// TestCase pushing Item to the tail of the list which consists only tail Item
func TestPushBackInListWithOneTailItem(t *testing.T) {
	l := List{}

	l.PushBack(0)
	l.PushBack(1)

	assert.Equal(t, 2, l.Len())

	assert.NotNil(t, l.First())
	assert.Equal(t, l.First().Value(), 0)
	assert.Equal(t, l.First().prev, l.Last())
	assert.Nil(t, l.First().next)

	assert.NotNil(t, l.Last())
	assert.Equal(t, l.Last().Value(), 1)
	assert.Equal(t, l.Last().next, l.First())
	assert.Nil(t, l.Last().prev)
}

// TestCase pushing Item to the tail of the list which consists only head Item
func TestPushBackInListWithOneHeadItem(t *testing.T) {
	l := List{}

	l.PushFront(0)
	l.PushBack(1)

	assert.Equal(t, 2, l.Len())

	assert.NotNil(t, l.First())
	assert.Equal(t, l.First().Value(), 0)
	assert.Equal(t, l.First().prev, l.Last())
	assert.Nil(t, l.First().next)

	assert.NotNil(t, l.Last())
	assert.Equal(t, l.Last().Value(), 1)
	assert.Equal(t, l.Last().next, l.First())
	assert.Nil(t, l.Last().prev)
}

// TestCase pushing Item to the tail of the list which consists two Items
func TestPushBackInListWithTwoItems(t *testing.T) {
	l := List{}

	l.PushFront(0)
	l.PushFront(1)
	l.PushBack(2)

	assert.Equal(t, 3, l.Len())

	assert.NotNil(t, l.First())
	assert.Equal(t, l.First().Value(), 1)
	assert.Equal(t, l.First().prev, l.Last().next)
	assert.Nil(t, l.First().next)

	assert.NotNil(t, l.First().prev)
	assert.Equal(t, l.First().prev.Value(), l.Last().next.Value())

	assert.NotNil(t, l.Last())
	assert.Equal(t, l.Last().Value(), 2)
	assert.Equal(t, l.Last().next, l.First().prev)
	assert.Nil(t, l.Last().prev)
}

// TestCase pushing Item to the head of the list which consists three Items
func TestPushBackInListWithThreeItems(t *testing.T) {
	l := List{}

	l.PushFront(0)
	l.PushFront(1)
	l.PushFront(2)
	l.PushBack(3)

	assert.Equal(t, 4, l.Len())

	assert.NotNil(t, l.First())
	assert.Equal(t, l.First().Value(), 2)
	assert.Nil(t, l.First().next)

	assert.NotNil(t, l.First().prev)
	assert.NotNil(t, l.Last().next.next)
	assert.Equal(t, l.First().prev, l.Last().next.next)
	assert.Equal(t, l.First().prev.Value(), l.Last().next.next.Value())
	assert.Equal(t, l.First().prev.Value(), 1)

	assert.NotNil(t, l.First().prev.prev)
	assert.NotNil(t, l.Last().next)
	assert.Equal(t, l.First().prev.prev, l.Last().next)
	assert.Equal(t, l.First().prev.prev.Value(), l.Last().next.Value())
	assert.Equal(t, l.First().prev.prev.Value(), 0)

	assert.NotNil(t, l.Last())
	assert.Equal(t, l.Last().Value(), 3)
	assert.Equal(t, l.Last().next, l.First().prev.prev)
	assert.Nil(t, l.Last().prev)
}

// Removing Item

// TestCase removing Item from the empty list
func TestRemovingItemFromEmptyList(t *testing.T) {
	l := List{}

	l.Remove(Item{})

	assert.Equal(t, 0, l.Len())
	assert.Nil(t, l.First())
	assert.Nil(t, l.Last())
}

// TestCase removing Item from the head of the list which consists only tail Item
func TestRemovingItemFromListWithOneTailItem(t *testing.T) {
	l := List{}

	l.PushBack(0)
	l.Remove(*l.Last())

	assert.Equal(t, 0, l.Len())

	assert.Nil(t, l.First())
	assert.Nil(t, l.Last())
}

// TestCase removing Item from the tail of the list which consists only head Item
func TestRemovingItemFromListWithOneHeadItem(t *testing.T) {
	l := List{}

	l.PushFront(0)
	l.Remove(*l.First())

	assert.Equal(t, 0, l.Len())

	assert.Nil(t, l.First())
	assert.Nil(t, l.Last())
}

// TestCase removing Item from the head of the list which consists two Items
func TestRemovingItemFromListWithTwoItemsFromHead(t *testing.T) {
	l := List{}

	l.PushFront(0)
	l.PushFront(1)
	l.Remove(*l.First())

	assert.Equal(t, 1, l.Len())

	assert.NotNil(t, l.First())
	assert.Equal(t, l.First().Value(), 0)
	assert.Nil(t, l.First().prev)
	assert.Nil(t, l.First().next)

	assert.NotNil(t, l.Last())
	assert.Equal(t, l.Last().Value(), 0)
	assert.Nil(t, l.Last().next)
	assert.Nil(t, l.Last().prev)
}

// TestCase removing Item from the tail of the list which consists two Items
func TestRemovingItemFromListWithTwoItemsFromTail(t *testing.T) {
	l := List{}

	l.PushFront(0)
	l.PushFront(1)
	l.Remove(*l.Last())

	assert.Equal(t, 1, l.Len())

	assert.NotNil(t, l.First())
	assert.Equal(t, l.First().Value(), 1)
	assert.Nil(t, l.First().prev)
	assert.Nil(t, l.First().next)

	assert.NotNil(t, l.Last())
	assert.Equal(t, l.Last().Value(), 1)
	assert.Nil(t, l.Last().next)
	assert.Nil(t, l.Last().prev)
}

// TestCase removing Item from the center of the list which consists three Items
func TestRemovingItemFromListWithThreeItemsFromCenter(t *testing.T) {
	l := List{}

	l.PushFront(0)
	l.PushFront(1)
	l.PushFront(2)
	l.Remove(*l.Last().next)

	assert.Equal(t, 2, l.Len())

	assert.NotNil(t, l.First())
	assert.Equal(t, l.First().Value(), 2)
	assert.Equal(t, l.First().prev, l.Last())
	assert.Nil(t, l.First().next)

	assert.NotNil(t, l.Last())
	assert.Equal(t, l.Last().Value(), 0)
	assert.Equal(t, l.Last().next, l.First())
	assert.Nil(t, l.Last().prev)
}

// TestCase removing all Items from the head of the list which consists three Items
func TestRemovingAllItemsFromListWithThreeItemsFromTheHead(t *testing.T) {
	l := List{}

	l.PushFront(0)
	l.PushFront(1)
	l.PushFront(2)
	l.Remove(*l.First())
	l.Remove(*l.First())
	l.Remove(*l.First())

	assert.Equal(t, 0, l.Len())

	assert.Nil(t, l.First())
	assert.Nil(t, l.Last())
}

// TestCase removing all Items from the tail of the list which consists three Items
func TestRemovingAllItemsFromListWithThreeItemsFromTheTail(t *testing.T) {
	l := List{}

	l.PushFront(0)
	l.PushFront(1)
	l.PushFront(2)
	l.Remove(*l.Last())
	l.Remove(*l.Last())
	l.Remove(*l.Last())

	assert.Equal(t, 0, l.Len())

	assert.Nil(t, l.First())
	assert.Nil(t, l.Last())
}
