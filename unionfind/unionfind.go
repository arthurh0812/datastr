package unionfind

import "sync"

type UnionFind struct {
	// parents maps indices to their parent nodes
	parents []int64 // keeps track of the parent node
	// setSizes maps root indices to the sizes of the component that they are the root of
	setSizes []int64

	size int64 // number of elements in union find
	numSets int64 // number of sets in union find

	mu sync.Mutex
}

func (u *UnionFind) Size() int64 {
	return u.size
}

func (u *UnionFind) setSize(s int64) {
	u.mu.Lock()
	u.size = s
	u.mu.Unlock()
}

func (u *UnionFind) NumSets() int64 {
	return u.numSets
}

func (u *UnionFind) setNumSets(n int64) {
	u.mu.Lock()
	u.numSets = n
	u.mu.Unlock()
}

func (u *UnionFind) Unify() {

}

func (u *UnionFind) findRoot(node int64) (root int64){
	root = node
	for root != u.parents[root] { // go as far as possible to root node
		root = u.parents[root]
	}
	return root
}

func (u *UnionFind) compressPath(node, root int64) {
	trav := node
	for trav != root {
		next := u.parents[trav]
		u.parents[trav] = root
		trav = next
	}
}

func (u *UnionFind) find(node int64) (root int64) {
	root = u.findRoot(node)
	u.compressPath(node, root)
	return root
}

// Find finds corresponding root node, compresses path and returns that root node.
func (u *UnionFind) Find(node int64) (root int64) {
	return u.find(node)
}


func (u *UnionFind) areConnected(first, second int64) bool {
	return u.find(first) == u.find(second) // both nodes have to have the same root node
}

func (u *UnionFind) ComponentSize(node int64) int64 {
	return u.setSizes[u.find(node)]
}