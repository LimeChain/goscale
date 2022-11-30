package goscale

import (
	"bytes"
	"testing"
)

func Test_EncodeVaryingData(t *testing.T) {
	var examples = []struct {
		label  string
		input  VaryingData
		expect []byte
	}{
		{label: "Encode VaryingData(42, true)", input: NewVaryingData(U8(42), Bool(true)), expect: []byte{0x0, 0x2a, 0x1, 0x1}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}

			// when:
			e.input.Encode(buffer)

			// then:
			assertEqual(t, buffer.Bytes(), e.expect)
		})
	}
}

func Test_NewVaryingData_InvalidLength(t *testing.T) {
	// given:
	values := make([]Encodable, 256)

	// then:
	assertPanic(t, func() {
		NewVaryingData(values...)
	})
}

func Test_VaryingData_Add_InvalidLength(t *testing.T) {
	// given:
	values := make([]Encodable, 256)
	varyingData := NewVaryingData()

	// then:
	assertPanic(t, func() {
		varyingData.Add(values...)
	})
}

func Test_VaryingData_Add(t *testing.T) {
	// given:
	values := []Encodable{Bool(true), U8(42)}
	varyingData := NewVaryingData()

	// when:
	varyingData.Add(values...)

	// then:
	for i, v := range values {
		assertEqual(t, varyingData.Values[U8(i)], v)
	}
}
