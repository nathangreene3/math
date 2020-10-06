package matrix2

import (
	"fmt"
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

// Dims ...
func (v *Vector) Dims() int {
	return len(v.vec)
}

// String ...
func (v *Vector) String() string {
	return fmt.Sprint(v.vec)
}

// Values ...
func (v *Vector) Values() []float64 {
	return append(make([]float64, 0, len(v.vec)), v.vec...)
}
