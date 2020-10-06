package matrix2

type matrixErr string

const (
	dimMismatch matrixErr = "Dimension mismatch"
	luFactor    matrixErr = "LU factorization not possible"
	invalidType matrixErr = "Invalid type"
	matrixShape matrixErr = "Matrix has an invalid shape"
)

// Error returns a string describing an error that occured in a matrix operation.
func (e matrixErr) Error() string {
	return string(e)
}
