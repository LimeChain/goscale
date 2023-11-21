package goscale

/*
	Ref: https://spec.polkadot.network/#defn-option-type)

	Option is a varying data structure that can store an Encodable Value.
	HasValue indicates if Value is available.
*/

import (
	"bytes"
	"errors"
)

var (
	errInvalidOptionBoolRepresentation = errors.New("invalid OptionBool representation")
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

func (o Option[T]) Encode(buffer *bytes.Buffer) error {
	encoder := Encoder{Writer: buffer}
	if !o.HasValue {
		err := encoder.EncodeByte(0)
		if err != nil {
			return err
		}
	} else {
		err := encoder.EncodeByte(1)
		if err != nil {
			return err
		}

		err = o.Value.Encode(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (o Option[T]) Bytes() []byte {
	return EncodedBytes(o)
}

func DecodeOption[T Encodable](buffer *bytes.Buffer) (Option[T], error) {
	b, err := DecodeBool(buffer)
	if err != nil {
		return Option[T]{}, err
	}

	option := Option[T]{
		HasValue: b,
	}

	if b {
		value, err := decodeByType(*new(T), buffer)
		if err != nil {
			return Option[T]{}, err
		}
		option.Value = value.(T)
	}

	return option, nil
}

func DecodeOptionWith[T Encodable](buffer *bytes.Buffer, decodeFunc func(buffer *bytes.Buffer) (T, error)) (Option[T], error) {
	option := Option[T]{HasValue: false}

	b, err := DecodeBool(buffer)
	if err != nil {
		return Option[T]{}, err
	}
	if b {
		option.HasValue = true
		val, err := decodeFunc(buffer)
		if err != nil {
			return Option[T]{}, err
		}
		option.Value = val
	}

	return option, nil
}

type OptionBool Option[Bool]

func (o OptionBool) Encode(buffer *bytes.Buffer) error {
	encoder := Encoder{Writer: buffer}
	if !o.HasValue {
		err := encoder.EncodeByte(0)
		if err != nil {
			return err
		}
	} else {
		if o.Value {
			err := encoder.EncodeByte(1)
			if err != nil {
				return err
			}
		} else {
			err := encoder.EncodeByte(2)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (o OptionBool) Bytes() []byte {
	buffer := &bytes.Buffer{}
	o.Encode(buffer)
	return buffer.Bytes()
}

func DecodeOptionBool(buffer *bytes.Buffer) (OptionBool, error) {
	decoder := Decoder{Reader: buffer}
	b, err := decoder.DecodeByte()
	if err != nil {
		return OptionBool{}, err
	}

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
		return OptionBool{}, errInvalidOptionBoolRepresentation
	}

	return result, nil
}
