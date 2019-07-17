package polynomial

import (
	"github.com/nathangreene3/math/stats"
)

// A Polynomial is an ordered set of weights f = [a0, a1, ..., an-1] such that f(x) = a0 + a1*x + ... + an-1 x^(n-1) for any real x.
type Polynomial []float64

func (f Polynomial) degree() int {
	return stats.MaxInt(len(f)-1, 0)
}

func (f Polynomial) evaluate(x float64) float64 {
	var y float64 // y = f(x)
	p := 1.0      // p = x^i
	for i := range f {
		y += f[i] * p
		p *= x
	}

	return y
}

func add(f, g Polynomial) Polynomial {
	m, n := f.degree(), g.degree()
	if n < m {
		return add(g, f)
	}

	h := make(Polynomial, 0, n)
	for i := 0; i < m; i++ {
		h = append(h, f[i]+g[i])
	}

	for i := m; i < n; i++ {
		h = append(h, g[i])
	}

	return h
}

func subtract(f, g Polynomial) Polynomial {
	return add(f, multiply(g, -1))
}

func multiply(f Polynomial, a float64) Polynomial {
	for i := range f {
		f[i] *= a
	}

	return f
}

func divide(f Polynomial, a float64) Polynomial {
	return multiply(f, 1.0/a)
}

func differentiate(f Polynomial) Polynomial {
	n := len(f)
	g := make(Polynomial, 0, n-1)
	for i := 1; i < n; i++ {
		g = append(g, float64(i)*f[i])
	}

	return g
}

func integrate(f Polynomial) Polynomial {
	g := make(Polynomial, 1, len(f)+1)
	for i := range f {
		g = append(g, f[i]/float64(i+1))
	}

	return g
}
