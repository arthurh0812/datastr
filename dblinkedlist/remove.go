package dblinkedlist

func (l *LinkedList) remove(n *node) {
	prev := n.prev
	next := n.next
	prev.next = next
	next.prev = prev
	n.prev = nil
	n.next = nil
	l.mu.Lock()
	l.len--
	l.mu.Unlock()
}

func (l *LinkedList) Remove(idx int64) {
	toRemove := l.traverse(idx)
	if toRemove == nil {
		return
	}
	l.remove(toRemove)
}

func (l *LinkedList) RemoveWhere(val interface{}) {
	toRemove := l.where(val)
	if toRemove == nil {
		return
	}
	l.remove(toRemove)
}

func (l *LinkedList) removeHead() *node {
	l.mu.Lock()
	head := l.head
	l.head = head.next
	head.next = nil
	l.len--
	l.mu.Unlock()
	return head
}

func (l *LinkedList) RemoveHead() interface{} {
	if l.IsEmpty() {
		return nil
	}
	h := l.removeHead()
	return h.val
}

func (l *LinkedList) pop() *node {
	tail := l.tail
	l.mu.Lock()
	l.tail = tail.prev
	l.tail.next = nil
	tail.prev = nil
	l.len--
	l.mu.Unlock()
	return tail
}

func (l *LinkedList) Pop() interface{} {
	if l.IsEmpty() {
		return nil
	}
	p := l.pop()
	return p.val
}

