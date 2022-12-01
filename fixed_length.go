package goscale

/*
	Ref: https://spec.polkadot.network/#defn-little-endian

	SCALE Fixed Length type translates to Go's fixed-width integer types.
	Values are encoded using a fixed-width, non-negative, little-endian format.
*/

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// TODO: handle *big.Int, *scale.Uint128

type U8 uint8

func (value U8) Encode(buffer *bytes.Buffer) {
	encoder := Encoder{Writer: buffer}
	encoder.EncodeByte(byte(value))
	// binary.Write(encoder.Writer, binary.LittleEndian, value)
}

func DecodeU8(buffer *bytes.Buffer) U8 {
	decoder := Decoder{Reader: buffer}
	result := make([]byte, 1)
	decoder.Read(result)
	return U8(result[0])
}

func (value U8) String() string {
	return fmt.Sprintf("%c", uint8(value))
}

type I8 int8

func (value I8) Encode(buffer *bytes.Buffer) {
	U8(value).Encode(buffer)
}

func DecodeI8(buffer *bytes.Buffer) I8 {
	decoder := Decoder{Reader: buffer}
	return I8(decoder.DecodeByte())
}

func (value I8) String() string {
	return fmt.Sprint(int8(value))
}

type U16 uint16

func (value U16) Encode(buffer *bytes.Buffer) {
	encoder := Encoder{Writer: buffer}
	result := make([]byte, 2)
	binary.LittleEndian.PutUint16(result, uint16(value))
	encoder.Write(result)
	// binary.Write(encoder.Writer, binary.LittleEndian, value)
}

func DecodeU16(buffer *bytes.Buffer) U16 {
	decoder := Decoder{Reader: buffer}
	result := make([]byte, 2)
	decoder.Read(result)
	return U16(binary.LittleEndian.Uint16(result))
}

func (value U16) String() string {
	return fmt.Sprint(uint16(value))
}

type I16 int16

func (value I16) Encode(buffer *bytes.Buffer) {
	U16(value).Encode(buffer)
}

func DecodeI16(buffer *bytes.Buffer) I16 {
	return I16(DecodeU16(buffer))
}

func (value I16) String() string {
	return fmt.Sprint(int16(value))
}

type U32 uint32

func (value U32) Encode(buffer *bytes.Buffer) {
	encoder := Encoder{Writer: buffer}
	result := make([]byte, 4)
	binary.LittleEndian.PutUint32(result, uint32(value))
	encoder.Write(result)
	// binary.Write(encoder.Writer, binary.LittleEndian, value)
}

func DecodeU32(buffer *bytes.Buffer) U32 {
	decoder := Decoder{Reader: buffer}
	result := make([]byte, 4)
	decoder.Read(result)
	return U32(binary.LittleEndian.Uint32(result))
}

func (value U32) String() string {
	return fmt.Sprint(uint32(value))
}

type I32 int32

func (value I32) Encode(buffer *bytes.Buffer) {
	U32(value).Encode(buffer)
}

func DecodeI32(buffer *bytes.Buffer) I32 {
	return I32(DecodeU32(buffer))
}

func (value I32) String() string {
	return fmt.Sprint(int32(value))
}

type U64 uint64

func (value U64) Encode(buffer *bytes.Buffer) {
	encoder := Encoder{Writer: buffer}
	result := make([]byte, 8)
	binary.LittleEndian.PutUint64(result, uint64(value))
	encoder.Write(result)
	// binary.Write(encoder.Writer, binary.LittleEndian, value)
}

func DecodeU64(buffer *bytes.Buffer) U64 {
	decoder := Decoder{Reader: buffer}
	result := make([]byte, 8)
	decoder.Read(result)
	return U64(binary.LittleEndian.Uint64(result))
}

func (value U64) String() string {
	return fmt.Sprint(uint64(value))
}

type I64 int64

func (value I64) Encode(buffer *bytes.Buffer) {
	U64(value).Encode(buffer)
}

func DecodeI64(buffer *bytes.Buffer) I64 {
	return I64(DecodeU64(buffer))
}

func (value I64) String() string {
	return fmt.Sprint(int64(value))
}
