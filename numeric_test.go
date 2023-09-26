package goscale

import (
	"math"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_anyIntegerToU128(t *testing.T) {
	u1, _ := NewU128FromString("340282366920938463463374607431768211455")

	testExamples := []struct {
		label       string
		input       any
		expectation U128
	}{
		{"uint(1)=>NewU128(1)", uint(1), NewU128(1)},
		{"uint8(math.MaxUint8)=>NewU128(math.MaxUint8)", uint8(math.MaxUint8), NewU128(math.MaxUint8)},
		{"uint16(math.MaxUint16)=>NewU128(math.MaxUint16)", uint16(math.MaxUint16), NewU128(math.MaxUint16)},
		{"uint32(math.MaxUint32)=>NewU128(math.MaxUint32)", uint32(math.MaxUint32), NewU128(math.MaxUint32)},
		{"uint64(math.MaxUint64)=>NewU128(uint64(math.MaxUint64))", uint64(math.MaxUint64), NewU128(uint64(math.MaxUint64))},
		{"big.NewInt(123456789)=>U128(123456789)", big.NewInt(123456789), NewU128(123456789)},
		{"NewU128(340282366920938463463374607431768211455)=>MaxU128()", u1, MaxU128()},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := anyIntegerTo128Bits[U128](testExample.input)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_anyIntegerToI128(t *testing.T) {
	testExamples := []struct {
		label       string
		input       any
		expectation I128
	}{
		{"int(-1))=>NewI128(-1)", int(-1), NewI128(-1)},
		{"int8(math.MinInt8)=>NewI128(math.MinInt8)", int8(math.MinInt8), NewI128(math.MinInt8)},
		{"int16(math.MinInt16)=>NewI128(math.MinInt16)", int16(math.MinInt16), NewI128(math.MinInt16)},
		{"int32(math.MinInt32)=>NewI128(math.MinInt32)", int32(math.MinInt32), NewI128(math.MinInt32)},
		{"int64(math.MinInt64)=>NewI128(math.MinInt64)", int64(math.MinInt64), NewI128(math.MinInt64)},
		{"big.NewInt(-123456789)=>NewI128(-123456789)", big.NewInt(-123456789), NewI128(-123456789)},
		{"NewI128(-1)=>NewI128(-1)", NewI128(-1), NewI128(-1)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := anyIntegerTo128Bits[I128](testExample.input)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}
