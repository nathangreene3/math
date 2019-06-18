package groups

// Permutation is an ordering of [0, 1, ..., n-1].
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
func Multiply(p, q Permutation) Permutation {
	n := len(p)
	if n != len(q) {
		panic("dimension mismatch")
	}

	r := make(Permutation, 0, n)
	for i := 0; i < n; i++ {
		r = append(r, p[q[i]])
	}

	return r
}

// Order returns the number of multiplications of p with itself until it becomes [0, 1, ..., n-1].
func Order(p Permutation) int {
	if !IsPermutation(p) {
		panic("not a permutation")
	}

	e := New(len(p))
	q := Copy(p)
	var c int
	for ; !Equal(q, e); c++ {
		q = Multiply(q, p)
	}

	return c
}

// Compare two permutations.
func Compare(p, q Permutation) int {
	n := len(p)
	if n != len(q) {
		panic("dimension mismatch")
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
