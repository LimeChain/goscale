package goscale

import "bytes"

type I64 int64

func (value I64) Encode(buffer *bytes.Buffer) {
	U64(value).Encode(buffer)
}

func (value I64) Bytes() []byte {
	return U64(value).Bytes()
}

func DecodeI64(buffer *bytes.Buffer) I64 {
	return I64(DecodeU64(buffer))
}
