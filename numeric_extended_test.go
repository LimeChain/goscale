package goscale

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {
	testExamples := []struct {
		label       string
		a           U64
		b           U64
		expectation U64
	}{
		{"MaxUint64 + 1", math.MaxUint64, 1, 0},
		{"MaxUint64 + MaxUint64", math.MaxUint64, math.MaxUint64, math.MaxUint64 - 1},
		{"1 + 2", 1, 2, 3},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Add(testExample.b)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_Sub(t *testing.T) {
	testExamples := []struct {
		label       string
		a           U64
		b           U64
		expectation U64
	}{
		{"MaxUint64 - MaxUint64", math.MaxUint64, math.MaxUint64, 0},
		{"2 - 1", 2, 1, 1},
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
		a           U64
		b           U64
		expectation U64
	}{
		{"MaxUint64 * MaxUint64", math.MaxUint64, math.MaxUint64, 1},
		{"MaxUint64 * 2", math.MaxUint64, 2, math.MaxUint64 - 1},
		{"2 * 3", 2, 3, 6},
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
		a           U64
		b           U64
		expectation U64
	}{
		{"MaxUint64 / 1", math.MaxUint64, 1, math.MaxUint64},
		{"6 / 3", 6, 3, 2},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Div(testExample.b)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_Max(t *testing.T) {
	testExamples := []struct {
		label       string
		a           U64
		b           U64
		expectation U64
	}{
		{"Max(1, MaxUint64)", 1, math.MaxUint64, math.MaxUint64},
		{"Max(MaxUint64, MaxUint64)", math.MaxUint64, math.MaxUint64, math.MaxUint64},
		{"Max(2, 1)", 2, 1, 2},
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
		a           U64
		b           U64
		expectation U64
	}{
		{"Min(1, MaxUint64)", 1, math.MaxUint64, 1},
		{"Min(MaxUint64, MaxUint64)", math.MaxUint64, math.MaxUint64, math.MaxUint64},
		{"Min(2, 1)", 2, 1, 1},
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
		a           U64
		expectation U64
	}{
		{"TrailingZeros(1)", 1, 0},
		{"TrailingZeros(2)", 2, 1},
		{"TrailingZeros(3)", 3, 0},
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
		a           U64
		minValue    U64
		maxValue    U64
		expectation U64
	}{
		{"Clamp(5, 1, 10)", 5, 1, 10, 5},
		{"Clamp(15, 1, 10)", 15, 1, 10, 10},
		{"Clamp(0, 1, 10)", 0, 1, 10, 1},
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
		a           U64
		b           U64
		expectation U64
	}{
		{"MaxUint64 + 1", math.MaxUint64, 1, math.MaxUint64},
		{"MaxUint64 + MaxUint64", math.MaxUint64, math.MaxUint64, math.MaxUint64},
		{"1 + 2", 1, 2, 3},
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
		a           U64
		b           U64
		expectation U64
	}{
		{"1 sub MaxUint64", 1, math.MaxUint64, 0},
		{"MaxUint64 sub MaxUint64", math.MaxUint64, math.MaxUint64, 0},
		{"2 sub 1", 2, 1, 1},
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
		a           U64
		b           U64
		expectation U64
	}{
		{"MaxUint64 mul MaxUint64", math.MaxUint64, math.MaxUint64, math.MaxUint64},
		{"MaxUint64 mul 2", math.MaxUint64, 2, math.MaxUint64},
		{"2 mul 3", 2, 3, 6},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.SaturatingMul(testExample.b)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_SaturatingDiv(t *testing.T) {
	testExamples := []struct {
		label       string
		a           U64
		b           U64
		expectation U64
	}{
		{"MaxUint64 div 0", math.MaxUint64, 0, math.MaxUint64},
		{"MaxUint64 div 1", math.MaxUint64, 1, math.MaxUint64},
		{"6 div 3", 6, 3, 2},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.SaturatingDiv(testExample.b)
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
		{"MaxUint64 add 1", math.MaxUint64, 1, 0, true},
		{"MaxUint64 add MaxUint64", math.MaxUint64, math.MaxUint64, 0, true},
		{"1 add 2", 1, 2, 3, false},
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
