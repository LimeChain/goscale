package goscale

import (
	"encoding/binary"
	"fmt"
)

/*
	https://spec.polkadot.network/#defn-little-endian

	SCALE Fixed Length type translates to Go's fixed-width integer types.
	Values are encoded using a fixed-width, non-negative, little-endian format.
*/

// TODO: handle *big.Int, *scale.Uint128

type U8 uint8

func (value U8) Encode(enc *Encoder) {
	enc.EncodeByte(byte(value))
	// binary.Write(enc.Writer, binary.LittleEndian, value)
}

func (dec *Decoder) DecodeU8() U8 {
	buf := make([]byte, 1)
	dec.Read(buf)
	return U8(buf[0])
}

func (value U8) String() string {
	return fmt.Sprintf("%c", uint8(value))
}

type I8 int8

func (value I8) Encode(enc *Encoder) {
	U8(value).Encode(enc)
}

func (dec *Decoder) DecodeI8() I8 {
	return I8(dec.DecodeByte())
}

func (value I8) String() string {
	return fmt.Sprint(int8(value))
}

type U16 uint16

func (value U16) Encode(enc *Encoder) {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, uint16(value))
	enc.Write(buf)
	// binary.Write(enc.Writer, binary.LittleEndian, value)
}

func (dec *Decoder) DecodeU16() U16 {
	buf := make([]byte, 2)
	dec.Read(buf)
	return U16(binary.LittleEndian.Uint16(buf))
}

func (value U16) String() string {
	return fmt.Sprint(uint16(value))
}

type I16 int16

func (value I16) Encode(enc *Encoder) {
	U16(value).Encode(enc)
}

func (dec *Decoder) DecodeI16() I16 {
	return I16(dec.DecodeU16())
}

func (value I16) String() string {
	return fmt.Sprint(int16(value))
}

type U32 uint32

func (value U32) Encode(enc *Encoder) {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(value))
	enc.Write(buf)
	// binary.Write(enc.Writer, binary.LittleEndian, value)
}

func (dec *Decoder) DecodeU32() U32 {
	buf := make([]byte, 4)
	dec.Read(buf)
	return U32(binary.LittleEndian.Uint32(buf))
}

func (value U32) String() string {
	return fmt.Sprint(uint32(value))
}

type I32 int32

func (value I32) Encode(enc *Encoder) {
	U32(value).Encode(enc)
}

func (dec *Decoder) DecodeI32() I32 {
	return I32(dec.DecodeU32())
}

func (value I32) String() string {
	return fmt.Sprint(int32(value))
}

type U64 uint64

func (value U64) Encode(enc *Encoder) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(value))
	enc.Write(buf)
	// binary.Write(enc.Writer, binary.LittleEndian, value)
}

func (dec *Decoder) DecodeU64() U64 {
	buf := make([]byte, 8)
	dec.Read(buf)
	return U64(binary.LittleEndian.Uint64(buf))
}

func (value U64) String() string {
	return fmt.Sprint(uint64(value))
}

type I64 int64

func (value I64) Encode(enc *Encoder) {
	U64(value).Encode(enc)
}

func (dec *Decoder) DecodeI64() I64 {
	return I64(dec.DecodeU64())
}

func (value I64) String() string {
	return fmt.Sprint(int64(value))
}

// func (enc Encoder) EncodeUint64(value uint64) {
// 	buf := make([]byte, 8)
// 	binary.LittleEndian.PutUint64(buf, value)
// 	enc.Write(buf)
// 	// binary.Write(enc.Writer, binary.LittleEndian, value)
// }

// func (dec Decoder) DecodeUint64() uint64 {
// 	buf := make([]byte, 8)
// 	dec.Read(buf)
// 	return binary.LittleEndian.Uint64(buf)
// }

// func (enc Encoder) EncodeInt64(value int64) {
// 	enc.EncodeUint64(uint64(value))
// }

// func (dec Decoder) DecodeInt64() int64 {
// 	return int64(dec.DecodeUint64())
// }

// func (enc Encoder) EncodeUint32(value uint32) {
// 	buf := make([]byte, 4)
// 	binary.LittleEndian.PutUint32(buf, value)
// 	enc.Write(buf)
// 	// binary.Write(enc.Writer, binary.LittleEndian, value)
// }

// func (dec Decoder) DecodeUint32() uint32 {
// 	buf := make([]byte, 4)
// 	dec.Read(buf)
// 	return binary.LittleEndian.Uint32(buf)
// }

// func (enc Encoder) EncodeInt32(value int32) {
// 	enc.EncodeUint32(uint32(value))
// }

// func (dec Decoder) DecodeInt32() int32 {
// 	return int32(dec.DecodeUint32())
// }

// func (enc Encoder) EncodeUint16(value uint16) {
// 	buf := make([]byte, 2)
// 	binary.LittleEndian.PutUint16(buf, value)
// 	enc.Write(buf)
// 	// binary.Write(enc.Writer, binary.LittleEndian, value)
// }

// func (dec Decoder) DecodeUint16() uint16 {
// 	buf := make([]byte, 2)
// 	dec.Read(buf)
// 	return binary.LittleEndian.Uint16(buf)
// }

// func (enc Encoder) EncodeInt16(value int16) {
// 	enc.EncodeUint16(uint16(value))
// }

// func (dec Decoder) DecodeInt16() int16 {
// 	return int16(dec.DecodeUint16())
// }

// func (enc Encoder) EncodeUint8(value uint8) {
// 	enc.EncodeByte(byte(value))
// 	// binary.Write(enc.Writer, binary.LittleEndian, value)
// }

// func (dec Decoder) DecodeUint8() uint8 {
// 	buf := make([]byte, 1)
// 	dec.Read(buf)
// 	return buf[0]
// }

// func (enc Encoder) EncodeInt8(value int8) {
// 	enc.EncodeUint8(byte(value))
// }

// func (dec Decoder) DecodeInt8() int8 {
// 	return int8(dec.DecodeByte())
// }
