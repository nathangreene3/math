package main

import (
	"fmt"

	mtx "github.com/nathangreene3/math/linalg/matrix"
	vtr "github.com/nathangreene3/math/linalg/vector"
)

func main() {
	A := mtx.New(
		vtr.New(1, 2),
		vtr.New(3, 4),
	)

	fmt.Println(mtx.Pow(mtx.Pow(A, 7), -1))
	A.Pow(-7)
	fmt.Println(A)
}
