package main

func main() {
	// fmt.Println(permTree(groups.Permutation{1, 3, 2, 0, 4}, groups.Permutation{0, 4, 1, 3, 2}))
}

// func testPerms0() {
// 	perms := set.Set{
// 		0: groups.Identity(3),
// 		1: groups.Permutation{1, 0, 2},
// 		2: groups.Permutation{0, 2, 1},
// 		3: groups.Permutation{2, 0, 1},
// 		4: groups.Permutation{1, 2, 0},
// 		5: groups.Permutation{2, 1, 0},
// 	}

// 	S3 := make(set.Set)
// 	for k := range perms {
// 		S3 = set.Union(S3, perms[k].(groups.Permutation).Generate())
// 	}

// 	fmt.Println(S3)
// }

// func permTree(a, b groups.Permutation) set.Set {
// 	S := set.New(nil)
// 	m, n := a.Order(), b.Order()
// 	a0, b0, a1, b1 := a.Copy(), b.Copy(), a.Copy(), b.Copy()
// 	for h := 0; h < m; h++ {
// 		for i := 0; i < n; i++ {
// 			for j := 0; j < m; j++ {
// 				for k := 0; k < n; k++ {
// 					S, _ = set.Insert(S, groups.Multiply(a0, b0, a1, b1))
// 					b1 = b1.Multiply(b)
// 				}
// 				a1 = a1.Multiply(a)
// 			}
// 			b0 = b0.Multiply(b)
// 		}
// 		a0 = a0.Multiply(a)
// 	}
// 	/*
// 		S := set.New(nil)
// 		S, _ = set.Insert(S, a)
// 		S, _ = set.Insert(S, b)
// 		a0, b0 := a.Multiply(a.Copy()), b.Multiply(b.Copy())
// 		a1, b1 := a0.Copy(), b0.Copy()
// 		for ; a0.CompareTo(a) != 0; a0 = a0.Multiply(a) {
// 			for ; b0.CompareTo(b) != 0; b0 = b0.Multiply(b) {
// 				for ; a1.CompareTo(a) != 0; a1 = a1.Multiply(a) {
// 					for ; b1.CompareTo(b) != 0; b1 = b1.Multiply(b) {
// 						S, _ = set.Insert(S, groups.ChainMultiply(a0, b0, a1, b1))
// 					}
// 				}
// 			}
// 		}
// 	*/

// 	return S
// }
