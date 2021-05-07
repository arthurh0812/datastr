package stack

import "github.com/arthurh0812/datastr/linkedlist"

type Stack struct {
	list *linkedlist.LinkedList
}

func (s *Stack) isEmpty() bool {
	return s == nil || s.list.IsEmpty()
}

func (s *Stack) IsEmpty() bool {
	return s.isEmpty()
}

func (s *Stack) Push(val interface{}) {
	if s == nil {
		return
	}
	s.list.Prepend(val)
}

// Pop removes and returns the first element.
func (s *Stack) Pop() (val interface{}) {
	if s == nil {
		return
	}
	return s.list.RemoveHead()
}

// Peek only returns the first element.
func (s *Stack) Peek() (val interface{}) {
	if s == nil {
		return nil
	}
	return s.list.Head()
}

func Empty() *Stack {
	return &Stack{
		list: linkedlist.Empty(),
	}
}

func New(val interface{}) *Stack {
	return &Stack{
		list: linkedlist.New(val, linkedlist.Empty()),
	}
}

