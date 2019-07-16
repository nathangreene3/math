package set

import (
	"sort"
)

// Comparable defines how similar types are compared.
type Comparable interface {
	CompareTo(v Comparable) int
}

// A Set is a collection of keyed values.
type Set map[int]Comparable

// New returns a set mapping the range of keys [0,n) to each value.
func New(S ...Comparable) Set {
	B := make(Set)
	for k := range S {
		B[k] = S[k]
	}

	return B
}

// Insert a value into a set. Returns the key. Duplicate values will not be inserted, but the current key of the existing value will be returned.
func (S Set) Insert(v Comparable) int {
	k, ok := S.Contains(v)
	if !ok {
		S[k] = v
	}

	return k
}

// Remove and return a value from a set. Returns nil if value not found.
func (S Set) Remove(k int) Comparable {
	if v, ok := S[k]; ok {
		delete(S, k)
		return v
	}

	return nil
}

// Contains returns the key of a value and true if it is found. If not, then the next available key and false is returned.
func (S Set) Contains(v Comparable) (int, bool) {
	n := len(S)
	if n == 0 {
		return 0, false
	}

	keys := make([]int, 0, n)
	for k := range S {
		if S[k].CompareTo(v) == 0 {
			return k, true
		}

		keys = append(keys, k)
	}

	sort.Ints(keys)
	for i := range keys {
		if i < keys[i] {
			return i, false
		}
	}

	return n, false
}

// Intersect of A and B returns a set containing values found in both A and B (AnB).
func (S Set) Intersect(T Set) Set {
	if len(S) < len(T) {
		var (
			U = make(Set)
			k int
		)
		for _, v := range S {
			if _, ok := T[k]; ok {
				U[k] = v
				k++
			}
		}

		return U
	}

	return T.Intersect(S)
}

// Intersection several sets.
func Intersection(S ...Set) Set {
	T := make(Set)
	for i := range S {
		T = T.Intersect(S[i])
	}

	return T
}

// Union of A and B returns a set containing values in either A or B (AuB). It is synonymous with joining or combining sets.
func (S Set) Union(T Set) Set {
	var (
		U = make(Set)
		k int
	)
	for _, v := range S {
		U[k] = v
		k++
	}

	for _, v := range T {
		if _, ok := S.Contains(v); !ok {
			U[k] = v
			k++
		}
	}

	return U
}

// Union several sets.
func Union(S ...Set) Set {
	T := make(Set)
	for i := range S {
		T = T.Union(S[i])
	}

	return T
}

// LeftDisjoint returs a set containing values in A, but not B (A-B).
func (S Set) LeftDisjoint(T Set) Set {
	var (
		U = make(Set)
		k int
	)
	for _, v := range S {
		if _, ok := T.Contains(v); !ok {
			U[k] = v
			k++
		}
	}

	return U
}

// RightDisjoint returns a set containing values in B, but not A (B-A).
func (S Set) RightDisjoint(T Set) Set {
	return T.LeftDisjoint(S)
}

// Disjoint returns a set containing values in A and B, but not in both A and B ((AuB)-(AnB)).
func (S Set) Disjoint(T Set) Set {
	var (
		U = make(Set)
		k int
	)
	for _, v := range S {
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

// Cardinality returns the number of values in a set.
func (S Set) Cardinality() int {
	return len(S)
}

// Equal returns true if A and B each contain the same values. Note that the keys may differ.
func (S Set) Equal(T Set) bool {
	if len(S) != len(T) {
		return false
	}

	for k := range S {
		if _, ok := T.Contains(S[k]); !ok {
			return false
		}
	}

	return true
}

// Copy a set.
func (S Set) Copy() Set {
	T := make(Set)
	for k := range S {
		T[k] = S[k]
	}

	return T
}

// ToSlice exports a set to a slice.
func (S Set) ToSlice() []Comparable {
	T := make([]Comparable, 0, len(S))
	for k := range S {
		T = append(T, S[k])
	}

	return T
}

// Sort returns a sorted set with new keys on the range [0,n).
func (S Set) Sort() Set {
	s := S.ToSlice()
	sort.Slice(s, func(i, j int) bool { return s[i].CompareTo(s[j]) < 0 })
	return New(s...)
}
