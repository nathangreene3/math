package main

import (
	"fmt"

	"github.com/nathangreene3/math/groups"
	"github.com/nathangreene3/math/set"
)

func main() {
	perms := set.Set{
		0: groups.New(3),
		1: groups.Permutation{1, 0, 2},
		2: groups.Permutation{0, 2, 1},
		3: groups.Permutation{2, 0, 1},
		4: groups.Permutation{1, 2, 0},
		5: groups.Permutation{2, 1, 0},
	}

	S3 := make(set.Set)
	for k := range perms {
		S3 = set.Union(S3, groups.Generate(perms[k].(groups.Permutation)))
	}

	fmt.Println(S3)
}
