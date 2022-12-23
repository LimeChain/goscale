package goscale

import (
	"bytes"
)

type Encodable interface {
	Encode(buffer *bytes.Buffer)
	Bytes() []byte
}

type Ordered interface {
	I8 | I16 | I32 | I64 | U8 | U16 | U32 | U64 | Str
}

func decodeByType(i interface{}, buffer *bytes.Buffer) Encodable {
	switch i.(type) {
	case Bool:
		return DecodeBool(buffer)
	case U8:
		return DecodeU8(buffer)
	case I8:
		return DecodeI8(buffer)
	case U16:
		return DecodeU16(buffer)
	case I16:
		return DecodeI16(buffer)
	case U32:
		return DecodeU32(buffer)
	case I32:
		return DecodeI32(buffer)
	case U64:
		return DecodeU64(buffer)
	case I64:
		return DecodeI64(buffer)
	case U128:
		return DecodeU128(buffer)
	case I128:
		return DecodeI128(buffer)
	case Compact:
		return DecodeCompact(buffer)
	case Sequence[U8]:
		return Sequence[U8](DecodeSliceU8(buffer))
	case Str:
		return DecodeStr(buffer)
	case Empty:
		return DecodeEmpty()
	// TODO:
	// case Result[Encodable]:
	// return DecodeResult(buffer)
	default:
		panic("type not found")
	}
}

func reverseSlice(a []byte) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func ToCompact(v uint64) Compact {
	return Compact(NewU128FromUint64(v))
}
