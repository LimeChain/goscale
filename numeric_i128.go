package goscale

import (
	"encoding/binary"
	"math"
	"math/big"
	"math/bits"
)

type I128 [2]U64

func NewI128(l, h uint64) Numeric {
	return I128{U64(l), U64(h)}
}

func NewI128FromBigInt(v big.Int) I128 {
	b := make([]byte, 16)
	v.FillBytes(b)

	reverseSlice(b)

	var result [2]U64
	result[0] = U64(binary.LittleEndian.Uint64(b[:8]))
	result[1] = U64(binary.LittleEndian.Uint64(b[8:]))

	if v.Sign() < 0 {
		result[0] = ^result[0]
		result[1] = ^result[1]

		result[0]++
		if result[0] == 0 {
			result[1]++
		}
	}

	return I128{
		result[0],
		result[1],
	}
}

func (value I128) ToBigInt() *big.Int {
	isNegative := value[1]&U64(1<<63) != 0
	if isNegative {
		if value[0] == 0 {
			value[1]--
		}
		value[0]--

		value[0] = ^value[0]
		value[1] = ^value[1]
	}

	bytes := make([]byte, 16)
	binary.BigEndian.PutUint64(bytes[:8], uint64(value[1]))
	binary.BigEndian.PutUint64(bytes[8:], uint64(value[0]))

	result := big.NewInt(0).SetBytes(bytes)
	if isNegative {
		result.Neg(result)
	}

	return result
}

func (a I128) Interface() Numeric {
	return a
}

func (a I128) Add(other Numeric) Numeric {
	var result I128

	result[1] = a[1] + other.(I128)[1]
	if result[1] < a[1] {
		result[0]++
	}
	result[0] += a[0] + other.(I128)[0]

	return result
}

func (a I128) Sub(other Numeric) Numeric {
	var diff I128

	diff[1] = a[1] - other.(I128)[1]
	if diff[1] > a[1] {
		diff[0]--
	}
	diff[0] -= a[0] - other.(I128)[0]

	return diff
}

func (a I128) Mul(other Numeric) Numeric {
	var product I128

	p1, p2 := bits.Mul64(uint64(a[0]), uint64(other.(I128)[0]))

	product[0], product[1] = U64(p1), U64(p2)

	return product
}

func (a I128) Div(other Numeric) Numeric {
	var quotient I128

	quotient[0] = a[0] / other.(I128)[0]

	return quotient
}

func (a I128) Mod(other Numeric) Numeric {
	var remainder I128

	remainder[0] = a[0] % other.(I128)[0]

	return remainder
}

func (a I128) Eq(other Numeric) bool {
	return a[0] == other.(I128)[0] && a[1] == other.(I128)[1]
}

func (a I128) Ne(other Numeric) bool {
	return a[0] != other.(I128)[0] || a[1] != other.(I128)[1]
}

func (a I128) Lt(other Numeric) bool {
	if a[1] < other.(I128)[1] {
		return true
	}

	if a[1] == other.(I128)[1] {
		return a[0] < other.(I128)[0]
	}

	return false
}

func (a I128) Lte(other Numeric) bool {
	if a[1] < other.(I128)[1] {
		return true
	}

	if a[1] == other.(I128)[1] {
		return a[0] <= other.(I128)[0]
	}

	return false
}

func (a I128) Gt(other Numeric) bool {
	if a[1] > other.(I128)[1] {
		return true
	}

	if a[1] == other.(I128)[1] {
		return a[0] > other.(I128)[0]
	}

	return false
}

func (a I128) Gte(other Numeric) bool {
	if a[1] > other.(I128)[1] {
		return true
	}

	if a[1] == other.(I128)[1] {
		return a[0] >= other.(I128)[0]
	}

	return false
}

func (a I128) Max(other Numeric) Numeric {
	if a.Gte(other) {
		return a
	}

	return other
}

func (a I128) Min(other Numeric) Numeric {
	if a.Lte(other) {
		return a
	}

	return other
}

func (a I128) Clamp(minValue, maxValue Numeric) Numeric {
	if a.Lte(minValue) {
		return minValue
	}

	if a.Gte(maxValue) {
		return maxValue
	}

	return a
}

func (a I128) TrailingZeros() Numeric {
	return U64(bits.TrailingZeros64(uint64(a[0])))
}

func (a I128) SaturatingAdd(b Numeric) Numeric {
	// check for overflow
	if a.Gt(NewI128FromBigInt(*big.NewInt(math.MaxInt64))) && b.(I128).Gt(NewI128FromBigInt(*big.NewInt(math.MaxInt64))) {
		return NewI128FromBigInt(*big.NewInt(math.MaxInt64))
	}

	return a.Add(b)
}

func (a I128) SaturatingSub(b Numeric) Numeric {
	if a.Lt(b) {
		return NewNumeric[U128](0)
	}
	return a.Sub(b)
}

func (a I128) SaturatingMul(b Numeric) Numeric {
	// check for overflow
	if a.Gt(NewI128FromBigInt(*big.NewInt(math.MaxInt64))) && b.(I128).Gt(NewI128FromBigInt(*big.NewInt(math.MaxInt64))) {
		return NewI128FromBigInt(*big.NewInt(math.MaxInt64))
	}

	return a.Mul(b)
}
