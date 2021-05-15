package binarytree

import (
	"sync"

	"github.com/arthurh0812/datastruct/types"
)

// apfelauto

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

func (t *Tree) Root() (val types.Value) {
	if t.isEmpty() {
		return nil
	}
	return t.root.val
}

func (t *Tree) Size() int64 {
	if t.isEmpty() {
		return 0
	}
	return t.size
}

func (t *Tree) isEmpty() bool {
	return t == nil || t.root == nil || t.size == 0
}

func (t *Tree) IsEmpty() bool {
	return t.isEmpty()
}

func (t *Tree) insert(n *node) {

}

func (t *Tree) Insert(val types.Value) {
	n := &node{
		val: val,
	}
	t.insert(n)
}

func (t *Tree) InsertMany(vals []types.Value) {
	for _, val := range vals {
		t.Insert(val)
	}
}