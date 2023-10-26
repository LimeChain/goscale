package goscale

/*
	Ref: https://spec.polkadot.network/#defn-varrying-data-type)

	SCALE Varying Data Type.
*/

import (
	"bytes"
	"errors"
	"math"
)

var (
	errDecodingFuncNotFound = errors.New("varying data: decode func not found")
	errExceedsU8Length      = errors.New("exceeds uint8 length")
)

type VaryingData []Encodable

func NewVaryingData(values ...Encodable) VaryingData {
	if len(values) > math.MaxUint8 {
		panic("exceeds uint8 length")
	}

	result := make([]Encodable, 0, len(values))
	result = append(result, values...)

	return result
}

func (vd VaryingData) Encode(buffer *bytes.Buffer) {
	for _, v := range vd {
		v.Encode(buffer)
	}
}

func DecodeVaryingData(decodeFuncs []func(buffer *bytes.Buffer) []Encodable, buffer *bytes.Buffer) (VaryingData, error) {
	funcsLen := len(decodeFuncs)
	if funcsLen > math.MaxUint8 {
		return VaryingData{}, errExceedsU8Length
	}

	index, err := DecodeU8(buffer)
	if err != nil {
		return VaryingData{}, err
	}
	if int(index) > funcsLen-1 {
		return VaryingData{}, errDecodingFuncNotFound
	}

	decoded := decodeFuncs[index](buffer)

	args := make([]Encodable, 0, len(decoded)+1)
	args = append(args, index)
	args = append(args, decoded...)

	return NewVaryingData(args...), nil
}

func (vd VaryingData) Bytes() []byte {
	return EncodedBytes(vd)
}
