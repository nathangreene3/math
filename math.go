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
