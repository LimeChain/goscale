package goscale

/*
	Ref: https://spec.polkadot.network/#defn-scale-list

	SCALE Fixed Sequence type translates to Go's array type.
*/

import "bytes"

type FixedSequence[T Encodable] struct {
	// TODO needs to be an array,
	// but currently it is not possible to be parameterized
	Values []T
}

func (fseq FixedSequence[T]) Encode(buffer *bytes.Buffer) {
	for _, value := range fseq.Values {
		value.Encode(buffer)
	}
}

func DecodeFixedSequence[T Encodable](size int, buffer *bytes.Buffer) FixedSequence[T] {
	result := make([]T, size)

	for i := 0; i < size; i++ {
		result[i] = decodeByType(*new(T), buffer).(T)
	}

	return FixedSequence[T]{Values: result}
}
