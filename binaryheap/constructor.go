package binaryheap

import "github.com/arthurh0812/datastr/types"

// Empty creates a new empty Heap.
func Empty() *Heap {
	return &Heap{
		arr: make([]types.Value, 0, 0),
		table: make(map[types.Value][]int, 0),
		max: false,
		size: 0,
	}
}

// New creates a new minimum Heap with the given values inserted into it.
func New(vals ...types.Value) *Heap {
	h := Empty()
	h.InsertAll(vals)
	return h
}

// NewMax creates a new maximum Heap with the given values inserted into it.
func NewMax(vals ...types.Value) *Heap {
	h := Empty()
	h.MakeMax()
	h.InsertAll(vals)
	return h
}
