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
		{label: "Encode Result(false, I128(min)", input: Result[Encodable]{false, I128{U64(0), U64(math.MaxInt64 + 1)}}, expect: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80}},

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

func Test_DecodeResultEmpty(t *testing.T) {
	var examples = []struct {
		label         string
		input         []byte
		bufferLenLeft int
		expect        Result[Empty]
	}{
		{
			label:         "Decode Result(false, empty)",
			input:         []byte{0x0},
			bufferLenLeft: 0,
			expect:        Result[Empty]{false, Empty{}},
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeResult[Empty](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			label:         "Decode Result(false, false)",
			input:         []byte{0x0, 0x0},
			bufferLenLeft: 0,
			expect:        Result[Bool]{false, Bool(false)},
		},
		{
			label:         "Decode Result(true,false)",
			input:         []byte{0x1, 0x0},
			bufferLenLeft: 0,
			expect:        Result[Bool]{true, Bool(false)},
		},
		{
			label:         "Decode Result(true,true)",
			input:         []byte{0x1, 0x1, 0x3},
			bufferLenLeft: 1,
			expect:        Result[Bool]{true, Bool(true)},
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeResult[Bool](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			label:         "Decode Result(true, U8(max))",
			input:         []byte{0x1, 0xff, 0xff},
			expect:        Result[U8]{true, U8(math.MaxUint8)},
			bufferLenLeft: 1,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeResult[U8](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			label:         "Decode Result(true, I8(min))",
			input:         []byte{0x1, 0x80},
			expect:        Result[I8]{true, I8(math.MinInt8)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeResult[I8](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			label:         "Decode Result(true, U16(max))",
			input:         []byte{0x1, 0xff, 0xff},
			expect:        Result[U16]{true, U16(math.MaxUint16)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeResult[U16](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			label:         "Decode Result(true, I16(min))",
			input:         []byte{0x1, 0x0, 0x80},
			expect:        Result[I16]{true, I16(math.MinInt16)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeResult[I16](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			label:         "Decode Result(true, U32(max))",
			input:         []byte{0x1, 0xff, 0xff, 0xff, 0xff},
			expect:        Result[U32]{true, U32(math.MaxUint32)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeResult[U32](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			label:         "Decode Result(true, I32(min))",
			input:         []byte{0x1, 0x0, 0x0, 0x0, 0x80},
			expect:        Result[I32]{true, I32(math.MinInt32)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeResult[I32](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			label:         "Decode Result(true, U64(max))",
			input:         []byte{0x1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			expect:        Result[U64]{true, U64(math.MaxUint64)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeResult[U64](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			label:         "Decode Result(true, I64(min))",
			input:         []byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80},
			expect:        Result[I64]{true, I64(math.MinInt64)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeResult[I64](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			label:         "Decode Result(false, I128(min))",
			input:         []byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80},
			expect:        Result[I128]{true, I128{U64(0), U64(math.MaxInt64 + 1)}},
			bufferLenLeft: 0,
			stringValue:   "-170141183460469231731687303715884105728",
		},
		{
			label:         "Decode Result(false, I128(max))",
			input:         []byte{0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f},
			expect:        Result[I128]{false, I128{U64(math.MaxUint64), U64(math.MaxInt64)}},
			bufferLenLeft: 0,
			stringValue:   "170141183460469231731687303715884105727",
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeResult[I128](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			input:         []byte{0x1, 0x13, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			expect:        Result[Compact]{true, Compact(math.MaxUint64)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeResult[Compact](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			input:         []byte{0x1, 0x4, 0x2a},
			expect:        Result[Sequence[U8]]{true, Sequence[U8]{[]U8{42}}},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeResult[Sequence[U8]](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
		})
	}
}
