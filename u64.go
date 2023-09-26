package goscale

import (
	"bytes"
	"encoding/binary"
)

type U64 uint64

func (value U64) Encode(buffer *bytes.Buffer) {
	encoder := Encoder{Writer: buffer}
	encoder.Write(value.Bytes())
}

func (value U64) Bytes() []byte {
	result := make([]byte, 8)
	binary.LittleEndian.PutUint64(result, uint64(value))

	return result
}

func DecodeU64(buffer *bytes.Buffer) U64 {
	decoder := Decoder{Reader: buffer}
	result := make([]byte, 8)
	decoder.Read(result)
	return U64(binary.LittleEndian.Uint64(result))
}
