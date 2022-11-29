package goscale

import (
	"bytes"
	"math"
	"math/big"
	"testing"
)

func Test_EncodeUint8(t *testing.T) {
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
		})
	}
}

func Test_DecodeUint8(t *testing.T) {
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

			testExample.expectation.Encode(buffer)

			result := DecodeU8(buffer)

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeInt8(t *testing.T) {
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
		})
	}
}

func Test_DecodeInt8(t *testing.T) {
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

			testExample.expectation.Encode(buffer)

			result := DecodeI8(buffer)

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeUint16(t *testing.T) {
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
		})
	}
}

func Test_DecodeUint16(t *testing.T) {
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

			testExample.expectation.Encode(buffer)

			result := DecodeU16(buffer)

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeInt16(t *testing.T) {
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
		})
	}
}

func Test_DecodeInt16(t *testing.T) {
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

			testExample.expectation.Encode(buffer)

			result := DecodeI16(buffer)

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeUint32(t *testing.T) {
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
		})
	}
}

func Test_DecodeUint32(t *testing.T) {
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

			testExample.expectation.Encode(buffer)

			result := DecodeU32(buffer)

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeInt32(t *testing.T) {
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
		})
	}
}

func Test_DecodeInt32(t *testing.T) {
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

			testExample.expectation.Encode(buffer)

			result := DecodeI32(buffer)

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeUint64(t *testing.T) {
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
		})
	}
}

func Test_DecodeUint64(t *testing.T) {
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

			testExample.expectation.Encode(buffer)

			result := DecodeU64(buffer)

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeInt64(t *testing.T) {
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
		})
	}
}

func Test_DecodeInt64(t *testing.T) {
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

			testExample.expectation.Encode(buffer)

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
		{label: "Encode U128 - (0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF9c)", input: "340282366920938463463374607431768211356", expect: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x9c}},
		{label: "Encode U128 - (0x2a)", input: "42", expect: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2a}},
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
		})
	}
}

func Test_DecodeU128(t *testing.T) {
	var examples = []struct {
		label  string
		input  []byte
		expect U128
	}{
		{label: "Decode U128 - (0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF)", input: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, expect: U128{math.MaxUint64, math.MaxUint64}},
		{label: "Decode U128 - (42)", input: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2a}, expect: U128{0, 3026418949592973312}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeU128(buffer)

			// then:
			assertEqual(t, result, e.expect)
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
