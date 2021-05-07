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

func TestLinkedList_Prepend(t *testing.T) {
	list := Empty()

	input := []int{4, 3, 6, 18, 32, 64, 37}
	output := []int{37, 64, 32, 18, 6, 3, 4}

	for _, n := range input {
		list.Prepend(n)
	}

	got := list.Values()

	if gotl, inputl := len(got), len(input); gotl != inputl {
		t.Fatalf("lenghts do not match: len(got) = %d, len(input) = %d", gotl, inputl)
	}

	for i, n := range list.Values() {
		if output[i] != n {
			t.Errorf("number %d: incorrect value: %v", i+1, n)
		}
	}
}

func TestLinkedList_RemoveHead(t *testing.T) {
	list := New("hello", Empty())

	list.Append("moin")
	list.Append("everything's fine")
	list.Prepend("not so ok")
	list.Prepend(75)
	list.Append(true)
	list.Prepend("i like")

	v := list.RemoveHead()
	if v != "i like" {
		t.Errorf("incorrect head value: %v", v)
	}

	t.Log(list.String())

	list.Clear() // list gets empty again

	v = list.RemoveHead()
	if v != nil {
		t.Errorf("incorrect head value: %v", v)
	}

	t.Log(list.String())
}