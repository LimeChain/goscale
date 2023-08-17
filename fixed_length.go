package goscale

/*
	Ref: https://spec.polkadot.network/#defn-little-endian

	SCALE Fixed Length type translates to Go's fixed-width integer types.
	Values are encoded using a fixed-width, non-negative, little-endian format.
*/

import (
	"bytes"
	"encoding/binary"
	"math/big"
	"reflect"
)

type U8 uint8

func (value U8) Encode(buffer *bytes.Buffer) {
	encoder := Encoder{Writer: buffer}
	encoder.Write(value.Bytes())
}

func (value U8) Bytes() []byte {
	return []byte{byte(value)}
}

func DecodeU8(buffer *bytes.Buffer) U8 {
	decoder := Decoder{Reader: buffer}
	result := make([]byte, 1)
	decoder.Read(result)
	return U8(result[0])
}

type I8 int8

func (value I8) Encode(buffer *bytes.Buffer) {
	U8(value).Encode(buffer)
}

func (value I8) Bytes() []byte {
	return U8(value).Bytes()
}

func DecodeI8(buffer *bytes.Buffer) I8 {
	decoder := Decoder{Reader: buffer}
	return I8(decoder.DecodeByte())
}

type U16 uint16

func (value U16) Encode(buffer *bytes.Buffer) {
	encoder := Encoder{Writer: buffer}
	encoder.Write(value.Bytes())
}

func (value U16) Bytes() []byte {
	result := make([]byte, 2)
	binary.LittleEndian.PutUint16(result, uint16(value))

	return result
}

func DecodeU16(buffer *bytes.Buffer) U16 {
	decoder := Decoder{Reader: buffer}
	result := make([]byte, 2)
	decoder.Read(result)
	return U16(binary.LittleEndian.Uint16(result))
}

type I16 int16

func (value I16) Encode(buffer *bytes.Buffer) {
	U16(value).Encode(buffer)
}

func (value I16) Bytes() []byte {
	return U16(value).Bytes()
}

func DecodeI16(buffer *bytes.Buffer) I16 {
	return I16(DecodeU16(buffer))
}

type U32 uint32

func (value U32) Encode(buffer *bytes.Buffer) {
	encoder := Encoder{Writer: buffer}
	encoder.Write(value.Bytes())
}

func (value U32) Bytes() []byte {
	result := make([]byte, 4)
	binary.LittleEndian.PutUint32(result, uint32(value))

	return result
}

func DecodeU32(buffer *bytes.Buffer) U32 {
	decoder := Decoder{Reader: buffer}
	result := make([]byte, 4)
	decoder.Read(result)
	return U32(binary.LittleEndian.Uint32(result))
}

type I32 int32

func (value I32) Encode(buffer *bytes.Buffer) {
	U32(value).Encode(buffer)
}

func (value I32) Bytes() []byte {
	return U32(value).Bytes()
}

func DecodeI32(buffer *bytes.Buffer) I32 {
	return I32(DecodeU32(buffer))
}

type U64 uint64

func (value U64) Encode(buffer *bytes.Buffer) {
	encoder := Encoder{Writer: buffer}
	encoder.Write(value.Bytes())
}

func (value U64) Bytes() []byte {
	result := make([]byte, 8)
	binary.LittleEndian.PutUint64(result, uint64(value))

	return result
}

func DecodeU64(buffer *bytes.Buffer) U64 {
	decoder := Decoder{Reader: buffer}
	result := make([]byte, 8)
	decoder.Read(result)
	return U64(binary.LittleEndian.Uint64(result))
}

type I64 int64

func (value I64) Encode(buffer *bytes.Buffer) {
	U64(value).Encode(buffer)
}

func (value I64) Bytes() []byte {
	return U64(value).Bytes()
}

func DecodeI64(buffer *bytes.Buffer) I64 {
	return I64(DecodeU64(buffer))
}

type U128 [2]U64

func NewU128FromUint64(v uint64) U128 {
	return NewU128FromBigInt(new(big.Int).SetUint64(v))
}

func NewU128FromBigInt(v *big.Int) U128 {
	b := make([]byte, 16)
	v.FillBytes(b)

	reverseSlice(b)

	return U128{
		U64(binary.LittleEndian.Uint64(b[:8])),
		U64(binary.LittleEndian.Uint64(b[8:])),
	}
}

func (u U128) Encode(buffer *bytes.Buffer) {
	u[0].Encode(buffer)
	u[1].Encode(buffer)
}

func (u U128) Bytes() []byte {
	return append(u[0].Bytes(), u[1].Bytes()...)
}

func (u U128) ToBigInt() *big.Int {
	return toBigInt(u)
}

func toBigInt(u U128) *big.Int {
	bytes := make([]byte, 16)
	binary.BigEndian.PutUint64(bytes[:8], uint64(u[1]))
	binary.BigEndian.PutUint64(bytes[8:], uint64(u[0]))

	return big.NewInt(0).SetBytes(bytes)
}

func DecodeU128(buffer *bytes.Buffer) U128 {
	decoder := Decoder{Reader: buffer}
	buf := make([]byte, 16)
	decoder.Read(buf)

	return U128{
		U64(binary.LittleEndian.Uint64(buf[:8])),
		U64(binary.LittleEndian.Uint64(buf[8:])),
	}
}

type I128 [2]U64

func NewI128FromBigInt(v big.Int) I128 {
	b := make([]byte, 16)
	v.FillBytes(b)

	reverseSlice(b)

	var result [2]U64
	result[0] = U64(binary.LittleEndian.Uint64(b[:8]))
	result[1] = U64(binary.LittleEndian.Uint64(b[8:]))

	if v.Sign() < 0 {
		result[0] = ^result[0]
		result[1] = ^result[1]

		result[0]++
		if result[0] == 0 {
			result[1]++
		}
	}

	return I128{
		result[0],
		result[1],
	}
}

func (value I128) Encode(buffer *bytes.Buffer) {
	value[0].Encode(buffer)
	value[1].Encode(buffer)
}

func (value I128) Bytes() []byte {
	return append(value[0].Bytes(), value[1].Bytes()...)
}

func (value I128) ToBigInt() *big.Int {
	isNegative := value[1]&U64(1<<63) != 0
	if isNegative {
		if value[0] == 0 {
			value[1]--
		}
		value[0]--

		value[0] = ^value[0]
		value[1] = ^value[1]
	}

	bytes := make([]byte, 16)
	binary.BigEndian.PutUint64(bytes[:8], uint64(value[1]))
	binary.BigEndian.PutUint64(bytes[8:], uint64(value[0]))

	result := big.NewInt(0).SetBytes(bytes)
	if isNegative {
		result.Neg(result)
	}

	return result
}

func DecodeI128(buffer *bytes.Buffer) I128 {
	return I128{
		DecodeU64(buffer),
		DecodeU64(buffer),
	}
}

func DecodeNumeric[T Numeric](buffer *bytes.Buffer) T {
	var result interface{}

	switch reflect.Zero(reflect.TypeOf(*new(T))).Interface().(type) {
	case U8:
		result = DecodeU8(buffer)
	case I8:
		result = DecodeI8(buffer)
	case U16:
		result = DecodeU16(buffer)
	case I16:
		result = DecodeI16(buffer)
	case U32:
		result = DecodeU32(buffer)
	case I32:
		result = DecodeI32(buffer)
	case U64:
		result = DecodeU64(buffer)
	case I64:
		result = DecodeI64(buffer)
	case U128:
		result = DecodeU128(buffer)
	case I128:
		result = DecodeI128(buffer)
	}

	return result.(T)
}
