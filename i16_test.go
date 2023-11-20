package goscale

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EncodeI16(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       I16
		expectation []byte
	}{
		{label: "int16(-128)", input: -128, expectation: []byte{0x80, 0xff}},
		{label: "int16(127)", input: 127, expectation: []byte{0x7f, 0x00}},
		{label: "int16(42)", input: 42, expectation: []byte{0x2a, 0x00}},
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

func Test_DecodeI16(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation I16
	}{
		{label: "(0x80ff)", input: []byte{0x80, 0xff}, expectation: -128},
		{label: "(0x2a00)", input: []byte{0x2a, 0x00}, expectation: 42},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			result, err := DecodeI16(buffer)

			assert.NoError(t, err)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}
