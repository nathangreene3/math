package function

import (
	"fmt"
	gomath "math"
	"testing"
)

func TestDiff(t *testing.T) {
	var (
		f     Function = func(x float64) float64 { return x * gomath.Exp(x) }       // x * x }
		dfdx  Function = func(x float64) float64 { return gomath.Exp(x) * (x + 1) } // 2 * x }
		x0, h float64  = 1, 0.1
	)

	for n := 1; n <= 10; n++ {
		t.Errorf("Error n=%d: %.16f", n, gomath.Abs(dfdx(x0)-Diff(f, x0, h, n)))
		t.Errorf("Error n=%d: %.16f", n, gomath.Abs(dfdx(x0)-Diff2(f, x0, h, n)))
	}
}

func BenchmarkDiff(b *testing.B) {
	var (
		f    Function = func(x float64) float64 { return x * gomath.Exp(x) }
		x, h float64  = 1, 0.1
	)
	for n := 1; n <= 10; n++ {
		benchmarkDiff(b, f, x, h, n)
		benchmarkDiff2(b, f, x, h, n)
	}
}

func benchmarkDiff(b *testing.B, f Function, x, h float64, n int) bool {
	g := func(b0 *testing.B) {
		for i := 0; i < b0.N; i++ {
			_ = Diff(f, x, h, n)
		}
	}

	return b.Run(fmt.Sprintf(" Diff(f,%f,%f,%d)", x, h, n), g)
}

func benchmarkDiff2(b *testing.B, f Function, x, h float64, n int) bool {
	g := func(b0 *testing.B) {
		for i := 0; i < b0.N; i++ {
			_ = Diff2(f, x, h, n)
		}
	}

	return b.Run(fmt.Sprintf("Diff2(f,%f,%f,%d)", x, h, n), g)
}
