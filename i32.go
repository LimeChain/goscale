package goscale

import "bytes"

type I32 int32

func (value I32) Encode(buffer *bytes.Buffer) error {
	return U32(value).Encode(buffer)
}

func (value I32) Bytes() []byte {
	return U32(value).Bytes()
}

func DecodeI32(buffer *bytes.Buffer) (I32, error) {
	dec32, err := DecodeU32(buffer)
	if err != nil {
		return 0, err
	}
	return I32(dec32), nil
}
