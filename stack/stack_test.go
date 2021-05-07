package stack

import "testing"

func TestStack_Push(t *testing.T) {
	stack := Empty()

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	output := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	for _, n := range input {
		stack.Push(n)
	}

	t.Log(stack.String())

	for _, n := range output {
		if got := stack.Pop(); got != n {
			t.Errorf("incorrect value: %v", got)
		}
	}
}
