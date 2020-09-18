package sequence

// Sequence is a function mapping from the integers to the reals.
type Sequence func(int) float64

// PartialProds ...
func PartialProds(a Sequence, i, n int) map[int]float64 {
	var (
		p float64 = 1
		m         = make(map[int]float64)
	)

	for ; i <= n; i++ {
		p *= a(i)
		m[i] = p
	}

	return m
}

// PartialSums ...
func PartialSums(a Sequence, i, n int) map[int]float64 {
	var (
		s float64
		m = make(map[int]float64)
	)

	for ; i <= n; i++ {
		s += a(i)
		m[i] = s
	}

	return m
}

// Prod returns the product of a sequence over [i,n].
func Prod(a Sequence, i, n int) float64 {
	var p float64 = 1
	for ; i <= n; i++ {
		p *= a(i)
	}

	return p
}

// Sum of a sequence over [i,n].
func Sum(a Sequence, i, n int) float64 {
	var s float64
	for ; i <= n; i++ {
		s += a(i)
	}

	return s
}

// ToMap ...
func ToMap(a Sequence, i, n int) map[int]float64 {
	m := make(map[int]float64)
	for ; i <= n; i++ {
		m[i] = a(i)
	}

	return m
}
