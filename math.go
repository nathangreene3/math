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

// fibonacci returns the nth fibonacci term, where the 0th and 1st terms are 1 and the nth term is the sum of the previous two terms.
func fibonacci(n int) int {
	a0, a1 := 1, 1
	var t int
	for ; 1 < n; n-- {
		t = a0 + a1
		a0 = a1
		a1 = t
	}

	return a1
}

// choose returns n-choose-k, or n!/(k!(n-k)!).
func choose(n, k int) int {
	return pascal(n)[n][k]
}

// pascal returns pascal's triangle.
func pascal(n int) [][]int {
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
