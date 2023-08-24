package goscale

import (
	"encoding/binary"
	"math/big"
)

type U128 [2]U64

func NewU128(l, h uint64) Numeric {
	return U128{U64(l), U64(h)}
}

func NewU128FromBigInt(b *big.Int) U128 {
	bytes := make([]byte, 16)
	b.FillBytes(bytes)
	reverseSlice(bytes)
	return U128{
		U64(binary.LittleEndian.Uint64(bytes[:8])),
		U64(binary.LittleEndian.Uint64(bytes[8:])),
	}
}

func (u U128) ToBigInt() *big.Int {
	bytes := make([]byte, 16)
	binary.BigEndian.PutUint64(bytes[:8], uint64(u[1]))
	binary.BigEndian.PutUint64(bytes[8:], uint64(u[0]))
	return big.NewInt(0).SetBytes(bytes)
}

func NewU128FromUint64(v uint64) U128 {
	return NewU128FromBigInt(new(big.Int).SetUint64(v))
}

func (a U128) Interface() Numeric {
	return a
}

func (a U128) Add(b Numeric) Numeric {
	return NewU128FromBigInt(
		new(big.Int).Add(a.ToBigInt(), b.(U128).ToBigInt()),
	)
}

func (a U128) Sub(b Numeric) Numeric {
	return NewU128FromBigInt(
		new(big.Int).Sub(a.ToBigInt(), b.(U128).ToBigInt()),
	)
}

func (a U128) Mul(b Numeric) Numeric {
	return NewU128FromBigInt(
		new(big.Int).Mul(a.ToBigInt(), b.(U128).ToBigInt()),
	)
}

func (a U128) Div(b Numeric) Numeric {
	return NewU128FromBigInt(
		new(big.Int).Div(a.ToBigInt(), b.(U128).ToBigInt()),
	)
}

func (a U128) Mod(b Numeric) Numeric {
	return NewU128FromBigInt(
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
	if a.Gte(b.(U128)) {
		return a
	}
	return b.(U128)
}

func (a U128) Min(b Numeric) Numeric {
	if a.Lte(b.(U128)) {
		return a
	}
	return b.(U128)
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
	return U128{U64(a.ToBigInt().TrailingZeroBits()), 0}
}

func (a U128) SaturatingAdd(b Numeric) Numeric {
	result := new(big.Int).Add(a.ToBigInt(), b.(U128).ToBigInt())

	// check for overflow
	maxValue := new(big.Int)
	maxValue.Lsh(big.NewInt(1), 128)
	maxValue.Sub(maxValue, big.NewInt(1))
	if result.Cmp(maxValue) > 0 {
		result.Set(maxValue)
	}

	return NewU128FromBigInt(result)
}

func (a U128) SaturatingSub(b Numeric) Numeric {
	result := new(big.Int).Sub(a.ToBigInt(), b.(U128).ToBigInt())

	// check for underflow
	if result.Sign() < 0 {
		result.SetInt64(0)
	}

	return NewU128FromBigInt(result)
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

	return NewU128FromBigInt(result)
}
