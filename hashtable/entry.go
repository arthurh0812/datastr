package hashtable

import (
	"fmt"

	"github.com/arthurh0812/datastruct/types"
)

type Entry struct {
	Key types.Value
	Val interface{}
	Hash int
}

func NewEntry(key types.Value, val interface{}, fn Function) *Entry {
	return &Entry{
		Key: key,
		Val: val,
		Hash: fn(key),
	}
}

func (e *Entry) IsEqualTo(other *Entry) bool {
	if e.Hash != other.Hash { // keys will never match if their hashes don't match
		return false
	}
	return e.Key.IsEqualTo(other.Key)
}

func (e *Entry) String() string {
	return fmt.Sprintf("%v => %v", e.Key, e.Val)
}