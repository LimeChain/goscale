package goscale

import (
	"math"
	"math/bits"
)

type I8 int8

func (a I8) ToNumeric() Numeric {
	return a
}

func NewI8(n int8) Numeric {
	return I8(n)
}

func (a I8) Add(b Numeric) Numeric {
	return a + b.(I8)
}

func (a I8) Sub(b Numeric) Numeric {
	return a - b.(I8)
}

func (a I8) Mul(b Numeric) Numeric {
	return a * b.(I8)
}

func (a I8) Div(b Numeric) Numeric {
	return a / b.(I8)
}

func (a I8) Mod(b Numeric) Numeric {
	return a % b.(I8)
}

func (a I8) Eq(b Numeric) bool {
	return a == b.(I8)
}

func (a I8) Ne(b Numeric) bool {
	return a != b.(I8)
}

func (a I8) Lt(b Numeric) bool {
	return a < b.(I8)
}

func (a I8) Lte(b Numeric) bool {
	return a <= b.(I8)
}

func (a I8) Gt(b Numeric) bool {
	return a > b.(I8)
}

func (a I8) Gte(b Numeric) bool {
	return a >= b.(I8)
}

func (a I8) Max(b Numeric) Numeric {
	if a > b.(I8) {
		return a
	}
	return b
}

func (a I8) Min(b Numeric) Numeric {
	if a < b.(I8) {
		return a
	}
	return b
}

func (a I8) Clamp(min, max Numeric) Numeric {
	if a < min.(I8) {
		return min
	} else if a > max.(I8) {
		return max
	} else {
		return a
	}
}

func (a I8) TrailingZeros() Numeric {
	return I8(bits.TrailingZeros(uint(a)))
}

func (a I8) SaturatingAdd(b Numeric) Numeric {
	sum := int16(a) + int16(b.(I8))

	if sum > math.MaxInt8 {
		return I8(math.MaxInt8)
	} else if sum < math.MinInt8 {
		return I8(math.MinInt8)
	}

	return I8(sum)
}

func (a I8) SaturatingSub(b Numeric) Numeric {
	if a.Lt(b) {
		return I8(math.MinInt8)
	}
	return a.Sub(b)
}
