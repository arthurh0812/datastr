package hashtable

import (
	"github.com/arthurh0812/datastruct/types"
)

func (h *HashTable) remove(key types.Value) *Entry {
	row := h.table.getRow(key, h.fn)
	for i, val := range row.Values() {
		if entry, ok := val.(*Entry); ok && entry.Key.IsEqualTo(key) {
			row.RemoveAt(int64(i))
			h.decreaseSize()
			return entry
		}
	}
	return nil
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
