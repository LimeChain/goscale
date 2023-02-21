package goscale

import (
	"bytes"
	"reflect"
)

/*
	https://spec.polkadot.network/#defn-scale-tuple

	SCALE Tuple type translates to Go's struct type.
*/

type Tuple struct {
	Encodable
}

func (t Tuple) Encode(buffer *bytes.Buffer) {
	panic("allows the Tuple type to conform to the Encodable interface")
}

func (t Tuple) Bytes() []byte {
	panic("allows the Tuple type to conform to the Encodable interface")
}

func EncodeTuple(t interface{}, buffer *bytes.Buffer) {
	tVal := reflect.ValueOf(t)

	if tVal.Kind() != reflect.Struct {
		panic("not a SCALE Tuple type")
	}

	// Tinygo does not support: reflect.VisibleFields(tVal.Type())
	for i := 0; i < tVal.NumField(); i++ {
		field := tVal.Field(i)

		if reflect.TypeOf(t).Field(i).IsExported() {

			switch field.Kind() {
			case reflect.Bool:
				ConvertTo[Bool](field).Encode(buffer)
			case reflect.Uint8:
				ConvertTo[U8](field).Encode(buffer)
			case reflect.Int8:
				ConvertTo[I8](field).Encode(buffer)
			case reflect.Uint16:
				ConvertTo[U16](field).Encode(buffer)
			case reflect.Int16:
				ConvertTo[I16](field).Encode(buffer)
			case reflect.Uint32:
				ConvertTo[U32](field).Encode(buffer)
			case reflect.Int32:
				ConvertTo[I32](field).Encode(buffer)
			case reflect.Uint64:
				ConvertTo[U64](field).Encode(buffer)
			case reflect.Int64:
				ConvertTo[I64](field).Encode(buffer)
			case reflect.String:
				ConvertTo[Str](field).Encode(buffer)
			case reflect.Array:
				// U128, I128, Compact
				switch field.Type() {
				case reflect.TypeOf(*new(U128)):
					ConvertTo[U128](field).Encode(buffer)
				case reflect.TypeOf(*new(I128)):
					ConvertTo[I128](field).Encode(buffer)
				case reflect.TypeOf(*new(Compact)):
					ConvertTo[Compact](field).Encode(buffer)
				default:
					panic("unreachable case (Array) in EncodeTuple")
				}
			case reflect.Slice:
				// Sequence[T], FixedSequence[T], VaryingData
				SequenceFieldEncode(field, buffer)
			case reflect.Map:
				DictionaryFieldEncode(field, buffer)
			case reflect.Struct:
				switch field.Type() {
				case reflect.TypeOf(*new(Empty)):
					EncodeTuple(field.Interface(), buffer)
				default:
					// Option[T], Result[T], Tuple
					if field.Kind() == reflect.Struct {
						EncodeTuple(field.Interface(), buffer)
					} else {
						panic("unreachable case (Struct) in EncodeTuple")
					}
				}
			case reflect.Int, reflect.Uint, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
				panic("encoding of T field is not supported")
			case reflect.Uintptr, reflect.UnsafePointer, reflect.Pointer, reflect.Chan, reflect.Func:
				panic("encoding of T field is not implemented")
			case reflect.Interface:
				/*
					Here it does nothing, but that allows the usage of the embedded Encodable
					in custom-defined structs which allows using them in places where Encodable
					is expected like in the case of Option[T], Result[T].
				*/
			default:
				panic("unreachable case in EncodeTuple")
			}

		}
	}
}

func SequenceFieldEncode(field reflect.Value, buffer *bytes.Buffer) {
	// Tinygo does not support: reflect.SliceOf(field.Type())
	switch field.Type().Elem() {
	case reflect.TypeOf(*new(Bool)):
		ConvertToSequence[Bool](field).Encode(buffer)
	case reflect.TypeOf(*new(U8)):
		ConvertToSequence[U8](field).Encode(buffer)
	case reflect.TypeOf(*new(I8)):
		ConvertToSequence[I8](field).Encode(buffer)
	case reflect.TypeOf(*new(U16)):
		ConvertToSequence[U16](field).Encode(buffer)
	case reflect.TypeOf(*new(I16)):
		ConvertToSequence[I16](field).Encode(buffer)
	case reflect.TypeOf(*new(U32)):
		ConvertToSequence[U32](field).Encode(buffer)
	case reflect.TypeOf(*new(I32)):
		ConvertToSequence[I32](field).Encode(buffer)
	case reflect.TypeOf(*new(U64)):
		ConvertToSequence[U64](field).Encode(buffer)
	case reflect.TypeOf(*new(I64)):
		ConvertToSequence[I64](field).Encode(buffer)
	case reflect.TypeOf(*new(U128)):
		ConvertToSequence[U128](field).Encode(buffer)
	case reflect.TypeOf(*new(I128)):
		ConvertToSequence[I128](field).Encode(buffer)
	case reflect.TypeOf(*new(Compact)):
		ConvertToSequence[Compact](field).Encode(buffer)
	case reflect.TypeOf(*new(Str)):
		ConvertToSequence[Str](field).Encode(buffer)
	case reflect.TypeOf(*new(VaryingData)):
		ConvertToSequence[VaryingData](field).Encode(buffer)

	// TODO: Sequence[Sequence[T]]
	// TODO: Sequence[FixedSequence[T]]
	// TODO: Sequence[Dictionary[T1, T2]]
	case reflect.TypeOf(*new(Sequence[Bool])):
		size := field.Len()
		ToCompact(size).Encode(buffer)
		for i := 0; i < field.Len(); i++ {
			SequenceFieldEncode(field.Index(i), buffer)
		}

	default:

		switch field.Type() {
		case reflect.TypeOf(*new(VaryingData)):
			ConvertTo[VaryingData](field).Encode(buffer)
		default:
			// Sequence[Option], Sequence[Result], Sequence[Tuple]

			// since there are infinite number of T we can't use switch
			size := field.Len()
			// TODO: if it is a FixedSequence[T] don't encode the length
			ToCompact(size).Encode(buffer)
			for i := 0; i < size; i++ {
				seqElem := field.Index(i)
				EncodeTuple(seqElem.Interface(), buffer)
			}
		}
	}
}

func DictionaryFieldEncode(field reflect.Value, buffer *bytes.Buffer) {
	// Tinygo does not support: reflect.MapOf(key.Type(), elem.Type())
	switch field.Type().Elem() {
	case reflect.TypeOf(*new(Bool)):
		ConvertToDictionary[Str, Bool](field).Encode(buffer)
	case reflect.TypeOf(*new(U8)):
		ConvertToDictionary[Str, U8](field).Encode(buffer)
	case reflect.TypeOf(*new(I8)):
		ConvertToDictionary[Str, I8](field).Encode(buffer)
	case reflect.TypeOf(*new(U16)):
		ConvertToDictionary[Str, U16](field).Encode(buffer)
	case reflect.TypeOf(*new(I16)):
		ConvertToDictionary[Str, I16](field).Encode(buffer)
	case reflect.TypeOf(*new(U32)):
		ConvertToDictionary[Str, U32](field).Encode(buffer)
	case reflect.TypeOf(*new(I32)):
		ConvertToDictionary[Str, I32](field).Encode(buffer)
	case reflect.TypeOf(*new(U64)):
		ConvertToDictionary[Str, U64](field).Encode(buffer)
	case reflect.TypeOf(*new(I64)):
		ConvertToDictionary[Str, I64](field).Encode(buffer)
	case reflect.TypeOf(*new(U128)):
		ConvertToDictionary[Str, U128](field).Encode(buffer)
	case reflect.TypeOf(*new(I128)):
		ConvertToDictionary[Str, I128](field).Encode(buffer)
	case reflect.TypeOf(*new(Compact)):
		ConvertToDictionary[Str, Compact](field).Encode(buffer)
	case reflect.TypeOf(*new(Str)):
		ConvertToDictionary[Str, Str](field).Encode(buffer)
	case reflect.TypeOf(*new(VaryingData)):
		ConvertToDictionary[Str, VaryingData](field).Encode(buffer)

	// TODO: Dictionary[Str, Sequence[T]]
	// TODO: Dictionary[Str, FixedSequence[T]]
	// TODO: Dictionary[Str, Dictionary]

	default:
		switch field.Type() {
		case reflect.TypeOf(*new(VaryingData)):
			ConvertTo[VaryingData](field).Encode(buffer)
		default:
			// Dictionary[Str, Option], Dictionary[Str, Result], Dictionary[Str, Tuple]

			// TODO: if it is a FixedSequence[T] don't encode the length
			size := field.Len()
			ToCompact(size).Encode(buffer)
			for i := 0; i < size; i++ {
				seqElem := field.Index(i)
				EncodeTuple(seqElem.Interface(), buffer)
			}
		}
	}
}

func ConvertToSequence[T Encodable](v reflect.Value) (value Encodable) {
	if v.Type() == reflect.TypeOf(*new(Sequence[T])) {
		value = ConvertTo[Sequence[T]](v)
	} else if v.Type() == reflect.TypeOf(*new(FixedSequence[T])) {
		value = ConvertTo[FixedSequence[T]](v)
	} else {
		panic("unable to convert to Sequence|FixedSequence[T]")
	}
	return value
}

func ConvertToDictionary[K Comparable, V Encodable](v reflect.Value) (value Encodable) {
	if v.Type() == reflect.TypeOf(*new(Dictionary[K, V])) {
		value = ConvertTo[Dictionary[K, V]](v)
	} else {
		panic("unable to convert to Dictionary[K, V]")
	}
	return value
}

func ConvertTo[T Encodable](v reflect.Value) Encodable {
	value, ok := v.Interface().(T)
	if !ok {
		panic("unable to convert to T")
	}
	return value
}
