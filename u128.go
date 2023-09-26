package goscale

import (
	"bytes"
	"encoding/binary"
	"math/big"
	"math/bits"
)

// little endian byte order
// [0] least significant bits
// [1] most significant bits
type U128 [2]U64

func NewU128[N Integer](n N) U128 {
	return anyIntegerTo128Bits[U128](n)
}

func NewU128FromString(n string) (U128, error) {
	return stringTo128Bits[U128](n)
}

func (n U128) Encode(buffer *bytes.Buffer) {
	n[0].Encode(buffer)
	n[1].Encode(buffer)
}

func (n U128) Bytes() []byte {
	return append(n[0].Bytes(), n[1].Bytes()...)
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

func (n U128) ToBigInt() *big.Int {
	bytes := make([]byte, 16)
	binary.BigEndian.PutUint64(bytes[:8], uint64(n[1]))
	binary.BigEndian.PutUint64(bytes[8:], uint64(n[0]))
	return big.NewInt(0).SetBytes(bytes)
}

func (n U128) Add(other U128) U128 {
	sumLow, carry := bits.Add64(uint64(n[0]), uint64(other[0]), 0)
	sumHigh, _ := bits.Add64(uint64(n[1]), uint64(other[1]), carry)
	return U128{U64(sumLow), U64(sumHigh)}
}

func (n U128) Sub(other U128) U128 {
	diffLow, borrow := bits.Sub64(uint64(n[0]), uint64(other[0]), 0)
	diffHigh, _ := bits.Sub64(uint64(n[1]), uint64(other[1]), borrow)
	return U128{U64(diffLow), U64(diffHigh)}
}

func (n U128) Mul(other U128) U128 {
	high, low := bits.Mul64(uint64(n[0]), uint64(other[0]))
	high += uint64(n[1])*uint64(other[0]) + uint64(n[0])*uint64(other[1])
	return U128{U64(low), U64(high)}
}

func (n U128) Div(other U128) U128 {
	return bigIntToU128(
		new(big.Int).Div(n.ToBigInt(), other.ToBigInt()),
	)
}

func (n U128) Mod(other U128) U128 {
	return bigIntToU128(
		new(big.Int).Mod(n.ToBigInt(), other.ToBigInt()),
	)
}

func (n U128) Eq(other U128) bool {
	return n.ToBigInt().Cmp(other.ToBigInt()) == 0
}

func (n U128) Ne(other U128) bool {
	return !n.Eq(other)
}

func (n U128) Lt(other U128) bool {
	return n.ToBigInt().Cmp(other.ToBigInt()) < 0
}

func (n U128) Lte(other U128) bool {
	return n.ToBigInt().Cmp(other.ToBigInt()) <= 0
}

func (n U128) Gt(other U128) bool {
	return n.ToBigInt().Cmp(other.ToBigInt()) > 0
}

func (n U128) Gte(other U128) bool {
	return n.ToBigInt().Cmp(other.ToBigInt()) >= 0
}

//func (n U128) SaturatingAdd(other U128) U128 {
//	sumLow, carry := bits.Add64(uint64(n[0]), uint64(other[0]), 0)
//	sumHigh, overflow := bits.Add64(uint64(n[1]), uint64(other[1]), carry)
//	// check for overflow
//	if overflow == 1 || (carry == 1 && sumHigh <= uint64(n[1]) && sumHigh <= uint64(other[1])) {
//		return MaxU128()
//	}
//	return U128{U64(sumLow), U64(sumHigh)}
//}
//
//func (n U128) SaturatingSub(other U128) U128 {
//	low, borrow := bits.Sub64(uint64(n[0]), uint64(other[0]), 0)
//	high, _ := bits.Sub64(uint64(n[1]), uint64(other[1]), borrow)
//	// check for underflow
//	if borrow == 1 || high > uint64(n[1]) {
//		return U128{0, 0}
//	}
//	return U128{U64(low), U64(high)}
//}
//
//func (n U128) SaturatingMul(other U128) U128 {
//	result := new(big.Int).Mul(n.ToBigInt(), other.ToBigInt())
//	// check for overflow
//	maxValue := new(big.Int)
//	maxValue.Lsh(big.NewInt(1), 128)
//	maxValue.Sub(maxValue, big.NewInt(1))
//	if result.Cmp(maxValue) > 0 {
//		result.Set(maxValue)
//	}
//	return bigIntToU128(result)
//}

func bigIntToU128(n *big.Int) U128 {
	bytes := make([]byte, 16)
	n.FillBytes(bytes)
	reverseSlice(bytes)

	return U128{
		U64(binary.LittleEndian.Uint64(bytes[:8])),
		U64(binary.LittleEndian.Uint64(bytes[8:])),
	}
}
