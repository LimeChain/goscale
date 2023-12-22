package goscale

import (
	"bytes"
	"encoding/binary"
	"math/big"
)

type U64 uint64

func (value U64) Encode(buffer *bytes.Buffer) error {
	encoder := Encoder{Writer: buffer}
	return encoder.Write(value.Bytes())
}

func (value U64) Bytes() []byte {
	result := make([]byte, 8)
	binary.LittleEndian.PutUint64(result, uint64(value))

	return result
}

func (value U64) Interface() Numeric {
	return value
}

func NewU64(n uint64) U64 {
	return U64(n)
}

func (value U64) ToBigInt() *big.Int {
	return new(big.Int).SetUint64(uint64(value))
}

func DecodeU64(buffer *bytes.Buffer) (U64, error) {
	decoder := Decoder{Reader: buffer}
	result := make([]byte, 8)
	err := decoder.Read(result)
	if err != nil {
		return 0, err
	}
	return U64(binary.LittleEndian.Uint64(result)), nil
}
