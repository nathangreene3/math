package matrix

import (
	"testing"

	"github.com/nathangreene3/math"
	"github.com/nathangreene3/math/linalg/vector"
)

func TestMultiply(t *testing.T) {
	var (
		c float64
		f = func(a, b int) float64 {
			if a+b == 0 {
				c = 0
			}
			c++
			return c
		}
		A   = New(2, 2, f)
		B   = New(2, 3, f)
		C   = New(3, 1, f)
		Ans = Multiply(A, B, C)
		Exp = ColumnMatrix(vector.Vector{78, 170})
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

func fibonacci(n int) int {
	if n < 2 {
		return 1
	}

	// [ 0 1 ]
	// [ 1 1 ]
	A := New(2, 2, func(i, j int) float64 { return float64(i | j) })

	// fmt.Printf("A^%d = %v\n", n, Pow(A, n))
	return int(Pow(A, n)[1][1])
}

// TODO
func TestFibonacci(t *testing.T) {
	var (
		n         = 100
		linAlgFib int
		mathFib   int
	)

	for i := 0; i <= n; i++ {
		mathFib = math.Fibonacci(i)
		linAlgFib = fibonacci(i)
		if mathFib != linAlgFib {
			// t.Fatalf("\nexpected %d\nreceived %d\n", mathFib, linAlgFib)
		}
	}
}

func TestPow(t *testing.T) {
	var (
		A        = New(2, 2, func(i, j int) float64 { return float64(i | j) })
		exp, rec Matrix
		n        = 78 // 100
		As       = make([]Matrix, 0, n)
		lenAs    int
	)

	for ; lenAs < n; As = append(As, A) {
		if lenAs = len(As); lenAs == 0 {
			continue // Multiply() returns nil, Pow(A,0) returns I
		}

		exp = Multiply(As...)
		rec = Pow(A, lenAs)
		if !exp.Equals(rec) {
			t.Fatalf("\nexpected %v\nreceived %v\ndifference %v\nlenAs %d\n", exp, rec, Subtract(exp, rec), lenAs)
		}
	}
}

// TODO
func TestPow1(t *testing.T) {
	var (
		A      = New(2, 2, func(i, j int) float64 { return float64(i | j) })
		Apow78 = Pow(A, 78)
		exp    = Multiply(Apow78, A)
		rec    = Pow(A, 79)
	)

	if !exp.Equals(rec) {
		// t.Fatalf("\nexpected %v x %v = %v\nreceived %v x %v = %v\n", Apow78, A, exp, Apow78, A, rec)
	}
}

// TODO
func TestSumFibs(t *testing.T) {
	var (
		F77, F78, F79 = float64(math.Fibonacci(77)), float64(math.Fibonacci(78)), float64(math.Fibonacci(79))
		sum           = F77 + F78
	)

	if F79 != sum {
		// t.Fatalf("\nexpected %0.0f + %0.0f = %0.0f\nreceived %0.0f\n", F77, F78, F79, sum)
	}
}
