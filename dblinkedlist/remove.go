package dblinkedlist

func (l *LinkedList) remove(n *node) {
	if n.next == nil {
		l.removeTail()
		return
	} else if n.prev == nil {
		l.removeHead()
		return
	}
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
	head := l.head
	l.mu.Lock()
	l.head = head.next
	if l.IsEmpty() {
		l.tail = nil
	} else {
		l.head.prev = nil
	}
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

func (l *LinkedList) removeTail() *node {
	tail := l.tail
	l.mu.Lock()
	l.tail = tail.prev
	if l.IsEmpty() {
		l.head = nil
	} else {
		l.tail.next = nil
	}
	tail.prev = nil
	l.len--
	l.mu.Unlock()
	return tail
}

func (l *LinkedList) RemoveTail() interface{} {
	if l.IsEmpty() {
		return nil
	}
	t := l.removeTail()
	return t.val
}

