# Matrix

A matrix is an alias for a `[]Vector`.  See the vector package for details and usage.

```go
go get github.com/nathangreene3/math/linalg/matrix
```

## Constructors and Operations on Matrices

New generates an m-by-n matrix with entries defined by a generating function f.

```go
func New(m, n int, f Generator) Matrix
```

Empty returns an m-by-n matrix with zeroes for all entries.

```go
func Empty(m, n int) Matrix
```

Identity returns the m-by-n identity matrix.

```go
func Identity(m, n int) Matrix
```
