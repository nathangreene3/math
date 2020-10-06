package matrix2

import (
	"fmt"
	gomath "math"
	"testing"

	vtr "github.com/nathangreene3/math/linalg/vector2"
	"github.com/nathangreene3/math/physconst"
)

func TestMult(t *testing.T) {
	{ // Pow
		A := New(2, 2, 1, 2, 3, 4)
		powToExp := map[int]*Matrix{
			// TODO: Negative powers
			0: Identity(2),
			1: A.Copy(),
			2: New(2, 2, 7, 10, 15, 22),
			3: New(2, 2, 37, 54, 81, 118),
			4: New(2, 2, 199, 290, 435, 634),
			5: New(2, 2, 1069, 1558, 2337, 3406),
			6: New(2, 2, 5743, 8370, 12555, 18298),
			7: New(2, 2, 30853, 44966, 67449, 98302),
			8: New(2, 2, 165751, 241570, 362355, 528106),
		}

		for p, exp := range powToExp {
			if rec := Pow(A, p); !exp.Equal(rec) {
				t.Errorf("\n   given %s^%d\nexpected %s\nreceived %s\n", A, p, exp, rec)
			}
		}
	}

	{
		// [1 2 3] [1 2]   [22 28]
		// [4 5 6]x[3 4] = [49 64]
		//         [5 6]
		var (
			c0, c1 float64
			A      = Gen(2, 3, func(i, j int) float64 { c0++; return c0 })
			B      = Gen(3, 2, func(i, j int) float64 { c1++; return c1 })
			exp    = New(2, 2, 22, 28, 49, 64)
			rec    = Mult(A, B)
		)

		if !exp.Equal(rec) {
			t.Errorf("\n   given %sx%s\nexpected %s\nreceived %s\n", A, B, exp, rec)
		}
	}

	{ // AN example from 50 Mathematical Ideas
		var (
			A   = New(3, 4, 7, 5, 0, 1, 0, 4, 3, 7, 3, 2, 0, 2)
			B   = New(4, 1, 3, 9, 8, 2)
			exp = New(3, 1, 68, 74, 31)
			rec = A.Mult(B)
		)

		if !exp.Equal(rec) {
			t.Errorf("\n   given %sx%s\nexpected %s\nreceived %s\n", A, B, exp, rec)
		}
	}
}

func TestREF(t *testing.T) {
	{
		// Fahrenheit to Celsius
		// A := New(2, 3, 212, 1, 100, 32, 1, 0)
	}

	{
		// Numerical Analysis, 7th Ed.
		// Ex 3, p 350
		//
		// [1  -1  2  -1 :  -8]
		// [2  -2  3  -3 : -20]
		// [1   1  1   0 :  -2]
		// [1  -1  4   3 :   4]

		A := New(4, 5, 1, -1, 2, -1, -8, 2, -2, 3, -3, -20, 1, 1, 1, 0, -2, 1, -1, 4, 3, 4)
		B := A.Copy()
		B.ref2()
		t.Errorf("\n%s\n%s\n", A, B)
		t.Errorf("\n%v\n", New(4, 4, 1, -1, 2, -1, 2, -2, 3, -3, 1, 1, 1, 0, 1, -1, 4, 3).Solve(vtr.New(-8, -20, -2, 4)))
	}

	{
		A := New(2, 2, 1, 2, 3, 4)
		A.Inverse()
		t.Errorf("\n%v\n", A)
	}

	{
		// A pion has kinetic energy of 90.0 MeV and rest energy 140. MeV. It
		// decays into two photons scattering at two potentially different
		// angles with respect to the direction of motion of the pion.
		var (
			K, E0          float64 = 90, 140
			theta0, theta1 float64 = 0, -gomath.Pi // Only some values work here: (0,-pi), (theta0,theta0) for some specific theta0.
			p                      = New(2, 2, gomath.Cos(theta0), gomath.Cos(theta1), 1, 1).Solve(vtr.New(gomath.Sqrt(gomath.Pow(K+E0, 2)-gomath.Pow(E0, 2))/physconst.SpeedOfLight, E0/physconst.SpeedOfLight))
		)

		t.Errorf("\n%v\n", p)
	}
}

func TestTrans(t *testing.T) {
	tests := []struct {
		A, exp *Matrix
	}{
		{
			A:   New(0, 0),
			exp: New(0, 0),
		},
		{
			A:   New(1, 1, 1),
			exp: New(1, 1, 1),
		},
		{
			A:   New(2, 2, 1, 2, 3, 4),
			exp: New(2, 2, 1, 3, 2, 4),
		},
		{
			A:   New(2, 3, 1, 2, 3, 4, 5, 6),
			exp: New(3, 2, 1, 4, 2, 5, 3, 6),
		},
	}

	for _, test := range tests {
		if rec := test.A.Trans(); !test.exp.Equal(rec) {
			t.Fatalf("\n   given %s\nexpected %s\nreceived %s", test.A, test.exp, rec)
		}
	}
}

func TestLU(t *testing.T) {
	tests := []struct {
		A, L, U *Matrix
	}{
		{
			A: New(4, 4, 1, 1, 0, 3, 2, 1, -1, 1, 3, -1, -1, 2, -1, 2, 3, -1),
			L: New(4, 4, 1, 0, 0, 0, 2, 1, 0, 0, 3, 4, 1, 0, -1, -3, 0, 1),
			U: New(4, 4, 1, 1, 0, 3, 0, -1, -1, -5, 0, 0, 3, 13, 0, 0, 0, -13),
		},
	}

	for _, test := range tests {
		var (
			L, U = test.A.LU(Doolitle)
			LU   = test.L.Mult(test.U)
		)

		if !test.L.Equal(L) {
			t.Errorf("\n   given A = %s\nexpected L = %s\nreceived L = %s\n", test.A, test.L, L)
		}

		if !test.U.Equal(U) {
			t.Errorf("\n   given A = %s\nexpected U = %s\nreceived U = %s\n", test.A, test.U, U)
		}

		if !test.A.Equal(LU) {
			t.Errorf("\n    given L = %s and U = %s\nexpected LU = %s\nreceived LU = %s\n", L, U, test.A, LU)
		}

		fmt.Printf("\n A = %s\nLU = %s\n", test.A, LU)
	}
}
