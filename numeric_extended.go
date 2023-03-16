package goscale

import (
	"errors"
	"math"
	"math/bits"
)

// TODO:
// refactor this to be a separate package
// that defines extended numeric types
// with operations handling overflow and underflow
//
// Example:
// package numeric
// type U8Ext uint8
//
// package goscale
// type U8 U8Ext

// TODO: I8, I16, I32, ...

func (a U8) Add(b U8) U8 {
	return a + b
}

func (a U8) Sub(b U8) U8 {
	return a - b
}

func (a U8) Mul(b U8) U8 {
	return a * b
}

func (a U8) Div(b U8) U8 {
	return a / b
}

func (a U8) Max(b U8) U8 {
	if a > b {
		return a
	}

	return b
}

func (a U8) Min(b U8) U8 {
	if a < b {
		return a
	}

	return b
}

func (a U8) TrailingZeros() U8 {
	return U8(bits.TrailingZeros(uint(a)))
}

func (a U8) Clamp(minValue, maxValue U8) U8 {
	if a < minValue {
		return minValue
	} else if a > maxValue {
		return maxValue
	} else {
		return a
	}
}

func (a U16) Add(b U16) U16 {
	return a + b
}

func (a U16) Sub(b U16) U16 {
	return a - b
}

func (a U16) Mul(b U16) U16 {
	return a * b
}

func (a U16) Div(b U16) U16 {
	return a / b
}

func (a U16) Max(b U16) U16 {
	if a > b {
		return a
	}

	return b
}

func (a U16) Min(b U16) U16 {
	if a < b {
		return a
	}

	return b
}

func (a U16) TrailingZeros() U16 {
	return U16(bits.TrailingZeros(uint(a)))
}

func (a U16) Clamp(minValue, maxValue U16) U16 {
	if a < minValue {
		return minValue
	} else if a > maxValue {
		return maxValue
	} else {
		return a
	}
}

func (a U32) Add(b U32) U32 {
	return a + b
}

func (a U32) Sub(b U32) U32 {
	return a - b
}

func (a U32) Mul(b U32) U32 {
	return a * b
}

func (a U32) Div(b U32) U32 {
	return a / b
}

func (a U32) Max(b U32) U32 {
	if a > b {
		return a
	}

	return b
}

func (a U32) Min(b U32) U32 {
	if a < b {
		return a
	}

	return b
}

func (a U32) TrailingZeros() U32 {
	return U32(bits.TrailingZeros(uint(a)))
}

func (a U32) Clamp(minValue, maxValue U32) U32 {
	if a < minValue {
		return minValue
	} else if a > maxValue {
		return maxValue
	} else {
		return a
	}
}

func (a U64) Add(b U64) U64 {
	return a + b
}

func (a U64) Sub(b U64) U64 {
	return a - b
}

func (a U64) Mul(b U64) U64 {
	return a * b
}

func (a U64) Div(b U64) U64 {
	return a / b
}

func (a U64) Max(b U64) U64 {
	if a > b {
		return a
	}

	return b
}

func (a U64) Min(b U64) U64 {
	if a < b {
		return a
	}

	return b
}

func (a U64) TrailingZeros() U64 {
	return U64(bits.TrailingZeros(uint(a)))
}

func (a U64) Clamp(minValue, maxValue U64) U64 {
	if a < minValue {
		return minValue
	} else if a > maxValue {
		return maxValue
	} else {
		return a
	}
}

func (a U8) SaturatingAdd(b U8) U8 {
	c := a + b

	if c < a {
		return U8(math.MaxUint8)
	}

	return c
}

func (a U8) SaturatingSub(b U8) U8 {
	c := a - b

	if c > a {
		return U8(0)
	}

	return c
}

func (a U8) SaturatingMul(b U8) U8 {
	if a == 0 || b == 0 {
		return U8(0)
	}

	c := a * b

	if c/a != b {
		return U8(math.MaxUint8)
	}

	return c
}

func (a U8) SaturatingDiv(b U8) U8 {
	if b == 0 {
		return U8(math.MaxUint8)
	}

	return a / b
}

func (a U16) SaturatingAdd(b U16) U16 {
	c := a + b

	if c < a {
		return U16(math.MaxUint16)
	}

	return c
}

func (a U16) SaturatingSub(b U16) U16 {
	c := a - b

	if c > a {
		return U16(0)
	}

	return c
}

func (a U16) SaturatingMul(b U16) U16 {
	if a == 0 || b == 0 {
		return U16(0)
	}

	c := a * b

	if c/a != b {
		return U16(math.MaxUint16)
	}

	return c
}

func (a U16) SaturatingDiv(b U16) U16 {
	if b == 0 {
		return U16(math.MaxUint16)
	}

	return a / b
}

func (a U32) SaturatingAdd(b U32) U32 {
	c := a + b

	if c < a {
		return U32(math.MaxUint32)
	}

	return c
}

func (a U32) SaturatingSub(b U32) U32 {
	c := a - b

	if c > a {
		return U32(0)
	}

	return c
}

func (a U32) SaturatingMul(b U32) U32 {
	if a == 0 || b == 0 {
		return U32(0)
	}

	c := a * b

	if c/a != b {
		return U32(math.MaxUint32)
	}

	return c
}

func (a U32) SaturatingDiv(b U32) U32 {
	if b == 0 {
		return U32(math.MaxUint32)
	}

	return a / b
}

func (a U64) SaturatingAdd(b U64) U64 {
	c := a + b

	if c < a {
		return U64(math.MaxUint64)
	}

	return c
}

func (a U64) SaturatingSub(b U64) U64 {
	c := a - b

	if c > a {
		return U64(0)
	}

	return c
}

func (a U64) SaturatingMul(b U64) U64 {
	if a == 0 || b == 0 {
		return U64(0)
	}

	c := a * b

	if c/a != b {
		return U64(math.MaxUint64)
	}

	return c
}

func (a U64) SaturatingDiv(b U64) U64 {
	if b == 0 {
		return U64(math.MaxUint64)
	}

	return a / b
}

// TODO: implement for other types
func (a U64) CheckedAdd(b U64) (U64, error) {
	c := a + b

	if c < a {
		return 0, errors.New("overflow")
	}

	return c, nil
}
