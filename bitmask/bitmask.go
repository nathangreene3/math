package bitmask

import (
	"math/big"
	"sort"
)

// Bitmask ...
type Bitmask struct {
	value big.Int
}

// New returns a new *Bitmask
func New(n *big.Int) *Bitmask {
	b := big.NewInt(0)
	if n == nil {
		return &Bitmask{value: *b.And(b, n)}
	}

	return &Bitmask{value: *b.And(b, n)}
}

// Clear ...
func (b *Bitmask) Clear(n *big.Int) *Bitmask {
	b.value.AndNot(&b.value, n)
	return b
}

// Copy ...
func (b *Bitmask) Copy() *Bitmask {
	return New(&b.value)
}

// IsSet ...
func (b *Bitmask) IsSet(n *big.Int) bool {
	return b.value.And(&b.value, n).Cmp(n) == 0
}

// Set ...
func (b *Bitmask) Set(n *big.Int) *Bitmask {
	b.value.Or(&b.value, n)
	return b
}

// String ...
func (b *Bitmask) String() string {
	return b.value.String()
}

// Eratosthenes returns a list of prime integers up to and including n.
func Eratosthenes(n int) []int {
	// TODO: Finish after bitmask is done.
	if n < 2 {
		return nil
	}

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

// func TestEratosthenes(t *testing.T) {
// 	primes := Eratosthenes(75000)
// 	for _, p := range primes {
// 		if !IsPrime(p) {
// 			t.Fatalf("\np = %d is composite, not prime\n", p)
// 		}
// 	}
// }
