package goscale

import (
	"bytes"
)

type VaryingData []Encodable

func NewVaryingData(values ...Encodable) VaryingData {
	if len(values) > 255 {
		panic("exceeds uint8 length")
	}

	var result []Encodable
	for _, v := range values {
		result = append(result, v)
	}

	return result
}

func (vd VaryingData) Encode(buffer *bytes.Buffer) {
	for k, v := range vd {
		U8(k).Encode(buffer)
		v.Encode(buffer)
	}
}

func DecodeVaryingData(values []Encodable, buffer *bytes.Buffer) VaryingData {
	vLen := len(values)
	if vLen > 255 {
		panic("exceeds uint8 length")
	}

	result := make([]Encodable, vLen)
	for i := 0; i < len(values); i++ {

		key := DecodeU8(buffer)
		value := decodeByType(values[key], buffer)

		result[key] = value
	}

	return result
}
