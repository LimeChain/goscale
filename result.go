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

func DecodeResult[T, E Encodable](buffer *bytes.Buffer, decodeValid func(*bytes.Buffer) (T, error), decodeErr func(*bytes.Buffer) (E, error)) (Result[Encodable], error) {
	hasError, err := DecodeBool(buffer)
	if err != nil {
		return Result[Encodable]{}, err
	}

	if hasError {
		value, err := decodeErr(buffer)
		if err != nil {
			return Result[Encodable]{}, err
		}

		return Result[Encodable]{
			HasError: hasError,
			Value:    value,
		}, nil
	}

	value, err := decodeValid(buffer)
	if err != nil {
		return Result[Encodable]{}, err
	}

	return Result[Encodable]{
		HasError: hasError,
		Value:    value,
	}, nil
}
