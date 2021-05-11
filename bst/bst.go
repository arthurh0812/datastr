package bst

import (
	"github.com/arthurh0812/datastruct/types"
	"sync"
)

type node struct {
	val types.Value
	left, right *node
}

func (n *node) isLeaf() bool {
	return n != nil && n.left == nil && n.right == nil
}

func (n *node) hasOnlyLeft() bool {
	return n != nil && n.left != nil && n.right == nil
}

func (n *node) hasOnlyRight() bool {
	return n != nil && n.right != nil && n.left == nil
}

type Tree struct {
	root *node
	size int64

	duplicate bool

	mu sync.Mutex
}

func (t *Tree) Size() int64 {
	return t.size
}

func (t *Tree) setSize(s int64) {
	t.mu.Lock()
	t.size = s
	t.mu.Lock()
}

func (t *Tree) increaseSize() {
	t.mu.Lock()
	t.size++
	t.mu.Unlock()
}

func (t *Tree) decreaseSize() {
	t.mu.Lock()
	t.size--
	t.mu.Unlock()
}

func (t *Tree) isEmpty() bool {
	return t == nil || t.size == 0 || t.root == nil
}

func (t *Tree) IsEmpty() bool {
	return t.isEmpty()
}

func (t *Tree) setRoot(n *node) {
	t.mu.Lock()
	t.root = n
	t.mu.Unlock()
}

func (t *Tree) allowDuplicate() {
	t.mu.Lock()
	t.duplicate = true
	t.mu.Unlock()
}

func (t *Tree) forbidDuplicate() {
	t.mu.Lock()
	t.duplicate = false
	t.mu.Unlock()
}

func (t *Tree) traverse(val types.Value) (pre *node) {
	trav := t.root
	for trav != nil {
		if val.IsGreaterThan(trav.val) && trav.right != nil {
			trav = trav.right
		} else if (val.IsLessThan(trav.val) || val.IsEqualTo(trav.val)) && trav.left != nil {
			trav = trav.left
		} else {
			return trav
		}
	}
	return nil // should never happen as traverse is only called if tree is not empty
}

// O(log(n)) average time complexity
func (t *Tree) find(val types.Value) (n *node) {
	trav := t.root
	for trav != nil {
		switch comparator(val, trav.val) {
		case 0:
			return trav
		case -1:
			trav = trav.left
		case 1:
			trav = trav.right
		}
	}
	return nil
}

func (t *Tree) findPre(val types.Value) (prev, found *node) {
	trav, prev := t.root, nil
	for trav != nil {
		switch comparator(val, trav.val) {
		case 0:
			return prev, trav
		case -1:
			prev = trav
			trav = trav.left
		case 1:
			prev = trav
			trav = trav.right
		}
	}
	return prev, trav
}

func (t *Tree) appendToNode(pre, toAdd *node) {
	if toAdd.val.IsLessThan(pre.val) || pre.val.IsEqualTo(toAdd.val) {
		pre.left = toAdd
	} else {
		pre.right = toAdd
	}
}

func (t *Tree) insert(n *node) {
	pre := t.traverse(n.val)
	if !t.duplicate && n.val.IsEqualTo(pre.val) { // return if no duplicates are allowed
		return
	}
	t.appendToNode(pre, n)
	t.increaseSize()
}

func (t *Tree) Insert(val types.Value) {
	n := &node{val: val}
	if t.isEmpty() {
		t.init(n)
		return
	}
	t.insert(n)
}

func (t *Tree) init(n *node) {
	t.setRoot(n)
	t.setSize(1)
}

func (t *Tree) clear() {
	t.mu.Lock()
	t.root = nil
	t.size = 0
	t.mu.Unlock()
}

func (t *Tree) Clear() {
	t.clear()
}

func comparator(first, second types.Value) int {
	if first.IsEqualTo(first) {
		return 0
	} else if first.IsGreaterThan(second) {
		return 1
	} else {
		return -1
	}
}