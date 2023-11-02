package goscale

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
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
			buffer := &bytes.Buffer{}

			err := e.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, buffer.Bytes(), e.expect)
			assert.Equal(t, e.input.Bytes(), []byte{})
		})
	}
}
