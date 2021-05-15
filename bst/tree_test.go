package bst

import (
	"github.com/arthurh0812/datastruct"
	"testing"
)

func TestTree_Clear(t *testing.T) {
	tree := New()

	vals := datastruct.NewIntSlice([]interface{}{8, 90, 30, 2, 34})
	for _, val := range vals {
		tree.Insert(val)
	}

	 if got, want := tree.Size(), int64(len(vals)); got != want {
	 	t.Fatalf("incorrect tree length: got %d, want %d", got, want)
	 }
	 tree.Clear()
	 if !tree.IsEmpty() {
	 	t.Fatalf("tree should be empty")
	 }
	 if got, want := tree.Size(), int64(0); got != want {
	 	t.Fatalf("empty tree should have a size of %d: got %d", want, got)
	 }
}

func TestTree_Insert_Remove(t *testing.T) {
	vals := datastruct.NewIntSlice([]interface{}{40, 50, 23, 17, 78, 29, 4, 32, 49})

	tree := New(vals...)

	tree.Remove(vals[1])

	if got, want := tree.Size(), int64(len(vals))-1; got != want {
		t.Fatalf("incorrect tree size: got %d, want %d", got, want)
	}

	if tree.Exists(vals[1]) {
		t.Fatalf("value %d should not exist inside tree", vals[1])
	}
}