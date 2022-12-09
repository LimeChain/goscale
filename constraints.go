package goscale

import (
	"bytes"
)

type Encodable interface {
	Encode(buffer *bytes.Buffer)
}

type Ordered interface {
	I8 | I16 | I32 | I64 | U8 | U16 | U32 | U64 | Str
}
