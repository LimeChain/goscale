package goscale

/*
	Ref: https://spec.polkadot.network/#defn-scale-list

	SCALE Sequence type translates to Go's slice and string types.
	SCALE FixedSequence type translates to Go's array type.
*/

import (
	"bytes"
)

type Sequence[T Encodable] []T
type FixedSequence[T Encodable] []T // TODO: https://github.com/LimeChain/goscale/issues/37

func (seq Sequence[Encodable]) Encode(buffer *bytes.Buffer) {
	Compact(len(seq)).Encode(buffer)
	for _, v := range seq {
		v.Encode(buffer)
	}
}

func DecodeSliceU8(buffer *bytes.Buffer) []U8 {
	size := DecodeCompact(buffer)
	values := make([]U8, size)
	for i := 0; i < len(values); i++ {
		values[i] = DecodeU8(buffer)
	}
	return values
}

func (fseq FixedSequence[T]) Encode(buffer *bytes.Buffer) {
	for _, value := range fseq {
		value.Encode(buffer)
	}
}

func DecodeFixedSequence[T Encodable](size int, buffer *bytes.Buffer) FixedSequence[T] {
	result := make([]T, size)
	for i := 0; i < size; i++ {
		result[i] = decodeByType(*new(T), buffer).(T)
	}
	return FixedSequence[T](result)
}

// additional helper type
type Str string

func (value Str) Encode(buffer *bytes.Buffer) {
	seq := Sequence[U8](StrToSliceU8(value))
	seq.Encode(buffer)
}

func DecodeStr(buffer *bytes.Buffer) Str {
	return SliceU8ToStr(DecodeSliceU8(buffer))
}

func StrToSliceU8(s Str) []U8 {
	result := make([]U8, len(s))
	for i, v := range []byte(s) {
		result[i] = U8(v) // TODO: https://github.com/LimeChain/goscale/issues/38
	}
	return result
}

func SliceU8ToStr(values []U8) Str {
	result := make([]byte, len(values))
	for i, v := range values {
		result[i] = byte(v)
	}
	return Str(result)
}
