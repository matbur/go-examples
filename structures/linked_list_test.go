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