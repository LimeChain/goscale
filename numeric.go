package goscale

// This could be extracted into a separate package. However,
// to define encoding/decoding methods, the type must be within
// the same package. This would require the definition of new
// types and their corresponding encoding/decoding methods,
// leading to frequent type conversions, which is not ideal.
// e.g. goscale.U8(numeric.U8(1)).

import (
	"errors"
	"math/big"
	"reflect"
)

var ErrOverflow = errors.New("overflow")

const MinInt64 = 1 << 63
const MaxInt64 = ^uint64(0) >> 1
const MaxUint64 = ^uint64(0)

// Go's primitive numeric types don't have a common interface.
// To be able to write generic code that works with any numeric
// type, including some custom types (U128/I128), we need to
// define a common interface.

type Numeric interface {
	Encodable

	Interface() Numeric

	Add(other Numeric) Numeric
	Sub(other Numeric) Numeric
	Mul(other Numeric) Numeric
	Div(other Numeric) Numeric
	Mod(other Numeric) Numeric

	SaturatingAdd(other Numeric) Numeric
	SaturatingSub(other Numeric) Numeric
	SaturatingMul(other Numeric) Numeric

	// TODO: finish implementing these
	// SaturatingDiv(other Numeric) Numeric
	// SaturatingMod(other Numeric) Numeric

	// CheckedAdd(other Numeric) (Numeric, error)
	// CheckedSub(other Numeric) (Numeric, error)
	// CheckedMul(other Numeric) (Numeric, error)
	// CheckedDiv(other Numeric) (Numeric, error)
	// CheckedMod(other Numeric) (Numeric, error)

	// And(other Numeric) Numeric
	// Or(other Numeric) Numeric
	// Xor(other Numeric) Numeric
	// Not() Numeric
	// Lsh(other Numeric) Numeric
	// Rsh(other Numeric) Numeric

	TrailingZeros() Numeric

	Eq(other Numeric) bool
	Ne(other Numeric) bool
	Lt(other Numeric) bool
	Lte(other Numeric) bool
	Gt(other Numeric) bool
	Gte(other Numeric) bool

	Max(other Numeric) Numeric
	Min(other Numeric) Numeric
	Clamp(minValue, maxValue Numeric) Numeric
}

// TODO: check when converting from bigger to smaller types (e.g. U64 -> U8)

// returns any Numeric type as a generic type N
func NewNumeric[N Numeric](n any) N {
	switch reflect.Zero(reflect.TypeOf(*new(N))).Interface().(type) {
	case U8:
		switch n := n.(type) {
		case uint:
			return U8(n).Interface().(N)
		case int:
			return U8(n).Interface().(N)
		case uint8:
			return U8(n).Interface().(N)
		case int8:
			return U8(n).Interface().(N)
		case uint16:
			return U8(n).Interface().(N)
		case int16:
			return U8(n).Interface().(N)
		case uint32:
			return U8(n).Interface().(N)
		case int32:
			return U8(n).Interface().(N)
		case uint64:
			return U8(n).Interface().(N)
		case int64:
			return U8(n).Interface().(N)
		case U8:
			return U8(n).Interface().(N)
		case I8:
			return U8(n).Interface().(N)
		case U16:
			return U8(n).Interface().(N)
		case I16:
			return U8(n).Interface().(N)
		case U32:
			return U8(n).Interface().(N)
		case I32:
			return U8(n).Interface().(N)
		case U64:
			return U8(n).Interface().(N)
		case I64:
			return U8(n).Interface().(N)
		case U128:
			return To[U8](n).Interface().(N)
		case I128:
			return To[U8](n).Interface().(N)
		default:
			panic("unknown type for U8")
		}

	case I8:
		switch n := n.(type) {
		case uint:
			return I8(n).Interface().(N)
		case int:
			return I8(n).Interface().(N)
		case uint8:
			return I8(n).Interface().(N)
		case int8:
			return I8(n).Interface().(N)
		case uint16:
			return I8(n).Interface().(N)
		case int16:
			return I8(n).Interface().(N)
		case uint32:
			return I8(n).Interface().(N)
		case int32:
			return I8(n).Interface().(N)
		case uint64:
			return I8(n).Interface().(N)
		case int64:
			return I8(n).Interface().(N)
		case U8:
			return I8(n).Interface().(N)
		case I8:
			return I8(n).Interface().(N)
		case U16:
			return I8(n).Interface().(N)
		case I16:
			return I8(n).Interface().(N)
		case U32:
			return I8(n).Interface().(N)
		case I32:
			return I8(n).Interface().(N)
		case U64:
			return I8(n).Interface().(N)
		case I64:
			return I8(n).Interface().(N)
		case U128:
			return To[I8](n).Interface().(N)
		case I128:
			return To[I8](n).Interface().(N)
		default:
			panic("unknown type for I8")
		}

	case U16:
		switch n := n.(type) {
		case uint:
			return U16(n).Interface().(N)
		case int:
			return U16(n).Interface().(N)
		case uint8:
			return U16(n).Interface().(N)
		case int8:
			return U16(n).Interface().(N)
		case uint16:
			return U16(n).Interface().(N)
		case int16:
			return U16(n).Interface().(N)
		case uint32:
			return U16(n).Interface().(N)
		case int32:
			return U16(n).Interface().(N)
		case uint64:
			return U16(n).Interface().(N)
		case int64:
			return U16(n).Interface().(N)
		case U8:
			return U16(n).Interface().(N)
		case I8:
			return U16(n).Interface().(N)
		case U16:
			return n.Interface().(N)
		case I16:
			return U16(n).Interface().(N)
		case U32:
			return U16(n).Interface().(N)
		case I32:
			return U16(n).Interface().(N)
		case U64:
			return U16(n).Interface().(N)
		case I64:
			return U16(n).Interface().(N)
		case U128:
			return To[U16](n).Interface().(N)
		case I128:
			return To[U16](n).Interface().(N)
		default:
			panic("unknown type for U16")
		}

	case I16:
		switch n := n.(type) {
		case uint:
			return I16(n).Interface().(N)
		case int:
			return I16(n).Interface().(N)
		case uint8:
			return I16(n).Interface().(N)
		case int8:
			return I16(n).Interface().(N)
		case uint16:
			return I16(n).Interface().(N)
		case int16:
			return I16(n).Interface().(N)
		case uint32:
			return I16(n).Interface().(N)
		case int32:
			return I16(n).Interface().(N)
		case uint64:
			return I16(n).Interface().(N)
		case int64:
			return I16(n).Interface().(N)
		case U8:
			return I16(n).Interface().(N)
		case I8:
			return I16(n).Interface().(N)
		case U16:
			return I16(n).Interface().(N)
		case I16:
			return n.Interface().(N)
		case U32:
			return I16(n).Interface().(N)
		case I32:
			return I16(n).Interface().(N)
		case U64:
			return I16(n).Interface().(N)
		case I64:
			return I16(n).Interface().(N)
		case U128:
			return To[I16](n).Interface().(N)
		case I128:
			return To[I16](n).Interface().(N)
		default:
			panic("unknown type for I16")
		}

	case U32:
		switch n := n.(type) {
		case uint:
			return U32(n).Interface().(N)
		case int:
			return U32(n).Interface().(N)
		case uint8:
			return U32(n).Interface().(N)
		case int8:
			return U32(n).Interface().(N)
		case uint16:
			return U32(n).Interface().(N)
		case int16:
			return U32(n).Interface().(N)
		case uint32:
			return U32(n).Interface().(N)
		case int32:
			return U32(n).Interface().(N)
		case uint64:
			return U32(n).Interface().(N)
		case int64:
			return U32(n).Interface().(N)
		case U8:
			return U32(n).Interface().(N)
		case I8:
			return U32(n).Interface().(N)
		case U16:
			return U32(n).Interface().(N)
		case I16:
			return U32(n).Interface().(N)
		case U32:
			return U32(n).Interface().(N)
		case I32:
			return U32(n).Interface().(N)
		case U64:
			return U32(n).Interface().(N)
		case I64:
			return U32(n).Interface().(N)
		case U128:
			return To[U32](n).Interface().(N)
		case I128:
			return To[U32](n).Interface().(N)
		default:
			panic("unknown type for U32")
		}

	case I32:
		switch n := n.(type) {
		case uint:
			return I32(n).Interface().(N)
		case int:
			return I32(n).Interface().(N)
		case uint8:
			return I32(n).Interface().(N)
		case int8:
			return I32(n).Interface().(N)
		case uint16:
			return I32(n).Interface().(N)
		case int16:
			return I32(n).Interface().(N)
		case uint32:
			return I32(n).Interface().(N)
		case int32:
			return I32(n).Interface().(N)
		case uint64:
			return I32(n).Interface().(N)
		case int64:
			return I32(n).Interface().(N)
		case U8:
			return I32(n).Interface().(N)
		case I8:
			return I32(n).Interface().(N)
		case U16:
			return I32(n).Interface().(N)
		case I16:
			return I32(n).Interface().(N)
		case U32:
			return I32(n).Interface().(N)
		case I32:
			return I32(n).Interface().(N)
		case U64:
			return I32(n).Interface().(N)
		case I64:
			return I32(n).Interface().(N)
		case U128:
			return To[I32](n).Interface().(N)
		case I128:
			return To[I32](n).Interface().(N)
		default:
			panic("unknown type for I32")
		}

	case U64:
		switch n := n.(type) {
		case uint:
			return U64(n).Interface().(N)
		case int:
			return U64(n).Interface().(N)
		case uint8:
			return U64(n).Interface().(N)
		case int8:
			return U64(n).Interface().(N)
		case uint16:
			return U64(n).Interface().(N)
		case int16:
			return U64(n).Interface().(N)
		case uint32:
			return U64(n).Interface().(N)
		case int32:
			return U64(n).Interface().(N)
		case uint64:
			return U64(n).Interface().(N)
		case int64:
			return U64(n).Interface().(N)
		case U8:
			return U64(n).Interface().(N)
		case I8:
			return U64(n).Interface().(N)
		case U16:
			return U64(n).Interface().(N)
		case I16:
			return U64(n).Interface().(N)
		case U32:
			return U64(n).Interface().(N)
		case I32:
			return U64(n).Interface().(N)
		case U64:
			return U64(n).Interface().(N)
		case I64:
			return U64(n).Interface().(N)
		case U128:
			return To[U64](n).Interface().(N)
		case I128:
			return To[U64](n).Interface().(N)
		default:
			panic("unknown type for U64")
		}

	case I64:
		switch n := n.(type) {
		case uint:
			return I64(n).Interface().(N)
		case int:
			return I64(n).Interface().(N)
		case uint8:
			return I64(n).Interface().(N)
		case int8:
			return I64(n).Interface().(N)
		case uint16:
			return I64(n).Interface().(N)
		case int16:
			return I64(n).Interface().(N)
		case uint32:
			return I64(n).Interface().(N)
		case int32:
			return I64(n).Interface().(N)
		case uint64:
			return I64(n).Interface().(N)
		case int64:
			return I64(n).Interface().(N)
		case U8:
			return I64(n).Interface().(N)
		case I8:
			return I64(n).Interface().(N)
		case U16:
			return I64(n).Interface().(N)
		case I16:
			return I64(n).Interface().(N)
		case U32:
			return I64(n).Interface().(N)
		case I32:
			return I64(n).Interface().(N)
		case U64:
			return I64(n).Interface().(N)
		case I64:
			return I64(n).Interface().(N)
		case U128:
			return To[I64](n).Interface().(N)
		case I128:
			return To[I64](n).Interface().(N)
		default:
			panic("unknown primitive type for I64")
		}

	case U128:
		switch n := n.(type) {
		case uint:
			return NewU128(uint64(n)).Interface().(N)
		case int:
			return NewU128(uint64(n)).Interface().(N)
		case uint8:
			return NewU128(uint64(n)).Interface().(N)
		case int8:
			return NewU128(uint64(n)).Interface().(N)
		case uint16:
			return NewU128(uint64(n)).Interface().(N)
		case int16:
			return NewU128(uint64(n)).Interface().(N)
		case uint32:
			return NewU128(uint64(n)).Interface().(N)
		case int32:
			return NewU128(uint64(n)).Interface().(N)
		case uint64:
			return NewU128(uint64(n)).Interface().(N)
		case int64:
			return NewU128(uint64(n)).Interface().(N)
		case U8:
			return NewU128(uint64(n)).Interface().(N)
		case I8:
			return NewU128(uint64(n)).Interface().(N)
		case U16:
			return NewU128(uint64(n)).Interface().(N)
		case I16:
			return NewU128(uint64(n)).Interface().(N)
		case U32:
			return NewU128(uint64(n)).Interface().(N)
		case I32:
			return NewU128(uint64(n)).Interface().(N)
		case U64:
			return NewU128(uint64(n)).Interface().(N)
		case I64:
			return NewU128(uint64(n)).Interface().(N)
		case U128:
			return n.Interface().(N)
		case I128:
			return U128(n).Interface().(N)
		default:
			panic("unknown primitive type for U128")
		}

	case I128:
		switch n := n.(type) {
		case uint:
			return NewI128(int64(n)).Interface().(N)
		case int:
			return NewI128(int64(n)).Interface().(N)
		case uint8:
			return NewI128(int64(n)).Interface().(N)
		case int8:
			return NewI128(int64(n)).Interface().(N)
		case uint16:
			return NewI128(int64(n)).Interface().(N)
		case int16:
			return NewI128(int64(n)).Interface().(N)
		case uint32:
			return NewI128(int64(n)).Interface().(N)
		case int32:
			return NewI128(int64(n)).Interface().(N)
		case uint64:
			return NewI128(int64(n)).Interface().(N)
		case int64:
			return NewI128(int64(n)).Interface().(N)
		case U8:
			return NewI128(int64(n)).Interface().(N)
		case I8:
			return NewI128(int64(n)).Interface().(N)
		case U16:
			return NewI128(int64(n)).Interface().(N)
		case I16:
			return NewI128(int64(n)).Interface().(N)
		case U32:
			return NewI128(int64(n)).Interface().(N)
		case I32:
			return NewI128(int64(n)).Interface().(N)
		case U64:
			return NewI128(int64(n)).Interface().(N)
		case I64:
			return NewI128(int64(n)).Interface().(N)
		case U128:
			return To[I128](n).Interface().(N)
		case I128:
			return n.Interface().(N)
		default:
			panic("unknown primitive type for I128")
		}

	default:
		panic("unknown numeric type N in NewNumeric[N]")
	}
}

// converts from any Numeric type to any other Numeric type
func To[N Numeric](n Numeric) N {
	switch reflect.Zero(reflect.TypeOf(*new(N))).Interface().(type) {
	case U8:
		switch n := n.(type) {
		case U8:
			return U8(n).Interface().(N)
		case I8:
			return U8(n).Interface().(N)
		case U16:
			return U8(n).Interface().(N)
		case I16:
			return U8(n).Interface().(N)
		case U32:
			return U8(n).Interface().(N)
		case I32:
			return U8(n).Interface().(N)
		case U64:
			return U8(n).Interface().(N)
		case I64:
			return U8(n).Interface().(N)
		case U128:
			return U8(n.ToBigInt().Uint64()).Interface().(N)
		case I128:
			return U8(n.ToBigInt().Uint64()).Interface().(N)
		default:
			panic("unknown numeric type for U8")
		}

	case I8:
		switch n := n.(type) {
		case U8:
			return I8(n).Interface().(N)
		case I8:
			return I8(n).Interface().(N)
		case U16:
			return I8(n).Interface().(N)
		case I16:
			return I8(n).Interface().(N)
		case U32:
			return I8(n).Interface().(N)
		case I32:
			return I8(n).Interface().(N)
		case U64:
			return I8(n).Interface().(N)
		case I64:
			return I8(n).Interface().(N)
		case U128:
			return I8(n.ToBigInt().Int64()).Interface().(N)
		case I128:
			return I8(n.ToBigInt().Int64()).Interface().(N)
		default:
			panic("unknown numeric type for I8")
		}

	case U16:
		switch n := n.(type) {
		case U8:
			return U16(n).Interface().(N)
		case I8:
			return U16(n).Interface().(N)
		case U16:
			return U16(n).Interface().(N)
		case I16:
			return U16(n).Interface().(N)
		case U32:
			return U16(n).Interface().(N)
		case I32:
			return U16(n).Interface().(N)
		case U64:
			return U16(n).Interface().(N)
		case I64:
			return U16(n).Interface().(N)
		case U128:
			return U16(n.ToBigInt().Uint64()).Interface().(N)
		case I128:
			return U16(n.ToBigInt().Uint64()).Interface().(N)
		default:
			panic("unknown numeric type for U16")
		}

	case I16:
		switch n := n.(type) {
		case U8:
			return I16(n).Interface().(N)
		case I8:
			return I16(n).Interface().(N)
		case U16:
			return I16(n).Interface().(N)
		case I16:
			return I16(n).Interface().(N)
		case U32:
			return I16(n).Interface().(N)
		case I32:
			return I16(n).Interface().(N)
		case U64:
			return I16(n).Interface().(N)
		case I64:
			return I16(n).Interface().(N)
		case U128:
			return I16(n.ToBigInt().Int64()).Interface().(N)
		case I128:
			return I16(n.ToBigInt().Int64()).Interface().(N)
		default:
			panic("unknown numeric type for I16")
		}

	case U32:
		switch n := n.(type) {
		case U8:
			return U32(n).Interface().(N)
		case I8:
			return U32(n).Interface().(N)
		case U16:
			return U32(n).Interface().(N)
		case I16:
			return U32(n).Interface().(N)
		case U32:
			return U32(n).Interface().(N)
		case I32:
			return U32(n).Interface().(N)
		case U64:
			return U32(n).Interface().(N)
		case I64:
			return U32(n).Interface().(N)
		case U128:
			return U32(n.ToBigInt().Uint64()).Interface().(N)
		case I128:
			return U32(n.ToBigInt().Uint64()).Interface().(N)
		default:
			panic("unknown numeric type for U32")
		}

	case I32:
		switch n := n.(type) {
		case U8:
			return I32(n).Interface().(N)
		case I8:
			return I32(n).Interface().(N)
		case U16:
			return I32(n).Interface().(N)
		case I16:
			return I32(n).Interface().(N)
		case U32:
			return I32(n).Interface().(N)
		case I32:
			return I32(n).Interface().(N)
		case U64:
			return I32(n).Interface().(N)
		case I64:
			return I32(n).Interface().(N)
		case U128:
			return I32(n.ToBigInt().Int64()).Interface().(N)
		case I128:
			return I32(n.ToBigInt().Int64()).Interface().(N)
		default:
			panic("unknown numeric type for I32")
		}

	case U64:
		switch n := n.(type) {
		case U8:
			return U64(n).Interface().(N)
		case I8:
			return U64(n).Interface().(N)
		case U16:
			return U64(n).Interface().(N)
		case I16:
			return U64(n).Interface().(N)
		case U32:
			return U64(n).Interface().(N)
		case I32:
			return U64(n).Interface().(N)
		case U64:
			return U64(n).Interface().(N)
		case I64:
			return U64(n).Interface().(N)
		case U128:
			return U64(n.ToBigInt().Uint64()).Interface().(N)
		case I128:
			return U64(n.ToBigInt().Uint64()).Interface().(N)
		default:
			panic("unknown numeric type for U64")
		}

	case I64:
		switch n := n.(type) {
		case U8:
			return I64(n).Interface().(N)
		case I8:
			return I64(n).Interface().(N)
		case U16:
			return I64(n).Interface().(N)
		case I16:
			return I64(n).Interface().(N)
		case U32:
			return I64(n).Interface().(N)
		case I32:
			return I64(n).Interface().(N)
		case U64:
			return I64(n).Interface().(N)
		case I64:
			return I64(n).Interface().(N)
		case U128:
			return I64(n.ToBigInt().Int64()).Interface().(N)
		case I128:
			return I64(n.ToBigInt().Int64()).Interface().(N)
		default:
			panic("unknown numeric type for I64")
		}

	case U128:
		switch n := n.(type) {
		case U8:
			return NewU128(uint64(n)).Interface().(N)
		case I8:
			return NewU128(uint64(n)).Interface().(N)
		case U16:
			return NewU128(uint64(n)).Interface().(N)
		case I16:
			return NewU128(uint64(n)).Interface().(N)
		case U32:
			return NewU128(uint64(n)).Interface().(N)
		case I32:
			return NewU128(uint64(n)).Interface().(N)
		case U64:
			return NewU128(uint64(n)).Interface().(N)
		case I64:
			return NewU128(uint64(n)).Interface().(N)
		case U128:
			return n.Interface().(N)
		case I128:
			return U128(n).Interface().(N)
		default:
			panic("unknown numeric type for U128")
		}

	case I128:
		switch n := n.(type) {
		case U8:
			return NewI128(int64(n)).Interface().(N)
		case I8:
			return NewI128(int64(n)).Interface().(N)
		case U16:
			return NewI128(int64(n)).Interface().(N)
		case I16:
			return NewI128(int64(n)).Interface().(N)
		case U32:
			return NewI128(int64(n)).Interface().(N)
		case I32:
			return NewI128(int64(n)).Interface().(N)
		case U64:
			return NewI128(int64(n)).Interface().(N)
		case I64:
			return NewI128(int64(n)).Interface().(N)
		case U128:
			return To[I128](n).Interface().(N)
		case I128:
			return n.Interface().(N)
		default:
			panic("unknown numeric type N in To[N]")
		}

	default:
		panic("unknown numeric type N in To[N]")
	}
}

func fromAnyNumberTo128[N Numeric](n any) N {
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
		return bigIntToGeneric[N](new(big.Int).SetInt64(int64(n)))
	case uint64:
		return bigIntToGeneric[N](new(big.Int).SetUint64(uint64(n)))
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
	case *big.Int:
		return bigIntToGeneric[N](n)
	case string:
		bn, ok := new(big.Int).SetString(n, 10)
		if !ok {
			panic("could not convert string to big.Int")
		}
		return bigIntToGeneric[N](bn)
	default:
		panic("unknown type in fromAnyNumberTo")
	}
}

func bigIntToGeneric[N Numeric](bn *big.Int) N {
	switch reflect.Zero(reflect.TypeOf(*new(N))).Interface().(type) {
	case I128:
		return bigIntToI128(bn).Interface().(N)
	case U128:
		return bigIntToU128(bn).Interface().(N)
	default:
		panic("unknown numeric type in bigIntToGeneric")
	}
}
