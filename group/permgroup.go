package group

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/nathangreene3/math/bitmask"
	"github.com/nathangreene3/math/set2"
)

// Permutation is an ordering of [0, 1, ..., n-1]. A complete set of all permutations of a given length n are denoted as Sn.
type Permutation []int

// New returns [0, 1, ..., n-1].
func New(n int) Permutation {
	a := make(Permutation, 0, n)
	for i := 0; i < n; i++ {
		a = append(a, i)
	}

	return a
}

// Identity returns [0, 1, ..., n-1].
func Identity(n int) Permutation {
	return New(n)
}

// Parse ...
func Parse(s string) Permutation {
	if 0 < len(s) && s[0] == '(' && s[len(s)-1] == ')' {
		var (
			ss = strings.Split(s, ",")
			a  = make(Permutation, 0, len(ss))
		)

		for i := 0; i < len(ss); i++ {
			ssi, err := strconv.Atoi(ss[i])
			if err != nil {
				panic(err.Error())
			}

			a = append(a, ssi)
		}

		if a.isPermutation() {
			return a
		}
	}

	panic("not a permutation")
}

// Cayley returns a Cayley table, a matrix consisting of each pair of permutations multiplied. The (i,j)th entry is the as[i]*as[j] result.
func Cayley(ps ...Permutation) [][]Permutation {
	var (
		n     = len(ps)
		table = make([][]Permutation, 0, n)
	)

	for i := 0; i < n; i++ {
		permutations := make([]Permutation, 0, n)
		for j := 0; j < n; j++ {
			permutations = append(permutations, ps[i].Multiply(ps[j]))
		}

		table = append(table, permutations)
	}

	return table
}

// Compare returns the comparison of a permutation to another permutation.
func (a Permutation) Compare(b set2.Comparable) int {
	c := b.(Permutation)
	switch {
	case len(a) != len(c):
		panic("dimension mismatch")
	case !a.isPermutation(), !c.isPermutation():
		panic("not a permutation")
	}

	for i := 0; i < len(a); i++ {
		switch {
		case a[i] < c[i]:
			return -1
		case c[i] < a[i]:
			return 1
		}
	}

	return 0
}

// Copy a permutation.
func (a Permutation) Copy() Permutation {
	b := make(Permutation, len(a))
	copy(b, a)
	return b
}

// Equal returns true if two permutations are Equal in each indexed value.
func (a Permutation) Equal(b Permutation) bool {
	return a.Compare(b) == 0
}

// generate TODO: generate entire (sub) group.
func generate(a ...Permutation) set2.Set {
	S := set2.New()
	for i := 0; i < len(a); i++ {
		S = S.Union(a[i].Generate())
	}

	return S
}

// Generate returns the subset <a> of Sn.
func (a Permutation) Generate() set2.Set {
	e := Identity(len(a))
	b := a.Copy()
	S := set2.New(b)
	for ; b.Compare(e) != 0; S.Insert(b) {
		b = b.Multiply(a)
	}

	return S
}

// isPermutation returns true if a permutation is an ordering of [0, 1, ..., n-1].
func (a Permutation) isPermutation() bool {
	bm := bitmask.New(big.NewInt(0))
	for i := 0; i < len(a); i++ {
		v := a[i]
		switch {
		case v < 0, len(a) <= v:
			return false
		case bm.IsSet(big.NewInt(int64(1 << v))):
			return false
		default:
			bm.Set(big.NewInt(int64(1 << v)))
		}
	}

	return true
}

// Multiply several permutations from left to right.
func Multiply(a ...Permutation) Permutation {
	if len(a) == 0 {
		return make(Permutation, 0)
	}

	b := a[0].Copy()
	for i := 1; i < len(a); i++ {
		b = b.Multiply(a[i])
	}

	return b
}

// Multiply returns ab.
func (a Permutation) Multiply(b Permutation) Permutation {
	n := len(a)
	if n != len(b) {
		panic("dimension mismatch")
	}

	if !a.isPermutation() || !b.isPermutation() {
		panic("not a permutation")
	}

	ab := make(Permutation, 0, n)
	for i := 0; i < n; i++ {
		ab = append(ab, a[b[i]])
	}

	return ab
}

// Order returns the number of multiplications of a with itself until it becomes [0, 1, ..., n-1].
func (a Permutation) Order() int {
	b := a.Multiply(a)
	n := 1
	for ; b.Compare(a) != 0; n++ {
		b = b.Multiply(a)
	}

	return n
}

// Pow returns a^p.
func (a Permutation) Pow(p int) Permutation {
	if p == -1 {
		// Todo: find inverse
	}

	// Yacas' method
	b := Identity(len(a))
	c := a.Copy()
	for ; 0 < p; p >>= 1 {
		if p&1 == 1 {
			b = b.Multiply(c)
		}

		c = c.Multiply(c)
	}

	return b
}

// String ...
func (a Permutation) String() string {
	var sb strings.Builder
	sb.WriteString("(" + strconv.Itoa(a[0]))
	for i := 1; i < len(a); i++ {
		sb.WriteString("," + strconv.Itoa(a[i]))
	}

	sb.WriteString(")")
	return sb.String()
}
