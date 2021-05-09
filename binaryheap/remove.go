package binaryheap

import "github.com/arthurh0812/datastruct/types"

func (h *Heap) removeLast() {
	h.mu.Lock()
	h.arr = h.arr[:h.size-1]
	h.mu.Unlock()
}

func (h *Heap) remove(idx int) {
	if idx < len(h.arr)-1 {  // swap with the last, if not already last
		h.swap(idx, h.size-1)
	}
	h.removeLast()
	h.decreaseSize()
	h.bubble(idx)
}

func (h *Heap) Remove(val types.Value) {
	if h.isEmpty() {
		return
	}
	idx := h.getIndex(val) // table lookup
	if idx == -1 {
		return
	}
	h.remove(idx) // removal of value at at heap index
	h.removeFromTable(val, idx) // removal of index from the table
}

func (h *Heap) RemoveAt(idx int) (val types.Value) {
	if h.isEmpty() || h.isOutOfBounds(idx) {
		return nil
	}
	val = h.getValue(idx) // heap lookup
	h.remove(idx) // removal of value at heap index
	h.removeFromTable(val, idx) // removal of index from the table
	return val
}