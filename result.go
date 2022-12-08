package goscale

/*
	Ref: https://spec.polkadot.network/#defn-result-type)

	SCALE Result Type.
*/

import (
	"bytes"
)

type Result[T Encodable] struct {
	ok    Bool
	value T
}

func (r Result[T]) Encode(buffer *bytes.Buffer) {
	r.ok.Encode(buffer)
	r.value.Encode(buffer)
}

func DecodeResult[T Encodable](buffer *bytes.Buffer) Result[T] {
	ok := DecodeBool(buffer)
	value := decodeByType(*new(T), buffer)
	return Result[T]{
		ok:    ok,
		value: value.(T),
	}
}
