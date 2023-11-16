package goscale

import "bytes"

type U8 uint8

func (value U8) Encode(buffer *bytes.Buffer) error {
	// do not use value.Bytes() here: https://github.com/LimeChain/goscale/issues/77
	encoder := Encoder{Writer: buffer}
	return encoder.EncodeByte(byte(value))
}

func (value U8) Bytes() []byte {
	return []byte{byte(value)}
}

func DecodeU8(buffer *bytes.Buffer) (U8, error) {
	decoder := Decoder{Reader: buffer}
	b, err := decoder.DecodeByte()
	return U8(b), err
}
