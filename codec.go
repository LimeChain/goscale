/*
Simple Concatenated Aggregate Little-Endian” (SCALE) codec

Polkadot Spec - https://spec.polkadot.network/#sect-scale-codec

Substrate Ref - https://docs.substrate.io/reference/scale-codec/
*/
package goscale

import (
	"errors"
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
		panic("can not write the provided " + strconv.Itoa(len(bytes)) + " bytes to writer")
	}
}

func (dec Decoder) Read(bytes []byte) error {
	n, err := dec.Reader.Read(bytes)
	if err != nil {
		return err
	}
	if n < len(bytes) {
		return errors.New("can not read the required number of bytes " + strconv.Itoa(len(bytes)) + ", only " + strconv.Itoa(n) + " available")
	}
	return nil
}

func (enc Encoder) EncodeByte(b byte) {
	buf := make([]byte, 1)
	buf[0] = b
	enc.Write(buf[:1])
}

func (dec Decoder) DecodeByte() (byte, error) {
	buf := make([]byte, 1)
	err := dec.Read(buf[:1])
	if err != nil {
		return 0, err
	}
	return buf[0], nil
}
