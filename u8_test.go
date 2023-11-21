package goscale

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
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

			err := testExample.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, testExample.expectation, buffer.Bytes())
			assert.Equal(t, testExample.expectation, testExample.input.Bytes())
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

			result, err := DecodeU8(buffer)

			assert.NoError(t, err)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}
