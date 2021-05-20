package hashtable

import (
	"errors"
	"sync"

	"github.com/arthurh0812/datastruct/linkedlist"
	"github.com/arthurh0812/datastruct/types"
)

var ErrInvalidCapacity = errors.New("table capacity must be greater than or equal to 0")
var ErrInvalidLoadFactor = errors.New("load factor must be greater than 0")
var ErrInvalidHashFunction = errors.New("hash function must not be nil")

// HashTable concrete data structure
type HashTable struct {
	fn Function // the hash function
	table []*linkedlist.LinkedList
	capacity, size int64
	loadFactor float64 // the factor to multiply the capacity with when full

	mu sync.Mutex
}

func (h *HashTable) Size() int64 {
	return h.size
}

func (h *HashTable) Capacity() int64 {
	return h.capacity
}

func (h *HashTable) isEmpty() bool {
	return h == nil || h.table == nil || h.capacity == 0 || h.size == 0
}

func (h *HashTable) IsEmpty() bool {
	return h.isEmpty()
}

func (h *HashTable) setTable(table []*linkedlist.LinkedList) {
	h.table = table
}

func (h *HashTable) setCapacity(cap int64) {
	h.mu.Lock()
	h.capacity = cap
	h.mu.Unlock()
}

func (h *HashTable) initCapacity(cap int64) error {
	if cap < 0 {
		return ErrInvalidCapacity
	}
	h.setCapacity(cap)
	return nil
}

func (h *HashTable) setFunction(fn Function) {
	h.mu.Lock()
	h.fn = fn
	h.mu.Unlock()
}

func (h *HashTable) SetFunction(fn Function) error {
	if fn == nil {
		return ErrInvalidHashFunction
	}
	h.setFunction(fn)
	return nil
}

func (h *HashTable) setLoadFactor(factor float64) {
	h.mu.Lock()
	h.loadFactor = factor
	h.mu.Unlock()
}


func (h *HashTable) SetLoadFactor(factor float64) (err error) {
	if factor <= 0 {
		return ErrInvalidLoadFactor
	}
	h.setLoadFactor(factor)
	return nil
}

func (h *HashTable) isOutOfBounds(idx int) bool {
	return idx < 0 || len(h.table)-1 < idx
}

func (h *HashTable) normalizeIndex(hash int) (idx int){
	return (hash & 0x7FFFFFFF) % 10
}

func (h *HashTable) Insert(key types.Value, val interface{}) {
	entry := NewEntry(key, val, h.fn)
	idx := h.normalizeIndex(entry.Hash)
	if h.isOutOfBounds(idx) {
		// increase hash table
	}
	items := h.table[idx]
	items.Append(val)
}

func (h *HashTable) clear() {
	h.mu.Lock()
	h.table = h.table[0:0]
	h.size = 0
	h.mu.Unlock()
}

func (h *HashTable) Clear() {
	h.clear()
}