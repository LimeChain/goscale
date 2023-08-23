package goscale

// TODO: this can be easily extracted into a separate package
//
// Example:
// package numeric
// type U8Ext uint8
//
// package goscale
// type U8 U8Ext

import (
	"errors"
	"reflect"
)

var ErrOverflow = errors.New("overflow")

// Go's primitive numeric types don't have a common interface.
// To be able to write generic code that works with any numeric type,
// including some custom types, we need to define a common interface.

type Numeric interface {
	Encodable // TODO: move it out of here

	ToNumeric() Numeric // TODO: rename to ToGenericNum ?

	Add(other Numeric) Numeric
	Sub(other Numeric) Numeric
	Mul(other Numeric) Numeric
	Div(other Numeric) Numeric
	Mod(other Numeric) Numeric

	Eq(other Numeric) bool
	Ne(other Numeric) bool
	Lt(other Numeric) bool
	Lte(other Numeric) bool
	Gt(other Numeric) bool
	Gte(other Numeric) bool

	Max(other Numeric) Numeric
	Min(other Numeric) Numeric
	Clamp(minValue, maxValue Numeric) Numeric
	TrailingZeros() Numeric

	SaturatingAdd(other Numeric) Numeric
	SaturatingSub(other Numeric) Numeric
	// SaturatingMul(other Numeric) Numeric
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
	// AndNot(other Numeric) Numeric

	// add more methods as needed
}

// TODO: add support for passing uint, int, U8, I8, etc
// TODO: finish the U128, I128 implementation
// TODO: rename to NewNum ?
func NewNumeric[N Numeric](n any) N {
	switch reflect.Zero(reflect.TypeOf(*new(N))).Interface().(type) {
	case U8:
		switch n := n.(type) {
		case uint8:
			return U8(n).ToNumeric().(N)
		case int8:
			return U8(n).ToNumeric().(N)
		case uint16:
			return U8(n).ToNumeric().(N)
		case int16:
			return U8(n).ToNumeric().(N)
		case uint32:
			return U8(n).ToNumeric().(N)
		case int32:
			return U8(n).ToNumeric().(N)
		case uint64:
			return U8(n).ToNumeric().(N)
		case int64:
			return U8(n).ToNumeric().(N)
		default:
			panic("unknown primitive type for U8")
		}

	case I8:
		switch n := n.(type) {
		case uint8:
			return I8(n).ToNumeric().(N)
		case int8:
			return I8(n).ToNumeric().(N)
		case uint16:
			return I8(n).ToNumeric().(N)
		case int16:
			return I8(n).ToNumeric().(N)
		case uint32:
			return I8(n).ToNumeric().(N)
		case int32:
			return I8(n).ToNumeric().(N)
		case uint64:
			return I8(n).ToNumeric().(N)
		case int64:
			return I8(n).ToNumeric().(N)
		default:
			panic("unknown primitive type for I8")
		}

	case U16:
		switch n := n.(type) {
		case uint8:
			return U16(n).ToNumeric().(N)
		case int8:
			return U16(n).ToNumeric().(N)
		case uint16:
			return U16(n).ToNumeric().(N)
		case int16:
			return U16(n).ToNumeric().(N)
		case uint32:
			return U16(n).ToNumeric().(N)
		case int32:
			return U16(n).ToNumeric().(N)
		case uint64:
			return U16(n).ToNumeric().(N)
		case int64:
			return U16(n).ToNumeric().(N)
		default:
			panic("unknown primitive type for U16")
		}

	case I16:
		switch n := n.(type) {
		case uint8:
			return I16(n).ToNumeric().(N)
		case int8:
			return I16(n).ToNumeric().(N)
		case uint16:
			return I16(n).ToNumeric().(N)
		case int16:
			return I16(n).ToNumeric().(N)
		case uint32:
			return I16(n).ToNumeric().(N)
		case int32:
			return I16(n).ToNumeric().(N)
		case uint64:
			return I16(n).ToNumeric().(N)
		case int64:
			return I16(n).ToNumeric().(N)
		default:
			panic("unknown primitive type for I16")
		}

	case U32:
		switch n := n.(type) {
		case uint8:
			return U32(n).ToNumeric().(N)
		case int8:
			return U32(n).ToNumeric().(N)
		case uint16:
			return U32(n).ToNumeric().(N)
		case int16:
			return U32(n).ToNumeric().(N)
		case uint32:
			return U32(n).ToNumeric().(N)
		case int32:
			return U32(n).ToNumeric().(N)
		case uint64:
			return U32(n).ToNumeric().(N)
		case int64:
			return U32(n).ToNumeric().(N)
		default:
			panic("unknown primitive type for U32")
		}

	case I32:
		switch n := n.(type) {
		case uint8:
			return I32(n).ToNumeric().(N)
		case int8:
			return I32(n).ToNumeric().(N)
		case uint16:
			return I32(n).ToNumeric().(N)
		case int16:
			return I32(n).ToNumeric().(N)
		case uint32:
			return I32(n).ToNumeric().(N)
		case int32:
			return I32(n).ToNumeric().(N)
		case uint64:
			return I32(n).ToNumeric().(N)
		case int64:
			return I32(n).ToNumeric().(N)
		default:
			panic("unknown primitive type for I32")
		}

	case U64:
		switch n := n.(type) {
		case uint8:
			return U64(n).ToNumeric().(N)
		case int8:
			return U64(n).ToNumeric().(N)
		case uint16:
			return U64(n).ToNumeric().(N)
		case int16:
			return U64(n).ToNumeric().(N)
		case uint32:
			return U64(n).ToNumeric().(N)
		case int32:
			return U64(n).ToNumeric().(N)
		case uint64:
			return U64(n).ToNumeric().(N)
		case int64:
			return U64(n).ToNumeric().(N)
		default:
			panic("unknown primitive type for U64")
		}

	case I64:
		switch n := n.(type) {
		case uint8:
			return I64(n).ToNumeric().(N)
		case int8:
			return I64(n).ToNumeric().(N)
		case uint16:
			return I64(n).ToNumeric().(N)
		case int16:
			return I64(n).ToNumeric().(N)
		case uint32:
			return I64(n).ToNumeric().(N)
		case int32:
			return I64(n).ToNumeric().(N)
		case uint64:
			return I64(n).ToNumeric().(N)
		case int64:
			return I64(n).ToNumeric().(N)
		default:
			panic("unknown primitive type for I64")
		}

	default:
		panic("unknown numeric type")
	}
}

// TODO: implement the rest of the methods

func ToU16(n Numeric) U16 {
	switch reflect.TypeOf(n) {
	case reflect.TypeOf(*new(U8)):
		return U16(n.(U8))
	case reflect.TypeOf(*new(U16)):
		return n.(U16)
	case reflect.TypeOf(*new(U32)):
		return U16(n.(U32))
	case reflect.TypeOf(*new(U64)):
		return U16(n.(U64))
	default:
		panic("unknown numeric type")
	}
}

func ToU64(n Numeric) U64 {
	switch reflect.TypeOf(n) {
	case reflect.TypeOf(*new(U8)):
		return U64(n.(U8))
	case reflect.TypeOf(*new(U16)):
		return U64(n.(U16))
	case reflect.TypeOf(*new(U32)):
		return U64(n.(U32))
	case reflect.TypeOf(*new(U64)):
		return n.(U64)
	default:
		panic("unknown numeric type")
	}
}
