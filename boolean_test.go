package goscale

import (
	"bytes"
	"testing"
)

func Test_EncodeBool(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Bool
		expectation []byte
	}{
		{label: "Bool(false)", input: Bool(false), expectation: []byte{0x00}},
		{label: "Bool(true)", input: Bool(true), expectation: []byte{0x01}},
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

func Test_DecodeBool(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation Bool
	}{
		{label: "(0x00)", input: []byte{0x00}, expectation: Bool(false)},
		{label: "(0x01)", input: []byte{0x01}, expectation: Bool(true)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.Write(testExample.input)

			dec := Decoder{Reader: &buffer}
			result := dec.DecodeBool()

			assertEqual(t, result, testExample.expectation)
		})
	}
}
