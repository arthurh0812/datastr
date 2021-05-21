package hashtable

import "github.com/arthurh0812/datastruct/types"

func (h *HashTable) find(key types.Value) *Entry {
	return h.table.getEntry(key, h.fn)
}

func (h *HashTable) Get(key types.Value) interface{} {
	if key == nil {
		return nil
	}
	e := h.find(key)
	if e == nil {
		return nil
	}
	return e.Val
}
