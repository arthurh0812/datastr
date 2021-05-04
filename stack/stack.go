package stack

import "github.com/arthurh0812/datastr/linkedlist"

type Stack struct {
	list *linkedlist.LinkedList
}

func (s *Stack) Push(val interface{}) {
	s.list.Prepend(val)
}

func New(val interface{}) *Stack {
	return &Stack{
		list: linkedlist.New(val, linkedlist.Empty()),
	}
}