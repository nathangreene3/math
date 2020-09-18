package linalg

// Matrix ...
type Matrix interface {
	Get(i, j int) float64
	Set(i, j int, v float64)
}

// LongMatrix ...
type LongMatrix struct {
	matrix []float64
	m, n   int
}

// Generator ...
type Generator func(i, j int) float64

// New ...
func New(m, n int, f Generator) *LongMatrix {
	A := LongMatrix{
		matrix: make([]float64, 0, m*n),
		m:      m,
		n:      n,
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			A.matrix = append(A.matrix, f(i, j))
		}
	}

	return &A
}
