package main

import (
	"fmt"

	"github.com/nathangreene3/math/groups"
)

func main() {
	perms := []groups.Permutation{
		// groups.New(3),
		groups.Permutation{1, 0, 2},
		groups.Permutation{0, 2, 1},
		groups.Permutation{2, 0, 1},
		groups.Permutation{1, 2, 0},
		groups.Permutation{2, 1, 0},
	}
	m := make([][]groups.Permutation, 0, 5)
	for i := 0; i < 5; i++ {
		m = append(m, make([]groups.Permutation, 0, 5))
		for j := 0; j < 5; j++ {
			m[i] = append(m[i], groups.Multiply(perms[i], perms[j]))
		}

		fmt.Println(m[i])
	}

}
