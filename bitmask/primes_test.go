package bitmask

import (
	"testing"

	"github.com/nathangreene3/math"
)

func TestEratosthenes(t *testing.T) {
	primes := Eratosthenes(2000000)
	for _, p := range primes {
		if !math.IsPrime(p) {
			t.Fatalf("\n%d = %v\n", p, math.Factor(p))
		}
	}
}
