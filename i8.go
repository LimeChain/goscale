package goscale

import "bytes"

type I8 int8

func (value I8) Encode(buffer *bytes.Buffer) {
	U8(value).Encode(buffer)
}

func (value I8) Bytes() []byte {
	return U8(value).Bytes()
}

func DecodeI8(buffer *bytes.Buffer) (I8, error) {
	decoder := Decoder{Reader: buffer}
	dec8, err := decoder.DecodeByte()
	if err != nil {
		return 0, err
	}
	return I8(dec8), nil
}
