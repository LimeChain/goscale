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
	encoder.Write(value.Bytes())
}

func (value Bool) Bytes() []byte {
	buf := make([]byte, 1)
	if value {
		buf[0] = 1
	}

	return buf
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

func (value Bool) IsTrue() bool {
	return value == Bool(true)
}

func (value Bool) String() string {
	return fmt.Sprint(bool(value))
}
