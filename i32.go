package goscale

import (
	"bytes"
	"math/big"
)

type I32 int32

func (value I32) Encode(buffer *bytes.Buffer) error {
	return U32(value).Encode(buffer)
}

func (value I32) Bytes() []byte {
	return U32(value).Bytes()
}

func NewI32(n int32) I32 {
	return I32(n)
}

func (value I32) ToBigInt() *big.Int {
	return new(big.Int).SetInt64(int64(value))
}

func DecodeI32(buffer *bytes.Buffer) (I32, error) {
	value, err := DecodeU32(buffer)
	if err != nil {
		return 0, err
	}
	return I32(value), nil
}
