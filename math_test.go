package math

import "testing"

func TestGCD(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
		actual   int
	}{
		{
			a:        0,
			b:        0,
			expected: 0,
		},
		{
			a:        0,
			b:        1,
			expected: 1,
		},
		{
			a:        1,
			b:        1,
			expected: 1,
		},
		{
			a:        1,
			b:        2,
			expected: 1,
		},
		{
			a:        2,
			b:        1,
			expected: 1,
		},
		{
			a:        2,
			b:        4,
			expected: 2,
		},
		{
			a:        4,
			b:        2,
			expected: 2,
		},
		{
			a:        5,
			b:        10,
			expected: 5,
		},
	}

	for _, test := range tests {
		test.actual = GCD(test.a, test.b)
		if test.expected != test.actual {
			t.Fatalf("\nexpected: %d\nreceived: %d\n", test.expected, test.actual)
		}
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		n                int
		expected, actual int
	}{
		{
			n:        0,
			expected: 1,
		},
		{
			n:        1,
			expected: 1,
		},
		{
			n:        2,
			expected: 2,
		},
		{
			n:        3,
			expected: 6,
		},
		{
			n:        4,
			expected: 24,
		},
		{
			n:        5,
			expected: 120,
		},
	}

	for _, test := range tests {
		test.actual = Factorial(test.n)
		if test.expected != test.actual {
			t.Fatalf("\nexpected: %d\nreceived: %d\n", test.expected, test.actual)
		}
	}
}

func TestChoose(t *testing.T) {
	tests := []struct {
		n, k             int
		expected, actual int
	}{
		{
			n: 5,
			k: 0,
		},
		{
			n: 5,
			k: 1,
		},
		{
			n: 5,
			k: 2,
		},
		{
			n: 5,
			k: 3,
		},
		{
			n: 5,
			k: 4,
		},
		{
			n: 5,
			k: 5,
		},
	}

	for _, test := range tests {
		test.expected = Factorial(test.n) / (Factorial(test.k) * Factorial(test.n-test.k))
		test.actual = Choose(test.n, test.k)
		if test.expected != test.actual {
			t.Fatalf("\nexpected: %d\nreceived: %d\n", test.expected, test.actual)
		}
	}
}
