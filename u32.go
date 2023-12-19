package goscale

import (
	"bytes"
	"encoding/binary"
	"math/big"
)

type U32 uint32

func (value U32) Encode(buffer *bytes.Buffer) error {
	encoder := Encoder{Writer: buffer}
	return encoder.Write(value.Bytes())
}

func (value U32) Bytes() []byte {
	result := make([]byte, 4)
	binary.LittleEndian.PutUint32(result, uint32(value))

	return result
}

func (value U32) ToBigInt() *big.Int {
	return new(big.Int).SetUint64(uint64(value))
}

func DecodeU32(buffer *bytes.Buffer) (U32, error) {
	decoder := Decoder{Reader: buffer}
	result := make([]byte, 4)
	err := decoder.Read(result)
	if err != nil {
		return 0, err
	}
	return U32(binary.LittleEndian.Uint32(result)), nil
}
