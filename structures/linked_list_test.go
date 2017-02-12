package structures

import "testing"

func TestNewNode(t *testing.T) {
	cases := []struct {
		in   int
		want *Node
	}{
		{0, &Node{0, nil}},
		{1, &Node{1, nil}},
		{2, &Node{2, nil}},
		{-1, &Node{-1, nil}},
	}

	for _, c := range cases {
		if got := NewNode(c.in); *got != *c.want {
			t.Errorf("NewNode(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestNewLinkedList(t *testing.T) {
	want := &LinkedList{nil, 0}
	got := NewLinkedList()
	if *got != *want {
		t.Errorf("NewLinkedList() == %q, want %q", got, want)
	}
}

func TestLinkedList_String(t *testing.T) {
	errorMessage := "NewLinkedList().String() == %q, want %q"

	ll := NewLinkedList()
	want := " 0: {}"
	got := ll.String()
	if got != want {
		t.Errorf(errorMessage, got, want)
	}

	ll.Append(3)
	want = " 1: {3}"
	got = ll.String()
	if got != want {
		t.Errorf(errorMessage, got, want)
	}
}

func TestLinkedList_Index(t *testing.T) {
	ll := NewLinkedList().Append(0).Append(1).Append(-1)
	cases := []struct {
		in, want int
	}{
		{0, 0},
		{-1, 2},
		{1, 1},
		{2, -1},
	}
	for _, c := range cases {
		got := ll.Index(c.in)
		if got != c.want {
			t.Errorf("ll.Index(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
