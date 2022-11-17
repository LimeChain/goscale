package goscale

import (
	"bytes"
	"testing"
)

func Test_EncodeUint64(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       uint64
		expectation []byte
	}{
		{label: "uint64(127)", input: uint64(127), expectation: []byte{0x7f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.EncodeUint64(testExample.input)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_DecodeUint64(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation uint64
	}{
		{label: "(0x7f00000000000000)", input: []byte{0x7f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, expectation: uint64(127)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.Write(testExample.input)

			dec := Decoder{Reader: &buffer}
			result := dec.DecodeUint64()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeInt64(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       int64
		expectation []byte
	}{
		{label: "int64(-128)", input: int64(-128), expectation: []byte{0x80, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.EncodeInt64(testExample.input)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_DecodeInt64(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation int64
	}{
		{label: "(0x80ffffffffffffff)", input: []byte{0x80, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, expectation: int64(-128)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.Write(testExample.input)

			dec := Decoder{Reader: &buffer}
			result := dec.DecodeInt64()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeUint32(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       uint32
		expectation []byte
	}{
		{label: "uint32(127)", input: uint32(127), expectation: []byte{0x7f, 0x00, 0x00, 0x00}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.EncodeUint32(testExample.input)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_DecodeUint32(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation uint32
	}{
		{label: "(0x7f000000)", input: []byte{0x7f, 0x00, 0x00, 0x00}, expectation: uint32(127)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.Write(testExample.input)

			dec := Decoder{Reader: &buffer}
			result := dec.DecodeUint32()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeInt32(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       int32
		expectation []byte
	}{
		{label: "int32(-128)", input: int32(-128), expectation: []byte{0x80, 0xff, 0xff, 0xff}},
		{label: "int32(16777215)", input: int32(16777215), expectation: []byte{0xff, 0xff, 0xff, 0x00}},
		{label: "int32(127)", input: int32(127), expectation: []byte{0x7f, 0x00, 0x00, 0x00}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.EncodeInt32(testExample.input)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_DecodeInt32(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation int32
	}{
		{label: "(0x80ffffff)", input: []byte{0x80, 0xff, 0xff, 0xff}, expectation: int32(-128)},
		{label: "(0xffffff00)", input: []byte{0xff, 0xff, 0xff, 0x00}, expectation: int32(16777215)},
		{label: "(0x7f000000)", input: []byte{0x7f, 0x00, 0x00, 0x00}, expectation: int32(127)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.Write(testExample.input)

			dec := Decoder{Reader: &buffer}
			result := dec.DecodeInt32()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeUint16(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       uint16
		expectation []byte
	}{
		{label: "uint16(127)", input: uint16(127), expectation: []byte{0x7f, 0x00}},
		{label: "uint16(42)", input: uint16(42), expectation: []byte{0x2a, 0x00}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.EncodeUint16(testExample.input)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_DecodeUint16(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation uint16
	}{
		{label: "(0x2a00)", input: []byte{0x2a, 0x00}, expectation: uint16(42)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.Write(testExample.input)

			dec := Decoder{Reader: &buffer}
			result := dec.DecodeUint16()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeInt16(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       int16
		expectation []byte
	}{
		{label: "int16(-128)", input: int16(-128), expectation: []byte{0x80, 0xff}},
		{label: "int16(127)", input: int16(127), expectation: []byte{0x7f, 0x00}},
		{label: "int16(42)", input: int16(42), expectation: []byte{0x2a, 0x00}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.EncodeInt16(testExample.input)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_DecodeInt16(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation int16
	}{
		{label: "(0x80ff)", input: []byte{0x80, 0xff}, expectation: int16(-128)},
		{label: "(0x2a00)", input: []byte{0x2a, 0x00}, expectation: int16(42)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.Write(testExample.input)

			dec := Decoder{Reader: &buffer}
			result := dec.DecodeInt16()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeUint8(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       uint8
		expectation []byte
	}{
		{label: "uint8(255)", input: uint8(255), expectation: []byte{0xff}},
		{label: "uint8(0)", input: uint8(0), expectation: []byte{0x00}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.EncodeUint8(testExample.input)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_DecodeUint8(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation uint8
	}{
		{label: "(0xff)", input: []byte{0xff}, expectation: uint8(255)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.Write(testExample.input)

			dec := Decoder{Reader: &buffer}
			result := dec.DecodeUint8()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeInt8(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       int8
		expectation []byte
	}{
		{label: "int8(0)", input: int8(0), expectation: []byte{0x00}},
		{label: "int8(-128)", input: int8(-128), expectation: []byte{0x80}},
		{label: "int8(127)", input: int8(127), expectation: []byte{0x7f}},
		{label: "int8(-1)", input: int8(-1), expectation: []byte{0xff}},
		{label: "int8(69)", input: int8(69), expectation: []byte{0x45}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.EncodeInt8(testExample.input)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_DecodeInt8(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation int8
	}{
		{label: "(0x80)", input: []byte{0x80}, expectation: int8(-128)},
		{label: "(0x7f)", input: []byte{0x7f}, expectation: int8(127)},
		{label: "(0xff)", input: []byte{0xff}, expectation: int8(-1)},
		{label: "(0x45)", input: []byte{0x45}, expectation: int8(69)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.Write(testExample.input)

			dec := Decoder{Reader: &buffer}
			result := dec.DecodeInt8()

			assertEqual(t, result, testExample.expectation)
		})
	}
}
