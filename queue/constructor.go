package queue

import "github.com/arthurh0812/datastr/linkedlist"

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
