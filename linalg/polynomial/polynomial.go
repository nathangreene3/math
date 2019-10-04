package polynomial

import (
	"github.com/nathangreene3/math"
	"github.com/nathangreene3/math/linalg/matrix"
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
func (f *Polynomial) Add(g Polynomial) {
	m, n := len(*f), len(g)
	switch {
	case m < n:
		*f = append(*f, make([]float64, n-m)...)
	case n < m:
		g = append(g, make([]float64, m-n)...)
	}

	for i := 0; i < n; i++ {
		(*f)[i] += g[i]
	}
}

// Compare two polynomials.
func (f *Polynomial) Compare(g Polynomial) int {
	trimmedF, trimmedG := f.Trim(), g.Trim()
	m, n := len(trimmedF), len(trimmedG)
	switch {
	case m < n:
		return 1
	case n < m:
		return -1
	}

	for i := 0; i < n; i++ {
		switch {
		case trimmedF[i] < trimmedG[i]:
			return 1
		case trimmedG[i] < trimmedF[i]:
			return -1
		}
	}

	return 0
}

// Copy a polynomial.
func (f *Polynomial) Copy() Polynomial {
	g := make(Polynomial, len(*f))
	copy(g, *f)
	return g
}

// Degree returns the highest power of f.
func (f *Polynomial) Degree() int {
	return math.MaxInt(len(*f)-1, 0)
}

// differentiate TODO
func differentiate(f Polynomial) Polynomial {
	return nil
}

// differentiate returns df/dx. TODO
func (f *Polynomial) differentiate() Polynomial {
	if len(*f) < 2 {
		return nil
	}

	g := New((*f)[1:]...)
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
func (f *Polynomial) Divide(a float64) {
	for i := range *f {
		(*f)[i] /= a
	}
}

// Equal compares two polynomials.
func (f *Polynomial) Equal(g Polynomial) bool {
	return f.Compare(g) == 0
}

// Evaluate returns f(x).
func (f *Polynomial) Evaluate(x float64) float64 {
	var (
		y float64 // y = f(x)
		p = 1.0   // p = x^i, for i = 0, 1, ..., n-1
	)

	for _, a := range *f {
		y += a * p
		p *= x
	}

	return y
}

// integrate returns the antiderivative of f. TODO
func (f *Polynomial) integrate() Polynomial {
	g := make(Polynomial, 1, len(*f)+1)
	for i := range *f {
		g = append(g, (*f)[i]/float64(i+1))
	}

	return g
}

// integrateRange returns the integral of f over [x0,x1]. TODO
func (f *Polynomial) integrateRange(x0, x1 float64) float64 {
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
func (f *Polynomial) Multiply(a float64) {
	for i := range *f {
		(*f)[i] *= a
	}
}

// of returns fog. This does not evaluate fog at x. To compute (fog)(x), use
// f.Of(g, x). TODO
func of(f, g Polynomial) Polynomial {
	return nil
}

// Of returns (fog)(x) = f(g(x)). This evaluates fog at x. To get the Polynomial
// fog, use Of(f, g).
func (f *Polynomial) Of(g Polynomial, x float64) float64 {
	return f.Evaluate(g.Evaluate(x))
}

// pow returns f^n. TODO
func (f *Polynomial) pow(n int) Polynomial {
	switch {
	case n < 1:
		panic("indeterminant form")
	case n == 1:
		return f.Copy()
	}

	// Given f = [a0 a1 ... an-1], f^n = F^(n-1)
	// F = [ a0   a0   ...   a0 ]
	//     [ a1   a1   ...   a1 ]
	//     [ ...  ...  ...  ... ]
	//     [ an-1 an-1 ... an-1 ]
	var (
		dims = len(*f)
		F    = matrix.Pow(matrix.New(dims, dims, func(i, j int) float64 { return (*f)[i] }), dims-1)
	)

	for i, v := range *f {
		F.MultiplyColumn(i, v)
	}

	// Each dimension g(k) = Sum F(i,j) for all k = i+j.
	g := make(Polynomial, n*dims)
	for i, r := range F {
		for j, v := range r {
			g[i+j] += v
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
func (f *Polynomial) Subtract(g Polynomial) {
	m, n := len(*f), len(g)
	switch {
	case m < n:
		*f = append(*f, make([]float64, n-m)...)
	case n < m:
		g = append(g, make([]float64, m-n)...)
	}

	for i := 0; i < n; i++ {
		(*f)[i] -= g[i]
	}
}

// Trim removes the higher powers that have zero valued coefficients. Only removes from the right.
func (f *Polynomial) Trim() Polynomial {
	n := len(*f)
	if n == 0 {
		return *f
	}

	for ; 0 < n && (*f)[n-1] == 0; n-- {
	}

	return (*f)[:n]
}
