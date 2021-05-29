package hashtable

import (
	"errors"
	"github.com/arthurh0812/datastruct/linkedlist"
)

var DefaultCapacity int64 = 3
var DefaultLoadFactor float64 = 0.75

var ErrInvalidKeyType = errors.New("the provided key type is not valid")

// Default creates and returns a new default hash map (= map[string]interface{})
func Default() *HashTable {
	h := &HashTable{}
	h.setTable(make([]*linkedlist.LinkedList, DefaultCapacity, 0))
	h.setCapacity(DefaultCapacity)
	h.setLoadFactor(DefaultLoadFactor)
	h.setFunction(GetFunction("string"))
	return h
}

func New(keyType string, cap int64) (*HashTable, error) {
	h := &HashTable{}
	err := h.initCapacity(cap)
	if err != nil {
		return nil, err
	}
	fn := GetFunction(keyType)
	if fn == nil {
		return nil, ErrInvalidKeyType
	}
	h.setFunction(fn)
	h.setTable(make([]*linkedlist.LinkedList, cap, 0))
	h.setLoadFactor(DefaultLoadFactor)
	return h, nil
}