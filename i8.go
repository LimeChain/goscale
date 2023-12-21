package goscale

import (
	"bytes"
	"math/big"
)

type I8 int8

func (value I8) Encode(buffer *bytes.Buffer) error {
	return U8(value).Encode(buffer)
}

func (value I8) Bytes() []byte {
	return U8(value).Bytes()
}

func NewI8(n int8) I8 {
	return I8(n)
}

func (value I8) ToBigInt() *big.Int {
	return new(big.Int).SetInt64(int64(value))
}

func DecodeI8(buffer *bytes.Buffer) (I8, error) {
	decoder := Decoder{Reader: buffer}
	value, err := decoder.DecodeByte()
	if err != nil {
		return 0, err
	}
	return I8(value), nil
}
