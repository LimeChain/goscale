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
				(field.Interface().(Str)).Encode(buffer)
			case reflect.Slice:
				// TODO: handle Sequence of { Sequence | Fixed Sequence | Dictionary | Tuple }
				elemType := reflect.SliceOf(field.Type())

				if elemType == reflect.TypeOf([][]Bool{}) {
					elemValues, ok := field.Interface().([]Bool)
					if !ok {
						panic("unable to convert to Sequence[Bool]")
					}
					Sequence[Bool]{Values: elemValues}.Encode(buffer)

				} else if elemType == reflect.TypeOf([][]Compact{}) {
					elemValues, ok := field.Interface().([]Compact)
					if !ok {
						panic("unable to convert to Sequence[Compact]")
					}
					Sequence[Compact]{Values: elemValues}.Encode(buffer)

				} else if elemType == reflect.TypeOf([][]U8{}) {
					elemValues, ok := field.Interface().([]U8)
					if !ok {
						panic("unable to convert to Sequence[U8]")
					}
					Sequence[U8]{Values: elemValues}.Encode(buffer)

				} else if elemType == reflect.TypeOf([][]I8{}) {
					elemValues, ok := field.Interface().([]I8)
					if !ok {
						panic("unable to convert to Sequence[I8]")
					}
					Sequence[I8]{Values: elemValues}.Encode(buffer)

				} else if elemType == reflect.TypeOf([][]U16{}) {
					elemValues, ok := field.Interface().([]U16)
					if !ok {
						panic("unable to convert to Sequence[U16]")
					}
					Sequence[U16]{Values: elemValues}.Encode(buffer)

				} else if elemType == reflect.TypeOf([][]I16{}) {
					elemValues, ok := field.Interface().([]I16)
					if !ok {
						panic("unable to convert to Sequence[I16]")
					}
					Sequence[I16]{Values: elemValues}.Encode(buffer)

				} else if elemType == reflect.TypeOf([][]U32{}) {
					elemValues, ok := field.Interface().([]U32)
					if !ok {
						panic("unable to convert to Sequence[U32]")
					}
					Sequence[U32]{Values: elemValues}.Encode(buffer)

				} else if elemType == reflect.TypeOf([][]I32{}) {
					elemValues, ok := field.Interface().([]I32)
					if !ok {
						panic("unable to convert to Sequence[I32]")
					}
					Sequence[I32]{Values: elemValues}.Encode(buffer)

				} else if elemType == reflect.TypeOf([][]U64{}) {
					elemValues, ok := field.Interface().([]U64)
					if !ok {
						panic("unable to convert to Sequence[U64]")
					}
					Sequence[U64]{Values: elemValues}.Encode(buffer)

				} else if elemType == reflect.TypeOf([][]I64{}) {
					elemValues, ok := field.Interface().([]I64)
					if !ok {
						panic("unable to convert to Sequence[I64]")
					}
					Sequence[I64]{Values: elemValues}.Encode(buffer)

				} else {
					// panic("Tuple type encoding is not implemented")
					EncodeTuple(field, buffer)
				}
			case reflect.Array:
				panic("encoding of type Fixed Sequence is not implemented")
			case reflect.Struct:
				EncodeTuple(field.Interface(), buffer)
			case reflect.Map:
				panic("encoding of type Dictionary is not implemented")
			// case reflect.Float32:
			// case reflect.Float64:
			// case reflect.Complex64:
			// case reflect.Complex128:
			// case reflect.Uintptr:
			// case reflect.UnsafePointer:
			// case reflect.Pointer:
			// case reflect.Chan:
			// case reflect.Func:
			// case reflect.Interface:
			default:
				panic("unreachable case")
			}
		}
	}
}
