package goscale

import "bytes"

type Encodable interface {
	Encode(buffer *bytes.Buffer) // TODO return an error
	String() string
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
