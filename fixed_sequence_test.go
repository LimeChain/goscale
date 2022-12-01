package goscale

import (
	"bytes"
	"testing"
)

func Test_EncodeFixedSequence(t *testing.T) {
	var examples = []struct {
		label  string
		input  FixedSequence[Encodable]
		expect []byte
	}{
		{label: "Encode FixedSequence[U8]", input: FixedSequence[Encodable]{[]Encodable{U8(5), U8(6), U8(7)}}, expect: []byte{0x5, 0x6, 0x7}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			e.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), e.expect)
		})
	}
}

func Test_DecodeFixedSequence(t *testing.T) {
	var examples = []struct {
		label  string
		input  []byte
		expect FixedSequence[Encodable]
	}{
		{label: "Decode FixedSequence[U8]", input: []byte{0x5, 0x6, 0x7}, expect: FixedSequence[Encodable]{[]Encodable{U8(5), U8(6), U8(7)}}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeFixedSequence(len(e.input), U8(0), buffer)

			// then:
			assertEqual(t, result, e.expect)
		})
	}
}
