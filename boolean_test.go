package goscale

import (
	"bytes"

	"testing"
)

func Test_EncodeBool(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       bool
		expectation []byte
	}{
		{label: "(false)", input: false, expectation: []byte{0x00}},
		{label: "(true)", input: true, expectation: []byte{0x01}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.EncodeBool(testExample.input)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_DecodeBool(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation bool
	}{
		{label: "(0x00)", input: []byte{0x00}, expectation: false},
		{label: "(0x01)", input: []byte{0x01}, expectation: true},
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
