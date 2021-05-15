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
	if tree.Contains(vals[1]) {
		t.Fatalf("value %d should not exist inside tree", vals[1])
	}

	tree.Remove(datastruct.NewInt(29))
	if got, want := tree.Size(), int64(len(vals))-2; got != want {
		t.Fatalf("incorrect tree size: got %d, want %d", got, want)
	}
	if tree.Contains(datastruct.NewInt(29)) {
		t.Fatalf("value %d should not exist inside tree", vals[1])
	}
}

func TestTree_InOrder(t *testing.T) {
	vals := datastruct.NewIntSlice([]interface{}{40, 50, 23, 17, 78, 29, 4, 32, 49})

	tree := New(vals...)

	inOrder := make([]interface{}, 0, len(vals))

	tree.InOrder(func(n *Node) {
		inOrder = append(inOrder, n.val)
	})

	t.Logf("%v\n", inOrder)
}

func TestTree_PreOrder(t *testing.T) {
	vals := datastruct.NewIntSlice([]interface{}{40, 50, 23, 17, 78, 29, 4, 32, 49})

	tree := New(vals...)

	preOrder := make([]interface{}, 0, len(vals))

	tree.PreOrder(func(n *Node) {
		preOrder = append(preOrder, n.val)
	})

	t.Logf("%v\n", preOrder)
}

func TestTree_PostOrder(t *testing.T) {
	vals := datastruct.NewIntSlice([]interface{}{40, 50, 23, 17, 78, 29, 4, 32, 49})

	tree := New(vals...)

	postOrder := make([]interface{}, 0, len(vals))

	tree.PostOrder(func(n *Node) {
		postOrder = append(postOrder, n.val)
	})

	t.Logf("%v\n", postOrder)
}

func TestTree_LevelOrder(t *testing.T) {
	vals := datastruct.NewIntSlice([]interface{}{40, 50, 23, 17, 78, 29, 4, 32, 49})

	tree := New(vals...)

	lvlOrder := make([]interface{}, 0, len(vals))

	tree.LevelOrder(func(n *Node) {
		lvlOrder = append(lvlOrder, n.val)
	})

	t.Logf("%v\n", lvlOrder)
}