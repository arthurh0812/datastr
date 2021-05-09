package stack

import "github.com/arthurh0812/datastruct/linkedlist"

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

func (s *Stack) String() string {
	if s.isEmpty() {
		return "[]"
	}
	return s.list.String()
}