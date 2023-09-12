package goscale

import (
	"math"
	"math/big"
	"math/bits"
)

type U64 uint64

func (n U64) Interface() Numeric {
	return n
}

func (a U64) Add(b Numeric) Numeric {
	return a + b.(U64)
}

func (a U64) Sub(b Numeric) Numeric {
	return a - b.(U64)
}

func (a U64) Mul(b Numeric) Numeric {
	return a * b.(U64)
}

func (a U64) Div(b Numeric) Numeric {
	return a / b.(U64)
}

func (a U64) Mod(b Numeric) Numeric {
	return a % b.(U64)
}

func (a U64) Eq(b Numeric) bool {
	return a == b.(U64)
}

func (a U64) Ne(b Numeric) bool {
	return a != b.(U64)
}

func (a U64) Lt(b Numeric) bool {
	return a < b.(U64)
}

func (a U64) Lte(b Numeric) bool {
	return a <= b.(U64)
}

func (a U64) Gt(b Numeric) bool {
	return a > b.(U64)
}

func (a U64) Gte(b Numeric) bool {
	return a >= b.(U64)
}

func (a U64) Max(b Numeric) Numeric {
	if a > b.(U64) {
		return a
	}
	return b
}

func (a U64) Min(b Numeric) Numeric {
	if a < b.(U64) {
		return a
	}
	return b
}

func (a U64) Clamp(min, max Numeric) Numeric {
	if a < min.(U64) {
		return min
	} else if a > max.(U64) {
		return max
	} else {
		return a
	}
}

func (a U64) TrailingZeros() Numeric {
	return U64(bits.TrailingZeros64(uint64(a)))
}

func (a U64) SaturatingAdd(b Numeric) Numeric {
	sum, carry := bits.Add64(uint64(a), uint64(b.(U64)), 0)
	if carry != 0 {
		return U64(math.MaxUint64)
	}
	return U64(sum)
}

func (a U64) SaturatingSub(b Numeric) Numeric {
	if a < b.(U64) {
		return NewNumeric[U64](uint64(0))
	}
	return a.Sub(b)
}

func (a U64) SaturatingMul(b Numeric) Numeric {
	if a == 0 || b.(U64) == 0 {
		return U64(0)
	}

	result := new(big.Int).Mul(new(big.Int).SetUint64(uint64(a)), new(big.Int).SetUint64(uint64(b.(U64))))
	// check for overflow
	if result.Cmp(new(big.Int).SetUint64(math.MaxUint64)) > 0 {
		return U64(math.MaxUint64)
	}
	return U64(result.Uint64())
}

func (a U64) CheckedAdd(b Numeric) (Numeric, error) {
	c := a + b.(U64)
	if c < a {
		return U64(0), ErrOverflow
	}
	return c, nil
}
