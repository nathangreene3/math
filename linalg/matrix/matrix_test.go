package matrix

import (
	"fmt"
	"testing"

	"github.com/nathangreene3/math"
	vtr "github.com/nathangreene3/math/linalg/vector"
)

func TestJoin(t *testing.T) {
	A := New(
		vtr.New(1, 2),
		vtr.New(3, 4),
	)

	B := New(
		vtr.New(5),
		vtr.New(6),
	)

	C := New(
		vtr.New(7, 8, 9),
		vtr.New(10, 11, 12),
	)

	fmt.Println(Join(A, B, C))
}
func TestGenerators(t *testing.T) {
	// TODO
}

func TestMultiply(t *testing.T) {
	{
		var (
			c float64
			f Generator = func(a, b int) float64 {
				// Generates matrices with entries 1 2 3 ...wrapped along each row. For example,
				// [1 2]
				// [3 4].
				if a+b == 0 {
					c = 0
				}

				c++
				return c
			}

			rec = Multiply(Gen(2, 2, f), Gen(2, 3, f), Gen(3, 1, f))
			exp = ColumnMatrix(vtr.New(78, 170))
		)

		if !exp.Equals(rec) {
			t.Fatalf("\nexpected %s\nreceived %s", exp.String(), rec.String())
		}
	}

	{ // An example from 50 Mathematical Ideas
		var (
			A = New(
				vtr.New(7, 5, 0, 1),
				vtr.New(0, 4, 3, 7),
				vtr.New(3, 2, 0, 2),
			)

			B   = ColumnMatrix(vtr.New(3, 9, 8, 2))
			exp = ColumnMatrix(vtr.New(68, 74, 31))
			rec = A.multiply(B)
		)

		if !exp.Equals(rec) {
			t.Fatalf("\nexpected %s\nreceived %s", exp.String(), rec.String())
		}
	}

	{ // Pow
		var (
			A = New(
				vtr.New(1, 2),
				vtr.New(3, 4),
			)

			exp = New(
				vtr.New(30853, 44966),
				vtr.New(67449, 98302),
			)

			rec = Pow(A, 7)
		)

		if !exp.Equals(rec) {
			t.Fatalf("\nexpected %s\nreceived %s", exp.String(), rec.String())
		}
	}

	{ // Pow, Inverse
		var (
			A = New(
				vtr.New(1, 2),
				vtr.New(3, 4),
			)

			exp = ScalarDivide(
				New(
					vtr.New(-98302, 44966),
					vtr.New(67499, -30853),
				),
				128,
			)

			rec = Pow(A, -7)
		)

		if !exp.Equals(rec) {
			t.Fatalf("\nexpected %s\nreceived %s", exp.String(), rec.String())
		}
	}
}

func TestSolve(t *testing.T) {
	{ // A simple test with expected values determined by WolframAlpha
		//  [1 2][x0] = [5]
		//  [3 4][x1]   [6]
		exp := vtr.New(-4, 4.5)
		rec := New(vtr.New(1, 2), vtr.New(3, 4)).Solve(vtr.New(5, 6))
		if !exp.Equal(rec) {
			t.Fatalf("\nexpected %v\nreceived %v\n", exp, rec)
		}
	}

	{ // Function converting Celsius to Farenheit
		// F(C) = x0*C + x1, F(0) = 32, F(100) = 212
		//
		// [  0 1][x0] = [ 32]
		// [100 1][x1]   [212]
		exp := vtr.New(1.8, 32)
		rec := New(vtr.New(0, 1), vtr.New(100, 1)).Solve(vtr.New(32, 212))
		if !exp.Equal(rec) {
			t.Fatalf("\nexpected %v\nreceived %v\n", exp, rec)
		}
	}

	{ // Function converting Farenheit to Celsius
		// C(F) = x0*F + x1, C(32) = 0, C(212) = 100
		//
		// [ 32 1][x0] = [  0]
		// [212 1][x1]   [100]
		exp := vtr.New(5.0/9.0, -160.0/9.0)
		rec := New(vtr.New(32, 1), vtr.New(212, 1)).Solve(vtr.New(0, 100))
		if !exp.Equal(rec) {
			t.Fatalf("\nexpected %v\nreceived %v\n", exp, rec)
		}
	}
}

func fibonacci(n int) int {
	if n < 2 {
		return 1
	}

	// [ 0 1 ]
	// [ 1 1 ]
	A := Gen(2, 2, func(i, j int) float64 { return float64(i | j) })

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
		A        = Gen(2, 2, func(i, j int) float64 { return float64(i | j) })
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
		A      = Gen(2, 2, func(i, j int) float64 { return float64(i | j) })
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

type matrix []vector
type vector []float64

func zeroes(m, n int) matrix {
	A := make(matrix, 0, m)
	for i := 0; i < m; i++ {
		A = append(A, make(vector, n))
	}

	return A
}

func (A matrix) copy0() matrix {
	B := make(matrix, 0, len(A))
	for i := 0; i < len(A); i++ {
		B = append(B, append(make(vector, 0, len(A[i])), A[i]...))
	}

	return B
}

func (A matrix) copy1() matrix {
	B := make(matrix, 0, len(A))
	for i := 0; i < len(A); i++ {
		B = append(B, A[i].copy())
	}

	return B
}

func addm0(A, B matrix) matrix {
	C := make(matrix, 0, len(A))
	for i := 0; i < len(A); i++ {
		C = append(C, append(make(vector, 0, len(A[i])), B[i]...))
	}

	return C
}

func addm1(A, B matrix) matrix {
	C := A.copy0()
	C.add0(B)
	return C
}

func addm2(A, B matrix) matrix {
	C := A.copy1()
	C.add1(B)
	return C
}

func (A matrix) add0(B matrix) {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			A[i][j] += B[i][j]
		}
	}
}

func (A matrix) add1(B matrix) {
	for i := 0; i < len(A); i++ {
		A[i].add1(B[i])
	}
}

func (A matrix) add2(B matrix) {
	for i := range A {
		A[i].add1(B[i])
	}
}

func (A matrix) add3(B matrix) {
	for i, a := range A {
		a.add1(B[i])
	}
}

func gen(n int, f func(i int) float64) vector {
	v := make(vector, 0, n)
	for i := 0; i < n; i++ {
		v = append(v, f(i))
	}

	return v
}

func (v vector) copy() vector {
	return append(make(vector, 0, len(v)), v...)
}

func add0(u, v vector) vector {
	w := make(vector, 0, len(u))
	for i := 0; i < len(u); i++ {
		w = append(w, u[i]+v[i])
	}

	return w
}

func add1(u, v vector) vector {
	w := u.copy()
	w.add1(v)
	return w
}

func add2(u, v vector) vector {
	return gen(len(u), func(i int) float64 { return u[i] + v[i] })
}

func (v vector) add0(w vector) {
	for i := 0; i < len(v); i++ {
		v[i] += w[i]
	}
}

func (v vector) add1(w vector) {
	for i, wi := range w {
		v[i] += wi
	}
}

func BenchmarkInlines(b *testing.B) {
	for n := 1; n <= 1024; n <<= 1 {
		benchmarkInline0(b, zeroes(n, n), zeroes(n, n))
		benchmarkInline1(b, zeroes(n, n), zeroes(n, n))
		benchmarkInline2(b, zeroes(n, n), zeroes(n, n))
	}
}

func benchmarkInline0(b *testing.B, A, B matrix) bool {
	f := func(b0 *testing.B) {
		for i := 0; i < b0.N; i++ {
			_ = addm0(A, B)
		}
	}

	return b.Run(fmt.Sprintf("dims: %dx%d", len(A), len(A[0])), f)
}

func benchmarkInline1(b *testing.B, A, B matrix) bool {
	f := func(b0 *testing.B) {
		for i := 0; i < b0.N; i++ {
			_ = addm1(A, B)
		}
	}

	return b.Run(fmt.Sprintf("len=%d", len(A)), f)
}

func benchmarkInline2(b *testing.B, A, B matrix) bool {
	f := func(b0 *testing.B) {
		for i := 0; i < b0.N; i++ {
			_ = addm2(A, B)
		}
	}

	return b.Run(fmt.Sprintf("len=%d", len(A)), f)
}
