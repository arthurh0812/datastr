package unionfind

import "sync"

type UnionFind struct {
	parents []int // parents[node_idx] -> parent_idx

	setSizes []int64 // setSizes[root_idx] -> set_size of that root

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

func (u *UnionFind) isEmpty() bool {
	return u == nil ||
		u.size == 0 ||
		u.numSets == 0 ||
		len(u.parents) == 0 ||
		len(u.setSizes) == 0
}

func (u *UnionFind) IsEmpty() bool {
	return u.isEmpty()
}

func (u *UnionFind) increaseNumSets() {
	u.mu.Lock()
	u.numSets++
	u.mu.Unlock()
}

func (u *UnionFind) decreaseNumSets() {
	u.mu.Lock()
	u.numSets--
	u.mu.Unlock()
}

// merge two sets together
func (u *UnionFind) unify(first, second int) {
	firstRoot := u.find(first)
	secondRoot := u.find(second)
	if firstRoot == secondRoot { // already in the same set
		return
	}
	if u.setSizes[first] < u.setSizes[second] { // the second set is bigger
		u.setSizes[second] += u.setSizes[first] // add the size of the first to second
		u.parents[first] = secondRoot
	} else { // the first set is bigger or equal to
		u.setSizes[first] += u.setSizes[second]
		u.parents[second] = firstRoot
	}
	u.decreaseNumSets()
}

func (u *UnionFind) Unify(first, second int) {
	u.unify(first, second)
}

func (u *UnionFind) findRoot(node int) (root int){
	root = node
	for root != u.parents[root] { // go as far as possible to root node
		root = u.parents[root]
	}
	return root
}

func (u *UnionFind) compressPath(node, root int) {
	trav := node
	for trav != root {
		next := u.parents[trav]
		u.parents[trav] = root
		trav = next
	}
}

// return index of root node
func (u *UnionFind) find(node int) (root int) {
	root = u.findRoot(node)
	u.compressPath(node, root)
	return root
}

// Find finds corresponding root node, compresses path and returns that root node.
func (u *UnionFind) Find(node int) (root int) {
	return u.find(node)
}

func (u *UnionFind) AreConnected(first, second int) bool {
	return u.find(first) == u.find(second) // both nodes have to have the same root node
}

func (u *UnionFind) ComponentSize(node int) int64 {
	return u.setSizes[u.find(node)]
}

func (u *UnionFind) clear() {
	u.mu.Lock()
	u.numSets = 0
	u.size = 0
	u.setSizes = make([]int64, 0, 0)
	u.parents = make([]int, 0, 0)
	u.mu.Unlock()
}

func (u *UnionFind) Clear() {
	u.clear()
}