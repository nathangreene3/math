package set

import (
	"sort"
)

// Comparable defines how similar types are compared.
type Comparable interface {
	Compare(v Comparable) int
}

// A Set is a collection of keyed values.
type Set map[int]Comparable

// New returns a set mapping the range of keys [0,n) to each value.
func New(S ...Comparable) Set {
	B := make(Set)
	for k, s := range S {
		B[k] = s
	}

	return B
}

// Contains returns the key of a value and true if it is found. If not, then the
// next available key and false is returned.
func (S *Set) Contains(value Comparable) (int, bool) {
	n := len(*S)
	if n == 0 {
		return 0, false
	}

	keys := make([]int, 0, n)
	for k, v := range *S {
		if v.Compare(value) == 0 {
			return k, true
		}

		keys = append(keys, k)
	}

	sort.Ints(keys)
	for i, k := range keys {
		if i < k {
			return i, false
		}
	}

	return n, false
}

// Cardinality returns the number of values in a set.
func (S *Set) Cardinality() int {
	return len(*S)
}

// Copy a set.
func (S *Set) Copy() Set {
	T := make(Set)
	for k, v := range *S {
		T[k] = v
	}

	return T
}

// Disjoint returns a set containing values in A and B, but not in both A and B
// (AuB-AnB).
func (S *Set) Disjoint(T Set) Set {
	var (
		U = make(Set)
		k int
	)

	for _, v := range *S {
		if _, ok := T.Contains(v); !ok {
			U[k] = v
			k++
		}
	}

	for _, v := range T {
		if _, ok := S.Contains(v); !ok {
			U[k] = v
			k++
		}
	}

	return U
}

// Equal returns true if A and B each contain the same values. Note that the
// keys may differ.
func (S *Set) Equal(T Set) bool {
	if len(*S) != len(T) {
		return false
	}

	for _, v := range *S {
		if _, ok := T.Contains(v); !ok {
			return false
		}
	}

	return true
}

// Insert a value into a set. Returns the key. Duplicate values will not be
// inserted, but the current key of the existing value will be returned.
func (S *Set) Insert(v Comparable) int {
	k, ok := S.Contains(v)
	if !ok {
		(*S)[k] = v
	}

	return k
}

// Intersect of A and B returns a set containing values found in both A and B
// (AnB).
func (S *Set) Intersect(T Set) Set {
	if len(*S) < len(T) {
		var (
			U = make(Set)
			k int
		)

		for _, v := range *S {
			if _, ok := T[k]; ok {
				U[k] = v
				k++
			}
		}

		return U
	}

	return T.Intersect(*S)
}

// Intersection several sets.
func Intersection(S ...Set) Set {
	T := make(Set)
	for _, s := range S {
		T = T.Intersect(s)
	}

	return T
}

// LeftDisjoint returs a set containing values in A, but not B (A-B).
func (S *Set) LeftDisjoint(T Set) Set {
	var (
		U = make(Set)
		k int
	)

	for _, v := range *S {
		if _, ok := T.Contains(v); !ok {
			U[k] = v
			k++
		}
	}

	return U
}

// Remove and return a value from a set. Returns nil if value not found.
func (S *Set) Remove(k int) Comparable {
	if v, ok := (*S)[k]; ok {
		delete(*S, k)
		return v
	}

	return nil
}

// RightDisjoint returns a set containing values in B, but not A (B-A).
func (S *Set) RightDisjoint(T Set) Set {
	return T.LeftDisjoint(*S)
}

// Sort returns a sorted set with new keys on the range [0,n).
func (S *Set) Sort() Set {
	s := (*S).ToSlice()
	sort.Slice(s, func(i, j int) bool { return s[i].Compare(s[j]) < 0 })
	return New(s...)
}

// ToSlice exports a set to a slice.
func (S *Set) ToSlice() []Comparable {
	T := make([]Comparable, 0, len(*S))
	for _, v := range *S {
		T = append(T, v)
	}

	return T
}

// Union of A and B returns a set containing values in either A or B (AuB). It
// is synonymous with joining or combining sets.
func (S *Set) Union(T Set) Set {
	var (
		U = make(Set)
		k int
	)

	for _, v := range *S {
		U[k] = v
		k++
	}

	for _, v := range T {
		if _, ok := (*S).Contains(v); !ok {
			U[k] = v
			k++
		}
	}

	return U
}

// Union several sets.
func Union(S ...Set) Set {
	T := make(Set)
	for _, s := range S {
		T = T.Union(s)
	}

	return T
}
