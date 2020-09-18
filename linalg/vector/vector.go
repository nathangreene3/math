package vector

import (
	gomath "math"
	"strconv"
	"strings"

	"github.com/nathangreene3/math"
)

// ------------------------------------------------------------------------------
// RESOURCES
// ------------------------------------------------------------------------------
// Most methods defined here are taken from or are inspired by Linear Algebra,
// 3rd Ed., by Stephen H. Friedberg, Arnold J. Insel, and Lawrence E. Spence.
// Any page references in comments are in reference to this source.
// ------------------------------------------------------------------------------

// Vector is an ordered n-tuple.
type Vector []float64

// Generator is a function defining the i-th entry of a vector.
type Generator func(i int) float64

// ------------------------------------------------------------------------------
// VECTOR CONSTRUCTORS
// ------------------------------------------------------------------------------

// Gen generates a vector of dimension n with entries defined by a generating
// function f.
func Gen(n int, f Generator) Vector {
	v := make(Vector, 0, n)
	for i := 0; i < n; i++ {
		v = append(v, f(i))
	}

	return v
}

// New values as a vector.
func New(values ...float64) Vector {
	return append(make(Vector, 0, len(values)), values...)
}

// Zeroes returns an n-dimensional vector with zeroes for all entries.
func Zeroes(n int) Vector {
	return make(Vector, n)
}

// ------------------------------------------------------------------------------
// EXPORTED OPERATIONS ON VECTORS
// ------------------------------------------------------------------------------

// Add several vectors to v.
func (v Vector) Add(ws ...Vector) {
	for i := 0; i < len(ws); i++ {
		if len(v) != len(ws[i]) {
			panic(dimErr)
		}

		for j := 0; j < len(v); j++ {
			v[j] += ws[i][j]
		}
	}
}

// Add returns the sum of several vectors.
func Add(vs ...Vector) Vector {
	switch len(vs) {
	case 0:
		return nil // TODO: Panic instead?
	case 1:
		return vs[0].Copy()
	default:
		v := vs[0].Copy()
		v.Add(vs[1:]...)
		return v
	}
}

// Angle returns the Angle between two vectors.
func (v Vector) Angle(w Vector) float64 {
	return gomath.Acos(Unit(v).Dot(Unit(w)))
}

// Approx returns true if v approximates w for a given tolerance on the range [0,1].
func (v Vector) Approx(w Vector, tol float64) bool {
	if tol < 0 || 1 < tol {
		panic("tolerance out of range")
	}

	if len(v) != len(w) {
		return false // TODO: Should this panic here?
	}

	for i := 0; i < len(v); i++ {
		if tol < gomath.Abs(v[i]-w[i]) {
			return false
		}
	}

	return true
}

// Compare returns -1, 0, or 1 indicating v precedes, is equal to, or follows
// w. Vectors v and w may be of different dimensions.
func (v Vector) Compare(w Vector) int {
	for i, min := 0, math.MinInt(len(v), len(w)); i < min; i++ {
		switch {
		case v[i] < w[i]:
			return -1
		case w[i] < v[i]:
			return 1
		}
	}

	switch {
	case len(v) < len(w):
		return -1
	case len(w) < len(v):
		return 1
	default:
		return 0
	}
}

// Compliment returns -v.
func Compliment(v Vector) Vector {
	return Multiply(v, -1)
}

// Compliment updates v as v *= -1 and returns v.
func (v Vector) Compliment() {
	v.Multiply(-1)
}

// Copy a vector.
func (v Vector) Copy() Vector {
	return append(make(Vector, 0, len(v)), v...)
}

// Cross returns the cross product of two vectors.
func (v Vector) Cross(w Vector) Vector {
	if len(v) != 3 || len(w) != 3 {
		panic("Cross: vectors must have length three")
	}

	return Vector{v[1]*w[2] - v[2]*w[1], -v[0]*w[2] + v[2]*w[0], v[0]*w[1] - v[1]*w[0]}
}

// Divide returns (1/a)*v.
func Divide(v Vector, a float64) Vector {
	w := v.Copy()
	w.Divide(a)
	return w
}

// Divide a vector v by a.
func (v Vector) Divide(a float64) {
	for i := 0; i < len(v); i++ {
		v[i] /= a
	}
}

// Dot returns v dot w.
func (v Vector) Dot(w Vector) float64 {
	if len(v) != len(w) {
		panic(dimErr)
	}

	var s float64
	for i := 0; i < len(v); i++ {
		s += v[i] * w[i]
	}

	return s
}

// Equal returns the comparison v = w.
func (v Vector) Equal(w Vector) bool {
	return v.Compare(w) == 0
}

// Format a vector as a string. TODO: Finish documentation.
func (v Vector) Format(fmt byte, prec int, left, right, sep rune) string {
	var sb strings.Builder
	sb.WriteRune(left)
	if 0 < len(v) {
		sb.WriteString(strconv.FormatFloat(v[0], fmt, prec, 64))
		for i := 1; i < len(v); i++ {
			sb.WriteRune(sep)
			sb.WriteString(strconv.FormatFloat(v[i], fmt, prec, 64))
		}
	}

	sb.WriteRune(right)
	return sb.String()
}

// IsMultOf returns true if either v or w is a multiple of the other
// (v = aw for some real a).
func (v Vector) IsMultOf(w Vector) bool {
	if v.Compare(w) == 0 {
		return true
	}

	if len(v) != len(w) {
		return false
	}

	// v and w are of the same dimension (n), but for one to be a
	// multiple of the other, all dimensions must either both be zero,
	// or neither be zero. This finds the first dimension i such that
	// v[i] and w[i] are both non-zero. Then it sets the ratio or
	// checks if ratios are consistent.
	var r float64 // Ratio of each non-zero dimension in v and w
	for i := 0; i < len(v); i++ {
		switch {
		case v[i] != 0:
			if w[i] != 0 {
				if 0 < r {
					if r != v[i]/w[i] {
						// Ratios not consistent
						return false
					}
				} else {
					// Ratio should be set only once
					r = v[i] / w[i]
				}
			} else {
				// v[i] != 0, but w[i] == 0
				return false
			}
		case w[i] != 0:
			// v[i] == 0, but w[i] != 0
			return false
		}
	}

	return true
}

// Len returns len(v).
func (v Vector) Len() int {
	return len(v)
}

// Mag returns |v|. This is NOT len(v).
func (v Vector) Mag() float64 {
	return gomath.Sqrt(v.Dot(v))
}

// Multiply returns a*v.
func Multiply(v Vector, a float64) Vector {
	w := v.Copy()
	w.Multiply(a)
	return w
}

// Multiply a vector v by a scalar a and return v.
func (v Vector) Multiply(a float64) {
	for i := 0; i < len(v); i++ {
		v[i] *= a
	}
}

// Set a vector.
func (v Vector) Set(f Generator) {
	for i := 0; i < len(v); i++ {
		v[i] = f(i)
	}
}

// String returns the default string-representation of a vector.
func (v Vector) String() string {
	return v.Format('f', -1, '[', ']', ' ')
}

// Subtract returns v-w.
func (v Vector) Subtract(ws ...Vector) {
	for i := 0; i < len(ws); i++ {
		if len(v) != len(ws[i]) {
			panic(dimErr)
		}

		for j := 0; j < len(v); j++ {
			v[j] -= ws[i][j]
		}
	}
}

// Subtract returns v - ws0 - ws1 - ...
func Subtract(v Vector, ws ...Vector) Vector {
	w := v.Copy()
	w.Subtract(ws...)
	return w
}

// Proj returns the projection of w onto v (proj_v(w)).
func (v Vector) Proj(w Vector) {
	r := v.Mag()
	v.Multiply(v.Dot(w) / (r * r))
}

// Proj ...
func Proj(v, w Vector) Vector {
	x := v.Copy()
	x.Proj(w)
	return x
}

// Unit returns v/|v|, a vector of length one pointing in the direction of v.
func (v Vector) Unit() {
	v.Divide(v.Mag())
}

// Unit ...
func Unit(v Vector) Vector {
	w := v.Copy()
	w.Unit()
	return w
}
