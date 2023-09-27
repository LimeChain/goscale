package goscale

import (
	"errors"
	"math"
	"math/bits"
)

var (
	errOverflow  = errors.New("overflow")
	errUnderflow = errors.New("underflow")
)

func Clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}

	return value
}

func Max64(a, b U64) U64 {
	if a > b {
		return a
	}
	return b
}

func Max128(a, b U128) U128 {
	if a.Gt(b) {
		return a
	}
	return b
}

// ff ff ff ff ff ff ff ff | 7f ff ff ff ff ff ff ff
func MaxI128() I128 {
	return I128{
		U64(^uint64(0)),
		U64(^uint64(0) >> 1),
	}
}

// ff ff ff ff ff ff ff ff | ff ff ff ff ff ff ff ff
func MaxU128() U128 {
	return U128{
		U64(^uint64(0)),
		U64(^uint64(0)),
	}
}

func Min64(a, b U64) U64 {
	if a < b {
		return a
	}
	return b
}

func Min128(a, b U128) U128 {
	if a.Lt(b) {
		return a
	}
	return b
}

// 00 00 00 00 00 00 00 00 | 80 00 00 00 00 00 00 00
func MinI128() I128 {
	return I128{
		U64(0),
		U64(1 << 63),
	}
}

// 00 00 00 00 00 00 00 00 | 00 00 00 00 00 00 00 00
func MinU128() U128 {
	return U128{
		U64(0),
		U64(0),
	}
}

func TrailingZeros128(n U128) uint {
	return n.ToBigInt().TrailingZeroBits()
}

func SaturatingAddU32(a, b U32) U32 {
	sum := uint64(a) + uint64(b)
	if sum > math.MaxUint32 {
		return U32(math.MaxUint32)
	}
	return U32(sum)
}

func SaturatingAddU64(a, b U64) U64 {
	sum, carry := bits.Add64(uint64(a), uint64(b), 0)
	if carry != 0 {
		return U64(math.MaxUint64)
	}
	return U64(sum)
}

func SaturatingAddU128(a, b U128) U128 {
	sumLow, carry := bits.Add64(uint64(a[0]), uint64(b[0]), 0)
	sumHigh, overflow := bits.Add64(uint64(a[1]), uint64(b[1]), carry)
	// check for overflow
	if overflow == 1 || (carry == 1 && sumHigh <= uint64(a[1]) && sumHigh <= uint64(b[1])) {
		return MaxU128()
	}
	return U128{U64(sumLow), U64(sumHigh)}
}

func SaturatingSubU64(a, b U64) U64 {
	diff, borrow := bits.Sub64(uint64(a), uint64(b), 0)
	if borrow != 0 {
		return U64(0)
	}
	return U64(diff)
}

func SaturatingSubU128(a, b U128) U128 {
	low, borrow := bits.Sub64(uint64(a[0]), uint64(b[0]), 0)
	high, _ := bits.Sub64(uint64(a[1]), uint64(b[1]), borrow)
	// check for underflow
	if borrow == 1 || high > uint64(a[1]) {
		return U128{0, 0}
	}
	return U128{U64(low), U64(high)}
}

func SaturatingMulU64(a, b U64) U64 {
	if a == 0 || b == 0 {
		return U64(0)
	}

	hi, lo := bits.Mul64(uint64(a), uint64(b))
	if hi != 0 {
		return U64(math.MaxUint64)
	}
	return U64(lo)
}

func CheckedAddU32(a, b U32) (U32, error) {
	sum, carry := bits.Add32(uint32(a), uint32(b), 0)
	if carry != 0 {
		return 0, errOverflow
	}
	return U32(sum), nil
}

func CheckedAddU64(a, b U64) (U64, error) {
	sum, carry := bits.Add64(uint64(a), uint64(b), 0)
	if carry != 0 {
		return 0, errOverflow
	}
	return U64(sum), nil
}
