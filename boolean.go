package goscale

/*
	Ref: https://spec.polkadot.network/#defn-scale-boolean

	SCALE Boolean type translates to Go's boolean type.
	Values are encoded using the least significant bit of a single byte.
*/

import (
	"bytes"
	"fmt"
)

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
	result := decoder.DecodeByte()
	switch result {
	case 0:
		return false
	case 1:
		return true
	default:
		panic("invalid bool representation")
	}
}

func (value Bool) String() string {
	return fmt.Sprint(bool(value))
}
