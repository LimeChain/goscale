package goscale

import (
	"encoding/binary"
	"math/big"
	"math/bits"
)

// little endian byte order
// [0] least significant bits
// [1] most significant bits
type I128 [2]U64

func NewI128(n any) I128 {
	return to128BitsNumber[I128](n)
}

func bigIntToI128(n *big.Int) I128 {
	bytes := make([]byte, 16)
	n.FillBytes(bytes)
	reverseSlice(bytes)

	var result = I128{
		U64(binary.LittleEndian.Uint64(bytes[:8])),
		U64(binary.LittleEndian.Uint64(bytes[8:])),
	}

	if n.Sign() < 0 {
		result = negateI128(result)
	}

	return result
}

func (n I128) ToBigInt() *big.Int {
	isNegative := n.isNegative()

	if isNegative {
		n = negateI128(n)
	}

	bytes := make([]byte, 16)
	binary.BigEndian.PutUint64(bytes[:8], uint64(n[1]))
	binary.BigEndian.PutUint64(bytes[8:], uint64(n[0]))
	result := big.NewInt(0).SetBytes(bytes)

	if isNegative {
		result.Neg(result)
	}

	return result
}

// ff ff ff ff ff ff ff ff | 7f ff ff ff ff ff ff ff
func MaxI128() I128 {
	return I128{
		U64(^uint64(0)),
		U64(^uint64(0) >> 1),
	}
}

// 00 00 00 00 00 00 00 00 | 80 00 00 00 00 00 00 00
func MinI128() I128 {
	return I128{
		U64(0),
		U64(1 << 63),
	}
}

func (n I128) Interface() Numeric {
	return n
}

func (a I128) Add(b Numeric) Numeric {
	sumLow, carry := bits.Add64(uint64(a[0]), uint64(b.(I128)[0]), 0)
	sumHigh, _ := bits.Add64(uint64(a[1]), uint64(b.(I128)[1]), carry)
	return I128{U64(sumLow), U64(sumHigh)}
}

func (a I128) Sub(b Numeric) Numeric {
	diffLow, borrow := bits.Sub64(uint64(a[0]), uint64(b.(I128)[0]), 0)
	diffHigh, _ := bits.Sub64(uint64(a[1]), uint64(b.(I128)[1]), borrow)
	return I128{U64(diffLow), U64(diffHigh)}
}

func (a I128) Mul(b Numeric) Numeric {
	negA := a[1]>>(64-1) == 1
	negB := b.(I128)[1]>>(64-1) == 1

	absA := a
	if negA {
		absA = negateI128(a)
	}

	absB := b.(I128)
	if negB {
		absB = negateI128(b.(I128))
	}

	high, low := bits.Mul64(uint64(absA[0]), uint64(absB[0]))
	high += uint64(absA[1])*uint64(absB[0]) + uint64(absA[0])*uint64(absB[1])

	result := I128{U64(low), U64(high)}

	// if one of the operands is negative the result is also negative
	if negA != negB {
		return negateI128(result)
	}
	return result
}

func (a I128) Div(b Numeric) Numeric {
	return bigIntToI128(
		new(big.Int).Div(a.ToBigInt(), b.(I128).ToBigInt()),
	)
}

func (a I128) Mod(b Numeric) Numeric {
	return bigIntToI128(
		new(big.Int).Mod(a.ToBigInt(), b.(I128).ToBigInt()),
	)
}

func (a I128) Eq(b Numeric) bool {
	return a.ToBigInt().Cmp(b.(I128).ToBigInt()) == 0
}

func (a I128) Ne(b Numeric) bool {
	return !a.Eq(b)
}

func (a I128) Lt(b Numeric) bool {
	return a.ToBigInt().Cmp(b.(I128).ToBigInt()) < 0
}

func (a I128) Lte(b Numeric) bool {
	return a.ToBigInt().Cmp(b.(I128).ToBigInt()) <= 0
}

func (a I128) Gt(b Numeric) bool {
	return a.ToBigInt().Cmp(b.(I128).ToBigInt()) > 0
}

func (a I128) Gte(b Numeric) bool {
	return a.ToBigInt().Cmp(b.(I128).ToBigInt()) >= 0
}

func (a I128) Max(b Numeric) Numeric {
	if a.Gt(b) {
		return a
	}
	return b
}

func (a I128) Min(b Numeric) Numeric {
	if a.Lt(b) {
		return a
	}
	return b
}

func (a I128) Clamp(minValue, maxValue Numeric) Numeric {
	if a.Lt(minValue) {
		return minValue
	}
	if a.Gt(maxValue) {
		return maxValue
	}
	return a
}

func (a I128) TrailingZeros() Numeric {
	return NewI128(a.ToBigInt().TrailingZeroBits())
}

func (a I128) SaturatingAdd(b Numeric) Numeric {
	sumLow, carry := bits.Add64(uint64(a[0]), uint64(b.(I128)[0]), 0)
	sumHigh, _ := bits.Add64(uint64(a[1]), uint64(b.(I128)[1]), carry)
	// check for overflow
	if a[1]&(1<<63) == 0 && b.(I128)[1]&(1<<63) == 0 && sumHigh&(1<<63) != 0 {
		return MaxI128()
	}
	// check for underflow
	if a[1]&(1<<63) != 0 && b.(I128)[1]&(1<<63) != 0 && sumHigh&(1<<63) == 0 {
		return MinI128()
	}
	return I128{U64(sumLow), U64(sumHigh)}
}

func (a I128) SaturatingSub(b Numeric) Numeric {
	diffLow, borrow := bits.Sub64(uint64(a[0]), uint64(b.(I128)[0]), 0)
	diffHigh, _ := bits.Sub64(uint64(a[1]), uint64(b.(I128)[1]), borrow)
	// check for overflow
	if a[1]&(1<<63) == 0 && b.(I128)[1]&(1<<63) != 0 && diffHigh&(1<<63) != 0 {
		return MaxI128()
	}
	// check for underflow
	if a[1]&(1<<63) != 0 && b.(I128)[1]&(1<<63) == 0 && diffHigh&(1<<63) == 0 {
		return MinI128()
	}
	return I128{U64(diffLow), U64(diffHigh)}
}

func (a I128) SaturatingMul(b Numeric) Numeric {
	result := new(big.Int).Mul(a.ToBigInt(), b.(I128).ToBigInt())
	// define the maximum and minimum representable I128 values
	maxI128 := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 127), big.NewInt(1))
	minI128 := new(big.Int).Neg(new(big.Int).Lsh(big.NewInt(1), 127))
	if result.Cmp(maxI128) > 0 {
		return bigIntToI128(maxI128)
	} else if result.Cmp(minI128) < 0 {
		return bigIntToI128(minI128)
	}
	return bigIntToI128(result)
}

func (n I128) isNegative() bool {
	return n[1]&U64(1<<63) != 0
}

func negateI128(n I128) I128 {
	// two's complement representation
	negLow, carry := bits.Add64(^uint64(n[0]), 1, 0)
	negHigh, _ := bits.Add64(^uint64(n[1]), 0, carry)
	return I128{U64(negLow), U64(negHigh)}
}
