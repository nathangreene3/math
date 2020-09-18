package main

import (
	"fmt"

	vtr "github.com/nathangreene3/math/linalg/vector"
)

func main() {
	v := vtr.New(1, 2, 3)
	u := vtr.Unit(v)
	u.Add(vtr.Gen(3, func(i int) float64 { return 0 }), vtr.Gen(3, func(i int) float64 { return 0 }))
	t := v.Equal(u)
	fmt.Println(t)
}
