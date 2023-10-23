package goscale

/*
	Ref: https://spec.polkadot.network/#defn-scale-boolean

	SCALE Boolean type translates to Go's boolean type.
	Values are encoded using the least significant bit of a single byte.
*/

import (
	"bytes"
	"errors"
)

type Bool bool

var (
	errInvalidBoolRepresentation = errors.New("invalid bool representation")
)

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

func DecodeBool(buffer *bytes.Buffer) (Bool, error) {
	decoder := Decoder{Reader: buffer}
	result := decoder.DecodeByte()
	switch result {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		return false, errInvalidBoolRepresentation
	}
}
