package goscale

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MaxU64(t *testing.T) {
	testExamples := []struct {
		label  string
		a      U64
		b      U64
		expect U64
	}{
		{"Max(1, 2)", 1, 2, 2},
		{"Max(3, 1)", 3, 1, 3},
		{"Max(1, MaxU64)", 1, math.MaxUint64, math.MaxUint64},
		{"Max(MaxU64, MaxU64)", math.MaxUint64, math.MaxUint64, math.MaxUint64},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := MaxU64(testExample.a, testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_Max128(t *testing.T) {
	testExamples := []struct {
		label  string
		a      U128
		b      U128
		expect U128
	}{
		{"Max(1, 2)", NewU128(1), NewU128(2), NewU128(2)},
		{"Max(1, MaxU128)", NewU128(1), MaxU128(), MaxU128()},
		{"Max(MaxU128, MaxU128)", MaxU128(), MaxU128(), MaxU128()},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := Max128(testExample.a, testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_MinU64(t *testing.T) {
	testExamples := []struct {
		label  string
		a      U64
		b      U64
		expect U64
	}{
		{"Min(1, 2)", 1, 2, 1},
		{"Min(1, MaxU64)", 1, math.MaxUint64, 1},
		{"Min(MaxU64, MaxU64)", math.MaxUint64, math.MaxUint64, math.MaxUint64},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := MinU64(testExample.a, testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_Min128(t *testing.T) {
	testExamples := []struct {
		label  string
		a      U128
		b      U128
		expect U128
	}{
		{"Min(1, 2)", NewU128(1), NewU128(2), NewU128(1)},
		{"Min(1, MaxU128)", NewU128(1), MaxU128(), NewU128(1)},
		{"Min(MaxU128, MaxU128)", MaxU128(), MaxU128(), MaxU128()},
		{"Min(MinU128, MaxU128)", MinU128(), MaxU128(), MinU128()},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := Min128(testExample.a, testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_Clamp(t *testing.T) {
	testExamples := []struct {
		label    string
		a        int
		minValue int
		maxValue int
		expect   int
	}{
		{"Clamp(5, 1, 10)", 5, 1, 10, 5},
		{"Clamp(15, 1, 10)", 15, 1, 10, 10},
		{"Clamp(0, 1, 10)", 0, 1, 10, 1},
		{"Clamp(-3, -2, 2)", -3, -2, 2, -2},
		{"Clamp(3, -2, 2)", 3, -2, 2, 2},
		{"Clamp(1, -2, 2)", 1, -2, 2, 1},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := Clamp(testExample.a, testExample.minValue, testExample.maxValue)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_SaturatingAddU32(t *testing.T) {
	testExamples := []struct {
		label  string
		a      U32
		b      U32
		expect U32
	}{
		{"2 + 1", 2, 1, 3},
		{"MaxU32+1", math.MaxUint32, 1, math.MaxUint32},
		{"MaxU32+MaxU32", math.MaxUint32, math.MaxUint32, math.MaxUint32},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := SaturatingAddU32(testExample.a, testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_SaturatingAddU64(t *testing.T) {
	testExamples := []struct {
		label  string
		a      U64
		b      U64
		expect U64
	}{
		{"2 + 1", 2, 1, 3},
		{"MaxU64+1", math.MaxUint64, 1, math.MaxUint64},
		{"MaxU64+MaxU64", math.MaxUint64, math.MaxUint64, math.MaxUint64},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := SaturatingAddU64(testExample.a, testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_SaturatingSubU64(t *testing.T) {
	testExamples := []struct {
		label  string
		a      U64
		b      U64
		expect U64
	}{
		{"0-1", 0, 1, 0},
		{"2-1", 2, 1, 1},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := SaturatingSubU64(testExample.a, testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_SaturatingMul(t *testing.T) {
	testExamples := []struct {
		label  string
		a      U64
		b      U64
		expect U64
	}{
		{"0*1", 1, 0, 0},
		{"2*0", 2, 0, 0},
		{"2*2", 2, 2, 4},
		{"2*3", 2, 3, 6},
		{"MaxU64*2", math.MaxUint64, 2, math.MaxUint64},
		{"MaxU64*MaxU64", math.MaxUint64, math.MaxUint64, math.MaxUint64},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := SaturatingMulU64(testExample.a, testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_CheckedAddU32(t *testing.T) {
	testExamples := []struct {
		label        string
		a            U32
		b            U32
		expect       U32
		hasExpectErr bool
	}{
		{"1+2", 1, 2, 3, false},
		{"MaxU32+1", math.MaxUint32, 1, 0, true},
		{"MaxU32+MaxU32", math.MaxUint32, math.MaxUint32, 0, true},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result, err := CheckedAddU32(testExample.a, testExample.b)

			assert.Equal(t, testExample.expect, result)
			if testExample.hasExpectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_CheckedAddU64(t *testing.T) {
	testExamples := []struct {
		label        string
		a            U64
		b            U64
		expect       U64
		hasExpectErr bool
	}{
		{"1+2", 1, 2, 3, false},
		{"MaxU64+1", math.MaxUint64, 1, 0, true},
		{"MaxU64+MaxU64", math.MaxUint64, math.MaxUint64, 0, true},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result, err := CheckedAddU64(testExample.a, testExample.b)

			assert.Equal(t, testExample.expect, result)
			if testExample.hasExpectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
