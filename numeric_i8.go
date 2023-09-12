package goscale

import (
	"math"
	"math/bits"
)

type I8 int8

func (n I8) Interface() Numeric {
	return n
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
	// check for overflow and underflow
	if sum > int16(math.MaxInt8) {
		return I8(math.MaxInt8)
	} else if sum < math.MinInt8 {
		return I8(math.MinInt8)
	}
	return I8(sum)
}

func (a I8) SaturatingSub(b Numeric) Numeric {
	diff := int16(a) - int16(b.(I8))
	// check for overflow
	if diff > int16(math.MaxInt8) {
		return I8(math.MaxInt8)
	}
	// check for underflow
	if diff < int16(math.MinInt8) {
		return I8(math.MinInt8)
	}
	return I8(diff)
}

func (a I8) SaturatingMul(b Numeric) Numeric {
	if a == 0 || b.(I8) == 0 {
		return U8(0)
	}

	product := int16(a) * int16(b.(I8))
	// check for overflow and underflow
	if product > int16(math.MaxInt8) {
		return I8(math.MaxInt8)
	} else if product < math.MinInt8 {
		return I8(math.MinInt8)
	}
	return I8(product)
}
