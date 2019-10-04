package matrix

import (
	gomath "math"
	"sort"
	"strings"

	"github.com/nathangreene3/math/linalg/vector"
)

// ------------------------------------------------------------------------------
// RESOURCES
// ------------------------------------------------------------------------------
// Most methods defined here are taken from or are inspired by Linear Algebra,
// 3rd Ed., by Stephen H. Friedberg, Arnold J. Insel, and Lawrence E. Spence.
// Any page references in comments are in reference to this source.
// ------------------------------------------------------------------------------

// Matrix is a set of vectors.
type Matrix []vector.Vector

// Generator is a function defining the (i,j)th entry of a matrix.
type Generator func(i, j int) float64

// ------------------------------------------------------------------------------
// MATRIX CONSTRUCTORS
// ------------------------------------------------------------------------------

// New generates an m-by-n matrix with entries defined by a generating function
// f.
func New(m, n int, f Generator) Matrix {
	A := make(Matrix, 0, m)
	for i := 0; i < m; i++ {
		A = append(A, vector.New(n, func(j int) float64 { return f(i, j) }))
	}

	return A
}

// Empty returns an m-by-n matrix with zeroes for all entries.
func Empty(m, n int) Matrix {
	return New(m, n, func(i, j int) float64 { return 0 })
}

// Identity returns the m-by-n identity matrix.
func Identity(m, n int) Matrix {
	// TODO: Determine if this should this panic when m,n = 0.

	// Note: this generating function is called the Kronecker delta
	// and is typically denoted as d(i,j).
	f := func(i, j int) float64 {
		if i == j {
			return 1
		}
		return 0
	}

	return New(m, n, f)
}

// ------------------------------------------------------------------------------
// OPERATIONS ON MATRICES
// ------------------------------------------------------------------------------
// In general, A.F(B) updates A and F(A,B) returns a new matrix.
// ------------------------------------------------------------------------------

// Add returns the sum of two matrices.
func Add(A, B Matrix) Matrix {
	C := A.Copy()
	C.Add(B)
	return C
}

// Add B to A.
func (A Matrix) Add(B Matrix) {
	ma, na := A.Dimensions()
	mb, nb := B.Dimensions()
	if ma != mb || na != nb {
		panic("matrices must have the same number of rows and columns")
	}

	for i := 0; i < ma; i++ {
		for j := 0; j < na; j++ {
			A[i][j] += B[i][j]
		}
	}
}

// AddColToCol returns A with column j to column i.
func AddColToCol(A Matrix, i, j int) Matrix {
	m, n := A.Dimensions()
	f := func(a, b int) float64 {
		if b == i {
			return A[a][b] + A[a][j]
		}
		return A[a][b]
	}

	return New(m, n, f)
}

// AddColToCol adds column j to column i.
func (A Matrix) AddColToCol(i, j int) {
	for _, a := range A {
		a[i] += a[j]
	}
}

// AddRowToRow returns A with row j added to row i.
func AddRowToRow(A Matrix, i, j int) Matrix {
	B := A.Copy()
	B.AddRowToRow(i, j)
	return B
}

// AddRowToRow adds row j to row i.
func (A Matrix) AddRowToRow(i, j int) {
	A[i].Add(A[j])
}

// AppendColumn returns a matrix that is the joining of a given matrix with a
// column Vector.
func (A Matrix) AppendColumn(x vector.Vector) Matrix {
	return A.Join(ColumnMatrix(x))
}

// AppendRow returns a matrix that is the joining of a given matrix with a row
// vector.
func (A Matrix) AppendRow(x vector.Vector) Matrix {
	if _, n := A.Dimensions(); n != len(x) {
		panic("matrix columns must be equal to vector dimensions")
	}

	return append(A, x)
}

// Approx returns true if A approximates B for a given precision on the range
// [0,1].
func (A Matrix) Approx(B Matrix, prec float64) bool {
	ma, na := A.Dimensions()
	mb, nb := B.Dimensions()
	if ma != mb || na != nb {
		return false
	}

	for i := 0; i < ma; i++ {
		if !A[i].Approx(B[i], prec) {
			return false
		}
	}

	return true
}

// At returns A[i][j].
func (A Matrix) At(i, j int) float64 {
	return A[i][j]
}

// ColumnMatrix converts a vector v to an n-by-1 matrix.
func ColumnMatrix(v vector.Vector) Matrix {
	return New(len(v), 1, func(i, j int) float64 { return v[i] })
}

// Compare returns -1, 0, 1 indicating A precedes, is equal to, or follows B.
func (A Matrix) Compare(B Matrix) int {
	ma, na := A.Dimensions()
	mb, nb := B.Dimensions()
	if ma != mb || na != nb {
		panic("dimension mismatch")
	}

	for i := 0; i < ma; i++ {
		for j := 0; j < na; j++ {
			switch {
			case A[i][j] < B[i][j]:
				return -1
			case B[i][j] < A[i][j]:
				return 1
			}
		}
	}

	return 0
}

// Copy returns a deep copied matrix.
func (A Matrix) Copy() Matrix {
	m, n := A.Dimensions()
	return New(m, n, func(i, j int) float64 { return A[i][j] })
}

// Cost returns the number of operations to compute AB.
func (A Matrix) Cost(B Matrix) int {
	ma, na := A.Dimensions()
	mb, nb := B.Dimensions()
	if na != mb {
		panic("dimension mismatch")
	}

	return ma * na * nb
}

// Determinant returns the Determinant of a square matrix.
func (A Matrix) Determinant() float64 {
	// TODO //
	m, n := A.Dimensions()
	if m == 0 || n == 0 {
		panic("cannot take determinant of empty matrix")
	}

	if m != n {
		panic("cannot take determinant of a non-square matrix")
	}

	switch m {
	case 1:
		return A[0][0]
	case 2:
		return A[0][0]*A[1][1] - A[0][1]*A[1][0]
	case 3:
		return A[0][0]*(A[1][2]*A[2][2]-A[1][2]*A[2][1]) - A[0][1]*(A[1][0]*A[2][2]-A[1][2]*A[2][0]) + A[0][2]*(A[1][0]*A[2][1]-A[1][1]*A[2][0])
	}

	// This is probably terrible...
	var (
		B   = A.RemoveRow(0)
		sgn = 1.0
		det float64
	)

	for i := 0; 0 < m; i++ {
		det += sgn * A[0][i] * B.RemoveColumn(i).Determinant()
		sgn *= -1
	}

	return det
}

// Dimensions returns the Dimensions (number of rows, number of columns) of a
// matrix.
func (A Matrix) Dimensions() (int, int) {
	m, n := len(A), len(A[0])
	for _, r := range A {
		if n != len(r) {
			panic("inconsistent matrix dimensions")
		}
	}

	return m, n
}

// Equals returns true if two matrices are equal in dimension and for each
// entry. Otherwise, it returns false.
func (A Matrix) Equals(B Matrix) bool {
	return A.Compare(B) == 0
}

// Join returns a matrix that is the joining of two given matrices.
func (A Matrix) Join(B Matrix) Matrix {
	ma, na := A.Dimensions()
	mb, nb := B.Dimensions()
	if ma != mb {
		panic("matrices must have equal number of rows")
	}

	f := func(i, j int) float64 {
		if j < na {
			return A[i][j]
		}
		return B[i][j-na]
	}

	return New(ma, na+nb, f)
}

// multiply returns AB. TODO: Correct the implementation of Winograd's algorithm.
func (A Matrix) multiply(B Matrix) Matrix {
	ma, na := A.Dimensions()
	mb, nb := B.Dimensions()
	if na != mb {
		panic("A and B are of incompatible dimensions")
	}

	f := func(i, j int) float64 {
		var v float64
		for k := 0; k < na; k++ {
			v += A[i][k] * B[k][j]
		}

		return v
	}

	return New(ma, nb, f)

	/*
		// Winograd's algorithm
		// Source: Analysis of Algorithms, 2nd Ed., by Jeffrey J.
		// McConnell, pg 139-140
		var (
			d      = na >> 1
			rfacts = make([]float64, 0, ma)
			cfacts = make([]float64, 0, nb)
		)

		for i := 0; i < ma; i++ {
			rfacts = append(rfacts, A[i][0]*A[i][1])
			for j := 1; j < d; j++ {
				rfacts[i] += A[i][j<<1] * A[i][j<<1+1]
			}
		}

		for i := 0; i < nb; i++ {
			cfacts = append(cfacts, B[0][i]*B[1][i])
			for j := 1; j < d; j++ {
				cfacts[i] += B[j<<1][i] * B[j<<1+1][i]
			}
		}

		C := Empty(ma, nb)
		for i := 0; i < ma; i++ {
			for j := 0; j < nb; j++ {
				C[i][j] = -rfacts[i] - cfacts[j]
				for k := 0; k < d; k++ {
					C[i][j] += (A[i][k<<1] + B[k<<1+1][j]) * (A[i][k<<1+1] + B[2*k][j])
				}
			}
		}

		if na%1 == 1 {
			for i := 0; i < ma; i++ {
				for j := 0; j < nb; j++ {
					C[i][j] += A[i][na-1] * B[na-1][j]
				}
			}
		}

		return C
	*/
}

// Multiply several matrices.
func Multiply(As ...Matrix) Matrix {
	// Source: https://home.cse.ust.hk/~dekai/271/notes/L12/L12.pdf

	n := len(As)
	switch n {
	case 0:
		return nil
	case 1:
		return As[0]
	case 2:
		return As[0].multiply(As[1])
	}

	var (
		dims    = make([]int, 0, n+1) // Matrix As[i] has dimension dims[i]xdims[i+1], for 0 <= i < n
		cache   = make([][]int, 0, n)
		order   = make([][]int, 0, n)
		j, cost int
	)

	for i := 0; i < n; i++ {
		dims = append(dims, len(As[i]))
		cache = append(cache, make([]int, n))
		order = append(order, make([]int, n))
	}

	dims = append(dims, len(As[n-1][0]))
	for h := 1; h < n; h++ {
		for i := 0; i < n-h; i++ {
			j = i + h
			cache[i][j] = gomath.MaxInt64
			for k := i; k < j; k++ {
				cost = cache[i][k] + cache[k+1][j] + dims[i]*dims[k+1]*dims[j+1]
				if cost < cache[i][j] {
					cache[i][j] = cost
					order[i][j] = k
				}
			}
		}
	}

	return multiplyByOrder(order, 0, n-1, As...)
}

// multiplyByOrder returns the product of matrices by multiplying by a given
// order. Initiate by calling on i = 0 and j = n-1.
func multiplyByOrder(s [][]int, i, j int, As ...Matrix) Matrix {
	if i < j {
		return multiplyByOrder(s, i, s[i][j], As...).multiply(multiplyByOrder(s, s[i][j]+1, j, As...))
	}

	return As[i]
}

// MultiplyColumn returns A with the ith col multiplied by a.
func MultiplyColumn(A Matrix, i int, a float64) Matrix {
	m, n := A.Dimensions()
	f := func(j, k int) float64 {
		if i == k {
			return a * A[j][k]
		}
		return A[j][k]
	}

	return New(m, n, f)
}

// MultiplyColumn by a.
func (A Matrix) MultiplyColumn(i int, a float64) {
	for j := range A {
		A[j][i] *= a
	}
}

// MultiplyRow returns A with the ith row multiplied by a.
func MultiplyRow(A Matrix, i int, a float64) Matrix {
	m, n := A.Dimensions()
	f := func(j, k int) float64 {
		if i == j {
			return a * A[j][k]
		}
		return A[j][k]
	}

	return New(m, n, f)
}

// MultiplyRow by a.
func (A Matrix) MultiplyRow(i int, a float64) {
	A[i].Multiply(a)
}

// Pow returns A^p, for square matrix A and -1 <= p. If p = -1, the inverse is
// returned. All other negative values for p will panic.
func Pow(A Matrix, p int) Matrix {
	m, n := A.Dimensions()
	switch {
	case m != n:
		panic("matrix must be square")
	case p < -1:
		panic("power must be non-negative, except for -1")
	case p == -1:
		return A.Inverse()
	}

	// Yacca's method
	B := Identity(m, n)
	C := A.Copy()
	for ; 0 < p; p >>= 1 {
		if p&1 == 1 {
			B = Multiply(B, C)
		}

		C = Multiply(C, C)
	}

	return B
}

// ScalarDivide returns (1/a)A.
func ScalarDivide(a float64, A Matrix) Matrix {
	return ScalarMultiply(1.0/a, A)
}

// ScalarDivide A by a.
func (A Matrix) ScalarDivide(a float64) {
	A.ScalarMultiply(1.0 / a)
}

// ScalarMultiply returns aA.
func ScalarMultiply(a float64, A Matrix) Matrix {
	m, n := A.Dimensions()
	return New(m, n, func(i, j int) float64 { return a * A[i][j] })
}

// ScalarMultiply A by a.
func (A Matrix) ScalarMultiply(a float64) {
	for _, r := range A {
		r.Multiply(a)
	}
}

// Subtract returns A-B.
func Subtract(A, B Matrix) Matrix {
	C := A.Copy()
	C.Subtract(B)
	return C
}

// Subtract B from A.
func (A Matrix) Subtract(B Matrix) {
	ma, na := A.Dimensions()
	mb, nb := B.Dimensions()
	if ma != mb || na != nb {
		panic("matrices must have the same number of rows and columns")
	}

	for i := 0; i < ma; i++ {
		for j := 0; j < na; j++ {
			A[i][j] -= B[i][j]
		}
	}
}

// SubtractRowFromRow returns A with row j subtracted from row i.
func SubtractRowFromRow(A Matrix, i, j int) Matrix {
	m, n := A.Dimensions()
	f := func(a, b int) float64 {
		if a == i {
			return A[a][b] - A[j][b]
		}
		return A[a][b]
	}

	return New(m, n, f)
}

// SubtractRowFromRow subtracts row j from row i.
func (A Matrix) SubtractRowFromRow(i, j int) {
	A[i].Subtract(A[j])
}

// SwapCols returns A with columns i and j swapped.
func SwapCols(A Matrix, i, j int) Matrix {
	m, n := A.Dimensions()
	f := func(a, b int) float64 {
		switch b {
		case i:
			return A[a][j]
		case j:
			return A[a][i]
		default:
			return A[a][b]
		}
	}

	return New(m, n, f)
}

// SwapCols swaps two columns.
func (A Matrix) SwapCols(i, j int) {
	for _, a := range A {
		a[i], a[j] = a[j], a[i]
	}
}

// SwapRows returns A with rows i and j swapped.
func SwapRows(A Matrix, i, j int) Matrix {
	m, n := A.Dimensions()
	f := func(a, b int) float64 {
		switch a {
		case i:
			return A[j][b]
		case j:
			return A[i][b]
		default:
			return A[a][b]
		}
	}

	return New(m, n, f)
}

// SwapRows swaps two rows.
func (A Matrix) SwapRows(i, j int) {
	A[i], A[j] = A[j], A[i]
}

// Solve Ax=y for x.
func (A Matrix) Solve(y vector.Vector) vector.Vector {
	B := A.Join(ColumnMatrix(y)).RemoveMultiples()
	m, n := B.Dimensions()
	for i := 1; i < m; i++ {
		for j := i; j < m; j++ {
			if B[j][j-1] != 0 {
				B[j].Divide(B[j][j-1])
				B[j].Multiply(B[j-1][j-1])
				B[j].Subtract(B[i-1])
			}
		}
	}

	for i := m - 2; 0 <= i; i-- {
		for j := i; 0 <= j; j-- {
			if B[j][j+1] != 0 {
				B[j].Divide(B[j][j+1])
				B[j].Multiply(B[j+1][j+1])
				B[j].Subtract(B[i+1])
			}
		}
	}

	for i, b := range B {
		if b[i] != 0 {
			b.Divide(b[i])
		}
	}

	return vector.New(m, func(i int) float64 { return B[i][n-1] }) // x = A^-1*b
}

// Inverse of a square matrix. Caution: not all matrices, even square ones, are
// guarenteed to be invertible.
func (A Matrix) Inverse() Matrix {
	m, n := A.Dimensions()
	if m != n {
		panic("invalid dimensions")
	}

	B := A.Join(Identity(m, m))
	for i := 1; i < m; i++ {
		for j := i; j < m; j++ {
			if B[j][j-1] != 0 {
				B[j].Divide(B[j][j-1])
				B[j].Multiply(B[j-1][j-1])
				B[j].Subtract(B[i-1])
			}
		}
	}

	for i := m - 2; 0 <= i; i-- {
		for j := i; 0 <= j; j-- {
			if B[j][j+1] != 0 {
				B[j].Divide(B[j][j+1])
				B[j].Multiply(B[j+1][j+1])
				B[j].Subtract(B[i+1])
			}
		}
	}

	for i, b := range B {
		if b[i] != 0 {
			b.Divide(b[i])
		}
	}

	return New(m, m, func(i, j int) float64 { return B[i+m][j] })
}

// RemoveColumn returns a copy of a matrix with column i removed.
func (A Matrix) RemoveColumn(i int) Matrix {
	m, n := A.Dimensions()
	f := func(a, b int) float64 {
		if b < i {
			return A[a][b]
		}
		return A[a][b+1]
	}

	return New(m, n-1, f)
}

// RemoveMultiples returns a sorted copy of a matrix with all row multiples
// removed.
func (A Matrix) RemoveMultiples() Matrix {
	B := A.Copy()
	B.Sort()

	m, _ := B.Dimensions()
	for i := 0; i+1 < m; i++ {
		for j := i + 1; j < m; j++ {
			if B[i].IsMultipleOf(B[j]) {
				B = B.RemoveRow(i)
			}
		}
	}

	return B
}

// RemoveRow returns A with row i removed.
func (A Matrix) RemoveRow(i int) Matrix {
	m, n := A.Dimensions()
	f := func(a, b int) float64 {
		if a < i {
			return A[a][b]
		}
		return A[a+1][b]
	}

	return New(m, n, f)
}

// RowMatrix converts a vector v to a 1-by-n matrix.
func RowMatrix(v vector.Vector) Matrix {
	return New(1, len(v), func(i, j int) float64 { return v[j] })
}

// Sort A such that the largest leading indices are at the top (index 0 is the
// top).
func (A Matrix) Sort() {
	sort.Slice(A, func(i, j int) bool { return 0 < A[i].Compare(A[j]) })
}

// String returns a formatted string representation of a matrix. TODO: Determine
// if this is needed.
func (A Matrix) String() string {
	var (
		m, n = A.Dimensions()
		sb   strings.Builder
	)

	sb.Grow(2*m*(n+1) + 1)
	sb.WriteByte(byte('['))
	sb.WriteString(A[0].String())
	for i := 1; i < len(A); i++ {
		sb.WriteByte(',')
		sb.WriteString(A[i].String())
	}

	sb.WriteByte(']')
	return sb.String()
}

// tildeA TODO: Rename this.
func (A Matrix) tildeA(i, j int) Matrix {
	// LinAlg pg 197
	return A.RemoveRow(i).RemoveColumn(j)
}

// Trace the main or secondary diagonal.
func (A Matrix) Trace(mainDiagonal bool) float64 {
	m, n := A.Dimensions()
	if m != n {
		panic("invalid dimensions")
	}

	var s float64
	if mainDiagonal {
		for i, r := range A {
			s += r[i]
		}
	} else {
		for i, r := range A {
			s += r[n-i-1]
		}
	}

	return s
}

// Transpose a matrix.
func (A Matrix) Transpose() Matrix {
	m, n := A.Dimensions()
	return New(n, m, func(i, j int) float64 { return A[j][i] })
}

// Vector converts a row or column matrix to a vector.
func (A Matrix) Vector() vector.Vector {
	m, n := A.Dimensions()
	switch {
	case m == 1:
		return vector.New(n, func(i int) float64 { return A[0][i] })
	case n == 1:
		return vector.New(m, func(i int) float64 { return A[i][0] })
	default:
		panic("invalid dimensions")
	}
}
