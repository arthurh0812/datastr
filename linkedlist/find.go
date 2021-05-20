package linkedlist

func (l *LinkedList) At(idx int) interface{} {
	n := l.traverse(int64(idx))
	return n.val
}

func (l *LinkedList) find(val interface{}) (idx int64) {
	for trav := l.head; trav != nil; trav = trav.next {
		if trav.val == val {
			return idx
		}
		idx++
	}
	return -1
}

func (l *LinkedList) Find(val interface{}) (idx int) {
	i := l.find(val)
	return int(i)
}
