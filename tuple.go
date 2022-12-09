package goscale

import (
	"bytes"
	"reflect"
)

/*
	https://spec.polkadot.network/#defn-scale-tuple

	SCALE Tuple type translates to Go's struct type.
*/

type Tuple[T any] struct {
	Data T
}

func (t Tuple[Encodable]) Encode(buffer *bytes.Buffer) {
	tVal := reflect.ValueOf(t.Data)
	tTyp := reflect.TypeOf(t.Data)

	if tTyp.Kind() != reflect.Struct {
		panic("not a SCALE Tuple type")
	}

	// for i, field := range reflect.VisibleFields(tt) {}
	for i := 0; i < tVal.NumField(); i++ {
		fieldT := tVal.Field(i)
		fieldV := tTyp.Field(i)

		if fieldV.IsExported() {
			println("field:", i, fieldV.Name)

			switch fieldT.Kind() {
			case reflect.Bool:
				(fieldT.Interface().(Bool)).Encode(buffer)
			case reflect.Uint:
				(fieldT.Interface().(Compact)).Encode(buffer)
			case reflect.Int:
				(fieldT.Interface().(Compact)).Encode(buffer)
			case reflect.Uint8:
				(fieldT.Interface().(U8)).Encode(buffer)
			case reflect.Int8:
				(fieldT.Interface().(I8)).Encode(buffer)
			case reflect.Uint16:
				(fieldT.Interface().(U16)).Encode(buffer)
			case reflect.Int16:
				(fieldT.Interface().(I16)).Encode(buffer)
			case reflect.Uint32:
				(fieldT.Interface().(U32)).Encode(buffer)
			case reflect.Int32:
				(fieldT.Interface().(I32)).Encode(buffer)
			case reflect.Uint64:
				(fieldT.Interface().(U64)).Encode(buffer)
			case reflect.Int64:
				(fieldT.Interface().(I64)).Encode(buffer)
			case reflect.String:
				(fieldT.Interface().(Str)).Encode(buffer)

			// TODO: handle Sequence of { Sequence | Fixed Sequence | Dictionary | Tuple }
			case reflect.Slice:
				elemType := reflect.SliceOf(fieldT.Type())

				if elemType == reflect.TypeOf([][]Bool{}) {
					elemValues, ok := fieldT.Interface().([]Bool)
					if !ok {
						panic("unable to convert to Sequence[Bool]")
					}
					Sequence[Bool](elemValues).Encode(buffer)

					// } else if elemType == reflect.TypeOf([][]Compact{}) {
					// 	elemValues, ok := fieldT.Interface().([]Compact)
					// 	if !ok {
					// 		panic("unable to convert to Sequence[Compact]")
					// 	}
					// 	Sequence[Compact](elemValues).Encode(buffer)

					// } else if elemType == reflect.TypeOf([][]U8{}) {
					// 	elemValues, ok := fieldT.Interface().([]U8)
					// 	if !ok {
					// 		panic("unable to convert to Sequence[U8]")
					// 	}
					// 	Sequence[U8](elemValues).Encode(buffer)

					// } else if elemType == reflect.TypeOf([][]I8{}) {
					// 	elemValues, ok := fieldT.Interface().([]I8)
					// 	if !ok {
					// 		panic("unable to convert to Sequence[I8]")
					// 	}
					// 	Sequence[I8](elemValues).Encode(buffer)

					// } else if elemType == reflect.TypeOf([][]U16{}) {
					// 	elemValues, ok := fieldT.Interface().([]U16)
					// 	if !ok {
					// 		panic("unable to convert to Sequence[U16]")
					// 	}
					// 	Sequence[U16](elemValues).Encode(buffer)

					// } else if elemType == reflect.TypeOf([][]I16{}) {
					// 	elemValues, ok := fieldT.Interface().([]I16)
					// 	if !ok {
					// 		panic("unable to convert to Sequence[I16]")
					// 	}
					// 	Sequence[I16](elemValues).Encode(buffer)

					// } else if elemType == reflect.TypeOf([][]U32{}) {
					// 	elemValues, ok := fieldT.Interface().([]U32)
					// 	if !ok {
					// 		panic("unable to convert to Sequence[U32]")
					// 	}
					// 	Sequence[U32](elemValues).Encode(buffer)

					// } else if elemType == reflect.TypeOf([][]I32{}) {
					// 	elemValues, ok := fieldT.Interface().([]I32)
					// 	if !ok {
					// 		panic("unable to convert to Sequence[I32]")
					// 	}
					// 	Sequence[I32](elemValues).Encode(buffer)

					// } else if elemType == reflect.TypeOf([][]U64{}) {
					// 	elemValues, ok := fieldT.Interface().([]U64)
					// 	if !ok {
					// 		panic("unable to convert to Sequence[U64]")
					// 	}
					// 	Sequence[U64](elemValues).Encode(buffer)

					// } else if elemType == reflect.TypeOf([][]I64{}) {
					// 	elemValues, ok := fieldT.Interface().([]I64)
					// 	if !ok {
					// 		panic("unable to convert to Sequence[I64]")
					// 	}
					// 	Sequence[I64](elemValues).Encode(buffer)

					// handle Sequence|FixedSequence types below
				} else if elemType == reflect.TypeOf([]Sequence[Bool]{}) {
					tryConvert[Sequence[Bool]](fieldT).Encode(buffer)
				} else if elemType == reflect.TypeOf([]Sequence[U8]{}) {
					tryConvert[Sequence[U8]](fieldT).Encode(buffer)
				} else if elemType == reflect.TypeOf([]Sequence[I8]{}) {
					tryConvert[Sequence[I8]](fieldT).Encode(buffer)
				} else if elemType == reflect.TypeOf([]Sequence[U16]{}) {
					tryConvert[Sequence[U16]](fieldT).Encode(buffer)
				} else if elemType == reflect.TypeOf([]Sequence[I16]{}) {
					tryConvert[Sequence[I16]](fieldT).Encode(buffer)
				} else if elemType == reflect.TypeOf([]Sequence[U32]{}) {
					tryConvert[Sequence[U32]](fieldT).Encode(buffer)
				} else if elemType == reflect.TypeOf([]Sequence[I32]{}) {
					tryConvert[Sequence[I32]](fieldT).Encode(buffer)
				} else if elemType == reflect.TypeOf([]Sequence[U64]{}) {
					tryConvert[Sequence[U64]](fieldT).Encode(buffer)
				} else if elemType == reflect.TypeOf([]Sequence[I64]{}) {
					tryConvert[Sequence[I64]](fieldT).Encode(buffer)
				} else {
					panic("case Slice: unknown type: " + fieldT.Type().Name())
				}
			// some SCALE defined types are represented as an array
			case reflect.Array:
				switch fieldT.Type().Name() {
				case "U128":
					fieldT.Interface().(U128).Encode(buffer)
				case "I128":
					fieldT.Interface().(I128).Encode(buffer)
				default:
					panic("case Array: " + fieldT.Type().Name())
				}
			// some SCALE defined types are represented as struct
			case reflect.Struct:
				switch fieldT.Type().Name() {
				case "Empty":
					fieldT.Interface().(Empty).Encode(buffer)
				case "CompactU128":
					fieldT.Interface().(CompactU128).Encode(buffer)
				default:
					panic("case Struct: " + fieldT.Type().Name())
				}
			case reflect.Map:
				panic("encoding of SCALE Dictionary field is not implemented")
			case reflect.Float32:
				panic("encoding of Float32 field is not implemented")
			case reflect.Float64:
				panic("encoding of Float64 field is not implemented")
			case reflect.Complex64:
				panic("encoding of Complex64 field is not implemented")
			case reflect.Complex128:
				panic("encoding of Complex128 field is not implemented")
			case reflect.Uintptr:
				panic("encoding of Uintptr field is not implemented")
			case reflect.UnsafePointer:
				panic("encoding of UnsafePointer field is not implemented")
			case reflect.Pointer:
				panic("encoding of Pointer field is not implemented")
			case reflect.Chan:
				panic("encoding of Chan field is not implemented")
			case reflect.Func:
				panic("encoding of Func field is not implemented")
			case reflect.Interface:
				panic("encoding of Interface field is not implemented")
			default:
				panic("unreachable case")
			}
		}
	}
}

func tryConvert[T Encodable](v reflect.Value) Encodable {
	value, ok := v.Interface().(T)
	if !ok {
		panic("unable to convert to []" + reflect.TypeOf(*new(T)).Name())
	}
	return value
}
