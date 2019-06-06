package matrix

import (
	"fmt"
	"testing"
)

func TestMultiply(t *testing.T) {
	// tests:=[]struct{
	// 	A Matrix
	// 	B Matrix
	// 	expected Matrix
	// 	actual Matrix
	// }{
	// 	{
	// 		A:Matrix{
	// 			vector.Vector{1,2,3,4},
	// 			vector.Vector{2,1,-1,3},
	// 			vector.Vector{4,0,1,2},
	// 		},
	// 		B:Matrix{

	// 			vector.Vector{2,1,-1,3},
	// 			vector.Vector{,,,},
	// 			vector.Vector{,,,},
	// 		}
	// 	}
	// }
}

func TestChainMultiply(t *testing.T) {
	A := Empty(2, 10)
	B := Empty(10, 3)
	C := Empty(3, 8)
	// ABC:   108 mults
	// A(BC): 400 mults
	fmt.Println(ChainMultiply(A, B, C).String())
}
