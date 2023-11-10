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
type I128 [2]U64

func NewI128[N Integer](n N) I128 {
	return anyIntegerTo128Bits[I128](n)
}

func NewI128FromString(n string) (I128, error) {
	return stringTo128Bits[I128](n)
}

func (n I128) Encode(buffer *bytes.Buffer) error {
	err := n[0].Encode(buffer)
	if err != nil {
		return err
	}
	err = n[1].Encode(buffer)
	if err != nil {
		return err
	}
	return nil
}

func (n I128) Bytes() []byte {
	return append(n[0].Bytes(), n[1].Bytes()...)
}

func DecodeI128(buffer *bytes.Buffer) (I128, error) {
	decU64One, err := DecodeU64(buffer)
	if err != nil {
		return I128{}, err
	}
	decU64Two, err := DecodeU64(buffer)
	if err != nil {
		return I128{}, err
	}
	return I128{
		decU64One,
		decU64Two,
	}, nil
}

func (n I128) ToBigInt() *big.Int {
	isNegative := n.isNegative()

	if isNegative {
		n = negateI128(n)
	}

	bytes := make([]byte, 16)
	binary.BigEndian.PutUint64(bytes[:8], uint64(n[1]))
	binary.BigEndian.PutUint64(bytes[8:], uint64(n[0]))
	result := big.NewInt(0).SetBytes(bytes)

	if isNegative {
		result.Neg(result)
	}

	return result
}

func (n I128) Add(other I128) I128 {
	sumLow, carry := bits.Add64(uint64(n[0]), uint64(other[0]), 0)
	sumHigh, _ := bits.Add64(uint64(n[1]), uint64(other[1]), carry)
	return I128{U64(sumLow), U64(sumHigh)}
}

func (n I128) Sub(other I128) I128 {
	diffLow, borrow := bits.Sub64(uint64(n[0]), uint64(other[0]), 0)
	diffHigh, _ := bits.Sub64(uint64(n[1]), uint64(other[1]), borrow)
	return I128{U64(diffLow), U64(diffHigh)}
}

func (n I128) Mul(other I128) I128 {
	negA := n[1]>>(64-1) == 1
	negB := other[1]>>(64-1) == 1

	absA := n
	if negA {
		absA = negateI128(n)
	}

	absB := other
	if negB {
		absB = negateI128(other)
	}

	high, low := bits.Mul64(uint64(absA[0]), uint64(absB[0]))
	high += uint64(absA[1])*uint64(absB[0]) + uint64(absA[0])*uint64(absB[1])

	result := I128{U64(low), U64(high)}

	// if one of the operands is negative the result is also negative
	if negA != negB {
		return negateI128(result)
	}
	return result
}

func (n I128) Div(other I128) I128 {
	return bigIntToI128(
		new(big.Int).Div(n.ToBigInt(), other.ToBigInt()),
	)
}

func (n I128) Mod(other I128) I128 {
	return bigIntToI128(
		new(big.Int).Mod(n.ToBigInt(), other.ToBigInt()),
	)
}

func (n I128) Eq(other I128) bool {
	return n.ToBigInt().Cmp(other.ToBigInt()) == 0
}

func (n I128) Ne(other I128) bool {
	return !n.Eq(other)
}

func (n I128) Lt(other I128) bool {
	return n.ToBigInt().Cmp(other.ToBigInt()) < 0
}

func (n I128) Lte(other I128) bool {
	return n.ToBigInt().Cmp(other.ToBigInt()) <= 0
}

func (n I128) Gt(other I128) bool {
	return n.ToBigInt().Cmp(other.ToBigInt()) > 0
}

func (n I128) Gte(other I128) bool {
	return n.ToBigInt().Cmp(other.ToBigInt()) >= 0
}

func bigIntToI128(n *big.Int) I128 {
	bytes := make([]byte, 16)
	n.FillBytes(bytes)
	reverseSlice(bytes)

	var result = I128{
		U64(binary.LittleEndian.Uint64(bytes[:8])),
		U64(binary.LittleEndian.Uint64(bytes[8:])),
	}

	if n.Sign() < 0 {
		result = negateI128(result)
	}

	return result
}

func (n I128) isNegative() bool {
	return n[1]&U64(1<<63) != 0
}

func negateI128(n I128) I128 {
	// two's complement representation
	negLow, carry := bits.Add64(^uint64(n[0]), 1, 0)
	negHigh, _ := bits.Add64(^uint64(n[1]), 0, carry)
	return I128{U64(negLow), U64(negHigh)}
}

//func (n I128) SaturatingAdd(b Numeric) Numeric {
//	sumLow, carry := bits.Add64(uint64(a[0]), uint64(b.(I128)[0]), 0)
//	sumHigh, _ := bits.Add64(uint64(a[1]), uint64(b.(I128)[1]), carry)
//	// check for overflow
//	if a[1]&(1<<63) == 0 && b.(I128)[1]&(1<<63) == 0 && sumHigh&(1<<63) != 0 {
//		return MaxI128()
//	}
//	// check for underflow
//	if a[1]&(1<<63) != 0 && b.(I128)[1]&(1<<63) != 0 && sumHigh&(1<<63) == 0 {
//		return MinI128()
//	}
//	return I128{U64(sumLow), U64(sumHigh)}
//}
//
//func (n I128) SaturatingSub(b Numeric) Numeric {
//	diffLow, borrow := bits.Sub64(uint64(a[0]), uint64(b.(I128)[0]), 0)
//	diffHigh, _ := bits.Sub64(uint64(a[1]), uint64(b.(I128)[1]), borrow)
//	// check for overflow
//	if a[1]&(1<<63) == 0 && b.(I128)[1]&(1<<63) != 0 && diffHigh&(1<<63) != 0 {
//		return MaxI128()
//	}
//	// check for underflow
//	if a[1]&(1<<63) != 0 && b.(I128)[1]&(1<<63) == 0 && diffHigh&(1<<63) == 0 {
//		return MinI128()
//	}
//	return I128{U64(diffLow), U64(diffHigh)}
//}
//
//func (n I128) SaturatingMul(b Numeric) Numeric {
//	result := new(big.Int).Mul(a.ToBigInt(), b.(I128).ToBigInt())
//	// define the maximum and minimum representable I128 values
//	maxI128 := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 127), big.NewInt(1))
//	minI128 := new(big.Int).Neg(new(big.Int).Lsh(big.NewInt(1), 127))
//	if result.Cmp(maxI128) > 0 {
//		return bigIntToI128(maxI128)
//	} else if result.Cmp(minI128) < 0 {
//		return bigIntToI128(minI128)
//	}
//	return bigIntToI128(result)
//}
