package goscale

import "bytes"

type FixedArray[T Encodable] struct {
	Value []T
}

func (fa FixedArray[T]) Encode(buffer *bytes.Buffer) {
	for _, value := range fa.Value {
		value.Encode(buffer)
	}
}

func DecodeFixedArray(len int, enc Encodable, buffer *bytes.Buffer) FixedArray[Encodable] {
	result := make([]Encodable, len)

	for i := 0; i < len; i++ {
		result[i] = decodeByType(enc, buffer)
	}

	return FixedArray[Encodable]{
		Value: result,
	}
}
