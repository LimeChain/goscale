package goscale

import "bytes"

/*
	https://spec.polkadot.network/#defn-option-type)

	SCALE Option Type ...

	https://spec.polkadot.network/#defn-result-type)

	SCALE Result Type ...
*/

// Option is a structure that can store a boolean or a missing value.
// Encoding rules are slightly different from other "Option" fields.
type Option[T Encodable] struct {
	hasValue bool
	Value    Encodable
}

func (o Option[T]) Encode(buffer *bytes.Buffer) {
	encoder := Encoder{Writer: buffer}
	if !o.hasValue {
		encoder.EncodeByte(0)
	} else {
		encoder.EncodeByte(1)
		o.Value.Encode(buffer)
	}
}

func DecodeOption[T Encodable](dec interface{}, buffer *bytes.Buffer) Option[T] {
	b := DecodeBool(buffer)

	option := Option[T]{
		hasValue: b == true,
	}

	if b {
		switch dec.(type) {
		case Bool:
			option.Value = DecodeBool(buffer)
		}
	}

	return option
}
