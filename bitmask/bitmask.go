package bitmask

import "math/big"

// Bitmask ...
type Bitmask uint64

// New returns b as a bitmask.
func New(b uint64) Bitmask {
	return Bitmask(b)
}

// Clear x from b.
func (b *Bitmask) Clear(x uint64) {
	*b &^= Bitmask(x)
}

// IsSet determines if b masked by x is invariant.
func (b *Bitmask) IsSet(x uint64) bool {
	y := Bitmask(x)
	return *b&y == y
}

// Set x within b.
func (b *Bitmask) Set(x uint64) {
	*b |= Bitmask(x)
}

// ------------------------------------------

// BoolBitmask ...
type BoolBitmask []bool

// NewBoolBitmask ...
func NewBoolBitmask(numBits uint64) BoolBitmask {
	return make(BoolBitmask, numBits)
}

// Clear ...
func (b *BoolBitmask) Clear(x BoolBitmask) {
	n := len(*b)
	if n != len(x) {
		panic("dimension mismatch")
	}

	for i := 0; i < n; i++ {
		if x[i] {
			(*b)[i] = false
		}
	}
}

// ClearIndices ...
func (b *BoolBitmask) ClearIndices(indices ...uint64) {
	for _, i := range indices {
		(*b)[i] = false
	}
}

// Copy ...
func (b *BoolBitmask) Copy() BoolBitmask {
	cpy := NewBoolBitmask(uint64(len(*b)))
	cpy.Set(*b)
	return cpy
}

// IndicesAreSet ...
func (b *BoolBitmask) IndicesAreSet(indices ...uint64) bool {
	for _, i := range indices {
		if !(*b)[i] {
			return false
		}
	}

	return true
}

// IsSet ...
func (b *BoolBitmask) IsSet(x BoolBitmask) bool {
	n := len(*b)
	if n != len(x) {
		panic("dimension mismatch")
	}

	for i := 0; i < n; i++ {
		if x[i] && !(*b)[i] {
			return false
		}
	}

	return true
}

// Resize ...
func (b *BoolBitmask) Resize(n uint64) {
	size := uint64(len(*b))
	switch {
	case n < size:
		*b = (*b)[:n]
	case size < n:
		*b = append(*b, make(BoolBitmask, n-size)...)
	}
}

// Set ...
func (b *BoolBitmask) Set(x BoolBitmask) {
	n := len(*b)
	if n != len(x) {
		panic("dimension mismatch")
	}

	for i := 0; i < n; i++ {
		(*b)[i] = (*b)[i] || x[i]
	}
}

// SetIndices ...
func (b *BoolBitmask) SetIndices(indices ...uint64) {
	for _, i := range indices {
		(*b)[i] = true
	}
}

// ----------------------------------------------------

// BigIntBitmask ...
type BigIntBitmask struct {
	value big.Int
}

// NewBigIntBitmask ...
func NewBigIntBitmask(n *big.Int) *BigIntBitmask {
	// TODO: check for n < 0.
	c := big.NewInt(0)
	return &BigIntBitmask{value: *c.And(c, n)}
}

// Clear ...
func (b *BigIntBitmask) Clear(n *big.Int) {
	(&b.value).AndNot(&b.value, n)
}

// Copy ...
func (b *BigIntBitmask) Copy() *BigIntBitmask {
	return NewBigIntBitmask(&b.value)
}

// IsSet ...
func (b *BigIntBitmask) IsSet(n *big.Int) bool {
	return (&b.value).And(&b.value, n).Cmp(n) == 0
}

// Set ...
func (b *BigIntBitmask) Set(n *big.Int) {
	(&b.value).Or(&b.value, n)
}

// String ...
func (b *BigIntBitmask) String() string {
	return b.value.String()
}
