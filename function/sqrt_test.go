package function

import (
	gomath "math"
	"math/rand"
	"testing"
)

func TestSqrt(t *testing.T) {
	for i := 0; i < 256; i++ {
		x := rand.Float64()
		if exp, rec := gomath.Sqrt(x), sqrtNewton2(x); exp != rec {
			t.Errorf("\n   given %f\nexpected %f\nreceived %f with abs error %e", x, exp, rec, gomath.Abs(exp-rec))
		}
	}
}
