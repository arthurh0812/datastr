package bst

import (
	"github.com/arthurh0812/datastruct/queue"
	"github.com/arthurh0812/datastruct/stack"
)

func (t *Tree) PreOrder(cb func(node *Node)) {
	trav := t.root
	nodeStack := stack.Empty()
	for trav != nil {
		nodeStack.Push(trav)
		cb(trav)
		trav = trav.left
	}
	for !nodeStack.IsEmpty() {
		last := nodeStack.Pop().(*Node)
		trav = last.right
		for trav != nil {
			nodeStack.Push(trav)
			cb(trav)
			trav = trav.left
		}
	}
}

func (t *Tree) InOrder(cb func(node *Node)) {
	trav := t.root
	nodeStack := stack.Empty()
	for trav != nil {
		nodeStack.Push(trav)
		trav = trav.left
	}
	for !nodeStack.IsEmpty() {
		last := nodeStack.Pop().(*Node)
		cb(last) // call the callback on the node
		trav = last.right
		for trav != nil {
			nodeStack.Push(trav)
			trav = trav.left
		}
	}
}

func (t *Tree) PostOrder(cb func(node *Node)) {
	//trav := t.root
	//nodeStack := stack.Empty()
	//for trav != nil {
	//	nodeStack.Push(trav)
	//	trav = trav.left
	//}
	//for !nodeStack.IsEmpty() {
	//	last := nodeStack.Pop().(*Node)
	//	cb(last)
	//	trav = last.right
	//	for trav != nil {
	//		nodeStack.Push(trav)
	//		trav = trav.left
	//	}
	//}
}

func (t *Tree) LevelOrder(cb func(node *Node)) {
	trav := t.root
	nodeQueue := queue.New(trav, queue.Empty())
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