package sequence

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		a    Sequence
		i, n int
		exp  float64
	}{
		{a: func(i int) float64 { return float64(i) }, i: 1, n: 5, exp: 5 * (5 + 1) / 2}, // 1+2+...+n
	}

	for _, test := range tests {
		if rec := Sum(test.a, test.i, test.n); test.exp != rec {
			t.Fatalf("\nexpected %f\nreceived %f\n", test.exp, rec)
		}
	}
}

func TestProd(t *testing.T) {
	tests := []struct {
		a    Sequence
		i, n int
		exp  float64
	}{
		{a: func(i int) float64 { return float64(i) }, i: 1, n: 5, exp: 1 * 2 * 3 * 4 * 5}, // 1*2*...*n
	}

	for _, test := range tests {
		if rec := Prod(test.a, test.i, test.n); test.exp != rec {
			t.Fatalf("\nexpected %f\nreceived %f\n", test.exp, rec)
		}
	}
}
