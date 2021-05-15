package bst

import (
	"github.com/arthurh0812/datastruct/types"
)

type Node struct {
	val types.Value
	left, right *Node
}

func (n *Node) isLeaf() bool {
	return n != nil && n.left == nil && n.right == nil
}

func (n *Node) hasOnlyLeft() bool {
	return n != nil && n.left != nil && n.right == nil
}

func (n *Node) hasOnlyRight() bool {
	return n != nil && n.right != nil && n.left == nil
}

func (n *Node) hasChildren() bool {
	return n != nil && (n.right != nil || n.left != nil)
}

// n and other mustn't be nil
func (n *Node) isEqualTo(other *Node) bool {
	return n.val.IsEqualTo(other.val)
}

// n and other mustn't be nil
func (n *Node) isGreaterThan(other *Node) bool {
	return n.val.IsGreaterThan(other.val)
}

// n and other mustn't be nil
func (n *Node) isLessThan(other *Node) bool {
	return n.val.IsLessThan(other.val)
}

func (n *Node) isLeft(other *Node) bool {
	return n.left == other
}

func (n *Node) isRight(other *Node) bool {
	return n.right == other
}

func (n *Node) insert(toInsert *Node) {
	if n == nil {
		return
	}
	if n.isEqualTo(toInsert) || n.isGreaterThan(toInsert) {
		n.insertLeft(toInsert)
	} else { // n.isLessThan(toInsert)
		n.insertRight(toInsert)
	}
}

func (n *Node) insertLeft(toInsert *Node) {
	n.left = toInsert
}

func (n *Node) insertRight(toInsert *Node) {
	n.right = toInsert
}

// swaps the values of the two Nodes
func (n *Node) swap(other *Node) {
	n.val, other.val = other.val, n.val
}

func (n *Node) removeChild(child *Node) {
	if n.left == child {
		n.left = nil
	} else if n.right == child {
		n.right = nil
	}
}

func (n *Node) removeLeft() {
	n.left = nil
}

func (n *Node) removeRight() {
	n.right = nil
}

func (n *Node) removeChildren() {
	n.left = nil
	n.right = nil
}

func (n *Node) removeChildAndJoinLeft(child *Node) {
	if n.isLeft(child) {
		n.left = child.left
		child.removeLeft()
	} else if n.isRight(child) {
		n.right = child.left
		child.removeLeft()
	}
}

func (n *Node) removeChildAndJoinRight(child *Node) {
	if n.isLeft(child) {
		n.left = child.right
		child.removeRight()
	} else if n.isRight(child) {
		n.right = child.right
		child.removeRight()
	}
}

// the left child of n must have a left subtree
func (n *Node) removeLeftAndFill() {
	child := n.left
	p, largest := child.findLargestLeft()
	child.swap(largest)
	chooseRemove(p, largest) // largest is always the right child of p
}

// the right child of n must have a left subtree
func (n *Node) removeRightAndFill() {
	child := n.right
	p, largest := child.findLargestRight()
	child.swap(largest)
	chooseRemove(p, largest)
}

// traverses always through subsequent right children of the left subtree of n
func (n *Node) findLargestLeft() (parent, largest *Node) {
	parent = n
	trav := n.left
	for trav != nil {
		if trav.right == nil {
			return parent, trav
		}
		parent = trav
		trav = trav.right
	}
	return nil, nil
}

// traverses always through subsequent right children of the right subtree of n
func (n *Node) findLargestRight() (parent, largest *Node) {
	parent = n
	trav := n.right
	for trav != nil {
		if trav.left == nil {
			return parent, trav
		}
		parent = trav
		trav = trav.left
	}
	return nil, nil
}