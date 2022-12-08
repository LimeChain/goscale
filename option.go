package goscale

/*
	Ref: https://spec.polkadot.network/#defn-option-type)

	Option is a varying data structure that can store an Encodable Value.
	HasValue indicates if Value is available.
*/

import "bytes"

type Option[T Encodable] struct {
	HasValue bool
	Value    T
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

func DecodeOption[T Encodable](buffer *bytes.Buffer) Option[T] {
	b := DecodeBool(buffer)

	option := Option[T]{
		HasValue: b == true,
	}

	if b {
		value := decodeByType(*new(T), buffer)
		option.Value = value.(T)
	}

	return option
}

type OptionBool Option[Bool]

func (o OptionBool) Encode(buffer *bytes.Buffer) {
	encoder := Encoder{Writer: buffer}
	if !o.HasValue {
		encoder.EncodeByte(0)
	} else {
		if o.Value {
			encoder.EncodeByte(1)
		} else {
			encoder.EncodeByte(2)
		}
	}
}

func DecodeOptionBool(buffer *bytes.Buffer) OptionBool {
	decoder := Decoder{Reader: buffer}
	b := decoder.DecodeByte()

	result := OptionBool{}

	switch b {
	case 0:
		result.HasValue = false
	case 1:
		result.HasValue = true
		result.Value = true
	case 2:
		result.HasValue = true
		result.Value = false
	default:
		panic("invalid OptionBool representation")
	}

	return result
}
