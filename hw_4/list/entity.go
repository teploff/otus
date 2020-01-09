package list

// List is data structure known as a Doubly linked list
// ref: https://en.wikipedia.org/wiki/Doubly_linked_list
type List struct {
	length int
	head   *Item
	tail   *Item
}

// Len method returns length of List data structure.
func (l List) Len() int {
	return l.length
}

// First method returns the first item.
func (l List) First() *Item {
	if l.head == nil && l.tail != nil {
		return l.tail
	}

	return l.head
}

// Last method returns the last item.
func (l List) Last() *Item {
	if l.tail == nil && l.head != nil {
		return l.head
	}

	return l.tail
}

// PushFront method adds item to the head of List data structure with value - v.
func (l *List) PushFront(v interface{}) {
	item := new(Item)

	if l.head == nil && l.tail == nil {
		item.value = v
		l.head = item
	} else if l.head != nil {
		item.prev = l.head
		item.value = v
		if l.tail == nil {
			l.tail = l.head
			l.tail.next = item
		}
		l.head.next = item
	} else {
		item.prev = l.tail
		item.value = v
		l.tail.next = item
	}

	l.length++
	l.head = item
}

// PushBack method adds item to the tail of List data structure with value - v.
func (l *List) PushBack(v interface{}) {
	item := new(Item)

	if l.head == nil && l.tail == nil {
		item.value = v
		l.tail = item
	} else if l.tail != nil {
		item.next = l.tail
		item.value = v
		if l.head == nil {
			l.head = l.tail
			l.head.prev = item
		}
		l.tail.prev = item
	} else {
		item.next = l.head
		item.value = v
		l.head.prev = item
	}

	l.length++
	l.tail = item
}

// Remove method removes item - Item of the List data structure.
func (l *List) Remove(i Item) {
	if l.length == 0 {
		return
	}

	if i.prev == nil && i.next == nil {
		l.head = nil
		l.tail = nil
	} else if i.prev != nil && i.next != nil {
		i.prev.next = i.next
		i.next.prev = i.prev
	} else if i.prev == nil {
		i.next.prev = nil
		l.tail = i.next
	} else {
		i.prev.next = nil
		l.head = i.prev
	}

	l.length--
}

// Item is an entity List data structure or node of it.
type Item struct {
	next  *Item
	prev  *Item
	value interface{}
}

// Next method returns the next Item instance behind current Item.
func (i Item) Next() *Item {
	return i.next
}

// Prev method returns the previous Item instance before current Item.
func (i Item) Prev() *Item {
	return i.prev
}

// Value method return value of the current Item.
func (i Item) Value() interface{} {
	return i.value
}
