package math

import (
	"fmt"
	"testing"
)

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
	}

	for _, test := range tests {
		test.rec = IsPrime(test.a)
		if test.exp != test.rec {
			t.Fatalf("\nexpected: %t\nreceived: %t\n", test.exp, test.rec)
		}
	}
}

func TestToBase(t *testing.T) {
	tests := []struct {
		n, b     int
		exp, rec []int
	}{
		{
			n:   15,
			b:   2,
			exp: []int{1, 1, 1, 1},
		},
		{
			n:   15,
			b:   3,
			exp: []int{0, 2, 1},
		},
		{
			n:   15,
			b:   10,
			exp: []int{5, 1},
		},
		{
			n:   42,
			b:   2,
			exp: []int{0, 1, 0, 1, 0, 1},
		},
	}

	for _, test := range tests {
		test.rec = Base(test.n, test.b)
		if !equalInts(test.exp, test.rec) {
			t.Fatalf("expected %v\nreceived %v\n", test.exp, test.rec)
		}

		fmt.Printf("len = %d, cap = %d\n", len(test.rec), cap(test.rec))
	}

	t.Fatalf("")
}

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
