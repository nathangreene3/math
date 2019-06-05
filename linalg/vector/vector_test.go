package vector

import "testing"

func TestIsMultipleOf(t *testing.T) {
	tests := []struct {
		v        Vector
		w        Vector
		expected bool
		actual   bool
	}{
		{
			v:        Vector{1, 2},
			w:        Vector{3, 1},
			expected: false,
		},
		{
			v:        Vector{1, 2},
			w:        Vector{3, 6},
			expected: true,
		},
		{
			v:        Vector{0, 1.1, 2},
			w:        Vector{0, 3, 6},
			expected: false,
		},
	}

	for _, test := range tests {
		test.actual = IsMultipleOf(test.v, test.w)
		if test.expected != test.actual {
			t.Fatalf("\nexpected: %t\nreceived: %t", test.expected, test.actual)
		}
	}
}
