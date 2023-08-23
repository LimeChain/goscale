package goscale

import (
	"math"
	"math/bits"
)

type I32 int32

func (a I32) ToNumeric() Numeric {
	return a
}

func NewI32(n int32) Numeric {
	return I32(n)
}

func (a I32) Add(b Numeric) Numeric {
	return a + b.(I32)
}

func (a I32) Sub(b Numeric) Numeric {
	return a - b.(I32)
}

func (a I32) Mul(b Numeric) Numeric {
	return a * b.(I32)
}

func (a I32) Div(b Numeric) Numeric {
	return a / b.(I32)
}

func (a I32) Mod(b Numeric) Numeric {
	return a % b.(I32)
}

func (a I32) Eq(b Numeric) bool {
	return a == b.(I32)
}

func (a I32) Ne(b Numeric) bool {
	return a != b.(I32)
}

func (a I32) Lt(b Numeric) bool {
	return a < b.(I32)
}

func (a I32) Lte(b Numeric) bool {
	return a <= b.(I32)
}

func (a I32) Gt(b Numeric) bool {
	return a > b.(I32)
}

func (a I32) Gte(b Numeric) bool {
	return a >= b.(I32)
}

func (a I32) Max(b Numeric) Numeric {
	if a > b.(I32) {
		return a
	}
	return b
}

func (a I32) Min(b Numeric) Numeric {
	if a < b.(I32) {
		return a
	}
	return b
}

func (a I32) Clamp(min, max Numeric) Numeric {
	if a < min.(I32) {
		return min
	} else if a > max.(I32) {
		return max
	} else {
		return a
	}
}

func (a I32) TrailingZeros() Numeric {
	return I32(bits.TrailingZeros(uint(a)))
}

func (a I32) SaturatingAdd(b Numeric) Numeric {
	sum := int64(a) + int64(b.(I32))

	if sum > math.MaxInt32 {
		return I32(math.MaxInt32)
	} else if sum < math.MinInt32 {
		return I32(math.MinInt32)
	}

	return I32(sum)
}

func (a I32) SaturatingSub(b Numeric) Numeric {
	if a.Lt(b) {
		return I32(math.MinInt32)
	}
	return a.Sub(b)
}
