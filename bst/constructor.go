package bst

import "github.com/arthurh0812/datastruct/types"

func Empty() *Tree {
	return &Tree{
		size: 0,
		root: nil,
	}
}

func New(vals ...types.Value) *Tree {
	t := Empty()
	for _, val := range vals {
		t.Insert(val)
	}
	return t
}
