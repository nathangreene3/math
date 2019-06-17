package matrix

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	// tests:=[]struct{
	// 	A Matrix
	// 	B Matrix
	// 	expected Matrix
	// 	actual Matrix
	// }{
	// 	{
	// 		A:Matrix{
	// 			vector.Vector{1,2,3,4},
	// 			vector.Vector{2,1,-1,3},
	// 			vector.Vector{4,0,1,2},
	// 		},
	// 		B:Matrix{

	// 			vector.Vector{2,1,-1,3},
	// 			vector.Vector{,,,},
	// 			vector.Vector{,,,},
	// 		}
	// 	}
	// }
}

func TestChainMultiply(t *testing.T) {
	// A := Empty(2, 10)
	// B := Empty(10, 3)
	// C := Empty(3, 8)
	var c float64
	A := New(2, 2, func(a, b int) float64 {
		if a == 0 && b == 0 {
			c = 0
		}

		c++
		return c
	})
	B := New(2, 3, func(a, b int) float64 {
		if a == 0 && b == 0 {
			c = 0
		}

		c++
		return c
	})
	C := New(3, 1, func(a, b int) float64 {
		if a == 0 && b == 0 {
			c = 0
		}

		c++
		return c
	})
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
