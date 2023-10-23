package goscale

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EncodeI64(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       I64
		expectation []byte
	}{
		{label: "int64(-128)", input: -128, expectation: []byte{0x80, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
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

func Test_DecodeI64(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation I64
	}{
		{label: "(0x80ffffffffffffff)", input: []byte{0x80, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, expectation: -128},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			result, _ := DecodeI64(buffer)

			assert.Equal(t, result, testExample.expectation)
		})
	}
}
