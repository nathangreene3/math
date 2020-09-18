package math

const (
	maxUInt8  = 1<<(1<<iota*8) - 1 // 255
	maxUInt16                      // 65535
	maxUInt32                      // 4294967295
	maxUInt64                      // 18446744073709551615
)

const (
	maxInt8  = 1<<(1<<iota*8-1) - 1 // 127
	maxInt16                        // 32767
	maxInt32                        // 2147483647
	maxInt64                        // ‭9223372036854775807‬
)

const (
	minInt8  = -1 << (1<<iota*8 - 1) // -128
	minInt16                         // -32768
	minInt32                         // -2147483648
	minInt64                         // -‭9223372036854775808
)
