package datastr

type node struct {
	val interface{}
	next *node
}

type LinkedList struct {
	head *node
	len int64
}

func NewLinkedList(val interface{}, next *LinkedList) *LinkedList {
	var nextLen int64
	var nextHead *node
	if next != nil {
		nextLen = next.len
		nextHead = next.head
	}
	return &LinkedList{
		head: &node{val: val, next: nextHead},
		len: 1 + nextLen,
	}
}
