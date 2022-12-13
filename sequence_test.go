package goscale

import (
	"bytes"
	"testing"
)

func Test_EncodeString(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Str
		expectation []byte
	}{
		{
			label:       "(abc)",
			input:       Str("abc"),
			expectation: []byte{0x0c, 0x61, 0x62, 0x63},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

func Test_DecodeString(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation Str
	}{
		{
			label:       "(0x0c616263)",
			input:       []byte{0x0c, 0x61, 0x62, 0x63},
			expectation: Str("abc"),
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.expectation.Encode(buffer)

			result := DecodeStr(buffer)

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeU8Sequence(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Sequence[U8]
		expectation []byte
	}{
		{
			label:       "(0x616263)",
			input:       Sequence[U8]{0x61, 0x62, 0x63},
			expectation: []byte{0x0c, 0x61, 0x62, 0x63},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

func Test_DecodeU8Sequence(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation Sequence[U8]
	}{
		{
			label:       "(0x616263)",
			input:       []byte{0x0c, 0x61, 0x62, 0x63},
			expectation: Sequence[U8]{0x61, 0x62, 0x63},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.expectation.Encode(buffer)

			result := Sequence[U8](DecodeSliceU8(buffer))

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeBoolSequence(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Sequence[Bool]
		expectation []byte
	}{
		{
			label:       "([false,true,true])",
			input:       Sequence[Bool]{false, true, true},
			expectation: []byte{0x0c, 0x00, 0x01, 0x01},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

func Test_EncodeCompactSequence(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Sequence[Compact]
		expectation []byte
	}{
		{
			label: "()",
			input: Sequence[Compact]{42, 63, 64, 65535, 1073741823},
			expectation: []byte{
				0x14,
				0xa8,
				0xfc,
				0x01, 0x01,
				0xfe, 0xff, 0x03, 0x00,
				0xfe, 0xff, 0xff, 0xff,
			},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

func Test_EncodeI8Sequence(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Sequence[I8]
		expectation []byte
	}{
		{
			label:       "(0x616263)",
			input:       Sequence[I8]{0x61, 0x62, 0x63},
			expectation: []byte{0x0c, 0x61, 0x62, 0x63},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

func Test_EncodeI16Sequence(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Sequence[I16]
		expectation []byte
	}{
		{
			label:       "[]int16{0,1,-1,2,-2,3,-3}",
			input:       Sequence[I16]{0, 1, -1, 2, -2, 3, -3},
			expectation: []byte{0x1c, 0x00, 0x00, 0x01, 0x00, 0xff, 0xff, 0x02, 0x00, 0xfe, 0xff, 0x03, 0x00, 0xfd, 0xff},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

func Test_EncodeU16Sequence(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Sequence[U16]
		expectation []byte
	}{
		{
			label:       "([4,8,15,16,23,42])",
			input:       Sequence[U16]{4, 8, 15, 16, 23, 42},
			expectation: []byte{0x18, 0x04, 0x00, 0x08, 0x00, 0x0f, 0x00, 0x10, 0x00, 0x17, 0x00, 0x2a, 0x00},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

func Test_EncodeNestedSequence(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Sequence[Sequence[U8]]
		expectation []byte
	}{
		{
			label: "()",
			input: Sequence[Sequence[U8]]{
				Sequence[U8]{0x33, 0x55},
				Sequence[U8]{0x77, 0x99},
			},
			expectation: []byte{0x08, 0x08, 0x33, 0x55, 0x08, 0x77, 0x99},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

func Test_EncodeStringSequence(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Sequence[Str]
		expectation []byte
	}{
		{
			label: "([\"a1\",\"b2\",\"c3\"])",
			input: Sequence[Str]{"a1", "b2", "c3"},
			expectation: []byte{
				0x0c,
				0x08, 0x61, 0x31,
				0x08, 0x62, 0x32,
				0x08, 0x63, 0x33,
			},
		},
		{
			label: "([\"Hamlet\",\"Война и мир\",\"三国演义\",\"أَلْف لَيْلَة وَلَيْلَة\u200e\"])",
			input: Sequence[Str]{"Hamlet", "Война и мир", "三国演义", "أَلْف لَيْلَة وَلَيْلَة\u200e"},
			expectation: []byte{
				0x10,
				0x18, 0x48, 0x61, 0x6d, 0x6c, 0x65, 0x74,
				0x50, 0xd0, 0x92, 0xd0, 0xbe, 0xd0, 0xb9, 0xd0, 0xbd, 0xd0, 0xb0, 0x20, 0xd0, 0xb8, 0x20, 0xd0, 0xbc, 0xd0, 0xb8, 0xd1, 0x80,
				0x30, 0xe4, 0xb8, 0x89, 0xe5, 0x9b, 0xbd, 0xe6, 0xbc, 0x94, 0xe4, 0xb9, 0x89,
				0xbc, 0xd8, 0xa3, 0xd9, 0x8e, 0xd9, 0x84, 0xd9, 0x92, 0xd9, 0x81, 0x20, 0xd9, 0x84, 0xd9, 0x8e, 0xd9, 0x8a, 0xd9, 0x92, 0xd9, 0x84, 0xd9, 0x8e, 0xd8, 0xa9, 0x20, 0xd9, 0x88, 0xd9, 0x8e, 0xd9, 0x84, 0xd9, 0x8e, 0xd9, 0x8a, 0xd9, 0x92, 0xd9, 0x84, 0xd9, 0x8e, 0xd8, 0xa9, 0xe2, 0x80, 0x8e,
			},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

func Test_EncodeFixedSequence(t *testing.T) {
	var examples = []struct {
		label  string
		input  FixedSequence[U8]
		expect []byte
	}{
		{
			label:  "Encode FixedSequence[U8]",
			input:  FixedSequence[U8]{5, 6, 7},
			expect: []byte{0x5, 0x6, 0x7},
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			e.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), e.expect)
		})
	}
}

func Test_DecodeFixedSequence(t *testing.T) {
	var examples = []struct {
		label  string
		input  []byte
		expect FixedSequence[U8]
	}{
		{
			label:  "Decode FixedSequence[U8]",
			input:  []byte{0x5, 0x6, 0x7},
			expect: FixedSequence[U8]{5, 6, 7},
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeFixedSequence[U8](len(e.input), buffer)

			// then:
			assertEqual(t, result, e.expect)
		})
	}
}

func Test_DecodeSequenceU8(t *testing.T) {
	var examples = []struct {
		label  string
		input  []byte
		expect Sequence[U8]
	}{
		{
			label:  "Decode Sequence[U8]",
			input:  []byte{0x0c, 0x61, 0x62, 0x63},
			expect: Sequence[U8]{0x61, 0x62, 0x63},
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeSequence[U8](buffer)

			// then:
			assertEqual(t, result, e.expect)
		})
	}
}
