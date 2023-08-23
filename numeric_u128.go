package goscale

import (
	"encoding/binary"
	"math/big"
	"math/bits"
)

type U128 [2]U64

func (a U128) ToNumeric() Numeric {
	return a
}

func NewU128(l, h uint64) Numeric {
	return U128{U64(l), U64(h)}
}

func (a U128) Add(b Numeric) Numeric {
	var result U128
	result[1], result[0] = a[0]+b.(U128)[0], a[1]+b.(U128)[1]

	// Check if there's a carry from the least significant 64 bits
	if result[0] < a[0] || (result[0] == a[0] && b.(U128)[0] != 0) {
		result[1]++
	}

	// Overflow of the most significant 64 bits cannot be detected directly
	// since we're working with fixed-size integers.
	// You might want to handle that based on the use-case.

	return result
}

func (a U128) Sub(b Numeric) Numeric {
	var result U128
	result[0], result[1] = a[0]-b.(U128)[0], a[1]-b.(U128)[1]

	// Check if there's a borrow from the least significant 64 bits
	if a[0] < b.(U128)[0] {
		result[1]--
	}

	// Underflow of the most significant 64 bits should also be handled based on the use-case.

	return result
}

func (a U128) Mul(b Numeric) Numeric {
	var result U128

	// Multiply and add the four products of the a and b parts
	// a[0] * b[0] (64 * 64 -> 128 bits)
	productLow := a[0] * b.(U128)[0]
	productHigh := (a[0] * b.(U128)[1]) + (a[1] * b.(U128)[0]) + (productLow >> 64)

	// Check for overflow on addition
	if productHigh < productLow>>64 {
		result[1]++
	}

	result[0] = a[1]*b.(U128)[1] + (productHigh >> 64)

	// Check for overflow on addition
	if result[0] < productHigh>>64 {
		// Handle overflow if needed
	}

	result[0] |= productHigh << 64
	result[1] |= productLow

	return result
}

func (a U128) Div(other Numeric) Numeric {
	var quotient U128

	quotient[0] = a[0] / other.(U128)[0]

	return quotient
}

func (a U128) Mod(other Numeric) Numeric {
	var remainder U128

	remainder[0] = a[0] % other.(U128)[0]

	return remainder
}

func (a U128) Eq(other Numeric) bool {
	return a[0] == other.(U128)[0] && a[1] == other.(U128)[1]
}

func (a U128) Ne(other Numeric) bool {
	return a[0] != other.(U128)[0] || a[1] != other.(U128)[1]
}

func (a U128) Lt(other Numeric) bool {
	if a[1] < other.(U128)[1] {
		return true
	}

	if a[1] == other.(U128)[1] {
		return a[0] < other.(U128)[0]
	}

	return false
}

func (a U128) Lte(other Numeric) bool {
	if a[1] < other.(U128)[1] {
		return true
	}

	if a[1] == other.(U128)[1] {
		return a[0] <= other.(U128)[0]
	}

	return false
}

func (a U128) Gt(other Numeric) bool {
	if a[1] > other.(U128)[1] {
		return true
	}

	if a[1] == other.(U128)[1] {
		return a[0] > other.(U128)[0]
	}

	return false
}

func (a U128) Gte(other Numeric) bool {
	if a[1] > other.(U128)[1] {
		return true
	}

	if a[1] == other.(U128)[1] {
		return a[0] >= other.(U128)[0]
	}

	return false
}

func (a U128) Max(other Numeric) Numeric {
	if a.Gte(other) {
		return a
	}

	return other
}

func (a U128) Min(other Numeric) Numeric {
	if a.Lte(other) {
		return a
	}

	return other
}

func (a U128) Clamp(minValue, maxValue Numeric) Numeric {
	if a.Lte(minValue) {
		return minValue
	}

	if a.Gte(maxValue) {
		return maxValue
	}

	return a
}

func (a U128) TrailingZeros() Numeric {
	return U8(bits.TrailingZeros64(uint64(a[0])))
}

func NewU128FromUint64(v uint64) U128 {
	return NewU128FromBigInt(new(big.Int).SetUint64(v))
}

func NewU128FromBigInt(v *big.Int) U128 {
	b := make([]byte, 16)
	v.FillBytes(b)

	reverseSlice(b)

	return U128{
		U64(binary.LittleEndian.Uint64(b[:8])),
		U64(binary.LittleEndian.Uint64(b[8:])),
	}
}

func (u U128) ToBigInt() *big.Int {
	return toBigInt(u)
}

func toBigInt(u U128) *big.Int {
	bytes := make([]byte, 16)
	binary.BigEndian.PutUint64(bytes[:8], uint64(u[1]))
	binary.BigEndian.PutUint64(bytes[8:], uint64(u[0]))

	return big.NewInt(0).SetBytes(bytes)
}

func (a U128) SaturatingAdd(b Numeric) Numeric {
	sum, carry := bits.Add64(uint64(a[0]), uint64(b.(U128)[0]), 0)

	var result U128
	result[0] = U64(sum)
	result[1] = a[1] + b.(U128)[1] + U64(carry)

	return result
}

func (a U128) SaturatingSub(b Numeric) Numeric {
	if a.Lt(b) {
		return U128{}
	}

	return a.Sub(b).(U128)
}
