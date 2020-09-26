package matrix2

import "testing"

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

func TestRemoveDups(t *testing.T) {
	tests := []struct {
		A, exp *Matrix
	}{
		// {
		// 	A:   New(3, 3, 1, 2, 3, 2, 4, 6, 3, 6, 9),
		// 	exp: New(1, 3, 1, 2, 3),
		// },
		// {
		// 	A:   New(3, 3, 1, 2, 3, 4, 5, 6, 2, 4, 6),
		// 	exp: New(2, 3, 1, 2, 3, 4, 5, 6),
		// },
		{
			A:   New(4, 5, 1, 2, 3, 4, 5, 1, 1, 1, 1, 1, 2, 4, 6, 8, 10, 2, 2, 2, 2, 2),
			exp: New(2, 5, 1, 2, 3, 4, 5, 1, 1, 1, 1, 1),
		},
	}

	for _, test := range tests {
		t.Fatalf("\n   given %s\nexpected %s\nreceived %s\n", test.A, test.exp, test.A.RemoveDupRows())
		if rec := test.A.RemoveDupRows(); !test.exp.Equal(rec) {
			t.Errorf("\n   given %s\nexpected %s\nreceived %s\n", test.A, test.exp, rec)
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
		A.ref()
		B.ref2()
		t.Errorf("\n%s\n%s\n", A, B)
		t.Errorf("\n%v\n", New(4, 4, 1, -1, 2, -1, 2, -2, 3, -3, 1, 1, 1, 0, 1, -1, 4, 3).Solve(NewVector(-8, -20, -2, 4)))
	}
}
