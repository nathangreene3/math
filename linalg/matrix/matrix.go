package matrix

import (
	gomath "math"
	"sort"
	"strings"

	vtr "github.com/nathangreene3/math/linalg/vector"
)

// ------------------------------------------------------------------------------
// RESOURCES
// ------------------------------------------------------------------------------
// Most methods defined here are taken from or are inspired by Linear Algebra,
// 3rd Ed., by Stephen H. Friedberg, Arnold J. Insel, and Lawrence E. Spence.
// Any page references in comments are in reference to this source.
// ------------------------------------------------------------------------------

// Matrix is a set of vectors.
type Matrix []vtr.Vector

// Generator is a function defining the (i,j)th entry of a matrix.
type Generator func(i, j int) float64

// ------------------------------------------------------------------------------
// MATRIX CONSTRUCTORS
// ------------------------------------------------------------------------------

// Gen generates an m-by-n matrix with entries defined by a generating function.
func Gen(m, n int, f Generator) Matrix {
	if m == 0 || n == 0 {
		return nil
	}

	A := make(Matrix, 0, m)
	for i := 0; i < m; i++ {
		A = append(A, make(vtr.Vector, 0, n))
		for j := 0; j < n; j++ {
			A[i] = append(A[i], f(i, j))
		}
	}

	return A
}

// Identity returns the n-by-n identity matrix.
func Identity(n int) Matrix {
	if n == 0 {
		return nil
	}

	A := make(Matrix, 0, n)
	for i := 0; i < n; i++ {
		A = append(A, make(vtr.Vector, n))
		A[i][i]++
	}

	return A
}

// New returns a list of several vectors into a matrix.
func New(vs ...vtr.Vector) Matrix {
	if len(vs) == 0 {
		return nil
	}

	A := make(Matrix, 0, len(vs))
	for i := 0; i < len(vs); i++ {
		if len(vs[0]) != len(vs[i]) {
			panic("invalid dimension")
		}

		A = append(A, append(make(vtr.Vector, 0, len(vs[i])), vs[i]...))
	}

	return A
}

// Zeroes returns an m-by-n matrix with zeroes for all entries.
func Zeroes(m, n int) Matrix {
	if m == 0 || n == 0 {
		return nil
	}

	A := make(Matrix, 0, m)
	for i := 0; i < m; i++ {
		A = append(A, make(vtr.Vector, n))
	}

	return A
}

// ------------------------------------------------------------------------------
// OPERATIONS ON MATRICES
// ------------------------------------------------------------------------------

// Add returns the sum of several matrices.
func Add(As ...Matrix) Matrix {
	switch len(As) {
	case 0:
		return nil
	case 1:
		return As[0].Copy()
	default:
		A := As[0].Copy()
		A.Add(As[1:]...)
		return A
	}
}

// Add several matrices to A and return A.
func (A Matrix) Add(Bs ...Matrix) {
	if 0 < len(Bs) {
		ma, na := A.Dimensions()
		for i := 0; i < len(Bs); i++ {
			if mb, nb := Bs[i].Dimensions(); ma != mb || na != nb {
				panic("matrices must have the same number of rows and columns")
			}

			for j := 0; j < ma; j++ {
				for k := 0; k < na; k++ {
					A[j][k] += Bs[i][j][k]
				}
			}
		}
	}
}

// AddColToCol returns A with column j to column i.
func AddColToCol(A Matrix, i, j int) Matrix {
	B := A.Copy()
	B.AddColToCol(i, j)
	return B
}

// AddColToCol adds column j to column i and returns A.
func (A Matrix) AddColToCol(i, j int) Matrix {
	for k := 0; k < len(A); k++ {
		A[k][i] += A[k][j]
	}

	return A
}

// AddRowToRow returns A with row j added to row i.
func AddRowToRow(A Matrix, i, j int) Matrix {
	B := A.Copy()
	B.AddRowToRow(i, j)
	return B
}

// AddRowToRow adds row j to row i and returns A.
func (A Matrix) AddRowToRow(i, j int) {
	A[i].Add(A[j])
}

// AppendColumn returns a matrix that is the joining of a given matrix with a
// column vtr.
func (A Matrix) AppendColumn(x vtr.Vector) Matrix {
	m, n := A.Dimensions()
	if m != len(x) {
		panic("matrix rows must be equal to vector length")
	}

	B := make(Matrix, 0, m)
	for i := 0; i < m; i++ {
		B = append(B, append(append(make(vtr.Vector, 0, n+1), A[i]...), x[i]))
	}

	return B
}

// AppendRow returns a matrix that is the joining of a given matrix with a row
// vtr.
func (A Matrix) AppendRow(x vtr.Vector) Matrix {
	m, n := A.Dimensions()
	if n != len(x) {
		panic("matrix columns must be equal to vector dimensions")
	}

	B := make(Matrix, 0, m+1)
	for i := 0; i < m; i++ {
		B = append(B, append(make(vtr.Vector, 0, n), A[i]...))
	}

	return append(B, append(make(vtr.Vector, 0, n), x...))
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

// ColumnMatrix converts a vector v to an n-by-1 matrix.
func ColumnMatrix(v vtr.Vector) Matrix {
	A := make(Matrix, 0, len(v))
	for i := 0; i < len(v); i++ {
		A = append(A, vtr.Vector{v[i]})
	}

	return A
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
	B := make(Matrix, 0, len(A))
	for i := 0; i < len(A); i++ {
		B = append(B, append(make(vtr.Vector, 0, len(A[i])), A[i]...))
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

// Determinant returns the Determinant of a square matrix.
func (A Matrix) Determinant() float64 {
	m, n := A.Dimensions()
	switch {
	case m == 0, n == 0:
		panic("cannot take determinant of empty matrix")
	case m != n:
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
		// TODO
		// This is probably terrible...
		var (
			B           = A.RemoveRow(0)
			sgn float64 = 1
			det float64
		)

		for i := 0; 0 < m; i++ {
			det += sgn * A[0][i] * B.RemoveColumn(i).Determinant()
			sgn *= -1
		}

		return det
	}
}

// Dimensions returns the Dimensions (number of rows, number of columns) of a
// matrix.
func (A Matrix) Dimensions() (int, int) {
	m, n := len(A), len(A[0])
	for i := 1; i < len(A); i++ {
		if n != len(A[i]) {
			panic("inconsistent matrix dimensions")
		}
	}

	return m, n
}

// Equal ...
func Equal(As ...Matrix) bool {
	for i := 1; i < len(As); i++ {
		if !As[0].Equals(As[i]) {
			return false
		}
	}

	return true
}

// Equals returns true if two matrices are equal in dimension and for each
// entry. Otherwise, it returns false.
func (A Matrix) Equals(B Matrix) bool {
	return A.Compare(B) == 0
}

// Get returns A[i][j].
func (A Matrix) Get(i, j int) float64 {
	return A[i][j]
}

// Join several matrices.
func Join(As ...Matrix) Matrix {
	switch len(As) {
	case 0:
		return nil
	case 1:
		return As[0].Copy()
	case 2:
		return As[0].Copy().Join(As[1])
	default:
		var (
			m0, n0 = As[0].Dimensions()
			ns     = append(make([]int, 0, len(As)), n0) // Partial sums up to ith n for ith A: {n0, n0+n1, n0+n1+n2, ...}
		)

		for j := 1; j < len(As); j++ {
			m, n := As[j].Dimensions()
			if m0 != m {
				panic("matrices must have equal number of rows")
			}

			ns = append(ns, n+ns[j-1])
		}

		var (
			i int // Index referencing each A
			f = func(j, k int) float64 {
				if k == 0 {
					// Reset index i
					i = 0
				}

				if ns[i] <= k {
					// Point to next A
					i++
				}

				if 0 < i {
					// k is larger than the width n of the kth A, so subtract off partial sum of ns for all previous As
					k -= ns[i-1]
				}

				return As[i][j][k]
			}
		)

		return Gen(m0, ns[len(ns)-1], f)
	}
}

// Join returns a matrix that is the joining of two given matrices.
func (A Matrix) Join(B Matrix) Matrix {
	ma, na := A.Dimensions()
	mb, nb := B.Dimensions()
	if ma != mb {
		panic("matrices must have equal number of rows")
	}

	nc := na + nb
	C := make(Matrix, 0, ma)
	for i := 0; i < ma; i++ {
		C = append(C, append(append(make(vtr.Vector, 0, nc), A[i]...), B[i]...))
	}

	return C
}

// multiply returns AB. TODO: Correct the implementation of Winograd's algorithm.
func (A Matrix) multiply(B Matrix) Matrix {
	ma, na := A.Dimensions()
	mb, nb := B.Dimensions()
	if na != mb {
		panic("A and B are of incompatible dimensions")
	}

	C := make(Matrix, 0, ma)
	for i := 0; i < ma; i++ {
		C = append(C, make(vtr.Vector, nb))
		for j := 0; j < nb; j++ {
			for k := 0; k < na; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	return C

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
	switch n := len(As); n {
	case 0:
		return nil
	case 1:
		return As[0]
	case 2:
		return As[0].multiply(As[1])
	default:
		// Source: https://home.cse.ust.hk/~dekai/271/notes/L12/L12.pdf
		var (
			dims    = make([]int, 0, n+1) // Matrix As[i] has dimension dims[i]xdims[i+1], for 0 <= i < n
			cache   = make([][]int, 0, n) // TODO: Describe this
			order   = make([][]int, 0, n) // TODO: Describe this
			j, cost int
		)

		for i := 0; i < n; i++ {
			dims = append(dims, len(As[i]))
			cache = append(cache, make([]int, n))
			order = append(order, make([]int, n))
		}

		dims = append(dims, len(As[n-1][0]))
		for h := 1; h < n; h++ {
			for i := 0; i+h < n; i++ {
				j = i + h
				cache[i][j] = gomath.MaxInt64
				for k := i; k < j; k++ {
					cost = cache[i][k] + cache[k+1][j] + dims[i]*dims[j+1]*dims[k+1]
					if cost < cache[i][j] {
						cache[i][j] = cost
						order[i][j] = k
					}
				}
			}
		}

		return multiplyByOrder(order, 0, n-1, As...)
	}
}

// multiplyByOrder returns the product of matrices by multiplying by a given
// order. Initiate by calling on i = 0 and j = n-1.
// TODO: Describe what order should be at a minimum, even though this function shouldn't be called by anyone/thing other than Multiply.
func multiplyByOrder(order [][]int, i, j int, As ...Matrix) Matrix {
	if i < j {
		return multiplyByOrder(order, i, order[i][j], As...).multiply(multiplyByOrder(order, order[i][j]+1, j, As...))
	}

	return As[i]
}

// MultiplyColumn returns A with the ith col multiplied by a.
func MultiplyColumn(A Matrix, i int, a float64) Matrix {
	B := A.Copy()
	A.MultiplyColumn(i, a)
	return B
}

// MultiplyColumn multiplies the ith column of A (A[*][i]) and returns A.
func (A Matrix) MultiplyColumn(i int, a float64) {
	for j := 0; j < len(A); j++ {
		A[j][i] *= a
	}
}

// MultiplyRow returns A with the ith row multiplied by a.
func MultiplyRow(A Matrix, i int, a float64) Matrix {
	B := A.Copy()
	B.MultiplyRow(i, a)
	return B
}

// MultiplyRow by a.
func (A Matrix) MultiplyRow(i int, a float64) {
	A[i].Multiply(a)
}

// Pow returns A^p given a square matrix A.
func Pow(A Matrix, p int) Matrix {
	B := A.Copy()
	B.Pow(p)
	return B
}

// Pow ...
func (A Matrix) Pow(p int) {
	switch m, n := A.Dimensions(); {
	case m != n:
		panic("matrix must be square")
	case p < -1:
		A.Pow(-p)
		A.Inverse()
	case p == -1:
		A.Inverse()
	case p == 0:
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				A[i][j] = 0
			}

			A[i][i]++
		}
	case p == 1:
	default:
		// Yacca's algorithm
		// TODO: Who was Yacca? Provide a source.
		P := A.Copy()

		// Reset A to be the identity
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				A[i][j] = 0
			}

			A[i][i]++
		}

		for ; 0 < p; p >>= 1 {
			if p&1 == 1 {
				A = Multiply(A, P)
			}

			P = Multiply(P, P)
		}
	}
}

// ScalarDivide returns (1/a)*A.
func ScalarDivide(A Matrix, a float64) Matrix {
	B := A.Copy()
	B.ScalarMultiply(1.0 / a)
	return B
}

// ScalarDivide A by a and return A.
func (A Matrix) ScalarDivide(a float64) {
	A.ScalarMultiply(1.0 / a)
}

// ScalarMultiply returns a*A.
func ScalarMultiply(a float64, A Matrix) Matrix {
	B := A.Copy()
	B.ScalarMultiply(a)
	return B
}

// ScalarMultiply A by a.
func (A Matrix) ScalarMultiply(a float64) {
	for i := 0; i < len(A); i++ {
		A[i].Multiply(a)
	}
}

// Set and return A.
func (A Matrix) Set(f Generator) Matrix {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			A[i][j] = f(i, j)
		}
	}

	return A
}

// Subtract returns A minus several matrices.
func Subtract(A Matrix, Bs ...Matrix) Matrix {
	B := A.Copy()
	B.Subtract(Bs...)
	return B
}

// Subtract several matrices from A and return A.
func (A Matrix) Subtract(Bs ...Matrix) {
	if 0 < len(Bs) {
		ma, na := A.Dimensions()
		for i := 0; i < len(Bs); i++ {
			if mb, nb := Bs[i].Dimensions(); ma != mb || na != nb {
				panic("matrices must have the same number of rows and columns")
			}

			for j := 0; j < ma; j++ {
				for k := 0; k < na; k++ {
					A[j][k] -= Bs[i][j][k]
				}
			}
		}
	}
}

// SubtractRowFromRow returns A with row j subtracted from row i.
func SubtractRowFromRow(A Matrix, i, j int) Matrix {
	B := A.Copy()
	B.SubtractRowFromRow(i, j)
	return B
}

// SubtractRowFromRow subtracts row j from row i.
func (A Matrix) SubtractRowFromRow(i, j int) {
	A[i].Subtract(A[j])
}

// SwapCols returns A with columns i and j swapped.
func SwapCols(A Matrix, i, j int) Matrix {
	B := A.Copy()
	B.SwapCols(i, j)
	return B
}

// SwapCols swaps two columns.
func (A Matrix) SwapCols(i, j int) {
	for k := 0; k < len(A); k++ {
		A[k][i], A[k][j] = A[k][j], A[k][i]
	}
}

// SwapRows returns A with rows i and j swapped.
func SwapRows(A Matrix, i, j int) Matrix {
	B := A.Copy()
	B.SwapRows(i, j)
	return B
}

// SwapRows swaps two rows and returns A.
func (A Matrix) SwapRows(i, j int) {
	A[i], A[j] = A[j], A[i]
}

// Solve A*x = y for x.
func (A Matrix) Solve(y vtr.Vector) vtr.Vector {
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

	for i := 0; i < m; i++ {
		if B[i][i] != 0 {
			B[i].Divide(B[i][i])
		}
	}

	x := make(vtr.Vector, 0, m)
	for i := 0; i < m; i++ {
		x = append(x, B[i][n-1])
	}

	return x // A^-1*y
}

// Inverse of a square matrix.
func (A Matrix) Inverse() Matrix {
	m, n := A.Dimensions()
	if m != n {
		panic("invalid dimensions")
	}

	B := A.Join(Identity(m))
	B.Sort()
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

	for i := 0; i < m; i++ {
		if B[i][i] != 0 {
			B[i].Divide(B[i][i])
		}
	}

	return Gen(m, n, func(i, j int) float64 { return B[i][j+n] })
}

// RemoveColumn returns a copy of a matrix with column i removed.
func (A Matrix) RemoveColumn(j int) Matrix {
	m, n := A.Dimensions()
	B := make(Matrix, 0, m)
	for i := 0; i < m; i++ {
		B = append(B, append(append(make(vtr.Vector, 0, n-1), A[i][:j]...), A[i][j+1:]...))
	}

	return B
}

// RemoveMultiples returns a sorted copy of a matrix with all row multiples
// removed.
func (A Matrix) RemoveMultiples() Matrix {
	B := A.Copy()
	B.Sort()
	for i, m := 0, len(B); i < m-1; i++ {
		for j := i + 1; j < m; j++ {
			if B[i].IsMultOf(B[j]) {
				B = B.RemoveRow(i)
			}
		}
	}

	for i := 0; i < len(B)-1; i++ {
		if B[i].IsMultOf(B[i+1]) {
			B = append(B[:i], B[i+1:]...)
		}
	}

	return B
}

// RemoveRow returns A with row i removed.
func (A Matrix) RemoveRow(i int) Matrix {
	B := make(Matrix, 0, len(A)-1)
	for j := 0; j < i; j++ {
		B = append(B, append(make(vtr.Vector, 0, len(B[j])), B[j]...))
	}

	for j := i + 1; j < len(A); j++ {
		B = append(B, append(make(vtr.Vector, 0, len(B[j])), B[j]...))
	}

	return B
}

// RowMatrix converts a vector v to a 1-by-n matrix.
func RowMatrix(v vtr.Vector) Matrix {
	return Matrix{append(make(vtr.Vector, 0, len(v)), v...)}
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

// tildeA ... TODO: Rename this.
func (A Matrix) tildeA(i, j int) Matrix {
	// LinAlg pg 197
	return A.RemoveRow(i).RemoveColumn(j)
}

// Trace the main or secondary diagonal.
func (A Matrix) Trace(mainDiagonal bool) float64 {
	n := len(A)
	if n != len(A[0]) {
		panic("invalid dimensions")
	}

	var s float64
	if mainDiagonal {
		for i := 0; i < n; i++ {
			s += A[i][i]
		}
	} else {
		for i := 0; i < n; i++ {
			s += A[i][n-i-1]
		}
	}

	return s
}

// Transpose a matrix.
func (A Matrix) Transpose() Matrix {
	m, n := A.Dimensions()
	B := make(Matrix, 0, n)
	for i := 0; i < n; i++ {
		B = append(B, make(vtr.Vector, 0, m))
		for j := 0; j < m; j++ {
			B[i] = append(B[i], A[j][i])
		}
	}

	return B
}

// Vector converts a row or column matrix to a vtr.
func (A Matrix) Vector() vtr.Vector {
	switch m, n := A.Dimensions(); {
	case m == 1:
		// Row
		return append(make(vtr.Vector, 0, n), A[0]...)
	case n == 1:
		// Column
		x := make(vtr.Vector, 0, m)
		for i := 0; i < m; i++ {
			x = append(x, A[i][0])
		}

		return x
	default:
		panic("invalid dimensions")
	}
}
