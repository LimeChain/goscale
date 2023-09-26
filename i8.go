package goscale

import "bytes"

type I8 int8

func (value I8) Encode(buffer *bytes.Buffer) {
	U8(value).Encode(buffer)
}

func (value I8) Bytes() []byte {
	return U8(value).Bytes()
}

func DecodeI8(buffer *bytes.Buffer) I8 {
	decoder := Decoder{Reader: buffer}
	return I8(decoder.DecodeByte())
}
