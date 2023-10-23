package goscale

import (
	"bytes"
	"encoding/binary"
)

type U16 uint16

func (value U16) Encode(buffer *bytes.Buffer) {
	encoder := Encoder{Writer: buffer}
	encoder.Write(value.Bytes())
}

func (value U16) Bytes() []byte {
	result := make([]byte, 2)
	binary.LittleEndian.PutUint16(result, uint16(value))

	return result
}

func DecodeU16(buffer *bytes.Buffer) (U16, error) {
	decoder := Decoder{Reader: buffer}
	result := make([]byte, 2)
	err := decoder.Read(result)
	if err != nil {
		return 0, err
	}
	return U16(binary.LittleEndian.Uint16(result)), nil
}
