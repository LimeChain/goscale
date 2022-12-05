package goscale

import (
	"bytes"
	"math"
	"testing"
)

type testEncodable struct {
}

func (testEncodable) Encode(*bytes.Buffer) {
}
func (testEncodable) String() string {
	return ""
}

func Test_EncodeOption(t *testing.T) {
	var examples = []struct {
		label  string
		input  Option[Encodable]
		expect []byte
	}{
		{label: "Encode Option(true, false)", input: Option[Encodable]{true, Bool(false)}, expect: []byte{0x1, 0x0}},
		{label: "Encode Option(true, true)", input: Option[Encodable]{true, Bool(true)}, expect: []byte{0x1, 0x1}},
		{label: "Encode Option(false, nil)", input: Option[Encodable]{false, nil}, expect: []byte{0x0}},
		{label: "Encode Option(false, true)", input: Option[Encodable]{false, Bool(true)}, expect: []byte{0x0}},

		{label: "Encode Option(true, U8(max))", input: Option[Encodable]{true, U8(math.MaxUint8)}, expect: []byte{0x1, 0xff}},
		{label: "Encode Option(true, I8(min))", input: Option[Encodable]{true, I8(math.MinInt8)}, expect: []byte{0x1, 0x80}},
		{label: "Encode Option(true, I8(max))", input: Option[Encodable]{true, I8(math.MaxInt8)}, expect: []byte{0x1, 0x7f}},
		{label: "Encode Option(true, U16(max))", input: Option[Encodable]{true, U16(math.MaxUint16)}, expect: []byte{0x1, 0xff, 0xff}},
		{label: "Encode Option(true, I16(min))", input: Option[Encodable]{true, I16(math.MinInt16)}, expect: []byte{0x1, 0x00, 0x80}},
		{label: "Encode Option(true, U32(max))", input: Option[Encodable]{true, U32(math.MaxUint32)}, expect: []byte{0x1, 0xff, 0xff, 0xff, 0xff}},
		{label: "Encode Option(true, I32(min))", input: Option[Encodable]{true, I32(math.MinInt32)}, expect: []byte{0x1, 0x0, 0x0, 0x0, 0x80}},
		{label: "Encode Option(true, U64(max))", input: Option[Encodable]{true, U64(math.MaxUint64)}, expect: []byte{0x1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{label: "Encode Option(true, I64(min))", input: Option[Encodable]{true, I64(math.MinInt64)}, expect: []byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80}},
		{label: "Encode Option(false, I64(min))", input: Option[Encodable]{false, I64(math.MinInt64)}, expect: []byte{0x0}},
		{label: "Encode Option(true, U128(max)", input: Option[Encodable]{true, U128{math.MaxUint64, math.MaxUint64}}, expect: []byte{0x1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},

		{label: "Encode Option(true, Compact(MaxUint64)", input: Option[Encodable]{true, Compact(math.MaxUint64)}, expect: []byte{0x1, 0x13, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},

		{label: "Encode Option(true, empty Seq[U8])", input: Option[Encodable]{true, Sequence[U8]{}}, expect: []byte{0x1, 0x0}},
		{label: "Encode Option(true, Seq[U8])", input: Option[Encodable]{true, Sequence[U8]{[]U8{42}}}, expect: []byte{0x1, 0x4, 0x2a}},

		{label: "Encode Option(true, Result(true, Seq[U8])", input: Option[Encodable]{true, Result[Encodable]{true, Sequence[U8]{[]U8{42}}}}, expect: []byte{0x1, 0x1, 0x4, 0x2a}},
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

func Test_EncodeOptionReverts(t *testing.T) {
	var examples = []struct {
		label string
		input Option[Encodable]
	}{
		{label: "Revert Option(true, nil)", input: Option[Encodable]{true, nil}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}

			// then:
			assertPanic(t, func() {
				e.input.Encode(buffer)
			})
		})
	}
}

func Test_DecodeOption(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		encodable     Encodable
		bufferLenLeft int
		expect        Option[Encodable]
	}{
		{
			label:         "Decode Option(false, nil)",
			input:         []byte{0x0},
			encodable:     Bool(false),
			bufferLenLeft: 0,
			expect:        Option[Encodable]{false, nil},
		},
		{
			label:         "Decode Option(true,false)",
			input:         []byte{0x1, 0x0},
			encodable:     Bool(false),
			bufferLenLeft: 0,
			expect:        Option[Encodable]{true, Bool(false)},
		},
		{
			label:         "Decode Option(true,true)",
			input:         []byte{0x1, 0x1, 0x3},
			encodable:     Bool(false),
			bufferLenLeft: 1,
			expect:        Option[Encodable]{true, Bool(true)},
		},
		{
			label:         "Decode Option(true, U8(max))",
			input:         []byte{0x1, 0xff, 0xff},
			encodable:     U8(0),
			expect:        Option[Encodable]{true, U8(math.MaxUint8)},
			bufferLenLeft: 1,
		},
		{
			label:         "Decode Option(true, I8(min))",
			input:         []byte{0x1, 0x80},
			encodable:     I8(0),
			expect:        Option[Encodable]{true, I8(math.MinInt8)},
			bufferLenLeft: 0,
		},
		{
			label:         "Decode Option(true, U16(max))",
			input:         []byte{0x1, 0xff, 0xff},
			encodable:     U16(0),
			expect:        Option[Encodable]{true, U16(math.MaxUint16)},
			bufferLenLeft: 0,
		},
		{
			label:         "Decode Option(true, I16(min))",
			input:         []byte{0x1, 0x0, 0x80},
			encodable:     I16(0),
			expect:        Option[Encodable]{true, I16(math.MinInt16)},
			bufferLenLeft: 0,
		},
		{
			label:         "Decode Option(true, U32(max))",
			input:         []byte{0x1, 0xff, 0xff, 0xff, 0xff},
			encodable:     U32(0),
			expect:        Option[Encodable]{true, U32(math.MaxUint32)},
			bufferLenLeft: 0,
		},
		{
			label:         "Decode Option(true, I32(min))",
			input:         []byte{0x1, 0x0, 0x0, 0x0, 0x80},
			encodable:     I32(0),
			expect:        Option[Encodable]{true, I32(math.MinInt32)},
			bufferLenLeft: 0,
		},
		{
			label:         "Decode Option(true, U64(max))",
			input:         []byte{0x1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			encodable:     U64(0),
			expect:        Option[Encodable]{true, U64(math.MaxUint64)},
			bufferLenLeft: 0,
		},
		{
			label:         "Decode Option(true, I64(min))",
			input:         []byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80},
			encodable:     I64(0),
			expect:        Option[Encodable]{true, I64(math.MinInt64)},
			bufferLenLeft: 0,
		},
		{
			label:         "Decode Option(true, U128(max))",
			input:         []byte{0x1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			encodable:     U128{},
			expect:        Option[Encodable]{true, U128{math.MaxUint64, math.MaxUint64}},
			bufferLenLeft: 0,
		},
		{
			label:         "Decode Compact(maxUint64)",
			input:         []byte{0x1, 0x13, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			encodable:     Compact(0),
			expect:        Option[Encodable]{HasValue: true, Value: Compact(math.MaxUint64)},
			bufferLenLeft: 0,
		},
		{
			label:         "Decode Seq[U8]",
			input:         []byte{0x1, 0x4, 0x2a},
			encodable:     Sequence[U8]{},
			expect:        Option[Encodable]{HasValue: true, Value: Sequence[U8]{[]U8{42}}},
			bufferLenLeft: 0,
		},
		// TODO: Decode Option<Result<true, any>>
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeOption[Encodable](e.encodable, buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
		})
	}
}

func Test_DecodeOptionReverts(t *testing.T) {
	var examples = []struct {
		label     string
		encodable Encodable
		input     []byte
	}{
		{label: "Revert Option(0x1)", encodable: Bool(false), input: []byte{0x1}},
		{label: "Revert Option(empty slice)", encodable: Bool(false), input: []byte{}},
		{label: "Revert Option(nil)", encodable: Bool(false), input: nil},
		{label: "Revert Option(different type)", encodable: testEncodable{}, input: []byte{0x1}},
		{label: "Revert Option(U16) - cannot read bytes, which are not found)", encodable: U16(0), input: []byte{0x1, 0x5}},
		{label: "Revert Option(U32) - cannot read bytes, which are not found)", encodable: U32(0), input: []byte{0x1, 0x5}},
		{label: "Revert Option(U64) - cannot read bytes, which are not found)", encodable: U64(0), input: []byte{0x1, 0x5}},
		{label: "Revert Option(I16) - cannot read bytes, which are not found)", encodable: I16(0), input: []byte{0x1, 0x5}},
		{label: "Revert Option(I32) - cannot read bytes, which are not found)", encodable: I32(0), input: []byte{0x1, 0x5}},
		{label: "Revert Option(I64) - cannot read bytes, which are not found)", encodable: I64(0), input: []byte{0x1, 0x5}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// then:
			assertPanic(t, func() {
				DecodeOption[Encodable](e.encodable, buffer)
			})
		})
	}
}
