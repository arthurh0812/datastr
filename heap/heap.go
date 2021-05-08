package heap

import "github.com/arthurh0812/datastr/types"

type Heap struct {
	arr []types.Value
	max bool // default=min
}

func (h *Heap) isEmpty() bool {
	return h == nil || len(h.arr) == 0
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

func (h *Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
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
		h.swap(l, p) // tie-case or left is smallest of all three
		return l
	}
	if right.IsLessThan(left) && right.IsLessThan(parent) {
		h.swap(r, p)
		return r
	}
	return -1
}

func (h *Heap) bubbleDownMax(p int) (child int) {
	l, r := h.getChildren(p)
	left, right, parent := h.arr[l], h.arr[r], h.arr[p]
	if (left.IsEqualTo(right)  || left.IsGreaterThan(right)) && left.IsGreaterThan(parent) {
		h.swap(l, p) // tie-case or left is smallest of all three
		return l
	}
	if right.IsGreaterThan(left) && right.IsGreaterThan(parent) {
		h.swap(r, p)
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

func (h *Heap) Insert(val types.Value) {
	// append the new value to the
	h.arr = append(h.arr, val)
	h.bubbleUp() // reorganize heap (upwards)
}

func (h *Heap) poll() (first types.Value) {
	first, last := h.arr[0], h.arr[len(h.arr)-1]
	h.arr[0] = last
	return first
}

func (h *Heap) Poll() (val types.Value) {
	 if h.isEmpty() {
	 	return nil
	 }
	 v := h.poll()
	 h.bubbleDown() // reorganize heap (downwards)
	 return v
}

func (h *Heap) Peek() (val types.Value) {
	if h.isEmpty() {
		return nil
	}
	return h.arr[0]
}

// MakeMin makes this heap a minimum priority queue; this is the default.
func (h *Heap) MakeMin() {
	h.max = false
}

// MakeMax makes this heap a maximum priority queue.
func (h *Heap) MakeMax() {
	h.max = true
}