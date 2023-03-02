package goscale

/*
	Ref: https://spec.polkadot.network/#defn-varrying-data-type)

	SCALE Varying Data Type.
*/

import (
	"bytes"
	"math"
)

type VaryingData []Encodable

func NewVaryingData(values ...Encodable) VaryingData {
	if len(values) > math.MaxUint8 {
		panic("exceeds uint8 length")
	}

	var result []Encodable
	result = append(result, values...)

	return result
}

func (vd VaryingData) Encode(buffer *bytes.Buffer) {
	for _, v := range vd {
		v.Encode(buffer)
	}
}

func DecodeVaryingData(decodeFuncs []func(buffer *bytes.Buffer) []Encodable, buffer *bytes.Buffer) VaryingData {
	funcsLen := len(decodeFuncs)
	if funcsLen > math.MaxUint8 {
		panic("exceeds uint8 length")
	}

	index := DecodeU8(buffer)
	if int(index) > funcsLen-1 {
		panic("varying data: decode func not found")
	}

	decoded := decodeFuncs[index](buffer)

	var args []Encodable
	args = append(args, index)
	args = append(args, decoded...)

	return NewVaryingData(args...)
}
func (vd VaryingData) Bytes() []byte {
	return EncodedBytes(vd)
}
