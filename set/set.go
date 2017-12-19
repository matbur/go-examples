// Package set provides implementation of set
package set

import (
	"strings"
	"fmt"
	"errors"
)

type Set struct {
	set map[int]struct{}
}

// Build an unordered collection of unique integers.
func NewSet(values ...int) *Set {
	set := make(map[int]struct{})
	for _, i := range values {
		set[i] = struct{}{}
	}
	s := Set{set}
	return &s
}

// Convert to string.
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

// Get number of items in set.
func (s *Set) Len() int {
	return len(s.set)
}

// Get rid of all values from set.
func (s *Set) Clear() *Set { return nil }

// Make copy of set.
func (s *Set) Copy() *Set { return nil }

// Add new value to set in place.
func (s *Set) Add(value int) *Set {
	s.set[value] = struct{}{}
	return s
}

// Check if set contains value.
func (s *Set) IsIn(value int) bool {
	_, ok := s.set[value]
	return ok
}

// Remove a value from set.
func (s *Set) Remove(value int) error {
	if ok := s.IsIn(value); !ok {
		return errors.New("there is no such value Set")
	}
	delete(s.set, value)
	return nil
}

func (s *Set) Difference(other *Set) *Set          { return nil }
func (s *Set) Intersection(other *Set) *Set        { return nil }
func (s *Set) Pop() int                            { return -1 }
func (s *Set) SymmetricDifference(other *Set) *Set { return nil }
func (s *Set) Union(other *Set) *Set               { return nil }

func (s *Set) IsDisjoint(other *Set) bool { return false }
func (s *Set) IsSubset(other *Set) bool   { return false }
func (s *Set) IsSuperset(other *Set) bool { return false }

func (s *Set) foo(other *Set) *Set { return nil }
func (s *Set) bar()                {}
