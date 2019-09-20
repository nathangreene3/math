package main

import (
	"fmt"

	"github.com/nathangreene3/math"
	"github.com/nathangreene3/table"
)

func main() {
	var (
		left, right    int
		xFacts, yFacts string
		n              = 1 << 10
		t              = table.New("", table.FltFmtNoExp, 0, n, 6)
	)

	t.SetHeader(table.Header{"x", "y", "(x^2+x)/2", "y^2", "Facts of x", "Facts of y"})
	for x := 0; x < n; x++ {
		left = x * (x + 1) >> 1
		if x == 0 {
			xFacts = ""
		} else {
			xFacts = fmt.Sprint(math.Factor(x))
		}

		for y := 0; y < n; y++ {
			right = y * y
			if y == 0 {
				yFacts = ""
			} else {
				yFacts = fmt.Sprint(math.Factor(y))
			}

			if left == right {
				t.AppendRow(table.Row{x, y, left, right, xFacts, yFacts})
			}
		}
	}

	fmt.Println(t.String())
}
