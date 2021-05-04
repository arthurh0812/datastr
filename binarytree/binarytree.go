package binarytree

type btNode struct {
	val interface{}
	left, right *btNode
}

type BinaryTree struct {
	root *btNode
	size int64
}

func leftBT(bt *BinaryTree) *BinaryTree {
	leftBT := &BinaryTree{
		root: bt.root.left,
	}
	leftBT.size = size(leftBT)
	return leftBT
}

func rightBT(bt *BinaryTree) *BinaryTree {
	 rightBT := &BinaryTree{
		root: bt.root.right,
	}
	rightBT.size = size(rightBT)
	return rightBT
}

var EmptyBinaryTree *BinaryTree

func NewBinaryTree(val interface{}, left, right *BinaryTree) *BinaryTree {
	subSize := getSubSize(left, right)
	lRoot, rRoot := getRoots(left, right)
	return &BinaryTree{
		root: &btNode{val: val, left: lRoot, right: rRoot},
		size: 1 + subSize,
	}
}

func size(bt *BinaryTree) int64 {
	size := int64(0)
	for curr := bt.root; curr != nil; curr = curr.right {
		for curr != nil {
			size++
			curr = curr.left
		}
		size++
	}
	return size
}

func getSubSize(left, right *BinaryTree) int64 {
	var leftSize, rightSize int64
	if left != nil {
		leftSize = left.size
	}
	if right != nil {
		rightSize = right.size
	}
	return leftSize + rightSize
}

func getRoots(left, right *BinaryTree) (lRoot, rRoot *btNode) {
	if left != nil {
		lRoot = left.root
	}
	if right != nil {
		rRoot = right.root
	}
	return
}

func IsEmptyBT(bt *BinaryTree) bool {
	return bt == nil
}

func (bst *BinaryTree) InsertBST(val interface{}) *BinaryTree {
	if IsEmptyBT(bst) {
		return NewBinaryTree(val, EmptyBinaryTree, EmptyBinaryTree)
	}
	if val.(int64) < bst.root.val.(int64) || val.(string) < bst.root.val.(string) {
		return NewBinaryTree(bst.root, leftBT(bst).InsertBST(val), rightBT(bst))
	}
	if val.(int64) > bst.root.val.(int64) || val.(string) > bst.root.val.(string) {
		return NewBinaryTree(bst.root, leftBT(bst), rightBT(bst).InsertBST(val))
	}
	panic("violated assumption in procedure 'InsertBST'")
}