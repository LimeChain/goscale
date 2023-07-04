package goscale

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EncodeDictionaryStrBool(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Dictionary[Str, Bool]
		expectation []byte
	}{
		{
			label: "Dictionary(aaa: true, bbb: false, ccc: true)",
			input: Dictionary[Str, Bool]{Str("aaa"): true, Str("bbb"): false, Str("ccc"): true},
			expectation: []byte{
				0x0c,
				0x0c, 0x61, 0x61, 0x61, 0x01,
				0x0c, 0x62, 0x62, 0x62, 0x00,
				0x0c, 0x63, 0x63, 0x63, 0x01,
			},
		},
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

func Test_DecodeDictionaryStrBool(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation Dictionary[Str, Bool]
	}{
		{
			label:       "(0x0c0c636363010c626262000c61616101)",
			input:       []byte{0x0c, 0x0c, 0x63, 0x63, 0x63, 0x01, 0x0c, 0x62, 0x62, 0x62, 0x00, 0x0c, 0x61, 0x61, 0x61, 0x01},
			expectation: Dictionary[Str, Bool]{Str("aaa"): true, Str("bbb"): false, Str("ccc"): true},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			result := DecodeDictionary[Str, Bool](buffer)

			assert.Equal(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeDictionaryU8Str(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Dictionary[U8, Str]
		expectation []byte
	}{
		{
			label: "Dictionary(1: aaa, 2: bbb, 3: ccc)",
			input: Dictionary[U8, Str]{1: Str("aaa"), 2: Str("bbb"), 3: Str("ccc")},
			expectation: []byte{
				0x0c,
				0x01, 0x0c, 0x61, 0x61, 0x61,
				0x02, 0x0c, 0x62, 0x62, 0x62,
				0x03, 0x0c, 0x63, 0x63, 0x63,
			},
		},
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

func Test_DecodeDictionaryU8Str(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation Dictionary[U8, Str]
	}{
		{
			label: "()",
			input: []byte{
				0x0c,
				0x03, 0x0c, 0x63, 0x63, 0x63,
				0x02, 0x0c, 0x62, 0x62, 0x62,
				0x01, 0x0c, 0x61, 0x61, 0x61,
			},
			expectation: Dictionary[U8, Str]{1: Str("aaa"), 2: Str("bbb"), 3: Str("ccc")},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			result := DecodeDictionary[U8, Str](buffer)

			assert.Equal(t, result, testExample.expectation)
		})
	}
}
