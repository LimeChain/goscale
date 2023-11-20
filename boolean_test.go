package goscale

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
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
			buffer := &bytes.Buffer{}

			err := testExample.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, testExample.expectation, buffer.Bytes())
			assert.Equal(t, testExample.expectation, testExample.input.Bytes())
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
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			result, err := DecodeBool(buffer)

			assert.NoError(t, err)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_DecodeBoolPanics(t *testing.T) {
	var testExamples = []struct {
		label string
		input []byte
	}{
		{label: "(0xff)", input: []byte{0xff}},
		{label: "(0x3)", input: []byte{0x3}},
	}

	var testExamplesEmpty = []struct {
		label string
		input []byte
	}{
		{label: "([])", input: []byte{}},
		{label: "(nil)", input: nil},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			_, err := DecodeBool(buffer)
			assert.ErrorIs(t, errInvalidBoolRepresentation, err)
		})
	}

	for _, testExample := range testExamplesEmpty {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			_, err := DecodeBool(buffer)
			assert.ErrorIs(t, io.EOF, err)
		})
	}
}
