package math

import (
	gomath "math"
)

// Approx returns true if |x-y| <= prec.
func Approx(x, y, prec float64) bool {
	if prec < 0 || 1 < prec {
		panic("precision must be on range [0,1]")
	}

	return gomath.Abs(x-y) <= prec
}

// Factor returns a map of factors to the number of times they divide
// an integer n. That is, for each key-value pair (k,v), k divides n
// v times. Each key will be a prime divisor, which means k will be
// at least two. An integer is prime if its only Factor is itself
// (and 1, which is called the empty prime).
func Factor(n int) map[int]int {
	if n < 1 {
		panic("cannot factor non-positive integer")
	}

	// Theorem: If p is prime greater than 3, then p = 6k-1 or 6k+1.
	// 1. Determine how many times 2 and 3 divide n, if at all.
	// 2. Check all integers 6k-1 and 6k+1 less than sqrt(n) to help eliminate multiples of 2 and 3. This speeds up factoring three times over.
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

// IsPrime indicates if n is prime.
func IsPrime(n int) bool {
	if n < 2 {
		return false // Prevents panic on non-positives and 1 isn't prime anyway
	}

	_, ok := Factor(n)[n]
	return ok
}

// Base converts a number into its base representation.
func Base(n, b int) []int {
	if n < 0 {
		panic("number must be non-negative")
	}

	if b < 2 {
		panic("base must be greater than one")
	}

	var (
		remainders = make([]int, 0, 64)
		d          int
	)

	for ; 0 < n; n = d {
		d = n / b
		remainders = append(remainders, n-d*b)
	}

	cpy := make([]int, len(remainders))
	copy(cpy, remainders)
	return cpy
}

// GCD returns the largest divisor of both a and b. If GCD(a,b) == 1,
// then a and b are relatively prime.
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

// Fibonacci returns the nth Fibonacci term, where the 0th and 1st
// terms are 1 and the nth term is the sum of the previous two terms.
func Fibonacci(n int) int {
	a0, a1 := 1, 1
	for ; 1 < n; n-- {
		a0, a1 = a1, a0+a1
	}

	return a1
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

// Choose returns n-Choose-k, or n!/(k!(n-k)!).
func Choose(n, k int) int {
	return Pascal(n + 1)[n][k]
}

// Pascal returns Pascal's triangle, consisting of n levels. The
// (n,k)th entry is the value n!/(k!(n-k)!).
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
