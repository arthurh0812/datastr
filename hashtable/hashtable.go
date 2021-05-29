package hashtable

import (
	"errors"
	"github.com/arthurh0812/datastruct/types"
	"sync"

	"github.com/arthurh0812/datastruct/linkedlist"
)

var ErrInvalidKey = errors.New("key must not be nil")
var ErrInvalidCapacity = errors.New("table capacity must be greater than or equal to 0")
var ErrInvalidLoadFactor = errors.New("load factor must be greater than 0")
var ErrInvalidHashFunction = errors.New("hash function must not be nil")

// HashTable concrete data structure
type HashTable struct {
	fn Function // the hash function
	table table // the table consisting of rows
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

// Only use is you are aware if you are certain that the keys are all of the same type.
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

func (h *HashTable) increaseSize() {
	h.mu.Lock()
	h.size++
	h.mu.Unlock()
}

func (h *HashTable) decreaseSize() {
	h.mu.Lock()
	h.size--
	h.mu.Unlock()
}

func (h *HashTable) isOutOfBounds(idx int) bool {
	return h.table.isOutOfBounds(idx)
}

// determines whether the table should be extended (depends on load factor)
func (h *HashTable) shouldExtend(idx int) bool {
	x := h.loadFactor * float64(h.size)
	aim := int(x)
	return idx >= aim
}

func (h *HashTable) normalizeIndex(hash int) (idx int){
	return (hash & 0x7FFFFFFF) % 10
}

func (h *HashTable) Keys() []types.Value {
	keys := make([]types.Value, 0, h.size)
	h.table.loop(func(e *Entry) {
		keys = append(keys, e.Key)
	})
	return keys
}

func (h *HashTable) Values() []interface{} {
	vals := make([]interface{}, 0, h.size)
	h.table.loop(func(e *Entry) {
		vals = append(vals, e.Val)
	})
	return vals
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