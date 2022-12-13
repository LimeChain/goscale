package goscale

import (
	"bytes"
	"reflect"
)

/*
	https://spec.polkadot.network/#defn-scale-tuple

	SCALE Tuple type translates to Go's struct type.
*/

func ConvertTo[T Encodable](v reflect.Value) Encodable {
	value, ok := v.Interface().(T)
	if !ok {
		panic("unable to convert to []" + reflect.TypeOf(*new(T)).Name())
	}
	return value
}

type Tuple interface {
	Encodable
}

func EncodeTuple(t interface{}, buffer *bytes.Buffer) {
	tVal := reflect.ValueOf(t)

	if tVal.Kind() != reflect.Struct {
		panic("not a SCALE Tuple type" + tVal.Type().Name())
	}

	for i := range reflect.VisibleFields(tVal.Type()) {
		field := tVal.Field(i)

		switch field.Kind() {
		case reflect.Bool:
			(field.Interface().(Bool)).Encode(buffer)
		case reflect.Uint:
			(field.Interface().(Compact)).Encode(buffer)
		case reflect.Int:
			(field.Interface().(Compact)).Encode(buffer)
		case reflect.Uint8:
			(field.Interface().(U8)).Encode(buffer)
		case reflect.Int8:
			(field.Interface().(I8)).Encode(buffer)
		case reflect.Uint16:
			(field.Interface().(U16)).Encode(buffer)
		case reflect.Int16:
			(field.Interface().(I16)).Encode(buffer)
		case reflect.Uint32:
			(field.Interface().(U32)).Encode(buffer)
		case reflect.Int32:
			(field.Interface().(I32)).Encode(buffer)
		case reflect.Uint64:
			(field.Interface().(U64)).Encode(buffer)
		case reflect.Int64:
			(field.Interface().(I64)).Encode(buffer)
		case reflect.String:
			(field.Interface().(Str)).Encode(buffer)
		case reflect.Array:
			// handles U128, I128
			switch field.Type().Name() {
			case "U128":
				field.Interface().(U128).Encode(buffer)
			case "I128":
				field.Interface().(I128).Encode(buffer)
			default:
				panic("unreachable case for Array: " + field.Type().Name())
			}
		case reflect.Slice:
			// handles Sequence[T], FixedSequence[T], VaryingData
			switch reflect.SliceOf(field.Type()) {
			case reflect.TypeOf([]Sequence[Bool]{}):
				ConvertTo[Sequence[Bool]](field).Encode(buffer)
			case reflect.TypeOf([]Sequence[U8]{}):
				ConvertTo[Sequence[U8]](field).Encode(buffer)
			case reflect.TypeOf([]Sequence[I8]{}):
				ConvertTo[Sequence[I8]](field).Encode(buffer)
			case reflect.TypeOf([]Sequence[U16]{}):
				ConvertTo[Sequence[U16]](field).Encode(buffer)
			case reflect.TypeOf([]Sequence[I16]{}):
				ConvertTo[Sequence[I16]](field).Encode(buffer)
			case reflect.TypeOf([]Sequence[U32]{}):
				ConvertTo[Sequence[U32]](field).Encode(buffer)
			case reflect.TypeOf([]Sequence[I32]{}):
				ConvertTo[Sequence[I32]](field).Encode(buffer)
			case reflect.TypeOf([]Sequence[U64]{}):
				ConvertTo[Sequence[U64]](field).Encode(buffer)
			case reflect.TypeOf([]Sequence[I64]{}):
				ConvertTo[Sequence[I64]](field).Encode(buffer)
			case reflect.TypeOf([]Sequence[Compact]{}):
				ConvertTo[Sequence[Compact]](field).Encode(buffer)
			case reflect.TypeOf([]Sequence[CompactU128]{}):
				ConvertTo[Sequence[CompactU128]](field).Encode(buffer)
			case reflect.TypeOf([]Sequence[Str]{}):
				ConvertTo[Sequence[Str]](field).Encode(buffer)
			case reflect.TypeOf([]Sequence[VaryingData]{}):
				panic("encoding of SCALE Sequence[VaryingData] is not implemented")
			default:
				// TODO T: {Sequence, FixedSequence, Dictionary, Option, Result, Tuple}

				switch field.Type().Name() {
				case "VaryingData":
					ConvertTo[VaryingData](field).Encode(buffer)
				default:
					println(reflect.SliceOf(field.Type()))
					// EncodeTuple(field.Interface(), buffer)
					panic("unreachable case for Slice: " + field.Type().Name())
				}
			}

		case reflect.Map:
			panic("encoding of SCALE Dictionary field is not implemented")
		case reflect.Struct:
			// handles Empty, CompactU128
			switch field.Type().Name() {
			case "Empty":
				EncodeTuple(field.Interface(), buffer)
			case "CompactU128":
				field.Interface().(CompactU128).Encode(buffer)
			default:
				// handles Option[T], Result[T], Tuple
				if field.Kind().String() == "struct" {
					EncodeTuple(field.Interface(), buffer)
				} else {
					panic("unreachable case for Struct: " + field.Type().Name())
				}
			}
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
			// Here it does nothing, but that allows the usage of the embedded Encodable
			// in custom-defined structs which allows using them in places where Encodable
			// is expected like in the case of Option[T], Result[T].
		default:
			panic("unreachable case")
		}
	}
}
