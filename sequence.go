package goscale

/*
	Ref: https://spec.polkadot.network/#defn-scale-list

	SCALE Sequence type translates to Go's slice and string types.
*/

import (
	"bytes"
)

type Sequence[T Encodable] struct {
	Values []T
}

func (seq Sequence[Encodable]) Encode(buffer *bytes.Buffer) {
	Compact(len(seq.Values)).Encode(buffer)
	for _, v := range seq.Values {
		v.Encode(buffer)
	}
}

func DecodeSequenceU8(buffer *bytes.Buffer) Sequence[U8] {
	return Sequence[U8]{Values: DecodeSliceU8(buffer)}
}

func DecodeSliceU8(buffer *bytes.Buffer) []U8 {
	size := DecodeCompact(buffer)
	values := make([]U8, size)
	for i := 0; i < len(values); i++ {
		values[i] = DecodeU8(buffer)
	}
	return values
}
