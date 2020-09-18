package function

import (
	"math"
	"testing"
)

func TestBisect(t *testing.T) {
	var (
		x0  = math.SmallestNonzeroFloat64
		x1  = math.MaxFloat64
		tol = 0.000001
		f   = func(x float64) float64 { return x*x - 1 }
	)

	t.Error(Bisect(f, x0, x1, tol))
}
