package goscale

// TODO: this should be extracted into a separate package
//
// package numeric
//
// type U8 uint8
//
// ...
//
// package goscale
//
// type U8 numeric.U8

import (
	"errors"
	"reflect"
)

var ErrOverflow = errors.New("overflow")

// Go's primitive numeric types don't have a common interface.
// To be able to write generic code that works with any numeric
// type, including some custom types (U128/I128), we need to
// define a common interface.

type Numeric interface {
	Encodable // TODO: this belongs to the goscale package

	Interface() Numeric

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
	SaturatingMul(other Numeric) Numeric
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

// TODO: check when converting from bigger to smaller types (e.g. U64 -> U8)

// TODO: rename NewNumeric to New when this is extracted
// into a separate package (num.New, num.U8, num.I16, etc)
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
			return n.Interface().(N)
		case I8:
			return n.Interface().(N)
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
			return n.Interface().(N)
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
			return n.Interface().(N)
		case I32:
			return U32(n).Interface().(N)
		case U64:
			return U32(n).Interface().(N)
		case I64:
			return U32(n).Interface().(N)

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
			return n.Interface().(N)
		case U64:
			return I32(n).Interface().(N)
		case I64:
			return I32(n).Interface().(N)

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
			return n.Interface().(N)
		case I64:
			return U64(n).Interface().(N)

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
			return n.Interface().(N)

		default:
			panic("unknown primitive type for I64")
		}

	case U128:
		switch n := n.(type) {
		case uint:
			return NewU128FromUint64(uint64(n)).Interface().(N)
		case int:
			return NewU128FromUint64(uint64(n)).Interface().(N)
		case uint8:
			return NewU128FromUint64(uint64(n)).Interface().(N)
		case int8:
			return NewU128FromUint64(uint64(n)).Interface().(N)
		case uint16:
			return NewU128FromUint64(uint64(n)).Interface().(N)
		case int16:
			return NewU128FromUint64(uint64(n)).Interface().(N)
		case uint32:
			return NewU128FromUint64(uint64(n)).Interface().(N)
		case int32:
			return NewU128FromUint64(uint64(n)).Interface().(N)
		case uint64:
			return NewU128FromUint64(uint64(n)).Interface().(N)
		case int64:
			return NewU128FromUint64(uint64(n)).Interface().(N)
		case U8:
			return NewU128FromUint64(uint64(n)).Interface().(N)
		case I8:
			return NewU128FromUint64(uint64(n)).Interface().(N)
		case U16:
			return NewU128FromUint64(uint64(n)).Interface().(N)
		case I16:
			return NewU128FromUint64(uint64(n)).Interface().(N)
		case U32:
			return NewU128FromUint64(uint64(n)).Interface().(N)
		case I32:
			return NewU128FromUint64(uint64(n)).Interface().(N)
		case U64:
			return NewU128FromUint64(uint64(n)).Interface().(N)
		case I64:
			return NewU128FromUint64(uint64(n)).Interface().(N)
		case U128:
			return n.Interface().(N)
		case I128:
			return U128(n).Interface().(N)

		default:
			panic("unknown primitive type for U128")
		}

	default:
		panic("unknown numeric type")
	}
}

// TODO: implement the rest of the cases for To[N]
// Maybe add it to the Numeric interface, instead of
// having it as a separate function

func To[N Numeric](n Numeric) N {
	switch reflect.Zero(reflect.TypeOf(*new(N))).Interface().(type) {
	case U16:
		switch reflect.TypeOf(n) {
		case reflect.TypeOf(*new(U8)):
			return U16(n.(U8)).Interface().(N)
		case reflect.TypeOf(*new(U16)):
			return n.(U16).Interface().(N)
		case reflect.TypeOf(*new(U32)):
			return U16(n.(U32)).Interface().(N)
		case reflect.TypeOf(*new(U64)):
			return U16(n.(U64)).Interface().(N)
		case reflect.TypeOf(*new(U128)):
			return U16(n.(U128)[0]).Interface().(N)
		default:
			panic("unknown numeric type for U16")
		}

	case U64:
		switch reflect.TypeOf(n) {
		case reflect.TypeOf(*new(U8)):
			return U64(n.(U8)).Interface().(N)
		case reflect.TypeOf(*new(I8)):
			return U64(n.(I8)).Interface().(N)
		case reflect.TypeOf(*new(U16)):
			return U64(n.(U16)).Interface().(N)
		case reflect.TypeOf(*new(I16)):
			return U64(n.(I16)).Interface().(N)
		case reflect.TypeOf(*new(U32)):
			return U64(n.(U32)).Interface().(N)
		case reflect.TypeOf(*new(I32)):
			return U64(n.(I32)).Interface().(N)
		case reflect.TypeOf(*new(U64)):
			return U64(n.(U64)).Interface().(N)
		case reflect.TypeOf(*new(I64)):
			return U64(n.(I64)).Interface().(N)
		case reflect.TypeOf(*new(U128)):
			return U64(n.(U128)[0]).Interface().(N)
		case reflect.TypeOf(*new(I128)):
			return U64(n.(I128)[0]).Interface().(N)
		default:
			panic("unknown numeric type for U64")
		}
	default:
		panic("unknown numeric type for N")
	}
}
