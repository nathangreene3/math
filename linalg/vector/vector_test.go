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
		test.actual = test.v.IsMultipleOf(test.w)
		if test.expected != test.actual {
			t.Fatalf("\nexpected: %t\nreceived: %t", test.expected, test.actual)
		}
	}
}

func TestApprox(t *testing.T) {
	var (
		v = []Vector{
			Vector{0},
			Vector{1},
		}
		w = []Vector{
			Vector{0.00000000000009383146683006760733897488065551829232273373104789015997084788978099822998046875},
			Vector{0.99999999999990596410981424924102611839771270751953125},
		}
		prec = 0.010
	)

	if !v[0].Approx(w[0], prec) {
		t.Fatalf("v[0] is not approximately w[0]:\nv = %s\nw[0] = %s", v[0], w[0])
	}

	if !v[1].Approx(w[1], prec) {
		t.Fatalf("v[1] is not approximately w[1]:\nv = %s\nw[1] = %s", v[1], w[1])
	}
}
