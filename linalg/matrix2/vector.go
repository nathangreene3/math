package matrix2

import "fmt"

// Vector ...
type Vector struct {
	vec []float64
}

// NewVector ...
func NewVector(values ...float64) *Vector {
	return &Vector{vec: append(make([]float64, 0, len(values)), values...)}
}

// ColMatrix ...
func (v *Vector) ColMatrix() *Matrix {
	return New(len(v.vec), 1, v.vec...)
}

// RowMatrix ...
func (v *Vector) RowMatrix() *Matrix {
	return New(1, len(v.vec), v.vec...)
}

// String ...
func (v *Vector) String() string {
	return fmt.Sprint(v.vec)
}
