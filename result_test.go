package goscale

import (
	"bytes"
	"io"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Result_Encode(t *testing.T) {
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

			err := e.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, buffer.Bytes())
			assert.Equal(t, e.expect, e.input.Bytes())
		})
	}
}

func Test_DecodeResult_ValidValue(t *testing.T) {
	// encoded Result consists of (false, U8)
	input := []byte{0, 10}
	buffer := bytes.NewBuffer(input)

	expect := Result[Encodable]{
		HasError: false,
		Value:    U8(10),
	}

	result, err := DecodeResult(buffer, DecodeU8, DecodeStr)

	assert.NoError(t, err)
	assert.Equal(t, expect, result)
	assert.Equal(t, 0, buffer.Len())
}

func Test_DecodeResult_ValidValue_Fails_EOF(t *testing.T) {
	// encoded Result consists of (true, U16)
	input := []byte{0}
	buffer := bytes.NewBuffer(input)

	result, err := DecodeResult(buffer, DecodeCompact, DecodeU16)

	assert.Equal(t, io.EOF, err)
	assert.Equal(t, Result[Encodable]{}, result)
}

func Test_DecodeResult_Error(t *testing.T) {
	// encoded Result consists of (true, U16)
	input := []byte{1, 10, 0}
	buffer := bytes.NewBuffer(input)

	expect := Result[Encodable]{
		HasError: true,
		Value:    U16(10),
	}

	result, err := DecodeResult(buffer, DecodeCompact, DecodeU16)

	assert.NoError(t, err)
	assert.Equal(t, expect, result)
	assert.Equal(t, 0, buffer.Len())
}

func Test_DecodeResult_ErrorValue_Fails_EOF(t *testing.T) {
	// encoded Result consists of (true, U16)
	input := []byte{1}
	buffer := bytes.NewBuffer(input)

	result, err := DecodeResult(buffer, DecodeCompact, DecodeU16)

	assert.Equal(t, io.EOF, err)
	assert.Equal(t, Result[Encodable]{}, result)
}

func Test_DecodeResult_CorrectLengthRead(t *testing.T) {
	// encoded Result consists of (true, U32)
	// U32 is 4 bytes, so the first 5 bytes will only be read
	input := []byte{1, 128, 0, 0, 0, 127, 126}
	buffer := bytes.NewBuffer(input)

	expect := Result[Encodable]{
		HasError: true,
		Value:    U32(128),
	}

	result, err := DecodeResult(buffer, DecodeCompact, DecodeU32)

	assert.NoError(t, err)
	assert.Equal(t, expect, result)
	assert.Equal(t, 2, buffer.Len())
	assert.Equal(t, []byte{127, 126}, buffer.Bytes())
}

func Test_DecodeResult_Fails_InvalidFirstByte(t *testing.T) {
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

			_, err := DecodeResult(buffer, DecodeBool, DecodeU8)
			assert.ErrorIs(t, errInvalidBoolRepresentation, err)
		})
	}
}
