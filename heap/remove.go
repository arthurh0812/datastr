package heap

import "github.com/arthurh0812/datastr/types"

func (h *Heap) removeLast() {
	h.arr = h.arr[:len(h.arr)-1]
}

func (h *Heap) remove(val types.Value) {
	idx := h.getIndex(val) // table lookup
	if idx < len(h.arr)-1 {
		h.swap(idx, len(h.arr)-1) // swap with the last
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
	h.remove(val)
}