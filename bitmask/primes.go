package bitmask

import (
	"math/big"
	"sort"
)

// Eratosthenes returns a list of prime integers up to and including n.
func Eratosthenes(n int) []int {
	pm := make(map[int]struct{})
	bm := New(big.NewInt(0))
	for k := 2; k <= n; k++ {
		pm[k] = struct{}{}
		bm.Set(big.NewInt(int64(k)))
	}

	for p := 2; p <= n; p++ {
		for k := range pm {
			if p != k && k/p*p == k {
				delete(pm, k)
				bm.Clear(big.NewInt(int64(k)))
			}
		}
	}

	primes := make([]int, 0, n)
	for p := range pm {
		primes = append(primes, p)
	}

	sort.Ints(primes)
	return primes
}

// Eratosthenes2 ...
func Eratosthenes2(n int) []int {
	// b := New(big.NewInt(math.PowInt(2,n)))

	return nil
}
