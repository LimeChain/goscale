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
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
			assertEqual(t, testExample.input.Bytes(), testExample.expectation)
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

			result := DecodeBool(buffer)

			assertEqual(t, result, testExample.expectation)
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
		{label: "([])", input: []byte{}},
		{label: "(nil)", input: nil},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			assertPanic(t, func() {
				DecodeBool(buffer)
			})
		})
	}
}
