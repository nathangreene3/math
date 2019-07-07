package matrix

import (
	"math"
	"sort"
	"strings"

	"github.com/nathangreene3/math/linalg/vector"
)

// ------------------------------------------------------------------
// RESOURCES
// ------------------------------------------------------------------
// Most methods defined here are taken from or are inspired by
// Linear Algebra, 3rd Ed., by Stephen H. Friedberg, Arnold J. Insel,
// and Lawrence E. Spence. Any page references in comments are in
// reference to this source.
// ------------------------------------------------------------------

// Matrix is a set of vectors.
type Matrix []vector.Vector

// Generator is a function defining the (i,j)th entry of a matrix.
type Generator func(i, j int) float64

// New generates an m-by-n matrix with entries defined by a
// generating function f.
func New(m, n int, f Generator) Matrix {
	A := make(Matrix, 0, m)
	for i := 0; i < m; i++ {
		A = append(A, make(vector.Vector, 0, n))
		for j := 0; j < n; j++ {
			A[i] = append(A[i], f(i, j))
		}
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
	return New(m, n, func(i, j int) float64 { return float64((i + j) % 2) }) // Note: this generating function is called the Kronecker delta and is typically denoted as d(i,j)
}

// Dimensions returns the Dimensions (number of rows, number of
// columns) of a matrix. Panics if number of columns is not constant
// for each row.
func (A Matrix) Dimensions() (int, int) {
	m := len(A)
	n := len(A[0])
	for i := range A {
		if n != len(A[i]) {
			panic("inconsistent matrix dimensions")
		}
	}

	return m, n
}

// ------------------------------------------------------------------
// ELEMENTARY OPERATIONS ON MATRICES
// ------------------------------------------------------------------
// Mathematical operations on matrices.
// ------------------------------------------------------------------

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

// Multiply several matrices.
func Multiply(A ...Matrix) Matrix {
	// Source: chrome-extension://oemmndcbldboiebfnladdacbdfmadadm/https://home.cse.ust.hk/~dekai/271/notes/L12/L12.pdf

	n := len(A)
	switch n {
	case 0:
		return nil
	case 1:
		return A[0]
	case 2:
		return A[0].multiply(A[1])
	}

	var (
		dims    = make([]int, 0, n+1) // Matrix As[i] has dimension dims[i]xdims[i+1], for 0 <= i < n
		cache   = make([][]int, 0, n)
		order   = make([][]int, 0, n)
		j, cost int
	)
	for i := 0; i < n; i++ {
		dims = append(dims, len(A[i]))
		cache = append(cache, make([]int, n))
		order = append(order, make([]int, n))
	}
	dims = append(dims, len(A[n-1][0]))

	for h := 1; h < n; h++ {
		for i := 0; i < n-h; i++ {
			j = i + h
			cache[i][j] = math.MaxInt64
			for k := i; k < j; k++ {
				cost = cache[i][k] + cache[k+1][j] + dims[i]*dims[k+1]*dims[j+1]
				if cost < cache[i][j] {
					cache[i][j] = cost
					order[i][j] = k
				}
			}
		}
	}

	return multiplyByOrder(order, 0, n-1, A...)
}

// multiplyByOrder returns the product of matrices by multiplying by
// a given order. Initiate by calling on i = 0 and j = n-1.
func multiplyByOrder(s [][]int, i, j int, As ...Matrix) Matrix {
	if i < j {
		return multiplyByOrder(s, i, s[i][j], As...).multiply(multiplyByOrder(s, s[i][j]+1, j, As...))
	}

	return As[i]
}

// multiply returns AB. To multiply by a vector, convert the vector
// to a column matrix.
func (A Matrix) multiply(B Matrix) Matrix {
	ma, na := A.Dimensions()
	mb, nb := B.Dimensions()
	if na != mb {
		panic("A and B are of incompatible dimensions")
	}

	/*
		f := func(i, j int) float64 {
			var v float64
			for k := 0; k < na; k++ {
				v += A[i][k] * B[k][j]
			}
			return v
		}
		return New(ma, nb, f)
	*/

	// Winograd's algorithm
	// Source: Analysis of Algorithms, 2nd Ed., by Jeffrey J.
	// McConnell, pg 139-140
	var (
		d      = na / 2
		rfacts = make([]float64, 0, ma)
		cfacts = make([]float64, 0, nb)
	)
	for i := 0; i < ma; i++ {
		rfacts = append(rfacts, A[i][0]*A[i][1])
		for j := 1; j < d; j++ {
			rfacts[i] += A[i][2*j] * A[i][2*j+1]
		}
	}

	for i := 0; i < nb; i++ {
		cfacts = append(cfacts, B[0][i]*B[1][i])
		for j := 1; j < d; j++ {
			cfacts[i] += B[2*j][i] * B[2*j+1][i]
		}
	}

	C := Empty(ma, nb)
	for i := 0; i < ma; i++ {
		for j := 0; j < nb; j++ {
			C[i][j] = -rfacts[i] - cfacts[j]
			for k := 0; k < d; k++ {
				C[i][j] += (A[i][2*k] + B[2*k+1][j]) * (A[i][2*k+1] + B[2*k][j])
			}
		}
	}

	if na%2 != 0 {
		for i := 0; i < ma; i++ {
			for j := 0; j < nb; j++ {
				C[i][j] += A[i][na-1] * B[na-1][j]
			}
		}
	}

	return C
}

// ScalarMultiply returns aA.
func ScalarMultiply(a float64, A Matrix) Matrix {
	m, n := A.Dimensions()
	return New(m, n, func(i, j int) float64 { return a * A[i][j] })
}

// ScalarMultiply A by a.
func (A Matrix) ScalarMultiply(a float64) {
	for i := range A {
		A[i].Multiply(a)
	}
}

// ScalarDivide returns (1/a)A.
func ScalarDivide(a float64, A Matrix) Matrix {
	return ScalarMultiply(1.0/a, A)
}

// ScalarDivide A by a.
func (A Matrix) ScalarDivide(a float64) {
	A.ScalarMultiply(1.0 / a)
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

// SubtractRowFromRow returns A with row j subtracted from row i.
func SubtractRowFromRow(A Matrix, i, j int) Matrix {
	var (
		m, n = A.Dimensions()
		f    = func(a, b int) float64 {
			if a == i {
				return A[a][b] - A[j][b]
			}
			return A[a][b]
		}
	)
	return New(m, n, f)
}

// SubtractRowFromRow subtracts row j from row i.
func (A Matrix) SubtractRowFromRow(i, j int) {
	A[i].Subtract(A[j])
}

// AddColToCol returns A with column j to column i.
func AddColToCol(A Matrix, i, j int) Matrix {
	var (
		m, n = A.Dimensions()
		f    = func(a, b int) float64 {
			if b == i {
				return A[a][b] + A[a][j]
			}
			return A[a][b]
		}
	)
	return New(m, n, f)
}

// AddColToCol adds column j to column i.
func (A Matrix) AddColToCol(i, j int) {
	for k := range A {
		A[k][i] += A[k][j]
	}
}

// SwapRows returns A with rows i and j swapped.
func SwapRows(A Matrix, i, j int) Matrix {
	var (
		m, n = A.Dimensions()
		f    = func(a, b int) float64 {
			switch a {
			case i:
				return A[j][b]
			case j:
				return A[i][b]
			default:
				return A[a][b]
			}
		}
	)
	return New(m, n, f)
}

// SwapRows swaps two rows.
func (A Matrix) SwapRows(i, j int) {
	temp := A[i]
	A[i] = A[j]
	A[j] = temp
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
	default:
		// This is probably terrible...
		var (
			a = 1.0
			d float64
			B = A.RemoveRow(0)
		)
		for i := 0; 0 < m; i++ {
			d += a * A[0][i] * B.RemoveColumn(i).Determinant()
			a *= -1
		}

		return d
	}
}

// tildeA TODO: Rename this.
func (A Matrix) tildeA(i, j int) Matrix {
	// LinAlg pg 197
	return A.RemoveRow(i).RemoveColumn(j)
}

// Transpose returns the transpose of a matrix.
func Transpose(A Matrix) Matrix {
	m, n := A.Dimensions()
	return New(n, m, func(i, j int) float64 { return A[j][i] })
}

// Pow returns A^p, for square matrix A and 0 <= p.
func Pow(A Matrix, p int) Matrix {
	if p < 0 {
		panic("power must be non-negative")
	}

	m, n := A.Dimensions()
	if m != n {
		panic("matrix must be square")
	}

	if p == 0 {
		return Identity(m, n)
	}

	B := A.Copy()
	for ; 1 < p; p /= 2 {
		B = Multiply(B, B)
	}

	if 0 < p {
		return Multiply(B, A)
	}

	return B
}

// Solve Ax=b for x.
func (A Matrix) Solve(b vector.Vector) vector.Vector {
	var (
		B    = A.Join(ColumnMatrix(b)).Reduce()
		m, n = B.Dimensions()
	)
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

	for i := range B {
		if B[i][i] != 0 {
			B[i].Divide(B[i][i])
		}
	}

	return vector.New(m, func(i int) float64 { return B[i][n-1] }) // x = A^-1*b
}

// Inverse of a square matrix. Caution: not all matrices, even square ones, are guarenteed to be invertible.
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

	for i := range B {
		if B[i][i] != 0 {
			B[i].Divide(B[i][i])
		}
	}

	return New(m, m, func(i, j int) float64 { return B[i+m][j] })
}

// Vector converts a row or column matrix to a vector.
func (A Matrix) Vector() vector.Vector {
	m, n := A.Dimensions()
	if m == 1 {
		return vector.New(n, func(i int) float64 { return A[0][i] })
	}

	if n == 1 {
		return vector.New(m, func(i int) float64 { return A[i][0] })
	}

	panic("invalid dimensions")
}

// ------------------------------------------------------------------
// ADDITIONAL OPERATIONS ON MATRICES
// ------------------------------------------------------------------
// Operations that assist elementary operations.
// ------------------------------------------------------------------

// CompareTo returns -1, 0, 1 indicating A precedes, is equal to, or
// follows B. Panics if matrices are not of equal dimension.
func (A Matrix) CompareTo(B Matrix) int {
	ma, na := A.Dimensions()
	mb, nb := B.Dimensions()
	if ma != mb || na != nb {
		panic("dimension mismatch")
	}

	for i := 0; i < ma; i++ {
		for j := 0; j < na; j++ {
			if A[i][j] < B[i][j] {
				return -1
			}

			if B[i][j] < A[i][j] {
				return 1
			}
		}
	}

	return 0
}

// Equals returns true if two matrices are equal in dimension and
// for each entry. Otherwise, it returns false.
func (A Matrix) Equals(B Matrix) bool {
	return A.CompareTo(B) == 0
}

// Sort A such that the largest leading indices are at the top
// (index 0 is the top).
func (A Matrix) Sort() {
	sort.SliceStable(A, func(i, j int) bool { return 0 < A[i].CompareTo(A[j]) })
}

// String returns a formatted string representation of a matrix.
func (A Matrix) String() string {
	var (
		sb   = strings.Builder{}
		m, n = A.Dimensions()
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

// Copy returns a deep copied matrix.
func (A Matrix) Copy() Matrix {
	m, n := A.Dimensions()
	return New(m, n, func(i, j int) float64 { return A[i][j] })
}

// RowMatrix converts a vector v to a 1-by-n matrix.
func RowMatrix(v vector.Vector) Matrix {
	return New(1, len(v), func(i, j int) float64 { return v[j] })
}

// ColumnMatrix converts a vector v to an n-by-1 matrix.
func ColumnMatrix(v vector.Vector) Matrix {
	return New(len(v), 1, func(i, j int) float64 { return v[i] })
}

// Join returns a matrix that is the joining of two given matrices.
// Panics if number of rows are not equal.
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

// AppendColumn returns a matrix that is the joining of a given
// matrix with a column Vector.
func (A Matrix) AppendColumn(x vector.Vector) Matrix {
	return A.Join(ColumnMatrix(x))
}

// AppendRow returns a matrix that is the joining of a given matrix
// with a row vector.
func (A Matrix) AppendRow(x vector.Vector) Matrix {
	if _, n := A.Dimensions(); n != len(x) {
		panic("matrix columns must be equal to vector dimensions")
	}

	return append(A, x)
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

// Reduce returns a sorted copy of a matrix with all row multiples
// removed.
func (A Matrix) Reduce() Matrix {
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

// Cost returns the number of operations to compute AB.
func (A Matrix) Cost(B Matrix) int {
	ma, na := A.Dimensions()
	mb, nb := B.Dimensions()
	if na != mb {
		panic("dimension mismatch")
	}

	return ma * na * nb
}
