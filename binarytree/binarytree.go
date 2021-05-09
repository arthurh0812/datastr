package binarytree

import (
	"sync"

	"github.com/arthurh0812/datastruct/types"
)

type node struct {
	val types.Value
	left, right *node
	subSize int64
}

type Tree struct {
	root *node
	size int64

	mu sync.Mutex
}

func (t *Tree) Insert(val types.Value) {

}

func (t *Tree) InsertMany(vals []types.Value) {
	for _, val := range vals {
		t.Insert(val)
	}
}
