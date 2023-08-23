package goscale

import (
	"math"
	"math/bits"
)

type U32 uint32

func (a U32) ToNumeric() Numeric {
	return a
}

func NewU32(n uint32) Numeric {
	return U32(n)
}

func (a U32) Add(b Numeric) Numeric {
	return a + b.(U32)
}

func (a U32) Sub(b Numeric) Numeric {
	return a - b.(U32)
}

func (a U32) Mul(b Numeric) Numeric {
	return a * b.(U32)
}

func (a U32) Div(b Numeric) Numeric {
	return a / b.(U32)
}

func (a U32) Mod(b Numeric) Numeric {
	return a % b.(U32)
}

func (a U32) Eq(b Numeric) bool {
	return a == b.(U32)
}

func (a U32) Ne(b Numeric) bool {
	return a != b.(U32)
}

func (a U32) Lt(b Numeric) bool {
	return a < b.(U32)
}

func (a U32) Lte(b Numeric) bool {
	return a <= b.(U32)
}

func (a U32) Gt(b Numeric) bool {
	return a > b.(U32)
}

func (a U32) Gte(b Numeric) bool {
	return a >= b.(U32)
}

func (a U32) Max(b Numeric) Numeric {
	if a > b.(U32) {
		return a
	}
	return b
}

func (a U32) Min(b Numeric) Numeric {
	if a < b.(U32) {
		return a
	}
	return b
}

func (a U32) Clamp(min, max Numeric) Numeric {
	if a < min.(U32) {
		return min
	} else if a > max.(U32) {
		return max
	} else {
		return a
	}
}

func (a U32) TrailingZeros() Numeric {
	return U32(bits.TrailingZeros(uint(a)))
}

func (a U32) SaturatingAdd(b Numeric) Numeric {
	sum := uint64(a) + uint64(b.(U32))
	if sum > math.MaxUint32 {
		return U32(math.MaxUint32)
	}
	return U32(sum)
}

func (a U32) SaturatingSub(b Numeric) Numeric {
	if a < b.(U32) {
		return NewNumeric[U32](uint32(0))
	}
	return a.Sub(b)
}

func (a U32) SaturatingMul(b Numeric) Numeric {
	if a == 0 || b.(U32) == 0 {
		return U32(0)
	}

	product := uint64(a) * uint64(b.(U32))
	if product > math.MaxUint32 {
		return U32(math.MaxUint32)
	}

	return U32(product)
}
