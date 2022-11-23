package goscale

import "strings"

/*
	https://spec.polkadot.network/#defn-scale-list

	SCALE Sequence type translates to Go's byte slice, string and array types.
*/

// TODO: handle fixed size sequence (Array type)

type Sequence[T Encodable] struct {
	Values []T
}

func (seq Sequence[Encodable]) Encode(enc *Encoder) {
	Compact(len(seq.Values)).Encode(enc)
	for _, v := range seq.Values {
		v.Encode(enc)
	}
}

func (dec *Decoder) DecodeSequenceU8() Sequence[U8] {
	return Sequence[U8]{Values: ToSliceU8(dec)}
}

func ToSliceU8(dec *Decoder) []U8 {
	size := dec.DecodeCompact()
	values := make([]U8, size)
	for i := 0; i < len(values); i++ {
		values[i] = dec.DecodeU8()
	}
	return values
}

func (seq Sequence[U8]) String() string {
	var res []string
	for _, v := range seq.Values {
		res = append(res, v.String())
	}
	return strings.Join(res, "")
}

// func (enc Encoder) EncodeByteSlice(value []byte) {
// 	size := len(value)
// 	enc.EncodeUintCompact(uint64(size))
// 	enc.Write(value)
// }

// func (dec Decoder) DecodeByteSlice() []byte {
// 	size := dec.DecodeUintCompact()
// 	value := make([]byte, size)
// 	dec.Read(value)
// 	return value
// }

// func (enc Encoder) EncodeString(value string) {
// 	enc.EncodeByteSlice([]byte(value))
// }

// func (dec Decoder) DecodeString() string {
// 	return string(dec.DecodeByteSlice())
// }

// func (enc Encoder) EncodeSlice(value interface{}) {
// 	if reflect.TypeOf(value).Kind() != reflect.Slice {
// 		panic("Not a Sequence type")
// 	}

// 	size := reflect.ValueOf(value).Len()
// 	enc.EncodeUintCompact(uint64(size))

// 	switch reflect.TypeOf(value).Elem().Kind() {
// 	case reflect.Bool:
// 		for _, v := range value.([]bool) {
// 			enc.EncodeBool(v)
// 		}
// 	case reflect.Int:
// 		for _, v := range value.([]int) {
// 			enc.EncodeUintCompact(uint64(v))
// 		}
// 	case reflect.Uint:
// 		for _, v := range value.([]uint) {
// 			enc.EncodeUintCompact(uint64(v))
// 		}
// 	case reflect.Int8:
// 		for _, v := range value.([]int8) {
// 			enc.EncodeInt8(v)
// 		}
// 	case reflect.Uint8:
// 		enc.Write(value.([]uint8))
// 	case reflect.Int16:
// 		for _, v := range value.([]int16) {
// 			enc.EncodeInt16(v)
// 		}
// 	case reflect.Uint16:
// 		for _, v := range value.([]uint16) {
// 			enc.EncodeUint16(v)
// 		}
// 	case reflect.Int32:
// 		for _, v := range value.([]int32) {
// 			enc.EncodeInt32(v)
// 		}
// 	case reflect.Uint32:
// 		for _, v := range value.([]uint32) {
// 			enc.EncodeUint32(v)
// 		}
// 	case reflect.Int64:
// 		for _, v := range value.([]int64) {
// 			enc.EncodeInt64(v)
// 		}
// 	case reflect.Uint64:
// 		for _, v := range value.([]uint64) {
// 			enc.EncodeUint64(v)
// 		}
// 	case reflect.String:
// 		for _, v := range value.([]string) {
// 			enc.EncodeString(v)
// 		}
// 	case reflect.Slice:
// 		switch value := value.(type) {
// 		case [][]uint8:
// 			for _, v := range value {
// 				enc.EncodeSlice(v)
// 			}
// 		}
// 	// case reflect.Array:
// 	case reflect.Struct:
// 		for _, v := range value.([]struct{}) {
// 			enc.EncodeTuple(v)
// 		}
// 	// case reflect.Float32:
// 	// case reflect.Float64:
// 	// case reflect.Complex64:
// 	// case reflect.Complex128:
// 	// case reflect.Uintptr:
// 	// case reflect.UnsafePointer:
// 	// case reflect.Pointer:
// 	// case reflect.Chan:
// 	// case reflect.Func:
// 	// case reflect.Map:
// 	// case reflect.Interface:
// 	default:
// 		// panic("Unreachable case")
// 	}
// }
