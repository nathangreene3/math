package matrix2

import (
	gomath "math"
	"strconv"
	"strings"

	"github.com/nathangreene3/math"
	vtr "github.com/nathangreene3/math/linalg/vector2"
)

// Matrix ...
type Matrix struct {
	mat  []float64
	m, n int // Rows, columns
}

// LUType dictates how LU factorization is performed.
type LUType uint8

const (
	// Doolitle dictates LU factorization has the property Lii = 1 for all i in [0,n).
	Doolitle LUType = iota
	// Crout dictates LU factorization has the property Uii = 1 for all i in [0,n).
	Crout
	// Cholski dictates LU factorization has the property Lii = Uii for all i in [0,n).
	Cholski
)

// ----------------------------------------------------------
// Matrix constructors
// ----------------------------------------------------------

// Cols ...
func Cols(vs ...*vtr.Vector) *Matrix {
	// TODO: Improve this to not allocate twice.
	return Rows(vs...).Trans()
}

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

// Rows ...
func Rows(vs ...*vtr.Vector) *Matrix {
	n := vs[0].Dims()
	mat := append(make([]float64, 0, n*len(vs)), vs[0].Values()...)
	for i := 1; i < len(vs); i++ {
		if n != vs[i].Dims() {
			panic("dimension mismatch")
		}

		mat = append(mat, vs[i].Values()...)
	}

	return &Matrix{mat: mat, m: len(vs), n: n}
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

// Add returns A+B.
func Add(A, B *Matrix) *Matrix {
	C := A.Copy()
	C.Add(B)
	return C
}

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
	mat := make([]float64, 0, A.m)
	for j := i; j < len(A.mat); j += A.n {
		mat = append(mat, A.mat[j])
	}

	return &Matrix{mat: mat, m: A.m, n: 1}
}

// Cols ...
func (A *Matrix) Cols() int {
	return A.n
}

// ColMatrix returns a matrix having shape n by 1.
func ColMatrix(v *vtr.Vector) *Matrix {
	return &Matrix{mat: v.Values(), m: v.Dims(), n: 1}
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
	if A.m != B.m || A.n != B.n || len(A.mat) != len(B.mat) {
		return false
	}

	for i := 0; i < len(A.mat); i++ {
		if A.mat[i] != B.mat[i] {
			return false
		}
	}

	return true
}

// Format ...
func (A *Matrix) Format() string {
	var sb strings.Builder
	for i := 0; i < A.m; i++ {
		sb.WriteString("[" + strconv.FormatFloat(A.mat[i*A.n], 'f', -1, 64))
		for j := 1; j < A.n; j++ {
			sb.WriteString(" " + strconv.FormatFloat(A.mat[i*A.n+j], 'f', -1, 64))
		}

		sb.WriteString("]\n")
	}

	return sb.String()
}

// Get ...
func (A *Matrix) Get(i, j int) float64 {
	return A.mat[i*A.n+j]
}

// Identity ...
func (A *Matrix) Identity() {
	for i := 0; i < A.m; i++ {
		A.mat[i*A.n+i] = 1
		for j, jmax := i+1, i+A.n; j < jmax; j += A.n {
			A.mat[j] = 0
		}
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
	B := Join(A, Identity(A.m))
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

// Join several matrices into one.
func Join(As ...*Matrix) *Matrix {
	n := As[0].n
	for i := 1; i < len(As); i++ {
		if As[0].m != As[i].m {
			panic("invalid dimension")
		}

		n += As[i].n
	}

	mat := make([]float64, 0, As[0].m*n)
	for i := 0; i < As[0].m; i++ {
		for j := 0; j < len(As); j++ {
			mat = append(mat, As[j].mat[i*As[j].n:(i+1)*As[j].n]...)
		}
	}

	return &Matrix{mat: mat, m: As[0].m, n: n}
}

// LU ...
func (A *Matrix) LU(t LUType) (*Matrix, *Matrix) {
	if A.m != A.n {
		panic(matrixShape)
	}

	if A.mat[0] == 0 {
		return nil, nil
	}

	switch t {
	case Doolitle:
		if A.m == 1 {
			return &Matrix{mat: []float64{1}, m: 1, n: 1}, &Matrix{mat: []float64{A.mat[0]}, m: 1, n: 1}
		}

		var (
			n    = A.m
			mn   = n * n
			lmat = make([]float64, mn)
			umat = make([]float64, mn)
		)

		// 1.
		lmat[0] = 1
		umat[0] = A.mat[0]

		// 2.
		for j := 1; j < n; j++ {
			j0 := j * n
			lmat[j0] = A.mat[j0] / umat[0]
			umat[j] = A.mat[j]
		}

		// 3.
		for i, imax := 1, n-1; i < imax; i++ {
			// 4.
			ii := i * (n + 1) // i*n + i
			lmat[ii] = 1
			umat[ii] = A.mat[ii]
			for k, kmax := 0, i; k < kmax; k++ {
				umat[ii] -= lmat[i*n+k] * umat[k*n+i]
			}

			if umat[ii] == 0 {
				return nil, nil
			}

			// 5.
			for j := i + 1; j < n; j++ {
				ij, ji := i*n+j, j*n+i

				// 5.1. Set Uij
				lmat[ji] = A.mat[ji]
				umat[ij] = A.mat[ij]
				for k, kmat := 0, i; k < kmat; k++ {
					lmat[ji] -= lmat[j*n+k] * umat[k*n+i]
					umat[ij] -= lmat[i*n+k] * umat[k*n+j]
				}

				lmat[ji] /= umat[ii]
				umat[ij] /= lmat[ii]
			}
		}

		// 6.
		nn := n*n - 1 // (n-1)*n + (n-1)
		lmat[nn] = 1
		umat[nn] = A.mat[nn]
		for k, kmax := 0, n-1; k < kmax; k++ {
			umat[nn] -= lmat[(n-1)*n+k] * umat[(k+1)*n-1] // (n-1)*n + k = n^2-n+k, k*n + (n-1) = (k+1)*n - 1
		}

		return &Matrix{mat: lmat, m: n, n: n}, &Matrix{mat: umat, m: n, n: n}
	case Crout:
		if A.m == 1 {
			return &Matrix{mat: []float64{A.mat[0]}, m: 1, n: 1}, &Matrix{mat: []float64{1}, m: 1, n: 1}
		}

		// TODO
		fallthrough
	default:
		panic(invalidType)
	}
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

// ref2 ...
func (A *Matrix) ref2() {
	// TODO: Do partial pivoting.
	// TODO: Return error when no unique solution is detected?
	if A.n < A.m {
		panic("invalid dimensions")
	}

	// ------------------------------------------------------------------------
	// Gaussian elimination (Algorithm 6.1)
	// Numerical Methods, 7th Ed.
	// Richard L. Burdent and J. Douglas Faires
	//
	// Backward substitution is performed in solving Ax = y and A^-1. See Solve
	// and Inverse.
	// ------------------------------------------------------------------------

	for i, imax := 0, A.m-1; i < imax; i++ {
		p := i // Pivot index: index of first row having non-zero entry at Api
		for ; p < A.m; p++ {
			if A.mat[p*A.n+i] != 0 {
				break
			}
		}

		if p == A.m {
			// No unique solution
			return
		}

		if i != p {
			// Swap rows i and p
			for j := 0; j < A.n; j++ {
				ij, pj := i*A.n+j, p*A.n+j
				A.mat[ij], A.mat[pj] = A.mat[pj], A.mat[ij]
			}
		}

		for j := i + 1; j < A.m; j++ {
			// Update Aj as Ajk := Ajk - Aji*Aik/Aii for each column k
			for k := 0; k < A.n; k++ {
				A.mat[j*A.n+k] -= A.mat[j*A.n+i] * A.mat[i*A.n+k] / A.mat[i*(A.n+1)]
			}
		}
	}

	// if A.mat[A.m*A.m-1] == 0 {
	// 	// No unique solution
	// 	return
	// }
}

// Row ...
func (A *Matrix) Row(i int) *vtr.Vector {
	v := make([]float64, 0, A.n)
	for j, jmax := i*A.n, i*A.n+A.n; j < jmax; j++ {
		v = append(v, A.mat[j])
	}

	return vtr.New(v...)
}

// Rows ...
func (A *Matrix) Rows() int {
	return A.m
}

// RowMatrix returns a matrix having shape 1 by n.
func RowMatrix(v *vtr.Vector) *Matrix {
	return &Matrix{mat: v.Values(), m: 1, n: v.Dims()}
}

// ScalMult ...
func (A *Matrix) ScalMult(a float64) {
	for i := 0; i < len(A.mat); i++ {
		A.mat[i] *= a
	}
}

// Set ...
func (A *Matrix) Set(i, j int, v float64) {
	A.mat[i*A.n+j] = v
}

// Solve Ax = y for x.
func (A *Matrix) Solve(y *vtr.Vector) *vtr.Vector {
	B := Join(A, ColMatrix(y))
	B.ref2()

	// TODO: Beef up vector so that there's not two allocations here.
	vec := append(make([]float64, B.m-1, B.m), B.mat[B.m*B.n-1]/B.mat[B.m*B.n-2])
	for i := B.m - 2; 0 <= i; i-- {
		vec[i] = B.mat[i*B.n+B.m]
		for j := i + 1; j < B.m; j++ {
			vec[i] -= B.mat[i*B.n+j] * vec[j]
		}

		vec[i] /= B.mat[i*(B.n+1)]
	}

	return vtr.New(vec...)
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
	mat := make([]float64, 0, len(A.mat))
	for j := 0; j < A.n; j++ {
		for i := 0; i < A.m; i++ {
			mat = append(mat, A.mat[i*A.n+j])
		}
	}

	return &Matrix{mat: mat, m: A.n, n: A.m}
}

// Vector ...
func (A *Matrix) Vector() *vtr.Vector {
	if A.m != len(A.mat) && A.n != len(A.mat) {
		panic("invalid dimensions")
	}

	return vtr.New(A.mat...)
}
