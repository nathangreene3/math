package bitmask

// Bitmask ...
type Bitmask interface {
	// Logical operations
	And(Bitmask) Bitmask
	Not() Bitmask
	Or(Bitmask) Bitmask
	Xor(Bitmask) Bitmask

	// Mask operations
	Clr(...Bitmask) Bitmask
	ClrBits(...uint) Bitmask
	Masks(Bitmask) bool
	MasksBit(uint) bool
	Set(...Bitmask) Bitmask
	SetBits(...uint) Bitmask

	// Shift operations
	Lsh(uint) Bitmask
	Rsh(uint) Bitmask

	// String operations
	Bin() string
	Dec() string
	Hex() string
	Oct() string
	String() string
}
