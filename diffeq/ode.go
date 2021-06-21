package diffeq

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	gomath "math"

	"github.com/nathangreene3/math"
	mtx "github.com/nathangreene3/math/linalg/matrix"
	vtr "github.com/nathangreene3/math/linalg/vector"
)

// An ODE (ordinary differential equation) is a function
// f(t, x0, x1, ...) = dx/dt for some x in {x0, x1, ...}.
type ODE func(t float64, v vtr.Vector) float64

// A SysODE (system of ODEs) is a list of ODEs.
type SysODE []ODE

// Eval ...
func (f SysODE) Eval(t float64, v0 vtr.Vector) vtr.Vector {
	n := len(f)
	if n != len(v0) {
		panic(dimErrStr)
	}

	v1 := vtr.Zeroes(n)
	for i := 0; i < n; i++ {
		v1[i] = f[i](t, v0)
	}

	return v1
}

// Euler ...
func Euler(f SysODE, v0 vtr.Vector, t0, t1, dt float64) vtr.Vector {
	n := len(f)
	if n != len(v0) {
		panic(dimErrStr)
	}

	var (
		d  = vtr.Zeroes(n)
		v1 = v0.Copy()
	)

	for ; t0 < t1; t0 += dt {
		for i := 0; i < n; i++ {
			d[i] = f[i](t0, v1) * dt
		}

		for i := 0; i < n; i++ {
			v1[i] += d[i]
		}
	}

	return v1
}

// EulerHistory ...
func EulerHistory(f SysODE, v0 vtr.Vector, t0, t1, dt float64) []vtr.Vector {
	n := len(f)
	if n != len(v0) {
		panic(dimErrStr)
	}

	vh := append(make([]vtr.Vector, 0, int((t1-t0)/dt)), v0.Copy())
	for ; t0 < t1; t0 += dt {
		vh = append(vh, vtr.Add(vh[len(vh)-1], vtr.Gen(n, func(i int) float64 { return f[i](t0, vh[len(vh)-1]) * dt })))
	}

	return vh
}

// Euler2 ...
func Euler2(f SysODE, v0 vtr.Vector, t0, t1, dt float64) vtr.Vector {
	n := len(f)
	if n != len(v0) {
		panic(dimErrStr)
	}

	var (
		k  = vtr.Zeroes(n) // f(t0, v1)
		w  = vtr.Zeroes(n) // Intermediate step
		v1 = v0.Copy()
	)

	for ; t0 < t1; t0 += dt {
		for i := 0; i < n; i++ {
			k[i] = f[i](t0, v1)
			w[i] = v1[i] + k[i]*dt
		}

		for i := 0; i < n; i++ {
			v1[i] += (k[i] + f[i](t0+dt, w)) * dt / 2
		}
	}

	return v1
}

// BookEuler2 ... (p. 105)
// func BookEuler2(f SysODE, v0 vtr.Vector, t0, t1, tol float64, M int) vtr.Vector {
// 	n := len(f)
// 	if n != len(v0) {
// 		panic(dimErrStr)
// 	}

// 	var (
// 		k  = vtr.Zeroes(n)
// 		w  = vtr.Zeroes(n)
// 		v1 = v0.Copy()
// 		z  = v0.Copy()
// 	)

// 	for i := 0; i < M; i++ {
// 		h := (t1 - t0) * float64(math.PowInt(2, -i))
// 		t := t0
// 		v := v1.Copy()
// 		for j := 0; j < math.PowInt(2, i); j++ {
// 			F := f.Eval(t, v)
// 			G := f.Eval(t+h, vtr.Add(v, vtr.Multiply(h, F)))
// 			t += h
// 			F.Add(G)
// 			F.Multiply(h / 2)
// 			v.Add(F)
// 		}

// 		if vtr.Approx(v,z)
// 	}

// 	return v1
// }

// Euler2History ...
func Euler2History(f SysODE, v0 vtr.Vector, t0, t1, dt float64) []vtr.Vector {
	n := len(f)
	if n != len(v0) {
		panic(dimErrStr)
	}

	var (
		k  = vtr.Zeroes(n) // f(t0, p0)
		w  = vtr.Zeroes(n) // Intermediate step
		v1 = append(make([]vtr.Vector, 0, int((t1-t0)/dt)), v0.Copy())
	)

	for ; t0 < t1; t0 += dt {
		for i := 0; i < n; i++ {
			k[i] = f[i](t0, v1[len(v1)-1])
			w[i] = v1[len(v1)-1][i] + k[i]*dt
		}

		v1 = append(v1, vtr.Gen(n, func(i int) float64 { return v1[len(v1)-1][i] + (k[i]+f[i](t0+dt, w))*dt/2 }))
	}

	return v1
}

// Heun ...
func Heun(f SysODE, v0 vtr.Vector, t0, t1, dt float64) vtr.Vector {
	n := len(f)
	if n != len(v0) {
		panic(dimErrStr)
	}

	var (
		k  = vtr.Zeroes(n) // f(t0, v1)
		w  = vtr.Zeroes(n) // Intermediate step
		v1 = v0.Copy()
	)

	for ; t0 < t1; t0 += dt {
		for i := 0; i < n; i++ {
			k[i] = f[i](t0, v1)
			w[i] = v1[i] + 2*k[i]*dt/3
		}

		for i := 0; i < n; i++ {
			v1[i] += (k[i] + 3*f[i](t0+2*dt/3, w)) * dt / 4
		}
	}

	return v1
}

// HeunHistory ...
func HeunHistory(f SysODE, v0 vtr.Vector, t0, t1, dt float64) []vtr.Vector {
	n := len(f)
	if n != len(v0) {
		panic(dimErrStr)
	}

	var (
		k  = vtr.Zeroes(n) // f(t0, p0)
		w  = vtr.Zeroes(n) // Intermediate step
		v1 = append(make([]vtr.Vector, 0, int((t1-t0)/dt)), v0.Copy())
	)

	for ; t0 < t1; t0 += dt {
		for i := 0; i < n; i++ {
			k[i] = f[i](t0, v1[len(v1)-1])
			w[i] = v1[len(v1)-1][i] + 2*k[i]*dt/3
		}

		v1 = append(v1, vtr.Gen(n, func(i int) float64 { return v1[len(v1)-1][i] + (k[i]+3*f[i](t0+2*dt/3, w))*dt/4 }))
	}

	return v1
}

// RungeKutta ...
func RungeKutta(f SysODE, v0 vtr.Vector, t0, t1, dt float64) vtr.Vector {
	n := len(f)
	if n != len(v0) {
		panic(dimErrStr)
	}

	var (
		// K0: f(t0, v1) dt
		// K1: f(t0+dt/2, v1+K0/2) dt
		// K2: f(t0+dt/2, v1+K1/2) dt
		// K3: f(t0+dt, v1+K2) dt
		K  = mtx.Zeroes(4, n)
		w  = vtr.Zeroes(n) // Intermediate step
		v1 = v0.Copy()
	)

	for ; t0 < t1; t0 += dt {
		for i := 0; i < n; i++ {
			K[0][i] = f[i](t0, v1) * dt
			w[i] = v1[i] + K[0][i]/2
		}

		for i := 0; i < n; i++ {
			K[1][i] = f[i](t0+dt/2, w) * dt
			w[i] = v1[i] + K[1][i]/2
		}

		for i := 0; i < n; i++ {
			K[2][i] = f[i](t0+dt/2, w) * dt
			w[i] = v1[i] + K[2][i]
		}

		for i := 0; i < n; i++ {
			K[3][i] = f[i](t0+dt, w) * dt
		}

		for i := 0; i < n; i++ {
			v1[i] += (K[0][i] + 2*(K[1][i]+K[2][i]) + K[3][i]) / 6
		}
	}

	return v1
}

// RungeKuttaHistory ...
func RungeKuttaHistory(f SysODE, v0 vtr.Vector, t0, t1, dt float64) []vtr.Vector {
	n := len(f)
	if n != len(v0) {
		panic(dimErrStr)
	}

	var (
		// K0: f(t0, v1) dt
		// K1: f(t0+dt/2, v1+K0/2) dt
		// K2: f(t0+dt/2, v1+K1/2) dt
		// K3: f(t0+dt, v1+K2) dt
		K  = mtx.Zeroes(4, n)
		w  = vtr.Zeroes(n) // Intermediate step
		v1 = append(make([]vtr.Vector, 0, int((t1-t0)/dt)), v0.Copy())
	)

	for ; t0 < t1; t0 += dt {
		for i := 0; i < n; i++ {
			K[0][i] = f[i](t0, v1[len(v1)-1]) * dt
			w[i] = v1[len(v1)-1][i] + K[0][i]/2
		}

		for i := 0; i < n; i++ {
			K[1][i] = f[i](t0+dt/2, w) * dt
			w[i] = v1[len(v1)-1][i] + K[1][i]/2
		}

		for i := 0; i < n; i++ {
			K[2][i] = f[i](t0+dt/2, w) * dt
			w[i] = v1[len(v1)-1][i] + K[2][i]
		}

		for i := 0; i < n; i++ {
			K[3][i] = f[i](t0+dt, w) * dt
		}

		v1 = append(v1, vtr.Gen(n, func(i int) float64 { return v1[len(v1)-1][i] + (K[0][i]+2*(K[1][i]+K[2][i])+K[3][i])/6 }))
	}

	return v1
}

// RungeKutta2 ...
func RungeKutta2(f SysODE, v0 vtr.Vector, t0, t1, tol float64, M int) vtr.Vector {
	n := len(f)
	if n != len(v0) {
		panic(dimErrStr)
	}

	var (
		dt = t1 - t0
		// K0: f(t0, v1) dt
		// K1: f(t0+dt/2, v1+K0/2) dt
		// K2: f(t0+dt/2, v1+K1/2) dt
		// K3: f(t0+dt, v1+K2) dt
		K  = mtx.Zeroes(4, n)
		w  = vtr.Zeroes(n) // Intermediate step
		v1 = v0.Copy()
		t  bool
		vi float64
	)

	for m := 0; m < M; m++ {
		for j := 0; j < math.PowInt(2, m); j++ {
			for i := 0; i < n; i++ {
				K[0][i] = f[i](t0, v1) * dt
				w[i] = v1[i] + K[0][i]/2
			}

			for i := 0; i < n; i++ {
				K[1][i] = f[i](t0+dt/2, w) * dt
				w[i] = v1[i] + K[1][i]/2
			}

			for i := 0; i < n; i++ {
				K[2][i] = f[i](t0+dt/2, w) * dt
				w[i] = v1[i] + K[2][i]
			}

			for i := 0; i < n; i++ {
				K[3][i] = f[i](t0+dt, w) * dt
			}

			t = true
			for i := 0; i < n; i++ {
				vi = (K[0][i] + 2*(K[1][i]+K[2][i]) + K[3][i]) / 6
				if t && tol <= gomath.Abs(vi) {
					t = false
				}

				v1[i] += vi
			}

			if t {
				return v1
			}
		}

		dt /= 2
	}

	return v1
}

// GIF2D ...
func GIF2D(history []vtr.Vector, w io.Writer) error {
	const (
		xMax  int   = 100 // Image x-axis will range over [-xMax, xMax]
		yMax  int   = 100 // Image y-axis will range over [-yMax, yMax]
		delay int   = 1   // x10ms
		black uint8 = 1   // Pixel color black
	)

	var (
		a = gif.GIF{LoopCount: len(history) / 10}
		p = []color.Color{color.White, color.Black}
		r = image.Rect(0, 0, 2*xMax+1, 2*yMax+2)
	)

	for i := 0; i < len(history); i += 10 {
		img := image.NewPaletted(r, p)
		img.SetColorIndex(xMax+int(history[i][0]*float64(xMax)), yMax+int(history[i][1]*float64(yMax)), black)
		a.Delay = append(a.Delay, delay)
		a.Image = append(a.Image, img)
	}

	return gif.EncodeAll(w, &a)
}
