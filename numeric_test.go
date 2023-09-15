package goscale

import (
	"math"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {
	bn1, _ := NewU128FromString("340282366920938463463374607431768211454")
	bn2, _ := NewI128FromString("170141183460469231731687303715884105726")

	testExamples := []struct {
		label       string
		a           Numeric
		b           Numeric
		expectation Numeric
	}{
		{"254+1", U8(254), U8(1), U8(math.MaxUint8)},
		{"126+1", I8(126), I8(1), I8(math.MaxInt8)},
		{"65534+1", U16(65534), U16(1), U16(math.MaxUint16)},
		{"32766+1", I16(32766), I16(1), I16(math.MaxInt16)},
		{"4294967294+1", U32(4294967294), U32(1), U32(math.MaxUint32)},
		{"2147483646+1", I32(2147483646), I32(1), I32(math.MaxInt32)},
		{"18446744073709551614+1", U64(18446744073709551614), U64(1), U64(math.MaxUint64)},
		{"9223372036854775806+1", I64(9223372036854775806), I64(1), I64(math.MaxInt64)},
		{"340282366920938463463374607431768211454+1", bn1, NewU128(1), MaxU128()},
		{"170141183460469231731687303715884105726+1", bn2, NewI128(1), MaxI128()},
		{"255+1", U8(math.MaxUint8), U8(1), U8(0)},
		{"65535+1", U16(math.MaxUint16), U16(1), U16(0)},
		{"4294967295+1", U32(math.MaxUint32), U32(1), U32(0)},
		{"18446744073709551615+1", U64(math.MaxUint64), U64(1), U64(0)},
		{"340282366920938463463374607431768211455+1", MaxU128(), NewU128(1), NewU128(0)},
		{"-2+(-1)", NewI128(-2), NewI128(-1), NewI128(-3)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Add(testExample.b)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_Sub(t *testing.T) {
	u1, _ := NewU128FromString("340282366920938463463374607431657675311")
	u2, _ := NewU128FromString("340282366920938463463374607421657675311")
	u3, _ := NewI128FromString("340282366920938463463374607431657675311")
	u4, _ := NewI128FromString("340282366920938463463374607421657675311")
	u5, _ := NewI128FromString("340282366920938463463374607431657675310")

	testExamples := []struct {
		label       string
		a           Numeric
		b           Numeric
		expectation Numeric
	}{
		{"2-1", U8(2), U8(1), U8(1)},
		{"0-1", U8(0), U8(1), U8(math.MaxUint8)},
		{"2-1", I8(2), I8(1), I8(1)},
		{"-128-1", I8(math.MinInt8), I8(1), I8(math.MaxInt8)},
		{"2-1", U16(2), U16(1), U16(1)},
		{"0-1", U16(0), U16(1), U16(math.MaxUint16)},
		{"2-1", I16(2), I16(1), I16(1)},
		{"-32768-1", I16(math.MinInt16), I16(1), I16(math.MaxInt16)},
		{"0-1", U32(0), U32(1), U32(math.MaxUint32)},
		{"2-1", U32(2), U32(1), U32(1)},
		{"2-1", I32(2), I32(1), I32(1)},
		{"-2147483648-1", I32(math.MinInt32), I32(1), I32(math.MaxInt32)},
		{"2-1", U64(2), U64(1), U64(1)},
		{"0-1", U64(0), U64(1), U64(math.MaxUint64)},
		{"2-1", I64(2), I64(1), I64(1)},
		{"-9223372036854775808-1", I64(math.MinInt64), I64(1), I64(math.MaxInt64)},
		{"2-1", NewU128(2), NewU128(1), NewU128(1)},
		{"0-1", NewU128(0), NewU128(1), MaxU128()},
		{"2-1", NewI128(2), NewI128(1), NewI128(1)},
		{"-2-(-1)", NewI128(-2), NewI128(-1), NewI128(-1)},
		{"-170141183460469231731687303715884105728-1", MinI128(), NewI128(1), MaxI128()},
		{"499999889463855-10000000000", NewU128(499999889463855), NewU128(10000000000), NewU128(499989889463855)},
		{"499999889463855-10000000000", NewI128(499999889463855), NewI128(10000000000), NewI128(499989889463855)},
		{"340282366920938463463374607431657675311-10000000000", u1, NewU128(10000000000), u2},
		{"340282366920938463463374607431657675311-10000000000", u3, NewI128(10000000000), u4},
		{"9889463854-10000000000", NewI128(9889463854), NewI128(10000000000), u5},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Sub(testExample.b)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_Mul(t *testing.T) {
	testExamples := []struct {
		label       string
		a           Numeric
		b           Numeric
		expectation Numeric
	}{
		{"2*3", U8(2), U8(3), U8(6)},
		{"MaxU8*0", U8(math.MaxUint8), U8(0), U8(0)},
		{"MaxU8*1", U8(math.MaxUint8), U8(1), U8(math.MaxUint8)},
		{"MaxU8*2", U8(math.MaxUint8), U8(2), U8(math.MaxUint8).Sub(U8(1))},
		{"MaxU8*MaxU8", U8(math.MaxUint8), U8(math.MaxUint8), U8(1)},

		{"2*3", I8(2), I8(3), I8(6)},
		{"MaxI8*0", I8(math.MaxInt8), I8(0), I8(0)},
		{"MaxI8*1", I8(math.MaxInt8), I8(1), I8(math.MaxInt8)},
		{"MaxI8*2", I8(math.MaxInt8), I8(2), I8(-2)},
		{"MaxI8*MaxI8", I8(math.MaxInt8), I8(math.MaxInt8), I8(1)},
		{"MinI8*0", I8(math.MinInt8), I8(0), I8(0)},
		{"MinI8*1", I8(math.MinInt8), I8(1), I8(math.MinInt8)},
		{"MinI8*2", I8(math.MinInt8), I8(2), I8(0)},
		{"MinI8*MinI8", I8(math.MinInt8), I8(math.MinInt8), I8(0)},

		{"2*3", U16(2), U16(3), U16(6)},
		{"MaxU16*0", U16(math.MaxUint16), U16(0), U16(0)},
		{"MaxU16*1", U16(math.MaxUint16), U16(1), U16(math.MaxUint16)},
		{"MaxU16*2", U16(math.MaxUint16), U16(2), U16(math.MaxUint16).Sub(U16(1))},
		{"MaxU16*MaxU16", U16(math.MaxUint16), U16(math.MaxUint16), U16(1)},

		{"2*3", I16(2), I16(3), I16(6)},
		{"MaxI16*0", I16(math.MaxInt16), I16(0), I16(0)},
		{"MaxI16*1", I16(math.MaxInt16), I16(1), I16(math.MaxInt16)},
		{"MaxI16*2", I16(math.MaxInt16), I16(2), I16(-2)},
		{"MaxI16*MaxI16", I16(math.MaxInt16), I16(math.MaxInt16), I16(1)},
		{"MinI16*0", I16(math.MinInt16), I16(0), I16(0)},
		{"MinI16*1", I16(math.MinInt16), I16(1), I16(math.MinInt16)},
		{"MinI16*2", I16(math.MinInt16), I16(2), I16(0)},
		{"MinI16*MinI16", I16(math.MinInt16), I16(math.MinInt16), I16(0)},

		{"2*3", U32(2), U32(3), U32(6)},
		{"MaxU32*0", U32(math.MaxUint32), U32(0), U32(0)},
		{"MaxU32*1", U32(math.MaxUint32), U32(1), U32(math.MaxUint32)},
		{"MaxU32*2", U32(math.MaxUint32), U32(2), U32(math.MaxUint32).Sub(U32(1))},
		{"MaxU32*MaxU32", U32(math.MaxUint32), U32(math.MaxUint32), U32(1)},

		{"2*3", I32(2), I32(3), I32(6)},
		{"MaxI32*0", I32(math.MaxInt32), I32(0), I32(0)},
		{"MaxI32*1", I32(math.MaxInt32), I32(1), I32(math.MaxInt32)},
		{"MaxI32*2", I32(math.MaxInt32), I32(2), I32(-2)},
		{"MaxI32*MaxI32", I32(math.MaxInt32), I32(math.MaxInt32), I32(1)},
		{"MinI32*0", I32(math.MinInt32), I32(0), I32(0)},
		{"MinI32*1", I32(math.MinInt32), I32(1), I32(math.MinInt32)},
		{"MinI32*2", I32(math.MinInt32), I32(2), I32(0)},
		{"MinI32*MinI32", I32(math.MinInt32), I32(math.MinInt32), I32(0)},

		{"2*3", U64(2), U64(3), U64(6)},
		{"MaxU64*0", U64(math.MaxUint64), U64(0), U64(0)},
		{"MaxU64*1", U64(math.MaxUint64), U64(1), U64(math.MaxUint64)},
		{"MaxU64*2", U64(math.MaxUint64), U64(2), U64(math.MaxUint64).Sub(U64(1))},
		{"MaxU64*MaxU64", U64(math.MaxUint64), U64(math.MaxUint64), U64(1)},

		{"2*3", I64(2), I64(3), I64(6)},
		{"MaxI64*0", I64(math.MaxInt64), I64(0), I64(0)},
		{"MaxI64*1", I64(math.MaxInt64), I64(1), I64(math.MaxInt64)},
		{"MaxI64*2", I64(math.MaxInt64), I64(2), I64(-2)},
		{"MaxI64*MaxI64", I64(math.MaxInt64), I64(math.MaxInt64), I64(1)},
		{"MinI64*0", I64(math.MinInt64), I64(0), I64(0)},
		{"MinI64*1", I64(math.MinInt64), I64(1), I64(math.MinInt64)},
		{"MinI64*2", I64(math.MinInt64), I64(2), I64(0)},
		{"MinI64*MinI64", I64(math.MinInt64), I64(math.MinInt64), I64(0)},

		{"2*3", NewU128(2), NewU128(3), NewU128(6)},
		{"MaxU128*0", MaxU128(), NewU128(0), NewU128(0)},
		{"MaxU128*1", MaxU128(), NewU128(1), MaxU128()},
		{"MaxU128*2", MaxU128(), NewU128(2), MaxU128().Sub(NewU128(1))},
		{"MaxU128*MaxU128", MaxU128(), MaxU128(), NewU128(1)},

		{"MaxI128*0", MaxI128(), NewI128(0), NewI128(0)},
		{"MaxI128*1", MaxI128(), NewI128(1), MaxI128()},
		{"MaxI128*2", MaxI128(), NewI128(2), NewI128(-2)},
		{"MaxI128*MaxI128", MaxI128(), MaxI128(), NewI128(1)},

		{"-2*3", NewI128(-2), NewI128(3), NewI128(-6)},
		{"-2*-3", NewI128(-2), NewI128(-3), NewI128(6)},
		{"1*0", NewI128(1), NewI128(0), NewI128(0)},
		{"-1*0", NewI128(-1), NewI128(0), NewI128(0)},
		{"1*1", NewI128(1), NewI128(1), NewI128(1)},
		{"-1*1", NewI128(-1), NewI128(1), NewI128(-1)},
		{"-1*-1", NewI128(-1), NewI128(-1), NewI128(1)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Mul(testExample.b)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_Div(t *testing.T) {
	testExamples := []struct {
		label       string
		a           Numeric
		b           Numeric
		expectation Numeric
	}{
		{"18446744073709551615/1", U64(math.MaxUint64), U64(1), U64(math.MaxUint64)},

		{"0/1", NewU128(0), NewU128(1), NewU128(0)},
		{"1/1", NewU128(1), NewU128(1), NewU128(1)},

		{"0/-1", NewI128(0), NewI128(-1), NewI128(0)},
		{"-1/1", NewI128(-1), NewI128(1), NewI128(-1)},
		{"-1/-1", NewI128(-1), NewI128(-1), NewI128(1)},

		{"6/3", NewU128(6), NewU128(3), NewU128(2)},
		{"-6/2", NewI128(-6), NewI128(2), NewI128(-3)},

		{"32/4", U64(32), U64(4), U64(8)},
		{"-32/4", I64(-32), I64(4), I64(-8)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Div(testExample.b)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_Mod(t *testing.T) {
	testExamples := []struct {
		label       string
		a           Numeric
		b           Numeric
		expectation Numeric
	}{
		{"1%1", U8(1), U8(1), U8(0)},
		{"1%2", U8(1), U8(2), U8(1)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Mod(testExample.b)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_Max(t *testing.T) {
	testExamples := []struct {
		label       string
		a           Numeric
		b           Numeric
		expectation Numeric
	}{
		{"Max(1, 2)", U8(1), U8(2), U8(2)},
		{"Max(1, MaxU8)", U8(1), U8(math.MaxUint8), U8(math.MaxUint8)},
		{"Max(MaxU8, MaxU8)", U8(math.MaxUint8), U8(math.MaxUint8), U8(math.MaxUint8)},
		{"Max(-2, -1)", I8(-2), I8(-1), I8(-1)},
		{"Max(-2, MaxI8)", I8(-2), I8(math.MaxInt8), I8(math.MaxInt8)},
		{"Max(MaxI8, MaxI8)", I8(math.MaxInt8), I8(math.MaxInt8), I8(math.MaxInt8)},
		{"Max(1, 2)", U16(1), U16(2), U16(2)},
		{"Max(1, MaxU16)", U16(1), U16(math.MaxUint16), U16(math.MaxUint16)},
		{"Max(MaxU16, MaxU16)", U16(math.MaxUint16), U16(math.MaxUint16), U16(math.MaxUint16)},
		{"Max(-2, -1)", I16(-2), I16(-1), I16(-1)},
		{"Max(-2, MaxI16)", I16(-2), I16(math.MaxInt16), I16(math.MaxInt16)},
		{"Max(MaxI16, MaxI16)", I16(math.MaxInt16), I16(math.MaxInt16), I16(math.MaxInt16)},
		{"Max(1, 2)", U32(1), U32(2), U32(2)},
		{"Max(1, MaxU32)", U32(1), U32(math.MaxUint32), U32(math.MaxUint32)},
		{"Max(MaxU32, MaxU32)", U32(math.MaxUint32), U32(math.MaxUint32), U32(math.MaxUint32)},
		{"Max(-2, -1)", I32(-2), I32(-1), I32(-1)},
		{"Max(-2, MaxI32)", I32(-2), I32(math.MaxInt32), I32(math.MaxInt32)},
		{"Max(MaxI32, MaxI32)", I32(math.MaxInt32), I32(math.MaxInt32), I32(math.MaxInt32)},
		{"Max(1, 2)", U64(1), U64(2), U64(2)},
		{"Max(1, MaxU64)", U64(1), U64(math.MaxUint64), U64(math.MaxUint64)},
		{"Max(MaxU64, MaxU64)", U64(math.MaxUint64), U64(math.MaxUint64), U64(math.MaxUint64)},
		{"Max(-2, -1)", I64(-2), I64(-1), I64(-1)},
		{"Max(-2, MaxI64)", I64(-2), I64(math.MaxInt64), I64(math.MaxInt64)},
		{"Max(MaxI64, MaxI64)", I64(math.MaxInt64), I64(math.MaxInt64), I64(math.MaxInt64)},
		{"Max(1, 2)", NewU128(1), NewU128(2), NewU128(2)},
		{"Max(1, MaxU128)", NewU128(1), MaxU128(), MaxU128()},
		{"Max(MaxU128, MaxU128)", MaxU128(), MaxU128(), MaxU128()},
		{"Max(-2, -1)", NewI128(-2), NewI128(-1), NewI128(-1)},
		{"Max(-2, MaxI128)", NewI128(-2), MaxI128(), MaxI128()},
		{"Max(MaxI128, MaxI128)", MaxI128(), MaxI128(), MaxI128()},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Max(testExample.b)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_Min(t *testing.T) {
	testExamples := []struct {
		label       string
		a           Numeric
		b           Numeric
		expectation Numeric
	}{
		{"Min(1, 2)", U8(1), U8(2), U8(1)},
		{"Min(1, MaxU8)", U8(1), U8(math.MaxUint8), U8(1)},
		{"Min(MaxU8, MaxU8)", U8(math.MaxUint8), U8(math.MaxUint8), U8(math.MaxUint8)},
		{"Min(-2, -1)", I8(-2), I8(-1), I8(-2)},
		{"Min(-2, MaxI8)", I8(-2), I8(math.MaxInt8), I8(-2)},
		{"Min(1, 2)", U16(1), U16(2), U16(1)},
		{"Min(1, MaxU16)", U16(1), U16(math.MaxUint16), U16(1)},
		{"Min(MaxU16, MaxU16)", U16(math.MaxUint16), U16(math.MaxUint16), U16(math.MaxUint16)},
		{"Min(-2, -1)", I16(-2), I16(-1), I16(-2)},
		{"Min(-2, MaxI16)", I16(-2), I16(math.MaxInt16), I16(-2)},
		{"Min(1, 2)", U32(1), U32(2), U32(1)},
		{"Min(1, MaxU32)", U32(1), U32(math.MaxUint32), U32(1)},
		{"Min(MaxU32, MaxU32)", U32(math.MaxUint32), U32(math.MaxUint32), U32(math.MaxUint32)},
		{"Min(-2, -1)", I32(-2), I32(-1), I32(-2)},
		{"Min(-2, MaxI32)", I32(-2), I32(math.MaxInt32), I32(-2)},
		{"Min(1, 2)", U64(1), U64(2), U64(1)},
		{"Min(1, MaxU64)", U64(1), U64(math.MaxUint64), U64(1)},
		{"Min(MaxU64, MaxU64)", U64(math.MaxUint64), U64(math.MaxUint64), U64(math.MaxUint64)},
		{"Min(-2, -1)", I64(-2), I64(-1), I64(-2)},
		{"Min(-2, MaxI64)", I64(-2), I64(math.MaxInt64), I64(-2)},
		{"Min(1, 2)", NewU128(1), NewU128(2), NewU128(1)},
		{"Min(1, MaxU128)", NewU128(1), MaxU128(), NewU128(1)},
		{"Min(MaxU128, MaxU128)", MaxU128(), MaxU128(), MaxU128()},
		{"Min(-2, -1)", NewI128(-2), NewI128(-1), NewI128(-2)},
		{"Min(-2, MaxI128)", NewI128(-2), MaxI128(), NewI128(-2)},
		{"Min(MinU128, MaxU128)", MinU128(), MaxU128(), MinU128()},
		{"Min(MinI128, MaxI128)", MinI128(), MaxI128(), MinI128()},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Min(testExample.b)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_TrailingZeros(t *testing.T) {
	testExamples := []struct {
		label       string
		a           Numeric
		expectation Numeric
	}{
		{"TrailingZeros(1)", U8(1), U8(0)},
		{"TrailingZeros(2)", U8(2), U8(1)},
		{"TrailingZeros(3)", U8(3), U8(0)},
		{"TrailingZeros(1)", I8(1), I8(0)},
		{"TrailingZeros(2)", I8(2), I8(1)},
		{"TrailingZeros(3)", I8(3), I8(0)},
		{"TrailingZeros(1)", U16(1), U16(0)},
		{"TrailingZeros(2)", U16(2), U16(1)},
		{"TrailingZeros(3)", U16(3), U16(0)},
		{"TrailingZeros(1)", I16(1), I16(0)},
		{"TrailingZeros(2)", I16(2), I16(1)},
		{"TrailingZeros(3)", I16(3), I16(0)},
		{"TrailingZeros(1)", U32(1), U32(0)},
		{"TrailingZeros(2)", U32(2), U32(1)},
		{"TrailingZeros(3)", U32(3), U32(0)},
		{"TrailingZeros(1)", I32(1), I32(0)},
		{"TrailingZeros(2)", I32(2), I32(1)},
		{"TrailingZeros(3)", I32(3), I32(0)},
		{"TrailingZeros(1)", U64(1), U64(0)},
		{"TrailingZeros(2)", U64(2), U64(1)},
		{"TrailingZeros(3)", U64(3), U64(0)},
		{"TrailingZeros(1)", I64(1), I64(0)},
		{"TrailingZeros(2)", I64(2), I64(1)},
		{"TrailingZeros(3)", I64(3), I64(0)},
		{"TrailingZeros(1)", NewU128(1), NewU128(0)},
		{"TrailingZeros(2)", NewU128(2), NewU128(1)},
		{"TrailingZeros(3)", NewU128(3), NewU128(0)},
		{"TrailingZeros(1)", NewI128(1), NewI128(0)},
		{"TrailingZeros(2)", NewI128(2), NewI128(1)},
		{"TrailingZeros(3)", NewI128(3), NewI128(0)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.TrailingZeros()
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_Clamp(t *testing.T) {
	testExamples := []struct {
		label       string
		a           Numeric
		minValue    Numeric
		maxValue    Numeric
		expectation Numeric
	}{
		{"Clamp(5, 1, 10)", U8(5), U8(1), U8(10), U8(5)},
		{"Clamp(15, 1, 10)", U8(15), U8(1), U8(10), U8(10)},
		{"Clamp(0, 1, 10)", U8(0), U8(1), U8(10), U8(1)},
		{"Clamp(-3, -2, 2)", I8(-3), I8(-2), I8(2), I8(-2)},
		{"Clamp(3, -2, 2)", I8(3), I8(-2), I8(2), I8(2)},
		{"Clamp(1, -2, 2)", I8(1), I8(-2), I8(2), I8(1)},
		{"Clamp(5, 1, 10)", U16(5), U16(1), U16(10), U16(5)},
		{"Clamp(15, 1, 10)", U16(15), U16(1), U16(10), U16(10)},
		{"Clamp(0, 1, 10)", U16(0), U16(1), U16(10), U16(1)},
		{"Clamp(-3, -2, 2)", I16(-3), I16(-2), I16(2), I16(-2)},
		{"Clamp(3, -2, 2)", I16(3), I16(-2), I16(2), I16(2)},
		{"Clamp(1, -2, 2)", I16(1), I16(-2), I16(2), I16(1)},
		{"Clamp(5, 1, 10)", U32(5), U32(1), U32(10), U32(5)},
		{"Clamp(15, 1, 10)", U32(15), U32(1), U32(10), U32(10)},
		{"Clamp(0, 1, 10)", U32(0), U32(1), U32(10), U32(1)},
		{"Clamp(-3, -2, 2)", I32(-3), I32(-2), I32(2), I32(-2)},
		{"Clamp(3, -2, 2)", I32(3), I32(-2), I32(2), I32(2)},
		{"Clamp(1, -2, 2)", I32(1), I32(-2), I32(2), I32(1)},
		{"Clamp(5, 1, 10)", U64(5), U64(1), U64(10), U64(5)},
		{"Clamp(15, 1, 10)", U64(15), U64(1), U64(10), U64(10)},
		{"Clamp(0, 1, 10)", U64(0), U64(1), U64(10), U64(1)},
		{"Clamp(-3, -2, 2)", I64(-3), I64(-2), I64(2), I64(-2)},
		{"Clamp(3, -2, 2)", I64(3), I64(-2), I64(2), I64(2)},
		{"Clamp(1, -2, 2)", I64(1), I64(-2), I64(2), I64(1)},
		{"Clamp(5, 1, 10)", NewU128(5), NewU128(1), NewU128(10), NewU128(5)},
		{"Clamp(15, 1, 10)", NewU128(15), NewU128(1), NewU128(10), NewU128(10)},
		{"Clamp(0, 1, 10)", NewU128(0), NewU128(1), NewU128(10), NewU128(1)},
		{"Clamp(-3, -2, 2)", NewI128(-3), NewI128(-2), NewI128(2), NewI128(-2)},
		{"Clamp(3, -2, 2)", NewI128(3), NewI128(-2), NewI128(2), NewI128(2)},
		{"Clamp(1, -2, 2)", NewI128(1), NewI128(-2), NewI128(2), NewI128(1)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Clamp(testExample.minValue, testExample.maxValue)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_SaturatingAdd(t *testing.T) {
	testExamples := []struct {
		label       string
		a           Numeric
		b           Numeric
		expectation Numeric
	}{
		{"MaxU8+1", U8(math.MaxUint8), U8(1), U8(math.MaxUint8)},
		{"MaxU8+MaxU8", U8(math.MaxUint8), U8(math.MaxUint8), U8(math.MaxUint8)},
		{"MaxI8+1", I8(math.MaxInt8), I8(1), I8(math.MaxInt8)},
		{"MaxI8+MaxI8", I8(math.MaxInt8), I8(math.MaxInt8), I8(math.MaxInt8)},
		{"MaxU16+1", U16(math.MaxUint16), U16(1), U16(math.MaxUint16)},
		{"MaxU16+MaxU16", U16(math.MaxUint16), U16(math.MaxUint16), U16(math.MaxUint16)},
		{"MaxI16+1", I16(math.MaxInt16), I16(1), I16(math.MaxInt16)},
		{"MaxI16+MaxI16", I16(math.MaxInt16), I16(math.MaxInt16), I16(math.MaxInt16)},
		{"MaxU32+1", U32(math.MaxUint32), U32(1), U32(math.MaxUint32)},
		{"MaxU32+MaxU32", U32(math.MaxUint32), U32(math.MaxUint32), U32(math.MaxUint32)},
		{"MaxI32+1", I32(math.MaxInt32), I32(1), I32(math.MaxInt32)},
		{"MaxI32+MaxI32", I32(math.MaxInt32), I32(math.MaxInt32), I32(math.MaxInt32)},
		{"MaxU64+1", U64(math.MaxUint64), U64(1), U64(math.MaxUint64)},
		{"MaxU64+MaxU64", U64(math.MaxUint64), U64(math.MaxUint64), U64(math.MaxUint64)},
		{"MaxI64+1", I64(math.MaxInt64), I64(1), I64(math.MaxInt64)},
		{"MaxI64+MaxI64", I64(math.MaxInt64), I64(math.MaxInt64), I64(math.MaxInt64)},
		{"MaxU128+1", MaxU128(), NewU128(1), MaxU128()},
		{"MaxU128+MaxU128", MaxU128(), MaxU128(), MaxU128()},
		{"MaxI128+1", MaxI128(), NewI128(1), MaxI128()},
		{"MaxI128+MaxI128", MaxI128(), MaxI128(), MaxI128()},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.SaturatingAdd(testExample.b)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_SaturatingSub(t *testing.T) {
	testExamples := []struct {
		label       string
		a           Numeric
		b           Numeric
		expectation Numeric
	}{
		{"0-1", U8(0), U8(1), U8(0)},
		{"0-MaxU8", U8(0), U8(math.MaxUint8), U8(0)},
		{"0-1", I8(0), I8(1), I8(-1)},
		{"MinU8-1", I8(math.MinInt8), I8(1), I8(math.MinInt8)},
		{"0-1", U16(0), U16(1), U16(0)},
		{"0-MaxU16", U16(0), U16(math.MaxUint16), U16(0)},
		{"0-1", I16(0), I16(1), I16(-1)},
		{"MinU16-1", I16(math.MinInt16), I16(1), I16(math.MinInt16)},
		{"0-1", U32(0), U32(1), U32(0)},
		{"0-MaxU32", U32(0), U32(math.MaxUint32), U32(0)},
		{"0-1", I32(0), I32(1), I32(-1)},
		{"MinU32-1", I32(math.MinInt32), I32(1), I32(math.MinInt32)},
		{"0-1", U64(0), U64(1), U64(0)},
		{"0-MaxU64", U64(0), U64(math.MaxUint64), U64(0)},
		{"0-1", I64(0), I64(1), I64(-1)},
		{"MinI64-1", I64(math.MinInt64), I64(1), I64(math.MinInt64)},
		{"0-1", NewU128(0), NewU128(1), NewU128(0)},
		{"0-MaxU128", NewU128(0), MaxU128(), NewU128(0)},
		{"0-1", NewI128(0), NewI128(1), NewI128(-1)},
		{"MinU128-1", MinI128(), NewI128(1), MinI128()},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.SaturatingSub(testExample.b)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_SaturatingMul(t *testing.T) {
	testExamples := []struct {
		label       string
		a           Numeric
		b           Numeric
		expectation Numeric
	}{
		{"MaxU8*2", U8(math.MaxUint8), U8(2), U8(math.MaxUint8)},
		{"MaxU8*MaxU8", U8(math.MaxUint8), U8(math.MaxUint8), U8(math.MaxUint8)},
		{"MaxI8*2", I8(math.MaxInt8), I8(2), I8(math.MaxInt8)},
		{"MinI8*2", I8(math.MinInt8), I8(2), I8(math.MinInt8)},
		{"MaxU16*2", U16(math.MaxUint16), U16(2), U16(math.MaxUint16)},
		{"MaxU16*MaxU16", U16(math.MaxUint16), U16(math.MaxUint16), U16(math.MaxUint16)},
		{"MaxI16*2", I16(math.MaxInt16), I16(2), I16(math.MaxInt16)},
		{"MinI16*2", I16(math.MinInt16), I16(2), I16(math.MinInt16)},
		{"MaxU32*2", U32(math.MaxUint32), U32(2), U32(math.MaxUint32)},
		{"MaxU32*MaxU32", U32(math.MaxUint32), U32(math.MaxUint32), U32(math.MaxUint32)},
		{"MaxI32*2", I32(math.MaxInt32), I32(2), I32(math.MaxInt32)},
		{"MinI32*2", I32(math.MinInt32), I32(2), I32(math.MinInt32)},
		{"MaxU64*2", U64(math.MaxUint64), U64(2), U64(math.MaxUint64)},
		{"MaxU64*MaxU64", U64(math.MaxUint64), U64(math.MaxUint64), U64(math.MaxUint64)},
		{"MaxI64*2", I64(math.MaxInt64), I64(2), I64(math.MaxInt64)},
		{"MinI64*2", I64(math.MinInt64), I64(2), I64(math.MinInt64)},
		{"MaxU128*2", MaxU128(), NewU128(2), MaxU128()},
		{"MaxU128*MaxU128", MaxU128(), MaxU128(), MaxU128()},
		{"MaxI128*2", MaxI128(), NewI128(2), MaxI128()},
		{"MinI128*2", MinI128(), NewI128(2), MinI128()},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.SaturatingMul(testExample.b)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_CheckedAdd(t *testing.T) {
	testExamples := []struct {
		label          string
		a              U64
		b              U64
		expectation    U64
		expectationErr bool
	}{
		{"1+2", U64(1), U64(2), U64(3), false},
		{"MaxU64+1", U64(math.MaxUint64), U64(1), U64(0), true},
		{"MaxU64+MaxU64", U64(math.MaxUint64), U64(math.MaxUint64), U64(0), true},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result, err := testExample.a.CheckedAdd(testExample.b)
			assert.Equal(t, testExample.expectation, result)
			if testExample.expectationErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_Eq(t *testing.T) {
	testExamples := []struct {
		label       string
		a           Numeric
		b           Numeric
		expectation bool
	}{
		{"1==1", U8(1), U8(1), true},
		{"1==2", U8(1), U8(2), false},

		{"1==1", I8(1), I8(1), true},
		{"1==2", I8(1), I8(2), false},
		{"-1==1", I8(-1), I8(1), false},
		{"-1==-1", I8(-1), I8(-1), true},

		{"1==1", U16(1), U16(1), true},
		{"1==2", U16(1), U16(2), false},

		{"1==1", I16(1), I16(1), true},
		{"1==2", I16(1), I16(2), false},
		{"-1==1", I16(-1), I16(1), false},
		{"-1==-1", I16(-1), I16(-1), true},

		{"1==1", U32(1), U32(1), true},
		{"1==2", U32(1), U32(2), false},

		{"1==1", I32(1), I32(1), true},
		{"1==2", I32(1), I32(2), false},
		{"-1==1", I32(-1), I32(1), false},
		{"-1==-1", I32(-1), I32(-1), true},

		{"1==1", U64(1), U64(1), true},
		{"1==2", U64(1), U64(2), false},

		{"1==1", I64(1), I64(1), true},
		{"1==2", I64(1), I64(2), false},
		{"-1==1", I64(-1), I64(1), false},
		{"-1==-1", I64(-1), I64(-1), true},

		{"1==1", NewU128(1), NewU128(1), true},
		{"1==2", NewU128(1), NewU128(2), false},

		{"1==1", NewI128(1), NewI128(1), true},
		{"1==2", NewI128(1), NewI128(2), false},
		{"-1==1", NewI128(-1), NewI128(1), false},
		{"-1==-1", NewI128(-1), NewI128(-1), true},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Eq(testExample.b)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_Lt(t *testing.T) {
	testExamples := []struct {
		label       string
		a           Numeric
		b           Numeric
		expectation bool
	}{
		{"1<1", U8(1), U8(1), false},
		{"1<1", U8(1), U8(1), false},
		{"1<2", U8(1), U8(2), true},

		{"1<1", I8(1), I8(1), false},
		{"1<2", I8(1), I8(2), true},
		{"-1<1", I8(-1), I8(1), true},
		{"-1<-1", I8(-1), I8(-1), false},

		{"1<1", U16(1), U16(1), false},
		{"1<2", U16(1), U16(2), true},

		{"1<1", I16(1), I16(1), false},
		{"1<2", I16(1), I16(2), true},
		{"-1<1", I16(-1), I16(1), true},
		{"-1<-1", I16(-1), I16(-1), false},

		{"1<1", U32(1), U32(1), false},
		{"1<2", U32(1), U32(2), true},

		{"1<1", I32(1), I32(1), false},
		{"1<2", I32(1), I32(2), true},
		{"-1<1", I32(-1), I32(1), true},
		{"-1<-1", I32(-1), I32(-1), false},

		{"1<1", U64(1), U64(1), false},
		{"1<2", U64(1), U64(2), true},

		{"1<1", I64(1), I64(1), false},
		{"1<2", I64(1), I64(2), true},
		{"-1<1", I64(-1), I64(1), true},
		{"-1<-1", I64(-1), I64(-1), false},

		{"1<1", NewU128(1), NewU128(1), false},
		{"1<2", NewU128(1), NewU128(2), true},

		{"1<1", NewI128(1), NewI128(1), false},
		{"1<2", NewI128(1), NewI128(2), true},
		{"-1<1", NewI128(-1), NewI128(1), true},
		{"-1<-1", NewI128(-1), NewI128(-1), false},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Lt(testExample.b)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_Lte(t *testing.T) {
	testExamples := []struct {
		label       string
		a           Numeric
		b           Numeric
		expectation bool
	}{
		{"1<=1", U8(1), U8(1), true},
		{"1<=2", U8(1), U8(2), true},
		{"1<=0", U8(1), U8(0), false},

		{"1<=1", I8(1), I8(1), true},
		{"1<=2", I8(1), I8(2), true},
		{"-1<=1", I8(-1), I8(1), true},
		{"-1<=-1", I8(-1), I8(-1), true},
		{"-1<=-2", I8(-1), I8(-2), false},

		{"1<=1", U16(1), U16(1), true},
		{"1<=2", U16(1), U16(2), true},
		{"1<=0", U16(1), U16(0), false},

		{"1<=1", I16(1), I16(1), true},
		{"1<=2", I16(1), I16(2), true},
		{"-1<=1", I16(-1), I16(1), true},
		{"-1<=-1", I16(-1), I16(-1), true},
		{"-1<=-2", I16(-1), I16(-2), false},

		{"1<=1", U32(1), U32(1), true},
		{"1<=2", U32(1), U32(2), true},
		{"1<=0", U32(1), U32(0), false},

		{"1<=1", I32(1), I32(1), true},
		{"1<=2", I32(1), I32(2), true},
		{"-1<=1", I32(-1), I32(1), true},
		{"-1<=-1", I32(-1), I32(-1), true},
		{"-1<=-2", I32(-1), I32(-2), false},

		{"1<=1", U64(1), U64(1), true},
		{"1<=2", U64(1), U64(2), true},
		{"1<=0", U64(1), U64(0), false},

		{"1<=1", I64(1), I64(1), true},
		{"1<=2", I64(1), I64(2), true},
		{"-1<=1", I64(-1), I64(1), true},
		{"-1<=-1", I64(-1), I64(-1), true},
		{"-1<=-2", I64(-1), I64(-2), false},

		{"1<=1", NewU128(1), NewU128(1), true},
		{"1<=2", NewU128(1), NewU128(2), true},
		{"1<=0", NewU128(1), NewU128(0), false},

		{"1<=1", NewI128(1), NewI128(1), true},
		{"1<=2", NewI128(1), NewI128(2), true},
		{"-1<=1", NewI128(-1), NewI128(1), true},
		{"-1<=-1", NewI128(-1), NewI128(-1), true},
		{"-1<=-2", NewI128(-1), NewI128(-2), false},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Lte(testExample.b)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_Gt(t *testing.T) {
	testExamples := []struct {
		label       string
		a           Numeric
		b           Numeric
		expectation bool
	}{
		{"1>1", U8(1), U8(1), false},
		{"1>2", U8(1), U8(2), false},
		{"2>1", U8(2), U8(1), true},

		{"1>1", I8(1), I8(1), false},
		{"1>2", I8(1), I8(2), false},
		{"1>-1", I8(1), I8(-1), true},
		{"-1>-1", I8(-1), I8(-1), false},

		{"1>1", U16(1), U16(1), false},
		{"1>2", U16(1), U16(2), false},
		{"2>1", U16(2), U16(1), true},

		{"1>1", I16(1), I16(1), false},
		{"1>2", I16(1), I16(2), false},
		{"1>-1", I16(1), I16(-1), true},
		{"-1>-1", I16(-1), I16(-1), false},

		{"1>1", U32(1), U32(1), false},
		{"1>2", U32(1), U32(2), false},
		{"2>1", U32(2), U32(1), true},

		{"1>1", I32(1), I32(1), false},
		{"1>2", I32(1), I32(2), false},
		{"1>-1", I32(1), I32(-1), true},
		{"-1>-1", I32(-1), I32(-1), false},

		{"1>1", U64(1), U64(1), false},
		{"1>2", U64(1), U64(2), false},
		{"2>1", U64(2), U64(1), true},

		{"1>1", I64(1), I64(1), false},
		{"1>2", I64(1), I64(2), false},
		{"1>-1", I64(1), I64(-1), true},
		{"-1>-1", I64(-1), I64(-1), false},

		{"1>1", NewU128(1), NewU128(1), false},
		{"1>2", NewU128(1), NewU128(2), false},
		{"2>1", NewU128(2), NewU128(1), true},

		{"1>1", NewI128(1), NewI128(1), false},
		{"1>2", NewI128(1), NewI128(2), false},
		{"1>-1", NewI128(1), NewI128(-1), true},
		{"-1>-1", NewI128(-1), NewI128(-1), false},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Gt(testExample.b)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_Gte(t *testing.T) {
	testExamples := []struct {
		label       string
		a           Numeric
		b           Numeric
		expectation bool
	}{
		{"1>=1", U8(1), U8(1), true},
		{"1>=2", U8(1), U8(2), false},
		{"2>=1", U8(2), U8(1), true},

		{"1>=1", I8(1), I8(1), true},
		{"1>=2", I8(1), I8(2), false},
		{"1>=-1", I8(1), I8(-1), true},
		{"-1>=-1", I8(-1), I8(-1), true},

		{"1>=1", U16(1), U16(1), true},
		{"1>=2", U16(1), U16(2), false},
		{"2>=1", U16(2), U16(1), true},

		{"1>=1", I16(1), I16(1), true},
		{"1>=2", I16(1), I16(2), false},
		{"1>=-1", I16(1), I16(-1), true},
		{"-1>=-1", I16(-1), I16(-1), true},

		{"1>=1", U32(1), U32(1), true},
		{"1>=2", U32(1), U32(2), false},
		{"2>=1", U32(2), U32(1), true},

		{"1>=1", I32(1), I32(1), true},
		{"1>=2", I32(1), I32(2), false},
		{"1>=-1", I32(1), I32(-1), true},
		{"-1>=-1", I32(-1), I32(-1), true},

		{"1>=1", U64(1), U64(1), true},
		{"1>=2", U64(1), U64(2), false},
		{"2>=1", U64(2), U64(1), true},

		{"1>=1", I64(1), I64(1), true},
		{"1>=2", I64(1), I64(2), false},
		{"1>=-1", I64(1), I64(-1), true},
		{"-1>=-1", I64(-1), I64(-1), true},

		{"1>=1", NewU128(1), NewU128(1), true},
		{"1>=2", NewU128(1), NewU128(2), false},
		{"2>=1", NewU128(2), NewU128(1), true},

		{"1>=1", NewI128(1), NewI128(1), true},
		{"1>=2", NewI128(1), NewI128(2), false},
		{"1>=-1", NewI128(1), NewI128(-1), true},
		{"-1>=-1", NewI128(-1), NewI128(-1), true},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Gte(testExample.b)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_ToU64(t *testing.T) {
	testExamples := []struct {
		label       string
		input       Numeric
		expectation U64
	}{
		{"U8(255))=>U64(255)", U8(math.MaxUint8), U64(math.MaxUint8)},
		{"I8(127)=>U64(127)", I8(math.MaxInt8), U64(math.MaxInt8)},
		{"U16(65535)=>U64(65535)", U16(math.MaxUint16), U64(math.MaxUint16)},
		{"I16(32767)=>U64(32767)", I16(math.MaxInt16), U64(math.MaxInt16)},
		{"U32(4294967295)=>U64(4294967295)", U32(math.MaxUint32), U64(math.MaxUint32)},
		{"I32(2147483647)=>U64(2147483647)", I32(math.MaxInt32), U64(math.MaxInt32)},
		{"U64(18446744073709551615)=>U64(18446744073709551615)", U64(math.MaxUint64), U64(math.MaxUint64)},
		{"I64(9223372036854775807)=>U64(9223372036854775807)", I64(math.MaxInt64), U64(math.MaxInt64)},
		{"U128(340282366920938463463374607431768211455)=>U64(18446744073709551615)", MaxU128(), U64(math.MaxUint64)},
		{"I128(170141183460469231731687303715884105727)=>U64(9223372036854775807)", MaxI128(), U64(math.MaxUint64)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := To[U64](testExample.input)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_ToI64(t *testing.T) {
	testExamples := []struct {
		label       string
		input       Numeric
		expectation I64
	}{}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := To[I64](testExample.input)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_ToU128(t *testing.T) {
	u1, _ := NewU128FromString("170141183460469231731687303715884105727")

	testExamples := []struct {
		label       string
		input       Numeric
		expectation U128
	}{
		{"U8(255))=>U128(255)", U8(math.MaxUint8), NewU128(math.MaxUint8)},
		{"I8(127)=>U128(127)", I8(math.MaxInt8), NewU128(math.MaxInt8)},
		{"U16(65535)=>U128(65535)", U16(math.MaxUint16), NewU128(math.MaxUint16)},
		{"I16(32767)=>U128(32767)", I16(math.MaxInt16), NewU128(math.MaxInt16)},
		{"U32(4294967295)=>U128(4294967295)", U32(math.MaxUint32), NewU128(math.MaxUint32)},
		{"I32(2147483647)=>U128(2147483647)", I32(math.MaxInt32), NewU128(math.MaxInt32)},
		{"U64(18446744073709551615)=>U128(18446744073709551615)", U64(math.MaxUint64), NewU128(uint64(math.MaxUint64))},
		{"I64(9223372036854775807)=>U128(9223372036854775807)", I64(math.MaxInt64), NewU128(math.MaxInt64)},
		{"U128(340282366920938463463374607431768211455)=>U128(340282366920938463463374607431768211455)", MaxU128(), MaxU128()},
		{"I128(170141183460469231731687303715884105727)=>U128(170141183460469231731687303715884105727)", MaxI128(), u1},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := To[U128](testExample.input)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_ToI128(t *testing.T) {
	u1, _ := NewU128FromString("170141183460469231731687303715884105727")
	i1, _ := NewI128FromString("170141183460469231731687303715884105727")

	testExamples := []struct {
		label       string
		input       Numeric
		expectation I128
	}{
		{"U8(255))=>I128(255)", U8(math.MaxUint8), NewI128(math.MaxUint8)},
		{"I8(127)=>I128(127)", I8(math.MinInt8), NewI128(math.MinInt8)},
		{"U16(65535)=>I128(65535)", U16(math.MaxUint16), NewI128(math.MaxUint16)},
		{"I16(32767)=>I128(32767)", I16(math.MinInt16), NewI128(math.MinInt16)},
		{"U32(4294967295)=>I128(4294967295)", U32(math.MaxUint32), NewI128(math.MaxUint32)},
		{"I32(2147483647)=>I128(2147483647)", I32(math.MinInt32), NewI128(math.MinInt32)},
		{"U64(18446744073709551615)=>I128(18446744073709551615)", U64(math.MaxUint64), NewI128(uint64(math.MaxUint64))},
		{"I64(9223372036854775807)=>I128(9223372036854775807)", I64(math.MinInt64), NewI128(math.MinInt64)},
		{"U128(170141183460469231731687303715884105727)=>I128(170141183460469231731687303715884105727)", u1, i1},
		{"I128(170141183460469231731687303715884105727)=>I128(170141183460469231731687303715884105727)", i1, i1},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := To[I128](testExample.input)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_U128ToBigInt(t *testing.T) {
	bnMaxU128Value, _ := new(big.Int).SetString("340282366920938463463374607431768211455", 10)

	testExamples := []struct {
		label       string
		input       U128
		expectation *big.Int
	}{
		{"340282366920938463463374607431768211455", MaxU128(), bnMaxU128Value},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.input.ToBigInt()
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_I128ToBigInt(t *testing.T) {
	bnMaxI128Value, _ := new(big.Int).SetString("170141183460469231731687303715884105727", 10)
	bnMinI128Value, _ := new(big.Int).SetString("-170141183460469231731687303715884105728", 10)

	testExamples := []struct {
		label       string
		input       I128
		expectation *big.Int
	}{
		{"170141183460469231731687303715884105727", MaxI128(), bnMaxI128Value},
		{"-170141183460469231731687303715884105728", MinI128(), bnMinI128Value},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.input.ToBigInt()
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

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
