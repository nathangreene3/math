package flts

import "sort"

// Set ...
type Set struct {
	values map[float64]struct{}
}

// New set.
func New(vs ...float64) *Set {
	s := Set{values: make(map[float64]struct{})}
	return s.Insert(vs...)
}

// Copy a set.
func (s *Set) Copy() *Set {
	t := Set{values: make(map[float64]struct{})}
	for v := range s.values {
		t.values[v] = struct{}{}
	}

	return &t
}

// Complement of a set with respect to another set. That is, A-B.
func (s *Set) Complement(t *Set) *Set {
	for v := range s.values {
		if _, ok := t.values[v]; ok {
			delete(s.values, v)
		}
	}

	return s
}

// Contains indicates if a value is in a set.
func (s *Set) Contains(v float64) bool {
	_, ok := s.values[v]
	return ok
}

// Equal compares two sets.
func (s *Set) Equal(t *Set) bool {
	for v := range s.values {
		if _, ok := t.values[v]; !ok {
			return false
		}
	}

	return true
}

// Insert values into a set.
func (s *Set) Insert(vs ...float64) *Set {
	for i := 0; i < len(vs); i++ {
		s.values[vs[i]] = struct{}{}
	}

	return s
}

// Intersect returns the intersection of two sets. That is, AnB.
func (s *Set) Intersect(t *Set) *Set {
	for v := range s.values {
		if _, ok := t.values[v]; !ok {
			delete(s.values, v)
		}
	}

	return s
}

// Intersection ...
func Intersection(ss ...*Set) *Set {
	switch len(ss) {
	case 0:
		return &Set{values: make(map[float64]struct{})}
	case 1:
		return ss[0].Copy()
	default:
		t := Set{values: make(map[float64]struct{})}
		var i int // Index of smallest set
		for j := 1; j < len(ss); j++ {
			if len(ss[j].values) < len(ss[i].values) {
				i = j
			}
		}

		for v := range ss[i].values {
			contains := true
			for j := 0; j < len(ss) && contains; j++ {
				if j != i {
					if _, ok := ss[j].values[v]; !ok {
						contains = false
					}
				}
			}

			if contains {
				t.values[v] = struct{}{}
			}
		}

		return &t
	}
}

// Size of a set.
func (s *Set) Size() int {
	return len(s.values)
}

// Slice returns a sorted slice.
func (s *Set) Slice() []float64 {
	slc := make([]float64, 0, len(s.values))
	for v := range s.values {
		slc = append(slc, v)
	}

	sort.Slice(len(slc), func(i, j int) bool { return slc[i] < slc[j] })
	return slc
}

// Union of two sets. That is, AuB.
func (s *Set) Union(t *Set) *Set {
	for v := range t.values {
		s.values[v] = struct{}{}
	}

	return s
}

// Union of several sets.
func Union(ss ...*Set) *Set {
	t := Set{values: make(map[float64]struct{})}
	for i := 0; i < len(ss); i++ {
		for v := range ss[i].values {
			t.values[v] = struct{}{}
		}
	}

	return &t
}

// Mapper ...
type Mapper func(v float64) float64

// Map ...
func Map(s *Set, f Mapper) *Set {
	t := Set{values: make(map[float64]struct{})}
	for v := range s.values {
		t.values[f(v)] = struct{}{}
	}

	return &t
}

// Filterer ...
type Filterer func(v float64) bool

// Filter ...
func Filter(s *Set, f Filterer) *Set {
	t := Set{values: make(map[float64]struct{})}
	for v := range s.values {
		if f(v) {
			t.values[v] = struct{}{}
		}
	}

	return &t
}

// Reducer ...
type Reducer func(u, v float64) float64

// Reduce ...
func Reduce(s *Set, f Reducer) float64 {
	var u float64
	for v := range s.values {
		u = f(u, v)
	}

	return u
}
