package goscale

import (
	"bytes"
	"math"
	"math/big"
	"testing"
)

func Test_EncodeU8(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       U8
		expectation []byte
	}{
		{label: "uint8(255)", input: U8(255), expectation: []byte{0xff}},
		{label: "uint8(0)", input: U8(0), expectation: []byte{0x00}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
			assertEqual(t, testExample.input.Bytes(), testExample.expectation)
		})
	}
}

func Test_DecodeU8(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation U8
	}{
		{label: "(0xff)", input: []byte{0xff}, expectation: U8(255)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			result := DecodeU8(buffer)

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeI8(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       I8
		expectation []byte
	}{
		{label: "int8(0)", input: I8(0), expectation: []byte{0x00}},
		{label: "int8(-128)", input: I8(-128), expectation: []byte{0x80}},
		{label: "int8(127)", input: I8(127), expectation: []byte{0x7f}},
		{label: "int8(-1)", input: I8(-1), expectation: []byte{0xff}},
		{label: "int8(69)", input: I8(69), expectation: []byte{0x45}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
			assertEqual(t, testExample.input.Bytes(), testExample.expectation)
		})
	}
}

func Test_DecodeI8(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation I8
	}{
		{label: "(0x80)", input: []byte{0x80}, expectation: I8(-128)},
		{label: "(0x7f)", input: []byte{0x7f}, expectation: I8(127)},
		{label: "(0xff)", input: []byte{0xff}, expectation: I8(-1)},
		{label: "(0x45)", input: []byte{0x45}, expectation: I8(69)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			result := DecodeI8(buffer)

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeU16(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       U16
		expectation []byte
	}{
		{label: "uint16(127)", input: U16(127), expectation: []byte{0x7f, 0x00}},
		{label: "uint16(42)", input: U16(42), expectation: []byte{0x2a, 0x00}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
			assertEqual(t, testExample.input.Bytes(), testExample.expectation)
		})
	}
}

func Test_DecodeU16(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation U16
	}{
		{label: "(0x2a00)", input: []byte{0x2a, 0x00}, expectation: U16(42)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			result := DecodeU16(buffer)

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeI16(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       I16
		expectation []byte
	}{
		{label: "int16(-128)", input: I16(-128), expectation: []byte{0x80, 0xff}},
		{label: "int16(127)", input: I16(127), expectation: []byte{0x7f, 0x00}},
		{label: "int16(42)", input: I16(42), expectation: []byte{0x2a, 0x00}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
			assertEqual(t, testExample.input.Bytes(), testExample.expectation)
		})
	}
}

func Test_DecodeI16(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation I16
	}{
		{label: "(0x80ff)", input: []byte{0x80, 0xff}, expectation: I16(-128)},
		{label: "(0x2a00)", input: []byte{0x2a, 0x00}, expectation: I16(42)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			result := DecodeI16(buffer)

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeU32(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       U32
		expectation []byte
	}{
		{label: "uint32(127)", input: U32(127), expectation: []byte{0x7f, 0x00, 0x00, 0x00}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
			assertEqual(t, testExample.input.Bytes(), testExample.expectation)
		})
	}
}

func Test_DecodeU32(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation U32
	}{
		{label: "(0x7f000000)", input: []byte{0x7f, 0x00, 0x00, 0x00}, expectation: U32(127)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			result := DecodeU32(buffer)

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeI32(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       I32
		expectation []byte
	}{
		{label: "int32(-128)", input: I32(-128), expectation: []byte{0x80, 0xff, 0xff, 0xff}},
		{label: "int32(16777215)", input: I32(16777215), expectation: []byte{0xff, 0xff, 0xff, 0x00}},
		{label: "int32(127)", input: I32(127), expectation: []byte{0x7f, 0x00, 0x00, 0x00}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
			assertEqual(t, testExample.input.Bytes(), testExample.expectation)
		})
	}
}

func Test_DecodeI32(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation I32
	}{
		{label: "(0x80ffffff)", input: []byte{0x80, 0xff, 0xff, 0xff}, expectation: I32(-128)},
		{label: "(0xffffff00)", input: []byte{0xff, 0xff, 0xff, 0x00}, expectation: I32(16777215)},
		{label: "(0x7f000000)", input: []byte{0x7f, 0x00, 0x00, 0x00}, expectation: I32(127)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			result := DecodeI32(buffer)

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeU64(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       U64
		expectation []byte
	}{
		{label: "uint64(127)", input: U64(127), expectation: []byte{0x7f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
			assertEqual(t, testExample.input.Bytes(), testExample.expectation)
		})
	}
}

func Test_DecodeU64(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation U64
	}{
		{label: "(0x7f00000000000000)", input: []byte{0x7f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, expectation: U64(127)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			result := DecodeU64(buffer)

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeI64(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       I64
		expectation []byte
	}{
		{label: "int64(-128)", input: I64(-128), expectation: []byte{0x80, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
			assertEqual(t, testExample.input.Bytes(), testExample.expectation)
		})
	}
}

func Test_DecodeI64(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation I64
	}{
		{label: "(0x80ffffffffffffff)", input: []byte{0x80, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, expectation: I64(-128)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			result := DecodeI64(buffer)

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeU128(t *testing.T) {
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
			// given:
			buffer := &bytes.Buffer{}

			value, ok := new(big.Int).SetString(e.input, 10)
			if !ok {
				panic("not ok")
			}
			input := NewU128FromBigInt(value)

			// when:
			input.Encode(buffer)

			// then:
			assertEqual(t, buffer.Bytes(), e.expect)
			// and:
			assertEqual(t, input.Bytes(), e.expect)
		})
	}
}

func Test_DecodeU128(t *testing.T) {
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
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeU128(buffer)
			bigInt := result.ToBigInt()

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, bigInt.String(), e.stringValue)
		})
	}
}

func Test_NewU128FromBigIntPanic(t *testing.T) {
	t.Run("Exceeds U128", func(t *testing.T) {
		// given:
		value, ok := new(big.Int).SetString("340282366920938463463374607431768211456", 10) // MaxUint128 + 1
		if !ok {
			panic("not ok")
		}

		// then:
		assertPanic(t, func() {
			NewU128FromBigInt(value)
		})
	})
}

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
			// given:
			buffer := &bytes.Buffer{}

			value, ok := new(big.Int).SetString(e.input, 10)
			if !ok {
				panic("not ok")
			}
			input := NewI128FromBigInt(*value)

			// when:
			input.Encode(buffer)

			// then:
			assertEqual(t, buffer.Bytes(), e.expect)
			// and:
			assertEqual(t, input.Bytes(), e.expect)
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
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeI128(buffer)
			bigInt := result.ToBigInt()

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, bigInt.String(), e.stringValue)
		})
	}
}
