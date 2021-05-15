package queue

import "github.com/arthurh0812/datastruct/linkedlist"

func Empty() *Queue {
	return &Queue{
		list: linkedlist.Empty(),
	}
}

// New creates a new Queue that is inserted the given value and appended with the given Queue.
func New(val interface{}, append *Queue) *Queue {
	toAppend := linkedlist.Empty()
	if !append.isEmpty() {
		toAppend = append.list
	}
	return &Queue{
		list: linkedlist.New(val, toAppend),
	}
}
