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

// func Test_EncodeOptionInt8(t *testing.T) {
// 	tests := map[OptionInt8]string{
// 		OptionInt8{true, 1}:  "01 01",
// 		OptionInt8{true, -1}: "01 ff",
// 		OptionInt8{false, 0}: "00",
// 	}

// 	for value, hex := range tests {
// 		pd := verifyEncodingReturnDecoder(t,
// 			func(pe Encoder) { value.Encode(pe) },
// 			hex,
// 		)
// 		v2 := OptionInt8{}
// 		v2.Decode(pd)
// 		assertEqual(t, v2, value)
// 	}
// }
