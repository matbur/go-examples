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
	want := " 0: []"
	got := ll.String()
	if got != want {
		t.Errorf(errorMessage, got, want)
	}

	ll.Append(3)
	want = " 1: [3]"
	got = ll.String()
	if got != want {
		t.Errorf(errorMessage, got, want)
	}
}

func TestLinkedList_Insert(t *testing.T) {
	ll := NewLinkedList()
	cases := []struct {
		in1, in2 int
		want     string
	}{
		{0, 4, " 1: [4]"},
		{0, 7, " 2: [7, 4]"},
		{0, -1, " 3: [-1, 7, 4]"},
		{0, 0, " 4: [0, -1, 7, 4]"},
		{99, 9, " 5: [0, -1, 7, 4, 9]"},
		{1, 1, " 6: [0, 1, -1, 7, 4, 9]"},
		{-1, 21, " 7: [0, 1, -1, 7, 4, 9, 21]"},
		{-100, 3, " 8: [3, 0, 1, -1, 7, 4, 9, 21]"},
		{-8, 1, " 9: [1, 3, 0, 1, -1, 7, 4, 9, 21]"},
	}
	for _, c := range cases {
		got := ll.Insert(c.in1, c.in2).String()
		if got != c.want {
			t.Errorf("ll.Insert(%q, %q) == %q, want %q", c.in1, c.in2, got, c.want)
		}
	}
}

func TestLinkedList_Append(t *testing.T) {
	ll := NewLinkedList()
	cases := []struct {
		in   int
		want string
	}{
		{4, " 1: [4]"},
		{7, " 2: [4, 7]"},
		{-1, " 3: [4, 7, -1]"},
		{0, " 4: [4, 7, -1, 0]"},
	}
	for _, c := range cases {
		got := ll.Append(c.in).String()
		if got != c.want {
			t.Errorf("ll.Append(%q) == %q, want %q", c.in, got, c.want)
		}
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

func TestLinkedList_Remove(t *testing.T) {
	ll := NewLinkedList()
	for _, i := range []int{1, 2, 3, 4, 5, 3, 4, 1, 1} {
		ll.Append(i)
	}
	cases := []struct {
		in, want1 int
		want2     string
	}{
		{3, 2, " 8: [1, 2, 4, 5, 3, 4, 1, 1]"},
		{17, -1, " 8: [1, 2, 4, 5, 3, 4, 1, 1]"},
		{5, 3, " 7: [1, 2, 4, 3, 4, 1, 1]"},
		{1, 0, " 6: [2, 4, 3, 4, 1, 1]"},
		{1, 4, " 5: [2, 4, 3, 4, 1]"},
		{4, 1, " 4: [2, 3, 4, 1]"},
		{2, 0, " 3: [3, 4, 1]"},
		{1, 2, " 2: [3, 4]"},
		{2, -1, " 2: [3, 4]"},
		{4, 1, " 1: [3]"},
		{3, 0, " 0: []"},
	}
	for _, c := range cases {
		got1 := ll.Remove(c.in)
		if got1 != c.want1 {
			t.Errorf("ll.Remove(%q) == %q, want %q", c.in, got1, c.want1)
		}
		got2 := ll.String()
		if got2 != c.want2 {
			t.Errorf("ll == %q, want %q", got2, c.want2)
		}
	}
}
