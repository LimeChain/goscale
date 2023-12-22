package goscale

import (
	"bytes"
	"math/big"
)

type U8 uint8

func (value U8) Encode(buffer *bytes.Buffer) error {
	// do not use value.Bytes() here: https://github.com/LimeChain/goscale/issues/77
	encoder := Encoder{Writer: buffer}
	return encoder.EncodeByte(byte(value))
}

func (value U8) Bytes() []byte {
	return []byte{byte(value)}
}

func NewU8(n uint8) U8 {
	return U8(n)
}

func (value U8) ToBigInt() *big.Int {
	return new(big.Int).SetUint64(uint64(value))
}

func (value U8) Interface() Numeric {
	return value
}

func DecodeU8(buffer *bytes.Buffer) (U8, error) {
	decoder := Decoder{Reader: buffer}
	b, err := decoder.DecodeByte()
	return U8(b), err
}
