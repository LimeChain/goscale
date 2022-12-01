package goscale

import "bytes"

type Empty struct {
}

func (e Empty) Encode(buffer *bytes.Buffer) {
	encoder := Encoder{Writer: buffer}

	encoder.Write([]byte{})
}

func (e Empty) String() string {
	return ""
}

func DecodeEmpty() Empty {
	return Empty{}
}
