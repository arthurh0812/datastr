package linkedlist

import "sync"

type node struct {
	val interface{}
	next *node
}

type LinkedList struct {
	head *node
	tail *node
	len int64
	mu sync.RWMutex // guards internal state of linked list
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
	if n.next != nil { // make sure the node doesnt hold a reference to anywhere else
		n.next = nil
	}
	l.mu.Lock()
	l.tail.next = n
	l.tail = n
	l.len++
	l.mu.Unlock()
}

func (l *LinkedList) Append(val interface{}) {
	newNode := &node{val: val, next: nil}
	if l.IsEmpty() { // special case for empty list
		l.init(newNode)
		return
	}
	l.append(newNode)
}

func (l *LinkedList) Prepend(val interface{}) {
	newNode := &node{val: val}
	l.prepend(newNode)
}

// adds node 'after' after curr
func (l *LinkedList) insert(curr, after *node) {
	l.mu.Lock()
	after.next = curr.next
	curr.next = after
	l.len++
	l.mu.Unlock()
}

func (l *LinkedList) prepend(pre *node) {
	l.mu.Lock()
	pre.next = l.head
	l.head = pre
	l.len++
	l.mu.Unlock()
}

func (l *LinkedList) traverse(idx int64) *node {
	trav := l.head
	for i := int64(0); i < idx; i++ {
		trav = trav.next
	}
	return trav
}

// traverses to the most previous node of the node with a value match
func (l *LinkedList) where(val interface{}) *node {
	search := l.head
	for next := search.next; search != nil && next != nil; search = search.next {
		if next.val == val {
			return search
		}
	}
	return nil
}

func (l *LinkedList) InsertAt(val interface{}, index int64) {
	if index < 0 || l.len <= index {
		return
	}
	newNode := &node{val: val, next: nil}
	if index == 0 { // insert before head, if index is 0
		l.prepend(newNode)
		return
	}
	prev := l.traverse(index-1) // travel to pre-index node
	l.insert(prev, newNode)
}

func (l *LinkedList) InsertWhere(val, whereVal interface{}) {
	if l.IsEmpty() {
		return
	}
	newNode := &node{val: val, next: nil}
	if l.head.val == whereVal { // insert before head, if match
		l.prepend(newNode)
		return
	}
	prevSearched := l.where(whereVal)
	if prevSearched != nil { // if a match has been found
		l.insert(prevSearched, newNode)
	}
}

func Empty() *LinkedList {
	return &LinkedList{
		head: nil,
		tail: nil,
		len: 0,
	}
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
