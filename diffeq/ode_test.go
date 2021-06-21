package diffeq

import (
	"math"
	"testing"

	mtx "github.com/nathangreene3/math/linalg/matrix"
	vtr "github.com/nathangreene3/math/linalg/vector"
)

func TestAll(t *testing.T) {
	/*
		{
			// Falling mass
			// A mass falls from an initial height and initial velocity. The range of motion is restricted to the vertical axis.
			var (
				g  float64 = 9.80665     // Accleration of gravity on Earth
				p0 Point   = Point{1, 0} // (y(t0), v(t0)) = (1, 0)
				t0 float64 = 0
				t1 float64 = (p0[1] + math.Sqrt(p0[1]*p0[1]+2*g*p0[0])) / g    // 0.451...
				p1 Point   = Point{p0[0] + p0[1]*t1 - g*t1*t1/2, p0[1] - g*t1} // (y(t1), v(t1)) = (0, -4.429...)
				f  SysODE  = SysODE{
					func(t float64, p Point) float64 { return p[1] }, // dy/dt = v
					func(t float64, p Point) float64 { return -g },   // dv/dt = -g
				}
				p1Approx Point = RungeKutta(f, p0, t0, t1, 0.000001)
			)

			t.Errorf(
				"\n(y(%v), v(%v)) = %v --> (y(%v), v(%v)) = %v\n"+
					"Runge-Kutta: %v\n"+
					"      Error: %v\n",
				t0, t0, p0,
				t1, t1, p1,
				p1Approx,
				Subtract(p1Approx, p1),
			)
		}
	*/

	/*
		{
			// Projectile motion
			// A mass is launched from an initial position at an initial velocity.
			var (
				g          float64 = 9.81 // Acceleration of gravity (m/s^2)
				t0         float64 = 0
				x0, y0     float64 = 0, 1 // Initial position (m, m)
				v0, theta0 float64 = 0, 0 // Initial velocity (m/s); Initial velocity angle (rad)
				p0         []Point = []Point{
					Point{x0, v0 * math.Cos(theta0)},      // (x(t0), vx(t0))
					Point{y0, v0*math.Sin(theta0) - g*t0}, // (y(t0), vy(t0))
				}
				t1 float64 = (p0[1][1] + math.Sqrt(p0[1][1]*p0[1][1]+2*g*p0[1][0])) / g
				p1 []Point = []Point{
					Point{p0[0][0], p0[0][1]*t1 + p0[0][1]},                    // (x(t1), vx(t1))
					Point{p0[1][0] + p0[1][1]*t1 - g*t1*t1/2, p0[1][1] - g*t1}, // (y(t1), vy(t1))
				}
				f []SysODE = []SysODE{
					SysODE{
						func(t float64, p Point) float64 { return v0 * math.Cos(theta0) }, // vx(t)
						func(t float64, p Point) float64 { return 0 },                     // 0
					},
					SysODE{
						func(t float64, p Point) float64 { return p[1] }, // vy(t)
						func(t float64, p Point) float64 { return -g },   // -g
					},
				}
				p1Approx []Point = SolveMultiDim(RungeKutta, f, p0, t0, t1, 0.000001)
			)

			t.Errorf(
				"\n[(x(%v), vx(%v)), (y(%v), vy(%v))] = [%v, %v] --> [(x(%v), vx(%v)), (y(%v), vy(%v))] = [%v, %v]\n"+
					"Runge-Kutta: [(x(%v), vx(%v)), (y(%v), vy(%v))] = [%v, %v]\n"+
					"      Error: [%v, %v]\n",
				t0, t0, t0, t0, p0[0], p0[1],
				t1, t1, t1, t1, p1[0], p1[1],
				t1, t1, t1, t1, p1Approx[0], p1Approx[1],
				Subtract(p1Approx[0], p1[0]), Subtract(p1Approx[1], p1[1]),
			)
		}
	*/

	/*
		{
			// Simple pendulum
			var (
				g  float64 = 9.81
				l  float64 = 1.0
				x0         = []float64{
					math.Pi / 4.0, // theta(0)
					0.0,           // velocity(0)
				}
				fs = SysODE{
					func(t float64, xs Point) float64 { return xs[1] / l },            // d/dt theta
					func(t float64, xs Point) float64 { return -g * math.Sin(xs[0]) }, // d/dt velocity
				}
			)

			t.Error(Euler(fs, x0, 0, 1, 0.001))
		}

		{
			// Chapter 5, Example 1 from Numerical Analysis, 7th Ed. by Burden and Faires
			// Solve dx/dt = x - t^2 + 1 for x(t1) given x(t0) = 0.5, t0 = 0, t1 = 2
			// General solution is x(t1) = (x(t0) - (t0+1)^2)*exp(t1-t0) + (t1+1)^2
			var (
				x0 float64 = 0.5
				t0 float64 = 0
				t1 float64 = 2
				dt float64 = 0.1                                                       // Step
				x  Point  d = Point{(x0-(t0+1)*(t0+1))*math.Exp(t1-t0) + (t1+1)*(t1+1)} // Correct answer
			)

			t.Errorf(
				"\n     Answer: %g\n"+
					"      Euler: %e\n"+
					"     Euler2: %e\n"+
					"       Heun: %e\n"+
					"Runge-Kutta: %e\n",
				x,
				err(x, Euler(SysODE{func(t float64, x Point) float64 { return x[0] - t*t + 1 }}, Point{x0}, t0, t1, dt)),
				err(x, Euler2(SysODE{func(t float64, x Point) float64 { return x[0] - t*t + 1 }}, Point{x0}, t0, t1, dt)),
				err(x, Heun(SysODE{func(t float64, x Point) float64 { return x[0] - t*t + 1 }}, Point{x0}, t0, t1, dt)),
				err(x, RungeKutta(SysODE{func(t float64, x Point) float64 { return x[0] - t*t + 1 }}, Point{x0}, t0, t1, dt)),
			)
		}
	*/

	/*
		{
			// Chapter 3, Example 3
			// Fundamentals of Differential Equations, erd Ed.

			// dx/dt = x^2, x0 = 1

			var (
				t0 = float64(0)
				t1 = float64(2)
				dt = float64(0.25)
				p0 = Point{1} // x0 = 1
				f  = SysODE{func(t float64, p Point) float64 { return p[0] * p[0] }}
			)

			r := RungeKuttaHistory(f, p0, t0, t1, dt)
			// t.Fatal(r)
			_ = r
		}
	*/

	// subtestEulersNumber(t)
	// subtestRosettaCode(t)
	// subtestPredatorPrey(t)
	// subtestSecondOrderLinear(t)
	subtestGravity(t)
}

// subtestEulersNumber approximates e.
func subtestEulersNumber(t *testing.T) bool {
	// Fundamentals of Differential Equations, 3rd Ed.
	// Chapter 3, p. 101

	// dx/dt = x --> x = ce^t

	r := t.Run(
		"Approximating Euler's Number",
		func(s *testing.T) {
			var (
				v0 = vtr.New(1) // x0 = e^0 = 1
				t0 = float64(0)
				t1 = float64(1)
				dt = float64(0.01)
				f  = SysODE{func(t float64, v vtr.Vector) float64 { return v[0] }}
				R  = mtx.New(
					Euler(f, v0, t0, t1, dt),
					Euler2(f, v0, t0, t1, dt),
					Heun(f, v0, t0, t1, dt),
					RungeKutta(f, v0, t0, t1, dt),
				)
			)

			s.Errorf(
				"\n+-------------+----------+----------+\n"+
					"| Method      |   Result |    Error |\n"+
					"+-------------+----------+----------+\n"+
					"| Actual      | %f |          |\n"+
					"| Euler       | %f | %e |\n"+
					"| Euler2      | %f | %e |\n"+
					"| Heun        | %f | %e |\n"+
					"| Runge-Kutta | %f | %e |\n"+
					"+-------------+----------+----------+",
				math.E,
				R[0][0],
				math.E-R[0][0],
				R[1][0],
				math.E-R[1][0],
				R[2][0],
				math.E-R[2][0],
				R[3][0],
				math.E-R[3][0],
			)
		},
	)

	return r
}

func subtestSecondOrderLinear(t *testing.T) bool {
	// Fundamentals of Differential Equations, 3rd Ed.
	// Chapter 5, p. 233, Ex. 2

	// x'' + 3x' +2x = 0

	r := t.Run(
		"Approximating x'' + 3x' +2x = 0",
		func(s *testing.T) {
			var (
				v0 = vtr.New(1, 1)
				t0 = float64(0)
				t1 = float64(1)
				dt = float64(0.125)
				f  = SysODE{
					func(t float64, v vtr.Vector) float64 { return v[1] },
					func(t float64, v vtr.Vector) float64 { return -2*v[0] - 3*v[1] },
				}
				R = mtx.New(
					Euler(f, v0, t0, t1, dt),
					Euler2(f, v0, t0, t1, dt),
					Heun(f, v0, t0, t1, dt),
					RungeKutta(f, v0, t0, t1, dt),
				)
				exp = float64(0.83295)
			)

			s.Errorf(
				"\n+-------------+----------+---------------+\n"+
					"| Method      |   Result |         Error |\n"+
					"+-------------+----------+---------------+\n"+
					"| Actual      | %f |               |\n"+
					"| Euler       | %f |  %e |\n"+
					"| Euler2      | %f |  %e |\n"+
					"| Heun        | %f |  %e |\n"+
					"| Runge-Kutta | %f | %e |\n"+
					"+-------------+----------+---------------+",
				exp,
				R[0][0],
				exp-R[0][0],
				R[1][0],
				exp-R[1][0],
				R[2][0],
				exp-R[2][0],
				R[3][0],
				exp-R[3][0],
			)
		},
	)

	return r
}

func subtestGravity(t *testing.T) bool {
	f := func(t0 *testing.T) {
		var (
			mu  float64 = 1
			p0          = vtr.New(1, 0, 0, 1) // r, theta, r', theta'
			sys SysODE  = SysODE{
				func(t float64, p vtr.Vector) float64 { return p[2] },                       // r'
				func(t float64, p vtr.Vector) float64 { return p[3] },                       // theta'
				func(t float64, p vtr.Vector) float64 { return p[0]*p[3] - mu/(p[0]*p[0]) }, // r''
				func(t float64, p vtr.Vector) float64 { return -2 * p[2] * p[3] / p[0] },    // theta''
			}
		)

		t.Fatal(RungeKutta(sys, p0, 0, 10, 0.01))
	}
	return t.Run("", f)
}

// subtestRosettaCode solves the problem used by Rosetta Code to demonstrate language differences.
// https://rosettacode.org/wiki/Runge-Kutta_method
func subtestRosettaCode(t *testing.T) bool {
	// dx/dt = tx^0.5, x0 = 1 --> x = (t^2+4)^2/16

	r := t.Run(
		"Rosetta Code Example",
		func(s *testing.T) {
			var (
				v0 = vtr.New(1) // x0 = 1
				t0 = float64(0)
				t1 = float64(10)
				dt = float64(0.0001)
				f  = SysODE{func(t float64, v vtr.Vector) float64 { return t * math.Sqrt(v[0]) }}
				x  = func(t float64) float64 { y := t*t + 4; return y * y / 16 }(t1)
				R  = mtx.New(
					Euler(f, v0, t0, t1, dt),
					Euler2(f, v0, t0, t1, dt),
					Heun(f, v0, t0, t1, dt),
					RungeKutta(f, v0, t0, t1, dt),
				)
			)

			s.Errorf(
				"\n+-------------+----------+----------+\n"+
					"| Method      |   Result |    Error |\n"+
					"+-------------+----------+----------+\n"+
					"| Actual      | %f |          |\n"+
					"| Euler       | %f | %e |\n"+
					"| Euler2      | %f | %e |\n"+
					"| Heun        | %f | %e |\n"+
					"| Runge-Kutta | %f | %e |\n"+
					"+-------------+----------+----------+\n",
				x,
				R[0][0],
				x-R[0][0],
				R[1][0],
				x-R[1][0],
				R[2][0],
				x-R[2][0],
				R[3][0],
				x-R[3][0],
			)
		},
	)

	return r
}

func subtestPredatorPrey(t *testing.T) bool {
	// Predators: dx/dt = 2x - 2xy, x0 = 1
	// Prey:      dy/dt = -y + xy,  y0 = 3

	r := t.Run(
		"Predator vs Prey Model",
		func(s *testing.T) {
			var (
				v0 = vtr.New(1, 3)
				t0 = float64(0)
				t1 = float64(1)
				dt = float64(0.01)
				f  = SysODE{
					func(t float64, v vtr.Vector) float64 { return 2*v[0] - 2*v[0]*v[1] },
					func(t float64, v vtr.Vector) float64 { return -v[1] + v[0]*v[1] },
				}

				R = mtx.New(
					Euler(f, v0, t0, t1, dt),
					Euler2(f, v0, t0, t1, dt),
					Heun(f, v0, t0, t1, dt),
					RungeKutta(f, v0, t0, t1, dt),
					RungeKutta2(f, v0, t0, t1, 0.0001, 10000),
				)
			)

			s.Errorf(
				"\n+-------------+----------+----------+\n"+
					"| Method      | Predator |     Prey |\n"+
					"+-------------+----------+----------+\n"+
					"| Euler       | %f | %f |\n"+
					"| Euler2      | %f | %f |\n"+
					"| Heun        | %f | %f |\n"+
					"| RungeKutta  | %f | %f |\n"+
					"| RungeKutta2 | %f | %f |\n"+
					"+-------------+----------+----------+\n",
				R[0][0], R[0][1],
				R[1][0], R[1][1],
				R[2][0], R[2][1],
				R[3][0], R[3][1],
				R[4][0], R[4][1],
			)
		},
	)

	return r
}

func TestHistory(t *testing.T) {
	/*
		{
			// Falling mass
			// A mass falls from an initial height and initial velocity. The range of motion is restricted to the vertical axis.
			var (
				g float64 = 9.80665 // Accleration of gravity on Earth
				// height = func(t float64, p0 Point) Point { return Point{p0[0] + p0[1]*t - g*t*t/2, p0[1] - g*t} } // Height of mass at time t given intial condition p0

				// Initial conditions
				t0 float64 = 0
				p0 Point   = Point{1, 0} // (y(t0), v(t0)) = (1, 0)

				// Final conditions (p1 is the expected solution)
				t1 float64 = (p0[1] + math.Sqrt(p0[1]*p0[1]+2*g*p0[0])) / g // 0.451...
				// p1 Point = height(t1, p0) // (y(t1), v(t1)) = (0, -4.429...)

				// System to solve that approximates p1
				f SysODE = SysODE{
					func(t float64, p Point) float64 { return p[1] }, // dy/dt = v
					func(t float64, p Point) float64 { return -g },   // dv/dt = -g
				}

				// Approximate solution
				p1Approx   []Point   = EulerHistory(f, p0, t0, t1, 0.1)  // Approximate p1
				heights    []float64 = make([]float64, 0, len(p1Approx)) // All ys from p1Approx to plot
				velocities []float64 = make([]float64, 0, len(p1Approx)) // All vs from p1Approx to plot
			)

			for i := 0; i < len(p1Approx); i++ {
				heights = append(heights, p1Approx[i][0])
				velocities = append(velocities, p1Approx[i][1])
			}

			t.Errorf(
				"\n%s\n\n%s\n",
				asciigraph.Plot(
					heights,
					asciigraph.Caption("Height [m] of a falling mass over time [s]"),
					asciigraph.Height(10),
				),
				asciigraph.Plot(
					velocities,
					asciigraph.Caption("Velocity [m/s] of a falling mass over time [s]"),
					asciigraph.Height(10),
				),
			)
		}
	*/

	{
		// Kirchhoff's law
		// Numerical Analyis, 7th Ed.
		// Ch 5, Ex 1

		// dx/dt = -4x + 3y + 6       --> ...TODO: Analytic solution
		// dy/dt = 0.6 dx/dt - 0.2y
		//       = -2.4x + 1.6y + 3.6

		var (
			v0 = vtr.New(0, 0) // (x0, y0)
			t0 = float64(0)
			t1 = float64(0.1)
			dt = float64(0.1)
			f  = SysODE{
				func(t float64, v vtr.Vector) float64 { return -4*v[0] + 3*v[1] + 6 },
				func(t float64, v vtr.Vector) float64 { return -2.4*v[0] + 1.6*v[1] + 3.6 },
			}
		)

		t.Error(RungeKutta(f, v0, t0, t1, dt))
	}

	// {
	// 	// Gravitational orbit
	// 	// A mass m1 orbits another mass m0 (taken as the origin of the system). It is assumed the acceleration d^2/dt^2(x,y) = -G*m0*m1/r^2 (x,y).
	// 	var (
	// 		G      float64 = physconst.Gravitation
	// 		m0, m1 float64 = 100, 1

	// 		p0 Point  = Point{1, 0, 0, 1} // (x0, y0, vx0, vy0)
	// 		f  SysODE = SysODE{
	// 			func(t float64, p Point) float64 { return p[2] },                                                     //  dx/dt = vx
	// 			func(t float64, p Point) float64 { return p[3] },                                                     //  dy/dt = vy
	// 			func(t float64, p Point) float64 { return -G * m0 * m1 * p[0] / math.Pow(p[0]*p[0]+p[1]*p[1], 1.5) }, // dvx/dt = -G*m0*m1*x/(x^2 + y^2)^(3/2)
	// 			func(t float64, p Point) float64 { return -G * m0 * m1 * p[1] / math.Pow(p[0]*p[0]+p[1]*p[1], 1.5) }, // dvy/dt = -G*m0*m1*y/(x^2 + y^2)^(3/2)
	// 		}

	// 		p1 []Point = RungeKuttaHistory(f, p0, 0, 100, 0.001)
	// 	)

	// 	t.Error(p1[len(p1)-1])
	// }
}

// BenchmarkEuler-4   	  750571	      1931 ns/op	      16 B/op	       2 allocs/op
func BenchmarkEuler(b *testing.B) {
	var (
		t0 = float64(0)
		t1 = float64(1)
		dt = float64(0.01)
		v0 = vtr.New(1)
		f  = SysODE{func(t float64, v vtr.Vector) float64 { return v[0] }}
	)

	for i := 0; i < b.N; i++ {
		_ = Euler(f, v0, t0, t1, dt)
	}
}

func BenchmarkLenAccessTime(b *testing.B) {
	subBenchLen(b)
	subBenchInt(b)
}

// subBenchLen ...
func subBenchLen(b *testing.B) bool {
	r := b.Run(
		"Accessing Slice Len",
		func(sub *testing.B) {
			var s []int
			for i := 0; i < sub.N; i++ {
				_ = len(s)
			}
		},
	)

	return r
}

// subBenchInt ...
func subBenchInt(b *testing.B) bool {
	r := b.Run(
		"Accessing int",
		func(sub *testing.B) {
			var n int
			for i := 0; i < sub.N; i++ {
				_ = n
			}
		},
	)

	return r
}
