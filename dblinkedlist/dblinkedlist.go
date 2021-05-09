package dblinkedlist

import (
	"fmt"
	"io"
	"strings"
	"sync"
)

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

func (l *LinkedList) isEmpty() bool {
	return l == nil || l.head == nil || l.tail == nil || l.len == 0
}

func (l *LinkedList) IsEmpty() bool {
	return l.isEmpty()
}

func (l *LinkedList) append(n *node) {
	l.mu.Lock()
	if n != nil {
		n.next = nil
		n.prev = l.tail
	}
	l.tail.next = n
	l.tail = n
	l.len++
	l.mu.Unlock()
}

func (l *LinkedList) prepend(n *node) {
	l.mu.Lock()
	if n != nil {
		n.prev = nil
		n.next = l.head
	}
	l.head.prev = n
	l.head = n
	l.len++
	l.mu.Unlock()
}

func (l *LinkedList) insertAfter(first, second *node) {
	second.next = first.next
	first.next = second
	second.prev = first
	l.mu.Lock()
	l.len++
	l.mu.Unlock()
}

func (l *LinkedList) insertBefore(first, second *node) {
	second.prev = first.prev
	first.prev = second
	second.next = first
	l.mu.Lock()
	l.len++
	l.mu.Unlock()
}

func (l *LinkedList) traverseUp(n int64) *node {
	trav := l.head
	for i := int64(0); i < n; i++ {
		if trav == nil {
			return nil
		}
		trav = trav.next
	}
	return trav
}

func (l *LinkedList) traverseDown(n int64) *node {
	trav := l.tail
	for i := l.len-1; i > n; i-- {
		if trav == nil {
			return nil
		}
		trav = trav.prev
	}
	return trav
}

// returns the node at index 'idx' (zero-based)
func (l *LinkedList) traverse(idx int64) *node {
	if idx < 0 || l.len <= idx {
		return nil
	}
	mid := l.len / 2
	if idx < mid { // traverse list starting at head
		return l.traverseUp(idx)
	} else { // traverse list starting at tail
		return l.traverseDown(idx)
	}
}

func (l *LinkedList) where(val interface{}) *node {
	for trav := l.head; trav != nil; trav = trav.next {
		if trav.val == val {
			return trav
		}
	}
	return nil
}

func (l *LinkedList) Append(val interface{}) {
	newNode := &node{val: val}
	if l.isEmpty() { // empty list case is always the same, no matter prepend/append etc.
		l.init(newNode)
		return
	}
	l.append(newNode)
}

func (l *LinkedList) Prepend(val interface{}) {
	newNode := &node{val: val}
	if l.isEmpty() { // empty list case
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

func (l *LinkedList) values() []interface{} {
	vals := make([]interface{}, 0, l.len)
	for trav := l.head; trav != nil; trav = trav.next {
		vals = append(vals, trav.val)
	}
	return vals
}

func (l *LinkedList) Values() []interface{} {
	return l.values()
}

func (l *LinkedList) string() string {
	b := &strings.Builder{}
	vals := l.values()
	b.WriteByte('[')
	writeArray(b, vals, " ; ")
	b.WriteByte(']')
	return b.String()
}

func (l *LinkedList) String() string {
	return l.string()
}

func writeArray(w io.Writer, arr []interface{}, sep string) {
	for i, obj := range arr {
		if 0 < i {
			w.Write([]byte(sep))
		}
		valString := fmt.Sprintf("%v", obj)
		if s, ok := obj.(string); ok {
			valString = fmt.Sprintf("%q", s)
		}
		w.Write([]byte(valString))
	}
}