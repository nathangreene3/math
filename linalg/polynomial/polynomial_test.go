package polynomial

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		f, g, exp, rec Polynomial
	}{
		{
			f:   New(),
			g:   New(),
			exp: New(),
		},
		{
			f:   New(1),
			g:   New(1),
			exp: New(2),
		},
		{
			f:   New(1, 2, 3),
			g:   New(4, 5, 6),
			exp: New(5, 7, 9),
		},
		{
			f:   New(1, 2, 3),
			g:   New(4, 5, 6, 7),
			exp: New(5, 7, 9, 7),
		},
	}

	for _, test := range tests {
		test.rec = Add(test.f, test.g)
		if !test.exp.Equal(test.rec) {
			t.Fatalf("\nexpected %v\nreceived %v\n", test.exp, test.rec)
		}
	}
}

func TestDifferentiate(t *testing.T) {
	tests := []struct {
		f, exp, rec Polynomial
	}{
		{
			f:   New(),
			exp: New(),
		},
		{
			f:   New(1),
			exp: New(),
		},
		{
			f:   New(1, 2),
			exp: New(2),
		},
		{
			f:   New(1, 2, 3),
			exp: New(2, 6),
		},
		{
			f:   New(1, 2, 3, 4),
			exp: New(2, 6, 12),
		},
	}

	for _, test := range tests {
		test.rec = test.f.differentiate()
		if !test.exp.Equal(test.rec) {
			t.Fatalf("\nexpected %v\nreceived %v\n", test.exp, test.rec)
		}
	}
}

func TestTrim(t *testing.T) {
	tests := []struct {
		f, exp, rec Polynomial
	}{
		{
			f:   New(),
			exp: New(),
		},
		{
			f:   New(1, 2, 3),
			exp: New(1, 2, 3),
		},
		{
			f:   New(1, 2, 3, 0, 0),
			exp: New(1, 2, 3),
		},
	}

	for _, test := range tests {
		test.rec = test.f.Trim()
		if !test.exp.Equal(test.rec) {
			t.Fatalf("\nexpected %v\nreceived %v\n", test.exp, test.rec)
		}
	}
}
