package stack

import "github.com/arthurh0812/datastr/linkedlist"

// Empty creates a new empty Stack that satisfies Stack.IsEmpty.
func Empty() *Stack {
	return &Stack{
		list: linkedlist.Empty(),
	}
}

// New creates a new Stack with the given value inserted at the head, and the given Stack appended.
func New(val interface{}, append *Stack) *Stack {
	toAppend := linkedlist.Empty()
	if !append.isEmpty() {
		toAppend = append.list
	}
	return &Stack{
		list: linkedlist.New(val, toAppend),
	}
}
