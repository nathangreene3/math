package matrix2

import (
	"fmt"
	gomath "math"

	"github.com/nathangreene3/math"
)

// Generator ...
type Generator func(i int) float64

// Vector ...
type Vector struct {
	vec []float64
}

// Gen ...
func Gen(n int, f Generator) *Vector {
	vec := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		vec = append(vec, f(i))
	}

	return &Vector{vec: vec}
}

// New ...
func New(values ...float64) *Vector {
	return &Vector{vec: append(make([]float64, 0, len(values)), values...)}
}

// Zeroes ...
func Zeroes(n int) *Vector {
	return &Vector{vec: make([]float64, n)}
}

// Add ...
func (v *Vector) Add(vs ...*Vector) {
	for i := 0; i < len(vs); i++ {
		if len(v.vec) != len(vs[i].vec) {
			panic("dimension mismatch")
		}

		for j := 0; j < len(v.vec); j++ {
			v.vec[j] += vs[i].vec[j]
		}
	}
}

// Angle ...
func (v *Vector) Angle(w *Vector) float64 {
	return gomath.Acos(Unit(v).Dot(Unit(w)))
}

// Compare ...
func (v *Vector) Compare(w *Vector) int {
	for i, min := 0, math.MinInt(len(v.vec), len(w.vec)); i < min; i++ {
		switch {
		case v.vec[i] < w.vec[i]:
			return -1
		case w.vec[i] < v.vec[i]:
			return 1
		}
	}

	switch {
	case len(v.vec) < len(w.vec):
		return -1
	case len(w.vec) < len(v.vec):
		return 1
	default:
		return 0
	}
}

// Copy ...
func (v *Vector) Copy() *Vector {
	vec := make([]float64, 0, len(v.vec))
	for i := 0; i < len(v.vec); i++ {
		vec = append(vec, v.vec[i])
	}

	return &Vector{vec: vec}
}

// Cross ...
func (v *Vector) Cross(w *Vector) *Vector {
	if len(v.vec) != 3 || len(w.vec) != 3 {
		panic("vectors must be of dimension 3")
	}

	return &Vector{vec: []float64{v.vec[1]*w.vec[2] - v.vec[2]*w.vec[1], -v.vec[0]*w.vec[2] + v.vec[2]*w.vec[0], v.vec[0]*w.vec[1] - v.vec[1]*w.vec[0]}}
}

// Dims ...
func (v *Vector) Dims() int {
	return len(v.vec)
}

// Div ...
func (v *Vector) Div(a float64) {
	for i := 0; i < len(v.vec); i++ {
		v.vec[i] /= a
	}
}

// Dot ...
func (v *Vector) Dot(w *Vector) float64 {
	if len(v.vec) != len(w.vec) {
		panic("dimension mismatch")
	}

	var d float64
	for i := 0; i < len(v.vec); i++ {
		d += v.vec[i] * w.vec[i]
	}

	return d
}

// Equal ...
func (v *Vector) Equal(w *Vector) bool {
	return v.Compare(w) == 0
}

// Mag ...
func (v *Vector) Mag() float64 {
	return gomath.Sqrt(v.Dot(v))
}

// Mult ...
func (v *Vector) Mult(a float64) {
	for i := 0; i < len(v.vec); i++ {
		v.vec[i] *= a
	}
}

// MultOf ...
func (v *Vector) MultOf(w *Vector) bool {
	switch {
	case len(v.vec) != len(w.vec):
		return false
	case v.Compare(w) == 0:
		return true
	default:
		// v and w are of the same dimension (n), but for one to be a
		// multiple of the other, all dimensions must either both be zero,
		// or neither be zero. This finds the first dimension i such that
		// v[i] and w[i] are both non-zero. Then it sets the ratio or
		// checks if ratios are consistent.
		var r float64 // Ratio of each non-zero dimension in v and w
		for i := 0; i < len(v.vec); i++ {
			switch {
			case v.vec[i] != 0:
				switch {
				case w.vec[i] == 0:
					// v[i] != 0, but w[i] == 0
					return false
				case 0 < r:
					if r != v.vec[i]/w.vec[i] {
						// Ratios not consistent
						return false
					}
				default:
					// Ratio should be set only once
					r = v.vec[i] / w.vec[i]
				}
			case w.vec[i] != 0:
				// v[i] == 0, but w[i] != 0
				return false
			}
		}

		return true
	}
}

// String ...
func (v *Vector) String() string {
	return fmt.Sprint(v.vec)
}

// Sub ...
func (v *Vector) Sub(vs ...*Vector) {
	for i := 0; i < len(vs); i++ {
		if len(v.vec) != len(vs[i].vec) {
			panic("dimension mismatch")
		}

		for j := 0; j < len(v.vec); j++ {
			v.vec[j] -= vs[i].vec[j]
		}
	}
}

// Values ...
func (v *Vector) Values() []float64 {
	return append(make([]float64, 0, len(v.vec)), v.vec...)
}

// Unit ...
func Unit(v *Vector) *Vector {
	cpy := v.Copy()
	cpy.Div(v.Mag())
	return cpy
}

// Unit ...
func (v *Vector) Unit() {
	v.Div(v.Mag())
}
