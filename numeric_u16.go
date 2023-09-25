package goscale

import (
	"math"
	"math/bits"
)

type U16 uint16

func (n U16) Interface() Numeric {
	return n
}

func (a U16) Add(b Numeric) Numeric {
	return a + b.(U16)
}

func (a U16) Sub(b Numeric) Numeric {
	return a - b.(U16)
}

func (a U16) Mul(b Numeric) Numeric {
	return a * b.(U16)
}

func (a U16) Div(b Numeric) Numeric {
	return a / b.(U16)
}

func (a U16) Mod(b Numeric) Numeric {
	return a % b.(U16)
}

func (a U16) Eq(b Numeric) bool {
	return a == b.(U16)
}

func (a U16) Ne(b Numeric) bool {
	return a != b.(U16)
}

func (a U16) Lt(b Numeric) bool {
	return a < b.(U16)
}

func (a U16) Lte(b Numeric) bool {
	return a <= b.(U16)
}

func (a U16) Gt(b Numeric) bool {
	return a > b.(U16)
}

func (a U16) Gte(b Numeric) bool {
	return a >= b.(U16)
}

func (a U16) Max(b Numeric) Numeric {
	if a > b.(U16) {
		return a
	}
	return b
}

func (a U16) Min(b Numeric) Numeric {
	if a < b.(U16) {
		return a
	}
	return b
}

func (a U16) Clamp(min, max Numeric) Numeric {
	if a < min.(U16) {
		return min
	} else if a > max.(U16) {
		return max
	} else {
		return a
	}
}

func (a U16) TrailingZeros() Numeric {
	return U16(bits.TrailingZeros(uint(a)))
}

func (a U16) SaturatingAdd(b Numeric) Numeric {
	sum := uint32(a) + uint32(b.(U16))
	// check for overflow
	if sum > math.MaxUint16 {
		return U16(math.MaxUint16)
	}
	return U16(sum)
}

func (a U16) SaturatingSub(b Numeric) Numeric {
	// check for underflow
	if a < b.(U16) {
		return U16(0)
	}
	return a.Sub(b)
}

func (a U16) SaturatingMul(b Numeric) Numeric {
	if a == 0 || b.(U16) == 0 {
		return U16(0)
	}
	// check for overflow
	product := uint32(a) * uint32(b.(U16))
	if product > math.MaxUint16 {
		return U16(math.MaxUint16)
	}
	return U16(product)
}
