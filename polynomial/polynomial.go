package polynomial

import "github.com/nathangreene3/math"

// A Polynomial is an ordered set of weights f = [a0, a1, ..., an-1] such that f(x) = a0 + a1*x + ... + an-1 x^(n-1) for any real x.
type Polynomial []float64

// New returns a polynomial defined as f = [a0, a1, ..., an-1] for each
// coeficient ak.
func New(coefs ...float64) Polynomial {
	f := make(Polynomial, 0, len(coefs))
	for _, a := range coefs {
		f = append(f, a)
	}

	return f
}

// add returns f+g.
func add(f, g Polynomial) Polynomial {
	h := f.copy()
	h.add(g)
	return h
}

// add g to f.
func (f Polynomial) add(g Polynomial) {
	n := len(f)
	if n != len(g) {
		panic("degree mismatch")
	}

	for i := 0; i < n; i++ {
		f[i] += g[i]
	}
}

// copy a polynomial.
func (f Polynomial) copy() Polynomial {
	return New(f...)
}

// degree returns the highest power of f.
func (f Polynomial) degree() int {
	return math.MaxInt(len(f)-1, 0)
}

// differentiate returns df/dx.
func (f Polynomial) differentiate() Polynomial {
	if len(f) < 2 {
		return nil
	}

	p := 1.0
	g := f.copy()[1:]
	for i := range g {
		g[i] *= p
		p++
	}

	return g
}

// divide returns 1/a*f.
func divide(a float64, f Polynomial) Polynomial {
	return multiply(1.0/a, f)
}

// divide f by a.
func (f Polynomial) divide(a float64) {
	for i := range f {
		f[i] /= a
	}
}

// evaluate ...
func (f Polynomial) evaluate(x float64) float64 {
	var y float64 // y = f(x)
	p := 1.0      // p = x^i
	for i := range f {
		y += f[i] * p
		p *= x
	}

	return y
}

// integrate ...
func (f Polynomial) integrate() Polynomial {
	g := make(Polynomial, 1, len(f)+1)
	for i := range f {
		g = append(g, f[i]/float64(i+1))
	}

	return g
}

// multiply returns a*f.
func multiply(a float64, f Polynomial) Polynomial {
	g := f.copy()
	g.multiply(a)
	return g
}

// multiply f by a.
func (f Polynomial) multiply(a float64) {
	for i := range f {
		f[i] *= a
	}
}

// Subtract returns f-g.
func subtract(f, g Polynomial) Polynomial {
	h := f.copy()
	h.subtract(g)
	return h
}

// Subtract g from f.
func (f Polynomial) subtract(g Polynomial) {
	n := len(f)
	if n != len(g) {
		panic("degree mismatch")
	}

	for i := 0; i < n; i++ {
		f[i] -= g[i]
	}
}

// trim removes the higher powers that have zero valued coefficients.
func (f Polynomial) trim() Polynomial {
	n := len(f)
	if n == 0 {
		return nil
	}

	for ; 0 < n && f[n-1] == 0; n-- {
	}

	return f[:n].copy()
}
