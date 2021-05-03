package linkedlist

func (l *LinkedList) remove(prev *node) {
	if prev == nil {
		return
	}
	l.mu.Lock()
	toRemove := prev.next
	if toRemove != nil {
		prev.next = toRemove.next
		toRemove.next = nil
	}
	l.mu.Unlock()
}

func (l *LinkedList) RemoveAt(idx int64) {
	if l.IsEmpty() {
		return
	}
	prev := l.traverse(idx)
	l.remove(prev)
}

func (l *LinkedList) RemoveWhere(whereVal interface{}) {
	if l.IsEmpty() {
		return
	}
	prevSearched := l.where(whereVal)
	l.remove(prevSearched)
}