package bst

import "github.com/arthurh0812/datastruct/types"

func (t *Tree) removeLeaf(prev, toRemove *node) {
	if toRemove == prev.left {
		prev.left = nil
	} else if toRemove == prev.right {
		prev.right = nil
	}
}

func (t *Tree) removeAndAddLeftSubtree(prev, toRemove *node) {
	if toRemove == prev.left {
		prev.left = toRemove.left
	} else if toRemove == prev.right {
		prev.right = toRemove.left
	}
	toRemove.left = nil
	toRemove.right = nil
}

func (t *Tree) removeAndAddRightSubtree(prev, toRemove *node) {
	if toRemove == prev.left {
		prev.left = toRemove.right
	} else if toRemove == prev.right {
		prev.right = toRemove.right
	}
	toRemove.left = nil
	toRemove.right = nil
}

func (t *Tree) remove(val types.Value) {
	prev, n := t.findPre(val)
	if prev == nil { // root
		return
	}
	if n.isLeaf() {
		t.removeLeaf(prev, n)
		return
	}
	if n.hasOnlyLeft() {
		t.removeAndAddLeftSubtree(prev, n)
	}
	if n.hasOnlyRight() {
		t.removeAndAddRightSubtree(prev, n)
	}
	// case where there is both a left and a right subtree
}
