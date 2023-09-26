package goscale

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EncodeI32(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       I32
		expectation []byte
	}{
		{label: "int32(-128)", input: -128, expectation: []byte{0x80, 0xff, 0xff, 0xff}},
		{label: "int32(16777215)", input: 16777215, expectation: []byte{0xff, 0xff, 0xff, 0x00}},
		{label: "int32(127)", input: 127, expectation: []byte{0x7f, 0x00, 0x00, 0x00}},
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

func Test_DecodeI32(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation I32
	}{
		{label: "(0x80ffffff)", input: []byte{0x80, 0xff, 0xff, 0xff}, expectation: -128},
		{label: "(0xffffff00)", input: []byte{0xff, 0xff, 0xff, 0x00}, expectation: 16777215},
		{label: "(0x7f000000)", input: []byte{0x7f, 0x00, 0x00, 0x00}, expectation: 127},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			result := DecodeI32(buffer)

			assert.Equal(t, result, testExample.expectation)
		})
	}
}