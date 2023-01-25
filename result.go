package goscale

/*
	Ref: https://spec.polkadot.network/#defn-result-type)

	SCALE Result Type.
*/

import (
	"bytes"
)

type Result[T Encodable] struct {
	HasError Bool
	Value    T
}

func (r Result[T]) Encode(buffer *bytes.Buffer) {
	(r.HasError).Encode(buffer)
	r.Value.Encode(buffer)
}

func (r Result[T]) Bytes() []byte {
	buffer := &bytes.Buffer{}
	r.Encode(buffer)

	return buffer.Bytes()
}

func DecodeResult[T Encodable](buffer *bytes.Buffer) Result[T] {
	hasError := DecodeBool(buffer)
	value := decodeByType(*new(T), buffer)

	return Result[T]{
		HasError: hasError,
		Value:    value.(T),
	}
}
