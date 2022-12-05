package goscale

import (
	"bytes"
	"math"
	"testing"
)

func Test_VaryingData_Encode(t *testing.T) {
	var examples = []struct {
		label  string
		input  VaryingData
		expect []byte
	}{
		{label: "Encode VaryingData(U8, Bool)", input: NewVaryingData(U8(42), Bool(true)), expect: []byte{0x0, 0x2a, 0x1, 0x1}},
		{label: "Encode VaryingData(U128, Empty)", input: NewVaryingData(U128{math.MaxUint64, math.MaxUint64}, Empty{}), expect: []byte{0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1}},
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

func Test_VaryingData_Decode(t *testing.T) {
	var examples = []struct {
		label  string
		input  []byte
		order  []Encodable
		expect VaryingData
	}{
		{
			label:  "Decode VaryingData(U8, Bool)",
			input:  []byte{0x0, 0x2a, 0x1, 0x1},
			order:  []Encodable{U8(1), Bool(false)},
			expect: NewVaryingData(U8(42), Bool(true))},
		{
			label:  "DecodeVaryingData(U128, Empty)",
			input:  []byte{0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1},
			order:  []Encodable{U128{0, 0}, Empty{}},
			expect: NewVaryingData(U128{math.MaxUint64, math.MaxUint64}, Empty{}),
		},
	}

	for _, e := range examples {
		// given:
		buffer := &bytes.Buffer{}
		buffer.Write(e.input)

		// when:
		result := DecodeVaryingData(e.order, buffer)

		// then:
		assertEqual(t, result, e.expect)
	}
}

func Test_VaryingData_Panic(t *testing.T) {
	// given:
	values := make([]Encodable, 256)

	// then:
	assertPanic(t, func() {
		DecodeVaryingData(values, &bytes.Buffer{})
	})
}
