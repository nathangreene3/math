package stats

import gomath "math"

// NormPDF ...
func NormPDF(x, mean, stDev float64) float64 {
	y := (x - mean) / stDev
	return gomath.Exp(-0.5*y*y) / (gomath.Sqrt2 * gomath.SqrtPi * stDev)
}

// StNormPDF ...
func StNormPDF(x float64) float64 {
	return gomath.Exp(-x*x/2) / (gomath.Sqrt2 * gomath.SqrtPi)
}
