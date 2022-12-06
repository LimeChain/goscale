package goscale

import (
	"bytes"
)

type Encodable interface {
	Encode(buffer *bytes.Buffer) // TODO return an error
}

type Ordered interface {
	I8 | I16 | I32 | I64 | U8 | U16 | U32 | U64 | Str
}

type Str string

func (value Str) Encode(buffer *bytes.Buffer) {
	seq := Sequence[U8]{Values: StringToSliceU8(string(value))}
	seq.Encode(buffer)
}

func DecodeStr(buffer *bytes.Buffer) Str {
	seq := Sequence[U8]{Values: DecodeSliceU8(buffer)}
	return Str(SliceU8ToString(seq.Values))
}

func SliceU8ToSequenceU8(values []U8) Sequence[U8] {
	return Sequence[U8]{Values: values}
}

func StringToSliceU8(s string) []U8 {
	result := make([]U8, len(s))

	for i, v := range []byte(s) {
		result[i] = U8(v)
		// TODO: fix
		// result = append(result, sc.U8(v)) -> panic: cannot convert pointer to integer -> /tinygo/interp/memory.go:541
	}

	return result
}

func SliceU8ToString(values []U8) string {
	result := make([]byte, len(values))

	for i, v := range values {
		result[i] = byte(v)
	}

	return string(result)
}
