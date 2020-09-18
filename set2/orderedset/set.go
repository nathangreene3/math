package orderedset

import (
	"sort"

	"github.com/nathangreene3/math"
)

// Set ...
type Set []Comparable

// Comparable ...
type Comparable interface {
	Compare(Comparable) int
}

// New ...
func New(values ...Comparable) Set {
	return make(Set, 0, math.NextPowOfTwo(len(values))).Insert(values...)
}

// Index ...
func (s Set) Index(value Comparable) (int, bool) {
	i := sort.Search(len(s), func(i int) bool { return value.Compare(s[i]) <= 0 })
	return i, value.Compare(s[i]) == 0
}

// Insert ...
func (s Set) Insert(values ...Comparable) Set {
	for i := 0; i < len(values); i++ {
		if _, ok := s.Index(values[i]); !ok {
			s = append(s, values[i])
			sort.SliceStable(s, s.Less)
		}
	}

	return s
}

// Len ...
func (s Set) Len() int {
	return len(s)
}

// Less ...
func (s Set) Less(i, j int) bool {
	return s[i].Compare(s[j]) < 0
}

// Remove ...
func (s Set) Remove(values ...Comparable) Set {
	for i := 0; i < len(values); i++ {
		if j, ok := s.Index(values[i]); ok {
			s = append(s[:j], s[j+1:]...)
		}
	}

	return s
}

// Swap ...
func (s Set) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Union ...
func (s Set) Union(t Set) Set {
	for i := 0; i < len(t); i++ {
		if _, ok := s.Index(t[i]); !ok {
			s = append(s, t[i])
			sort.SliceStable(s, s.Less)
		}
	}

	return s
}
