package math

import (
	"fmt"
	gomath "math"
	"strconv"
	"strings"
	"testing"
)

func ackermann(x, y int) int {
	switch {
	case x == 0:
		return y + 1
	case y == 0:
		return ackermann(x-1, 1)
	default:
		return ackermann(x-1, ackermann(x, y-1))
	}
}

func TestAckermann(t *testing.T) {
	t.Error(ackermann(2, 2))
}

// equalInts determines if two slices of integers are equal.
func equalInts(a, b []int) bool {
	n := len(a)
	if n != len(b) {
		return false
	}

	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// equalIntsToMap determines if a slice of integers and a map of integer keys are equal.
func equalIntsToMap(s []int, m map[int]struct{}) bool {
	n := len(s)
	if n != len(m) {
		return false
	}

	for i := 0; i < n; i++ {
		if _, ok := m[s[i]]; !ok {
			return false
		}
	}

	return true
}

func TestBase(t *testing.T) {
	tests := []struct {
		n0, n1, b int
		exp, rec  []int
	}{
		{
			n0:  15,
			b:   2,
			exp: []int{1, 1, 1, 1},
		},
		{
			n0:  15,
			b:   3,
			exp: []int{0, 2, 1},
		},
		{
			n0:  15,
			b:   10,
			exp: []int{5, 1},
		},
		{
			n0:  42,
			b:   2,
			exp: []int{0, 1, 0, 1, 0, 1},
		},
	}

	for _, test := range tests {
		test.rec = Base(test.n0, test.b)
		if !equalInts(test.exp, test.rec) {
			t.Fatalf("expected %v\nreceived %v\n", test.exp, test.rec)
		}

		test.n1 = Base10(test.rec, test.b)
		if test.n0 != test.n1 {
			t.Fatalf("\nexpected %d\nreceived %d\n", test.n0, test.n1)
		}
	}
}

func TestBasePows(t *testing.T) {
	tests := []struct {
		n, b     int
		exp, rec []int
	}{
		{
			n:   1,
			b:   2,
			exp: []int{1},
		},
		{
			n:   2,
			b:   2,
			exp: []int{0, 2},
		},
		{
			n:   3,
			b:   2,
			exp: []int{1, 2},
		},
		{
			n:   15,
			b:   2,
			exp: []int{1, 2, 4, 8},
		},
		{
			n:   15,
			b:   3,
			exp: []int{0 * 1, 2 * 3, 1 * 9},
		},
		{
			n:   15,
			b:   10,
			exp: []int{5, 10},
		},
		{
			n:   42,
			b:   2,
			exp: []int{0, 2, 0, 8, 0, 32},
		},
	}

	for _, test := range tests {
		test.rec = BasePows(test.n, test.b)
		if !equalInts(test.exp, test.rec) {
			t.Fatalf("expected %v\nreceived %v\n", test.exp, test.rec)
		}
	}
}

func TestChoose(t *testing.T) {
	tests := []struct {
		n, k             int
		expected, actual int
	}{
		{n: 5, k: 0},
		{n: 5, k: 1},
		{n: 5, k: 2},
		{n: 5, k: 3},
		{n: 5, k: 4},
		{n: 5, k: 5},
	}

	for _, test := range tests {
		test.expected = Fact(test.n) / (Fact(test.k) * Fact(test.n-test.k))
		test.actual = Choose(test.n, test.k)
		if test.expected != test.actual {
			t.Fatalf("\nexpected: %d\nreceived: %d\n", test.expected, test.actual)
		}
	}
}

func TestFactor(t *testing.T) {
	tests := []struct {
		n        int
		expected map[int]int
	}{
		{n: 1, expected: map[int]int{}},
		{n: 2, expected: map[int]int{2: 1}},
		{n: 3, expected: map[int]int{3: 1}},
		{n: 4, expected: map[int]int{2: 2}},
		{n: 5, expected: map[int]int{5: 1}},
		{n: 6, expected: map[int]int{2: 1, 3: 1}},
		{n: 7, expected: map[int]int{7: 1}},
		{n: 8, expected: map[int]int{2: 3}},
		{n: 9, expected: map[int]int{3: 2}},
		{n: 10, expected: map[int]int{2: 1, 5: 1}},
		{n: 11, expected: map[int]int{11: 1}},
		{n: 12, expected: map[int]int{2: 2, 3: 1}},
		{n: 13, expected: map[int]int{13: 1}},
		{n: 14, expected: map[int]int{2: 1, 7: 1}},
		{n: 15, expected: map[int]int{3: 1, 5: 1}},
		{n: 16, expected: map[int]int{2: 4}},
		{n: 17, expected: map[int]int{17: 1}},
		{n: 18, expected: map[int]int{2: 1, 3: 2}},
		{n: 19, expected: map[int]int{19: 1}},
		{n: 20, expected: map[int]int{2: 2, 5: 1}},
		{n: 21, expected: map[int]int{3: 1, 7: 1}},
		{n: 22, expected: map[int]int{2: 1, 11: 1}},
		{n: 23, expected: map[int]int{23: 1}},
		{n: 24, expected: map[int]int{2: 3, 3: 1}},
		{n: 25, expected: map[int]int{5: 2}},
		{n: 49, expected: map[int]int{7: 2}},

		// Highly composite numbers
		{n: 24, expected: map[int]int{2: 3, 3: 1}},
		{n: 36, expected: map[int]int{2: 2, 3: 2}},
		{n: 48, expected: map[int]int{2: 4, 3: 1}},
		{n: 60, expected: map[int]int{2: 2, 3: 1, 5: 1}},
		{n: 120, expected: map[int]int{2: 3, 3: 1, 5: 1}},

		// Largest prime for each order of 10
		{n: 97, expected: map[int]int{97: 1}},
		{n: 997, expected: map[int]int{997: 1}},
		{n: 7919, expected: map[int]int{7919: 1}},

		// Mersenne primes
		{n: 3, expected: map[int]int{3: 1}},
		{n: 7, expected: map[int]int{7: 1}},
		{n: 31, expected: map[int]int{31: 1}},
		{n: 127, expected: map[int]int{127: 1}},
		{n: 8191, expected: map[int]int{8191: 1}},
		{n: 131071, expected: map[int]int{131071: 1}},
		{n: 524287, expected: map[int]int{524287: 1}},
		{n: 2147483647, expected: map[int]int{2147483647: 1}},
		// {n: 2305843009213693951, expected: map[int]int{2305843009213693951: 1}}, // Largest possible without overflowing
	}

	for _, test := range tests {
		received := Factor(test.n)
		for k, expected := range test.expected {
			if r, ok := received[k]; !ok || expected != r {
				t.Fatalf("expected %d\nreceived %d", test.expected, received)
			}
		}

		// received = factor2(test.n)
		// for k, expected := range test.expected {
		// 	if r, ok := received[k]; !ok || expected != r {
		// 		t.Fatalf("expected %d\nreceived %d", test.expected, received)
		// 	}
		// }

		received = factor3(test.n)
		for k, expected := range test.expected {
			if r, ok := received[k]; !ok || expected != r {
				t.Fatalf("expected %d\nreceived %d", test.expected, received)
			}
		}

		// received = factor4(test.n)
		// for k, expected := range test.expected {
		// 	if r, ok := received[k]; !ok || expected != r {
		// 		t.Fatalf("expected %d\nreceived %d", test.expected, received)
		// 	}
		// }
	}
}

func TestFactorCount(t *testing.T) {
	var (
		b   strings.Builder
		max int
	)

	for n := 2; n <= 4096; n++ {
		if f := Factor(n); max < len(f) {
			if 0 < b.Len() {
				b.WriteByte('\n')
			}

			b.WriteString(fmt.Sprintf("%d: %v", n, f))
			max = len(f)
		}
	}

	t.Fatal("\n" + b.String())
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		n, exp int
	}{
		{n: 0, exp: 1},
		{n: 1, exp: 1},
		{n: 2, exp: 2},
		{n: 3, exp: 6},
		{n: 4, exp: 24},
		{n: 5, exp: 120},
		{n: 6, exp: 720},
	}

	for _, test := range tests {
		if rec := Fact(test.n); test.exp != rec {
			t.Fatalf("\n   given %d\nexpected %d\nreceived %d\n", test.n, test.exp, rec)
		}

		if rec := factRec(test.n); test.exp != rec {
			t.Fatalf("\n   given %d\nexpected %d\nreceived %d\n", test.n, test.exp, rec)
		}

		if rec := factTail(test.n); test.exp != rec {
			t.Fatalf("\n   given %d\nexpected %d\nreceived %d\n", test.n, test.exp, rec)
		}
	}
}

func TestGCD(t *testing.T) {
	tests := []struct {
		a   int
		b   int
		exp int
	}{
		{a: 1, b: 1, exp: 1},
		{a: 1, b: 2, exp: 1},
		{a: 2, b: 1, exp: 1},
		{a: 2, b: 4, exp: 2},
		{a: 4, b: 2, exp: 2},
		{a: 5, b: 10, exp: 5},
		{a: 10, b: 5, exp: 5},
		{a: 2 * 3 * 5 * 7, b: 2 * 3 * 5, exp: 5},
	}

	for _, test := range tests {
		if rec := GCD(test.a, test.b); test.exp != rec {
			t.Fatalf("\nexpected: %d\nreceived: %d\n", test.exp, rec)
		}
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		a        int
		exp, rec bool
	}{

		{a: 0, exp: false},
		{a: 1, exp: false},
		{a: 2, exp: true},
		{a: 3, exp: true},
		{a: 4, exp: false},
		{a: 5, exp: true},
		{a: 6, exp: false},
		{a: 7, exp: true},
		{a: 8, exp: false},
		{a: 9, exp: false},
		{a: 10, exp: false},
		{a: 11, exp: true},
		{a: 12, exp: false},
		{a: 13, exp: true},
		{a: 14, exp: false},
		{a: 15, exp: false},
		{a: 16, exp: false},
		{a: 17, exp: true},
		{a: 18, exp: false},
		{a: 19, exp: true},
		{a: 20, exp: false},

		// Mersenne Prime: Mp := 2^p-1
		// Not all Mersenne numbers are prime.
		{a: PowInt(2, 2) - 1, exp: true},
		{a: PowInt(2, 3) - 1, exp: true},
		{a: PowInt(2, 5) - 1, exp: true},
		{a: PowInt(2, 7) - 1, exp: true},
		{a: PowInt(2, 13) - 1, exp: true},
		{a: PowInt(2, 17) - 1, exp: true},
		{a: PowInt(2, 19) - 1, exp: true},
	}

	for _, test := range tests {
		test.rec = IsPrime(test.a)
		if test.exp != test.rec {
			t.Fatalf("\na = %d\nexpected: %t\nreceived: %t\n", test.a, test.exp, test.rec)
		}
	}
}

func TestPrimes(t *testing.T) {
	var (
		n          int              = 1024
		p0, p1, p3 []int            = primes(n), primes1(n), primes3(n)
		p2         map[int]struct{} = primes2(n)
	)

	if !equalInts(p0, p1) {
		t.Fatalf("\nexpected equality")
	}

	if !equalIntsToMap(p0, p2) {
		t.Fatalf("\nexpected equality\n")
	}

	if !equalInts(p0, p3) {
		t.Fatalf("\nexpected equality")
	}
}

func TestPowInt(t *testing.T) {
	tests := []struct {
		a, p     int
		exp, rec int
	}{
		{
			a:   0,
			p:   1,
			exp: 0,
		},
		{
			a:   1,
			p:   0,
			exp: 1,
		},
		{
			a:   2,
			p:   0,
			exp: 1,
		},
		{
			a:   2,
			p:   1,
			exp: 2,
		},
		{
			a:   2,
			p:   2,
			exp: 4,
		},
		{
			a:   2,
			p:   3,
			exp: 8,
		},
		{
			a:   2,
			p:   4,
			exp: 16,
		},
		{
			a:   2,
			p:   5,
			exp: 32,
		},
		{
			a:   2,
			p:   6,
			exp: 64,
		},
		{
			a:   2,
			p:   7,
			exp: 128,
		},
		{
			a:   3,
			p:   0,
			exp: 1,
		},
		{
			a:   3,
			p:   1,
			exp: 3,
		},
		{
			a:   3,
			p:   2,
			exp: 9,
		},
		{
			a:   3,
			p:   3,
			exp: 27,
		},
		{
			a:   3,
			p:   4,
			exp: 81,
		},
	}

	for _, test := range tests {
		test.rec = PowInt(test.a, test.p)
		if test.exp != test.rec {
			t.Fatalf("expected %v\nreceived %v\n", test.exp, test.rec)
		}
	}
}

func TestTotient(t *testing.T) {
	tests := []struct {
		n   int
		exp int
		rec int
	}{
		{
			n:   1,
			exp: 1,
		},
		{
			n:   2,
			exp: 1,
		},
		{
			n:   3,
			exp: 2,
		},
		{
			n:   4,
			exp: 2,
		},
		{
			n:   5,
			exp: 4,
		},
		{
			n:   6,
			exp: 2,
		},
		{
			n:   7,
			exp: 6,
		},
		{
			n:   8,
			exp: 4,
		},
		{
			n:   9,
			exp: 6,
		},
		{
			n:   10,
			exp: 4,
		},
	}

	for _, test := range tests {
		test.rec = Totient(test.n)
		if test.exp != test.rec {
			t.Fatalf("\nphi(%d)\nexpected %d\nreceived %d\n", test.n, test.exp, test.rec)
		}
	}
}

func TestTernaryAndTwitch(t *testing.T) {
	t.Fatal(Twitch(2, -1, Tuple{0, 0}, Tuple{1, 1}, Tuple{2, 0}))
}

// TODO: Finish testing for 32 and 64 bits.
func TestNextPowOfTwo(t *testing.T) {
	tests := []struct {
		n, exp int
	}{
		{n: -8, exp: -8},
		{n: -7, exp: -8},
		{n: -6, exp: -8},
		{n: -5, exp: -8},
		{n: -4, exp: -4},
		{n: -3, exp: -4},
		{n: -2, exp: -2},
		{n: -1, exp: -1},
		{n: 0, exp: 1},
		{n: 1, exp: 1},
		{n: 2, exp: 2},
		{n: 3, exp: 4},
		{n: 4, exp: 4},
		{n: 5, exp: 8},
		{n: 6, exp: 8},
		{n: 7, exp: 8},
		{n: 8, exp: 8},
	}

	for _, test := range tests {
		rec := NextPowOfTwo(test.n)
		if test.exp != rec {
			t.Fatalf("\n   NextPowOfTwo(%d)\nexpected %d\nreceived %d\n", test.n, test.exp, rec)
		}
	}
}

func BenchmarkFactor(b *testing.B) {
	tests := []int{
		// Mersenne primes should be hard to factor...
		3,
		7,
		31,
		127,
		8191,
		131071,
		524287,
		// 2147483647,
		// 2305843009213693951,
	}

	for i := 0; i < len(tests); i++ {
		benchmarkFactor(b, tests[i])
		// benchmarkFactor2(b, tests[i])
		// benchmarkFactor3(b, tests[i])
		benchmarkFactor4(b, tests[i])
	}
}

func benchmarkFactor(b *testing.B, n int) {
	b.Run(
		"Factoring "+strconv.Itoa(n),
		func(b1 *testing.B) {
			for i := 0; i < b1.N; i++ {
				_ = Factor(n)
			}
		},
	)
}

func benchmarkFactor2(b *testing.B, n int) {
	b.Run(
		"Factoring "+strconv.Itoa(n),
		func(b1 *testing.B) {
			for i := 0; i < b1.N; i++ {
				_ = factor2(n)
			}
		},
	)
}

func benchmarkFactor3(b *testing.B, n int) {
	b.Run(
		"Factoring "+strconv.Itoa(n),
		func(b1 *testing.B) {
			for i := 0; i < b1.N; i++ {
				_ = factor3(n)
			}
		},
	)
}

func benchmarkFactor4(b *testing.B, n int) {
	b.Run(
		"Factoring "+strconv.Itoa(n),
		func(b1 *testing.B) {
			for i := 0; i < b1.N; i++ {
				_ = factor4(n)
			}
		},
	)
}

func BenchmarkFactorial(b *testing.B) {
	// Linear increase in n
	var _min, _max, step int = 1, 8, 1
	for n := _min; n <= _max; n += step {
		subbenchmarkFactorial(b, n)
	}

	for n := _min; n <= _max; n += step {
		benchmarkFactorialRec(b, n)
	}

	for n := _min; n <= _max; n += step {
		benchmarkFactorialTail(b, n)
	}
}

func subbenchmarkFactorial(b *testing.B, n int) bool {
	f := func(b1 *testing.B) {
		for i := 0; i < b1.N; i++ {
			_ = Fact(n)
		}
	}

	return b.Run(fmt.Sprintf("Factorial(%d)", n), f)
}

func benchmarkFactorialRec(b *testing.B, n int) bool {
	f := func(b1 *testing.B) {
		for i := 0; i < b1.N; i++ {
			_ = factRec(n)
		}
	}

	return b.Run(fmt.Sprintf("FactorialRec(%d)", n), f)
}

func benchmarkFactorialTail(b *testing.B, n int) bool {
	f := func(b1 *testing.B) {
		for i := 0; i < b1.N; i++ {
			_ = factTail(n)
		}
	}

	return b.Run(fmt.Sprintf("FactorialTail(%d)", n), f)
}

func BenchmarkFibonacci(b *testing.B) {
	tests := []int{0, 1, 2, 3, 4, 5}
	for i := 0; i < len(tests); i++ {
		benchmarkFibonacci(b, tests[i])
		benchmarkFibonacciTail(b, tests[i])
	}
}

func benchmarkFibonacci(b *testing.B, n int) {
	b.Run(
		"Fibonacci of "+strconv.Itoa(n),
		func(b1 *testing.B) {
			for i := 0; i < b1.N; i++ {
				Fibonacci(n)
			}
		},
	)
}

func benchmarkFibonacciTail(b *testing.B, n int) {
	b.Run(
		"FibonacciTail of "+strconv.Itoa(n),
		func(b1 *testing.B) {
			for i := 0; i < b1.N; i++ {
				fibonacciTail(n)
			}
		},
	)
}

func BenchmarkPowInt(b *testing.B) {
	{
		// Linear increase in a and p
		var _min, _max, step int = 1, 8, 1
		for a := _min; a <= _max; a += step {
			for p := _min; p <= _max; p += step {
				subbenchmarkPowInt(b, a, p)
			}
		}

		for a := _min; a <= _max; a += step {
			for p := _min; p <= _max; p += step {
				subbenchmarkGoMathPow(b, a, p)
			}
		}
	}

	{
		// Exponential increase in a and p
		var _min, _max, step int = 1, 256, 1
		for a := _min; a <= _max; a <<= step {
			for p := _min; p <= _max; p <<= step {
				subbenchmarkPowInt(b, a, p)
			}
		}

		for a := _min; a <= _max; a <<= step {
			for p := _min; p <= _max; p <<= step {
				subbenchmarkGoMathPow(b, a, p)
			}
		}
	}
}

func subbenchmarkPowInt(b *testing.B, a, p int) bool {
	f := func(b0 *testing.B) {
		for i := 0; i < b0.N; i++ {
			_ = PowInt(a, p)
		}
	}

	return b.Run(fmt.Sprintf("PowInt(%d,%d)", a, p), f)
}

func subbenchmarkGoMathPow(b *testing.B, a, p int) bool {
	f := func(b0 *testing.B) {
		for i := 0; i < b0.N; i++ {
			_ = gomath.Pow(float64(a), float64(p)) // Converting to floats is intended to be tested as well
		}
	}

	return b.Run(fmt.Sprintf("GoPow(%d,%d)", a, p), f)
}

func BenchmarkPrimes(b *testing.B) {
	{
		// Linear increase in n
		i0, max, step := 256, 2048, 256
		for i := i0; i <= max; i += step {
			subbenchPrimes(b, i)
		}

		for i := i0; i <= max; i += step {
			subbenchPrimes1(b, i)
		}

		for i := i0; i <= max; i += step {
			subbenchPrimes2(b, i)
		}

		for i := i0; i <= max; i += step {
			subbenchPrimes3(b, i)
		}
	}

	{
		// Exponential increase in n
		i0, max, step := 1, 256, 1
		for i := i0; i <= max; i <<= step {
			subbenchPrimes(b, i)
		}

		for i := i0; i <= max; i <<= step {
			subbenchPrimes1(b, i)
		}

		for i := i0; i <= max; i <<= step {
			subbenchPrimes2(b, i)
		}

		for i := i0; i <= max; i <<= step {
			subbenchPrimes3(b, i)
		}
	}
}

func subbenchPrimes(b *testing.B, n int) bool {
	f := func(b0 *testing.B) {
		for i := 0; i < b0.N; i++ {
			_ = primes(n)
		}
	}

	return b.Run(fmt.Sprintf("primes(%d)", n), f)
}

func subbenchPrimes1(b *testing.B, n int) bool {
	f := func(b0 *testing.B) {
		for i := 0; i < b0.N; i++ {
			_ = primes1(n)
		}
	}

	return b.Run(fmt.Sprintf("primes1(%d)", n), f)
}

func subbenchPrimes2(b *testing.B, n int) bool {
	f := func(b0 *testing.B) {
		for i := 0; i < b0.N; i++ {
			_ = primes2(n)
		}
	}

	return b.Run(fmt.Sprintf("primes2(%d)", n), f)
}

func subbenchPrimes3(b *testing.B, n int) bool {
	f := func(b0 *testing.B) {
		for i := 0; i < b0.N; i++ {
			_ = primes3(n)
		}
	}

	return b.Run(fmt.Sprintf("primes3(%d)", n), f)
}

func BenchmarkSum(b *testing.B) {
	s := make([]float64, 256)
	for i := 0; i < b.N; i++ {
		_ = Sum(s...)
	}
}

func BenchmarkAbs(b *testing.B) {
	for n := -4.0; n <= 4; n++ {
		benchmarkGoMathAbs(b, n)
	}

	for n := -4.0; n <= 4; n++ {
		benchmarkAbs(b, n)
	}

	benchmarkGoMathAbs(b, -gomath.MaxFloat64)
	benchmarkGoMathAbs(b, gomath.MaxFloat64)
	benchmarkAbs(b, -gomath.MaxFloat64)
	benchmarkAbs(b, gomath.MaxFloat64)
}

func benchmarkGoMathAbs(b *testing.B, x float64) bool {
	f := func(b0 *testing.B) {
		for i := 0; i < b0.N; i++ {
			_ = gomath.Abs(x)
		}
	}

	return b.Run(fmt.Sprintf("|%f|", x), f)
}

// Abs ...
func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}

	return x
}

func benchmarkAbs(b *testing.B, x float64) bool {
	f := func(b0 *testing.B) {
		for i := 0; i < b0.N; i++ {
			_ = Abs(x)
		}
	}

	return b.Run(fmt.Sprintf("|%f|", x), f)
}
