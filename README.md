# math

```go
go get github.com/nathangreene3/math
```

## bitmask

```go
go get github.com/nathangreene3/math/bitmask
```

The bitmask package provides bitmasking functionality.

## linalg

The linear algebra package contains the vector and matrix sub-packages. A vector is defined as a `[]float64`, and a matrix is defined as a `[]vector`.

### matrix

```go
go get github.com/nathangreene3/math/linalg/matrix
```

### vector

```go
go get github.com/nathangreene3/math/vector
```

## sequence

```go
go get github.com/nathangreene3/math/sequence
```

A sequence is defined as a `func(int) float64`.

## set

```go
go get github.com/nathangreene3/math/set
```

A set is a `map[int]Comparable` that allows anything to be placed in it that is comparable, that is something that implements `Compare(Comparable) int`. The keys are simply the indices of the elements as they are inserted into the set, but accessing the set at an indexed value is not intended. Keys are for internal use only.
