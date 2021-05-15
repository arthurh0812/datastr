package bst

import (
	"github.com/arthurh0812/datastruct/queue"
	"github.com/arthurh0812/datastruct/stack"
)

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

// InOrder traverses through the binary tree in sorting order.
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

// PostOrder is a more complex traversal.
func (t *Tree) PostOrder(cb func(node *Node)) {
	trav, prev := t.root, (*Node)(nil)
	nodeStack := stack.Empty()
	for trav != nil || !nodeStack.IsEmpty() { // traverse right-wards
		if trav != nil { // traverse leftwards
			nodeStack.Push(trav)
			trav = trav.left
			continue
		}
		last := nodeStack.Last().(*Node)
		if last.right == nil || last.right == prev {
			nodeStack.Pop()
			cb(last)
			prev, trav = last, nil
		} else {
			trav = last.right
		}
	}
}

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