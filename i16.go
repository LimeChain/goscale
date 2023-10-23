package goscale

import "bytes"

type I16 int16

func (value I16) Encode(buffer *bytes.Buffer) {
	U16(value).Encode(buffer)
}

func (value I16) Bytes() []byte {
	return U16(value).Bytes()
}

func DecodeI16(buffer *bytes.Buffer) (I16, error) {
	return I16(DecodeU16(buffer)), nil
}
