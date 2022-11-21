/*
	Simple Concatenated Aggregate Little-Endian” (SCALE) codec

	Polkadot Spec - https://spec.polkadot.network/#sect-scale-codec

	Substrate Ref - https://docs.substrate.io/reference/scale-codec/
*/
package goscale

import (
	"io"
	"strconv"
)

type Encoder struct {
	Writer io.Writer
}

type Decoder struct {
	Reader io.Reader
}

func (enc Encoder) Write(bytes []byte) {
	n, err := enc.Writer.Write(bytes)
	if err != nil {
		panic(err.Error())
	}
	if n < len(bytes) {
		panic("Can not write the provided " + strconv.Itoa(len(bytes)) + " bytes to writer")
	}
}

func (dec Decoder) Read(bytes []byte) {
	n, err := dec.Reader.Read(bytes)
	if err != nil {
		panic(err.Error())
	}
	if n < len(bytes) {
		panic("Can not read the required number of bytes " + strconv.Itoa(len(bytes)) + ", only " + strconv.Itoa(n) + " available")
	}
}

func (enc Encoder) EncodeByte(b byte) {
	buf := make([]byte, 1)
	buf[0] = b
	enc.Write(buf[:1])
}

func (dec Decoder) DecodeByte() byte {
	buf := make([]byte, 1)
	dec.Read(buf[:1])
	return buf[0]
}

type Encodable interface {
	Encode(enc *Encoder) // TODO return an error
}
