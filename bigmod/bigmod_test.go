package bigmod

import (
	"testing"
)

func TestAddSub(t *testing.T) {
	tests := []struct {
		x, y, n, exp int
		sub          bool
	}{
		{
			x:   11, //  102
			y:   5,  // + 12
			n:   3,  // ----
			exp: 16, //  121
		},
		{
			x:   11, //   102
			y:   -5, // + -12
			n:   3,  // -----
			exp: 6,  //    20
		},
		{
			x:   -11, //  -102
			y:   5,   // +  21
			n:   3,   // -----
			exp: -6,  //   -20
		},
		{
			x:   -11, //  -102
			y:   -5,  // + -12
			n:   3,   // -----
			exp: -16, //  -121
		},

		// Subtract
		{
			x:   11,
			y:   5,
			n:   3,
			exp: 6,
			sub: true,
		},
		{
			x:   11,
			y:   -5,
			n:   3,
			exp: 16,
			sub: true,
		},
		{
			x:   -11,
			y:   5,
			n:   3,
			exp: -16,
			sub: true,
		},
		{
			x:   -11,
			y:   -5,
			n:   3,
			exp: -6,
			sub: true,
		},
	}

	for _, test := range tests {
		if test.sub {
			if exp, rec := New(test.exp, test.n), New(test.x, test.n).Subtract(New(test.y, test.n)); exp.Compare(rec) != 0 {
				// TODO
				// t.Errorf("\n"+
				// 	"%d - %d (mod %d)\n"+
				// 	"expected %d as %v\n"+
				// 	"received %v\n",
				// 	test.x, test.y, test.n,
				// 	test.exp, exp,
				// 	rec,
				// )
			}
		} else {
			if exp, rec := New(test.exp, test.n), New(test.x, test.n).Add(New(test.y, test.n)); exp.Compare(rec) != 0 {
				t.Errorf("\n"+
					"%d + %d (mod %d)\n"+
					"expected %d as %v\n"+
					"received %v\n",
					test.x, test.y, test.n,
					test.exp, exp,
					rec,
				)
			}
		}
	}
}

func BenchmarkAddSubtract(b *testing.B) {

}
