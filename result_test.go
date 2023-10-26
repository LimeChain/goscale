package goscale

import (
	"bytes"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EncodeResult(t *testing.T) {
	var examples = []struct {
		label  string
		input  Result[Encodable]
		expect []byte
	}{
		{label: "Encode Result(false, false)", input: Result[Encodable]{HasError: false, Value: Bool(false)}, expect: []byte{0x0, 0x0}},
		{label: "Encode Result(false, true)", input: Result[Encodable]{HasError: false, Value: Bool(true)}, expect: []byte{0x0, 0x1}},
		{label: "Encode Result(true, empty)", input: Result[Encodable]{HasError: true, Value: Empty{}}, expect: []byte{0x1}},
		{label: "Encode Result(true, true)", input: Result[Encodable]{HasError: true, Value: Bool(true)}, expect: []byte{0x1, 0x1}},

		{label: "Encode Result(false, U8(max))", input: Result[Encodable]{HasError: false, Value: U8(math.MaxUint8)}, expect: []byte{0x0, 0xff}},
		{label: "Encode Result(false, I8(min))", input: Result[Encodable]{HasError: false, Value: I8(math.MinInt8)}, expect: []byte{0x0, 0x80}},
		{label: "Encode Result(false, I8(max))", input: Result[Encodable]{HasError: false, Value: I8(math.MaxInt8)}, expect: []byte{0x0, 0x7f}},
		{label: "Encode Result(false, U16(max))", input: Result[Encodable]{HasError: false, Value: U16(math.MaxUint16)}, expect: []byte{0x0, 0xff, 0xff}},
		{label: "Encode Result(false, I16(min))", input: Result[Encodable]{HasError: false, Value: I16(math.MinInt16)}, expect: []byte{0x0, 0x00, 0x80}},
		{label: "Encode Result(false, U32(max))", input: Result[Encodable]{HasError: false, Value: U32(math.MaxUint32)}, expect: []byte{0x0, 0xff, 0xff, 0xff, 0xff}},
		{label: "Encode Result(false, I32(min))", input: Result[Encodable]{HasError: false, Value: I32(math.MinInt32)}, expect: []byte{0x0, 0x0, 0x0, 0x0, 0x80}},
		{label: "Encode Result(false, U64(max))", input: Result[Encodable]{HasError: false, Value: U64(math.MaxUint64)}, expect: []byte{0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{label: "Encode Result(false, I64(min))", input: Result[Encodable]{HasError: false, Value: I64(math.MinInt64)}, expect: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80}},
		{label: "Encode Result(true, I64(min))", input: Result[Encodable]{HasError: true, Value: I64(math.MinInt64)}, expect: []byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80}},
		{label: "Encode Result(false, U128(max))", input: Result[Encodable]{HasError: false, Value: U128{math.MaxUint64, math.MaxUint64}}, expect: []byte{0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{label: "Encode Result(true, I128(min)", input: Result[Encodable]{HasError: true, Value: I128{U64(0), U64(math.MaxInt64 + 1)}}, expect: []byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80}},
		{label: "Encode Result(false, Compact(MaxUint64)", input: Result[Encodable]{HasError: false, Value: ToCompact(uint64(math.MaxUint64))}, expect: []byte{0x0, 0x13, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},

		{label: "Encode Result(false, empty Seq[U8])", input: Result[Encodable]{HasError: false, Value: Sequence[U8]{}}, expect: []byte{0x0, 0x0}},
		{label: "Encode Result(false, Seq[U8])", input: Result[Encodable]{HasError: false, Value: Sequence[U8]{42}}, expect: []byte{0x0, 0x4, 0x2a}},
		{label: "Encode Result(false, Result(false, Seq[U8])", input: Result[Encodable]{HasError: false, Value: Result[Encodable]{HasError: true, Value: Sequence[U8]{42, 43}}}, expect: []byte{0x0, 0x1, 0x8, 0x2a, 0x2b}},
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

func Test_DecodeResultEmpty(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Result[Empty]
	}{
		{
			label:         "Decode Result(true, empty)",
			input:         []byte{0x1},
			bufferLenLeft: 0,
			expect:        Result[Empty]{true, Empty{}},
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeResult[Empty](buffer)
			assert.NoError(t, err)

			assert.Equal(t, result, e.expect)
			assert.Equal(t, buffer.Len(), e.bufferLenLeft)
		})
	}
}

func Test_DecodeResultBool(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Result[Bool]
	}{
		{
			label:         "Decode Result(true, false)",
			input:         []byte{0x1, 0x0},
			bufferLenLeft: 0,
			expect:        Result[Bool]{true, Bool(false)},
		},
		{
			label:         "Decode Result(false,false)",
			input:         []byte{0x0, 0x0},
			bufferLenLeft: 0,
			expect:        Result[Bool]{false, Bool(false)},
		},
		{
			label:         "Decode Result(false,true)",
			input:         []byte{0x0, 0x1, 0x3},
			bufferLenLeft: 1,
			expect:        Result[Bool]{false, Bool(true)},
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeResult[Bool](buffer)
			assert.NoError(t, err)

			assert.Equal(t, result, e.expect)
			assert.Equal(t, buffer.Len(), e.bufferLenLeft)
		})
	}
}

func Test_DecodeResultU8(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Result[U8]
	}{
		{
			label:         "Decode Result(false, U8(max))",
			input:         []byte{0x0, 0xff, 0xff},
			expect:        Result[U8]{false, U8(math.MaxUint8)},
			bufferLenLeft: 1,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeResult[U8](buffer)
			assert.NoError(t, err)

			assert.Equal(t, result, e.expect)
			assert.Equal(t, buffer.Len(), e.bufferLenLeft)
		})
	}
}

func Test_DecodeResultI8(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Result[I8]
	}{
		{
			label:         "Decode Result(false, I8(min))",
			input:         []byte{0x0, 0x80},
			expect:        Result[I8]{false, I8(math.MinInt8)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeResult[I8](buffer)
			assert.NoError(t, err)

			assert.Equal(t, result, e.expect)
			assert.Equal(t, buffer.Len(), e.bufferLenLeft)
		})
	}
}

func Test_DecodeResultU16(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Result[U16]
	}{
		{
			label:         "Decode Result(false, U16(max))",
			input:         []byte{0x0, 0xff, 0xff},
			expect:        Result[U16]{false, U16(math.MaxUint16)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeResult[U16](buffer)
			assert.NoError(t, err)

			assert.Equal(t, result, e.expect)
			assert.Equal(t, buffer.Len(), e.bufferLenLeft)
		})
	}
}

func Test_DecodeResultI16(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Result[I16]
	}{
		{
			label:         "Decode Result(false, I16(min))",
			input:         []byte{0x0, 0x0, 0x80},
			expect:        Result[I16]{false, I16(math.MinInt16)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeResult[I16](buffer)
			assert.NoError(t, err)

			assert.Equal(t, result, e.expect)
			assert.Equal(t, buffer.Len(), e.bufferLenLeft)
		})
	}
}

func Test_DecodeResultU32(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Result[U32]
	}{
		{
			label:         "Decode Result(false, U32(max))",
			input:         []byte{0x0, 0xff, 0xff, 0xff, 0xff},
			expect:        Result[U32]{false, U32(math.MaxUint32)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeResult[U32](buffer)
			assert.NoError(t, err)

			assert.Equal(t, result, e.expect)
			assert.Equal(t, buffer.Len(), e.bufferLenLeft)
		})
	}
}

func Test_DecodeResultI32(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Result[I32]
	}{

		{
			label:         "Decode Result(false, I32(min))",
			input:         []byte{0x0, 0x0, 0x0, 0x0, 0x80},
			expect:        Result[I32]{false, I32(math.MinInt32)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeResult[I32](buffer)
			assert.NoError(t, err)

			assert.Equal(t, result, e.expect)
			assert.Equal(t, buffer.Len(), e.bufferLenLeft)
		})
	}
}

func Test_DecodeResultU64(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Result[U64]
	}{
		{
			label:         "Decode Result(false, U64(max))",
			input:         []byte{0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			expect:        Result[U64]{false, U64(math.MaxUint64)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeResult[U64](buffer)
			assert.NoError(t, err)

			assert.Equal(t, result, e.expect)
			assert.Equal(t, buffer.Len(), e.bufferLenLeft)
		})
	}
}

func Test_DecodeResultI64(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Result[I64]
	}{
		{
			label:         "Decode Result(false, I64(min))",
			input:         []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80},
			expect:        Result[I64]{false, I64(math.MinInt64)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeResult[I64](buffer)
			assert.NoError(t, err)

			assert.Equal(t, result, e.expect)
			assert.Equal(t, buffer.Len(), e.bufferLenLeft)
		})
	}
}

func Test_DecodeResultI128(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Result[I128]
		stringValue   string
	}{
		{
			label:         "Decode Result(true, I128(min))",
			input:         []byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80},
			expect:        Result[I128]{true, I128{U64(0), U64(math.MaxInt64 + 1)}},
			bufferLenLeft: 0,
			stringValue:   "-170141183460469231731687303715884105728",
		},
		{
			label:         "Decode Result(true, I128(max))",
			input:         []byte{0x1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f},
			expect:        Result[I128]{true, I128{U64(math.MaxUint64), U64(math.MaxInt64)}},
			bufferLenLeft: 0,
			stringValue:   "170141183460469231731687303715884105727",
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeResult[I128](buffer)
			assert.NoError(t, err)

			assert.Equal(t, result, e.expect)
			assert.Equal(t, buffer.Len(), e.bufferLenLeft)
		})
	}
}

func Test_DecodeResultCompact(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Result[Compact]
	}{
		{
			label:         "Decode Compact(maxUint64)",
			input:         []byte{0x0, 0x13, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			expect:        Result[Compact]{false, ToCompact(uint64(math.MaxUint64))},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeResult[Compact](buffer)
			assert.NoError(t, err)

			assert.Equal(t, result, e.expect)
			assert.Equal(t, buffer.Len(), e.bufferLenLeft)
		})
	}
}

func Test_DecodeResultSeqU8(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Result[Sequence[U8]]
	}{
		{
			label:         "Decode Seq[U8]",
			input:         []byte{0x0, 0x4, 0x2a},
			expect:        Result[Sequence[U8]]{false, Sequence[U8]{42}},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeResult[Sequence[U8]](buffer)
			assert.NoError(t, err)

			assert.Equal(t, result, e.expect)
			assert.Equal(t, buffer.Len(), e.bufferLenLeft)
		})
	}
}

func Test_DecodeResultErrorInvalidFirstByte(t *testing.T) {
	var testExamples = []struct {
		label string
		input []byte
	}{
		{label: "Panic DecodeResult(0xff, 0x1)", input: []byte{0xff, 0x1}},
		{label: "Panic DecodeResult(0x3, 0x1)", input: []byte{0x3, 0x1}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			_, err := DecodeResult[Bool](buffer)
			assert.ErrorIs(t, err, errInvalidBoolRepresentation)
		})
	}
}
