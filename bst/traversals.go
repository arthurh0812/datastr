package bst

import (
	"github.com/arthurh0812/datastruct/queue"
	"github.com/arthurh0812/datastruct/stack"
)

// PreOrder is a traversal operation that iteratively calls the cb function on each each
// top node, then traverses the left subtree of that node and after that its right subtree.
func (t *Tree) PreOrder(cb func(node *Node)) {
	nodeStack := stack.New(t.root, stack.Empty())
	for !nodeStack.IsEmpty() {
		last := nodeStack.Pop().(*Node)
		cb(last)
		if last.right != nil {
			nodeStack.Push(last.right)
		}
		if last.left != nil {
			nodeStack.Push(last.left)
		}
	}
}

// InOrder is a traversal operation that iteratively calls the cb function on each node
// of the tree in increasing order.
func (t *Tree) InOrder(cb func(node *Node)) {
	trav := t.root
	nodeStack := stack.Empty()
	for trav != nil || !nodeStack.IsEmpty() { // traverse right-wards
		if trav != nil { // traverse leftwards
			nodeStack.Push(trav)
			trav = trav.left
			continue
		}
		last := nodeStack.Pop().(*Node)
		cb(last)
		trav = last.right
	}
}

// PostOrder is a traversal operation that iteratively traverses the left and then the right
// subtree calling the cb on each node, and only after those two operations calls the cb
// function on the top node.
func (t *Tree) PostOrder(cb func(node *Node)) {
	trav, prev := t.root, (*Node)(nil)
	nodeStack := stack.Empty()
	for trav != nil || !nodeStack.IsEmpty() { // traverse right-wards
		if trav != nil { // traverse leftwards
			nodeStack.Push(trav)
			trav = trav.left
			continue
		} // up to this point trav must be nil here
		last := nodeStack.Last().(*Node)
		if last.right != nil && last.right != prev {
			// go right if there is a right node and if that node is not the previously most recently popped node
			trav = last.right
			continue
		}
		nodeStack.Pop()
		cb(last)
		prev, trav = last, nil
	}
}

// LevelOrder is a BFS-algorithm that iteratively calls the given cb function on each
// traversed node, from left to right, level by level.
func (t *Tree) LevelOrder(cb func(node *Node)) {
	nodeQueue := queue.New(t.root, queue.Empty())
	for !nodeQueue.IsEmpty() {
		parent := nodeQueue.Dequeue().(*Node)
		cb(parent)
		if parent.left != nil {
			nodeQueue.Enqueue(parent.left)
		}
		if parent.right != nil {
			nodeQueue.Enqueue(parent.right)
		}
	}
}