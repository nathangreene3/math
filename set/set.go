package set

import "sort"

// A Set is a collection of keyed values.
type Set map[int]interface{}

// New returns a set mapping the range of keys [0,n) to each value.
func New(S []interface{}) Set {
	B := make(Set)
	for k := range S {
		B[k] = S[k]
	}

	return B
}

// Copy a set.
func Copy(A Set) Set {
	B := make(Set)
	for k := range A {
		B[k] = A[k]
	}

	return B
}

// ToSlice exports a set to a slice.
func ToSlice(A Set) []interface{} {
	S := make([]interface{}, 0, len(A))
	for k := range A {
		S = append(S, A[k])
	}

	return S
}

// Equal returns true if A and B each contain the same values. Note that the keys may differ.
func Equal(A, B Set) bool {
	if len(A) != len(B) {
		return false
	}

	for k := range A {
		if _, ok := Contains(B, A[k]); !ok {
			return false
		}
	}

	return true
}

// Sort returns a sorted set with new keys on the range [0,n).
func Sort(A Set, less func(i, j int) bool) Set {
	S := ToSlice(A)
	sort.Slice(S, less)
	return New(S)
}

// Insert a value into a set. Returns the key. Duplicate values will not be inserted, but the current key of the existing value will be returned.
func Insert(A Set, v interface{}) (Set, int) {
	k, ok := Contains(A, v)
	if !ok {
		A[k] = v
	}

	return A, k
}

// Remove and return a value from a set.
func Remove(A Set, k int) interface{} {
	if v, ok := A[k]; ok {
		delete(A, k)
		return v
	}

	return nil
}

// Contains returns the key of a value and true if it is found. If not, then the next available key and false is returned.
func Contains(A Set, v interface{}) (int, bool) {
	n := len(A)
	if n == 0 {
		return 0, false
	}

	keys := make([]int, 0, n)
	for k := range A {
		if A[k] == v {
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

// Intersection of A and B returns a set containing values found in both A and B (AnB).
func Intersection(A, B Set) Set {
	if len(A) < len(B) {
		C := make(Set)
		var k int
		for _, v := range A {
			if _, ok := B[k]; ok {
				C[k] = v
				k++
			}
		}

		return C
	}

	return Intersection(B, A)
}

// IntersectionMany several sets.
func IntersectionMany(As ...Set) Set {
	A := make(Set)
	for i := range As {
		A = Intersection(A, As[i])
	}

	return A
}

// Union of A and B returns a set containing values in either A or B (AuB). It is synonymous with joining or combining sets.
func Union(A, B Set) Set {
	C := make(Set)
	var k int
	for _, v := range A {
		C[k] = v
		k++
	}

	for _, v := range B {
		if _, ok := Contains(A, v); !ok {
			C[k] = v
			k++
		}
	}

	return C
}

// UnionMany several sets.
func UnionMany(As ...Set) Set {
	A := make(Set)
	for i := range As {
		A = Union(A, As[i])
	}

	return A
}

// LeftDisjoint returs a set containing values in A, but not B (A-B).
func LeftDisjoint(A, B Set) Set {
	C := make(Set)
	var k int
	for _, v := range A {
		if _, ok := Contains(B, v); !ok {
			C[k] = v
			k++
		}
	}

	return C
}

// RightDisjoint returns a set containing values in B, but not A (B-A).
func RightDisjoint(A, B Set) Set {
	return LeftDisjoint(B, A)
}

// Disjoint returns a set containing values in A and B, but not in both A and B ((AuB)-(AnB)).
func Disjoint(A, B Set) Set {
	C := make(Set)
	var k int
	for _, v := range A {
		if _, ok := Contains(B, v); !ok {
			C[k] = v
			k++
		}
	}

	for _, v := range B {
		if _, ok := Contains(A, v); !ok {
			C[k] = v
			k++
		}
	}

	return C
}

// Cardinality returns the number of values in a set.
func Cardinality(A Set) int {
	return len(A)
}
