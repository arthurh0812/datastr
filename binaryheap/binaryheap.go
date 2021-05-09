package binaryheap

import (
	"sync"

	"github.com/arthurh0812/datastr/types"
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

func (h *Heap) replaceTable(key types.Value, toReplace, replaceWith int) {
	indices := h.table[key]
	for j, el := range indices {
		if el == toReplace { // linear search is enough because indices are unique
			h.mu.Lock()
			indices[j] = replaceWith
			h.mu.Unlock()
			return
		}
	}
}

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