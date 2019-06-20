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

// Multiply returns ab.
func Multiply(a, b Permutation) Permutation {
	n := len(a)
	if n != len(b) {
		panic("dimension mismatch")
	}

	if !IsPermutation(a) || !IsPermutation(b) {
		panic("not a permutation")
	}

	ab := make(Permutation, 0, n)
	for i := 0; i < n; i++ {
		ab = append(ab, a[b[i]])
	}

	return ab
}

// Multiply returns ab.
func (a Permutation) Multiply(b Permutation) Permutation {
	return Multiply(a, b)
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

// Generate returns the subset <a> of Sn.
func (a Permutation) Generate() set.Set {
	return Generate(a)
}

// Pow returns a^n.
func Pow(a Permutation, n int) Permutation {
	if !IsPermutation(a) {
		panic("not a permutation")
	}

	if n == -1 {
		// Todo: find inverse
	}

	if n < 1 {
		panic("invalid power")
	}

	b := Copy(a)
	for ; 1 < n; n /= 2 {
		b = Multiply(b, b)
	}

	if 0 < n {
		return Multiply(b, a)
	}

	return b
}

/*
func Inverse(a Permutation) Permutation {
	if !IsPermutation(a) {
		panic("not a permutation")
	}
	e := New(len(a))
	b := Copy(a)
	for ; Compare(Multiply(b, a), e) != 0; b = Multiply(b, a) {
	}
	return b
}
func (a Permutation)Inverse()Permutation{
}
*/

// Order returns the number of multiplications of a with itself until it becomes [0, 1, ..., n-1].
func Order(a Permutation) int {
	return len(Generate(a))
}

// Order returns the number of multiplications of a with itself until it becomes [0, 1, ..., n-1].
func (a Permutation) Order() int {
	return Order(a)
}

// Compare two permutations.
func Compare(a, b Permutation) int {
	n := len(a)
	if n != len(b) {
		panic("dimension mismatch")
	}

	if !IsPermutation(a) || !IsPermutation(b) {
		panic("not a permutation")
	}

	for i := 0; i < n; i++ {
		if a[i] < b[i] {
			return -1
		}

		if b[i] < a[i] {
			return 1
		}
	}

	return 0
}

// CompareTo returns the comparison of a permutation to another permutation.
func (a Permutation) CompareTo(b set.Comparable) int {
	return Compare(a, b.(Permutation))
}

// Equal returns true if two permutations are Equal in each indexed value.
func Equal(a, b Permutation) bool {
	return Compare(a, b) == 0
}

// Equal returns true if two permutations are Equal in each indexed value.
func (a Permutation) Equal(b Permutation) bool {
	return Equal(a, b)
}

// IsPermutation returns true if a permutation is an ordering of [0, 1, ..., n-1].
func IsPermutation(a Permutation) bool {
	n := len(a)
	mask := make([]bool, n)
	for _, v := range a {
		if v < 0 || n <= v {
			return false
		}

		if mask[v] {
			return false
		}

		mask[v] = true
	}

	return true
}

// IsPermutation returns true if a permutation is an ordering of [0, 1, ..., n-1].
func (a Permutation) IsPermutation() bool {
	return IsPermutation(a)
}

// Copy a permutation.
func Copy(a Permutation) Permutation {
	b := make(Permutation, len(a))
	copy(b, a)
	return b
}

// Copy a permutation.
func (a Permutation) Copy() Permutation {
	return Copy(a)
}
