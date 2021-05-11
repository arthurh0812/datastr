package binaryheap

import (
	"sync"

	"github.com/arthurh0812/datastruct/types"
)

type Heap struct {
	arr []types.Value
	table map[types.Value][]int
	max bool // default=min
	size int

	mu sync.Mutex
}

func (h *Heap) getAll() []types.Value {
	if h.isEmpty() {
		return nil
	}
	return h.arr
}

func (h *Heap) getFirst() types.Value {
	if h.isEmpty() {
		return nil
	}
	return h.arr[0]
}

// sets max to true
func (h *Heap) setMax() {
	h.mu.Lock()
	h.max = true
	h.mu.Unlock()
}

// sets max to false
func (h *Heap) setMin() {
	h.mu.Lock()
	h.max = false
	h.mu.Unlock()
}

func (h *Heap) increaseSize() {
	h.mu.Lock()
	h.size++
	h.mu.Unlock()
}

func (h *Heap) decreaseSize() {
	h.mu.Lock()
	h.size--
	h.mu.Unlock()
}

func (h *Heap) isEmpty() bool {
	return h == nil || h.size == 0
}

func (h *Heap) IsEmpty() bool {
	return h.isEmpty()
}

func (h *Heap) isOutOfBounds(idx int) bool {
	return idx < 0 || h.size-1 < idx
}

func (h *Heap) getChildren(parent int) (left, right int) {
	if h == nil {
		return -1, -1
	}
	return parent*2+1, parent*2+2
}

func (h *Heap) getParent(child int) (parent int) {
	if h == nil {
		return 0
	}
	return (child-1)/2
}

func (h *Heap) swap(f, s int) {
	first, sec := h.arr[f], h.arr[s]
	h.mu.Lock()
	h.replaceTable(first, f, s) // change index in table
	h.replaceTable(sec, s, f) // change index in table
	h.arr[f], h.arr[s] = sec, first
	h.mu.Unlock()
}

func (h *Heap) appendArray(val types.Value) {
	h.mu.Lock()
	h.arr = append(h.arr, val)
	h.mu.Unlock()
}

func (h *Heap) appendTable(key types.Value, idx int) {
	h.mu.Lock()
	h.table[key] = append(h.table[key], idx)
	h.mu.Unlock()
}

func (h *Heap) getIndex(val types.Value) (idx int) {
	indices := h.table[val]
	if len(indices) == 0 {
		return -1
	}
	return indices[0] // retrieve first of the indices
}

func (h *Heap) getValue(idx int) (val types.Value) {
	if h.isEmpty() || h.isOutOfBounds(idx) {
		return nil
	}
	return h.arr[idx]
}

func (h *Heap) searchTable(key types.Value, idx int) (tableIdx int) {
	for tableIdx, heapIdx := range h.table[key] {
		if heapIdx == idx {
			return tableIdx
		}
	}
	return -1
}

func (h *Heap) replaceTable(key types.Value, toReplace, replaceWith int) {
	tableIdx := h.searchTable(key, toReplace)
	if tableIdx < 0 {
		return
	}
	h.mu.Lock()
	h.table[key][tableIdx] = replaceWith
	h.mu.Unlock()
}

func (h *Heap) removeFromTable(key types.Value, idx int) {
	tableIdx := h.searchTable(key, idx)
	if tableIdx < 0 {
		return
	}
	indices := h.table[key]
	h.mu.Lock()
	h.table[key] = append(indices[:tableIdx], indices[tableIdx+1:]...) // remove the index of the array of the table
	h.mu.Unlock()
}

func (h *Heap) append(val types.Value) {
	h.appendArray(val) // append the value to the end of the array
	h.appendTable(val, h.size-1) // add the new last index to the table for key 'val'
	h.increaseSize()
}

func (h *Heap) Insert(val types.Value) {
	h.append(val) // append the new value to array and table
	h.bubbleUp() // reorganize binary heap (upwards)
}

func (h *Heap) InsertAll(vals []types.Value) {
	for _, val := range vals {
		h.Insert(val)
	}
}

func (h *Heap) Poll() (val types.Value) {
	 if h.isEmpty() {
	 	return nil
	 }
	 v := h.getFirst()
	 h.remove(0) // remove the item at index 0 (first)
	 return v
}

func (h *Heap) Peek() (val types.Value) {
	if h.isEmpty() {
		return nil
	}
	return h.getFirst()
}

// MakeMin makes this heap a minimum priority queue; this is the default.
func (h *Heap) MakeMin() {
	if !h.max { // if it already is a minimum heap
		return
	}
	all := h.getAll()
	h.clear()
	h.setMin()
	h.InsertAll(all)
}

// MakeMax makes this heap a maximum priority queue.
func (h *Heap) MakeMax() {
	if h.max { // if it already is a maximum heap
		return
	}
	all := h.getAll()
	h.clear()
	h.setMax()
	h.InsertAll(all)
}

func (h *Heap) clear() {
	h.mu.Lock()
	h.arr = make([]types.Value, 0, 0)
	h.table = make(map[types.Value][]int)
	h.max = false
	h.size = 0
	h.mu.Unlock()
}

func (h *Heap) Clear() {
	h.clear()
}