package dblinkedlist

// Empty creates a new empty doubly linked list that satisfies LinkedList.IsEmpty.
func Empty() *LinkedList {
	return &LinkedList{
		tail: nil,
		head: nil,
		len: 0,
	}
}

// New creates a new doubly linked list with the given value inserted at the head, and the given list appended.
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

