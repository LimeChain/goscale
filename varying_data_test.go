package goscale

import (
	"bytes"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
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
			buffer := &bytes.Buffer{}

			e.input.Encode(buffer)

			assert.Equal(t, buffer.Bytes(), e.expect)
			assert.Equal(t, e.input.Bytes(), e.expect)
		})
	}
}

func Test_NewVaryingData_InvalidLength(t *testing.T) {
	values := make([]Encodable, math.MaxUint8+1)

	assert.Panics(t, func() {
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
					resultDecode, _ := DecodeU8(buffer)
					return []Encodable{resultDecode}
				},
				func(buffer *bytes.Buffer) []Encodable {
					resultDecode, _ := DecodeBool(buffer)
					return []Encodable{resultDecode}
				},
			},
			expect: NewVaryingData(U8(0), U8(42)),
		},
		{
			label: "Decode VaryingData(U128, Empty)",
			input: []byte{0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			decodeFuncs: []func(buffer *bytes.Buffer) []Encodable{
				func(buffer *bytes.Buffer) []Encodable {
					resultDecode, _ := DecodeU128(buffer)
					return []Encodable{resultDecode}
				},
			},
			expect: NewVaryingData(U8(0), U128{math.MaxUint64, math.MaxUint64}),
		},
		{
			label: "Decode VaryingData(U64,U32,Sequence[U8])",
			input: []byte{0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x4, 0x2a},
			decodeFuncs: []func(buffer *bytes.Buffer) []Encodable{
				func(buffer *bytes.Buffer) []Encodable {
					resultDecodeU64, _ := DecodeU64(buffer)
					resultDecodeU32, _ := DecodeU32(buffer)
					resultDecodeSeqU8, _ := DecodeSequence[U8](buffer)
					return []Encodable{resultDecodeU64, resultDecodeU32, resultDecodeSeqU8}
				},
			},
			expect: NewVaryingData(U8(0), U64(math.MaxUint64), U32(math.MaxUint32), Sequence[U8]{42}),
		},
		{
			label: "Decode VaryingData(I8,U16,I16,CompactUint,CompactUint,I32,I64)",
			input: []byte{0x0, 0x80, 0xff, 0xff, 0x00, 0x80, 0x0b, 0x00, 0x40, 0x7a, 0x10, 0xf3, 0x5a, 0x14, 0x0, 0x0, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80},
			decodeFuncs: []func(buffer *bytes.Buffer) []Encodable{
				func(buffer *bytes.Buffer) []Encodable {
					resultDecodeI8, _ := DecodeI8(buffer)
					resultDecodeU16, _ := DecodeU16(buffer)
					resultDecodeI16, _ := DecodeI16(buffer)
					resultDecodeCompactOne, _ := DecodeCompact(buffer)
					resultDecodeCompactTwo, _ := DecodeCompact(buffer)
					resultDecodeI32, _ := DecodeI32(buffer)
					resultDecodeI64, _ := DecodeI64(buffer)
					return []Encodable{resultDecodeI8, resultDecodeU16, resultDecodeI16, resultDecodeCompactOne, resultDecodeCompactTwo, resultDecodeI32, resultDecodeI64}
				},
			},
			expect: NewVaryingData(U8(0), I8(math.MinInt8), U16(math.MaxUint16), I16(math.MinInt16), ToCompact(100000000000000), ToCompact(5), I32(math.MinInt32), I64(math.MinInt64)),
		},
	}

	for _, e := range examples {
		buffer := &bytes.Buffer{}
		buffer.Write(e.input)

		result, _ := DecodeVaryingData(e.decodeFuncs, buffer)

		assert.Equal(t, result, e.expect)
	}
}

func Test_VaryingData_Decode_Error_ExceedsLength(t *testing.T) {
	values := make([]func(buffer *bytes.Buffer) []Encodable, math.MaxUint8+1)

	_, err := DecodeVaryingData(values, &bytes.Buffer{})
	assert.ErrorIs(t, err, errExceedsU8Length)
}

func Test_VaryingData_Decode_Error_Index_NotFound(t *testing.T) {
	values := make([]func(buffer *bytes.Buffer) []Encodable, 1)

	buffer := &bytes.Buffer{}
	buffer.Write(U8(1).Bytes())

	_, err := DecodeVaryingData(values, buffer)

	assert.ErrorIs(t, err, errDecodingFuncNotFound)
}
