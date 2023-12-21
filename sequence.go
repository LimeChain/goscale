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

func (seq Sequence[T]) Encode(buffer *bytes.Buffer) error {
	err := ToCompact(len(seq)).Encode(buffer)
	if err != nil {
		return err
	}

	for _, v := range seq {
		//if reflect.TypeOf(v).Kind() == reflect.Struct {
		//	EncodeTuple(v, buffer)
		//} else {
		err := v.Encode(buffer)
		if err != nil {
			return err
		}
		//}
	}

	return nil
}

func (seq Sequence[T]) Bytes() []byte {
	return EncodedBytes(seq)
}

func DecodeSequence[T Encodable](buffer *bytes.Buffer) (Sequence[T], error) {
	size, err := DecodeCompact[BigNumbers](buffer)
	if err != nil {
		return Sequence[T]{}, err
	}
	v := size.ToBigInt()
	values := make([]T, v.Int64())

	for i := 0; i < len(values); i++ {
		t, err := decodeByType(*new(T), buffer)
		if err != nil {
			return Sequence[T]{}, err
		}
		values[i] = t.(T)
	}
	return values, nil
}

func DecodeSequenceWith[T Encodable](buffer *bytes.Buffer, decodeFunc func(buffer *bytes.Buffer) (T, error)) (Sequence[T], error) {
	size, err := DecodeCompact[BigNumbers](buffer)
	if err != nil {
		return Sequence[T]{}, err
	}
	v := size.ToBigInt()
	values := make([]T, v.Int64())

	for i := 0; i < len(values); i++ {
		dec, err := decodeFunc(buffer)
		if err != nil {
			return Sequence[T]{}, err
		}
		values[i] = dec
	}
	return values, nil
}

func DecodeSliceU8(buffer *bytes.Buffer) ([]U8, error) {
	sequence, err := DecodeSequence[U8](buffer)
	if err != nil {
		return make([]U8, 0), err
	}
	return sequence, nil
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

func (fseq FixedSequence[T]) Encode(buffer *bytes.Buffer) error {
	for _, v := range fseq {
		//if reflect.TypeOf(v).Kind() == reflect.Struct {
		//	EncodeTuple(v, buffer)
		//} else {
		err := v.Encode(buffer)
		if err != nil {
			return err
		}
		//}
	}
	return nil
}

func (fseq FixedSequence[T]) Bytes() []byte {
	return EncodedBytes(fseq)
}

func DecodeFixedSequence[T Encodable](size int, buffer *bytes.Buffer) (FixedSequence[T], error) {
	result := make([]T, size)
	for i := 0; i < size; i++ {
		t, err := decodeByType(*new(T), buffer)
		if err != nil {
			return FixedSequence[T]{}, err
		}
		result[i] = t.(T)
	}
	return result, nil
}

// additional helper type
type Str string

func (value Str) Encode(buffer *bytes.Buffer) error {
	_, err := buffer.Write(value.Bytes())
	return err
}

func (value Str) Bytes() []byte {
	return Sequence[U8](StrToSliceU8(value)).Bytes()
}

func DecodeStr(buffer *bytes.Buffer) (Str, error) {
	decodeSlice, err := DecodeSliceU8(buffer)
	if err != nil {
		return "", err
	}
	return SliceU8ToStr(decodeSlice), nil
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
		result[i] = U8(v)
	}
	return result
}

func SequenceU8ToBytes(bytes Sequence[U8]) []byte {
	result := make([]byte, len(bytes))
	for i, v := range bytes {
		result[i] = byte(v)
	}
	return result
}

func FixedSequenceU8ToBytes(bytes FixedSequence[U8]) []byte {
	result := make([]byte, len(bytes))
	for i, v := range bytes {
		result[i] = byte(v)
	}
	return result
}

func BytesToSequenceU8(bytes []byte) Sequence[U8] {
	result := make(Sequence[U8], len(bytes))
	for i, v := range bytes {
		result[i] = U8(v)
	}
	return result
}

func BytesToFixedSequenceU8(bytes []byte) FixedSequence[U8] {
	result := make(FixedSequence[U8], len(bytes))
	for i, v := range bytes {
		result[i] = U8(v)
	}
	return result
}
