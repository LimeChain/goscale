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
			input:  NewVaryingData(U8(0), U8(42)),
			expect: []byte{0x0, 0x2a}},
		{
			label:  "Encode VaryingData(U128, Empty)",
			input:  NewVaryingData(U8(0), U128{math.MaxUint64, math.MaxUint64}, Empty{}),
			expect: []byte{0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{
			label:  "Encode VaryingData(U64,U32,Sequence[U8])",
			input:  NewVaryingData(U8(0), U64(math.MaxUint64), U32(math.MaxUint32), Sequence[U8]{42}),
			expect: []byte{0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x4, 0x2a}},
		{
			label:  "Encode VaryingData(I8,U16,I16,CompactUint,CompactUint,I32,I64)",
			input:  NewVaryingData(U8(0), I8(math.MinInt8), U16(math.MaxUint16), I16(math.MinInt16), ToCompact(100000000000000), ToCompact(5), I32(math.MinInt32), I64(math.MinInt64)),
			expect: []byte{0x0, 0x80, 0xff, 0xff, 0x00, 0x80, 0x0b, 0x00, 0x40, 0x7a, 0x10, 0xf3, 0x5a, 0x14, 0x0, 0x0, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80},
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
		label       string
		input       []byte
		decodeFuncs []func(buffer *bytes.Buffer) []Encodable
		expect      VaryingData
	}{
		{
			label: "Decode VaryingData(U8, Bool)",
			input: []byte{0x0, 0x2a},
			decodeFuncs: []func(buffer *bytes.Buffer) []Encodable{
				func(buffer *bytes.Buffer) []Encodable {
					return []Encodable{DecodeU8(buffer)}
				},
				func(buffer *bytes.Buffer) []Encodable {
					return []Encodable{DecodeBool(buffer)}
				},
			},
			expect: NewVaryingData(U8(0), U8(42)),
		},
		{
			label: "Decode VaryingData(U128, Empty)",
			input: []byte{0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			decodeFuncs: []func(buffer *bytes.Buffer) []Encodable{
				func(buffer *bytes.Buffer) []Encodable {
					return []Encodable{DecodeU128(buffer)}
				},
			},
			expect: NewVaryingData(U8(0), U128{math.MaxUint64, math.MaxUint64}),
		},
		{
			label: "Decode VaryingData(U64,U32,Sequence[U8])",
			input: []byte{0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x4, 0x2a},
			decodeFuncs: []func(buffer *bytes.Buffer) []Encodable{
				func(buffer *bytes.Buffer) []Encodable {
					return []Encodable{DecodeU64(buffer), DecodeU32(buffer), DecodeSequence[U8](buffer)}
				},
			},
			expect: NewVaryingData(U8(0), U64(math.MaxUint64), U32(math.MaxUint32), Sequence[U8]{42}),
		},
		{
			label: "Decode VaryingData(I8,U16,I16,CompactUint,CompactUint,I32,I64)",
			input: []byte{0x0, 0x80, 0xff, 0xff, 0x00, 0x80, 0x0b, 0x00, 0x40, 0x7a, 0x10, 0xf3, 0x5a, 0x14, 0x0, 0x0, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80},
			decodeFuncs: []func(buffer *bytes.Buffer) []Encodable{
				func(buffer *bytes.Buffer) []Encodable {
					return []Encodable{DecodeI8(buffer), DecodeU16(buffer), DecodeI16(buffer), DecodeCompact(buffer), DecodeCompact(buffer), DecodeI32(buffer), DecodeI64(buffer)}
				},
			},
			expect: NewVaryingData(U8(0), I8(math.MinInt8), U16(math.MaxUint16), I16(math.MinInt16), ToCompact(100000000000000), ToCompact(5), I32(math.MinInt32), I64(math.MinInt64)),
		},
	}

	for _, e := range examples {
		// given:
		buffer := &bytes.Buffer{}
		buffer.Write(e.input)

		// when:
		result := DecodeVaryingData(e.decodeFuncs, buffer)

		// then:
		assertEqual(t, result, e.expect)
	}
}

func Test_VaryingData_Decode_Panic_ExceedsLength(t *testing.T) {
	// given:
	values := make([]func(buffer *bytes.Buffer) []Encodable, math.MaxUint8+1)

	// then:
	assertPanic(t, func() {
		DecodeVaryingData(values, &bytes.Buffer{})
	})
}

func Test_VaryingData_Decode_Panic_Index_NotFound(t *testing.T) {
	// given:
	values := make([]func(buffer *bytes.Buffer) []Encodable, 1)

	buffer := &bytes.Buffer{}
	buffer.Write(U8(1).Bytes())

	// then:
	assertPanic(t, func() {
		DecodeVaryingData(values, buffer)
	})
}
