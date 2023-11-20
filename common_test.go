package goscale

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type encodableType struct{}

func (e encodableType) Encode(buffer *bytes.Buffer) error {
	return errors.New("error")
}

func (e encodableType) Bytes() []byte {
	return []byte{}
}

func Test_EncodeEach(t *testing.T) {
	buffer := &bytes.Buffer{}

	err := EncodeEach(buffer, Bool(true), U8(127), I16(-128))

	assert.NoError(t, err)
	assert.Equal(t, []byte{0x01, 0x7f, 0x80, 0xff}, buffer.Bytes())
}

func Test_EncodeEach_Error(t *testing.T) {
	buffer := &bytes.Buffer{}

	err := EncodeEach(buffer, Bool(true), encodableType{}, U8(127))

	assert.Error(t, err)
	assert.Equal(t, []byte{0x01}, buffer.Bytes())
}
