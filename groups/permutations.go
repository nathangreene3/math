package groups

import "github.com/nathangreene3/math/set"

// Permutation is an ordering of [0, 1, ..., n-1]. A complete set of all permutations of a given length n are denoted as Sn.
type Permutation []int

// New returns [0, 1, ..., n-1].
func New(n int) Permutation {
	p := make(Permutation, 0, n)
	for i := 0; i < n; i++ {
		p = append(p, i)
	}

	return p
}

// Multiply returns pq.
func Multiply(a, b Permutation) Permutation {
	n := len(a)
	if n != len(b) {
		panic("dimension mismatch")
	}

	if !IsPermutation(a) || !IsPermutation(b) {
		panic("not a permutation")
	}

	c := make(Permutation, 0, n)
	for i := 0; i < n; i++ {
		c = append(c, a[b[i]])
	}

	return c
}

// Generate returns the subset <a> of Sn.
func Generate(a Permutation) set.Set {
	if !IsPermutation(a) {
		panic("not a permutation")
	}

	S := make(set.Set)
	e := New(len(a))
	b := Copy(a)
	set.Insert(S, b)
	for ; !Equal(b, e); set.Insert(S, b) {
		b = Multiply(b, a)
	}

	return S
}

// Order returns the number of multiplications of p with itself until it becomes [0, 1, ..., n-1].
func Order(a Permutation) int {
	return len(Generate(a))
}

// Compare two permutations.
func Compare(p, q Permutation) int {
	n := len(p)
	if n != len(q) {
		panic("dimension mismatch")
	}

	if !IsPermutation(p) || !IsPermutation(q) {
		panic("not a permutation")
	}

	for i := 0; i < n; i++ {
		if p[i] < q[i] {
			return -1
		}

		if q[i] < p[i] {
			return 1
		}
	}

	return 0
}

// Equal returns true if two permutations are Equal in each indexed value.
func Equal(p, q Permutation) bool {
	return Compare(p, q) == 0
}

// IsPermutation returns true if a permutation is an ordering of [0, 1, ..., n-1].
func IsPermutation(p Permutation) bool {
	n := len(p)
	q := make([]bool, n)
	for _, v := range p {
		if v < 0 || n <= v {
			return false
		}

		if q[v] {
			return false
		}

		q[v] = true
	}

	return true
}

// Copy a permutation.
func Copy(p Permutation) Permutation {
	q := make(Permutation, len(p))
	copy(q, p)
	return q
}
