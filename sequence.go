package goscale

/*
	Ref: https://spec.polkadot.network/#defn-scale-list

	SCALE Sequence type translates to Go's slice and string types.
	SCALE FixedSequence type translates to Go's array type.
*/

import (
	"bytes"
	"reflect"
)

type Sequence[T Encodable] []T

func (seq Sequence[Encodable]) Encode(buffer *bytes.Buffer) {
	Compact(len(seq)).Encode(buffer)

	for _, v := range seq {
		if reflect.TypeOf(v).Kind() == reflect.Struct {
			EncodeTuple(v, buffer)
		} else {
			v.Encode(buffer)
		}
	}
}

func (seq Sequence[Encodable]) Bytes() []byte {
	buffer := &bytes.Buffer{}
	seq.Encode(buffer)

	return buffer.Bytes()
}

func DecodeSequence[T Encodable](buffer *bytes.Buffer) Sequence[T] {
	size := DecodeCompact(buffer)
	values := make([]T, size)

	for i := 0; i < len(values); i++ {
		values[i] = decodeByType(*new(T), buffer).(T)
	}
	return values
}

func DecodeSliceU8(buffer *bytes.Buffer) []U8 {
	return DecodeSequence[U8](buffer)
}

type FixedSequence[T Encodable] []T // TODO: https://github.com/LimeChain/goscale/issues/37

func (fseq FixedSequence[T]) Encode(buffer *bytes.Buffer) {
	for _, v := range fseq {
		if reflect.TypeOf(v).Kind() == reflect.Struct {
			EncodeTuple(v, buffer)
		} else {
			v.Encode(buffer)
		}
	}
}

func (fseq FixedSequence[T]) Bytes() []byte {
	buffer := &bytes.Buffer{}
	fseq.Encode(buffer)

	return buffer.Bytes()
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
	buffer.Write(value.Bytes())
}

func (value Str) Bytes() []byte {
	return Sequence[U8](StrToSliceU8(value)).Bytes()
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
