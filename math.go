package math

import (
	gomath "math"
	"math/big"
	"math/bits"
	"sort"

	"github.com/nathangreene3/math/bitmask"
)

// Approx returns true if |x-y| <= prec, where prec in [0,1].
func Approx(x, y, prec float64) bool {
	if prec < 0 || 1 < prec {
		panic("precision must be on range [0,1]")
	}

	return gomath.Abs(x-y) <= prec
}

// Base10 converts a number n represented in base b to decimal.
func Base10(n []int, b int) int {
	var x int
	for i, v := range n {
		x += v * PowInt(b, i)
	}

	return x
}

// Base converts a number into its base representation.
func Base(n, b int) []int {
	switch {
	case n < 0:
		panic("number must be non-negative")
	case b < 2:
		panic("base must be greater than one")
	}

	var (
		remainders = make([]int, 0)
		k          int
	)

	for ; 0 < n; n = k {
		k = n / b
		remainders = append(remainders, n-k*b)
	}

	return remainders
}

// BasePows returns the powers of base b that sum to a number n.
func BasePows(n, b int) []int {
	switch {
	case n < 0:
		panic("number must be non-negative")
	case b < 2:
		panic("base must be greater than one")
	}

	var pows []int
	for bp := 1; bp <= n; bp *= b {
		pows = append(pows, bp)
	}

	for i := len(pows) - 1; 0 <= i; i-- {
		if n < pows[i] {
			pows[i] = 0
			continue
		}

		var c int // Number of times each power contributes to n
		for ; pows[i] <= n; c++ {
			n -= pows[i]
		}

		pows[i] *= c
	}

	return pows
}

// Choose returns n-Choose-k, or n!/(k!(n-k)!).
func Choose(n, k int) int {
	return Pascal(n + 1)[n][k]
}

// Cototient returns n-phi(n).
func Cototient(n int) int {
	return n - Totient(n)
}

// CoVar returns the covariance of two sets of values.
func CoVar(x, y []float64) float64 {
	n := len(x)
	if n != len(y) {
		panic("dimension mismatch")
	}

	var (
		mx, my = Mean(x...), Mean(y...)
		cv     float64
	)

	for i := 0; i < n; i++ {
		cv += (x[i] - mx) * (y[i] - my)
	}

	return cv / float64(n)
}

// Eratosthenes returns a list of prime integers up to and including n.
func Eratosthenes(n int) []int {
	// TODO: Finish after bitmask is done.
	if n < 2 {
		return nil
	}

	pm := make(map[int]struct{})
	bm := bitmask.New(big.NewInt(0))
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

// Factor returns a map of factors to the number of times they divide an integer
// n. That is, for each key-value pair (k,v), k divides n a total of v times.
// Each key will be a prime divisor, which means k will be at least two. An
// integer is prime if its only Factor is itself (and 1, which is called the
// empty prime).
func Factor(n int) map[int]int {
	if n < 1 {
		panic("cannot factor non-positive integer")
	}

	// Theorem: If p is prime greater than 3, then p = 6k-1 or 6k+1 for some
	// maximal k > 0.

	// This does NOT mean all 6k-1 or 6k+1 are prime, only that we have to check
	// numbers of the form 6k-1 and 6k+1 to be prime. This speeds up factoring
	// three times over.

	// 1. Determine how many times 2 and 3 divide n, if at all.
	// 2. Check all integers 6k-1 and 6k+1 less than sqrt(n) to help eliminate
	//    multiples of 2 and 3.
	factors := make(map[int]int)
	for ; n&1 == 0; n >>= 1 {
		factors[2]++
	}

	for ; n%3 == 0; n /= 3 {
		factors[3]++
	}

	for d := 5; d <= n; d += 4 {
		for ; n%d == 0; n /= d {
			factors[d]++
		}

		d += 2
		for ; n%d == 0; n /= d {
			factors[d]++
		}
	}

	return factors
}

// Factorial returns n!
func Factorial(n int) int {
	if n < 0 {
		panic("n must be non-negative")
	}

	f := 1
	for ; 1 < n; n-- {
		f *= n
	}

	return f
}

// Fibonacci returns the nth Fibonacci term, where the 0th and 1st terms are 1
// and the nth term is the sum of the previous two terms.
func Fibonacci(n int) int {
	a0, a1 := 1, 1
	for ; 1 < n; n-- {
		a0, a1 = a1, a0+a1
	}

	return a1
}

// GCD returns the largest divisor of both a and b. If GCD(a,b) == 1, then a and
// b are relatively prime.
func GCD(a, b int) int {
	if a < 0 || b < 0 {
		panic("invalid sign")
	}

	if a < b {
		a, b = b, a
	}

	for 0 < b {
		a, b = b, a%b
	}

	return a
}

// IsPrime indicates if n is prime.
func IsPrime(n int) bool {
	if n < 2 {
		return false // Prevents panic on non-positives and 1 isn't prime anyway
	}

	_, ok := Factor(n)[n]
	return ok
}

// LCM returns the largest multiple of a and b divisible by a and b.
func LCM(a, b int) int {
	if a < 1 || b < 1 {
		panic("a and b must be positive")
	}

	a1, b1 := a, b
	for a1 != b1 {
		for ; a1 < b1; a1 += a {
		}

		for ; b1 < a1; b1 += b {
		}
	}

	return a1
}

// Max returns the maximum of a list of values.
func Max(xs ...float64) float64 {
	max := xs[0]
	for _, x := range xs[1:] {
		if max < x {
			max = x
		}
	}

	return max
}

// MaxInt returns the maximum of a list of values.
func MaxInt(xs ...int) int {
	max := xs[0]
	for _, x := range xs[1:] {
		if max < x {
			max = x
		}
	}

	return max
}

// Mean returns the Mean (or average) of a list of values.
func Mean(xs ...float64) float64 {
	return Sum(xs...) / float64(len(xs))
}

// Min returns the minimum of a list of values.
func Min(xs ...float64) float64 {
	min := xs[0]
	for _, x := range xs[1:] {
		if x < min {
			min = x
		}
	}

	return min
}

// MinInt returns the minimum of a list of values.
func MinInt(xs ...int) int {
	min := xs[0]
	for _, x := range xs[1:] {
		if x < min {
			min = x
		}
	}

	return min
}

// NextPowOfTwo returns 2^k greater than or equal to n for the smallest k >= 0.
func NextPowOfTwo(n int) int {
	switch {
	case n < 1:
		return 1
	case n&(n-1) == 0:
		return n
	default:
		return 1 << bits.Len(uint(n))
	}
}

// Pascal returns Pascal's triangle, consisting of n levels. The (n,k)th entry
// is the value n!/(k!(n-k)!).
func Pascal(n int) [][]int {
	if n < 1 {
		panic("n must be positive")
	}

	tri := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		tri = append(tri, make([]int, 0, i+1))
		for j := 0; j < i+1; j++ {
			switch j {
			case 0, i:
				tri[i] = append(tri[i], 1)
			default:
				tri[i] = append(tri[i], tri[i-1][j-1]+tri[i-1][j])
			}
		}
	}

	return tri
}

// PowInt returns a^p for any integer a and non-zero integer p (exception: 0^0
// is undefined and will panic unlike most libraries).
func PowInt(a, p int) int {
	switch {
	case a == 0:
		if p == 0 {
			panic("indeterminant form")
		}
		return 0
	case p < 0:
		panic("p must be non-negative")
	case p == 0:
		return 1
	}

	// Yacca's method
	y := 1
	for ; 0 < p; p >>= 1 {
		if p&1 == 1 {
			y *= a
		}

		a *= a
	}

	return y
}

// StDev returns the standard deviation of a list of values.
func StDev(xs ...float64) float64 {
	return gomath.Sqrt(Var(xs...))
}

// Sum returns the sum of a list of values.
func Sum(xs ...float64) float64 {
	var s float64
	for _, x := range xs {
		s += x
	}

	return s
}

// SumInts returns the sum of a list of values.
func SumInts(xs ...int) int {
	var s int
	for _, x := range xs {
		s += x
	}

	return s
}

// Totient returns phi(n) = n prod(1-1/p) for all primes p such that p|n.
func Totient(n int) int {
	phi := 1
	for p, k := range Factor(n) {
		phi *= PowInt(p, k-1) * (p - 1)
	}

	return phi
}

// Var returns the Var of a list of values.
func Var(xs ...float64) float64 {
	var (
		m    = Mean(xs...)
		v, t float64
	)

	for _, x := range xs {
		t = x - m
		v += t * t
	}

	return v / float64(len(xs)-1)
}
