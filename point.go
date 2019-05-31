package math

import "github.com/nathangreene3/math/stats"

// Point is an n-tuple.
type Point []float64

// Origin returns the zero-valued point of a given number of dimensions.
func Origin(dims int) Point {
	return make(Point, dims)
}

// ComparePoints returns -1, 0, or 1 indicating if x precedes, is equal to, or follows y.
func ComparePoints(x, y Point) int {
	m, n := len(x), len(y)
	min := stats.MinInt(m, n)
	for i := 0; i < min; i++ {
		if x[i] < y[i] {
			return -1
		}

		if y[i] < x[i] {
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

// equalPoints returns true if points are equal in dimension and in each value comparison.
func equalPoints(x, y Point) bool {
	return ComparePoints(x, y) == 0
}

// copyPoint returns a copy of a point.
func copyPoint(point Point) Point {
	cpy := make(Point, len(point))
	copy(cpy, point)
	return cpy
}
