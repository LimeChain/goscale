package goscale

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

			err := testExample.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, buffer.Bytes(), testExample.expectation)
			assert.Equal(t, testExample.input.Bytes(), testExample.expectation)
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

			result, err := DecodeU64(buffer)

			assert.NoError(t, err)
			assert.Equal(t, result, testExample.expectation)
		})
	}
}
