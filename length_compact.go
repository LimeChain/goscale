package goscale

/*
	Ref: https://spec.polkadot.network/#sect-sc-length-and-compact-encoding

	SCALE Length and Compact Encoding translates to Go's integer types of variying sizes.
*/

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
)

/*
	Error: 0x0100: Zero encoded in mode 1)

	0b00 00 00 00 / 00 00 00 00 / 00 00 00 00 / 00 00 00 00`
	xx xx xx 00                                           (0 ... 2**6 - 1)		   (u8)
	yL yL yL 01 / yH yH yH yL                             (2**6 ... 2**14 - 1)	 (u8, u16)  low LH high
	zL zL zL 10 / zM zM zM zL / zM zM zM zM / zH zH zH zM (2**14 ... 2**30 - 1)	 (u16, u32)  low LMMH high
	nn nn nn 11 [ / zz zz zz zz ]{4 + n}                  (2**30 ... 2**536 - 1) (u32, u64, u128, U256, U512, U520) straight LE-encoded
*/

type CompactU8 uint
type CompactU16 uint
type CompactU32 uint
type CompactU64 uint

type Compact uint

func (value Compact) Encode(buffer *bytes.Buffer) {
	encoder := Encoder{Writer: buffer}

	intBuf := make([]byte, 8)
	if value < 1<<30 {
		if value < 1<<6 {
			// 0b00: single-byte mode;
			// upper six bits are the LE encoding of the value (valid only for values of 0-63).
			// (1<<6 - 1 => 63) => (00111111) =>
			// 111111|00
			// binary.Write(encoder.Writer, binary.LittleEndian, uint8(n)<<2)
			encoder.EncodeByte(byte(value) << 2)
		} else if value < 1<<14 {
			// 0b01: two-byte mode:
			// upper six bits and the following byte is the LE encoding of the value (valid only for values 64-(2**14-1)).
			// (1<<14 - 1 => 16383) => (11111111 00111111) << 2 + 1 =>
			// 111111|01 11111111
			// binary.Write(encoder.Writer, binary.LittleEndian, uint16(n<<2)+1)
			buf := intBuf[:2]
			binary.LittleEndian.PutUint16(buf, uint16(value<<2)+1)
			encoder.Write(buf)
		} else {
			// 0b10: four-byte mode:
			// upper six bits and the following three bytes are the LE encoding of the value (valid only for values (2**14)-(2**30-1)).
			// (1<<30 - 1 => 1073741823) => (11111111 11111111 11111111 00111111) << 2 + 2 =>
			// (111111|10 11111111 11111111 11111111)
			// binary.Write(encoder.Writer, binary.LittleEndian, uint32(n<<2)+2)
			buf := intBuf[:4]
			binary.LittleEndian.PutUint32(buf, uint32(value<<2)+2)
			encoder.Write(buf)
		}
		return
	}

	// 0b11: Big-integer mode:
	// The upper six bits are the number of bytes following, plus four. The value is contained, LE encoded, in the bytes following.
	// The final (most significant) byte must be non-zero. Valid only for values (2**30)-(2**536-1).
	// => (001100|11 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000) =>

	n := byte(0)
	limit := uint64(1 << 32)
	// when overflows, limit will be < 256
	for uint64(value) >= limit && limit > 256 {
		n++
		limit <<= 8
	}
	if n > 4 {
		panic("assertion error: n>4 needed to compact-encode uint64")
	}
	encoder.EncodeByte((n << 2) + 3)
	binary.LittleEndian.PutUint64(intBuf[:8], uint64(value))
	encoder.Write(intBuf[:4+n])
}

func DecodeCompact(buffer *bytes.Buffer) Compact {
	decoder := Decoder{Reader: buffer}
	intBuf := make([]byte, 8)
	b := decoder.DecodeByte()
	mode := b & 3
	switch mode {
	case 0:
		return Compact(b >> 2)
	case 1:
		r := uint64(decoder.DecodeByte())
		r <<= 6
		r += uint64(b >> 2)
		return Compact(r)
	case 2:
		buf := intBuf[:4]
		buf[0] = b
		decoder.Read(intBuf[1:4])
		r := binary.LittleEndian.Uint32(buf)
		r >>= 2
		return Compact(r)
	case 3:
		n := b >> 2
		if n > 4 {
			panic("not supported: n>4 encountered when decoding a compact-encoded uint")
		}
		decoder.Read(intBuf[:n+4])
		for i := n + 4; i < 8; i++ {
			intBuf[i] = 0
		}
		return Compact(binary.LittleEndian.Uint64(intBuf[:8]))
	default:
		panic("code should be unreachable")
	}
}

func (value Compact) String() string {
	return fmt.Sprint(uint(value))
}

type CompactU128 struct {
	U128
}

func (c CompactU128) Encode(buffer *bytes.Buffer) {
	encoder := Encoder{Writer: buffer}
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
				encoder.EncodeByte(byte(value) << 2)
			} else if value < 1<<14 {
				// 0b01: two-byte mode:
				// upper six bits and the following byte is the LE encoding of the value (valid only for values 64-(2**14-1)).
				// (1<<14 - 1 => 16383) => (11111111 00111111) << 2 + 1 =>
				// 111111|01 11111111
				// binary.Write(encoder.Writer, binary.LittleEndian, uint16(n<<2)+1)
				buf := make([]byte, 2)
				binary.LittleEndian.PutUint16(buf, uint16(value<<2)+1)
				encoder.Write(buf)
			} else {
				// 0b10: four-byte mode:
				// upper six bits and the following three bytes are the LE encoding of the value (valid only for values (2**14)-(2**30-1)).
				// (1<<30 - 1 => 1073741823) => (11111111 11111111 11111111 00111111) << 2 + 2 =>
				// (111111|10 11111111 11111111 11111111)
				// binary.Write(encoder.Writer, binary.LittleEndian, uint32(n<<2)+2)
				buf := make([]byte, 4)
				binary.LittleEndian.PutUint32(buf, uint32(value<<2)+2)
				encoder.Write(buf)
			}
			return
		}
	}

	// 0b11: Big-integer mode:
	// The upper six bits are the number of bytes following, plus four. The value is contained, LE encoded, in the bytes following.
	// The final (most significant) byte must be non-zero. Valid only for values (2**30)-(2**536-1).
	// => (001100|11 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000) =>
	numBytes := len(bn.Bytes())
	topSixBits := uint8(numBytes - 4)

	encoder.EncodeByte((topSixBits << 2) + 3)
	value := make([]byte, 16)
	binary.LittleEndian.PutUint64(value[:8], uint64(c.U128[0]))
	binary.LittleEndian.PutUint64(value[8:], uint64(c.U128[1]))

	index := 0
	for i, v := range value {
		if v != 0 {
			index = i
			break
		}
	}

	result := value[index:]
	reverseSlice(result)

	encoder.Write(result)
}

func DecodeCompactU128(buffer *bytes.Buffer) CompactU128 {
	decoder := Decoder{Reader: buffer}
	result := make([]byte, 16)
	b := decoder.DecodeByte()
	mode := b & 3
	switch mode {
	case 0:
		return CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(uint64(b >> 2)))}
	case 1:
		r := uint64(decoder.DecodeByte())
		r <<= 6
		r += uint64(b >> 2)
		return CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(r))}
	case 2:
		buf := result[:4]
		buf[0] = b
		decoder.Read(result[1:4])
		r := binary.LittleEndian.Uint32(buf)
		r >>= 2
		return CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(uint64(r)))}
	case 3:
		n := b >> 2
		if n > 63 {
			panic("not supported: n>63 encountered when decoding a compact-encoded uint")
		}
		decoder.Read(result[:n+4])
		reverseSlice(result)
		return CompactU128{NewU128FromBigInt(big.NewInt(0).SetBytes(result))}
	default:
		panic("code should be unreachable")
	}
}
