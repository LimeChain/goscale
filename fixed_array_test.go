package goscale

import (
	"bytes"
	"testing"
)

func Test_EncodeFixedArray(t *testing.T) {
	var examples = []struct {
		label  string
		input  FixedArray[Encodable]
		expect []byte
	}{
		{label: "Encode FixedArray[U8]", input: FixedArray[Encodable]{[]Encodable{U8(5), U8(6), U8(7)}}, expect: []byte{0x5, 0x6, 0x7}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			e.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), e.expect)
		})
	}
}

func Test_DecodeFixedArray(t *testing.T) {
	var examples = []struct {
		label  string
		input  []byte
		expect FixedArray[Encodable]
	}{
		{label: "Decode FixedArray[U8]", input: []byte{0x5, 0x6, 0x7}, expect: FixedArray[Encodable]{[]Encodable{U8(5), U8(6), U8(7)}}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeFixedArray(len(e.input), U8(0), buffer)

			// then:
			assertEqual(t, result, e.expect)
		})
	}
}
