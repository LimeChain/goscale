package goscale

/*
	https://spec.polkadot.network/#defn-scale-list

	SCALE Sequence type translates to Go's byte slice, string and array types.
*/

import (
	"reflect"
)

// TODO: handle array types

func (enc Encoder) EncodeByteSlice(value []byte) {
	size := len(value)
	enc.EncodeUintCompact(uint64(size))
	enc.Write(value)
}

func (dec Decoder) DecodeByteSlice() []byte {
	size := dec.DecodeUintCompact()
	value := make([]byte, size)
	dec.Read(value)
	return value
}

func (enc Encoder) EncodeString(value string) {
	enc.EncodeByteSlice([]byte(value))
}

func (dec Decoder) DecodeString() string {
	return string(dec.DecodeByteSlice())
}

func (enc Encoder) EncodeSlice(value interface{}) {
	if reflect.TypeOf(value).Kind() != reflect.Slice {
		panic("Not a Sequence type")
	}

	size := reflect.ValueOf(value).Len()
	enc.EncodeUintCompact(uint64(size))

	switch reflect.TypeOf(value).Elem().Kind() {
	case reflect.Bool:
		for _, v := range value.([]bool) {
			enc.EncodeBool(v)
		}
	case reflect.Int:
		for _, v := range value.([]int) {
			enc.EncodeUintCompact(uint64(v))
		}
	case reflect.Uint:
		for _, v := range value.([]uint) {
			enc.EncodeUintCompact(uint64(v))
		}
	case reflect.Int8:
		for _, v := range value.([]int8) {
			enc.EncodeInt8(v)
		}
	case reflect.Uint8:
		enc.Write(value.([]uint8))
	case reflect.Int16:
		for _, v := range value.([]int16) {
			enc.EncodeInt16(v)
		}
	case reflect.Uint16:
		for _, v := range value.([]uint16) {
			enc.EncodeUint16(v)
		}
	case reflect.Int32:
		for _, v := range value.([]int32) {
			enc.EncodeInt32(v)
		}
	case reflect.Uint32:
		for _, v := range value.([]uint32) {
			enc.EncodeUint32(v)
		}
	case reflect.Int64:
		for _, v := range value.([]int64) {
			enc.EncodeInt64(v)
		}
	case reflect.Uint64:
		for _, v := range value.([]uint64) {
			enc.EncodeUint64(v)
		}
	// case reflect.Float32:
	// case reflect.Float64:
	// case reflect.Complex64:
	// case reflect.Complex128:
	// case reflect.Uintptr:
	case reflect.String:
		for _, v := range value.([]string) {
			enc.EncodeString(v)
		}
	// case reflect.UnsafePointer:
	// case reflect.Pointer:
	// case reflect.Slice:
	// case reflect.Array:
	// case reflect.Chan:
	// case reflect.Struct:
	// case reflect.Func:
	// case reflect.Map:
	// case reflect.Interface:
	default:
		panic("Unreachable case")
	}
}
