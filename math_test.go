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
