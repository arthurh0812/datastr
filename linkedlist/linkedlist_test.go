package linkedlist

import "testing"

func TestLinkedList_IsEmpty(t *testing.T) {
	tests := []struct{
		isEmpty bool
		list *LinkedList
	}{
		{
			isEmpty: false,
			list: New(3, Empty()),
		},
		{
			isEmpty: true,
			list: Empty(),
		},
	}

	for i, test := range tests {
		if got, exp := test.list.IsEmpty(), test.isEmpty; got != exp {
			t.Errorf("Test No.%d: expected %t, got %v", i+1, exp, got)
		}
	}
}

func TestLinkedList_Append(t *testing.T) {
	list := Empty()

	numbers := []int{4, 3, 10, 4832, 78, 929}
	for _, n := range numbers {
		list.Append(n)
	}

	got := list.Values()

	if len(got) != len(numbers) {
		t.Fatalf("len(got) = %d != len(numbers) = %d", len(got), len(numbers))
	}

	for i, n := range list.Values() {
		if numbers[i] != n {
			t.Errorf("number %d: incorrect value: %d", i+1, n)
		}
	}
}