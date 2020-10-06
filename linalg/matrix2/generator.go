package matrix2

// Generator ...
type Generator func(i, j int) float64

// Delta is a generator that produces an identity matrix.
func Delta(i, j int) float64 {
	if i == j {
		return 1
	}

	return 0
}
