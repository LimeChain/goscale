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
)

var (
	errCouldNotDecodeCompact = errors.New("could not decode compact")
	errNotSupported          = errors.New("not supported: n>63 encountered when decoding a compact-encoded uint")
)

type Compact U128

func (c Compact) Encode(buffer *bytes.Buffer) {
	encoder := Encoder{Writer: buffer}
	encoder.Write(c.Bytes())
}

func (c Compact) ToBigInt() *big.Int {
	return U128(c).ToBigInt()
}

func (c Compact) Bytes() []byte {
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

func DecodeCompact(buffer *bytes.Buffer) (Compact, error) {
	decoder := Decoder{Reader: buffer}
	result := make([]byte, 16)
	b, errDecode := decoder.DecodeByte()
	if errDecode != nil {
		return Compact{}, errDecode
	}
	mode := b & 3
	switch mode {
	case 0:
		return Compact(NewU128(big.NewInt(0).SetUint64(uint64(b >> 2)))), nil
	case 1:
		db, err := decoder.DecodeByte()
		if err != nil {
			return Compact{}, err
		}
		r := uint64(db)
		r <<= 6
		r += uint64(b >> 2)
		return Compact(NewU128(big.NewInt(0).SetUint64(r))), nil
	case 2:
		buf := result[:4]
		buf[0] = b
		err := decoder.Read(result[1:4])
		if err != nil {
			return Compact{}, err
		}
		r := binary.LittleEndian.Uint32(buf)
		r >>= 2
		return Compact(NewU128(big.NewInt(0).SetUint64(uint64(r)))), nil
	case 3:
		n := b >> 2
		if n > 63 {
			return Compact(NewU128(0)), errNotSupported
		}
		err := decoder.Read(result[:n+4])
		if err != nil {
			return Compact{}, err
		}
		reverseSlice(result)
		return Compact(NewU128(big.NewInt(0).SetBytes(result))), nil
	default:
		return Compact{}, errCouldNotDecodeCompact
	}
}
