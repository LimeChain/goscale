package goscale

import "bytes"

type I32 int32

func (value I32) Encode(buffer *bytes.Buffer) {
	U32(value).Encode(buffer)
}

func (value I32) Bytes() []byte {
	return U32(value).Bytes()
}

func DecodeI32(buffer *bytes.Buffer) I32 {
	return I32(DecodeU32(buffer))
}
