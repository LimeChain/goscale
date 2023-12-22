package goscale

import (
	"bytes"
	"encoding/binary"
	"math/big"
)

type U16 uint16

func (value U16) Encode(buffer *bytes.Buffer) error {
	encoder := Encoder{Writer: buffer}
	return encoder.Write(value.Bytes())
}

func NewU16(n uint16) U16 {
	return U16(n)
}

func (value U16) ToBigInt() *big.Int {
	return new(big.Int).SetUint64(uint64(value))
}

func (value U16) Bytes() []byte {
	result := make([]byte, 2)
	binary.LittleEndian.PutUint16(result, uint16(value))

	return result
}

func DecodeU16(buffer *bytes.Buffer) (U16, error) {
	decoder := Decoder{Reader: buffer}
	result := make([]byte, 2)
	err := decoder.Read(result)
	if err != nil {
		return 0, err
	}
	return U16(binary.LittleEndian.Uint16(result)), nil
}
