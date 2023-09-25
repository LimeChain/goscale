package goscale

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
	"math/big"
	"reflect"
)

var (
	errOverflow  = errors.New("overflow")
	errUnderflow = errors.New("underflow")
)

type Numeric interface {
	int | int8 | int16 | int32 | int64 | I128 | uint | uint8 | uint16 | uint32 | uint64 | U128
}

type Num[T Numeric] struct {
	value *big.Int
}

func NewNum[T Numeric](n *big.Int) *Num[T] {
	result := new(Num[T])

	err := result.check(n)
	if err != nil {
		panic(err)
	}

	result.value = n

	return result
}

// TODO:
func newNum[T Numeric](n *big.Int) Num[T] {
	return *NewNum[T](n)
}

func (n Num[T]) Add(other *Num[T]) *Num[T] {
	res := new(big.Int).Add(n.value, other.value)

	return &Num[T]{
		value: n.saturate(res),
	}
}

func (n Num[T]) Sub(other *Num[T]) *Num[T] {
	return &Num[T]{
		value: new(big.Int).Sub(n.value, other.value),
	}
}

func (n Num[T]) Mul(other *Num[T]) *Num[T] {
	return &Num[T]{
		value: new(big.Int).Mul(n.value, other.value),
	}
}

func (n Num[T]) SaturatingAdd(other *Num[T]) *Num[T] {
	return &Num[T]{
		value: new(big.Int).Add(n.value, other.value),
	}
}

func (n Num[T]) SaturatingSub(other *Num[T]) *Num[T] {
	res := new(big.Int).Sub(n.value, other.value)

	return &Num[T]{
		value: n.saturate(res),
	}
}

func (n Num[T]) SaturatingMul(other *Num[T]) *Num[T] {
	res := new(big.Int).Mul(n.value, other.value)

	return &Num[T]{
		value: n.saturate(res),
	}
}

func (n Num[T]) Div(other *Num[T]) *Num[T] {
	res := new(big.Int).Div(n.value, other.value)

	return &Num[T]{
		value: n.saturate(res),
	}
}

func (n Num[T]) CheckedAdd(other *Num[T]) (*Num[T], error) {
	value := new(big.Int).Add(n.value, other.value)

	return &Num[T]{
		value,
	}, n.check(value)
}

func (n Num[T]) Int64() int64 {
	return n.value.Int64()
}

func (n Num[T]) Uint64() uint64 {
	return n.value.Uint64()
}

func (n Num[T]) Cmp(other *Num[T]) int {
	return n.value.Cmp(other.value)
}

//func (n *Num[T]) SetUint64(value uint64) *Num[T] {
//	n.value.SetUint64(value)
//
//	return n
//}

func (n Num[T]) Encode(buffer *bytes.Buffer) {
	switch reflect.Zero(reflect.TypeOf(*new(T))).Interface().(type) {
	case int8, uint8:
		buffer.Write(n.value.Bytes())
	case int16, uint16:
		result := make([]byte, 2)
		binary.LittleEndian.PutUint16(result, uint16(n.value.Uint64()))
	case int32, uint32:
		result := make([]byte, 4)
		binary.LittleEndian.PutUint32(result, uint32(n.value.Uint64()))
	case int64, uint64:
		result := make([]byte, 8)
		binary.LittleEndian.PutUint64(result, n.value.Uint64())
	case I128:
		buffer.Write(NewI128FromBigInt(*n.value).Bytes())
	case U128:
		buffer.Write(NewU128FromBigInt(n.value).Bytes())
	default:
		panic("type not found")
	}
}

func (n Num[T]) Bytes() []byte {
	return EncodedBytes(n)
}

func (n Num[T]) saturate(v *big.Int) *big.Int {
	switch reflect.Zero(reflect.TypeOf(*new(T))).Interface().(type) {
	case int8:
		if v.Int64() > int64(math.MaxInt8) {
			return new(big.Int).SetInt64(math.MaxInt8)
		} else if v.Int64() < int64(math.MinInt8) {
			return new(big.Int).SetInt64(math.MinInt8)
		}
	case int16:
		if v.Int64() > int64(math.MaxInt16) {
			return new(big.Int).SetInt64(math.MaxInt16)
		} else if v.Int64() < int64(math.MinInt16) {
			return new(big.Int).SetInt64(math.MinInt16)
		}
	case int32:
		if v.Int64() > int64(math.MaxInt32) {
			return new(big.Int).SetInt64(math.MaxInt32)
		} else if v.Int64() < int64(math.MinInt32) {
			return new(big.Int).SetInt64(math.MinInt32)
		}
	case int64:
		if v.Int64() > math.MaxInt64 {
			return new(big.Int).SetInt64(math.MaxInt64)
		} else if v.Int64() < int64(math.MinInt64) {
			return new(big.Int).SetInt64(math.MinInt64)
		}
	case I128:
		if v.Cmp(maxI128()) > 0 {
			return maxI128()
		} else if v.Cmp(minI128()) < 0 {
			return minI128()
		}
	case uint8:
		if v.Uint64() > uint64(math.MaxUint8) {
			return big.NewInt(math.MaxUint8)
		} else if v.Cmp(big.NewInt(0)) < 0 {
			return big.NewInt(0)
		}
	case uint16:
		if v.Uint64() > uint64(math.MaxUint16) {
			return big.NewInt(math.MaxUint8)
		} else if v.Cmp(big.NewInt(0)) < 0 {
			return big.NewInt(0)
		}
	case uint32:
		if v.Uint64() > uint64(math.MaxUint32) {
			return big.NewInt(math.MaxUint32)
		} else if v.Cmp(big.NewInt(0)) < 0 {
			return big.NewInt(0)
		}
	case uint64:
		if v.Uint64() > uint64(math.MaxUint64) {
			return new(big.Int).SetUint64(math.MaxUint64)
		} else if v.Cmp(big.NewInt(0)) < 0 {
			return big.NewInt(0)
		}
	case U128:
		if v.Cmp(maxU128()) > 0 {
			return maxU128()
		} else if v.Cmp(big.NewInt(0)) < 0 {
			return big.NewInt(0)
		}
	default:
		panic("type not found")
	}

	return v
}

func (n Num[T]) check(v *big.Int) error {
	switch reflect.Zero(reflect.TypeOf(*new(T))).Interface().(type) {
	case int8:
		if v.Int64() > int64(math.MaxInt8) {
			return errOverflow
		} else if v.Int64() < int64(math.MinInt8) {
			return errUnderflow
		}
	case int16:
		if v.Int64() > int64(math.MaxInt16) {
			return errOverflow
		} else if v.Int64() < int64(math.MinInt16) {
			return errUnderflow
		}
	case int32:
		if v.Int64() > int64(math.MaxInt32) {
			return errOverflow
		} else if v.Int64() < int64(math.MinInt32) {
			return errUnderflow
		}
	case int64:
		if v.Int64() > math.MaxInt64 {
			return errOverflow
		} else if v.Int64() < int64(math.MinInt64) {
			return errUnderflow
		}
	case I128:
		if v.Cmp(maxI128()) > 0 {
			return errOverflow
		} else if v.Cmp(minI128()) < 0 {
			return errUnderflow
		}
	case uint8:
		if v.Uint64() > uint64(math.MaxUint8) {
			return errOverflow
		} else if v.Cmp(big.NewInt(0)) < 0 {
			return errUnderflow
		}
	case uint16:
		if v.Uint64() > uint64(math.MaxUint16) {
			return errOverflow
		} else if v.Cmp(big.NewInt(0)) < 0 {
			return errUnderflow
		}
	case uint32:
		if v.Uint64() > uint64(math.MaxUint32) {
			return errOverflow
		} else if v.Cmp(big.NewInt(0)) < 0 {
			return errUnderflow
		}
	case uint64:
		if v.Uint64() > uint64(math.MaxUint64) {
			return errOverflow
		} else if v.Cmp(big.NewInt(0)) < 0 {
			return errUnderflow
		}
	case U128:
		if v.Cmp(maxU128()) > 0 {
			return errOverflow
		} else if v.Cmp(big.NewInt(0)) < 0 {
			return errUnderflow
		}
	default:
		panic("type not found")
	}

	return nil
}

func DecodeNum[T Numeric](buffer *bytes.Buffer) Num[T] {
	var result interface{}

	switch reflect.Zero(reflect.TypeOf(*new(T))).Interface().(type) {
	case uint8:
		result = newNum[uint8](big.NewInt(int64(DecodeU8(buffer))))
	case int8:
		result = newNum[int8](big.NewInt(int64(DecodeI8(buffer))))
	case uint16:
		result = newNum[uint16](big.NewInt(int64(DecodeU16(buffer))))
	case int16:
		result = newNum[int16](big.NewInt(int64(DecodeI16(buffer))))
	case uint32:
		result = newNum[uint32](big.NewInt(int64(DecodeU32(buffer))))
	case int32:
		result = newNum[int32](big.NewInt(int64(DecodeI32(buffer))))
	case uint64:
		result = newNum[uint64](big.NewInt(int64(DecodeU64(buffer))))
	case int64:
		result = newNum[int64](big.NewInt(int64(DecodeI64(buffer))))
	case U128:
		result = newNum[U128](DecodeU128(buffer).ToBigInt())
	case I128:
		result = newNum[I128](big.NewInt(int64(DecodeU8(buffer))))
	}

	return result.(Num[T])
}

func maxU128() *big.Int {
	maxU128 := new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil)

	return maxU128.Sub(maxU128, big.NewInt(1))
}

func maxI128() *big.Int {
	maxI128 := new(big.Int).Exp(big.NewInt(2), big.NewInt(127), nil)

	return maxI128.Sub(maxI128, big.NewInt(1))
}

func minI128() *big.Int {
	minI128 := new(big.Int).Exp(big.NewInt(2), big.NewInt(127), nil)

	return minI128.Neg(minI128)
}

func EncodeNumeric[T Numeric](num T, buffer *bytes.Buffer) {
	var result []byte

	switch reflect.Zero(reflect.TypeOf(*new(T))).Interface().(type) {
	case int8:
		buffer.WriteByte(byte(reflect.ValueOf(num).Int()))
	case int16:
		result = make([]byte, 2)
		binary.LittleEndian.PutUint16(result, uint16(reflect.ValueOf(num).Int()))
	case int32:
		result = make([]byte, 4)
		binary.LittleEndian.PutUint32(result, uint32(reflect.ValueOf(num).Int()))
	case int64:
		result = make([]byte, 8)
		binary.LittleEndian.PutUint64(result, uint64(reflect.ValueOf(num).Int()))
	case I128:
		result = reflect.ValueOf(num).Interface().(I128).Bytes()
	case uint8:
		buffer.WriteByte(byte(reflect.ValueOf(num).Uint()))
	case uint16:
		result = make([]byte, 2)
		binary.LittleEndian.PutUint16(result, uint16(reflect.ValueOf(num).Uint()))
	case uint32:
		result = make([]byte, 4)
		binary.LittleEndian.PutUint32(result, uint32(reflect.ValueOf(num).Uint()))
	case uint64:
		result = make([]byte, 8)
		binary.LittleEndian.PutUint64(result, reflect.ValueOf(num).Uint())
	case U128:
		result = reflect.ValueOf(num).Interface().(U128).Bytes()
	default:
		panic("type not found")
	}

	buffer.Write(result)
}
