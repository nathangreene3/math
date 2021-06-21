package main

import (
	"bytes"
	"image"
	"image/color"
	"image/gif"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"

	ode "github.com/nathangreene3/math/diffeq"
	vtr "github.com/nathangreene3/math/linalg/vector"
)

func main() {
	var (
		G  = float64(1)
		m0 = float64(1)
		m1 = float64(1)
		v0 = vtr.New(1.25, 0, 0, 0.5)
		t0 = float64(0)
		t1 = float64(1)
		dt = float64(0.001)
		f  = ode.SysODE{
			func(t float64, v vtr.Vector) float64 { return v[2] },                                          // dx/dt = vx                                                 // dx/dt = -vx
			func(t float64, v vtr.Vector) float64 { return -G * m0 * m1 * v[0] / (v[0]*v[0] + v[1]*v[1]) }, // dvx/dt = -G*m0*m1*x/r^2
			func(t float64, v vtr.Vector) float64 { return v[3] },                                          // dy/dt = vy                                                 // dx/dt = -vx
			func(t float64, v vtr.Vector) float64 { return -G * m0 * m1 * v[1] / (v[0]*v[0] + v[1]*v[1]) }, // dvy/dt = -G*m0*m1*y/r^2
		}

		buf = bytes.NewBuffer(make([]byte, 0, 1024))
	)

	if err := ode.GIF2D(ode.RungeKuttaHistory(f, v0, t0, t1, dt), buf); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("gravity.gif", buf.Bytes(), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func lissajous(w io.Writer) error {
	const (
		white  uint8   = 0
		black  uint8   = 1
		xMax   int     = 100 // x-axis ranges over [-xMax, xMax], 2*xMax+1 distinct points wide/tall
		frames int     = 64
		delay  int     = 8 // x10ms
		cycles float64 = 5
		res    float64 = 0.001
		d      float64 = 0.1
	)

	var (
		p  = []color.Color{color.White, color.Black}
		f  = 3 * rand.Float64()
		a  = gif.GIF{LoopCount: frames}
		r  = image.Rect(0, 0, xMax, xMax)
		ph = float64(0)
	)

	for i := 0; i < frames; i++ {
		img := image.NewPaletted(r, p)
		for t := float64(0); t < 2*cycles*math.Pi; t += res {
			img.SetColorIndex(xMax+int(math.Sin(t)*float64(xMax)+0.5), xMax+int(math.Sin(t*f+ph)*float64(xMax)+0.5), black)
		}

		ph += d
		a.Delay = append(a.Delay, delay)
		a.Image = append(a.Image, img)
	}

	return gif.EncodeAll(w, &a)
}
