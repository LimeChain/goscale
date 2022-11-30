package goscale

import (
	"bytes"
)

type VaryingData struct {
	Values map[U8]Encodable
}

func NewVaryingData(values ...Encodable) VaryingData {
	if len(values) > 255 {
		panic("invalid length")
	}
	i := 0
	result := map[U8]Encodable{}

	for _, v := range values {
		result[U8(i)] = v
		i++
	}

	return VaryingData{
		Values: result,
	}
}

func (vd *VaryingData) Add(values ...Encodable) {
	l := len(vd.Values)
	if l+len(values) > 255 {
		panic("exceeds uint8 length")
	}

	for _, v := range values {
		vd.Values[U8(l)] = v
		l++
	}
}

func (vd *VaryingData) Encode(buffer *bytes.Buffer) {
	for k, v := range vd.Values {
		k.Encode(buffer)
		v.Encode(buffer)
	}
}
