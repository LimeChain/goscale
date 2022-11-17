package goscale

import (
	"bytes"
	"reflect"
)

/*
	https://spec.polkadot.network/#defn-scale-tuple

	SCALE Tuple type translates to Go's struct type.
*/

type Tuple interface{}

func EncodeTuple(value Tuple, buffer *bytes.Buffer) {
	t := reflect.TypeOf(value)

	if t.Kind() != reflect.Struct {
		panic("not a Tuple type")
	}

	v := reflect.ValueOf(value)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if t.Field(i).IsExported() {
			switch field.Kind() {
			case reflect.Bool:
				(field.Interface().(Bool)).Encode(buffer)
			case reflect.Int:
				(field.Interface().(Compact)).Encode(buffer)
			case reflect.Uint:
				(field.Interface().(Compact)).Encode(buffer)
			case reflect.Int8:
				(field.Interface().(I8)).Encode(buffer)
			case reflect.Uint8:
				(field.Interface().(U8)).Encode(buffer)
			case reflect.Int16:
				(field.Interface().(I16)).Encode(buffer)
			case reflect.Uint16:
				(field.Interface().(U16)).Encode(buffer)
			case reflect.Int32:
				(field.Interface().(I32)).Encode(buffer)
			case reflect.Uint32:
				(field.Interface().(U32)).Encode(buffer)
			case reflect.Int64:
				(field.Interface().(I64)).Encode(buffer)
			case reflect.Uint64:
				(field.Interface().(U64)).Encode(buffer)
			case reflect.String:
				SliceU8ToSequenceU8(StringToSliceU8(field.Interface().(string))).Encode(buffer)
			case reflect.Slice:
				// TODO handle the cases of sequence of any type
				// field.Interface().(Any type).Encode(buffer)
				SliceU8ToSequenceU8(StringToSliceU8(string(field.Interface().([]U8)))).Encode(buffer)
			// case reflect.Array:
			case reflect.Struct:
				EncodeTuple(field.Interface(), buffer)
			// case reflect.Float32:
			// case reflect.Float64:
			// case reflect.Complex64:
			// case reflect.Complex128:
			// case reflect.Uintptr:
			// case reflect.UnsafePointer:
			// case reflect.Pointer:
			// case reflect.Chan:
			// case reflect.Func:
			// case reflect.Map:
			// case reflect.Interface:
			default:
				panic("unreachable case")
			}
		}
	}
}
