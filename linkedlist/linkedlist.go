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
	l.tail.next = n
	l.tail = n
	l.len++
}

func (l *LinkedList) Append(val interface{}) {
	newNode := &node{val: val, next: nil}
	l.mu.Unlock()
	defer l.mu.Unlock()
	if l.IsEmpty() { // special case for empty list
		l.init(newNode)
		return
	}
	l.append(newNode)
}

func (l *LinkedList) Prepend(val interface{}) {
	newNode := &node{val: val}
	l.mu.Lock()
	defer l.mu.Unlock()
	l.prepend(newNode)
}

// adds node 'after' after curr
func (l *LinkedList) insert(curr, after *node) {
	after.next = curr.next
	curr.next = after
	l.len++
}

func (l *LinkedList) prepend(pre *node) {
	pre.next = l.head
	l.head = pre
	l.len++
}

func (l *LinkedList) traverse(idx int64) *node {
	trav := l.head
	for i := int64(0); i < idx; i++ {
		trav = trav.next
	}
	return trav
}

func (l *LinkedList) search(val interface{}) *node {
	search := l.head
	for search != nil {
		if search.val == val {
			return search
		}
		search = search.next
	}
	return nil
}

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
	preSearched := l.where(whereVal)
	if preSearched != nil { // if a match has been found
		l.insert(preSearched, newNode)
	}
}

func Empty() *LinkedList {
	return &LinkedList{
		head: nil,
		len: 0,
	}
}

func (l *LinkedList) init(n *node) {
	if n == nil {
		return
	}
	l.head = n
	l.tail = n
	l.len = 1
}

func New(val interface{}, next *LinkedList) *LinkedList {
	var nextLen int64
	var nextHead *node
	if next != nil {
		nextLen = next.len
		nextHead = next.head
	}
	initNode := &node{val: val, next: nextHead}
	return &LinkedList{
		head: initNode,
		tail: initNode,
		len: 1 + nextLen,
	}
}


