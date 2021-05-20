package hashtable

import "github.com/arthurh0812/datastruct/linkedlist"

var DefaultCapacity int64 = 3
var DefaultLoadFactor float64 = 0.75

func Default() *HashTable {
	h := &HashTable{}
	h.setTable(make([]*linkedlist.LinkedList, DefaultCapacity, 0))
	h.setCapacity(DefaultCapacity)
	h.setLoadFactor(DefaultLoadFactor)
	h.setFunction(DefaultFunction)
	return h
}

func New(cap int64) (*HashTable, error) {
	h := &HashTable{}
	err := h.initCapacity(cap)
	if err != nil {
		return nil, err
	}
	h.setFunction(DefaultFunction)
	h.setTable(make([]*linkedlist.LinkedList, cap, 0))
	h.setLoadFactor(DefaultLoadFactor)
	return h, nil
}