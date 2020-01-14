package list

// Item is an entity List data structure or node of it.
type Item struct {
	next, prev *Item
	list       *List
	value      interface{}
}

// Next method returns the next Item instance behind current Item.
func (i *Item) Next() *Item {
	if p := i.next; i.list != nil && p != &i.list.head {
		return p
	}

	return nil
}

// Prev method returns the previous Item instance before current Item.
func (i *Item) Prev() *Item {
	if p := i.prev; i.list != nil && p != &i.list.head {
		return p
	}

	return nil
}

// Value method return value of the current Item.
func (i *Item) Value() interface{} {
	return i.value
}

// List is data structure known as a Doubly linked list
// ref: https://en.wikipedia.org/wiki/Doubly_linked_list
type List struct {
	head   Item
	length int
}

// Init initializes or clears list l.
func (l *List) Init() *List {
	l.head.next = &l.head
	l.head.prev = &l.head
	l.length = 0

	return l
}

// NewList returns an initialized list.
func NewList() *List { return new(List).Init() }

// Len method returns length of List data structure.
func (l *List) Len() int { return l.length }

// First method returns the first item.
func (l *List) First() *Item {
	if l.length == 0 {
		return nil
	}

	return l.head.next
}

// Last method returns the last item.
func (l *List) Last() *Item {
	if l.length == 0 {
		return nil
	}

	return l.head.prev
}

// lazyInit lazily initializes a zero List value.
func (l *List) lazyInit() {
	if l.head.next == nil {
		l.Init()
	}
}

// insert method inserts item after at with value v increments l.len, and returns inserted item.
func (l *List) insert(v interface{}, at *Item) *Item {
	item := &Item{value: v}

	nextItem := at.next
	at.next = item
	item.prev = at
	item.next = nextItem
	nextItem.prev = item
	item.list = l
	l.length++

	return item
}

// PushFront method adds item to the head of List data structure with value - v.
func (l *List) PushFront(v interface{}) *Item {
	l.lazyInit()

	return l.insert(v, &l.head)
}

// PushBack method adds item to the tail of List data structure with value - v.
func (l *List) PushBack(v interface{}) *Item {
	l.lazyInit()
	return l.insert(v, l.head.prev)
}

// Remove method removes item - Item of the List data structure.
func (l *List) Remove(item *Item) interface{} {
	if item.list == l {
		item.prev.next = item.next
		item.next.prev = item.prev
		item.next = nil
		item.prev = nil
		item.list = nil
		l.length--
	}

	return item.Value
}
