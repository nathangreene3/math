package function

import "github.com/nathangreene3/math"

// Bisect ...
func Bisect(f Function, x0, x1, precision float64) float64 {
	y0, y1 := f(x0), f(x1)
	switch {
	case y0 == 0:
		return x0
	case y1 == 0:
		return x1
	case 0 < y0*y1:
		panic("cannot find root on given range")
	case x1 < x0:
		x0, x1 = x1, x0
	}

	for {
		r := x0 + (x1-x0)/2.0
		y := f(r)
		switch {
		case f(x0)*y < 0:
			x1 = r
		case y*f(x1) < 0:
			x0 = r
		default:
			return r
		}

		if math.Approx(x1, x0, precision) {
			return r
		}
	}
}

func newton(f Function,x0,tol float64) float64 {
	// for x1:=x0-f(x0)/Diff2(f,x0,0.1,2);
	return 0
}
