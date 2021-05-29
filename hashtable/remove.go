package hashtable

import (
	"github.com/arthurh0812/datastruct/linkedlist"
	"github.com/arthurh0812/datastruct/types"
)

func (h *HashTable) remove(key types.Value) *Entry {
	row := h.table.getRow(key, h.fn)
	for i, val := range row.Values() {
		if entry, ok := val.(*Entry); ok && entry.Key.IsEqualTo(key) {
			h.removeEntry(row, i)
			return entry
		}
	}
	return nil
}

func (h *HashTable) removeEntry(row *linkedlist.LinkedList, idx int) {
	row.RemoveAt(int64(idx))
	h.decreaseSize()
}

func (h *HashTable) Remove(key types.Value) interface{} {
	if key == nil {
		return nil
	}
	e := h.remove(key)
	if e == nil {
		return nil
	}
	return e.Val
}
