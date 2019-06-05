package vector

import (
	"fmt"
	"math"

	"github.com/nathangreene3/math/stats"
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

// New generates a vector of dimension n with entries defined by a generating function f.
func New(n int, f Generator) Vector {
	v := make(Vector, 0, n)
	for i := 0; i < n; i++ {
		v = append(v, f(i))
	}

	return v
}

// At returns the ith value.
func (v Vector) At(i int) float64 {
	return v[i]
}

// Length returns |v|.
func Length(v Vector) float64 {
	return math.Sqrt(Dot(v, v))
}

// Length returns |v|.
func (v Vector) Length() float64 {
	return math.Sqrt(Dot(v, v))
}

// Add returns v+w.
func Add(v, w Vector) Vector {
	n := len(v)
	if n != len(w) {
		panic("dimension mismatch")
	}

	u := make(Vector, 0, n)
	for i := 0; i < n; i++ {
		u = append(u, v[i]+w[i])
	}

	return u
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

// Subtract w from v.
func (v Vector) Subtract(w Vector) {
	n := len(v)
	if n != len(w) {
		panic("dimension mismatch")
	}

	for i := 0; i < n; i++ {
		v[i] -= w[i]
	}
}

// Subtract returns v-w.
func Subtract(v, w Vector) Vector {
	return Add(v, Multiply(-1, w))
}

// Multiply returns av.
func Multiply(a float64, v Vector) Vector {
	w := make(Vector, 0, len(v))
	for i := range v {
		w = append(w, a*v[i])
	}

	return w
}

// Multiply each value by a.
func (v Vector) Multiply(a float64) {
	for i := range v {
		v[i] *= a
	}
}

// divide returns v/a.
func divide(a float64, v Vector) Vector {
	return Multiply(1.0/a, v)
}

// divide each value by a.
func (v Vector) divide(a float64) {
	v.Multiply(1.0 / a)
}

// Dot returns v dot w.
func Dot(v, w Vector) float64 {
	n := len(v)
	if n != len(w) {
		panic("dimension mismatch")
	}

	var s float64
	for i := 0; i < n; i++ {
		s += v[i] * w[i]
	}

	return s
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

// Unit returns v/|v|, a vector of length one pointing in the direction of v.
func Unit(v Vector) Vector {
	return Multiply(1.0/v.Length(), v)
}

// Unit returns v/|v|, a vector of length one pointing in the direction of v.
func (v Vector) Unit() Vector {
	return Multiply(1.0/v.Length(), v)
}

// Angle returns the angle between two vectors.
func Angle(v, w Vector) float64 {
	return math.Acos(Dot(Unit(v), Unit(w)))
}

// Angle returns the Angle between two vectors.
func (v Vector) Angle(w Vector) float64 {
	return math.Acos(v.Unit().Dot(w.Unit()))
}

// Projection returns the projection of w onto v (proj_v(w)).
func Projection(v, w Vector) Vector {
	lenv := w.Length()
	return Multiply(Dot(v, w)/(lenv*lenv), w)
}

// Projection returns the projection of w onto v (proj_v(w)).
func (v Vector) Projection(w Vector) Vector {
	r := v.Length()
	return Multiply(Dot(w, v)/(r*r), v)
}

// Rotate2D returns a vector rotated from v's position by an angle b in radians.
func Rotate2D(v Vector, angle float64) Vector {
	sin, cos := math.Sin(angle), math.Cos(angle)
	return Vector{
		v[0]*cos - v[1]*sin,
		v[0]*sin + v[1]*cos,
	}
}

// Rotate2D returns a vector rotated from v's position by an angle in radians.
func (v Vector) Rotate2D(angle float64) {
	v0, v1 := v[0], v[1]
	sin, cos := math.Sin(angle), math.Cos(angle)
	v[0], v[1] = v0*cos-v1*sin, v0*sin+v1*cos
}

// OrthonormalBasis returns the typical set of unit vectors spanning R^n.
func OrthonormalBasis(n int) []Vector {
	b := make([]Vector, 0, n)
	for i := 0; i < n; i++ {
		b = append(b, OrthonormalVector(i, n))
	}

	return b
}

// OrthonormalVector returns the vector (0,...,0,1,0,...,0) of length n and the ith value set to 1.
func OrthonormalVector(i, n int) Vector {
	v := make(Vector, n)
	v[i]++
	return v
}

// Copy a vector.
func (v Vector) Copy() Vector {
	w := make(Vector, len(v))
	copy(w, v)
	return w
}

// Compare returns -1, 0, or 1 indicating v precedes, is equal to, or follows w. Vectors v and w may be of different dimensions.
func Compare(v, w Vector) int {
	m, n := len(v), len(w)
	min := stats.MinInt(m, n)
	for i := 0; i < min; i++ {
		if v[i] < w[i] {
			return -1
		}

		if w[i] < v[i] {
			return 1
		}
	}

	if m < n {
		return -1
	}

	if n < m {
		return 1
	}

	return 0
}

// Compare returns -1, 0, or 1 indicating v precedes, is equal to, or follows w. Vectors v and w may be of different dimensions.
func (v Vector) Compare(w Vector) int {
	m, n := len(v), len(w)
	min := stats.MinInt(m, n)
	for i := 0; i < min; i++ {
		if v[i] < w[i] {
			return -1
		}

		if w[i] < v[i] {
			return 1
		}
	}

	if m < n {
		return -1
	}

	if n < m {
		return 1
	}

	return 0
}

// Equal returns the comparison v = w.
func Equal(v, w Vector) bool {
	return Compare(v, w) == 0
}

// Equal returns the comparison v = w.
func (v Vector) Equal(w Vector) bool {
	return v.Compare(w) == 0
}

// String returns a string-representation of a vector.
func (v Vector) String() string {
	return fmt.Sprintf("%0.3f", v) // TODO: Find the fastest way to stringify slices.
}

// IsMultipleOf returns true if either v or w is a multiple of the other (v = aw for some real a).
func IsMultipleOf(v, w Vector) bool {
	n := len(v)
	if n != len(w) {
		return false
	}

	vToW := v.Compare(w)
	if vToW == 0 || n == 0 {
		return true
	}

	z := make(Vector, n)
	vToZ := v.Compare(z)
	wToZ := w.Compare(z)
	if vToZ == 0 || wToZ == 0 {
		return false
	}

	v = v.Copy()
	w = w.Copy()
	if vToZ < 0 {
		v = Multiply(-1, v)
	}

	if wToZ < 0 {
		w = Multiply(-1, w)
	}

	vToW = v.Compare(w)
	if vToW < 0 {
		for {
			w.Subtract(v)
			if vToW = v.Compare(w); vToW == 0 {
				return true
			}

			if 0 < vToW {
				return false
			}
		}
	}

	if 0 < vToW {
		for {
			v.Subtract(w)
			if vToW = v.Compare(w); vToW == 0 {
				return true
			}

			if vToW < 0 {
				return false
			}
		}
	}

	return true
}
