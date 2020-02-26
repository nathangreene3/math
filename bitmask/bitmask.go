package bitmask

import "math/big"

// Bitmask ...
type Bitmask struct {
	value big.Int
}

// New returns a new bitmask.
func New(n *big.Int) *Bitmask {
	b := big.NewInt(0)
	return &Bitmask{value: *b.And(b, n)}
}

// NewInt returns a new bitmask.
func NewInt(n int) *Bitmask {
	return &Bitmask{value: *big.NewInt(int64(n))}
}

// Clear a value from the bitmask.
func (b *Bitmask) Clear(n *big.Int) *Bitmask {
	b.value.AndNot(&b.value, n)
	return b
}

// ClearInt a value from the bitmask.
func (b *Bitmask) ClearInt(n int) *Bitmask {
	return b.Clear(big.NewInt(int64(n)))
}

// Copy a bitmask.
func (b *Bitmask) Copy() *Bitmask {
	return New(&b.value)
}

// IsSet determines if a bitmask contains a value.
func (b *Bitmask) IsSet(n *big.Int) bool {
	return b.value.And(&b.value, n).Cmp(n) == 0
}

// IsSetInt determines if a bitmask contains a value.
func (b *Bitmask) IsSetInt(n int) bool {
	return b.IsSet(big.NewInt(int64(n)))
}

// Set a value to a bitmask.
func (b *Bitmask) Set(n *big.Int) *Bitmask {
	b.value.Or(&b.value, n)
	return b
}

// SetInt sets a value to a bitmask.
func (b *Bitmask) SetInt(n int) *Bitmask {
	return b.Set(big.NewInt(int64(n)))
}

// String returns the string representation of a bitmask.
func (b *Bitmask) String() string {
	return b.value.String()
}
