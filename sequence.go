package goscale

/*
	https://spec.polkadot.network/#defn-scale-list

	SCALE Sequence type translates to Go's byte slice, string and array types.
*/

// TODO: handle Array types

type Sequence[T Encodable] struct {
	Values []T
}

func (seq Sequence[Encodable]) Encode(enc *Encoder) {
	Compact(len(seq.Values)).Encode(enc)
	for _, v := range seq.Values {
		v.Encode(enc)
	}
}

func DecodeSequence[T Encodable](dec *Decoder) Sequence[T] {
	size := dec.DecodeCompact()
	value := make([]byte, size)
	dec.Read(value)
	seq := Sequence[T]{Values: []T{}}
	return seq
}

func ToU8(s string) []U8 {

	for i, v := range []byte(s) {
		// TODO: check
		// result = append(result, sc.U8(v)) -> panic: cannot convert pointer to integer -> /tinygo/interp/memory.go:541
		result[i] = U8(v)
	}

	return result
}

func (seq Sequence[U8]) String() string {
	res := []byte{}

	// for _, v := range seq.Values {
	// 	res = append(res, byte(v))
	// }

	return string(res)
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
