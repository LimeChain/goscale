package goscale

import (
	"bytes"
	"math"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EncodeI128(t *testing.T) {
	var examples = []struct {
		label  string
		input  string
		expect []byte
	}{
		{label: "Encode I128 - (MinI128)", input: "-170141183460469231731687303715884105728", expect: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80}},
		{label: "Encode I128 - (-123456789)", input: "-123456789", expect: []byte{0xeb, 0x32, 0xa4, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{label: "Encode I128 - (-42)", input: "-42", expect: []byte{0xd6, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{label: "Encode I128 - (-1)", input: "-1", expect: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{label: "Encode I128 - (0)", input: "0", expect: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}},
		{label: "Encode I128 - (1)", input: "1", expect: []byte{0x01, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}},
		{label: "Encode I128 - (42)", input: "42", expect: []byte{0x2a, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}},
		{label: "Encode I128 - (123456789)", input: "123456789", expect: []byte{0x15, 0xcd, 0x5b, 0x07, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}},
		{label: "Encode I128 - (MaxInt128)", input: "170141183460469231731687303715884105727", expect: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			value, ok := new(big.Int).SetString(e.input, 10)
			if !ok {
				panic("not ok")
			}
			input := bigIntToI128(value)

			err := input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, buffer.Bytes())
			assert.Equal(t, e.expect, input.Bytes())
		})
	}
}

func Test_DecodeI128(t *testing.T) {
	var examples = []struct {
		label       string
		input       []byte
		expect      I128
		stringValue string
	}{
		{label: "Decode I128 - (MinInt128 == -170141183460469231731687303715884105728", input: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80}, expect: I128{U64(0), U64(math.MaxInt64 + 1)}, stringValue: "-170141183460469231731687303715884105728"},
		{label: "Decode I128 - (-123456789)", input: []byte{0xeb, 0x32, 0xa4, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, expect: I128{U64(math.MaxUint64 - 123456789 + 1), U64(math.MaxUint64)}, stringValue: "-123456789"},
		{label: "Decode I128 - (-42)", input: []byte{0xd6, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, expect: I128{U64(math.MaxUint64 - 41), U64(math.MaxUint64)}, stringValue: "-42"},
		{label: "Decode I128 - (-1)", input: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, expect: I128{U64(math.MaxUint64), U64(math.MaxUint64)}, stringValue: "-1"},
		{label: "Encode I128 - (0)", input: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, expect: I128{U64(0), U64(0)}, stringValue: "0"},
		{label: "Encode I128 - (1)", input: []byte{0x01, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, expect: I128{U64(1), U64(0)}, stringValue: "1"},
		{label: "Decode I128 - (42)", input: []byte{0x2a, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, expect: I128{U64(42), U64(0)}, stringValue: "42"},
		{label: "Encode I128 - (123456789)", input: []byte{0x15, 0xcd, 0x5b, 0x07, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, expect: I128{U64(123456789)}, stringValue: "123456789"},
		{label: "Decode I128 - (MaxInt128 == 170141183460469231731687303715884105727)", input: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f}, expect: I128{U64(math.MaxUint64), U64(math.MaxInt64)}, stringValue: "170141183460469231731687303715884105727"},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeI128(buffer)
			assert.NoError(t, err)
			bigInt := result.ToBigInt()

			assert.Equal(t, e.expect, result)
			assert.Equal(t, e.stringValue, bigInt.String())
		})
	}
}

func Test_I128_Add(t *testing.T) {
	testExamples := []struct {
		label  string
		a      I128
		b      I128
		expect I128
	}{
		{"-2+(-1)", NewI128(-2), NewI128(-1), NewI128(-3)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Add(testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_I128_Sub(t *testing.T) {
	u1, _ := NewI128FromString("340282366920938463463374607431657675311")
	u2, _ := NewI128FromString("340282366920938463463374607421657675311")
	u3, _ := NewI128FromString("340282366920938463463374607431657675310")

	testExamples := []struct {
		label  string
		a      I128
		b      I128
		expect I128
	}{
		{"2-1", NewI128(2), NewI128(1), NewI128(1)},
		{"-2-(-1)", NewI128(-2), NewI128(-1), NewI128(-1)},
		{"-170141183460469231731687303715884105728-1", MinI128(), NewI128(1), MaxI128()},
		{"499999889463855-10000000000", NewI128(499999889463855), NewI128(10000000000), NewI128(499989889463855)},
		{"340282366920938463463374607431657675311-10000000000", u1, NewI128(10000000000), u2},
		{"9889463854-10000000000", NewI128(9889463854), NewI128(10000000000), u3},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Sub(testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_I128_Mul(t *testing.T) {
	testExamples := []struct {
		label  string
		a      I128
		b      I128
		expect I128
	}{
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
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_I128_Div(t *testing.T) {
	testExamples := []struct {
		label  string
		a      I128
		b      I128
		expect I128
	}{
		{"0/-1", NewI128(0), NewI128(-1), NewI128(0)},
		{"-1/1", NewI128(-1), NewI128(1), NewI128(-1)},
		{"-1/-1", NewI128(-1), NewI128(-1), NewI128(1)},
		{"-6/2", NewI128(-6), NewI128(2), NewI128(-3)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Div(testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_I128_Eq(t *testing.T) {
	testExamples := []struct {
		label  string
		a      I128
		b      I128
		expect bool
	}{
		{"1==1", NewI128(1), NewI128(1), true},
		{"1==2", NewI128(1), NewI128(2), false},
		{"-1==1", NewI128(-1), NewI128(1), false},
		{"-1==-1", NewI128(-1), NewI128(-1), true},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Eq(testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_I128_Ne(t *testing.T) {
	testExamples := []struct {
		label  string
		a      I128
		b      I128
		expect bool
	}{
		{"1!=1", NewI128(U8(1)), NewI128(U16(1)), false},
		{"1!=2", NewI128(U32(1)), NewI128(U64(2)), true},
		{"-1!=1", NewI128(I8(-1)), NewI128(I16(1)), true},
		{"-1!=-1", NewI128(I32(-1)), NewI128(I64(-1)), false},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Ne(testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_I128_Lt(t *testing.T) {
	testExamples := []struct {
		label  string
		a      I128
		b      I128
		expect bool
	}{
		{"1<1", NewI128(1), NewI128(1), false},
		{"1<2", NewI128(1), NewI128(2), true},
		{"-1<1", NewI128(-1), NewI128(1), true},
		{"-1<-1", NewI128(-1), NewI128(-1), false},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Lt(testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_I128_Lte(t *testing.T) {
	testExamples := []struct {
		label  string
		a      I128
		b      I128
		expect bool
	}{
		{"1<=1", NewI128(1), NewI128(1), true},
		{"1<=2", NewI128(1), NewI128(2), true},
		{"-1<=1", NewI128(-1), NewI128(1), true},
		{"-1<=-1", NewI128(-1), NewI128(-1), true},
		{"-1<=-2", NewI128(-1), NewI128(-2), false},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Lte(testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_I128_Gt(t *testing.T) {
	testExamples := []struct {
		label  string
		a      I128
		b      I128
		expect bool
	}{
		{"1>1", NewI128(1), NewI128(1), false},
		{"1>2", NewI128(1), NewI128(2), false},
		{"1>-1", NewI128(1), NewI128(-1), true},
		{"-1>-1", NewI128(-1), NewI128(-1), false},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Gt(testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_I128_Gte(t *testing.T) {
	testExamples := []struct {
		label  string
		a      I128
		b      I128
		expect bool
	}{
		{"1>=1", NewI128(1), NewI128(1), true},
		{"1>=2", NewI128(1), NewI128(2), false},
		{"1>=-1", NewI128(1), NewI128(-1), true},
		{"-1>=-1", NewI128(-1), NewI128(-1), true},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.a.Gte(testExample.b)
			assert.Equal(t, testExample.expect, result)
		})
	}
}

func Test_I128_ToBigInt(t *testing.T) {
	bnMaxI128Value, _ := new(big.Int).SetString("170141183460469231731687303715884105727", 10)
	bnMinI128Value, _ := new(big.Int).SetString("-170141183460469231731687303715884105728", 10)

	testExamples := []struct {
		label  string
		input  I128
		expect *big.Int
	}{
		{"170141183460469231731687303715884105727", MaxI128(), bnMaxI128Value},
		{"-170141183460469231731687303715884105728", MinI128(), bnMinI128Value},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			result := testExample.input.ToBigInt()
			assert.Equal(t, testExample.expect, result)
		})
	}
}
