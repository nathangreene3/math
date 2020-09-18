package math

import (
	"fmt"
	gomath "math"
	"math/bits"
	"strconv"
)

// Approx determines if the absolute difference in two values is less than or equal to some tolerance. That is, |x-y| <= tol is returned.
func Approx(x, y, tol float64) bool {
	if tol < 0 || 1 < tol {
		panic("tolerance must be on range [0,1]")
	}

	return gomath.Abs(x-y) <= tol
}

// Base10 converts a number n represented in base b to decimal.
func Base10(n []int, b int) int {
	var x int
	for i := 0; i < len(n); i++ {
		x += n[i] * PowInt(b, i)
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
		rs = make([]int, 0) // Remainders of n/b
		k  int              // n/b
	)

	for ; 0 < n; n = k {
		k = n / b
		rs = append(rs, n-k*b)
	}

	return rs
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
	nFact, kFact, nkFact := 1, 1, 1 // n!, k!, (n-k)!
	if k<<1 < n {                   // k < n-k
		for i := 2; i <= k; i++ {
			kFact *= i
		}

		nkFact = kFact
		for i := k + 1; i+k <= n; i++ {
			nkFact *= i
		}

		nFact = nkFact
		for i := n - k + 1; i <= n; i++ {
			nFact *= i
		}
	} else {
		for i := 2; i+k <= n; i++ {
			nkFact *= i
		}

		kFact = nkFact
		for i := n - k + 1; i <= k; i++ {
			kFact *= i
		}

		nFact = kFact
		for i := k + 1; i <= n; i++ {
			nFact *= i
		}
	}

	return nFact / (kFact * nkFact)
}

// Cototient returns n-phi(n).
func Cototient(n int) int {
	return n - Totient(n)
}

// CoVar returns the covariance of two sets of values.
func CoVar(x, y []float64) float64 {
	if len(x) != len(y) {
		panic("dimension mismatch")
	}

	var (
		mx, my = Mean(x...), Mean(y...)
		cv     float64
	)

	for i := 0; i < len(x); i++ {
		cv += (x[i] - mx) * (y[i] - my)
	}

	return cv / float64(len(x))
}

// Fact returns n!
func Fact(n int) int {
	if n < 0 {
		panic("n must be non-negative")
	}

	f := 1
	for ; 1 < n; n-- {
		f *= n
	}

	return f
}

func factRec(n int) int {
	switch {
	case n < 0:
		panic("n must be non-negative")
	case n < 2:
		return 1
	default:
		return n * factRec(n-1)
	}
}

func factTail(n int) int {
	return factAcc(n, 1)
}

func factAcc(n, a int) int {
	switch {
	case n < 0:
		panic("factAcc: n must be non-negative")
	case n < 2:
		return 1
	default:
		return factAcc(n-1, n*a)
	}
}

// Factor returns a map of factors to the number of times they divide an integer
// n. That is, for each key-value pair (k,v), k divides n a total of v times.
// Each key will be a prime divisor, which means k will be at least two. An
// integer is prime if its only factor is itself (and 1, which is called the
// empty prime). Since the values are the number of times a factor divides a the
// given number, one is not included as it can divide anything infinitely many
// times.
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
	// 2. Check all integers 6k-1 and 6k+1 less than sqrt(n) to eliminate
	//    multiples of 2 and 3.
	var (
		factors = make(map[int]int)
		c       int // Number of times a divisor divides n
	)

	// 1.a Check how many times 2 divides n
	for ; n&1 == 0; n >>= 1 {
		c++
	}

	if 0 < c {
		factors[2] = c
		c = 0
	}

	// 1.b Check how many times 3 divides n
	for ; n%3 == 0; n /= 3 {
		c++
	}

	if 0 < c {
		factors[3] = c
		c = 0
	}

	// 2. Check how many times each integer 6k+/-1 divides n
	for d := 5; d <= n; d += 4 {
		for ; n%d == 0; n /= d {
			c++
		}

		if 0 < c {
			factors[d] = c
			c = 0
		}

		for d += 2; n%d == 0; n /= d {
			c++
		}

		if 0 < c {
			factors[d] = c
			c = 0
		}

		if int(gomath.Sqrt(float64(n))) <= d {
			factors[n]++
			break
		}
	}

	return factors
}

func factor2(n int) map[int]int {
	var (
		f  = make(map[int]int)
		ps = primes(n)
	)

	for i := 0; i < len(ps); i++ {
		for p := ps[i]; n%p == 0; n /= p {
			f[p]++
		}
	}

	return f
}

// primes256 is a list of all primes less than 256.
var primes256 = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251}

func factor3(n int) map[int]int {
	factors := make(map[int]int)
	for i := 0; i < len(primes256) && primes256[i] <= n; i++ {
		for ; n%primes256[i] == 0; n /= primes256[i] {
			factors[primes256[i]]++
		}
	}

	// 6*43 - 1 = 257
	for d := 257; d <= n; d += 4 {
		for ; n%d == 0; n /= d {
			factors[d]++
		}

		for d += 2; n%d == 0; n /= d {
			factors[d]++
		}
	}

	return factors
}

func factor4(n int) map[int]int {
	if n < 1 {
		panic("cannot factor non-positive integer")
	}

	var (
		factors = make(map[int]int)
		c       int // Number of times a divisor divides n
	)

	for ; n&1 == 0; n >>= 1 {
		c++
	}

	if 0 < c {
		factors[2] = c
		c = 0
	}

	for d := fermatFactor(n); 1 < n; d = fermatFactor(n) {
		switch d {
		case 0:
			panic("n is even?")
		case 1:
			// n is now an odd prime
			factors[n]++
			return factors
		default:
			factors[d]++
			n /= d
		}
	}

	return factors
}

func fermatFactor(n int) int {
	if n&1 == 0 {
		panic("only odd numbers can be factored as a^2-b^2")
	}

	var (
		fn = float64(n)
		a  = gomath.Ceil(gomath.Sqrt(fn))
		b  = gomath.Sqrt(a*a - fn)
	)

	for ; float64(int64(b)) != b; a, b = a+1, gomath.Sqrt(a*a-fn) {
	}

	return int(a - b)
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

// fibonacciRec is the naive solution.
func fibonacciRec(n int) int {
	if n < 2 {
		return 1
	}

	return fibonacciRec(n-2) + fibonacciRec(n-1)
}

func fibonacciTail(n int) int {
	return fibonacciAcc(n, 1, 1)
}

func fibonacciAcc(n, a, b int) int {
	if n < 2 {
		return a
	}

	return fibonacciAcc(n-1, b, a+b)
}

// Freq ...
func Freq(xs ...float64) map[float64]int {
	m := make(map[float64]int)
	for i := 0; i < len(xs); i++ {
		m[xs[i]]++
	}

	return m
}

// GCD returns the largest divisor of both a and b. If GCD(a,b) == 1, then a and
// b are relatively prime.
func GCD(a, b int) int {
	if a < 1 || b < 1 {
		panic("invalid sign")
	}

	if b < a {
		a, b = b, a
	}

	for ; 0 < a; a, b = b%a, a {
	}

	return b
}

// IsInt determines if a value is an integer.
func IsInt(n float64) bool {
	return float64(int64(n)) == n
}

// IsEven determines if a value is even.
func IsEven(n int) bool {
	return n&1 == 0
}

// IsOdd determines if a value is odd.
func IsOdd(n int) bool {
	return n&1 == 1
}

// IsPrime indicates if n is prime.
func IsPrime(n int) bool {
	return 1 < n && Factor(n)[n] == 1
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
	m := xs[0]
	for i := 1; i < len(xs); i++ {
		if m < xs[i] {
			m = xs[i]
		}
	}

	return m
}

// MaxInt returns the maximum of a list of values.
func MaxInt(xs ...int) int {
	m := xs[0]
	for i := 1; i < len(xs); i++ {
		if m < xs[i] {
			m = xs[i]
		}
	}

	return m
}

// Mean returns the Mean (or average) of a list of values.
func Mean(xs ...float64) float64 {
	return Sum(xs...) / float64(len(xs))
}

// Min returns the minimum of a list of values.
func Min(xs ...float64) float64 {
	m := xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] < m {
			m = xs[i]
		}
	}

	return m
}

// MinInt returns the minimum of a list of values.
func MinInt(xs ...int) int {
	m := xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] < m {
			m = xs[i]
		}
	}

	return m
}

// Mode returns the mode of a list of values.
func Mode(xs ...float64) float64 {
	var (
		m  = map[float64]int{xs[0]: 1}
		md = xs[0]
	)

	for i := 1; i < len(xs); i++ {
		m[xs[i]]++
		if m[md] < m[xs[i]] {
			md = xs[i]
		}
	}

	return md
}

// NextPowOfTwo returns 2^k greater than or equal to n for minimal k >= 0.
func NextPowOfTwo(n int) int {
	switch {
	case n < 0:
		return -NextPowOfTwo(-n)
	case n == 0:
		return 1
	case n&(n-1) == 0:
		// n = 2^k, for some k >= 0
		return n
	default:
		// 2^k > n where k bits represent n in base 2 (disregarding the leading sign bit)
		return 1 << bits.Len(uint(n))
	}
}

// prevPowOfTwo: TODO
func prevPowOfTwo(n int) int {
	switch {
	case n < 0:
		return -prevPowOfTwo(-n)
	case n < 2:
		panic("")
	case n&(n-1) == 0:
		return n
	default:
		return 1 << (bits.Len(uint(n)) - 1)
	}
}

// Pascal returns Pascal's triangle, consisting of n levels. The (n,k)th entry
// is the value n!/(k!(n-k)!), where 0 <= n and 0 <= k <= n.
func Pascal(n int) [][]int {
	if n < 0 {
		panic("n must be positive")
	}

	p := make([][]int, 0, n+1)
	for i := 0; i <= n; i++ {
		p = append(p, make([]int, 0, i+1))
		for j := 0; j < i+1; j++ {
			switch j {
			case 0, i:
				p[i] = append(p[i], 1)
			default:
				p[i] = append(p[i], p[i-1][j-1]+p[i-1][j])
			}
		}
	}

	return p
}

// PowInt returns a^p for any integer a and non-negative integer p (exception: 0^0
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
	}

	// Successive squaring
	// Yacca's method (source?)
	y := 1
	for ; 0 < p; p >>= 1 {
		if p&1 == 1 {
			y *= a
		}

		a *= a
	}

	return y
}

func primes(n int) []int {
	if maxInt16 < n {
		panic(fmt.Sprintf("Primes: %d is the largest value supported", maxInt16))
	}

	// Seive of Eratosthenes
	b := make([]bool, 2, n+1) // b = [false, false] to begin with
	for i := 2; i <= n; i++ {
		b = append(b, true)
	}

	p := make([]int, 0, n+1)
	for i := 2; i <= n; i++ {
		if b[i] {
			p = append(p, i)
			for j := 2 * i; j <= n; j += i {
				b[j] = false
			}
		}
	}

	return p
}

func primes1(n int) []int {
	if maxInt16 < n {
		panic(fmt.Sprintf("Primes: %d is the largest value supported", maxInt16))
	}

	// Seive of Eratosthenes
	b := make(map[int]struct{})
	for i := 2; i <= n; i++ {
		b[i] = struct{}{}
	}

	p := make([]int, 0, n+1)
	for i := 2; i <= n; i++ {
		if _, ok := b[i]; ok {
			p = append(p, i)
			for j := 2 * i; j <= n; j += i {
				delete(b, j)
			}
		}
	}

	return p
}

func primes2(n int) map[int]struct{} {
	if maxInt16 < n {
		panic(strconv.Itoa(maxInt16) + " is the largest value supported")
	}

	// Seive of Eratosthenes
	p := make(map[int]struct{})
	for i := 2; i <= n; i++ {
		p[i] = struct{}{}
	}

	for i := 2; i <= n; i++ {
		if _, ok := p[i]; ok {
			for j := 2 * i; j <= n; j += i {
				delete(p, j)
			}
		}
	}

	return p
}

func primes3(n int) []int {
	p := make([]int, 0, n)
	for a := 2; a <= n; a++ {
		var (
			sr  = gomath.Sqrt(float64(a))
			fsr = int(sr)
		)

		if float64(fsr) != sr {
			// * Squares are not prime
			// * Any divisor will be less than or equal to sqrt(a)
			var i int
			for ; i < len(p) && p[i] <= fsr && a%p[i] != 0; i++ {
			}

			if len(p) <= i || fsr < p[i] {
				p = append(p, a)
			}
		}
	}

	return p
}

// Prod returns the product of a list of values.
func Prod(values ...float64) float64 {
	p := values[0]
	for i := 1; i < len(values); i++ {
		p *= values[i]
	}

	return p
}

// Sgn returns the signum of a given number. That is, if x is
// negative, -1 is returned, if x is positive, 1 is returned, and 0
// is returned otherwise.
func Sgn(x float64) float64 {
	switch {
	case x < 0:
		return -1
	case 0 < x:
		return 1
	default:
		return 0
	}
}

// SgnInt ...
func SgnInt(n int) int {
	switch {
	case n < 0:
		return -1
	case 0 < n:
		return 1
	default:
		return 0
	}
}

// StDev returns the standard deviation of a list of values.
func StDev(mean float64, xs ...float64) float64 {
	return gomath.Sqrt(Var(mean, xs...))
}

// Sum returns the sum of a list of values.
func Sum(xs ...float64) float64 {
	var s float64
	for i := 0; i < len(xs); i++ {
		s += xs[i]
	}

	return s
}

// SumInts returns the sum of a list of values.
func SumInts(xs ...int) int {
	var s int
	for i := 0; i < len(xs); i++ {
		s += xs[i]
	}

	return s
}

// Totient returns phi(n) = n prod(1-1/p) for all primes p such that p|n.
func Totient(n int) int {
	var phi int = 1
	for p, k := range Factor(n) {
		phi *= PowInt(p, k-1) * (p - 1)
	}

	return phi
}

// Var returns the Var of a list of values.
func Var(mean float64, xs ...float64) float64 {
	var v float64
	for i := 0; i < len(xs); i++ {
		t := xs[i] - mean
		v += t * t
	}

	return v / float64(len(xs)-1)
}

// ------------
// EXPERIMENTAL
// ------------

// Ternary ...
func Ternary(condition bool, trueResult, falseResult float64) float64 {
	if condition {
		return trueResult
	}

	return falseResult
}

// Tuple ...
type Tuple []float64

// Twitch ...
func Twitch(v, d float64, t ...Tuple) float64 {
	for i := 0; i < len(t); i++ {
		if v == t[i][0] {
			return t[i][1]
		}
	}

	return d
}
