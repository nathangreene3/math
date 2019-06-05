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
