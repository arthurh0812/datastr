package unionfind

func Empty() *UnionFind {
	return &UnionFind{
		size: 0,
		parents: make([]int64, 0, 0),
		setSizes: make([]int64, 0, 0),
	}
}

// New creates a new union find data structure.
func New(size int64) *UnionFind {
	if size <= 0 {
		return Empty()
	}
	u := &UnionFind{}
	u.setSize(size)
	u.setNumSets(size) // each element is a component in the beginning, so size = numSets
	u.parents = make([]int64, size)
	u.setSizes = make([]int64, size)
	for i := int64(0); i < size; i++ {
		u.parents[i] = i
		u.setSizes[i] = 1
	}
	return u
}
