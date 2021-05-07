package stack

import "github.com/arthurh0812/datastr/linkedlist"

type Stack struct {
	list *linkedlist.LinkedList
}

func (s *Stack) First() interface{} {
	if s.isEmpty() {
		return nil
	}
	return s.list.Tail()
}

func (s *Stack) Last() interface{} {
	if s.isEmpty() {
		return nil
	}
	return s.list.Head()
}

func (s *Stack) isEmpty() bool {
	return s == nil || s.list.IsEmpty()
}

func (s *Stack) IsEmpty() bool {
	return s.isEmpty()
}

func (s *Stack) Push(val interface{}) {
	if s.isEmpty() {
		return
	}
	s.list.Prepend(val)
}

// Pop removes and returns the first element.
func (s *Stack) Pop() (val interface{}) {
	if s.isEmpty() {
		return
	}
	return s.list.RemoveHead()
}

func Empty() *Stack {
	return &Stack{
		list: linkedlist.Empty(),
	}
}

func New(val interface{}, append *Stack) *Stack {
	toAppend := linkedlist.Empty()
	if !append.isEmpty() {
		toAppend = append.list
	}
	return &Stack{
		list: linkedlist.New(val, toAppend),
	}
}

func (s *Stack) String() string {
	if s.isEmpty() {
		return "[]"
	}
	return s.list.String()
}

