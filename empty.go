package goscale

/*
	Ref: https://spec.polkadot.network/#defn-scale-empty

	SCALE Empty type.
	Values are encoded as byte array of zero length.
*/

import "bytes"

type Empty struct{}

func (e Empty) Encode(buffer *bytes.Buffer) error {
	encoder := Encoder{Writer: buffer}
	return encoder.Write(e.Bytes())
}

func (e Empty) Bytes() []byte {
	return []byte{}
}

func DecodeEmpty() (Empty, error) {
	return Empty{}, nil
}
