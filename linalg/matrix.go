package linalg

// Matrix ...
type Matrix interface {
	Get(i, j int) float64
	Set(i, j int, v float64)

	Add(B Matrix)
	Subtract(B Matrix)
	Multiply(B Matrix) Matrix

	Solve(v Vector) Vector
}

// Vector ...
type Vector interface {
	Get(i int) float64
	Set(i int, v float64)

	Add(w Vector)
	Subtract(w Vector)
	Multiply(a float64)
	Divide(a float64)
}
