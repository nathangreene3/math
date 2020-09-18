package stats

import (
	gomath "math"

	"github.com/nathangreene3/math"
)

// BinomPDF returns the probability of X = x successes in n Bernouli trials each with probability of success p.
func BinomPDF(x, n int, p float64) float64 {
	return float64(math.Choose(n, x)) * gomath.Pow(p, float64(x)) * gomath.Pow(1-p, float64(n-x))
}

// BinomCDF returns the cumulative probability of X = x0, ..., x1 successes in n Bernouli trials each with probability of success p.
func BinomCDF(x0, x1, n int, p float64) float64 {
	var prob float64
	for ; x0 <= x1; x0++ {
		prob += BinomPDF(x0, n, p)
	}

	return prob
}

// PoissonPDF ...
func PoissonPDF(x int, r float64) float64 {
	return gomath.Exp(float64(-x)) * gomath.Pow(r, float64(x)) / float64(math.Fact(x))
}

// PoissonCDF ...
func PoissonCDF(x0, x1 int, r float64) float64 {
	var prob float64
	for ; x0 <= x1; x0++ {
		prob += PoissonPDF(x0, r)
	}

	return prob
}
