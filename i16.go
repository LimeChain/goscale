package goscale

import (
	"bytes"
)

type I16 int16

func (value I16) Encode(buffer *bytes.Buffer) error {
	return U16(value).Encode(buffer)
}

func (value I16) Bytes() []byte {
	return U16(value).Bytes()
}

func DecodeI16(buffer *bytes.Buffer) (I16, error) {
	value, err := DecodeU16(buffer)
	if err != nil {
		return 0, err
	}
	return I16(value), nil
}
