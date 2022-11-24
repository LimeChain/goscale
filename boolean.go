package goscale

import (
	"bytes"
	"fmt"
)

/*
	https://spec.polkadot.network/#defn-scale-boolean

	SCALE's Boolean type translates to Go's boolean type.
	Values are encoded using the least significant bit of a single byte.
*/

type Bool bool

func (value Bool) Encode(buffer *bytes.Buffer) {
	encoder := Encoder{Writer: buffer}
	if value {
		encoder.EncodeByte(0x01)
	} else {
		encoder.EncodeByte(0x00)
	}
}

func DecodeBool(buffer *bytes.Buffer) Bool {
	decoder := Decoder{Reader: buffer}
	result := make([]byte, 1)
	decoder.Read(result)
	return Bool(result[0] > 0)
}

func (value Bool) String() string {
	return fmt.Sprint(bool(value))
}
