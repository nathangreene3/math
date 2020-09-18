package diffeq

import (
	"strconv"
	"strings"
)

// A Point is a vector p = (p0, p1, ...) in a phase space defined by a system of ODEs.
type Point []float64

// Add returns x+y.
func Add(x, y Point) Point {
	n := len(x)
	if n != len(y) {
		panic("dimension mismatch")
	}

	z := make(Point, 0, n)
	for i := 0; i < n; i++ {
		z = append(z, x[i]+y[i])
	}

	return z
}

// Copy a point.
func (p Point) Copy() Point {
	cpy := make(Point, len(p))
	copy(cpy, p)
	return cpy
}

// Divide returns x/a.
func Divide(a float64, x Point) Point {
	y := make(Point, 0, len(x))
	for i := 0; i < len(x); i++ {
		y = append(y, x[i]/a)
	}

	return y
}

// Multiply returns a*x.
func Multiply(a float64, x Point) Point {
	y := make(Point, 0, len(x))
	for i := 0; i < len(x); i++ {
		y = append(y, a*x[i])
	}

	return y
}

// String representation of a point.
func (p Point) String() string {
	if len(p) == 0 {
		return "()"
	}

	var sb strings.Builder
	sb.WriteString("(" + strconv.FormatFloat(p[0], 'g', -1, 64))
	for i := 1; i < len(p); i++ {
		sb.WriteString("," + strconv.FormatFloat(p[i], 'g', -1, 64))
	}

	return sb.String() + ")"
}

// Subtract returns x-y.
func Subtract(x, y Point) Point {
	n := len(x)
	if n != len(y) {
		panic("dimension mismatch")
	}

	z := make(Point, 0, n)
	for i := 0; i < n; i++ {
		z = append(z, x[i]-y[i])
	}

	return z
}
