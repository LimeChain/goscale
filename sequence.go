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

func (seq Sequence[Encodable]) Encode(buffer *bytes.Buffer) {
	ToCompact(len(seq)).Encode(buffer)

	for _, v := range seq {
		//if reflect.TypeOf(v).Kind() == reflect.Struct {
		//	EncodeTuple(v, buffer)
		//} else {
		v.Encode(buffer)
		//}
	}
}

func (seq Sequence[Encodable]) Bytes() []byte {
	return EncodedBytes(seq)
}

func DecodeSequence[T Encodable](buffer *bytes.Buffer) Sequence[T] {
	size := DecodeCompact(buffer)
	v := size.ToBigInt()
	values := make([]T, v.Int64())

	for i := 0; i < len(values); i++ {
		values[i] = decodeByType(*new(T), buffer).(T)
	}
	return values
}

func DecodeSequenceWith[T Encodable](buffer *bytes.Buffer, decodeFunc func(buffer *bytes.Buffer) T) Sequence[T] {
	size := DecodeCompact(buffer)
	v := size.ToBigInt()
	values := make([]T, v.Int64())

	for i := 0; i < len(values); i++ {
		values[i] = decodeFunc(buffer)
	}
	return values
}

func DecodeSliceU8(buffer *bytes.Buffer) []U8 {
	return DecodeSequence[U8](buffer)
}

type FixedSequence[T Encodable] []T // TODO: https://github.com/LimeChain/goscale/issues/37

// Initializes with the specified size and provides size checks (at least at runtime)
func NewFixedSequence[T Encodable](size int, values ...T) FixedSequence[T] {
	if len(values) != size {
		panic("values size is out of the specified bound")
	}

	fseq := make(FixedSequence[T], size)
	for i, v := range values {
		fseq[i] = T(v)
	}
	return fseq
}

func (fseq FixedSequence[T]) Encode(buffer *bytes.Buffer) {
	for _, v := range fseq {
		//if reflect.TypeOf(v).Kind() == reflect.Struct {
		//	EncodeTuple(v, buffer)
		//} else {
		v.Encode(buffer)
		//}
	}
}

func (fseq FixedSequence[T]) Bytes() []byte {
	return EncodedBytes(fseq)
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

func SliceU8ToStr(values []U8) Str {
	result := make([]byte, len(values))
	for i, v := range values {
		result[i] = byte(v)
	}
	return Str(result)
}

func StrToSliceU8(s Str) []U8 {
	result := make([]U8, len(s))
	for i, v := range []byte(s) {
		result[i] = U8(v) // TODO: https://github.com/LimeChain/goscale/issues/38
	}
	return result
}

// TODO: need to update Tinygo
// type SequentialU8 interface {
// 	Sequence[U8] | FixedSequence[U8]
// }

// func SequentialU8ToBytes[S SequentialU8](bytes S) []byte {
// 	result := make([]byte, len(bytes))
// 	for i, v := range bytes {
// 		result[i] = byte(v) // TODO: https://github.com/LimeChain/goscale/issues/38
// 	}
// 	return result
// }

func SequenceU8ToBytes(bytes Sequence[U8]) []byte {
	result := make([]byte, len(bytes))
	for i, v := range bytes {
		result[i] = byte(v) // TODO: https://github.com/LimeChain/goscale/issues/38
	}
	return result
}

func FixedSequenceU8ToBytes(bytes FixedSequence[U8]) []byte {
	result := make([]byte, len(bytes))
	for i, v := range bytes {
		result[i] = byte(v) // TODO: https://github.com/LimeChain/goscale/issues/38
	}
	return result
}

func BytesToSequenceU8(bytes []byte) Sequence[U8] {
	result := make(Sequence[U8], len(bytes))
	for i, v := range bytes {
		result[i] = U8(v) // TODO: https://github.com/LimeChain/goscale/issues/38
	}
	return result
}

func BytesToFixedSequenceU8(bytes []byte) FixedSequence[U8] {
	result := make(FixedSequence[U8], len(bytes))
	for i, v := range bytes {
		result[i] = U8(v) // TODO: https://github.com/LimeChain/goscale/issues/38
	}
	return result
}
