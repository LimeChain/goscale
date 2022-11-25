package goscale

import "bytes"

// Option is a varying data structure that can store an Encodable Value.
// HasValue indicates if Value is available.
// Ref: https://spec.polkadot.network/#defn-option-type)
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
		switch dec.(type) {
		case Bool:
			option.Value = DecodeBool(buffer)
		case U8:
			option.Value = DecodeU8(buffer)
		case I8:
			option.Value = DecodeI8(buffer)
		case U16:
			option.Value = DecodeU16(buffer)
		case I16:
			option.Value = DecodeI16(buffer)
		case U32:
			option.Value = DecodeU32(buffer)
		case I32:
			option.Value = DecodeI32(buffer)
		case U64:
			option.Value = DecodeU64(buffer)
		case I64:
			option.Value = DecodeI64(buffer)
		case Compact:
			option.Value = DecodeCompact(buffer)
		case Sequence[U8]:
			option.Value = DecodeSequenceU8(buffer)
		default:
			panic("type not found")
		}
	}

	return option
}
