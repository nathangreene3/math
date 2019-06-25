package groups

import "github.com/nathangreene3/math/set"

// Permutation is an ordering of [0, 1, ..., n-1]. A complete set of all permutations of a given length n are denoted as Sn.
type Permutation []int

// Identity returns [0, 1, ..., n-1].
func Identity(n int) Permutation {
	a := make(Permutation, 0, n)
	for i := 0; i < n; i++ {
		a = append(a, i)
	}

	return a
}

// Multiply returns ab.
func (a Permutation) Multiply(b Permutation) Permutation {
	n := len(a)
	if n != len(b) {
		panic("dimension mismatch")
	}

	if !a.IsPermutation() || !b.IsPermutation() {
		panic("not a permutation")
	}

	ab := make(Permutation, 0, n)
	for i := 0; i < n; i++ {
		ab = append(ab, a[b[i]])
	}

	return ab
}

// Multiply several permutations from left to right.
func Multiply(a ...Permutation) Permutation {
	n := len(a)
	if n == 0 {
		return nil
	}

	b := a[0].Copy()
	for i := 1; i < n; i++ {
		b = b.Multiply(a[i])
	}

	return b
}

// generate TODO: generate entire (sub) group.
func generate(a ...Permutation) set.Set {
	S := set.New()
	for i := range a {
		S = S.Union(a[i].Generate())
	}

	return S
}

// Generate returns the subset <a> of Sn.
func (a Permutation) Generate() set.Set {
	if !a.IsPermutation() {
		panic("not a permutation")
	}

	S := make(set.Set)
	e := Identity(len(a))
	b := a.Copy()
	S.Insert(b)
	for ; b.CompareTo(e) != 0; S.Insert(b) {
		b = b.Multiply(a)
	}

	return S
}

// Pow returns a^n.
func (a Permutation) Pow(n int) Permutation {
	if !a.IsPermutation() {
		panic("not a permutation")
	}

	if n == -1 {
		// Todo: find inverse
	}

	// Yacas' method
	b := Identity(len(a))
	c := a.Copy()
	for ; 0 < n; n /= 2 {
		if n%2 == 0 {
			b = b.Multiply(c)
		}

		c = c.Multiply(c)
	}

	return b
}

// Order returns the number of multiplications of a with itself until it becomes [0, 1, ..., n-1].
func (a Permutation) Order() int {
	b := a.Multiply(a)
	n := 1
	for ; b.CompareTo(a) != 0; n++ {
		b = b.Multiply(a)
	}

	return n
}

// Cayley returns a Cayley table, a matrix consisting of each pair of permutations multiplied. The (i,j)th entry is the as[i]*as[j] result.
func Cayley(a ...Permutation) [][]Permutation {
	n := len(a)
	table := make([][]Permutation, 0, n)
	for i := range a {
		table = append(table, make([]Permutation, 0, n))
		for j := range a {
			table[i] = append(table[i], a[i].Multiply(a[j]))
		}
	}

	return table
}

// CompareTo returns the comparison of a permutation to another permutation.
func (a Permutation) CompareTo(b set.Comparable) int {
	n := len(a)
	if n != len(b.(Permutation)) {
		panic("dimension mismatch")
	}

	if !a.IsPermutation() || !b.(Permutation).IsPermutation() {
		panic("not a permutation")
	}

	for i := 0; i < n; i++ {
		switch {
		case a[i] < b.(Permutation)[i]:
			return -1
		case b.(Permutation)[i] < a[i]:
			return 1
		}
	}

	return 0
}

// Equal returns true if two permutations are Equal in each indexed value.
func (a Permutation) Equal(b Permutation) bool {
	return a.CompareTo(b) == 0
}

// IsPermutation returns true if a permutation is an ordering of [0, 1, ..., n-1].
func (a Permutation) IsPermutation() bool {
	n := len(a)
	mask := make([]bool, n)
	for _, v := range a {
		switch {
		case v < 0, n <= v:
			return false
		case mask[v]:
			return false
		default:
			mask[v] = true
		}
	}

	return true
}

// Copy a permutation.
func (a Permutation) Copy() Permutation {
	b := make(Permutation, len(a))
	copy(b, a)
	return b
}
