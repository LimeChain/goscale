package goscale

import "bytes"

type I64 int64

func (value I64) Encode(buffer *bytes.Buffer) error {
	return U64(value).Encode(buffer)
}

func (value I64) Bytes() []byte {
	return U64(value).Bytes()
}

func DecodeI64(buffer *bytes.Buffer) (I64, error) {
	dec64, err := DecodeU64(buffer)
	if err != nil {
		return 0, err
	}
	return I64(dec64), nil
}
