package math

import (
	gomath "math"
	"testing"
)

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
		test.expected = Factorial(test.n) / (Factorial(test.k) * Factorial(test.n-test.k))
		test.actual = Choose(test.n, test.k)
		if test.expected != test.actual {
			t.Fatalf("\nexpected: %d\nreceived: %d\n", test.expected, test.actual)
		}
	}
}

func TestFactor(t *testing.T) {
	tests := []struct {
		n                  int
		expected, received map[int]int
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

		// Highly composite numbers
		{n: 24, expected: map[int]int{2: 3, 3: 1}},
		{n: 36, expected: map[int]int{2: 2, 3: 2}},
		{n: 48, expected: map[int]int{2: 4, 3: 1}},
		{n: 60, expected: map[int]int{2: 2, 3: 1, 5: 1}},
		{n: 120, expected: map[int]int{2: 3, 3: 1, 5: 1}},

		// Largest prime for each order of 10
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
		// {n: 2305843009213693951, expected: map[int]int{2305843009213693951: 1}}, // Largest possible without overflow2305843009213693951:1ing
	}

	for _, test := range tests {
		test.received = Factor(test.n)
		for k, expected := range test.expected {
			received, ok := test.received[k]
			if !ok || expected != received {
				t.Fatalf("expected %d\nreceived %d", test.expected, test.received)
			}
		}
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		n, expected, actual int
	}{
		{n: 0, expected: 1},
		{n: 1, expected: 1},
		{n: 2, expected: 2},
		{n: 3, expected: 6},
		{n: 4, expected: 24},
		{n: 5, expected: 120},
	}

	for _, test := range tests {
		test.actual = Factorial(test.n)
		if test.expected != test.actual {
			t.Fatalf("\nexpected: %d\nreceived: %d\n", test.expected, test.actual)
		}
	}
}

func TestGCD(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
		actual   int
	}{
		{a: 0, b: 0, expected: 0},
		{a: 0, b: 1, expected: 1},
		{a: 1, b: 1, expected: 1},
		{a: 1, b: 2, expected: 1},
		{a: 2, b: 1, expected: 1},
		{a: 2, b: 4, expected: 2},
		{a: 4, b: 2, expected: 2},
		{a: 5, b: 10, expected: 5},
		{a: 10, b: 5, expected: 5},
	}

	for _, test := range tests {
		test.actual = GCD(test.a, test.b)
		if test.expected != test.actual {
			t.Fatalf("\nexpected: %d\nreceived: %d\n", test.expected, test.actual)
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
			t.Fatalf("\n   given %d\nexpected %d\nreceived %d\n", test.n, test.exp, rec)
		}
	}
}

func BenchmarkPowInt(b *testing.B) {
	a, p := 10, 10
	for i := 0; i < b.N; i++ {
		_ = PowInt(a, p)
	}
}

func BenchmarkGomathPow(b *testing.B) {
	a, p := 10, 10
	for i := 0; i < b.N; i++ {
		_ = int(gomath.Pow(float64(a), float64(p)))
	}
}

var s = make([]int, 256)

func BenchmarkFor1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s); j++ {
			_ = j
		}
	}
}

func BenchmarkFor2(b *testing.B) {
	n := len(s)
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			_ = j
		}
	}
}

func BenchmarkFor3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := range s {
			_ = j
		}
	}
}

func BenchmarkSum(b *testing.B) {
	s := make([]float64, 256)
	for i := 0; i < b.N; i++ {
		_ = Sum(s...)
	}
}
