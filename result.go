package goscale

/*
	Ref: https://spec.polkadot.network/#defn-result-type)

	SCALE Result Type.
*/

import (
	"bytes"
)

type Result[T Encodable] struct {
	Ok    Bool
	Value T
}

func (r Result[T]) Encode(buffer *bytes.Buffer) {
	(!r.Ok).Encode(buffer)
	r.Value.Encode(buffer)
}

func (r Result[T]) Bytes() []byte {
	buffer := &bytes.Buffer{}
	r.Encode(buffer)

	return buffer.Bytes()
}

func DecodeResult[T Encodable](buffer *bytes.Buffer) Result[T] {
	ok := !DecodeBool(buffer)
	value := decodeByType(*new(T), buffer)

	return Result[T]{
		Ok:    ok,
		Value: value.(T),
	}
}
