package goscale

import (
	"math"
	"math/bits"
)

type I64 int64

func (n I64) Interface() Numeric {
	return n
}

func (a I64) Add(b Numeric) Numeric {
	return a + b.(I64)
}

func (a I64) Sub(b Numeric) Numeric {
	return a - b.(I64)
}

func (a I64) Mul(b Numeric) Numeric {
	return a * b.(I64)
}

func (a I64) Div(b Numeric) Numeric {
	return a / b.(I64)
}

func (a I64) Mod(b Numeric) Numeric {
	return a % b.(I64)
}

func (a I64) Eq(b Numeric) bool {
	return a == b.(I64)
}

func (a I64) Ne(b Numeric) bool {
	return a != b.(I64)
}

func (a I64) Lt(b Numeric) bool {
	return a < b.(I64)
}

func (a I64) Lte(b Numeric) bool {
	return a <= b.(I64)
}

func (a I64) Gt(b Numeric) bool {
	return a > b.(I64)
}

func (a I64) Gte(b Numeric) bool {
	return a >= b.(I64)
}

func (a I64) Max(b Numeric) Numeric {
	if a > b.(I64) {
		return a
	}
	return b
}

func (a I64) Min(b Numeric) Numeric {
	if a < b.(I64) {
		return a
	}
	return b
}

func (a I64) Clamp(min, max Numeric) Numeric {
	if a < min.(I64) {
		return min
	} else if a > max.(I64) {
		return max
	} else {
		return a
	}
}

func (a I64) TrailingZeros() Numeric {
	return I64(bits.TrailingZeros64(uint64(a)))
}

func (a I64) SaturatingAdd(b Numeric) Numeric {
	sum, _ := bits.Add64(uint64(a), uint64(b.(I64)), 0)
	// check for overflow
	if a >= 0 && b.(I64) >= 0 && int64(sum) < int64(a) {
		return I64(math.MaxInt64)
	}
	// check for underflow
	if a < 0 && b.(I64) < 0 && int64(sum) >= int64(a) {
		return I64(math.MinInt64)
	}
	return I64(sum)
}

func (a I64) SaturatingSub(b Numeric) Numeric {
	diff, borrow := bits.Sub64(uint64(a), uint64(b.(I64)), 0)
	// check for underflow
	if a >= 0 && b.(I64) < 0 && (borrow == 1 || int64(diff) < int64(a)) {
		return I64(math.MaxInt64)
	}
	// check for overflow
	if a < 0 && b.(I64) >= 0 && (borrow == 1 || int64(diff) >= int64(a)) {
		return I64(math.MinInt64)
	}
	return I64(diff)
}

func (a I64) SaturatingMul(b Numeric) Numeric {
	if uint64(a) == 0 || uint64(b.(I64)) == 0 {
		return I64(0)
	}

	hi, lo := bits.Mul64(uint64(a), uint64(b.(I64)))
	// if hi is not zero, there is an overflow
	if hi != 0 || (a > 0 && b.(I64) > 0 && int64(lo) < 0) || (a < 0 && b.(I64) < 0 && int64(lo) > 0) {
		if (a < 0 && b.(I64) >= 0) || (a >= 0 && b.(I64) < 0) {
			// negative overflow
			return I64(math.MinInt64)
		}
		// positive overflow
		return I64(math.MaxInt64)
	}
	return I64(int64(lo))
}
