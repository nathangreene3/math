package set2

import "sort"

// Set is a map of comparable keys.
type Set map[Comparable]struct{}

// Comparable defines the comparison of two values.
type Comparable interface {
	Compare(Comparable) int
}

// Filterer defines the retaining of a value.
type Filterer func(v Comparable) bool

// Folder defines the reduction of a generic value with a comparable value.
type Folder func(x interface{}, v Comparable) interface{}

// Generator defines the ith value on the range [0, n).
type Generator func(i int) Comparable

// Mapper defines the mapping of a value.
type Mapper func(v Comparable) Comparable

// Reducer defines the reduction of two values into one.
type Reducer func(u, v Comparable) Comparable

// New set.
func New(vs ...Comparable) Set {
	return make(Set).Insert(vs...)
}

// Empty set.
func Empty() Set {
	return make(Set)
}

// Generate a set of n values.
func Generate(n int, f Generator) Set {
	s := make(Set)
	for i := 0; i < n; i++ {
		s[f(i)] = struct{}{}
	}

	return s
}

// Copy a set.
func (s Set) Copy() Set {
	t := make(Set)
	for v := range s {
		t[v] = struct{}{}
	}

	return t
}

// CountGreater returns the number of values in a set greater than a given value.
func (s Set) CountGreater(v Comparable) int {
	var c int
	for u := range s {
		if v.Compare(u) < 0 {
			c++
		}
	}

	return c
}

// CountGreaterEqual returns the number of values in a set greater than or equal to a given value.
func (s Set) CountGreaterEqual(v Comparable) int {
	var c int
	for u := range s {
		if v.Compare(u) <= 0 {
			c++
		}
	}

	return c
}

// CountLess returns the number of values in a set less than a given value.
func (s Set) CountLess(v Comparable) int {
	var c int
	for u := range s {
		if u.Compare(v) < 0 {
			c++
		}
	}

	return c
}

// CountLessEqual returns the number of values in a set less than or equal to a given value.
func (s Set) CountLessEqual(v Comparable) int {
	var c int
	for u := range s {
		if u.Compare(v) <= 0 {
			c++
		}
	}

	return c
}

// Complement of a set with respect to another set. Modifies and returns s.
func (s Set) Complement(t Set) Set {
	for v := range s {
		if _, ok := t[v]; ok {
			delete(s, v)
		}
	}

	return s
}

// Contains indicates if a value is in a set.
func (s Set) Contains(v Comparable) bool {
	_, ok := s[v]
	return ok
}

// ContainsSubset determines if a set contains all of another set.
func (s Set) ContainsSubset(t Set) bool {
	if len(s) < len(t) {
		return false
	}

	for v := range t {
		if _, ok := s[v]; !ok {
			return false
		}
	}

	return true
}

// Equal compares two sets.
func (s Set) Equal(t Set) bool {
	if len(s) != len(t) {
		return false
	}

	for v := range s {
		if _, ok := t[v]; !ok {
			return false
		}
	}

	return true
}

// Filter returns a new set with filtered values.
func (s Set) Filter(f Filterer) Set {
	t := make(Set)
	for v := range s {
		if f(v) {
			t[v] = struct{}{}
		}
	}

	return t
}

// Fold a set of values into a generic value given a generic seed.
func (s Set) Fold(x interface{}, f Folder) interface{} {
	for v := range s {
		x = f(x, v)
	}

	return x
}

// Insert values into a set. Modifies and returns s.
func (s Set) Insert(vs ...Comparable) Set {
	for i := 0; i < len(vs); i++ {
		s[vs[i]] = struct{}{}
	}

	return s
}

// Intersect returns the intersection of two sets. Modifies and returns s.
func (s Set) Intersect(t Set) Set {
	for v := range s {
		if _, ok := t[v]; !ok {
			delete(s, v)
		}
	}

	return s
}

// Intersects determines if two sets intersect.
func (s Set) Intersects(t Set) bool {
	if len(s) < len(t) {
		for v := range s {
			if _, ok := t[v]; ok {
				return true
			}
		}
	} else {
		for v := range t {
			if _, ok := s[v]; ok {
				return true
			}
		}
	}

	return false
}

// Intersection of several sets.
func Intersection(ss ...Set) Set {
	switch len(ss) {
	case 0:
		return make(Set)
	case 1:
		return ss[0].Copy()
	default:
		t := make(Set)
		var i int // Index of smallest set
		for j := 1; j < len(ss); j++ {
			if len(ss[j]) < len(ss[i]) {
				i = j
			}
		}

		for v := range ss[i] {
			contains := true
			for j := 0; j < len(ss) && contains; j++ {
				if j != i {
					if _, ok := ss[j][v]; !ok {
						contains = false
					}
				}
			}

			if contains {
				t[v] = struct{}{}
			}
		}

		return t
	}
}

// IsSubsetOf determines if a set is a subset of another set.
func (s Set) IsSubsetOf(t Set) bool {
	if len(t) < len(s) {
		return false
	}

	for v := range s {
		if _, ok := t[v]; !ok {
			return false
		}
	}

	return true
}

// Map a set to a new set.
func (s Set) Map(f Mapper) Set {
	t := make(Set)
	for v := range s {
		t[f(v)] = struct{}{}
	}

	return t
}

// Max value in a set.
func (s Set) Max() Comparable {
	var (
		m     Comparable
		first = true
	)

	for v := range s {
		if first {
			m = v
			first = false
		} else {
			if m.Compare(v) < 0 {
				m = v
			}
		}
	}

	return m
}

// Min value in a set.
func (s Set) Min() Comparable {
	var (
		m     Comparable
		first = true
	)

	for v := range s {
		if first {
			m = v
			first = false
		} else {
			if v.Compare(m) < 0 {
				m = v
			}
		}
	}

	return m
}

// Reduce a set to a single value.
func (s Set) Reduce(f Reducer) Comparable {
	var (
		u     Comparable
		first = true
	)

	for v := range s {
		if first {
			u = v
			first = false
		} else {
			u = f(u, v)
		}
	}

	return u
}

// Remove several values from a set. Modifies and returns s.
func (s Set) Remove(vs ...Comparable) Set {
	for i := 0; i < len(vs); i++ {
		delete(s, vs[i])
	}

	return s
}

// Size of a set.
func (s Set) Size() int {
	return len(s)
}

// Slice returns a sorted slice of values in a set.
func (s Set) Slice() []Comparable {
	slc := make([]Comparable, 0, len(s))
	for v := range s {
		slc = append(slc, v)
	}

	sort.Slice(slc, func(i, j int) bool { return slc[i].Compare(slc[j]) < 0 })
	return slc
}

// SubsetGreater returns a set of values greater than a given value.
func (s Set) SubsetGreater(v Comparable) Set {
	t := make(Set)
	for u := range s {
		if v.Compare(u) < 0 {
			t[u] = struct{}{}
		}
	}

	return t
}

// SubsetGreaterEqual returns a set of values greater than or equal to a given value.
func (s Set) SubsetGreaterEqual(v Comparable) Set {
	t := make(Set)
	for u := range s {
		if v.Compare(u) <= 0 {
			t[u] = struct{}{}
		}
	}

	return t
}

// SubsetLess returns a set of values less than a given value.
func (s Set) SubsetLess(v Comparable) Set {
	t := make(Set)
	for u := range s {
		if u.Compare(v) < 0 {
			t[u] = struct{}{}
		}
	}

	return t
}

// SubsetLessEqual returns a set of values less than or equal to a given value.
func (s Set) SubsetLessEqual(v Comparable) Set {
	t := make(Set)
	for u := range s {
		if u.Compare(v) <= 0 {
			t[u] = struct{}{}
		}
	}

	return t
}

// SubsetRange returns a set of values between two values.
func (s Set) SubsetRange(min, max Comparable) Set {
	if max.Compare(min) < 0 {
		min, max = max, min
	}

	t := make(Set)
	for v := range s {
		if min.Compare(v) < 0 && v.Compare(max) < 0 {
			t[v] = struct{}{}
		}
	}

	return t
}

// SubsetRangeEqual returns a set of values between two values. The given values are potentially included in the returned set.
func (s Set) SubsetRangeEqual(min, max Comparable) Set {
	if max.Compare(min) < 0 {
		min, max = max, min
	}

	t := make(Set)
	for v := range s {
		if min.Compare(v) <= 0 && v.Compare(max) <= 0 {
			t[v] = struct{}{}
		}
	}

	return t
}

// SubsetOmitRange returns a set of values outside a given range.
func (s Set) SubsetOmitRange(min, max Comparable) Set {
	if max.Compare(min) < 0 {
		min, max = max, min
	}

	t := make(Set)
	for u := range s {
		if u.Compare(min) < 0 || max.Compare(u) < 0 {
			t[u] = struct{}{}
		}
	}

	return t
}

// SubsetOmitRangeEqual returns a set of values outside a given range. The given values are potentially included in the returned set.
func (s Set) SubsetOmitRangeEqual(min, max Comparable) Set {
	if max.Compare(min) < 0 {
		min, max = max, min
	}

	t := make(Set)
	for u := range s {
		if u.Compare(min) <= 0 || max.Compare(u) <= 0 {
			t[u] = struct{}{}
		}
	}

	return t
}

// Union of two sets. Modifies and returns s.
func (s Set) Union(t Set) Set {
	for v := range t {
		s[v] = struct{}{}
	}

	return s
}

// Union of several sets.
func Union(ss ...Set) Set {
	t := make(Set)
	for i := 0; i < len(ss); i++ {
		for v := range ss[i] {
			t[v] = struct{}{}
		}
	}

	return t
}
