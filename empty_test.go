package goscale

import (
	"bytes"
	"testing"
)

func Test_EncodeEmpty(t *testing.T) {
	var examples = []struct {
		label  string
		input  Empty
		expect []byte
	}{
		{
			label:  "Empty",
			input:  Empty{},
			expect: nil,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}

			// when:
			e.input.Encode(buffer)

			// then:
			assertEqual(t, buffer.Bytes(), e.expect)
		})
	}
}
