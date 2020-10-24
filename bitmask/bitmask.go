package bitmask

// Bitmask is defined by bitwise logic, masking, shifting, and string operations.
type Bitmask interface {
	// Logical operations
	And(Bitmask) Bitmask
	Not() Bitmask
	Or(Bitmask) Bitmask
	Xor(Bitmask) Bitmask

	// Mask operations
	Bits() uint
	Clr(...Bitmask) Bitmask
	ClrBits(...uint) Bitmask
	Count() uint
	Masks(Bitmask) bool
	MasksBit(uint) bool
	Set(...Bitmask) Bitmask
	SetBits(...uint) Bitmask

	// Shift operations
	Lsh(uint) Bitmask
	Rsh(uint) Bitmask

	// String operations
	Base(uint) string
	Bin() string
	Dec() string
	Hex() string
	Oct() string
	String() string
}
