package vector

import (
	"fmt"
	"math"

	"github.com/nathangreene3/math/stats"
)

// Vector is an ordered n-tuple.
type Vector []float64

// Length returns |v|.
func (v Vector) Length() float64 {
	return math.Sqrt(Dot(v, v))
}

// Multiply returns av.
func Multiply(a float64, v Vector) Vector {
	w := make(Vector, 0, len(v))
	for i := range v {
		w = append(w, a*v[i])
	}

	return w
}

// Unit returns v/|v|, a vector of length one pointing in the direction of v.
func Unit(v Vector) Vector {
	return Multiply(1.0/v.Length(), v)
}

// Add returns u+v.
func Add(u, v Vector) Vector {
	n := len(u)
	if n != len(v) {
		panic("dimension mismatch")
	}

	w := make(Vector, 0, n)
	for i := 0; i < n; i++ {
		w = append(w, u[i]+v[i])
	}

	return w
}

// Subtract returns u-v.
func Subtract(u, v Vector) Vector {
	return Add(u, Multiply(-1, v))
}

// Dot returns u dot v.
func Dot(u, v Vector) float64 {
	n := len(u)
	if n != len(v) {
		panic("dimension mismatch")
	}

	var s float64
	for i := 0; i < n; i++ {
		s += u[i] * v[i]
	}

	return s
}

// Projection returns the projection ov u onto v (proj_v(u)).
func Projection(u, v Vector) Vector {
	lenv := v.Length()
	return Multiply(Dot(u, v)/(lenv*lenv), v)
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

// Angle returns the angle between two vectors.
func Angle(u, v Vector) float64 {
	return math.Acos(Dot(Unit(u), Unit(v)))
}

// Compare returns -1, 0, or 1 indicating u precedes, is equal to, or follows v. Vectors u and v may be of different lengths.
func Compare(u, v Vector) int {
	m, n := len(u), len(v)
	min := stats.MinInt(m, n)
	for i := 0; i < min; i++ {
		if u[i] < v[i] {
			return -1
		}

		if v[i] < u[i] {
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

// Equal returns the comparison u = v.
func Equal(u, v Vector) bool {
	return Compare(u, v) == 0
}

// String returns a string-representation of a vector.
func (v Vector) String() string {
	return fmt.Sprintf("%0.3f", v) // TODO: Find the fastest way to stringify slices.
}
