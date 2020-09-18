package linalg

// Matrix ...
type Matrix struct {
	mat  []float64
	m, n int
}

// Generator ...
type Generator func(i, j int) float64

// Identity ...
func Identity(n int) *Matrix {
	mat := make([]float64, n*n)
	for i := 0; i < len(mat); i += n {
		mat[i]++
	}

	return &Matrix{mat: mat, m: n, n: n}
}

// New ...
func New(m, n int, f Generator) *Matrix {
	A := Matrix{
		mat: make([]float64, 0, m*n),
		m:   m,
		n:   n,
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			A.mat = append(A.mat, f(i, j))
		}
	}

	return &A
}

// Zeroes ...
func Zeroes(m, n int) *Matrix {
	return &Matrix{mat: make([]float64, m*n), m: m, n: n}
}

// Add ...
func (A *Matrix) Add(B *Matrix) {
	if len(A.mat) != len(B.mat) || A.m != B.m || A.n != B.n {
		panic("")
	}

	for i, imax := 0, A.m*A.n; i < imax; i++ {
		A.mat[i] += B.mat[i]
	}
}

// Compare ...
func (A *Matrix) Compare(B *Matrix) int {
	if len(A.mat) != len(B.mat) || A.m != B.m || A.n != B.n {
		panic("")
	}

	for i := 0; i < len(A.mat); i++ {
		switch {
		case A.mat[i] < B.mat[i]:
			return -1
		case B.mat[i] < A.mat[i]:
			return 1
		}
	}

	return 0
}

// Copy ...
func (A *Matrix) Copy() *Matrix {
	mat := make([]float64, 0, len(A.mat))
	for i := 0; i < len(A.mat); i++ {
		mat = append(mat, A.mat[i])
	}

	return &Matrix{mat: mat, m: A.m, n: A.n}
}

// Dims ...
func (A *Matrix) Dims() (int, int) {
	return A.m, A.n
}

// Equal ...
func (A *Matrix) Equal(B *Matrix) bool {
	return A.Compare(B) == 0
}

// Get ...
func (A *Matrix) Get(i, j int) float64 {
	return A.mat[i*A.m+j]
}

// Join ...
func (A *Matrix) Join(B *Matrix) *Matrix {
	mat := make([]float64, 0, len(A.mat)+len(B.mat))
	for i := 0; i < A.m; i++ {
		for j, jmax := i*A.m, i*A.m+A.n; j < jmax; j++ {
			mat = append(mat, A.mat[j])
		}

		for j, jmax := i*B.m, i*B.m+B.n; j < jmax; j++ {
			mat = append(mat, B.mat[j])
		}
	}

	return &Matrix{mat: mat, m: A.m, n: A.n + B.n}
}

// Multiply ...
func (A *Matrix) Multiply(a float64) {
	for i := 0; i < len(A.mat); i++ {
		A.mat[i] *= a
	}
}

// Set ...
func (A *Matrix) Set(i, j int, v float64) {
	A.mat[i*A.m+j] = v
}

// Subtract ...
func (A *Matrix) Subtract(B *Matrix) {
	if len(A.mat) != len(B.mat) || A.m != B.m || A.n != B.n {
		panic("")
	}

	for i, imax := 0, A.m*A.n; i < imax; i++ {
		A.mat[i] -= B.mat[i]
	}
}
