package goscale

import (
	"bytes"
	"fmt"
)

type Result[T Encodable] struct {
	ok    Bool
	value Encodable
}

func (r Result[T]) Encode(buffer *bytes.Buffer) {
	r.ok.Encode(buffer)
	r.value.Encode(buffer)
}

func DecodeResult[T Encodable](dec Encodable, buffer *bytes.Buffer) Result[T] {
	return Result[T]{
		ok:    DecodeBool(buffer),
		value: decodeByType(dec, buffer),
	}
}

func (r Result[T]) String() string {
	return fmt.Sprintf(r.ok.String(), r.value.String())
}
