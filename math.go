package math

// GCD returns the largest divisor of both a and b.
func GCD(a, b int) int {
	if a < 0 || b < 0 {
		panic("invalid sign")
	}

	var c int
	if a < b {
		c = a
		a = b
		b = c
	}

	for 0 < b {
		c = a % b
		a = b
		b = c
	}

	return a
}

// LCM returns the largest multiple of a and b divisible by a and b.
func LCM(a, b int) int {
	if a < 1 || b < 1 {
		panic("a and b must be positive")
	}

	m, n := a, b
	for m != n {
		for m < n {
			m += a
		}

		for n < m {
			n += b
		}
	}

	return m
}

// Fibonacci returns the nth Fibonacci term, where the 0th and 1st terms are 1 and the nth term is the sum of the previous two terms.
func Fibonacci(n int) int {
	a0, a1 := 1, 1
	var t int
	for ; 1 < n; n-- {
		t = a0 + a1
		a0 = a1
		a1 = t
	}

	return a1
}

func fibDyn(n int) int {
	if n < 2 {
		return 1
	}

	cache := map[int]int{0: 1, 1: 1}
	for i := 2; i <= n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}

	return cache[n]
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

// Pascal returns Pascal's triangle, consisting of n levels. The (n,k) entry is the value n!/(k!(n-k)!).
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
