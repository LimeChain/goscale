package goscale

/*
	Ref: https://spec.polkadot.network/#defn-option-type)

	Option is a varying data structure that can store an Encodable Value.
	HasValue indicates if Value is available.
*/

import "bytes"

type Option[T Encodable] struct {
	HasValue bool
	Value    Encodable
}

func (o Option[T]) Encode(buffer *bytes.Buffer) {
	encoder := Encoder{Writer: buffer}
	if !o.HasValue {
		encoder.EncodeByte(0)
	} else {
		encoder.EncodeByte(1)
		o.Value.Encode(buffer)
	}
}

func DecodeOption[T Encodable](dec Encodable, buffer *bytes.Buffer) Option[T] {
	b := DecodeBool(buffer)

	option := Option[T]{
		HasValue: b == true,
	}

	if b {
		option.Value = decodeByType(dec, buffer)
	}

	return option
}
