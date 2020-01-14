package main

import (
	"fmt"
	"github.com/teploff/otus/hw_4/list"
)

func main() {
	// Usage example
	l := list.List{}

	l.PushFront(0)
	l.PushBack(1)
	l.PushFront(2)
	l.PushBack(3)
	l.Remove(l.Last())
	l.Remove(l.First())
	l.Remove(l.Last())
	l.Remove(l.First())

	fmt.Println(l.Len())
}
