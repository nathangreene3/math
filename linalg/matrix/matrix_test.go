package matrix

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	var (
		c float64
		f = func(a, b int) float64 {
			if a == 0 && b == 0 {
				c = 0
			}
			c++
			return c
		}
	)
	A := New(2, 2, f)
	B := New(2, 3, f)
	C := New(3, 1, f)
	D := Multiply(A, B, C)
	E := New(2, 1, func(i, j int) float64 {
		if i == 0 {
			return 78
		}
		return 170
	})
	if !D.Equals(E) {
		t.Fatalf("\nexpected %s\nreceived %s", E.String(), D.String())
	}
}
