package goscale

import "bytes"

type Encodable interface {
	Encode(buffer *bytes.Buffer) // TODO return an error
	String() string
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
