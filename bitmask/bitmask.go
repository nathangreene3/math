package bitmask

// Bitmask is defined by bitwise logic, masking, shifting, and string operations.
type Bitmask interface {
	// Logical operations
	And(Bitmask) Bitmask
	Not() Bitmask
	Or(Bitmask) Bitmask
	Xor(Bitmask) Bitmask

	// Mask operations
	Bits() int
	Clr(...Bitmask) Bitmask
	ClrBits(...int) Bitmask
	Count() int
	Masks(Bitmask) bool
	MasksBit(int) bool
	Set(...Bitmask) Bitmask
	SetBits(...int) Bitmask

	// Shift operations
	Lsh(int) Bitmask
	Rsh(int) Bitmask

	// String operations
	Base(int) string
	Bin() string
	Dec() string
	Hex() string
	Oct() string
	String() string
}
