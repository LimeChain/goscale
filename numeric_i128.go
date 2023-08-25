package goscale

import (
	"encoding/binary"
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

// TODO: implement

func (a I128) Add(other Numeric) Numeric {
	return NewI128FromBigInt(*big.NewInt(0))
}

func (a I128) Sub(other Numeric) Numeric {
	return NewI128FromBigInt(*big.NewInt(0))
}

func (a I128) Mul(other Numeric) Numeric {
	return NewI128FromBigInt(*big.NewInt(0))
}

func (a I128) Div(other Numeric) Numeric {
	return NewI128FromBigInt(*big.NewInt(0))
}

func (a I128) Mod(other Numeric) Numeric {
	return NewI128FromBigInt(*big.NewInt(0))
}

func (a I128) Eq(other Numeric) bool {
	return false
}

func (a I128) Ne(other Numeric) bool {
	return false
}

func (a I128) Lt(other Numeric) bool {
	return false
}

func (a I128) Lte(other Numeric) bool {
	return false
}

func (a I128) Gt(other Numeric) bool {
	return false
}

func (a I128) Gte(other Numeric) bool {
	return false
}

func (a I128) Max(other Numeric) Numeric {
	return NewI128FromBigInt(*big.NewInt(0))
}

func (a I128) Min(other Numeric) Numeric {
	return NewI128FromBigInt(*big.NewInt(0))
}

func (a I128) Clamp(minValue, maxValue Numeric) Numeric {
	return NewI128FromBigInt(*big.NewInt(0))
}

func (a I128) TrailingZeros() Numeric {
	return U64(bits.TrailingZeros64(uint64(a[0])))
}

func (a I128) SaturatingAdd(b Numeric) Numeric {
	return NewI128FromBigInt(*big.NewInt(0))
}

func (a I128) SaturatingSub(b Numeric) Numeric {
	return NewI128FromBigInt(*big.NewInt(0))
}

func (a I128) SaturatingMul(b Numeric) Numeric {
	return NewI128FromBigInt(*big.NewInt(0))
}
