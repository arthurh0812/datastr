package linkedlist

// Empty creates a new empty LinkedList.
func Empty() *LinkedList {
	return &LinkedList{
		head: nil,
		tail: nil,
		len: 0,
	}
}

// New creates a new LinkedList with the given value inserted to the head and the given LinkedList appended.
func New(val interface{}, next *LinkedList) *LinkedList {
	var nextLen int64
	var nextHead *node
	if next != nil {
		nextLen = next.len
		nextHead = next.head
	}
	initNode := &node{val: val}
	ll := Empty()
	ll.init(initNode)
	initNode.next = nextHead
	ll.len += nextLen
	return ll
}
