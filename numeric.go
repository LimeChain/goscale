package goscale

import (
	"errors"
	"math/big"
	"reflect"
)

// Signed integer constraint, for type safety checks
type SignedPrimitiveInteger interface {
	int | int8 | int16 | int32 | int64 | I8 | I16 | I32 | I64
}

// Unsigned integer constraint, for type safety checks
type UnsignedPrimitiveInteger interface {
	uint | uint8 | uint16 | uint32 | uint64 | U8 | U16 | U32 | U64
}

type Integer128 interface {
	U128 | I128
}

// Signed/Unsigned integer constraint, for type safety checks
type Integer interface {
	SignedPrimitiveInteger | UnsignedPrimitiveInteger | Integer128 | *big.Int
}

// Converts any integer value to 128 bits representation
func anyIntegerTo128Bits[N Integer128](n any) N {
	switch n := n.(type) {
	case int:
		return bigIntToGeneric[N](new(big.Int).SetInt64(int64(n)))
	case uint:
		return bigIntToGeneric[N](new(big.Int).SetUint64(uint64(n)))
	case int8:
		return bigIntToGeneric[N](new(big.Int).SetInt64(int64(n)))
	case uint8:
		return bigIntToGeneric[N](new(big.Int).SetUint64(uint64(n)))
	case int16:
		return bigIntToGeneric[N](new(big.Int).SetInt64(int64(n)))
	case uint16:
		return bigIntToGeneric[N](new(big.Int).SetUint64(uint64(n)))
	case int32:
		return bigIntToGeneric[N](new(big.Int).SetInt64(int64(n)))
	case uint32:
		return bigIntToGeneric[N](new(big.Int).SetUint64(uint64(n)))
	case int64:
		return bigIntToGeneric[N](new(big.Int).SetInt64(n))
	case uint64:
		return bigIntToGeneric[N](new(big.Int).SetUint64(n))
	case I8:
		return bigIntToGeneric[N](new(big.Int).SetInt64(int64(n)))
	case U8:
		return bigIntToGeneric[N](new(big.Int).SetUint64(uint64(n)))
	case I16:
		return bigIntToGeneric[N](new(big.Int).SetInt64(int64(n)))
	case U16:
		return bigIntToGeneric[N](new(big.Int).SetUint64(uint64(n)))
	case I32:
		return bigIntToGeneric[N](new(big.Int).SetInt64(int64(n)))
	case U32:
		return bigIntToGeneric[N](new(big.Int).SetUint64(uint64(n)))
	case I64:
		return bigIntToGeneric[N](new(big.Int).SetInt64(int64(n)))
	case U64:
		return bigIntToGeneric[N](new(big.Int).SetUint64(uint64(n)))
	case U128:
		return bigIntToGeneric[N](n.ToBigInt())
	case I128:
		return bigIntToGeneric[N](n.ToBigInt())
	case *big.Int:
		return bigIntToGeneric[N](n)
	default:
		panic("unknown type in anyIntegerTo128Bits")
	}
}

func stringTo128Bits[N Integer128](s string) (N, error) {
	bn, ok := new(big.Int).SetString(s, 10)
	if !ok {
		return bigIntToGeneric[N](big.NewInt(0)), errors.New("can not convert string to big.Int")
	}
	return bigIntToGeneric[N](bn), nil
}

func bigIntToGeneric[N Integer128](bn *big.Int) N {
	switch reflect.Zero(reflect.TypeOf(*new(N))).Interface().(type) {
	case I128:
		return N(bigIntToI128(bn))
	case U128:
		return N(bigIntToU128(bn))
	default:
		panic("unknown numeric type in bigIntToGeneric")
	}
}
