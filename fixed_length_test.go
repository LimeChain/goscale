package goscale

import (
	"bytes"
	"testing"
)

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
			buffer := bytes.Buffer{}

			enc := &Encoder{Writer: &buffer}
			testExample.input.Encode(enc)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
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
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.Write(testExample.input)

			dec := Decoder{Reader: &buffer}
			result := dec.DecodeU64()

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
			buffer := bytes.Buffer{}

			enc := &Encoder{Writer: &buffer}
			testExample.input.Encode(enc)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
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
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.Write(testExample.input)

			dec := Decoder{Reader: &buffer}
			result := dec.DecodeI64()

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
			buffer := bytes.Buffer{}

			enc := &Encoder{Writer: &buffer}
			testExample.input.Encode(enc)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
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
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.Write(testExample.input)

			dec := Decoder{Reader: &buffer}
			result := dec.DecodeU32()

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
			buffer := bytes.Buffer{}

			enc := &Encoder{Writer: &buffer}
			testExample.input.Encode(enc)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
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
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.Write(testExample.input)

			dec := Decoder{Reader: &buffer}
			result := dec.DecodeI32()

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
			buffer := bytes.Buffer{}

			enc := &Encoder{Writer: &buffer}
			testExample.input.Encode(enc)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
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
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.Write(testExample.input)

			dec := Decoder{Reader: &buffer}
			result := dec.DecodeU16()

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
			buffer := bytes.Buffer{}

			enc := &Encoder{Writer: &buffer}
			testExample.input.Encode(enc)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
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
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.Write(testExample.input)

			dec := Decoder{Reader: &buffer}
			result := dec.DecodeI16()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

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
			buffer := bytes.Buffer{}

			enc := &Encoder{Writer: &buffer}
			testExample.input.Encode(enc)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
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
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.Write(testExample.input)

			dec := Decoder{Reader: &buffer}
			result := dec.DecodeU8()

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
			buffer := bytes.Buffer{}

			enc := &Encoder{Writer: &buffer}
			testExample.input.Encode(enc)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
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
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.Write(testExample.input)

			dec := Decoder{Reader: &buffer}
			result := dec.DecodeI8()

			assertEqual(t, result, testExample.expectation)
		})
	}
}
