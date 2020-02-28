package math

// Zn = {0, 1, ..., n-1},  n > 0
//    = {n+1, ..., -1, 0}, n < 0 (my extension to the definition)

// Addition modulo n: a mod n = r such that a = kn + r for some integer k and r
// in [0,n) if n > 0 and (n,0] if n < 0.

// The values "carry" and "borrow" refer to k and are useful in addition and
// subtraction over the external direct product of several sets of Zn. See J.A.
// Gallian's Contemporary Abstract Algebra, 6th Ed., chapters 1, 2, and 8.

// A possible property: if x >= 0 and n > 0, then x mod -n = -(-x mod n).

// AddWithCarry returns r = (a+b) mod n with the carried amount k such that
// a+b = kn+r for a given modulus n != 0. If n > 0, then 0 <= r < n. If n < 0,
// then n < r <= 0.
func AddWithCarry(a, b, modulus int) (int, int) {
	var (
		ka, ra = EuclidsCoeffs(a, modulus)
		kb, rb = EuclidsCoeffs(b, modulus)
		k, r   = EuclidsCoeffs(ra+rb, modulus)
	)

	return r, ka + kb + k
}

// EuclidsCoeffs returns (k,r) such that a = kn + r for a given modulus n != 0.
// If n > 0, then 0 <= r < n. Otherwise, n < r <= 0. In either case, k = (x-r)/n.
func EuclidsCoeffs(a, modulus int) (int, int) {
	r := Mod(a, modulus)
	return (a - r) / modulus, r
}

// Mod returns r such that a = kn + r for some integer k given modulus n != 0. If
// n > 0, then 0 <= r < n. If n < 0, then n < r <= 0. Note this is NOT equivalent
// to a % n as the % operator returns the remainder of a/n.
func Mod(a, modulus int) int {
	return (a%modulus + modulus) % modulus
}

// SubtractWithBorrow returns r = (a-b) mod n with the carried amount k such that
// a-b = -kn+r for a given modulus n != 0. If n > 0, then 0 <= r < n. If n < 0,
// then n < r <= 0.
func SubtractWithBorrow(a, b, modulus int) (int, int) {
	var (
		ka, ra = EuclidsCoeffs(a, modulus)
		kb, rb = EuclidsCoeffs(b, modulus)
		k, r   = EuclidsCoeffs(ra-rb, modulus)
	)

	return r, -ka + kb - k
}
