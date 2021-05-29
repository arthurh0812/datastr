package hashtable

import "github.com/arthurh0812/datastruct/types"

func (h *HashTable) find(key types.Value) *Entry {
	return h.table.getEntry(key, h.fn)
}

func (h *HashTable) findValue(val interface{}) *Entry {
	return h.table.getEntryByValue(val)
}

func (h *HashTable) Get(key types.Value) (val interface{}) {
	return h.get(key)
}

func (h *HashTable) get(key types.Value) interface{} {
	if key == nil {
		return nil
	}
	e := h.find(key)
	if e == nil {
		return nil
	}
	return e.Val
}

func (h *HashTable) GetKey(val interface{}) types.Value {
	return h.getKey(val)
}

func (h *HashTable) getKey(val interface{}) types.Value {
	e := h.findValue(val)
	if e == nil {
		return nil
	}
	return e.Key
}

func (h *HashTable) Contains(key types.Value) bool {
	return h.contains(key)
}

func (h *HashTable) contains(key types.Value) bool {
	return key != nil && h.find(key) != nil
}

func (h *HashTable) ContainsValue(val types.Value) bool {
	return h.findValue(val) != nil
}