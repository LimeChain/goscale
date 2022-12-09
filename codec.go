/*
	Simple Concatenated Aggregate Little-Endian” (SCALE) codec

	Polkadot Spec - https://spec.polkadot.network/#sect-scale-codec

	Substrate Ref - https://docs.substrate.io/reference/scale-codec/
*/
package goscale

import (
	"bytes"
	"io"
	"strconv"
)

type Encoder struct {
	Writer io.Writer
}

type Decoder struct {
	Reader io.Reader
}

func (enc Encoder) Write(bytes []byte) {
	n, err := enc.Writer.Write(bytes)
	if err != nil {
		panic(err.Error())
	}
	if n < len(bytes) {
		panic("can not write the provided " + strconv.Itoa(len(bytes)) + " bytes to writer")
	}
}

func (dec Decoder) Read(bytes []byte) {
	n, err := dec.Reader.Read(bytes)
	if err != nil {
		panic(err.Error())
	}
	if n < len(bytes) {
		panic("can not read the required number of bytes " + strconv.Itoa(len(bytes)) + ", only " + strconv.Itoa(n) + " available")
	}
}

func (enc Encoder) EncodeByte(b byte) {
	buf := make([]byte, 1)
	buf[0] = b
	enc.Write(buf[:1])
}

func (dec Decoder) DecodeByte() byte {
	buf := make([]byte, 1)
	dec.Read(buf[:1])
	return buf[0]
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
