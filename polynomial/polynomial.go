package polynomial

import (
	"github.com/nathangreene3/math"
	"github.com/nathangreene3/math/linalg/matrix"
	"github.com/nathangreene3/math/linalg/vector"
)

// A Polynomial is an ordered set of weights f = [a0, a1, ..., an-1] such that f(x) = a0 + a1*x + ... + an-1 x^(n-1) for any real x.
type Polynomial []float64

// New returns a polynomial defined as f = [a0, a1, ..., an-1] for each
// coeficient ak.
func New(coefs ...float64) Polynomial {
	return append(make(Polynomial, 0, len(coefs)), coefs...)
}

// Add returns f+g.
func Add(f, g Polynomial) Polynomial {
	h := f.Copy()
	h.Add(g)
	return h
}

// Add g to f.
func (f Polynomial) Add(g Polynomial) {
	n := len(f)
	if n != len(g) {
		panic("degree mismatch")
	}

	for i := 0; i < n; i++ {
		f[i] += g[i]
	}
}

// Copy a polynomial.
func (f Polynomial) Copy() Polynomial {
	g := make(Polynomial, len(f))
	copy(g, f)
	return g
}

// Degree returns the highest power of f.
func (f Polynomial) Degree() int {
	return math.MaxInt(len(f)-1, 0)
}

// differentiate returns df/dx. TODO
func (f Polynomial) differentiate() Polynomial {
	if len(f) < 2 {
		return nil
	}

	g := f.Copy()[1:]
	p := 1.0
	for i := range g {
		g[i] *= p
		p++
	}

	return g
}

// Divide returns 1/a*f.
func Divide(a float64, f Polynomial) Polynomial {
	g := f.Copy()
	g.Divide(a)
	return g
}

// Divide f by a.
func (f Polynomial) Divide(a float64) {
	for i := range f {
		f[i] /= a
	}
}

// Evaluate returns f(x).
func (f Polynomial) Evaluate(x float64) float64 {
	var (
		y float64 // y = f(x)
		p = 1.0   // p = x^i, for i = 0, 1, ..., n-1
	)

	for _, a := range f {
		y += a * p
		p *= x
	}

	return y
}

// integrate returns the antiderivative of f. TODO
func (f Polynomial) integrate() Polynomial {
	g := make(Polynomial, 1, len(f)+1)
	for i := range f {
		g = append(g, f[i]/float64(i+1))
	}

	return g
}

// integrateRange returns the integral of f over [x0,x1]. TODO
func (f Polynomial) integrateRange(x0, x1 float64) float64 {
	g := f.integrate()
	return g.Evaluate(x1) - g.Evaluate(x0)
}

// Multiply returns a*f.
func Multiply(a float64, f Polynomial) Polynomial {
	g := f.Copy()
	g.Multiply(a)
	return g
}

// Multiply f by a.
func (f Polynomial) Multiply(a float64) {
	for i := range f {
		f[i] *= a
	}
}

// of returns fog. This does not evaluate fog at x. To compute (fog)(x), use
// f.Of(g, x). TODO
func of(f, g Polynomial) Polynomial {
	return nil
}

// Of returns (fog)(x) = f(g(x)). This evaluates fog at x. To get the Polynomial
// fog, use Of(f, g).
func (f Polynomial) Of(g Polynomial, x float64) float64 {
	return f.Evaluate(g.Evaluate(x))
}

// pow returns f^n. TODO
func (f Polynomial) pow(n int) Polynomial {
	switch {
	case n < 1:
		panic("indeterminant form")
	case n == 1:
		return f.Copy()
	}

	// Given G = [a0 a1 ... an-1], f^n = F^(n-1)*f, where
	// F = [ a0   a0   ...   a0 ]
	//     [ a1   a1   ...   a1 ]
	//     [ ...  ...  ...  ... ]
	//     [ an-1 an-1 ... an-1 ]
	// The coefficients of g = f^n are defined by summing the terms off the
	// secondary diagonals.
	var (
		dims = len(f)
		G    = matrix.Multiply(matrix.Pow(matrix.New(dims, dims, func(i, j int) float64 { return f[i] }), n-1), matrix.ColumnMatrix(vector.Vector(f)))
		g    = make(Polynomial, n*dims)
		k    int
	)

	for i, r := range G {
		k = i
		for _, v := range r {
			g[k] += v
			k++
		}
	}

	return g
}

// Subtract returns f-g.
func Subtract(f, g Polynomial) Polynomial {
	h := f.Copy()
	h.Subtract(g)
	return h
}

// Subtract g from f.
func (f Polynomial) Subtract(g Polynomial) {
	n := len(f)
	if n != len(g) {
		panic("degree mismatch")
	}

	for i := 0; i < n; i++ {
		f[i] -= g[i]
	}
}

// Trim removes the higher powers that have zero valued coefficients.
func (f Polynomial) Trim() Polynomial {
	// [1,2,3,0,0] --> [1,2,3]

	n := len(f)
	for ; 0 < n && f[n-1] == 0; n-- {
	}

	return f[:n].Copy()
}
