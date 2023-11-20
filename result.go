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

func (r Result[T]) Encode(buffer *bytes.Buffer) error {
	err := (r.HasError).Encode(buffer)
	if err != nil {
		return err
	}
	err = r.Value.Encode(buffer)
	if err != nil {
		return err
	}
	return nil
}

func (r Result[T]) Bytes() []byte {
	return EncodedBytes(r)
}

func DecodeResult[T Encodable](buffer *bytes.Buffer) (Result[T], error) {
	hasError, err := DecodeBool(buffer)
	if err != nil {
		return Result[T]{}, err
	}
	value, err := decodeByType(*new(T), buffer)
	if err != nil {
		return Result[T]{}, err
	}

	return Result[T]{
		HasError: hasError,
		Value:    value.(T),
	}, nil
}
