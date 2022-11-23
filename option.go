package goscale

/*
	https://spec.polkadot.network/#defn-option-type)

	SCALE Option Type ...

	https://spec.polkadot.network/#defn-result-type)

	SCALE Result Type ...
*/

import "strconv"

// Encodeable is an interface that defines a custom encoding rules for a data type.
// Should be defined for structs (not pointers to them).
type Encodeable interface {
	Encode(enc Encoder)
}

// Decodeable is an interface that defines a custom encoding rules for a data type.
// Should be defined for pointers to structs.
// Decode populates this structure from a stream (overwriting the current contents), return false on failure
type Decodeable interface {
	Decode(dec Decoder)
}

// Option is a structure that can store a boolean or a missing value.
// Encoding rules are slightly different from other "Option" fields.
type Option[T any] struct {
	hasValue bool
	value    T
}

func NewOptionEmpty() Option[bool] {
	return Option[bool]{false, false}
}

func NewOptionBool(value bool) Option[bool] {
	return Option[bool]{true, value}
}

func (enc Encoder) EncodeOptionBool(o Option[bool]) {
	if !o.hasValue {
		enc.EncodeByte(0x00)
	} else {
		switch o.value {
		case true:
			enc.EncodeByte(0x01)
		case false:
			enc.EncodeByte(0x02)
		}
	}
}

func (dec Decoder) DecodeOptionBool() {
	o := Option[bool]{}

	b := dec.DecodeByte()

	switch b {
	case 0:
		o.hasValue = false
		o.value = false
	case 1:
		o.hasValue = true
		o.value = true
	case 2:
		o.hasValue = true
		o.value = false
	default:
		panic("Unknown byte prefix for encoded OptionBool: " + strconv.Itoa(int(b)))
	}
}

func (enc Encoder) EncodeOption(hasValue bool, value Encodeable) {
	if !hasValue {
		enc.EncodeByte(0)
	} else {
		enc.EncodeByte(1)
		value.Encode(enc)
	}
}

// DecodeOption decodes an optionally available value into a boolean presence field and a value.
func (dec Decoder) DecodeOption(hasValue *bool, valuePointer Decodeable) {
	b := dec.DecodeByte()
	switch b {
	case 0:
		*hasValue = false
	case 1:
		*hasValue = true
		valuePointer.Decode(dec)
	default:
		panic("Unknown byte prefix for encoded Option: " + strconv.Itoa(int(b)))
	}
}
