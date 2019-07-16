# math

```go
go get github.com/nathangreene3/math
```

## linalg

The linear algebra package contains the vector and matrix packages. A vector is simply defined as a `[]float64`, and a matrix is a slice of vectors. All (or most) elementary operations on vectors and matrices are provided.

### matrix

```go
go get github.com/nathangreene3/math/linalg/matrix
```

### vector

```go
go get github.com/nathangreene3/math/vector
```

## set

```go
go get github.com/nathangreene3/math/set
```

A set is a `map[int]Comparable` that allows anything to be placed in it that is comparable, that is something that implements `CompareTo(Comparable) int`. The keys are simply the indices of the elements as they are inserted into the set, but accessing the set at an indexed value is not intended. Keys are for internal use only.

## stats
