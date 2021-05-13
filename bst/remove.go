package bst

import "github.com/arthurh0812/datastruct/types"

func (t *Tree) remove(n *node) {
	parent, child := t.findPre(n)
	if parent == nil || child == nil { // root or not found
		return
	}
	t.chooseRemove(parent, child)
	t.decreaseSize()
}

func (t *Tree) chooseRemove(parent, child *node) {
	if child.isLeaf() {
		parent.removeChild(child)
		return
	}
	if child.hasOnlyLeft() {
		parent.removeChildAndJoinLeft(child)
		return
	}
	if child.hasOnlyRight() {
		parent.removeChildAndJoinRight(child)
		return
	}
	// the child has both a left and a right subtree
	if parent.isLeft(child) {
		parent.removeLeftAndFill()
	} else if parent.isRight(child) {
		parent.removeRightAndFill()
	}
}

// Remove tries to find a node that holds val and removes it from the tree.
func (t *Tree) Remove(val types.Value) {
	n := &node{val: val}
	t.remove(n)
}