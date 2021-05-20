package hashtable

import "github.com/arthurh0812/datastruct/types"

func (h *HashTable) find(key types.Value) *Entry {
	idx := h.normalizeIndex(h.fn(key))
	if h.isOutOfBounds(idx) {
		return nil
	}
	items := h.table[idx]
	for _, val := range items.Values() {
		if entry, ok := val.(*Entry); ok && entry.Key.IsEqualTo(key) {
			return entry
		}
	}
	return nil
}

func (h *HashTable) Get(key types.Value) interface{} {
	e := h.find(key)
	if e == nil {
		return nil
	}
	return e.Val
}