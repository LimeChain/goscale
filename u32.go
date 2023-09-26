package goscale

import (
	"bytes"
	"encoding/binary"
)

type U32 uint32

func (value U32) Encode(buffer *bytes.Buffer) {
	encoder := Encoder{Writer: buffer}
	encoder.Write(value.Bytes())
}

func (value U32) Bytes() []byte {
	result := make([]byte, 4)
	binary.LittleEndian.PutUint32(result, uint32(value))

	return result
}

func DecodeU32(buffer *bytes.Buffer) U32 {
	decoder := Decoder{Reader: buffer}
	result := make([]byte, 4)
	decoder.Read(result)
	return U32(binary.LittleEndian.Uint32(result))
}
