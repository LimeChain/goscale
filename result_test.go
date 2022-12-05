package goscale

import (
	"bytes"
	"math"
	"testing"
)

func Test_EncodeResult(t *testing.T) {
	var examples = []struct {
		label  string
		input  Result[Encodable]
		expect []byte
	}{
		{label: "Encode Result(true, false)", input: Result[Encodable]{true, Bool(false)}, expect: []byte{0x1, 0x0}},
		{label: "Encode Result(true, true)", input: Result[Encodable]{true, Bool(true)}, expect: []byte{0x1, 0x1}},
		{label: "Encode Result(false, empty)", input: Result[Encodable]{false, Empty{}}, expect: []byte{0x0}},
		{label: "Encode Result(false, true)", input: Result[Encodable]{false, Bool(true)}, expect: []byte{0x0, 0x1}},

		{label: "Encode Result(true, U8(max))", input: Result[Encodable]{true, U8(math.MaxUint8)}, expect: []byte{0x1, 0xff}},
		{label: "Encode Result(true, I8(min))", input: Result[Encodable]{true, I8(math.MinInt8)}, expect: []byte{0x1, 0x80}},
		{label: "Encode Result(true, I8(max))", input: Result[Encodable]{true, I8(math.MaxInt8)}, expect: []byte{0x1, 0x7f}},
		{label: "Encode Result(true, U16(max))", input: Result[Encodable]{true, U16(math.MaxUint16)}, expect: []byte{0x1, 0xff, 0xff}},
		{label: "Encode Result(true, I16(min))", input: Result[Encodable]{true, I16(math.MinInt16)}, expect: []byte{0x1, 0x00, 0x80}},
		{label: "Encode Result(true, U32(max))", input: Result[Encodable]{true, U32(math.MaxUint32)}, expect: []byte{0x1, 0xff, 0xff, 0xff, 0xff}},
		{label: "Encode Result(true, I32(min))", input: Result[Encodable]{true, I32(math.MinInt32)}, expect: []byte{0x1, 0x0, 0x0, 0x0, 0x80}},
		{label: "Encode Result(true, U64(max))", input: Result[Encodable]{true, U64(math.MaxUint64)}, expect: []byte{0x1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{label: "Encode Result(true, I64(min))", input: Result[Encodable]{true, I64(math.MinInt64)}, expect: []byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80}},
		{label: "Encode Result(false, I64(min))", input: Result[Encodable]{false, I64(math.MinInt64)}, expect: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80}},
		{label: "Encode Result(true, U128(max))", input: Result[Encodable]{true, U128{math.MaxUint64, math.MaxUint64}}, expect: []byte{0x1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},

		{label: "Encode Result(true, Compact(MaxUint64)", input: Result[Encodable]{true, Compact(math.MaxUint64)}, expect: []byte{0x01, 0x13, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},

		{label: "Encode Result(true, empty Seq[U8])", input: Result[Encodable]{true, Sequence[U8]{}}, expect: []byte{0x1, 0x0}},
		{label: "Encode Result(true, Seq[U8])", input: Result[Encodable]{true, Sequence[U8]{[]U8{42}}}, expect: []byte{0x1, 0x4, 0x2a}},
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

func Test_DecodeResult(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		encodable     Encodable
		bufferLenLeft int
		expect        Result[Encodable]
	}{
		{
			label:         "Decode Result(false, empty)",
			input:         []byte{0x0},
			encodable:     Empty{},
			bufferLenLeft: 0,
			expect:        Result[Encodable]{false, Empty{}},
		},
		{
			label:         "Decode Result(false, false)",
			input:         []byte{0x0, 0x0},
			encodable:     Bool(false),
			bufferLenLeft: 0,
			expect:        Result[Encodable]{false, Bool(false)},
		},
		{
			label:         "Decode Result(true,false)",
			input:         []byte{0x1, 0x0},
			encodable:     Bool(false),
			bufferLenLeft: 0,
			expect:        Result[Encodable]{true, Bool(false)},
		},
		{
			label:         "Decode Result(true,true)",
			input:         []byte{0x1, 0x1, 0x3},
			encodable:     Bool(false),
			bufferLenLeft: 1,
			expect:        Result[Encodable]{true, Bool(true)},
		},
		{
			label:         "Decode Result(true, U8(max))",
			input:         []byte{0x1, 0xff, 0xff},
			encodable:     U8(0),
			expect:        Result[Encodable]{true, U8(math.MaxUint8)},
			bufferLenLeft: 1,
		},
		{
			label:         "Decode Result(true, I8(min))",
			input:         []byte{0x1, 0x80},
			encodable:     I8(0),
			expect:        Result[Encodable]{true, I8(math.MinInt8)},
			bufferLenLeft: 0,
		},
		{
			label:         "Decode Result(true, U16(max))",
			input:         []byte{0x1, 0xff, 0xff},
			encodable:     U16(0),
			expect:        Result[Encodable]{true, U16(math.MaxUint16)},
			bufferLenLeft: 0,
		},
		{
			label:         "Decode Result(true, I16(min))",
			input:         []byte{0x1, 0x0, 0x80},
			encodable:     I16(0),
			expect:        Result[Encodable]{true, I16(math.MinInt16)},
			bufferLenLeft: 0,
		},
		{
			label:         "Decode Result(true, U32(max))",
			input:         []byte{0x1, 0xff, 0xff, 0xff, 0xff},
			encodable:     U32(0),
			expect:        Result[Encodable]{true, U32(math.MaxUint32)},
			bufferLenLeft: 0,
		},
		{
			label:         "Decode Result(true, I32(min))",
			input:         []byte{0x1, 0x0, 0x0, 0x0, 0x80},
			encodable:     I32(0),
			expect:        Result[Encodable]{true, I32(math.MinInt32)},
			bufferLenLeft: 0,
		},
		{
			label:         "Decode Result(true, U64(max))",
			input:         []byte{0x1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			encodable:     U64(0),
			expect:        Result[Encodable]{true, U64(math.MaxUint64)},
			bufferLenLeft: 0,
		},
		{
			label:         "Decode Result(true, I64(min))",
			input:         []byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80},
			encodable:     I64(0),
			expect:        Result[Encodable]{true, I64(math.MinInt64)},
			bufferLenLeft: 0,
		},
		{
			label:         "Decode Compact(maxUint64)",
			input:         []byte{0x1, 0x13, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			encodable:     Compact(0),
			expect:        Result[Encodable]{true, Compact(math.MaxUint64)},
			bufferLenLeft: 0,
		},
		{
			label:         "Decode Seq[U8]",
			input:         []byte{0x1, 0x4, 0x2a},
			encodable:     Sequence[U8]{},
			expect:        Result[Encodable]{true, Sequence[U8]{[]U8{42}}},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeResult[Encodable](e.encodable, buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
		})
	}
}
