package matrix2

import (
	gomath "math"
	"sort"
	"strconv"
	"strings"

	"github.com/nathangreene3/math"
	vtr "github.com/nathangreene3/math/linalg/vector"
)

// Matrix ...
type Matrix struct {
	mat  []float64
	m, n int
}

// ----------------------------------------------------------
// Matrix constructors
// ----------------------------------------------------------

// Generator ...
type Generator func(i, j int) float64

// Identity ...
func Identity(n int) *Matrix {
	if n < 0 {
		panic("")
	}

	mat := make([]float64, n*n)
	for i, n1 := 0, n+1; i < len(mat); i += n1 {
		mat[i]++
	}

	return &Matrix{mat: mat, m: n, n: n}
}

// Gen ...
func Gen(m, n int, f Generator) *Matrix {
	if m < 0 || n < 0 {
		panic("")
	}

	mat := make([]float64, 0, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mat = append(mat, f(i, j))
		}
	}

	return &Matrix{mat: mat, m: m, n: n}
}

// New ...
func New(m, n int, mat ...float64) *Matrix {
	if m*n != len(mat) {
		panic("")
	}

	return &Matrix{mat: append(make([]float64, 0, len(mat)), mat...), m: m, n: n}
}

// Zeroes ...
func Zeroes(m, n int) *Matrix {
	if m < 0 || n < 0 {
		panic("")
	}

	return &Matrix{mat: make([]float64, m*n), m: m, n: n}
}

// ----------------------------------------------------------
// Matrix operations
// ----------------------------------------------------------

// Add ...
func (A *Matrix) Add(B *Matrix) {
	if len(A.mat) != len(B.mat) {
		panic("")
	}

	for i := 0; i < len(A.mat); i++ {
		A.mat[i] += B.mat[i]
	}
}

// Approx ...
func (A *Matrix) Approx(B *Matrix, tol float64) bool {
	if A.m != B.m || A.n != B.n {
		return false
	}

	for i := 0; i < len(A.mat); i++ {
		if !math.Approx(A.mat[i], B.mat[i], tol) {
			return false
		}
	}

	return true
}

// Col ...
func (A *Matrix) Col(i int) *Matrix {
	v := make([]float64, 0, A.m)
	for j := i; j < len(A.mat); j += A.n {
		v = append(v, A.mat[j])
	}

	return &Matrix{mat: v, m: A.m, n: 1}
}

// ColMatrix ...
func ColMatrix(v vtr.Vector) *Matrix {
	return &Matrix{mat: append(make([]float64, 0, len(v)), v...), m: len(v), n: 1}
}

// Compare ...
func (A *Matrix) Compare(B *Matrix) int {
	if len(A.mat) != len(B.mat) {
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

// CompareCols ...
func (A *Matrix) CompareCols(i, j int) int {
	for k := 0; k < A.m; k++ {
		switch {
		case A.mat[i*A.n+k] < A.mat[j*A.n+k]:
			return -1
		case A.mat[j*A.n+k] < A.mat[i*A.n+k]:
			return 1
		}
	}

	return 0
}

// CompareRows ...
func (A *Matrix) CompareRows(i, j int) int {
	for k := 0; k < A.n; k++ {
		switch {
		case A.mat[i*A.n+k] < A.mat[j*A.n+k]:
			return -1
		case A.mat[j*A.n+k] < A.mat[i*A.n+k]:
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

// Det ...
func (A *Matrix) Det() float64 {
	if A.m != A.n {
		panic("")
	}

	switch A.m {
	case 1:
		return A.mat[0]
	case 2:
		return A.mat[0]*A.mat[3] - A.mat[1]*A.mat[2]
	case 3:
		return A.mat[0]*(A.mat[4]*A.mat[8]-A.mat[5]*A.mat[7]) - A.mat[1]*(A.mat[3]*A.mat[8]-A.mat[5]*A.mat[6]) + A.mat[2]*(A.mat[3]*A.mat[7]-A.mat[4]*A.mat[6])
	default:
		// TODO
		return 0
	}
}

// Dims ...
func (A *Matrix) Dims() (int, int) {
	return A.m, A.n
}

// Equal ...
func (A *Matrix) Equal(B *Matrix) bool {
	if len(A.mat) != len(B.mat) {
		return false
	}

	for i := 0; i < len(A.mat); i++ {
		if A.mat[i] != B.mat[i] {
			return false
		}
	}

	return true
}

// Get ...
func (A *Matrix) Get(i, j int) float64 {
	return A.mat[i*A.n+j]
}

// Identity ...
func (A *Matrix) Identity() {
	for i := 0; i < A.m; i++ {
		for j, jmax := i+1, i+A.n; j < jmax; j += A.n {
			A.mat[j] = 0
		}
	}

	for i := 0; i < len(A.mat); i += A.n + 1 {
		A.mat[i] = 1
	}
}

// Inverse ...
func Inverse(A *Matrix) *Matrix {
	B := A.Copy()
	B.Inverse()
	return B
}

// Inverse ...
func (A *Matrix) Inverse() {
	B := A.Join(Identity(A.m))
	B.ref2()

	for k := 0; k < A.n; k++ {
		for i := B.m - 2; 0 <= i; i-- {
			A.mat[i*A.n+k] = B.mat[i*B.n+B.m]
			for j := i + 1; j < B.m; j++ {
				A.mat[i*A.n+k] -= B.mat[i*B.n+j] * A.mat[j*A.n+k]
			}
		}
	}
}

// Join ...
func (A *Matrix) Join(B *Matrix) *Matrix {
	mat := make([]float64, 0, len(A.mat)+len(B.mat))
	for i := 0; i < A.m; i++ {
		for j, jmax := i*A.n, (i+1)*A.n; j < jmax; j++ {
			mat = append(mat, A.mat[j])
		}

		for j, jmax := i*B.n, (i+1)*B.n; j < jmax; j++ {
			mat = append(mat, B.mat[j])
		}
	}

	return &Matrix{mat: mat, m: A.m, n: A.n + B.n}
}

// Len ...
func (A *Matrix) Len() int {
	return A.m
}

// Less ...
func (A *Matrix) Less(rowi, rowj int) bool {
	return 0 < A.CompareRows(rowi, rowj)
}

// Mult ...
func Mult(As ...*Matrix) *Matrix {
	switch n := len(As); n {
	case 0:
		return nil
	case 1:
		return As[0].Copy()
	case 2:
		return As[0].Mult(As[1])
	default:
		// Source: https://home.cse.ust.hk/~dekai/271/notes/L12/L12.pdf
		var (
			dims  = make([]int, 0, n+1) // Matrix As[i] has dimension dims[i]xdims[i+1], for 0 <= i < n
			cache = make([][]int, 0, n)
			order = make([][]int, 0, n)
		)

		for i := 0; i < n; i++ {
			dims = append(dims, As[i].m)
			cache = append(cache, make([]int, n))
			order = append(order, make([]int, n))
		}

		dims = append(dims, As[n-1].n)
		for h := 1; h < n; h++ {
			for i, imax := 0, n-h; i < imax; i++ {
				j := i + h
				cache[i][j] = gomath.MaxInt64
				for k := i; k < j; k++ {
					cost := cache[i][k] + cache[k+1][j] + dims[i]*dims[j+1]*dims[k+1]
					if cost < cache[i][j] {
						cache[i][j] = cost
						order[i][j] = k
					}
				}
			}
		}

		// ordMult returns the product of matrices by multiplying by
		// a given order. Initiate by calling on i = 0 and j = n-1.
		var ordMult func(order [][]int, i, j int, As ...*Matrix) *Matrix
		ordMult = func(order [][]int, i, j int, As ...*Matrix) *Matrix {
			if i < j {
				return ordMult(order, i, order[i][j], As...).Mult(ordMult(order, order[i][j]+1, j, As...))
			}

			return As[0]
		}

		return ordMult(order, 0, n-1, As...)
	}
}

// Mult ...
func (A *Matrix) Mult(B *Matrix) *Matrix {
	if A.n != B.m {
		panic("")
	}

	// TODO: Use Winograd's algorithm.
	mat := make([]float64, A.m*B.n)
	for i := 0; i < A.m; i++ {
		for j := 0; j < B.n; j++ {
			for k := 0; k < A.n; k++ {
				mat[i*B.n+j] += A.mat[i*A.n+k] * B.mat[k*B.n+j]
			}
		}
	}

	return &Matrix{mat: mat, m: A.m, n: B.n}
}

// Pow ...
func Pow(A *Matrix, p int) *Matrix {
	B := A.Copy()
	B.Pow(p)
	return B
}

// Pow ...
func (A *Matrix) Pow(p int) {
	switch {
	case A.m != A.n:
		panic("")
	case p < -1:
		A.Pow(-p)
		A.Inverse()
	case p == -1:
		A.Inverse()
	case p == 0:
		A.Identity()
	case p == 1:
	default:
		// Yacca's algorithm
		// TODO: Who was Yacca? Provide a source.
		P := A.Copy()
		B := Identity(A.m)
		for ; 0 < p; p >>= 1 {
			if p&1 == 1 {
				B = B.Mult(P)
			}

			P = P.Mult(P)
		}

		copy(A.mat, B.mat)
	}
}

// ref ...
func (A *Matrix) ref() {
	for i := 1; i < A.m; i++ {
		for j := i; j < A.m; j++ {
			if A.mat[j*(A.n+1)-1] != 0 {
				A.ScalDivRow(j, A.mat[j*(A.n+1)-1])
				A.ScalMultRow(j, A.mat[(j-1)*(A.n+1)])
				A.SubRows(j, i-1)
			}
		}
	}

	for i := A.m - 2; 0 <= i; i-- {
		for j := i; 0 <= j; j-- {
			if A.mat[j*(A.n+1)+1] != 0 {
				A.ScalDivRow(j, A.mat[j*(A.n+1)+1])
				A.ScalMultRow(j, A.mat[(j+1)*(A.n+1)])
				A.SubRows(j, i+1)
			}
		}
	}

	for i := 0; i < A.m; i++ {
		if A.mat[i*(A.n+1)] != 0 {
			A.ScalDivRow(i, A.mat[i*(A.n+1)])
		}
	}
}

// ref2 ...
func (A *Matrix) ref2() {
	if A.n < A.m {
		panic("invalid dimensions")
	}

	// ------------------------------------------------------------------------
	// Gaussian elimination (Algorithm 6.1)
	// Numerical Methods, 7th Ed.
	// Richard L. Burdent and J. Douglas Faires
	//
	// Backward substitution is performed in solving Ax = y and A^-1.
	// ------------------------------------------------------------------------

	for i, imax := 0, A.m-1; i < imax; i++ {
		p := A.m
		for j := i; j < A.m; j++ {
			if A.mat[j*A.n+i] != 0 {
				p = j
				break
			}
		}

		if p == A.m {
			// No unique solution
			return
		}

		if i != p {
			A.Swap(i, p)
		}

		for j := i + 1; j < A.m; j++ {
			A.addEqRow(j, 1, i, -A.mat[j*A.n+i]/A.mat[i*(A.n+1)])
		}
	}

	if A.mat[A.m*A.m-1] == 0 {
		// No unique solution
		return
	}
}

// RemoveDupRows ...
func (A *Matrix) RemoveDupRows() *Matrix {
	// areRowsMults ...
	areRowsMults := func(rowi, rowj int) bool {
		// For Ai to be a multiple of Aj, all column values must either both be
		// zero, or neither be zero. This finds the first column k such that Aik
		// and Ajk are both non-zero. Then it sets the ratio r or checks if ratios
		// are consistent.
		var r float64 // Ratio of each Aik to Ajk, assuming Aik, Ajk != 0
		for k := 0; k < A.n; k++ {
			switch Aik, Ajk := A.mat[rowi*A.n+k], A.mat[rowj*A.n+k]; {
			case Aik != 0:
				if Ajk != 0 {
					if 0 < r {
						if r != Aik/Ajk {
							// r not consistent
							return false
						}
					} else {
						// r should be set only once
						r = Aik / Ajk
					}
				} else {
					// Aik != 0, but Ajk == 0
					return false
				}
			case Ajk != 0:
				// Aik == 0, but Ajk != 0
				return false
			}
		}

		// r is consistant for all non-zero Aik, Ajk
		return true
	}

	// TODO
	var (
		mat  = append(make([]float64, 0, len(A.mat)), A.mat[:A.n]...)
		dups = make([]bool, A.m)
		m    = 1
	)

	for i, imax := 0, A.m; i < imax; i++ {
		if dups[i] {
			continue
		}

		for j := i + 1; j < A.m; j++ {
			switch {
			case dups[j]:
			case areRowsMults(i, j):
				dups[j] = true
			default:
				mat = append(mat, A.mat[j*A.n:(j+1)*A.n]...)
				m++
			}
		}
	}

	return &Matrix{mat: mat, m: m, n: A.n}
}

// Row ...
func (A *Matrix) Row(i int) *Matrix {
	v := make([]float64, 0, A.n)
	for j, jmax := i*A.n, i*A.n+A.n; j < jmax; j++ {
		v = append(v, A.mat[j])
	}

	return &Matrix{mat: v, m: 1, n: A.n}
}

// RowMatrix ...
func RowMatrix(v vtr.Vector) *Matrix {
	return &Matrix{mat: append(make([]float64, 0, len(v)), v...), m: 1, n: len(v)}
}

// ScalDivRow ...
func (A *Matrix) ScalDivRow(i int, a float64) {
	for j, jmax := i*A.n, (i+1)*A.n; j < jmax; j++ {
		A.mat[j] /= a
	}
}

// ScalMult ...
func (A *Matrix) ScalMult(a float64) {
	for i := 0; i < len(A.mat); i++ {
		A.mat[i] *= a
	}
}

// ScalMultRow ...
func (A *Matrix) ScalMultRow(i int, a float64) {
	for j, jmax := i*A.n, (i+1)*A.n; j < jmax; j++ {
		A.mat[j] *= a
	}
}

// Set ...
func (A *Matrix) Set(i, j int, v float64) {
	A.mat[i*A.n+j] = v
}

// Solve ...
func (A *Matrix) Solve(y *Vector) *Vector {
	B := A.Join(y.ColMatrix())
	B.ref2()

	vec := append(make([]float64, B.m-1, B.m), B.mat[B.m*B.n-1]/B.mat[B.m*B.n-2])
	for i := B.m - 2; 0 <= i; i-- {
		vec[i] = B.mat[i*B.n+B.m]
		for j := i + 1; j < B.m; j++ {
			vec[i] -= B.mat[i*B.n+j] * vec[j]
		}

		vec[i] /= B.mat[i*(B.n+1)]
	}

	return &Vector{vec: vec}
}

// Sort ...
func (A *Matrix) Sort() {
	sort.Sort(A)
}

// String ...
func (A *Matrix) String() string {
	var sb strings.Builder
	sb.Grow(len(A.mat) + 2*(A.m+1))
	sb.WriteByte('[')
	for i := 0; i < len(A.mat); i += A.n {
		sb.WriteByte('[')
		sb.WriteString(strconv.FormatFloat(A.mat[i], 'f', -1, 64))
		for j, jmax := i+1, i+A.n; j < jmax; j++ {
			sb.WriteString(" " + strconv.FormatFloat(A.mat[j], 'f', -1, 64))
		}

		sb.WriteByte(']')
	}

	sb.WriteByte(']')
	return sb.String()
}

// Sub ...
func (A *Matrix) Sub(B *Matrix) {
	if len(A.mat) != len(B.mat) {
		panic("")
	}

	for i, imax := 0, A.m*A.n; i < imax; i++ {
		A.mat[i] -= B.mat[i]
	}
}

// SubRows ...
func (A *Matrix) SubRows(fromi, subj int) {
	for k := 0; k < A.n; k++ {
		A.mat[fromi*A.n+k] -= A.mat[subj*A.n+k]
	}
}

// addEqRow updates row i as A[i] := aA[i] + bA[j]. A[j] is unchanged.
func (A *Matrix) addEqRow(rowi int, a float64, rowj int, b float64) {
	for k := 0; k < A.n; k++ {
		A.mat[rowi*A.n+k] *= a
		A.mat[rowi*A.n+k] += b * A.mat[rowj*A.n+k]
	}
}

// Swap ...
func (A *Matrix) Swap(rowi, rowj int) {
	for k := 0; k < A.n; k++ {
		A.mat[rowi*A.n+k], A.mat[rowj*A.n+k] = A.mat[rowj*A.n+k], A.mat[rowi*A.n+k]
	}
}

// SwapCols ...
func (A *Matrix) SwapCols(i, j int) {
	for k := 0; k < A.m; k++ {
		A.mat[k*A.n+i], A.mat[k*A.n+j] = A.mat[k*A.n+j], A.mat[k*A.n+i]
	}
}

// Trace ...
func (A *Matrix) Trace(mainDiag bool) float64 {
	if A.m != A.n {
		panic("")
	}

	var s float64
	if mainDiag {
		for i, n := 0, A.n+1; i < len(A.mat); i += n {
			s += A.mat[i]
		}
	} else {
		n := A.n - 1
		for i, imax := n, len(A.mat)-1; i < imax; i += n {
			s += A.mat[i]
		}
	}

	return s
}

// Trans ...
func (A *Matrix) Trans() *Matrix {
	// TODO
	mat := make([]float64, 0, A.m*A.n)
	return &Matrix{mat: mat, m: A.n, n: A.m}
}

// Wid ...
func (A *Matrix) Wid() int {
	return A.n
}

// Vector ...
func (A *Matrix) Vector() *Vector {
	if A.m != len(A.mat) && A.n != len(A.mat) {
		panic("invalid dimensions")
	}

	return &Vector{vec: append(make([]float64, 0, len(A.mat)), A.mat...)}
}
