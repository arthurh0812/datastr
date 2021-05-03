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
	mu sync.RWMutex
}

func (l *LinkedList) Head() interface{} {
	if l.IsEmpty() {
		return nil
	}
	return l.head.val
}

func (l *LinkedList) Tail() interface{} {
	if l.IsEmpty() {
		return nil
	}
	return l.tail.val
}

func (l *LinkedList) Len() int64 {
	return l.len
}

func (l *LinkedList) IsEmpty() bool {
	return l == nil || l.head == nil || l.tail == nil || l.len == 0
}

func (l *LinkedList) append(n *node) {
	if n != nil {
		n.next = nil
		n.prev = l.tail
	}
	l.mu.Lock()
	l.tail.next = n
	l.tail = n
	l.len++
	l.mu.Unlock()
}

func (l *LinkedList) prepend(n *node) {
	if n != nil {
		n.prev = nil
		n.next = l.head
	}
	l.mu.Lock()
	l.head.prev = n
	l.head = n
	l.len++
	l.mu.Unlock()
}

func (l *LinkedList) Append(val interface{}) {
	newNode := &node{val: val}
	if l.IsEmpty() {
		l.init(newNode)
		return
	}
	l.append(newNode)
}

func (l *LinkedList) Prepend(val interface{}) {
	newNode := &node{val: val}
	if l.IsEmpty() {
		l.init(newNode)
		return
	}
	l.prepend(newNode)
}

func (l *LinkedList) init(n *node) {
	if n == nil {
		return
	}
	n.next = nil
	n.prev = nil
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
	initNode := &node{val: val}
	ll := &LinkedList{}
	ll.init(initNode)
	initNode = nextHead
	ll.len += nextLen
	return ll
}
