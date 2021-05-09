package heap

import "github.com/arthurh0812/datastr/types"

func (h *Heap) removeLast() {
	h.mu.Lock()
	h.arr = h.arr[:len(h.arr)-1]
	h.mu.Unlock()
}

func (h *Heap) remove(idx int) {
	if idx < len(h.arr)-1 {  // swap with the last, if not already last
		h.swap(idx, len(h.arr)-1)
	}
	h.removeLast()
	down := h.decideBubble(idx)
	if down {
		h.bubbleDown()
	} else {
		h.bubbleUp()
	}
}

func (h *Heap) Remove(val types.Value) {
	idx := h.getIndex(val) // table lookup
	h.remove(idx) // removal at index
}