package goscale

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EncodeI8(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       I8
		expectation []byte
	}{
		{label: "int8(0)", input: 0, expectation: []byte{0x00}},
		{label: "int8(-128)", input: -128, expectation: []byte{0x80}},
		{label: "int8(127)", input: 127, expectation: []byte{0x7f}},
		{label: "int8(-1)", input: -1, expectation: []byte{0xff}},
		{label: "int8(69)", input: 69, expectation: []byte{0x45}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			err := testExample.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, testExample.expectation, buffer.Bytes())
			assert.Equal(t, testExample.expectation, testExample.input.Bytes())
		})
	}
}

func Test_DecodeI8(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation I8
	}{
		{label: "(0x80)", input: []byte{0x80}, expectation: -128},
		{label: "(0x7f)", input: []byte{0x7f}, expectation: 127},
		{label: "(0xff)", input: []byte{0xff}, expectation: -1},
		{label: "(0x45)", input: []byte{0x45}, expectation: 69},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			result, err := DecodeI8(buffer)

			assert.NoError(t, err)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_DecodeI8_Empty(t *testing.T) {
	buffer := &bytes.Buffer{}

	result, err := DecodeI8(buffer)

	assert.Equal(t, io.EOF, err)
	assert.Equal(t, I8(0), result)
}
