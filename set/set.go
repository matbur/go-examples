package set

import (
	"strings"
	"fmt"
	"errors"
)

type Set struct {
	set map[int]struct{}
}

func NewSet(tab ...int) *Set {
	set := make(map[int]struct{})
	for _, i := range tab {
		set[i] = struct{}{}
	}
	s := Set{set}
	return &s
}

func (s *Set) String() string {
	res := []string{"{"}
	for k := range s.set {
		v := fmt.Sprintf(", %d", k)
		if len(res) == 1 {
			v = fmt.Sprintf("%d", k)
		}
		res = append(res, v)
	}
	res = append(res, "}")
	return strings.Join(res, "")
}

func (s *Set) Len() int {
	return len(s.set)
}

func (s *Set) Add(value int) *Set {
	s.set[value] = struct{}{}
	return s
}

func (s *Set) Remove(value int) error {
	if ok := s.IsIn(value); !ok {
		return errors.New("there is no such value Set")
	}
	delete(s.set, value)
	return nil
}

func (s *Set) IsIn(value int) bool {
	_, ok := s.set[value]
	return ok
}

func (s *Set) foo() {}
