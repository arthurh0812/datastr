package linkedlist

func (l *LinkedList) remove(prev *node) {
	if tail := prev.next; tail != nil && tail.next == nil {
		l.removeTail()
		return
	}
	l.mu.Lock()
	toRemove := prev.next
	if toRemove != nil {
		prev.next = toRemove.next
		toRemove.next = nil
	}
	l.len--
	l.mu.Unlock()
}

func (l *LinkedList) RemoveAt(idx int64) {
	if l.isEmpty() {
		return
	}
	prev := l.traverse(idx)
	l.remove(prev)
}

func (l *LinkedList) RemoveWhere(whereVal interface{}) {
	if l.isEmpty() {
		return
	}
	prevSearched := l.where(whereVal)
	l.remove(prevSearched)
}

// list must not be empty; removes and returns head node
func (l *LinkedList) removeHead() *node {
	head := l.head
	l.mu.Lock()
	l.head = head.next
	head.next = nil
	if l.isEmpty() {
		l.tail = nil
	}
	l.len--
	l.mu.Unlock()
	return head
}

func (l *LinkedList) RemoveHead() interface{} {
	if l.isEmpty() {
		return nil
	}
	h := l.removeHead()
	return h.val
}

// list must not be empty; removes and returns tail node
func (l *LinkedList) removeTail() *node {
	tailIdx := l.len-1
	preTail := l.traverse(tailIdx-1)
	if preTail == nil { // only possible case: if tailIdx-1 is -1 which means head=tail
		return l.removeHead()
	}
	tail := preTail.next
	preTail.next = nil
	l.mu.Lock()
	l.tail = preTail
	l.len--
	l.mu.Unlock()
	return tail
}

func (l *LinkedList) RemoveTail() interface{} {
	if l.isEmpty() {
		return nil
	}
	t := l.removeTail()
	return t.val
}