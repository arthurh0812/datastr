package queue

import (
	"github.com/arthurh0812/datastr/linkedlist"
)

type Queue struct {
	list *linkedlist.LinkedList
}

func (q *Queue) isEmpty() bool {
	return q == nil || q.list.IsEmpty()
}

func (q *Queue) Last() interface{} {
	if q.isEmpty() {
		return nil
	}
	return q.list.Head()
}

func (q *Queue) Peek() interface{} {
	if q.isEmpty() {
		return nil
	}
	return q.list.Tail()
}

func (q *Queue) IsEmpty() bool {
	return q.isEmpty()
}

func (q *Queue) Enqueue(val interface{}) {
	if q.isEmpty() {
		return
	}
	q.list.Prepend(val)
}

func (q *Queue) Dequeue() (val interface{}) {
	if q.isEmpty() {
		return nil
	}
	return q.list.RemoveTail()
}

func (q *Queue) String() string {
	if q.isEmpty() {
		return "[]"
	}
	return q.list.String()
}