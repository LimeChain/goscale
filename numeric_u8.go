package goscale

import (
	"math"
	"math/bits"
)

type U8 uint8

func (a U8) ToNumeric() Numeric {
	return a
}

func NewU8(n uint8) Numeric {
	return U8(n)
}

func (a U8) Add(b Numeric) Numeric {
	return a + b.(U8)
}

func (a U8) Sub(b Numeric) Numeric {
	return a - b.(U8)
}

func (a U8) Mul(b Numeric) Numeric {
	return a * b.(U8)
}

func (a U8) Div(b Numeric) Numeric {
	return a / b.(U8)
}

func (a U8) Mod(b Numeric) Numeric {
	return a % b.(U8)
}

func (a U8) Eq(b Numeric) bool {
	return a == b.(U8)
}

func (a U8) Ne(b Numeric) bool {
	return a != b.(U8)
}

func (a U8) Lt(b Numeric) bool {
	return a < b.(U8)
}

func (a U8) Lte(b Numeric) bool {
	return a <= b.(U8)
}

func (a U8) Gt(b Numeric) bool {
	return a > b.(U8)
}

func (a U8) Gte(b Numeric) bool {
	return a >= b.(U8)
}

func (a U8) Max(b Numeric) Numeric {
	if a > b.(U8) {
		return a
	}
	return b
}

func (a U8) Min(b Numeric) Numeric {
	if a < b.(U8) {
		return a
	}
	return b
}

func (a U8) Clamp(min, max Numeric) Numeric {
	if a < min.(U8) {
		return min
	} else if a > max.(U8) {
		return max
	} else {
		return a
	}
}

func (a U8) TrailingZeros() Numeric {
	return U8(bits.TrailingZeros(uint(a)))
}

func (a U8) SaturatingAdd(b Numeric) Numeric {
	sum := U16(a) + U16(b.(U8))
	if sum > math.MaxUint8 {
		return U8(math.MaxUint8)
	}
	return U8(sum)
}

func (a U8) SaturatingSub(b Numeric) Numeric {
	if a < b.(U8) {
		return NewNumeric[U8](uint8(0))
	}
	return a.Sub(b)
}

func (a U8) SaturatingMul(b Numeric) Numeric {
	if a == 0 || b.(U8) == 0 {
		return U8(0)
	}

	product := uint16(a) * uint16(b.(U8))
	if product > math.MaxUint8 {
		return U8(math.MaxUint8)
	}

	return U8(product)
}
