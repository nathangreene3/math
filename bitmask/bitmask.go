package bitmask

import "math/big"

// Bitmask ...
type Bitmask struct {
	value big.Int
}

// New returns a new *Bitmask
func New(n *big.Int) *Bitmask {
	b := big.NewInt(0)
	return &Bitmask{value: *b.And(b, n)}
}

// Clear ...
func (b *Bitmask) Clear(n *big.Int) *Bitmask {
	b.value.AndNot(&b.value, n)
	return b
}

// Copy ...
func (b *Bitmask) Copy() *Bitmask {
	return New(&b.value)
}

// IsSet ...
func (b *Bitmask) IsSet(n *big.Int) bool {
	return b.value.And(&b.value, n).Cmp(n) == 0
}

// Set ...
func (b *Bitmask) Set(n *big.Int) *Bitmask {
	b.value.Or(&b.value, n)
	return b
}

// String ...
func (b *Bitmask) String() string {
	return b.value.String()
}
