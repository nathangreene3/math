package matrix

import (
	"testing"

	"github.com/nathangreene3/math/linalg/vector"
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
		A   = New(2, 2, f)
		B   = New(2, 3, f)
		C   = New(3, 1, f)
		Ans = Multiply(A, B, C)
		Exp = New(2, 1, func(i, j int) float64 {
			if i == 0 {
				return 78
			}
			return 170
		})
	)
	if !Ans.Equals(Exp) {
		t.Fatalf("\nexpected %s\nreceived %s", Exp.String(), Ans.String())
	}

	// An example from 50 Mathematical Ideas
	A = Matrix{
		vector.Vector{7, 5, 0, 1},
		vector.Vector{0, 4, 3, 7},
		vector.Vector{3, 2, 0, 2},
	}
	B = ColumnMatrix(vector.Vector{3, 9, 8, 2})
	Exp = ColumnMatrix(vector.Vector{68, 74, 31})
	Ans = A.multiply(B)
	if !Exp.Equals(Ans) {
		t.Fatalf("\nexpected %s\nreceived %s", Ans.String(), Exp.String())
	}
}

func TestSolve(t *testing.T) {
	var (
		c    float64
		A    Matrix
		x, y vector.Vector
	)
	A = New(2, 2, func(i, j int) float64 {
		c++
		return c
	})
	x = A.Solve(vector.New(2, func(i int) float64 {
		c++
		return c
	}))
	y = vector.New(2, func(i int) float64 {
		if i == 0 {
			return -4
		}
		return 4.5
	})
	if !x.Equal(y) {
		t.Fatalf("expected %v, received %v", y, x)
	}

	// Function converting Celsius to Farenheit
	A = New(2, 2, func(i, j int) float64 {
		if j == 0 {
			if i == 0 {
				return 0
			}
			return 100
		}
		return 1
	})
	x = A.Solve(vector.New(2, func(i int) float64 {
		if i == 0 {
			return 32
		}
		return 212
	}))
	y = vector.New(2, func(i int) float64 {
		if i == 0 {
			return 1.8
		}
		return 32
	})
	if !x.Equal(y) {
		t.Fatalf("expected %v, received %v", y, x)
	}
}
