package goscale

import (
	"bytes"
	"errors"
	"io"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewOption(t *testing.T) {
	var examples = []struct {
		label  string
		input  Option[U32]
		expect Option[U32]
	}{
		{
			label:  "NewOption[U32](7)",
			input:  NewOption[U32](U32(7)),
			expect: Option[U32]{HasValue: true, Value: 7},
		},
		{
			label:  "NewOption[U32](nil)",
			input:  NewOption[U32](nil),
			expect: Option[U32]{HasValue: false},
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			err := e.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, e.input)
		})
	}
}

func Test_NewOption_Panics(t *testing.T) {
	assert.PanicsWithValue(t,
		"invalid value type for Option[T]",
		func() {
			NewOption[U32](U64(5))
		})
}

type testEncodable struct {
}

func (testEncodable) Encode(*bytes.Buffer) error {
	return nil
}

func (testEncodable) Bytes() []byte {
	return []byte{}
}

func Test_EncodeOptionBoolGeneric(t *testing.T) {
	var examples = []struct {
		label  string
		input  Option[Bool]
		expect []byte
	}{
		{label: "Encode Option(true, false)", input: NewOption[Bool](Bool(false)), expect: []byte{0x1, 0x0}},
		{label: "Encode Option(true, true)", input: NewOption[Bool](Bool(true)), expect: []byte{0x1, 0x1}},
		{label: "Encode Option(false, nil)", input: NewOption[Bool](nil), expect: []byte{0x0}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			err := e.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, buffer.Bytes())
			assert.Equal(t, e.expect, e.input.Bytes())
		})
	}
}

func Test_EncodeOptionU8(t *testing.T) {
	var examples = []struct {
		label  string
		input  Option[U8]
		expect []byte
	}{
		{label: "Encode Option(true, U8(max))", input: NewOption[U8](U8(math.MaxUint8)), expect: []byte{0x1, 0xff}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			err := e.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, buffer.Bytes())
			assert.Equal(t, e.expect, e.input.Bytes())
		})
	}
}

func Test_EncodeOptionI8(t *testing.T) {
	var examples = []struct {
		label  string
		input  Option[I8]
		expect []byte
	}{
		{label: "Encode Option(true, I8(min))", input: NewOption[I8](I8(math.MinInt8)), expect: []byte{0x1, 0x80}},
		{label: "Encode Option(true, I8(max))", input: NewOption[I8](I8(math.MaxInt8)), expect: []byte{0x1, 0x7f}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			err := e.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, buffer.Bytes())
			assert.Equal(t, e.expect, e.input.Bytes())
		})
	}
}

func Test_EncodeOptionU16(t *testing.T) {
	var examples = []struct {
		label  string
		input  Option[U16]
		expect []byte
	}{
		{label: "Encode Option(true, U16(max))", input: NewOption[U16](U16(math.MaxUint16)), expect: []byte{0x1, 0xff, 0xff}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			err := e.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, buffer.Bytes())
			assert.Equal(t, e.expect, e.input.Bytes())
		})
	}
}

func Test_EncodeOptionI16(t *testing.T) {
	var examples = []struct {
		label  string
		input  Option[I16]
		expect []byte
	}{
		{label: "Encode Option(true, I16(min))", input: NewOption[I16](I16(math.MinInt16)), expect: []byte{0x1, 0x00, 0x80}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			err := e.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, buffer.Bytes())
			assert.Equal(t, e.expect, e.input.Bytes())
		})
	}
}

func Test_EncodeOptionU32(t *testing.T) {
	var examples = []struct {
		label  string
		input  Option[U32]
		expect []byte
	}{
		{label: "Encode Option(true, U32(max))", input: NewOption[U32](U32(math.MaxUint32)), expect: []byte{0x1, 0xff, 0xff, 0xff, 0xff}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			err := e.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, buffer.Bytes())
			assert.Equal(t, e.expect, e.input.Bytes())
		})
	}
}

func Test_EncodeOptionI32(t *testing.T) {
	var examples = []struct {
		label  string
		input  Option[I32]
		expect []byte
	}{
		{label: "Encode Option(true, I32(min))", input: NewOption[I32](I32(math.MinInt32)), expect: []byte{0x1, 0x0, 0x0, 0x0, 0x80}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			err := e.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, buffer.Bytes())
			assert.Equal(t, e.expect, e.input.Bytes())
		})
	}
}

func Test_EncodeOptionU64(t *testing.T) {
	var examples = []struct {
		label  string
		input  Option[U64]
		expect []byte
	}{
		{label: "Encode Option(true, U64(max))", input: NewOption[U64](U64(math.MaxUint64)), expect: []byte{0x1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			err := e.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, buffer.Bytes())
			assert.Equal(t, e.expect, e.input.Bytes())
		})
	}
}

func Test_EncodeOptionI64(t *testing.T) {
	var examples = []struct {
		label  string
		input  Option[I64]
		expect []byte
	}{
		{label: "Encode Option(true, I64(min))", input: NewOption[I64](I64(math.MinInt64)), expect: []byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80}},
		{label: "Encode Option(false, I64(min))", input: NewOption[I64](nil), expect: []byte{0x0}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			err := e.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, buffer.Bytes())
			assert.Equal(t, e.expect, e.input.Bytes())
		})
	}
}

func Test_EncodeOptionU128(t *testing.T) {
	var examples = []struct {
		label  string
		input  Option[U128]
		expect []byte
	}{
		{label: "Encode Option(true, U128(max)", input: NewOption[U128](U128{math.MaxUint64, math.MaxUint64}), expect: []byte{0x1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			err := e.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, buffer.Bytes())
			assert.Equal(t, e.expect, e.input.Bytes())
		})
	}
}

func Test_EncodeOptionI128(t *testing.T) {
	var examples = []struct {
		label  string
		input  Option[I128]
		expect []byte
	}{
		{label: "Encode Option(true, I128(min)", input: NewOption[I128](I128{U64(0), U64(math.MaxInt64 + 1)}), expect: []byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			err := e.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, buffer.Bytes())
			assert.Equal(t, e.expect, e.input.Bytes())
		})
	}
}

func Test_EncodeOptionCompact(t *testing.T) {
	var examples = []struct {
		label  string
		input  Option[Compact]
		expect []byte
	}{
		{label: "Encode Option(true, Compact(MaxUint64)", input: NewOption[Compact](ToCompact(uint64(math.MaxUint64))), expect: []byte{0x1, 0x13, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			err := e.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, buffer.Bytes())
			assert.Equal(t, e.expect, e.input.Bytes())
		})
	}
}

func Test_EncodeOptionSequenceU8(t *testing.T) {
	var examples = []struct {
		label  string
		input  Option[Sequence[U8]]
		expect []byte
	}{
		{label: "Encode Option(true, empty Seq[U8])", input: NewOption[Sequence[U8]](Sequence[U8]{}), expect: []byte{0x1, 0x0}},
		{label: "Encode Option(true, Seq[U8])", input: NewOption[Sequence[U8]](Sequence[U8]{42}), expect: []byte{0x1, 0x4, 0x2a}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			err := e.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, buffer.Bytes())
			assert.Equal(t, e.expect, e.input.Bytes())
		})
	}
}

func Test_EncodeOptionResult(t *testing.T) {
	type ResultValue = Result[Sequence[U8]]

	var examples = []struct {
		label  string
		input  Option[ResultValue]
		expect []byte
	}{
		{
			label:  "Encode Option(true, Result(true, Seq[U8])",
			input:  NewOption[ResultValue](ResultValue{HasError: false, Value: Sequence[U8]{42}}),
			expect: []byte{0x1, 0x0, 0x4, 0x2a}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			err := e.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, buffer.Bytes())
			assert.Equal(t, e.expect, e.input.Bytes())
		})
	}
}

func Test_EncodeOptionBool(t *testing.T) {
	var examples = []struct {
		label  string
		input  OptionBool
		expect []byte
	}{
		{label: "Encode OptionBool(true, false)", input: OptionBool{true, Bool(false)}, expect: []byte{0x2}},
		{label: "Encode OptionBool(true, true)", input: OptionBool{true, Bool(true)}, expect: []byte{0x1}},
		{label: "Encode Option(false, nil)", input: OptionBool{HasValue: false}, expect: []byte{0x0}},
		{label: "Encode Option(false, true)", input: OptionBool{false, Bool(true)}, expect: []byte{0x0}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			err := e.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, buffer.Bytes())
			assert.Equal(t, e.expect, e.input.Bytes())
		})
	}
}

func Test_EncodeOptionPanics(t *testing.T) {
	var examples = []struct {
		label string
		input Option[Encodable]
	}{
		{label: "Panic EncodeOption(true, nil)", input: Option[Encodable]{true, nil}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			assert.Panics(t, func() {
				e.input.Encode(buffer)
			})
		})
	}
}

func Test_DecodeOptionNil(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Option[Encodable]
	}{
		{
			label:         "Decode Option(false, nil)",
			input:         []byte{0x0},
			bufferLenLeft: 0,
			expect:        NewOption[Encodable](nil),
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeOption[Encodable](buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, result)
			assert.Equal(t, e.bufferLenLeft, buffer.Len())
		})
	}
}

func Test_DecodeOptionFromBool(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Option[Bool]
	}{
		{
			label:         "Decode Option(true,false)",
			input:         []byte{0x1, 0x0},
			bufferLenLeft: 0,
			expect:        NewOption[Bool](Bool(false)),
		},
		{
			label:         "Decode Option(true,true)",
			input:         []byte{0x1, 0x1, 0x3},
			bufferLenLeft: 1,
			expect:        NewOption[Bool](Bool(true)),
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeOption[Bool](buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, result)
			assert.Equal(t, e.bufferLenLeft, buffer.Len())
		})
	}
}

func Test_DecodeOptionBool(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        OptionBool
	}{
		{
			label:         "Decode OptionBool(true,nil)",
			input:         []byte{0x0},
			bufferLenLeft: 0,
			expect:        OptionBool{HasValue: false},
		},
		{
			label:         "Decode OptionBool(true,false)",
			input:         []byte{0x2},
			bufferLenLeft: 0,
			expect:        OptionBool{true, Bool(false)},
		},
		{
			label:         "Decode OptionBool(true,true)",
			input:         []byte{0x1, 0x3},
			bufferLenLeft: 1,
			expect:        OptionBool{true, Bool(true)},
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeOptionBool(buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, result)
			assert.Equal(t, e.bufferLenLeft, buffer.Len())
		})
	}
}

func Test_DecodeOptionBool_InvalidOptionBoolError(t *testing.T) {
	var examples = []struct {
		label string
		input []byte
	}{
		{
			label: "Panic DecodeOptionBool(0x03)",
			input: []byte{0x3},
		},
		{
			label: "Panic DecodeOptionBool(0xff)",
			input: []byte{0xff},
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			_, err := DecodeOptionBool(buffer)
			assert.ErrorIs(t, errInvalidOptionBoolRepresentation, err)
		})
	}
}

func Test_DecodeOptionU8(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Option[U8]
	}{
		{
			label:         "Decode Option(true, U8(max))",
			input:         []byte{0x1, 0xff, 0xff},
			expect:        NewOption[U8](U8(math.MaxUint8)),
			bufferLenLeft: 1,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeOption[U8](buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, result)
			assert.Equal(t, e.bufferLenLeft, buffer.Len())
		})
	}
}

func Test_DecodeOptionI8(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Option[I8]
	}{
		{
			label:         "Decode Option(true, I8(min))",
			input:         []byte{0x1, 0x80},
			expect:        NewOption[I8](I8(math.MinInt8)),
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeOption[I8](buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, result)
			assert.Equal(t, e.bufferLenLeft, buffer.Len())
		})
	}
}

func Test_DecodeOptionU16(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Option[U16]
	}{
		{
			label:         "Decode Option(true, U16(max))",
			input:         []byte{0x1, 0xff, 0xff},
			expect:        NewOption[U16](U16(math.MaxUint16)),
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeOption[U16](buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, result)
			assert.Equal(t, e.bufferLenLeft, buffer.Len())
		})
	}
}

func Test_DecodeOptionI16(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Option[I16]
	}{
		{
			label:         "Decode Option(true, I16(min))",
			input:         []byte{0x1, 0x0, 0x80},
			expect:        NewOption[I16](I16(math.MinInt16)),
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeOption[I16](buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, result)
			assert.Equal(t, e.bufferLenLeft, buffer.Len())
		})
	}
}

func Test_DecodeOptionU32(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Option[U32]
	}{
		{
			label:         "Decode Option(true, U32(max))",
			input:         []byte{0x1, 0xff, 0xff, 0xff, 0xff},
			expect:        NewOption[U32](U32(math.MaxUint32)),
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeOption[U32](buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, result)
			assert.Equal(t, e.bufferLenLeft, buffer.Len())
		})
	}
}

func Test_DecodeOptionI32(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Option[I32]
	}{
		{
			label:         "Decode Option(true, I32(min))",
			input:         []byte{0x1, 0x0, 0x0, 0x0, 0x80},
			expect:        NewOption[I32](I32(math.MinInt32)),
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeOption[I32](buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, result)
			assert.Equal(t, e.bufferLenLeft, buffer.Len())
		})
	}
}

func Test_DecodeOptionU64(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Option[U64]
	}{
		{
			label:         "Decode Option(true, U64(max))",
			input:         []byte{0x1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			expect:        NewOption[U64](U64(math.MaxUint64)),
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeOption[U64](buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, result)
			assert.Equal(t, e.bufferLenLeft, buffer.Len())
		})
	}
}

func Test_DecodeOptionI64(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Option[I64]
	}{
		{
			label:         "Decode Option(true, I64(min))",
			input:         []byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80},
			expect:        NewOption[I64](I64(math.MinInt64)),
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeOption[I64](buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, result)
			assert.Equal(t, e.bufferLenLeft, buffer.Len())
		})
	}
}

func Test_DecodeOptionI128(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Option[I128]
		stringValue   string
	}{
		{
			label:         "Decode Option(true, I128(min))",
			input:         []byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80},
			expect:        NewOption[I128](I128{U64(0), U64(math.MaxInt64 + 1)}),
			bufferLenLeft: 0,
		},
		{
			label:         "Decode Option(true, I128(max))",
			input:         []byte{0x1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f},
			expect:        NewOption[I128](I128{U64(math.MaxUint64), U64(math.MaxInt64)}),
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeOption[I128](buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, result)
			assert.Equal(t, e.bufferLenLeft, buffer.Len())
		})
	}
}

func Test_DecodeOptionU128(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Option[U128]
	}{
		{
			label:         "Decode Option(true, U128(max))",
			input:         []byte{0x1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			expect:        NewOption[U128](U128{math.MaxUint64, math.MaxUint64}),
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeOption[U128](buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, result)
			assert.Equal(t, e.bufferLenLeft, buffer.Len())
		})
	}
}

func Test_DecodeOptionCompact(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Option[Compact]
	}{
		{
			label:         "Decode Compact(maxUint64)",
			input:         []byte{0x1, 0x13, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			expect:        NewOption[Compact](ToCompact(uint64(math.MaxUint64))),
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeOption[Compact](buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, result)
			assert.Equal(t, e.bufferLenLeft, buffer.Len())
		})
	}
}

func Test_DecodeOptionSequenceU8(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Option[Sequence[U8]]
	}{
		{
			label:         "Decode Seq[U8]",
			input:         []byte{0x1, 0x4, 0x2a},
			expect:        NewOption[Sequence[U8]](Sequence[U8]{42}),
			bufferLenLeft: 0,
		},
		// TODO: Decode Option<Result<true, any>>
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeOption[Sequence[U8]](buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, result)
			assert.Equal(t, e.bufferLenLeft, buffer.Len())
		})
	}
}

func Test_DecodeOptionPanicsMissingBoolBytes(t *testing.T) {
	var examples = []struct {
		label string
		input []byte
	}{
		{label: "Panic DecodeOption(0x1)", input: []byte{0x1}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			_, err := DecodeOption[Bool](buffer)
			assert.ErrorIs(t, io.EOF, err)
		})
	}
}

func Test_DecodeOptionPanicsInvalidFirstBoolByte(t *testing.T) {
	var examples = []struct {
		label string
		input []byte
	}{
		{label: "Panic DecodeOption(invalid first byte)", input: []byte{0x3}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			_, err := DecodeOption[Bool](buffer)

			assert.ErrorIs(t, errInvalidBoolRepresentation, err)
		})
	}
}

func Test_DecodeOptionErrorEmptySlice(t *testing.T) {
	var examples = []struct {
		label string
		input []byte
	}{
		{label: "Panic DecodeOption(empty slice)", input: []byte{}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			_, err := DecodeOption[Bool](buffer)
			assert.ErrorIs(t, io.EOF, err)
		})
	}
}

func Test_DecodeOptionErrorNil(t *testing.T) {
	var examples = []struct {
		label string
		input []byte
	}{
		{label: "Panic DecodeOption(nil)", input: nil},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			_, err := DecodeOption[Bool](buffer)
			assert.ErrorIs(t, io.EOF, err)
		})
	}
}

func Test_DecodeOptionErrorDifferentType(t *testing.T) {
	var examples = []struct {
		label string
		input []byte
	}{
		{label: "Panic DecodeOption(different type)", input: []byte{0x1}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			_, err := DecodeOption[testEncodable](buffer)

			assert.ErrorIs(t, errTypeNotFound, err)
		})
	}
}

func Test_DecodeOptionErrorU8MissingBytes(t *testing.T) {
	var examples = []struct {
		label string
		input []byte
	}{
		{label: "Panic DecodeOption(U16 - cannot read bytes, which are not found)", input: []byte{0x1, 0x5}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			_, err := DecodeOption[U16](buffer)

			expectedErr := errors.New("can not read the required number of bytes 2, only 1 available")

			assert.Equal(t, expectedErr, err)
		})
	}
}
