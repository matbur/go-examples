package double_linked_list

import "testing"

func TestNewNode(t *testing.T) {
	cases := []struct {
		in   int
		want *node
	}{
		{0, &node{0, nil, nil}},
		{1, &node{1, nil, nil}},
		{2, &node{2, nil, nil}},
		{-1, &node{-1, nil, nil}},
	}

	for _, c := range cases {
		if got := newNode(c.in); *got != *c.want {
			t.Errorf("newNode(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestNewDoubleLinkedList(t *testing.T) {
	want := &DoubleLinkedList{nil, nil, 0}
	got := NewDoubleLinkedList()
	if *got != *want {
		t.Errorf("NewDoubleLinkedList() == %q, want %q", got, want)
	}
}

func TestDoubleLinkedList_AppendRigth(t *testing.T) {
	ll := NewDoubleLinkedList()
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
		got := ll.AppendRight(c.in).String()
		if got != c.want {
			t.Errorf("dll.AppendRigth(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestDoubleLinkedList_AppendLeft(t *testing.T) {
	ll := NewDoubleLinkedList()
	cases := []struct {
		in   int
		want string
	}{
		{4, " 1: [4]"},
		{7, " 2: [7, 4]"},
		{-1, " 3: [-1, 7, 4]"},
		{0, " 4: [0, -1, 7, 4]"},
	}
	for _, c := range cases {
		got := ll.AppendLeft(c.in).String()
		if got != c.want {
			t.Errorf("dll.AppendLeft(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestDoubleLinkedList_PopRight(t *testing.T) {
	dll := NewDoubleLinkedList()
	for _, i := range []int{3, 2, 1, 0, -1} {
		dll.AppendRight(i)
	}
	cases := []struct {
		want1 int
		want2 string
	}{
		{-1, " 4: [3, 2, 1, 0]"},
		{0, " 3: [3, 2, 1]"},
		{1, " 2: [3, 2]"},
		{2, " 1: [3]"},
		{3, " 0: []"},
	}
	for _, c := range cases {
		got1 := dll.PopRight()
		if got1 != c.want1 {
			t.Errorf("dll.PopRight() == %q, want %q", got1, c.want1)
		}
		got2 := dll.String()
		if got2 != c.want2 {
			t.Errorf("dll.String() == %q, want %q", got2, c.want2)
		}
	}

	defer func() {
		want := "pop from empty list"
		if got := recover(); got != want {
			t.Errorf("panic == '%q', want '%q'", got, want)
		}
	}()
	dll.PopRight()
}

func TestDoubleLinkedList_PopLeft(t *testing.T) {
	dll := NewDoubleLinkedList()
	for _, i := range []int{3, 2, 1, 0, -1} {
		dll.AppendRight(i)
	}
	cases := []struct {
		want1 int
		want2 string
	}{
		{3, " 4: [2, 1, 0, -1]"},
		{2, " 3: [1, 0, -1]"},
		{1, " 2: [0, -1]"},
		{0, " 1: [-1]"},
		{-1, " 0: []"},
	}
	for _, c := range cases {
		got1 := dll.PopLeft()
		if got1 != c.want1 {
			t.Errorf("dll.PopLeft() == %q, want %q", got1, c.want1)
		}
		got2 := dll.String()
		if got2 != c.want2 {
			t.Errorf("dll.String() == %q, want %q", got2, c.want2)
		}
	}

	defer func() {
		want := "pop from empty list"
		if got := recover(); got != want {
			t.Errorf("panic == '%q', want '%q'", got, want)
		}
	}()
	dll.PopLeft()
}
