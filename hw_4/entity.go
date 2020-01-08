package hw_4

type List struct {
	length int
	head   *Item
	tail   *Item
}

func (l List) Len() int {
	return l.length
}

func (l List) First() Item {
	if l.head == nil && l.tail != nil {
		return *l.tail
	} else {
		return *l.head
	}
}

func (l List) Last() Item {
	if l.tail == nil && l.head != nil {
		return *l.head
	} else {
		return *l.tail
	}
}

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

func (l List) PushBack(v interface{}) {
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

func (l List) Remove(i Item) {
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

type Item struct {
	next  *Item
	prev  *Item
	value interface{}
}

func (i Item) Next() *Item {
	return i.next
}

func (i Item) Prev() *Item {
	return i.prev
}

func (i Item) Value() interface{} {
	return i.value
}
