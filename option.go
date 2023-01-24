package goscale

/*
	Ref: https://spec.polkadot.network/#defn-option-type)

	Option is a varying data structure that can store an Encodable Value.
	HasValue indicates if Value is available.
*/

import (
	"bytes"
)

type Option[T Encodable] struct {
	HasValue Bool
	Value    T
}

func NewOption[T Encodable](value Encodable) Option[T] {
	switch value := value.(type) {
	case T:
		return Option[T]{HasValue: true, Value: value}
	case nil:
		return Option[T]{HasValue: false}
	default:
		panic("invalid value type for Option[T]")
	}
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
		HasValue: b,
	}

	if b {
		value := decodeByType(*new(T), buffer)
		option.Value = value.(T)
	}

	return option
}

func (o Option[T]) Bytes() []byte {
	buffer := &bytes.Buffer{}
	o.Encode(buffer)
	return buffer.Bytes()
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

func (o OptionBool) Bytes() []byte {
	buffer := &bytes.Buffer{}
	o.Encode(buffer)

	return buffer.Bytes()
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
