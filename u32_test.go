package goscale

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

			assert.Equal(t, buffer.Bytes(), testExample.expectation)
			assert.Equal(t, testExample.input.Bytes(), testExample.expectation)
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

			result, _ := DecodeU32(buffer)

			assert.Equal(t, result, testExample.expectation)
		})
	}
}
