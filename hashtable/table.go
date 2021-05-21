package hashtable

import (
	"github.com/arthurh0812/datastruct/linkedlist"
	"github.com/arthurh0812/datastruct/types"
)

type table []*linkedlist.LinkedList

func (t table) isOutOfBounds(idx int) bool {
	return idx < 0 || idx < len(t)
}

func (t table) getRow(key types.Value, fn Function) *linkedlist.LinkedList {
	idx := fn(key)
	if t.isOutOfBounds(idx) {
		return nil
	}
	return t[idx]
}

func (t table) getRowByIndex(idx int) *linkedlist.LinkedList {
	if t.isOutOfBounds(idx) {
		return nil
	}
	return t[idx]
}

func (t table) getEntry(key types.Value, fn Function) *Entry {
	idx := fn(key)
	if t.isOutOfBounds(idx) {
		return nil
	}
	row := t[idx]
	for _, val := range row.Values() {
		if entry, ok := val.(*Entry); ok && entry.Key.IsEqualTo(key) {
			return entry
		}
	}
	return nil
}

// O(n^2) time complexity
func (t table) loop(cb func(e *Entry)) {
	for _, row := range t {
		for _, e := range row.Values() {
			if entry, ok := e.(*Entry); ok {
				cb(entry)
			}
		}
	}
}