package set

import (
	"testing"
	"fmt"
)

func TestNewSet(t *testing.T) {
	for _, i := range []struct {
		tab []int
		str string
	}{
		{[]int{}, "{}"},
		{[]int{1}, "{1}"},
		{[]int{1, -2}, "{1, -2}"},
		{[]int{1, 1}, "{1}"},
		{[]int{1, 2, 1}, "{1, 2}"},
	} {
		tab, want := i.tab, i.str
		get := NewSet(tab...).String()
		if get != want {
			t.Errorf("NewSet(%v) == %v, want %v", tab, get, want)
		}
	}
}

func TestSet_Len(t *testing.T) {
	for _, i := range []struct {
		s *Set
		l int
	}{
		{NewSet(), 0},
		{NewSet(1), 1},
		{NewSet(1, -2), 2},
		{NewSet(1, 1), 1},
		{NewSet(1, 2, 1), 2},
	} {
		s, want := i.s, i.l
		get := s.Len()
		if get != want {
			t.Errorf("%v.Len() == %v, want %v", s, get, want)
		}
	}
}

func TestSet_Add(t *testing.T) {
	for _, i := range []struct {
		s *Set
		v int
		l int
	}{
		{NewSet(), 4, 1},
		{NewSet(1), 4, 2},
		{NewSet(1, -2), 4, 3},
		{NewSet(1), 1, 1},
		{NewSet(1, 2), 1, 2},
	} {
		s, v, want := i.s, i.v, i.l
		oldS := s.String()
		get := s.Add(v).Len()
		if get != want {
			t.Errorf("%v.Add(%v).Len() == %v, want %v", oldS, v, get, want)
		}
	}
}

func TestSet_sth(t *testing.T) {
	s := NewSet(1, 2, 3)
	fmt.Println(s)
	err := s.Remove(3)
	fmt.Println(s)
	fmt.Println(err)
}
