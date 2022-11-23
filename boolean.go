package goscale

/*
	https://spec.polkadot.network/#defn-scale-boolean

	SCALE's Boolean type translates to Go's boolean type.
	Values are encoded using the least significant bit of a single byte.
*/

type Bool bool

func (value Bool) Encode(enc *Encoder) {
	if value {
		enc.EncodeByte(0x01)
	} else {
		enc.EncodeByte(0x00)
	}
}

func (dec *Decoder) DecodeBool() Bool {
	buf := make([]byte, 1)

	dec.Read(buf)
	return Bool(buf[0] > 0)
	// return dec.DecodeByte() > 0
}

// func (enc Encoder) EncodeBool(value bool) {
// 	if value {
// 		enc.EncodeByte(0x01)
// 	} else {
// 		enc.EncodeByte(0x00)
// 	}
// }

// func (dec Decoder) DecodeBool() bool {
// 	buf := make([]byte, 1)
// 	dec.Read(buf)
// 	return buf[0] > 0
// 	// return dec.DecodeByte() > 0
// }
