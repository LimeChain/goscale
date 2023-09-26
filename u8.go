package goscale

import "bytes"

type U8 uint8

func (value U8) Encode(buffer *bytes.Buffer) {
	// do not use value.Bytes() here: https://github.com/LimeChain/goscale/issues/77
	encoder := Encoder{Writer: buffer}
	encoder.EncodeByte(byte(value))
}

func (value U8) Bytes() []byte {
	return []byte{byte(value)}
}

func DecodeU8(buffer *bytes.Buffer) U8 {
	decoder := Decoder{Reader: buffer}
	result := make([]byte, 1)
	decoder.Read(result)
	return U8(result[0])
}
