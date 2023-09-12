package goscale

import (
	"encoding/binary"
	"math/big"
	"math/bits"
)

// little endian byte order
// [0] least significant bits
// [1] most significant bits
type U128 [2]U64

func NewU128(n any) U128 {
	return fromAnyNumberTo128Bits[U128](n)
}

func bigIntToU128(n *big.Int) U128 {
	bytes := make([]byte, 16)
	n.FillBytes(bytes)
	reverseSlice(bytes)

	return U128{
		U64(binary.LittleEndian.Uint64(bytes[:8])),
		U64(binary.LittleEndian.Uint64(bytes[8:])),
	}
}

func (n U128) ToBigInt() *big.Int {
	bytes := make([]byte, 16)
	binary.BigEndian.PutUint64(bytes[:8], uint64(n[1]))
	binary.BigEndian.PutUint64(bytes[8:], uint64(n[0]))
	return big.NewInt(0).SetBytes(bytes)
}

// ff ff ff ff ff ff ff ff | ff ff ff ff ff ff ff ff
func MaxU128() U128 {
	return U128{
		U64(^uint64(0)),
		U64(^uint64(0)),
	}
}

// 00 00 00 00 00 00 00 00 | 00 00 00 00 00 00 00 00
func MinU128() U128 {
	return U128{
		U64(0),
		U64(0),
	}
}

func (n U128) Interface() Numeric {
	return n
}

func (a U128) Add(b Numeric) Numeric {
	sumLow, carry := bits.Add64(uint64(a[0]), uint64(b.(U128)[0]), 0)
	sumHigh, _ := bits.Add64(uint64(a[1]), uint64(b.(U128)[1]), carry)
	return U128{U64(sumLow), U64(sumHigh)}
}

func (a U128) Sub(b Numeric) Numeric {
	diffLow, borrow := bits.Sub64(uint64(a[0]), uint64(b.(U128)[0]), 0)
	diffHigh, _ := bits.Sub64(uint64(a[1]), uint64(b.(U128)[1]), borrow)
	return U128{U64(diffLow), U64(diffHigh)}
}

func (a U128) Mul(b Numeric) Numeric {
	high, low := bits.Mul64(uint64(a[0]), uint64(b.(U128)[0]))
	high += uint64(a[1])*uint64(b.(U128)[0]) + uint64(a[0])*uint64(b.(U128)[1])
	return U128{U64(low), U64(high)}
}

func (a U128) Div(b Numeric) Numeric {
	return bigIntToU128(
		new(big.Int).Div(a.ToBigInt(), b.(U128).ToBigInt()),
	)
}

func (a U128) Mod(b Numeric) Numeric {
	return bigIntToU128(
		new(big.Int).Mod(a.ToBigInt(), b.(U128).ToBigInt()),
	)
}

func (a U128) Eq(b Numeric) bool {
	return a.ToBigInt().Cmp(b.(U128).ToBigInt()) == 0
}

func (a U128) Ne(b Numeric) bool {
	return !a.Eq(b)
}

func (a U128) Lt(b Numeric) bool {
	return a.ToBigInt().Cmp(b.(U128).ToBigInt()) < 0
}

func (a U128) Lte(b Numeric) bool {
	return a.ToBigInt().Cmp(b.(U128).ToBigInt()) <= 0
}

func (a U128) Gt(b Numeric) bool {
	return a.ToBigInt().Cmp(b.(U128).ToBigInt()) > 0
}

func (a U128) Gte(b Numeric) bool {
	return a.ToBigInt().Cmp(b.(U128).ToBigInt()) >= 0
}

func (a U128) Max(b Numeric) Numeric {
	if a.Gt(b) {
		return a
	}
	return b
}

func (a U128) Min(b Numeric) Numeric {
	if a.Lt(b) {
		return a
	}
	return b
}

func (a U128) Clamp(minValue, maxValue Numeric) Numeric {
	if a.Lt(minValue) {
		return minValue
	}
	if a.Gt(maxValue) {
		return maxValue
	}
	return a
}

func (a U128) TrailingZeros() Numeric {
	return NewU128(a.ToBigInt().TrailingZeroBits())
}

func (a U128) SaturatingAdd(b Numeric) Numeric {
	sumLow, carry := bits.Add64(uint64(a[0]), uint64(b.(U128)[0]), 0)
	sumHigh, overflow := bits.Add64(uint64(a[1]), uint64(b.(U128)[1]), carry)
	// check for overflow
	if overflow == 1 || (carry == 1 && sumHigh <= uint64(a[1]) && sumHigh <= uint64(b.(U128)[1])) {
		return MaxU128()
	}
	return U128{U64(sumLow), U64(sumHigh)}
}

func (a U128) SaturatingSub(b Numeric) Numeric {
	low, borrow := bits.Sub64(uint64(a[0]), uint64(b.(U128)[0]), 0)
	high, _ := bits.Sub64(uint64(a[1]), uint64(b.(U128)[1]), borrow)
	// check for underflow
	if borrow == 1 || high > uint64(a[1]) {
		return U128{0, 0}
	}
	return U128{U64(low), U64(high)}
}

func (a U128) SaturatingMul(b Numeric) Numeric {
	result := new(big.Int).Mul(a.ToBigInt(), b.(U128).ToBigInt())
	// check for overflow
	maxValue := new(big.Int)
	maxValue.Lsh(big.NewInt(1), 128)
	maxValue.Sub(maxValue, big.NewInt(1))
	if result.Cmp(maxValue) > 0 {
		result.Set(maxValue)
	}
	return bigIntToU128(result)
}
