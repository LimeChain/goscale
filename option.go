package goscale

/*
	https://spec.polkadot.network/#defn-option-type)

	SCALE Option Type ...

	https://spec.polkadot.network/#defn-result-type)

	SCALE Result Type ...
*/

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
type Option[T Encodable] struct {
	hasValue bool
	value    T
}

func (o Option[Encodable]) Encode(encoder *Encoder) {
	if !o.hasValue {
		encoder.EncodeByte(0)
	} else {
		encoder.EncodeByte(1)
		o.value.Encode(encoder)
	}
}

func DecodeOption[T Encodable](decoder *Decoder) Option[T] {
	b := decoder.DecodeBool()
	if b {
		//TODO: decode generic type
	}

	return Option[T]{
		hasValue: b == true,
		//value:
	}
}
