package goscale

import (
	"bytes"
	"testing"
)

func Test_EncodeOptionBool(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Option[bool]
		expectation []byte
	}{
		{label: "Option(nil)", input: NewOptionEmpty(), expectation: []byte{0x00}},
		{label: "Option(true)", input: NewOptionBool(true), expectation: []byte{0x01}},
		{label: "Option(false)", input: NewOptionBool(false), expectation: []byte{0x02}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.EncodeOptionBool(testExample.input)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeOptionInt8(t *testing.T) {

	var testExamples = []struct {
		label       string
		input       Option[U8]
		expectation []byte
	}{
		{label: "Option(U8)", input: Option[U8]{true, 255}, expectation: []byte{0xff}},
	}

	for _, example := range testExamples {
		t.Run(example.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.Encode(example.input)
		})
		buffer := bytes.Buffer{}
		pd := verifyEncodingReturnDecoder(t,
			func(pe Encoder) { value.Encode(pe) },
			hex,
		)
		v2 := Option[U8]{}
		v2(pd)
		assertEqual(t, v2, value)
	}
}
