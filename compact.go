package goscale

/*
	Ref: https://spec.polkadot.network/#sect-sc-length-and-compact-encoding

	SCALE Length and Compact Encoding translates to Go's integer types of variying sizes.
*/

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math/big"
	"reflect"
)

var (
	errCouldNotDecodeCompact = errors.New("could not decode compact")
	errNotSupported          = errors.New("not supported: n>63 encountered when decoding a compact-encoded uint")
)

type Numeric interface {
	ToBigInt() *big.Int
}

type Compact[T Numeric] struct {
	Number T
}

func (c Compact[T]) Encode(buffer *bytes.Buffer) error {
	encoder := Encoder{Writer: buffer}
	return encoder.Write(c.Bytes())
}

func (c Compact[T]) ToBigInt() *big.Int {
	return c.Number.ToBigInt()
}

func (c Compact[T]) Bytes() []byte {
	bn := c.ToBigInt()

	if bn.IsUint64() {
		value := bn.Uint64()
		if value < 1<<30 {
			if value < 1<<6 {
				// 0b00: single-byte mode;
				// upper six bits are the LE encoding of the value (valid only for values of 0-63).
				// (1<<6 - 1 => 63) => (00111111) =>
				// 111111|00
				// binary.Write(encoder.Writer, binary.LittleEndian, uint8(n)<<2)
				return []byte{byte(value) << 2}
			} else if value < 1<<14 {
				// 0b01: two-byte mode:
				// upper six bits and the following byte is the LE encoding of the value (valid only for values 64-(2**14-1)).
				// (1<<14 - 1 => 16383) => (11111111 00111111) << 2 + 1 =>
				// 111111|01 11111111
				// binary.Write(encoder.Writer, binary.LittleEndian, uint16(n<<2)+1)
				buf := make([]byte, 2)
				binary.LittleEndian.PutUint16(buf, uint16(value<<2)+1)
				return buf
			} else {
				// 0b10: four-byte mode:
				// upper six bits and the following three bytes are the LE encoding of the value (valid only for values (2**14)-(2**30-1)).
				// (1<<30 - 1 => 1073741823) => (11111111 11111111 11111111 00111111) << 2 + 2 =>
				// (111111|10 11111111 11111111 11111111)
				// binary.Write(encoder.Writer, binary.LittleEndian, uint32(n<<2)+2)
				buf := make([]byte, 4)
				binary.LittleEndian.PutUint32(buf, uint32(value<<2)+2)
				return buf
			}
		}
	}

	// 0b11: Big-integer mode:
	// The upper six bits are the number of bytes following, plus four. The value is contained, LE encoded, in the bytes following.
	// The final (most significant) byte must be non-zero. Valid only for values (2**30)-(2**536-1).
	// => (001100|11 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000) =>
	b := bn.Bytes()
	numBytes := len(b)
	topSixBits := uint8(numBytes - 4)

	reverseSlice(b)

	return append([]byte{(topSixBits << 2) + 3}, b...)
}

func DecodeCompact[T Numeric](buffer *bytes.Buffer) (Compact[T], error) {
	decoder := Decoder{Reader: buffer}
	result := make([]byte, 16)
	b, err := decoder.DecodeByte()
	if err != nil {
		return Compact[T]{}, err
	}
	mode := b & 3
	var value Numeric
	switch mode {
	case 0:
		switch reflect.TypeOf(*new(T)) {
		case reflect.TypeOf(*new(U128)):
			value = NewU128(big.NewInt(0).SetUint64(uint64(b >> 2)))
		case reflect.TypeOf(*new(U64)):
			value = NewU64(uint64(b >> 2))
		case reflect.TypeOf(*new(U32)):
			value = NewU32(uint32(b >> 2))
		case reflect.TypeOf(*new(U16)):
			value = NewU16(uint16(b >> 2))
		case reflect.TypeOf(*new(U8)):
			value = NewU8(b >> 2)
		default:
			value = NewU128(big.NewInt(0).SetUint64(uint64(b >> 2)))
		}
		return Compact[T]{Number: interface{}(value).(T)}, nil
	case 1:
		db, err := decoder.DecodeByte()
		if err != nil {
			return Compact[T]{}, err
		}
		r := uint64(db)
		r <<= 6
		r += uint64(b >> 2)
		switch reflect.TypeOf(*new(T)) {
		case reflect.TypeOf(*new(U128)):
			value = NewU128(r)
		case reflect.TypeOf(*new(U64)):
			value = NewU64(r)
		case reflect.TypeOf(*new(U32)):
			value = NewU32(uint32(r))
		case reflect.TypeOf(*new(U8)):
			value = NewU8(uint8(r))
		default:
			value = NewU128(r)
		}

		return Compact[T]{interface{}(value).(T)}, nil
	case 2:
		buf := result[:4]
		buf[0] = b
		err := decoder.Read(result[1:4])
		if err != nil {
			return Compact[T]{}, err
		}
		r := binary.LittleEndian.Uint32(buf)
		r >>= 2
		switch reflect.TypeOf(*new(T)) {
		case reflect.TypeOf(*new(U128)):
			value = NewU128(uint64(r))
		case reflect.TypeOf(*new(U64)):
			value = NewU64(uint64(r))
		case reflect.TypeOf(*new(U32)):
			value = NewU32(r)
		case reflect.TypeOf(*new(U8)):
			value = NewU8(uint8(r))
		default:
			value = NewU128(r)
		}
		return Compact[T]{interface{}(value).(T)}, nil
	case 3:
		n := b >> 2
		if n > 63 {
			value := NewU64(0)
			result := interface{}(value).(T)
			return Compact[T]{result}, errNotSupported
		}
		err := decoder.Read(result[:n+4])
		if err != nil {
			return Compact[T]{}, err
		}
		reverseSlice(result)
		switch reflect.TypeOf(*new(T)) {
		case reflect.TypeOf(*new(U128)):
			value = NewU128(big.NewInt(0).SetBytes(result))
		case reflect.TypeOf(*new(U64)):
			value = NewU64(big.NewInt(0).SetBytes(result).Uint64())
		case reflect.TypeOf(*new(U32)):
			value = NewU32(uint32(big.NewInt(0).SetBytes(result).Uint64()))
		default:
			value = NewU128(big.NewInt(0).SetBytes(result))
		}
		result := interface{}(value).(T)
		return Compact[T]{result}, nil
	default:
		return Compact[T]{}, errCouldNotDecodeCompact
	}
}
