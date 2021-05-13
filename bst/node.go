package bst

import (
	"github.com/arthurh0812/datastruct/types"
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

func (n *node) hasChildren() bool {
	return n != nil && (n.right != nil || n.left != nil)
}

// n and other mustn't be nil
func (n *node) isEqualTo(other *node) bool {
	return n.val.IsEqualTo(other.val)
}

// n and other mustn't be nil
func (n *node) isGreaterThan(other *node) bool {
	return n.val.IsGreaterThan(other.val)
}

// n and other mustn't be nil
func (n *node) isLessThan(other *node) bool {
	return n.val.IsLessThan(other.val)
}

func (n *node) isLeft(other *node) bool {
	return n.left == other
}

func (n *node) isRight(other *node) bool {
	return n.right == other
}

func (n *node) insert(toInsert *node) {
	if n == nil {
		return
	}
	if n.isEqualTo(toInsert) || n.isGreaterThan(toInsert) {
		n.insertLeft(toInsert)
	} else { // n.isLessThan(toInsert)
		n.insertRight(toInsert)
	}
}

func (n *node) insertLeft(toInsert *node) {
	n.left = toInsert
}

func (n *node) insertRight(toInsert *node) {
	n.right = toInsert
}

// swaps the values of the two nodes
func (n *node) swap(other *node) {
	n.val, other.val = other.val, n.val
}

func (n *node) removeChild(child *node) {
	if n.left == child {
		n.left = nil
	} else if n.right == child {
		n.right = nil
	}
}

func (n *node) removeLeft() {
	n.left = nil
}

func (n *node) removeRight() {
	n.right = nil
}

func (n *node) removeChildren() {
	n.left = nil
	n.right = nil
}

func (n *node) removeChildAndJoinLeft(child *node) {
	if n.isLeft(child) {
		n.left = child.left
		child.removeLeft()
	} else if n.isRight(child) {
		n.right = child.left
		child.removeLeft()
	}
}

func (n *node) removeChildAndJoinRight(child *node) {
	if n.isLeft(child) {
		n.left = child.right
		child.removeRight()
	} else if n.isRight(child) {
		n.right = child.right
		child.removeRight()
	}
}

// the left child of n must have a left subtree
func (n *node) removeLeftAndFill() {
	child := n.left
	p, largest := child.left.findLargest()
	child.swap(largest)
	p.removeRight() // largest is always the right child of p
}

// the right child of n must have a left subtree
func (n *node) removeRightAndFill() {
	child := n.right
	p, largest := child.left.findLargest()
	child.swap(largest)
	p.removeRight()
}

// traverses always through subsequent right children of n (to find the largest subsequent node)
func (n *node) findLargest() (parent, largest *node) {
	trav := n
	for trav != nil {
		if trav.right == nil {
			return parent, trav
		}
		parent = trav
		trav = trav.right
	}
	return nil, nil
}

// traverses always through subsequent left children of n (to find the smallest subsequent node)
func (n *node) findSmallest() (parent, smallest *node) {
	trav := n
	for trav != nil {
		if trav.left == nil {
			return parent, trav
		}
		parent = trav
		trav = trav.left
	}
	return nil, nil
}