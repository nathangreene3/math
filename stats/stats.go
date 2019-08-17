package stats

import "math"

// Sum returns the sum of a list of values.
func Sum(x []float64) float64 {
	var s float64
	for i := range x {
		s += x[i]
	}

	return s
}

// Mean returns the Mean (or average) of a list of values.
func Mean(x []float64) float64 {
	return Sum(x) / float64(len(x))
}

// Var returns the Var of a list of values.
func Var(x []float64) float64 {
	m := Mean(x)
	var v, t float64
	for i := range x {
		t = x[i] - m
		v += t * t
	}

	return v / float64(len(x)-1)
}

// StDev returns the standard deviation of a list of values.
func StDev(x []float64) float64 {
	return math.Sqrt(Var(x))
}

// CoVar returns the covariance of two sets of values.
func CoVar(x, y []float64) float64 {
	n := len(x)
	if n != len(y) {
		panic("dimension mismatch")
	}

	mx, my := Mean(x), Mean(y)
	var cv float64
	for i := 0; i < n; i++ {
		cv += (x[i] - mx) * (y[i] - my)
	}

	return cv / float64(n)
}

// Max returns the maximum of x and y.
func Max(x, y float64) float64 {
	if x < y {
		return y
	}

	return x
}

// Min returns the minimum of x and y.
func Min(x, y float64) float64 {
	if x < y {
		return x
	}

	return y
}

// MaxInt returns the maximum of x and y.
func MaxInt(x, y int) int {
	if x < y {
		return y
	}

	return x
}

// MinInt returns the minimum of x and y.
func MinInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}
