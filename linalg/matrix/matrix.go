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
	ma, na := A.Dimensions()
	mb, nb := B.Dimensions()
	if ma != mb || na != nb {
		panic("matrices must have the same number of rows and columns")
	}

	return New(ma, na, func(i, j int) float64 { return A[i][j] + B[i][j] })
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
	ma, na := A.Dimensions()
	mb, nb := B.Dimensions()
	if ma != mb || na != nb {
		panic("matrices must have the same number of rows and columns")
	}

	return New(ma, na, func(i, j int) float64 { return A[i][j] - B[i][j] })
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
func Multiply(As ...Matrix) Matrix {
	// Source: chrome-extension://oemmndcbldboiebfnladdacbdfmadadm/https://home.cse.ust.hk/~dekai/271/notes/L12/L12.pdf

	n := len(As)
	switch n {
	case 0:
		return nil
	case 1:
		return As[0]
	case 2:
		return multiply(As[0], As[1])
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

	return multiplyByOrder(order, 0, n-1, As...)
}

// multiplyByOrder returns the product of matrices by multiplying by a given order. Initiate by calling on i = 0 and j = n-1.
func multiplyByOrder(s [][]int, i, j int, As ...Matrix) Matrix {
	if i < j {
		return multiply(multiplyByOrder(s, i, s[i][j], As...), multiplyByOrder(s, s[i][j]+1, j, As...))
	}

	return As[i]
}

// multiply returns AB. To multiply by a vector, convert the vector to a column matrix.
func multiply(A, B Matrix) Matrix {
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
	// Source: Analysis of Algorithms, 2nd Ed., by Jeffrey J. McConnell, pg 139-140
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
	for j := range A[i] {
		A[i][j] *= a
	}
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
	m, n := A.Dimensions()
	f := func(a, b int) float64 {
		if a == i {
			return A[i][b] + A[j][b]
		}

		return A[a][b]
	}

	return New(m, n, f)
}

// AddRowToRow adds row j to row i.
func (A Matrix) AddRowToRow(i, j int) {
	for k := range A[i] {
		A[i][k] += A[j][k]
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
	for k := range A[i] {
		A[i][k] -= A[j][k]
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
	for k := range A {
		A[k][i] += A[k][j]
	}
}

// SwapRows returns A with rows i and j swapped.
func SwapRows(i, j int, A Matrix) Matrix {
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
	for k, v := range A[i] {
		A[i][k] = A[j][k]
		A[i][k] = v
	}
}

// Determinant returns the Determinant of a square matrix. Panics if matrix is empty or not a square.
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

	var (
		a = float64(1)
		d float64
	)
	for i := 0; 0 < m; i++ {
		d += a * A[0][i] * RemoveColumn(RemoveRow(A, 0), i).Determinant()
		a *= -1
	}

	return d
}

// tildeA TODO: Rename this.
func (A Matrix) tildeA(i, j int) Matrix {
	// LinAlg pg 197
	return RemoveColumn(RemoveRow(A, i), j)
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

// ------------------------------------------------------------------
// ADDITIONAL OPERATIONS ON MATRICES
// ------------------------------------------------------------------
// Operations that assist elementary operations.
// ------------------------------------------------------------------

// CompareTo returns -1, 0, 1 indicating A precedes, is equal to, or
// follows B. Panics if matrices are not of equal dimension.
func CompareTo(A, B Matrix) int {
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

// CompareTo returns -1, 0, 1 indicating A precedes, is equal to, or
// follows B. Panics if matrices are not of equal dimension.
func (A Matrix) CompareTo(B Matrix) int {
	return CompareTo(A, B)
}

// Equals returns true if two matrices are equal in dimension and
// for each entry. Otherwise, it returns false.
func Equals(A, B Matrix) bool {
	return CompareTo(A, B) == 0
}

// Equals returns true if two matrices are equal in dimension and
// for each entry. Otherwise, it returns false.
func (A Matrix) Equals(B Matrix) bool {
	return Equals(A, B)
}

// Sort A such that the largest leading indices are at the top
// (index 0 is the top).
func (A Matrix) Sort() {
	sort.SliceStable(A, func(i, j int) bool { return 0 < A[i].CompareTo(A[j]) })
}

// String returns a formatted string representation of a matrix.
func (A Matrix) String() string {
	sb := strings.Builder{}

	m, n := A.Dimensions()
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
func Join(A, B Matrix) Matrix {
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
func AppendColumn(A Matrix, x vector.Vector) Matrix {
	return Join(A, ColumnMatrix(x))
}

// AppendRow returns a matrix that is the joining of a given matrix
// with a row vector. Panics if the vector dimensions are not equal
// to the number of matrix columns.
func AppendRow(A Matrix, x vector.Vector) Matrix {
	m, n := A.Dimensions()
	if n != len(x) {
		panic("matrix columns must be equal to vector dimensions")
	}

	f := func(i, j int) float64 {
		if i < m {
			return A[i][j]
		}

		return x[j]
	}

	return New(m+1, n, f)
}

// RemoveRow returns A with row i removed.
func RemoveRow(A Matrix, i int) Matrix {
	if m, _ := A.Dimensions(); i+1 < m {
		return append(A[:i], A[i+1:]...)
	}

	return A[:i]
}

// RemoveColumn returns A with column i removed.
func RemoveColumn(A Matrix, i int) Matrix {
	if _, n := A.Dimensions(); i+1 < n {
		for j := range A {
			A[j] = append(A[j][:i], A[j][i+1:]...)
		}
	} else {
		for j := range A {
			A[j] = A[j][:i]
		}
	}

	return A
}

// Reduce returns A with all row multiples removed.
func Reduce(A Matrix) Matrix {
	A.Sort()
	m, _ := A.Dimensions()
	for i := 0; i+1 < m; i++ {
		for j := i + 1; j < m; j++ {
			if vector.IsMultipleOf(A[i], A[j]) {
				A = RemoveRow(A, i)
			}
		}
	}

	return A
}

// Cost returns the number of operations to compute AB.
func Cost(A, B Matrix) int {
	ma, na := A.Dimensions()
	mb, nb := B.Dimensions()
	if na != mb {
		panic("dimension mismatch")
	}

	return ma * na * nb
}
