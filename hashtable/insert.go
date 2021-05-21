package hashtable

import (
	"github.com/arthurh0812/datastruct/linkedlist"
	"github.com/arthurh0812/datastruct/types"
)

func (h *HashTable) Insert(key types.Value, val interface{}) (*Entry, error) {
	if key == nil {
		return nil, ErrInvalidKey
	}
	entry := NewEntry(key, val, h.fn)
	idx := h.normalizeIndex(entry.Hash)
	h.insertEntry(entry, idx)
	return entry, nil
}

// expects: e != nil && idx = normalized
func (h *HashTable) insertEntry(e *Entry, idx int) {
	if h.shouldExtend(idx) {
		h.extendTable(idx)
	}
	row := h.table.getRowByIndex(idx)
	row.Append(e.Val)
	h.increaseSize()
}

func (h *HashTable) extendTable(idx int) {
	h.mu.Lock()
	var newTable []*linkedlist.LinkedList
	newlen := idx
	if newlen <= cap(h.table) { // there is room to grow, extend the slice
		newTable = h.table[:newlen]
	} else { // there is no room to grow, create a new underlying array
		newcap := newlen
		if newcap < 2*len(h.table) {
			newcap = 2*len(h.table)
		}
		newTable = make([]*linkedlist.LinkedList, newlen, newcap)
		copy(newTable, h.table)
	}
	newTable[newlen-1] = linkedlist.Empty()
	h.table = newTable
	h.mu.Unlock()
}