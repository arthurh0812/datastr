package datastr

type btNode struct {
	val interface{}
	left, right *btNode
}

type BinaryTree struct {
	root *btNode
	size int64
}

func NewBinaryTree(val interface{}, left, right *BinaryTree) *BinaryTree {
	subSize := getSubSize(left, right)
	lRoot, rRoot := getRoots(left, right)
	return &BinaryTree{
		root: &btNode{val: val, left: lRoot, right: rRoot},
		size: 1 + subSize,
	}
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