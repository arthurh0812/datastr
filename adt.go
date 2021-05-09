package datastruct

import "github.com/arthurh0812/datastruct/types"

// ADT is the interface that all kinds of abstract data types implement
type ADT interface {
	IsEmpty() bool
	Clear()
}

// LinkedList Abstract Data Type
type LinkedList interface {
	ADT // implements Abstract Data Type
	Head() (val interface{})
	Tail() (val interface{})
	Append(val interface{})
	Prepend(val interface{})
	RemoveHead() (val interface{})
	RemoveTail() (val interface{})
}

// Stack Abstract Data Type
type Stack interface {
	ADT // implements Abstract Data Type
	First() (val interface{})
	Last() (val interface{})
	Push(val interface{})
	Pop() (val interface{})
}

// Queue Abstract Data Type
type Queue interface {
	ADT // implements Abstract Data Type
	Peek() (val interface{})
	Last() (val interface{})
	Enqueue(val interface{})
	Dequeue() (val interface{})
}

// PriorityQueue Abstract Data Type
type PriorityQueue interface {
	ADT // implements Abstract Data Type
	Peek() types.Value
	Poll() types.Value
	Insert(val types.Value)
	Remove(val types.Value)
	MakeMin()
	MakeMax()
}

// Tree Abstract Data Type
type Tree interface {
	Root() types.Value
	Size() int64
}

// UnionFind Abstract Data Type
type UnionFind interface {

}