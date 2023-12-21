package goscale

import (
	"bytes"
	"errors"
)

var (
	errTypeNotFound = errors.New("type not found")
)

type Encodable interface {
	Encode(buffer *bytes.Buffer) error
	Bytes() []byte
}

type Ordered interface {
	I8 | I16 | I32 | I64 | U8 | U16 | U32 | U64 | Str
}

func EncodedBytes(e Encodable) []byte {
	buffer := &bytes.Buffer{}
	e.Encode(buffer)
	return buffer.Bytes()
}

func EncodeEach(buffer *bytes.Buffer, encodables ...Encodable) error {
	for _, encodable := range encodables {
		err := encodable.Encode(buffer)
		if err != nil {
			return err
		}
	}
	return nil
}

func decodeByType(i interface{}, buffer *bytes.Buffer) (Encodable, error) {
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
	case Compact[BigNumbers]:
		return DecodeCompact[BigNumbers](buffer)
	case Sequence[U8]:
		dec, err := DecodeSliceU8(buffer)
		if err != nil {
			return Empty{}, err
		}
		return Sequence[U8](dec), nil
	case Str:
		return DecodeStr(buffer)
	case Empty:
		return DecodeEmpty()
	// TODO:
	// case Result[Encodable]:
	// return DecodeResult(buffer)
	default:
		return Empty{}, errTypeNotFound
	}
}

func reverseSlice(a []byte) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func ToCompact(v interface{}) Compact[BigNumbers] {
	switch v := v.(type) {
	case int:
		return Compact[BigNumbers]{NewU128(v)} //TODO: Platform dependent ?
	case uint:
		return Compact[BigNumbers]{NewU128(v)}
	case int8:
		return Compact[BigNumbers]{NewI8(v)}
	case I8:
		return Compact[BigNumbers]{NewI8(int8(v))}
	case uint8:
		return Compact[BigNumbers]{NewU8(v)}
	case U8:
		return Compact[BigNumbers]{NewU8(uint8(v))}
	case int16:
		return Compact[BigNumbers]{NewI16(v)}
	case I16:
		return Compact[BigNumbers]{NewI16(int16(v))}
	case uint16:
		return Compact[BigNumbers]{NewU16(v)}
	case U16:
		return Compact[BigNumbers]{NewU16(uint16(v))}
	case int32:
		return Compact[BigNumbers]{NewI32(v)}
	case I32:
		return Compact[BigNumbers]{NewI32(int32(v))}
	case uint32:
		return Compact[BigNumbers]{NewU32(v)}
	case U32:
		return Compact[BigNumbers]{NewU32(uint32(v))}
	case int64:
		return Compact[BigNumbers]{NewI64(v)}
	case I64:
		return Compact[BigNumbers]{NewI64(int64(v))}
	case uint64:
		return Compact[BigNumbers]{NewU64(v)}
	case U64:
		return Compact[BigNumbers]{NewU64(uint64(v))}
	case U128:
		return Compact[BigNumbers]{NewU128(v)}
	case I128:
		return Compact[BigNumbers]{NewI128(v)}
	default:
		panic("invalid numeric type in ToCompact()")
	}
}
