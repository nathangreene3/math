package vector

import (
	"fmt"
	gomath "math"

	"github.com/nathangreene3/math"
)

// ------------------------------------------------------------------
// RESOURCES
// ------------------------------------------------------------------
// Most methods defined here are taken from or are inspired by
// Linear Algebra, 3rd Ed., by Stephen H. Friedberg, Arnold J. Insel,
// and Lawrence E. Spence. Any page references in comments are in
// reference to this source.
// ------------------------------------------------------------------

// Vector is an ordered n-tuple.
type Vector []float64

// Generator is a function defining the i-th entry of a vector.
type Generator func(i int) float64

// ------------------------------------------------------------------
// VECTOR CONSTRUCTORS
// ------------------------------------------------------------------

// New generates a vector of dimension n with entries defined by a
// generating function f.
func New(n int, f Generator) Vector {
	v := make(Vector, 0, n)
	for i := 0; i < n; i++ {
		v = append(v, f(i))
	}

	return v
}

// Zero returns the zero vector of n dimensions.
func Zero(n int) Vector {
	return make(Vector, n)
}

// ------------------------------------------------------------------
// EXPORTED OPERATIONS ON VECTORS
// ------------------------------------------------------------------

// Approx returns true if v approximates w for a given precision on the range [0,1].
func (v Vector) Approx(w Vector, prec float64) bool {
	n := len(v)
	if n != len(w) {
		return false
	}

	for i := 0; i < n; i++ {
		if !math.Approx(v[i], w[i], prec) {
			return false
		}
	}

	return true
}

// Add returns v+w.
func Add(v, w Vector) Vector {
	n := len(v)
	if n != len(w) {
		panic("dimension mismatch")
	}

	return New(n, func(i int) float64 { return v[i] + w[i] })
}

// Add w to v.
func (v Vector) Add(w Vector) {
	n := len(v)
	if n != len(w) {
		panic("dimension mismatch")
	}

	for i := 0; i < n; i++ {
		v[i] += w[i]
	}
}

// Angle returns the Angle between two vectors.
func (v Vector) Angle(w Vector) float64 {
	return gomath.Acos(v.Unit().Dot(w.Unit()))
}

// CompareTo returns -1, 0, or 1 indicating v precedes, is equal to,
// or follows w. Vectors v and w may be of different dimensions.
func (v Vector) CompareTo(w Vector) int {
	m, n := len(v), len(w)
	min := math.MinInt(m, n)
	for i := 0; i < min; i++ {
		switch {
		case v[i] < w[i]:
			return -1
		case w[i] < v[i]:
			return 1
		}
	}

	switch {
	case m < n:
		return -1
	case n < m:
		return 1
	default:
		return 0
	}
}

// Complimentary returns -v.
func (v Vector) Complimentary() Vector {
	return Multiply(-1, v)
}

// Copy a vector.
func (v Vector) Copy() Vector {
	w := make(Vector, len(v))
	copy(w, v)
	return w
}

// Dimensions returns len(v).
func (v Vector) Dimensions() int {
	return len(v)
}

// Divide returns v/a.
func Divide(a float64, v Vector) Vector {
	return New(len(v), func(i int) float64 { return v[i] / a })
}

// Divide each value by a.
func (v Vector) Divide(a float64) {
	for i := range v {
		v[i] /= a
	}
}

// Dot returns v dot w.
func (v Vector) Dot(w Vector) float64 {
	n := len(w)
	if n != len(v) {
		panic("dimension mismatch")
	}

	var s float64
	for i := 0; i < n; i++ {
		s += w[i] * v[i]
	}

	return s
}

// Equal returns the comparison v = w.
func (v Vector) Equal(w Vector) bool {
	return v.CompareTo(w) == 0
}

// IsMultipleOf returns true if either v or w is a multiple of the
// other (v = aw for some real a).
func (v Vector) IsMultipleOf(w Vector) bool {
	if v.CompareTo(w) == 0 {
		return true
	}

	n := len(v)
	if n != len(w) {
		return false
	}

	// v and w are of the same dimension (n), but for one to be a
	// multiple of the other, all dimensions must either both be zero,
	// or neither be zero. This finds the first dimension i such that
	// v[i] and w[i] are both non-zero. Then it sets the ratio or
	// checks if ratios are consistent.
	var r float64 // Ratio of each non-zero dimension in v and w
	for i := 0; i < n; i++ {
		switch {
		case v[i] != 0:
			switch {
			case w[i] != 0:
				if 0 < r {
					if r != v[i]/w[i] {
						return false // Ratios not consistent
					}
				} else {
					r = v[i] / w[i] // Ratio should be set only once
				}
			default:
				return false // v[i] != 0, but w[i] == 0
			}
		case w[i] != 0:
			return false // v[i] == 0, but w[i] != 0
		}
	}

	return true
}

// Length returns |v|. This is NOT len(v).
func (v Vector) Length() float64 {
	return gomath.Sqrt(v.Dot(v))
}

// Multiply returns av.
func Multiply(a float64, v Vector) Vector {
	return New(len(v), func(i int) float64 { return a * v[i] })
}

// Multiply each value by a.
func (v Vector) Multiply(a float64) {
	for i := range v {
		v[i] *= a
	}
}

// String returns a string-representation of a vector.
func (v Vector) String() string {
	return fmt.Sprintf("%0.3f", v) // TODO: Find the fastest way to stringify slices.
}

// Subtract returns v-w.
func Subtract(v, w Vector) Vector {
	n := len(v)
	if n != len(w) {
		panic("dimension mismatch")
	}

	return New(n, func(i int) float64 { return v[i] - w[i] })
}

// Subtract w from v.
func (v Vector) Subtract(w Vector) {
	v.Add(Multiply(-1, w))
}

// Projection returns the projection of w onto v (proj_v(w)).
func (v Vector) Projection(w Vector) Vector {
	r := v.Length()
	return Multiply(w.Dot(v)/(r*r), v)
}

// Unit returns v/|v|, a vector of length one pointing in the
// direction of v.
func (v Vector) Unit() Vector {
	return Divide(v.Length(), v)
}
