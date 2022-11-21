package goscale

/*
	https://spec.polkadot.network/#defn-scale-tuple

	SCALE Tuple type translates to Go's struct type.
*/

type Tuple interface{}

// func (enc Encoder) EncodeTuple(value interface{}) {
// 	t := reflect.TypeOf(value)

// 	if t.Kind() != reflect.Struct {
// 		panic("Not a Tuple type")
// 	}

// 	v := reflect.ValueOf(value)

// 	for i := 0; i < v.NumField(); i++ {
// 		field := v.Field(i)

// 		if t.Field(i).IsExported() {
// 			switch field.Kind() {
// 			// case reflect.Bool:
// 			// 	enc.EncodeBool(field.Interface().(bool))
// 			// case reflect.Int:
// 			// 	enc.EncodeUintCompact(uint64(field.Interface().(int)))
// 			// case reflect.Uint:
// 			// 	enc.EncodeUintCompact(uint64(field.Interface().(uint)))
// 			// case reflect.Int8:
// 			// 	enc.EncodeInt8(field.Interface().(int8))
// 			case reflect.Uint8:
// 				// enc.EncodeUint8(field.Interface().(uint8))
// 			// case reflect.Int16:
// 			// 	enc.EncodeInt16(field.Interface().(int16))
// 			// case reflect.Uint16:
// 			// 	enc.EncodeUint16(field.Interface().(uint16))
// 			// case reflect.Int32:
// 			// 	enc.EncodeInt32(field.Interface().(int32))
// 			case reflect.Uint32:
// 				// enc.EncodeUint32(field.Interface().(uint32))
// 			// case reflect.Int64:
// 			// 	enc.EncodeInt64(field.Interface().(int64))
// 			// case reflect.Uint64:
// 			// 	enc.EncodeUint64(field.Interface().(uint64))
// 			case reflect.String:
// 				enc.EncodeString(field.Interface().(string))
// 			case reflect.Slice:
// 				enc.EncodeSlice(field.Interface())
// 			// case reflect.Array:
// 			case reflect.Struct:
// 				enc.EncodeTuple(field.Interface())
// 			// case reflect.Float32:
// 			// case reflect.Float64:
// 			// case reflect.Complex64:
// 			// case reflect.Complex128:
// 			// case reflect.Uintptr:
// 			// case reflect.UnsafePointer:
// 			// case reflect.Pointer:
// 			// case reflect.Chan:
// 			// case reflect.Func:
// 			// case reflect.Map:
// 			// case reflect.Interface:
// 			default:
// 				// panic("Unreachable case")
// 			}
// 		}
// 	}
// }
