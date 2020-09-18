package function

// Function ...
type Function func(float64) float64

// Diff approximates f'(x) via Richardson's extrapolation.
func Diff(f Function, x, h float64, n int) float64 {
	// ------------------------------------------------------
	// Richardson's extrapolation
	// ------------------------------------------------------
	// Numerical Mathematics: Texts in Applied Mathematics
	// By Alfio Quarteroni, Riccardo Sacco, and Fausto Saleri
	// Page 387
	// ------------------------------------------------------
	// Given an estimation function of f'(x), the array A
	// represents further refinements upon f'(x) and is
	// defined by the recurrence relation
	//
	// 	A[i,0] = A(d^i*h)
	// 	A[i,j] = (A[i,j-1] - d^j*A[i-1,j-1]) / (1 - d^j).
	//
	// The array A is a lower triangular matrix shown below
	// with each entry A[i,j] having precision
	//
	// 	A[i,j] = f'(x) + O((d^i*h)^(j+1)), 0<d<1, 0<=i,j<n.
	//
	// The lower right entry A[n-1,n-1] is the best
	// approximation for f'(x).
	//
	// 	A[0,0]
	//         ↘
	// 	A[1,0] → A[1,1]
	// 	       ↘        ↘
	// 	A[2,0] → A[2,1] → A[2,2]
	//         ↘        ↘        ↘
	// 	...        ...        ...              ...
	// 	A[n-1,0] → A[n-1,1] → A[n-1,2] → ... → A[n-1,n-1]
	// ------------------------------------------------------

	var (
		d = 0.5
		A = make([][]float64, 0, n)
	)

	for i := 0; i < n; i++ {
		A = append(A, append(make([]float64, 0, i+1), ctdDiff(f, x, h)))
		h *= d
		for j, e := 1, d; j <= i; j++ {
			A[i] = append(A[i], (A[i][j-1]-e*A[i-1][j-1])/(1-e))
			e *= d
		}
	}

	return A[n-1][n-1]
}

// Diff2 ...
func Diff2(f Function, x, h float64, n int) float64 {
	// ------------------------------------------------------
	// Richardson's extrapolation
	// ------------------------------------------------------
	// Numerical Mathematics: Texts in Applied Mathematics
	// By Alfio Quarteroni, Riccardo Sacco, and Fausto Saleri
	// Page 387
	// ------------------------------------------------------
	// This is implementation improves space usage from O(n^2) to O(n), limits allocations to one, and speeds up processing by a factor of two.
	// ------------------------------------------------------

	if n == 1 {
		return ctdDiff(f, x, h)
	}

	var (
		d = 0.5
		A = make([]float64, 0, n)
	)

	for i := 0; i < n; i++ {
		A = append(A, ctdDiff(f, x, h))
		h *= d
	}

	e := d // d^i
	for i := 1; i < n-1; i++ {
		for j := n - 1; i <= j; j-- {
			A[j] = (A[j] - e*A[j-1]) / (1 - e)
		}

		e *= d
	}

	return (A[n-1] - e*A[n-2]) / (1 - e)
}

// ctdDiff returns the centered-differential approximation for f'(x).
func ctdDiff(f Function, x, dx float64) float64 {
	return (f(x+dx) - f(x-dx)) / (2 * dx)
}

// fwdDiff returns the forward-differential approximation for f'(x).
func fwdDiff(f Function, x, dx float64) float64 {
	return (f(x+dx) - f(x)) / dx
}
