package bigmod

import (
	"strconv"
	"strings"

	"github.com/nathangreene3/math"
	mod "github.com/nathangreene3/math/mod"
)

// Z ...
type Z struct {
	value    []int
	modulus  int
	negative bool
}

// New ...
func New(value int, modulus int) *Z {
	if value < 0 {
		return &Z{value: math.Base(-value, modulus), modulus: modulus, negative: true}
	}

	return &Z{value: math.Base(value, modulus), modulus: modulus}
}

// Bin ...
func Bin(value int) *Z {
	return New(value, 2)
}

// Oct ...
func Oct(value int) *Z {
	return New(value, 8)
}

// Dec ...
func Dec(value int) *Z {
	return New(value, 10)
}

// Hex ...
func Hex(value int) *Z {
	return New(value, 16)
}

// One is equivalent to New(1,n).
func One(modulus int) *Z {
	return New(1, modulus)
}

// Zero is equivalent to New(0,n).
func Zero(modulus int) *Z {
	return New(0, modulus)
}

// Abs ...
func (x *Z) Abs() *Z {
	y := x.Copy()
	y.negative = false
	return y
}

// Add y to x.
func (x *Z) Add(y *Z) *Z {
	n := x.modulus
	if n != y.modulus {
		panic("unequal moduli")
	}

	var (
		xLen, yLen = len(x.value), len(y.value)
		minLen     = math.MinInt(xLen, yLen)
		v, k0, k1  int
	)

	if x.negative != y.negative {
		if x.CompareAbs(y) < 0 {
			x.Negate()
		}

		return x.subtractIgnoreSign(y)
	}

	return x.addIgnoreSign(y)
}

func (x *Z) addIgnoreSign(y *Z) *Z {
	n := x.modulus
	if n != y.modulus {
		panic("unequal moduli")
	}

	var (
		xLen, yLen = len(x.value), len(y.value)
		minLen     = math.MinInt(xLen, yLen)
		v, k0, k1  int
	)

	for i := 0; i < minLen; i++ {
		v, k0 = mod.AddWithCarry(x.value[i], k1, n)         // x(i) + k(i-1)
		x.value[i], k1 = mod.AddWithCarry(v, y.value[i], n) // (x(i) + k(i-1)) + y(i)
		k1 += k0
	}

	switch minLen {
	case xLen:
		for i := minLen; i < yLen; i++ {
			v, k1 = mod.AddWithCarry(y.value[i], k1, n) // y(i) + k(i-1)
			x.value = append(x.value, v)                // x(i) = y(i) + k(i-1)
		}
	case yLen:
		for i := minLen; i < xLen && k1 != 0; i++ {
			x.value[i], k1 = mod.AddWithCarry(x.value[i], k1, n) // x(i) = x(i) + k(i-1)
		}
	}

	for k1 != 0 {
		v, k1 = mod.AddWithCarry(k1, 0, n)
		x.value = append(x.value, v)
	}

	return x.clean()
}

func (x *Z) subtractIgnoreSign(y *Z) *Z {
	n := x.modulus
	if n != y.modulus {
		panic("unequal moduli")
	}

	if x.negative != y.negative && x.CompareAbs(y) < 0 {
		x.Negate()
	}

	var (
		xLen, yLen = len(x.value), len(y.value)
		minLen     = math.MinInt(xLen, yLen)
		v, k0, k1  int
	)

	for i := 0; i < minLen; i++ {
		v, k0 = mod.SubtractWithBorrow(x.value[i], k1, n)         // x(i) - k(i-1)
		x.value[i], k1 = mod.SubtractWithBorrow(v, y.value[i], n) // (x(i) - k(i-1)) - y(i)
		k1 += k0
	}

	switch minLen {
	case xLen:
		for i := minLen; i < yLen; i++ {
			v, k1 = mod.SubtractWithBorrow(y.value[i], k1, n) // y(i) - k(i-1)
			x.value = append(x.value, v)                      // x(i) = (y(i) - k(i-1)) - x(i)
		}
	case yLen:
		for i := minLen; i < xLen && k1 != 0; i++ {
			x.value[i], k1 = mod.SubtractWithBorrow(x.value[i], k1, n) // x(i) = x(i) - k(i-1)
		}
	}

	for k1 != 0 {
		v, k1 = mod.SubtractWithBorrow(k1, 0, n)
		x.value = append(x.value, v)
	}

	return x.clean()
}

func (x *Z) addInt(y int) *Z {
	return x.set(x.Integer() + y)
}

// clean calls trim and normalize.
func (x *Z) clean() *Z {
	return x.trim().normalize()
}

// Compare ...
func (x *Z) Compare(y *Z) int {
	switch {
	case x.modulus != y.modulus:
		panic("unequal moduli")
	case x.negative:
		if y.negative {
			xLen, yLen := len(x.value), len(y.value)
			switch {
			case xLen < yLen:
				return 1
			case yLen < xLen:
				return -1
			default:
				for i := 0; i < xLen; i++ {
					switch {
					case x.value[i] < y.value[i]:
						return 1
					case y.value[i] < x.value[i]:
						return -1
					}
				}

				return 0
			}
		}

		return -1
	default:
		if y.negative {
			return 1
		}

		xLen, yLen := len(x.value), len(y.value)
		switch {
		case xLen < yLen:
			return -1
		case yLen < xLen:
			return 1
		default:
			for i := 0; i < xLen; i++ {
				switch {
				case x.value[i] < y.value[i]:
					return -1
				case y.value[i] < x.value[i]:
					return 1
				}
			}

			return 0
		}
	}
}

// CompareAbs ...
func (x *Z) CompareAbs(y *Z) int {
	if x.modulus != y.modulus {
		panic("unequal moduli")
	}

	xLen, yLen := len(x.value), len(y.value)
	switch {
	case xLen < yLen:
		return -1
	case yLen < xLen:
		return 1
	default:
		for i := 0; i < xLen; i++ {
			switch {
			case x.value[i] < y.value[i]:
				return -1
			case y.value[i] < x.value[i]:
				return 1
			}
		}

		return 0
	}
}

// Copy ...
func (x *Z) Copy() *Z {
	cpy := Z{
		value:    make([]int, len(x.value)),
		modulus:  x.modulus,
		negative: x.negative,
	}

	copy(cpy.value, x.value)
	return &cpy
}

func (x *Z) divideInt(y int) *Z {
	return x.set(x.Integer() / y)
}

// Integer returns the base-10 integer value.
// TODO: removed.
func (x *Z) Integer() int {
	n := math.Base10(x.value, x.modulus)
	if x.negative {
		n *= -1
	}

	return n
}

// IsEven ...
func (x *Z) IsEven() bool {
	return len(x.value) == 0 || x.value[0]%2 == 0
}

// IsNegative ...
func (x *Z) IsNegative() bool {
	return x.negative
}

// IsOdd ...
func (x *Z) IsOdd() bool {
	return len(x.value) != 0 && x.value[0]%2 != 0
}

// IsPositive ...
func (x *Z) IsPositive() bool {
	return !x.negative
}

// IsZero ...
func (x *Z) IsZero() bool {
	return len(x.value) == 0
}

// Mulitply ...
func (x *Z) Mulitply(y *Z) *Z {
	return x.multiplyInt(y.Integer())
}

func (x *Z) multiplyInt(y int) *Z {
	return x.set(x.Integer() * y)
}

// Negate ...
func (x *Z) Negate() *Z {
	x.negative = !x.negative
	return x
}

// normalize each indexed value to Z[n], where n is the modulus.
func (x *Z) normalize() *Z {
	var (
		i, k int
		n    = len(x.value)
	)

	for ; i < n; i++ {
		x.value[i], k = mod.AddWithCarry(x.value[i], k, x.modulus)
	}

	if i < n {
		// Ran out of value before i iterated to n
		x.value = x.value[:i]
		return x
	}

	var v int
	for k != 0 {
		v, k = mod.AddWithCarry(k, 0, x.modulus)
		x.value = append(x.value, v)
	}

	return x
}

func (x *Z) set(value int) *Z {
	if x.negative = value < 0; x.negative {
		value *= -1
	}

	var (
		i int
		n = len(x.value)
	)

	for ; i < n && value != 0; i++ {
		x.value[i], value = mod.AddWithCarry(value, 0, x.modulus)
	}

	if i < n {
		// Ran out of value before i iterated to n
		x.value = x.value[:i]
		return x
	}

	var v int
	for value != 0 {
		v, value = mod.AddWithCarry(value, 0, x.modulus)
		x.value = append(x.value, v)
	}

	return x
}

func (x *Z) String() string {
	n := len(x.value)
	if n == 0 {
		return "(0) base (" + strconv.Itoa(x.modulus) + ")"
	}

	var b strings.Builder
	if x.negative {
		b.WriteByte('-')
	}

	b.WriteString("(" + strconv.Itoa(x.value[n-1]))
	for i := n - 2; 0 <= i; i-- {
		b.WriteString("," + strconv.Itoa(x.value[i]))
	}

	b.WriteString(") (base " + strconv.Itoa(x.modulus) + ")")
	return b.String()
}

// Subtract ...
func (x *Z) Subtract(y *Z) *Z {
	// var (
	// 	xLen,yLen=len(x.value),len(y.value)
	// 	minLen=math.MinInt(xLen,yLen)
	// 	v,k0,k1 int
	// )

	// if x.negative!=y.negative{

	// }

	// for i:=0;i<minLen;i++{
	// 	v,k0=
	// }

	return x.trim()
}

func (x *Z) subtractInt(y int) *Z {
	return x.set(x.Integer() - y)
}

// trim all leading zeroes.
func (x *Z) trim() *Z {
	n := len(x.value)
	for i := n - 1; 0 <= i && x.value[i] == 0; i-- {
		n--
	}

	x.value = x.value[:n]
	if n == 0 {
		x.negative = false
	}

	return x
}
