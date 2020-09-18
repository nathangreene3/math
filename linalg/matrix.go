package linalg

// Matrix ...
type Matrix interface {
	MatrixReader
	MatrixWriter
}

// MatrixReader ...
type MatrixReader interface {
	Get(i, j int) float64
}

// MatrixWriter ...
type MatrixWriter interface{ Set(i, j int) float64 }

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

// Add ...
func (A *LongMatrix) Add(B *LongMatrix) {
	switch {
	case len(A.matrix) != len(B.matrix):
	case A.m != B.m:
	case A.n != B.n:
	}

	for i := 0; i < A.m*A.n; i++ {
		A.matrix[i] += B.matrix[i]
	}
}

// Get ...
func (A *LongMatrix) Get(i, j int) float64 {
	return A.matrix[i*A.m+j]
}

// Set ...
func (A *LongMatrix) Set(i, j int, v float64) {
	A.matrix[i*A.m+j] = v
}
