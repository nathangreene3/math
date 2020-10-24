package umask

import (
	"strconv"

	bm "github.com/nathangreene3/math/bitmask"
)

// UMask is a uint that is a bitmask.
type UMask uint

const (
	// Zero is the value of 0.
	Zero UMask = 0

	// One is the value of 1.
	One UMask = 1

	// Max is the largest umask.
	Max UMask = ^Zero

	// Bits is the number of bits in a umask.
	Bits uint = strconv.IntSize
)

// New ...
func New() bm.Bitmask {
	return Zero
}

// And returns a bitmask with only the bits set that are common to both bitmasks.
func (a UMask) And(b bm.Bitmask) bm.Bitmask {
	return a & b.(UMask)
}

// Base returns a string representing a bitmask in a given base n where 2 <= n <= 36.
func (a UMask) Base(n uint) string {
	return strconv.FormatUint(uint64(a), int(n))
}

// Bin returns a string representing a bitmask in binary.
func (a UMask) Bin() string {
	return strconv.FormatUint(uint64(a), 2)
}

// Clr returns a bitmask with the bits of each given bitmask b cleared from a.
func (a UMask) Clr(b ...bm.Bitmask) bm.Bitmask {
	for i := 0; i < len(b); i++ {
		a &^= b[i].(UMask)
	}

	return a
}

// ClrBits returns a bitmask with the given bits cleared from a.
func (a UMask) ClrBits(b ...uint) bm.Bitmask {
	for i := 0; i < len(b); i++ {
		a &^= 1 << b[i]
	}

	return a
}

// Dec returns a string representing a bitmask in decimal.
func (a UMask) Dec() string {
	return strconv.FormatUint(uint64(a), 10)
}

// Hex returns a string representing a bitmask in hexidecimal.
func (a UMask) Hex() string {
	return strconv.FormatUint(uint64(a), 16)
}

// Lsh returns a Bitmask shifted to the left n times.
func (a UMask) Lsh(n uint) bm.Bitmask {
	return a << n
}

// Masks determines if the bits set in b are set in a.
func (a UMask) Masks(b bm.Bitmask) bool {
	return a&b.(UMask) == b.(UMask)
}

// MasksBit determines if a bit is set.
func (a UMask) MasksBit(b uint) bool {
	var c UMask = 1 << b
	return a&c == c
}

// Not inverts a bitmask. This is equivalent to calling Max.Xor(a).
func (a UMask) Not() bm.Bitmask {
	return ^a
}

// Oct returns a string representing a bitmask in decimal.
func (a UMask) Oct() string {
	return strconv.FormatUint(uint64(a), 8)
}

// Or returns a bitmask with the bits set in either a or b.
func (a UMask) Or(b bm.Bitmask) bm.Bitmask {
	return a | b.(UMask)
}

// Rsh returns a Bitmask shifted to the right n times.
func (a UMask) Rsh(n uint) bm.Bitmask {
	return a >> n
}

// Set returns a Bitmask with bits set in each b. This is equivalent to repeatedly calling a.Or(b) for each b.
func (a UMask) Set(b ...bm.Bitmask) bm.Bitmask {
	for i := 0; i < len(b); i++ {
		a |= b[i].(UMask)
	}

	return a
}

// SetBits ...
func (a UMask) SetBits(b ...uint) bm.Bitmask {
	for i := 0; i < len(b); i++ {
		a |= 1 << b[i]
	}

	return a
}

// String returns a string representing a bitmask in decimal.
func (a UMask) String() string {
	return strconv.FormatUint(uint64(a), 10)
}

// Xor returns the bits of a and b that are set, but not simultaneously set in both a and b.
func (a UMask) Xor(b bm.Bitmask) bm.Bitmask {
	return a ^ b.(UMask)
}
