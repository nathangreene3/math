package bitmask0

import "strconv"

// Bitmask ...
type Bitmask uint64

// New ...
func New() Bitmask {
	return 0
}

// And returns a Bitmask with only the bits set that are common to both bitmask.
func (a Bitmask) And(b Bitmask) Bitmask {
	return a & b
}

// AndNot returns a Bitmask with the bits common to b cleared from a. This is equivalent to a.And(b.Not()).
func (a Bitmask) AndNot(b Bitmask) Bitmask {
	return a &^ b
}

// Base returns a string representing a bitmask in a given base n where 2 <= n <= 36.
func (a Bitmask) Base(n uint64) string {
	return strconv.FormatUint(uint64(a), int(n))
}

// Bin returns a string representing a bitmask in binary.
func (a Bitmask) Bin() string {
	return strconv.FormatUint(uint64(a), 2)
}

// Clr returns a Bitmask with the bits common to each b cleared from a. This is equivalent to repeatedly calling a.AndNot(b) for each b.
func (a Bitmask) Clr(b ...Bitmask) Bitmask {
	for i := 0; i < len(b); i++ {
		a &^= b[i]
	}

	return a
}

// Dec returns a string representing a bitmask in decimal.
func (a Bitmask) Dec() string {
	return strconv.FormatUint(uint64(a), 10)
}

// Hex returns a string representing a bitmask in hexidecimal.
func (a Bitmask) Hex() string {
	return strconv.FormatUint(uint64(a), 16)
}

// Masks determines if the bits set in b are set in a.
func (a Bitmask) Masks(b Bitmask) bool {
	return a&b == b
}

// Lsh returns a Bitmask shifted to the left n times.
func (a Bitmask) Lsh(n uint64) Bitmask {
	return a << n
}

// Nand returns the bits not set in both a and b. This is equivalent to a.And(b).Not().
func (a Bitmask) Nand(b Bitmask) Bitmask {
	return ^(a & b)
}

// Not inverts a.  This is equivalent to (1<<64-1).Xor(a).
func (a Bitmask) Not() Bitmask {
	return ^a
}

// Oct returns a string representing a bitmask in decimal.
func (a Bitmask) Oct() string {
	return strconv.FormatUint(uint64(a), 8)
}

// Or returns a Bitmask with the bits set in either a or b.
func (a Bitmask) Or(b Bitmask) Bitmask {
	return a | b
}

// Rsh returns a Bitmask shifted to the right n times.
func (a Bitmask) Rsh(n uint64) Bitmask {
	return a >> n
}

// Set returns a Bitmask with bits set in each b. This is equivalent to repeatedly calling a.Or(b) for each b.
func (a Bitmask) Set(b ...Bitmask) Bitmask {
	for i := 0; i < len(b); i++ {
		a |= b[i]
	}

	return a
}

// String returns a string representing a bitmask in decimal.
func (a Bitmask) String() string {
	return strconv.FormatUint(uint64(a), 10)
}

// Xor returns the bits of a and b that are set, but not simultaneously set in both a and b.
func (a Bitmask) Xor(b Bitmask) Bitmask {
	return a ^ b
}
