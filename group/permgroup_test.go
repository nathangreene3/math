package group

import "testing"

func TestIsPermutation(t *testing.T) {
	tests := []struct {
		a        Permutation
		exp, rec bool
	}{
		{
			a:   New(5),
			exp: true,
		},
		{
			a:   New(5),
			exp: true,
		},
	}

	for _, test := range tests {
		test.rec = test.a.isPermutation()
		if test.exp != test.rec {
			t.Fatalf("expected %v\nreceived %v\n", test.exp, test.rec)
		}
	}
}
