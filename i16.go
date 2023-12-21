package goscale

import (
	"bytes"
	"math/big"
)

type I16 int16

func (value I16) Encode(buffer *bytes.Buffer) error {
	return U16(value).Encode(buffer)
}

func NewI16(n int16) I16 {
	return I16(n)
}

func (value I16) ToBigInt() *big.Int {
	return new(big.Int).SetInt64(int64(value))
}

func (value I16) Bytes() []byte {
	return U16(value).Bytes()
}

func DecodeI16(buffer *bytes.Buffer) (I16, error) {
	value, err := DecodeU16(buffer)
	if err != nil {
		return 0, err
	}
	return I16(value), nil
}
