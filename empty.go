package goscale

/*
	Ref: https://spec.polkadot.network/#defn-scale-empty

	SCALE Empty type.
	Values are encoded as byte array of zero length.
*/

import "bytes"

type Empty struct{}

func (e Empty) Encode(buffer *bytes.Buffer) {
	encoder := Encoder{Writer: buffer}

	encoder.Write(e.Bytes())
}

func (e Empty) Bytes() []byte {
	return []byte{}
}

func (e Empty) String() string {
	return ""
}

func DecodeEmpty() Empty {
	return Empty{}
}
