package binaryheap

// All of the bubble functionality

func (h *Heap) decideBubble(idx int) (down bool) {
	val := h.arr[idx]
	p := h.getParent(idx)
	l, r := h.getChildren(idx)
	if h.max {
		if val.IsGreaterThan(h.arr[p]) || val.IsLessThan(h.arr[l]) || val.IsLessThan(h.arr[r]) {
			down = false
		}
	} else {
		if val.IsLessThan(h.arr[p]) || val.IsGreaterThan(h.arr[l]) || val.IsGreaterThan(h.arr[r]) {
			down = false
		}
	}
	return true
}

func (h *Heap) bubbleUpMin(child int) (parent int) {
	p := h.getParent(child)
	if h.arr[p].IsGreaterThan(h.arr[child]) { // if parent value is greater than child value
		h.swap(p, child)
		return p
	}
	return -1
}

func (h *Heap) bubbleUpMax(child int) (parent int) {
	p := h.getParent(child)
	if h.arr[p].IsLessThan(h.arr[child]) { // if parent value is less than child value
		h.swap(p, child)
		return p
	}
	return -1
}

func (h *Heap) bubbleUp() {
	curr := len(h.arr)-1 // start with the last item
	for 0 < curr {
		var parent int
		if h.max {
			parent = h.bubbleUpMax(curr) // make bubble check (maximum) and swap
		} else if !h.max {
			parent = h.bubbleUpMin(curr) // make bubble check (minimum) and swap
		}
		curr = parent
	}
}

func (h *Heap) bubbleDownMin(p int) (child int) {
	l, r := h.getChildren(p)
	left, right, parent := h.arr[l], h.arr[r], h.arr[p]
	if (left.IsEqualTo(right)  || left.IsLessThan(right)) && left.IsLessThan(parent) {
		h.swap(l, p) // either tie-case (+ left is smaller than parent) or left is smallest of all three
		return l
	}
	if right.IsLessThan(left) && right.IsLessThan(parent) {
		h.swap(r, p) // right is smallest of all three
		return r
	}
	return -1
}

func (h *Heap) bubbleDownMax(p int) (child int) {
	l, r := h.getChildren(p)
	left, right, parent := h.arr[l], h.arr[r], h.arr[p]
	if (left.IsEqualTo(right)  || left.IsGreaterThan(right)) && left.IsGreaterThan(parent) {
		h.swap(l, p) // either tie-case (+ left is greater than parent) or left is greatest of all three
		return l
	}
	if right.IsGreaterThan(left) && right.IsGreaterThan(parent) {
		h.swap(r, p) // right is greatest of all three
		return r
	}
	return -1
}

func (h *Heap) bubbleDown() {
	curr := 0
	for -1 < curr && curr < len(h.arr)-1 {
		var child int
		if h.max {
			child = h.bubbleDownMax(curr)
		} else {
			child = h.bubbleDownMin(curr)
		}
		curr = child
	}
}

// groups all bubble functions together
func (h *Heap) bubble(idx int) {
	down := h.decideBubble(idx)
	if down {
		h.bubbleDown()
	} else {
		h.bubbleUp()
	}
}
