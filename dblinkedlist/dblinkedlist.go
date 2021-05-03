package dblinkedlist

import "sync"

type node struct {
	val interface{}
	prev, next *node
}

type LinkedList struct {
	head *node
	tail *node
	len int64
	mu sync.Mutex
}

func (l *LinkedList) init(n *node) {
	if n == nil {
		return
	}
	l.mu.Lock()
	l.head = n
	l.tail = n
	l.len = 1
	l.mu.Unlock()
}

func (l *LinkedList) clear() {
	l.mu.Lock()
	l.head = nil
	l.tail = nil
	l.len = 0
	l.mu.Unlock()
}

func (l *LinkedList) Clear() {
	l.clear()
}

func New(val interface{}, next *LinkedList) *LinkedList {
	var nextLen int64
	var nextHead *node
	if next != nil {
		nextLen = next.len
		nextHead = next.head
	}
	initNode := &node{val: val, next: nextHead}
	ll := &LinkedList{}
	ll.init(initNode)
	ll.len += nextLen
	return ll
}
