package goscale

/*
	https://spec.polkadot.network/#defn-little-endian

	SCALE Fixed Length type translates to Go's fixed-width integer types.
	Values are encoded using a fixed-width, non-negative, little-endian format.
*/

import (
	"encoding/binary"
)

// TODO: handle *big.Int, *scale.Uint128

func (enc Encoder) EncodeUint64(value uint64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, value)
	enc.Write(buf)
	// binary.Write(enc.Writer, binary.LittleEndian, value)
}

func (dec Decoder) DecodeUint64() uint64 {
	buf := make([]byte, 8)
	dec.Read(buf)
	return binary.LittleEndian.Uint64(buf)
}

func (enc Encoder) EncodeInt64(value int64) {
	enc.EncodeUint64(uint64(value))
}

func (dec Decoder) DecodeInt64() int64 {
	return int64(dec.DecodeUint64())
}

func (enc Encoder) EncodeUint32(value uint32) {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, value)
	enc.Write(buf)
	// binary.Write(enc.Writer, binary.LittleEndian, value)
}

func (dec Decoder) DecodeUint32() uint32 {
	buf := make([]byte, 4)
	dec.Read(buf)
	return binary.LittleEndian.Uint32(buf)
}

func (enc Encoder) EncodeInt32(value int32) {
	enc.EncodeUint32(uint32(value))
}

func (dec Decoder) DecodeInt32() int32 {
	return int32(dec.DecodeUint32())
}

func (enc Encoder) EncodeUint16(value uint16) {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, value)
	enc.Write(buf)
	// binary.Write(enc.Writer, binary.LittleEndian, value)
}

func (dec Decoder) DecodeUint16() uint16 {
	buf := make([]byte, 2)
	dec.Read(buf)
	return binary.LittleEndian.Uint16(buf)
}

func (enc Encoder) EncodeInt16(value int16) {
	enc.EncodeUint16(uint16(value))
}

func (dec Decoder) DecodeInt16() int16 {
	return int16(dec.DecodeUint16())
}

func (enc Encoder) EncodeUint8(value uint8) {
	enc.EncodeByte(byte(value))
	// binary.Write(enc.Writer, binary.LittleEndian, value)
}

func (dec Decoder) DecodeUint8() uint8 {
	buf := make([]byte, 1)
	dec.Read(buf)
	return buf[0]
}

func (enc Encoder) EncodeInt8(value int8) {
	enc.EncodeUint8(byte(value))
}

func (dec Decoder) DecodeInt8() int8 {
	return int8(dec.DecodeByte())
}
