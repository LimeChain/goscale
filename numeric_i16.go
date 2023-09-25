package goscale

import (
	"math"
	"math/bits"
)

type I16 int16

func (n I16) Interface() Numeric {
	return n
}

func (a I16) Add(b Numeric) Numeric {
	return a + b.(I16)
}

func (a I16) Sub(b Numeric) Numeric {
	return a - b.(I16)
}

func (a I16) Mul(b Numeric) Numeric {
	return a * b.(I16)
}

func (a I16) Div(b Numeric) Numeric {
	return a / b.(I16)
}

func (a I16) Mod(b Numeric) Numeric {
	return a % b.(I16)
}

func (a I16) Eq(b Numeric) bool {
	return a == b.(I16)
}

func (a I16) Ne(b Numeric) bool {
	return a != b.(I16)
}

func (a I16) Lt(b Numeric) bool {
	return a < b.(I16)
}

func (a I16) Lte(b Numeric) bool {
	return a <= b.(I16)
}

func (a I16) Gt(b Numeric) bool {
	return a > b.(I16)
}

func (a I16) Gte(b Numeric) bool {
	return a >= b.(I16)
}

func (a I16) Max(b Numeric) Numeric {
	if a > b.(I16) {
		return a
	}
	return b
}

func (a I16) Min(b Numeric) Numeric {
	if a < b.(I16) {
		return a
	}
	return b
}

func (a I16) Clamp(min, max Numeric) Numeric {
	if a < min.(I16) {
		return min
	} else if a > max.(I16) {
		return max
	} else {
		return a
	}
}

func (a I16) TrailingZeros() Numeric {
	return I16(bits.TrailingZeros(uint(a)))
}

func (a I16) SaturatingAdd(b Numeric) Numeric {
	sum := int32(a) + int32(b.(I16))
	// check for overflow and underflow
	if sum > math.MaxInt16 {
		return I16(math.MaxInt16)
	} else if sum < math.MinInt16 {
		return I16(math.MinInt16)
	}
	return I16(sum)
}

func (a I16) SaturatingSub(b Numeric) Numeric {
	diff := int32(a) - int32(b.(I16))
	// check for overflow
	if diff > int32(math.MaxInt16) {
		return I16(math.MaxInt16)
	}
	// check for underflow
	if diff < int32(math.MinInt16) {
		return I16(math.MinInt16)
	}
	return I16(diff)
}

func (a I16) SaturatingMul(b Numeric) Numeric {
	if a == 0 || b.(I16) == 0 {
		return I16(0)
	}

	product := int32(a) * int32(b.(I16))
	// check for overflow and underflow
	if product > math.MaxInt16 {
		return I16(math.MaxInt16)
	} else if product < math.MinInt16 {
		return I16(math.MinInt16)
	}
	return I16(product)
}
