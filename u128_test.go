package goscale

import (
	"bytes"
	"math"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_U128_Encode(t *testing.T) {
	var examples = []struct {
		label  string
		input  string
		expect []byte
	}{
		{label: "Encode U128 - (0)", input: "0", expect: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}},
		{label: "Encode U128 - (42)", input: "42", expect: []byte{0x2a, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}},
		{label: "Encode U128 - (const.MaxInt64 + 1)", input: "9223372036854775808", expect: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}},
		{label: "Encode U128 - (0x9cFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF)", input: "340282366920938463463374607431768211356", expect: []byte{0x9c, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{label: "Encode U128 - (MaxInt128)", input: "340282366920938463463374607431768211455", expect: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			value, ok := new(big.Int).SetString(e.input, 10)
			if !ok {
				panic("not ok")
			}
			input := NewU128(value)

			input.Encode(buffer)

			assert.Equal(t, buffer.Bytes(), e.expect)
			assert.Equal(t, input.Bytes(), e.expect)
		})
	}
}

func Test_U128_Decode(t *testing.T) {
	var examples = []struct {
		label       string
		input       []byte
		expect      U128
		stringValue string
	}{
		{label: "Decode U128 - (0)", input: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, expect: U128{U64(0), U64(0)}, stringValue: "0"},
		{label: "Decode U128 - (42)", input: []byte{0x2a, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, expect: U128{42, 0}, stringValue: "42"},
		{label: "Decode U128 - (math.MaxInt64 + 1)", input: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, expect: U128{U64(math.MaxInt64 + 1), 0}, stringValue: "9223372036854775808"},
		{label: "Decode U128 - (0x9cFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF)", input: []byte{0x9c, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, expect: U128{18446744073709551516, 18446744073709551615}, stringValue: "340282366920938463463374607431768211356"},
		{label: "Decode U128 - (MaxInt128)", input: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, expect: U128{math.MaxUint64, math.MaxUint64}, stringValue: "340282366920938463463374607431768211455"},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, _ := DecodeU128(buffer)
			bigInt := result.ToBigInt()

			assert.Equal(t, result, e.expect)
			assert.Equal(t, bigInt.String(), e.stringValue)
		})
	}
}

func Test_NewU128FromBigIntPanic(t *testing.T) {
	t.Run("Exceeds U128", func(t *testing.T) {
		value, ok := new(big.Int).SetString("340282366920938463463374607431768211456", 10) // MaxU128 + 1
		if !ok {
			panic("not ok")
		}

		assert.Panics(t, func() { NewU128(value) })
	})
}

func Test_U128_Add(t *testing.T) {
	bn1, _ := NewU128FromString("340282366920938463463374607431768211454")

	testExamples := []struct {
		label  string
		a      U128
		b      U128
		expect U128
	}{
		{"340282366920938463463374607431768211454+1", bn1, NewU128(1), MaxU128()},
		{"340282366920938463463374607431768211455+1", MaxU128(), NewU128(1), NewU128(0)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Add(testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_U128_Sub(t *testing.T) {
	u1, _ := NewU128FromString("340282366920938463463374607431657675311")
	u2, _ := NewU128FromString("340282366920938463463374607421657675311")

	testExamples := []struct {
		label  string
		a      U128
		b      U128
		expect U128
	}{
		{"2-1", NewU128(2), NewU128(1), NewU128(1)},
		{"0-1", NewU128(0), NewU128(1), MaxU128()},
		{"499999889463855-10000000000", NewU128(499999889463855), NewU128(10000000000), NewU128(499989889463855)},
		{"340282366920938463463374607431657675311-10000000000", u1, NewU128(10000000000), u2},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Sub(testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_U128_Mul(t *testing.T) {
	testExamples := []struct {
		label  string
		a      U128
		b      U128
		expect U128
	}{
		{"2*3", NewU128(2), NewU128(3), NewU128(6)},
		{"MaxU128*0", MaxU128(), NewU128(0), NewU128(0)},
		{"MaxU128*1", MaxU128(), NewU128(1), MaxU128()},
		{"MaxU128*2", MaxU128(), NewU128(2), MaxU128().Sub(NewU128(1))},
		{"MaxU128*MaxU128", MaxU128(), MaxU128(), NewU128(1)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Mul(testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_U128_Div(t *testing.T) {
	testExamples := []struct {
		label  string
		a      U128
		b      U128
		expect U128
	}{
		{"0/1", NewU128(0), NewU128(1), NewU128(0)},
		{"1/1", NewU128(1), NewU128(1), NewU128(1)},
		{"6/3", NewU128(6), NewU128(3), NewU128(2)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Div(testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_U128_Eq(t *testing.T) {
	testExamples := []struct {
		label  string
		a      U128
		b      U128
		expect bool
	}{
		{"1==1", NewU128(1), NewU128(1), true},
		{"1==2", NewU128(1), NewU128(2), false},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Eq(testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_U128_Ne(t *testing.T) {
	testExamples := []struct {
		label  string
		a      U128
		b      U128
		expect bool
	}{
		{"1!=1", NewU128(1), NewU128(1), false},
		{"1!=2", NewU128(1), NewU128(2), true},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Ne(testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_U128_Lt(t *testing.T) {
	testExamples := []struct {
		label  string
		a      U128
		b      U128
		expect bool
	}{
		{"1<1", NewU128(1), NewU128(1), false},
		{"1<2", NewU128(1), NewU128(2), true},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Lt(testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_U128_Lte(t *testing.T) {
	testExamples := []struct {
		label  string
		a      U128
		b      U128
		expect bool
	}{
		{"1<=1", NewU128(1), NewU128(1), true},
		{"1<=2", NewU128(1), NewU128(2), true},
		{"1<=0", NewU128(1), NewU128(0), false},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Lte(testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_U128_Gt(t *testing.T) {
	testExamples := []struct {
		label  string
		a      U128
		b      U128
		expect bool
	}{
		{"1>1", NewU128(1), NewU128(1), false},
		{"1>2", NewU128(1), NewU128(2), false},
		{"2>1", NewU128(2), NewU128(1), true},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Gt(testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_U128_Gte(t *testing.T) {
	testExamples := []struct {
		label  string
		a      U128
		b      U128
		expect bool
	}{
		{"1>=1", NewU128(1), NewU128(1), true},
		{"1>=2", NewU128(1), NewU128(2), false},
		{"2>=1", NewU128(2), NewU128(1), true},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Gte(testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_U128_ToBigInt(t *testing.T) {
	bnMaxU128Value, _ := new(big.Int).SetString("340282366920938463463374607431768211455", 10)

	testExamples := []struct {
		label  string
		input  U128
		expect *big.Int
	}{
		{"340282366920938463463374607431768211455", MaxU128(), bnMaxU128Value},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.input.ToBigInt()
			assert.Equal(t, testExample.expect, result)
		})
	}
}
