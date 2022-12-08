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
	for i, v := range vd {
		U8(i).Encode(buffer)
		v.Encode(buffer)
	}
}

func DecodeVaryingData(values []Encodable, buffer *bytes.Buffer) VaryingData {
	vLen := len(values)
	if vLen > math.MaxUint8 {
		panic("exceeds uint8 length")
	}

	result := make([]Encodable, vLen)
	for i := 0; i < vLen; i++ {
		index := DecodeU8(buffer)
		value := decodeByType(values[index], buffer)

		result[index] = value
	}

	return result
}
