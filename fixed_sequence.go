package goscale

import "bytes"

type FixedSequence[T Encodable] struct {
	Value []T
}

func (fa FixedSequence[T]) Encode(buffer *bytes.Buffer) {
	for _, value := range fa.Value {
		value.Encode(buffer)
	}
}

func DecodeFixedSequence(len int, enc Encodable, buffer *bytes.Buffer) FixedSequence[Encodable] {
	result := make([]Encodable, len)

	for i := 0; i < len; i++ {
		result[i] = decodeByType(enc, buffer)
	}

	return FixedSequence[Encodable]{
		Value: result,
	}
}
