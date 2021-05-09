package binarytree

import (
	"sync"

	"github.com/arthurh0812/datastruct/types"
)

// Empty creates an empty binary tree.
func Empty() *Tree {
	return &Tree{
		root: nil,
		size: 0,
		mu: sync.Mutex{},
	}
}

// New creates a new binary tree with the given values inserted into it.
func New(vals ...types.Value) *Tree {
	t := Empty()
	t.InsertMany(vals)
	return t
}