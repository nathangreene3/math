package bitmask

// Bitmask ...
type Bitmask uint64

// New returns b as a bitmask.
func New(b uint64) Bitmask {
	return Bitmask(b)
}

// Set x within b.
func (b *Bitmask) Set(x uint64) {
	*b |= Bitmask(x)
}

// Has determines if b masked by x is invariant.
func (b *Bitmask) Has(x uint64) bool {
	y := Bitmask(x)
	return *b&y == y
}

// Clear x from b.
func (b *Bitmask) Clear(x uint64) {
	*b &^= Bitmask(x)
}

// ------------------------------------------

// BigBitmask ...
type BigBitmask []bool

// NewBigBitmask ...
func NewBigBitmask(numBits uint64) BigBitmask {
	return make(BigBitmask, numBits)
}

// Set ...
func (b *BigBitmask) Set(x BigBitmask) {
	n := len(*b)
	if n != len(x) {
		panic("")
	}

	for i := 0; i < n; i++ {
		(*b)[i] = (*b)[i] || x[i]
	}
}

// Has ...
func (b *BigBitmask) Has(x BigBitmask) bool {
	n := len(*b)
	if n != len(x) {
		panic("")
	}

	for i := 0; i < n; i++ {
		if x[i] && !(*b)[i] {
			return false
		}
	}

	return true
}

// Clear ...
func (b *BigBitmask) Clear(x BigBitmask) {
	n := len(*b)
	if n != len(x) {
		panic("")
	}

	for i := 0; i < n; i++ {
		if x[i] {
			(*b)[i] = false
		}
	}
}

// SetIndices ...
func (b *BigBitmask) SetIndices(indices ...uint64) {
	for _, i := range indices {
		(*b)[i] = true
	}
}

// ClearIndices ...
func (b *BigBitmask) ClearIndices(indices ...uint64) {
	for _, i := range indices {
		(*b)[i] = false
	}
}

// Copy ...
func (b *BigBitmask) Copy() BigBitmask {
	cpy := NewBigBitmask(uint64(len(*b)))
	cpy.Set(*b)
	return cpy
}

// func(b *BigBitmask)Resize(n uint64){
// 	cpy:=b.Copy()
// 	(*b)=
// }
