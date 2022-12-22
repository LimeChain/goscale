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
		{
			label:  "Encode VaryingData(U8, Bool)",
			input:  NewVaryingData(U8(42), Bool(true)),
			expect: []byte{0x0, 0x2a, 0x1, 0x1}},
		{
			label:  "Encode VaryingData(U128, Empty)",
			input:  NewVaryingData(U128{math.MaxUint64, math.MaxUint64}, Empty{}),
			expect: []byte{0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1}},
		{
			label:  "Encode VaryingData(U64,U32,Sequence[U8])",
			input:  NewVaryingData(U64(math.MaxUint64), U32(math.MaxUint32), Sequence[U8]{42}),
			expect: []byte{0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1, 0xff, 0xff, 0xff, 0xff, 0x2, 0x4, 0x2a}},
		{
			label:  "Encode VaryingData(I8,U16,I16,CompactUint,CompactUint,I32,I64)",
			input:  NewVaryingData(I8(math.MinInt8), U16(math.MaxUint16), I16(math.MinInt16), toCompact(100000000000000), toCompact(5), I32(math.MinInt32), I64(math.MinInt64)),
			expect: []byte{0x0, 0x80, 0x1, 0xff, 0xff, 0x2, 0x00, 0x80, 0x3, 0x0b, 0x00, 0x40, 0x7a, 0x10, 0xf3, 0x5a, 0x4, 0x14, 0x5, 0x0, 0x0, 0x0, 0x80, 0x6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80},
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}

			// when:
			e.input.Encode(buffer)

			// then:
			assertEqual(t, buffer.Bytes(), e.expect)
			// and:
			assertEqual(t, e.input.Bytes(), e.expect)
		})
	}
}

func Test_NewVaryingData_InvalidLength(t *testing.T) {
	// given:
	values := make([]Encodable, math.MaxUint8+1)

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
			label:  "Decode VaryingData(U128, Empty)",
			input:  []byte{0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1},
			order:  []Encodable{U128{0, 0}, Empty{}},
			expect: NewVaryingData(U128{math.MaxUint64, math.MaxUint64}, Empty{}),
		},
		{
			label:  "Decode VaryingData(U64,U32,Sequence[U8])",
			input:  []byte{0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1, 0xff, 0xff, 0xff, 0xff, 0x2, 0x4, 0x2a},
			order:  []Encodable{U64(0), U32(0), Sequence[U8]{}},
			expect: NewVaryingData(U64(math.MaxUint64), U32(math.MaxUint32), Sequence[U8]{42}),
		},
		{
			label:  "Decode VaryingData(I8,U16,I16,CompactUint,CompactUint,I32,I64)",
			input:  []byte{0x0, 0x80, 0x1, 0xff, 0xff, 0x2, 0x00, 0x80, 0x3, 0x0b, 0x00, 0x40, 0x7a, 0x10, 0xf3, 0x5a, 0x4, 0x14, 0x5, 0x0, 0x0, 0x0, 0x80, 0x6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80},
			order:  []Encodable{I8(0), U16(0), I16(0), toCompact(0), toCompact(0), I32(0), I64(0)},
			expect: NewVaryingData(I8(math.MinInt8), U16(math.MaxUint16), I16(math.MinInt16), toCompact(100000000000000), toCompact(5), I32(math.MinInt32), I64(math.MinInt64)),
		},
		{
			label:  "Decode VaryingData(U8, Bool, Compact, I8) with mixed bytes order",
			input:  []byte{0x1, 0x1, 0x3, 0x80, 0x0, 0x5, 0x2, 0x0b, 0x00, 0x40, 0x7a, 0x10, 0xf3, 0x5a},
			order:  []Encodable{U8(0), Bool(false), toCompact(0), I8(0)},
			expect: NewVaryingData(U8(5), Bool(true), toCompact(100000000000000), I8(math.MinInt8)),
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

func Test_VaryingData_Decode_Panic(t *testing.T) {
	// given:
	values := make([]Encodable, math.MaxUint8+1)

	// then:
	assertPanic(t, func() {
		DecodeVaryingData(values, &bytes.Buffer{})
	})
}
