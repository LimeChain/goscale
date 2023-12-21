package goscale

import (
	"bytes"
	"math/big"
)

type I64 int64

func (value I64) Encode(buffer *bytes.Buffer) error {
	return U64(value).Encode(buffer)
}

func (value I64) Bytes() []byte {
	return U64(value).Bytes()
}

func NewI64(n int64) I64 {
	return I64(n)
}

func (value I64) ToBigInt() *big.Int {
	return new(big.Int).SetInt64(int64(value))
}

func DecodeI64(buffer *bytes.Buffer) (I64, error) {
	value, err := DecodeU64(buffer)
	if err != nil {
		return 0, err
	}
	return I64(value), nil
}
