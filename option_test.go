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

func (testEncodable) Bytes() []byte {
	return []byte{}
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
		{label: "Encode Option(true, I128(min)", input: Option[Encodable]{true, I128{U64(0), U64(math.MaxInt64 + 1)}}, expect: []byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80}},

		{label: "Encode Option(true, Compact(MaxUint64)", input: Option[Encodable]{true, Compact(math.MaxUint64)}, expect: []byte{0x1, 0x13, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},

		{label: "Encode Option(true, empty Seq[U8])", input: Option[Encodable]{true, Sequence[U8]{}}, expect: []byte{0x1, 0x0}},
		{label: "Encode Option(true, Seq[U8])", input: Option[Encodable]{true, Sequence[U8]{42}}, expect: []byte{0x1, 0x4, 0x2a}},

		{label: "Encode Option(true, Result(true, Seq[U8])", input: Option[Encodable]{true, Result[Encodable]{true, Sequence[U8]{42}}}, expect: []byte{0x1, 0x1, 0x4, 0x2a}},
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
			// given:
			buffer := &bytes.Buffer{}

			// when:
			e.input.Encode(buffer)

			// then:
			assertEqual(t, buffer.Bytes(), e.expect)
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
			// given:
			buffer := &bytes.Buffer{}

			// then:
			assertPanic(t, func() {
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
			expect:        Option[Encodable]{false, nil},
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeOption[Encodable](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			expect:        Option[Bool]{true, Bool(false)},
		},
		{
			label:         "Decode Option(true,true)",
			input:         []byte{0x1, 0x1, 0x3},
			bufferLenLeft: 1,
			expect:        Option[Bool]{true, Bool(true)},
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeOption[Bool](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeOptionBool(buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
		})
	}
}

func Test_DecodeOptionBoolPanics(t *testing.T) {
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
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// then:
			assertPanic(t, func() {
				DecodeOptionBool(buffer)
			})
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
			expect:        Option[U8]{true, U8(math.MaxUint8)},
			bufferLenLeft: 1,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeOption[U8](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			expect:        Option[I8]{true, I8(math.MinInt8)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeOption[I8](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			expect:        Option[U16]{true, U16(math.MaxUint16)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeOption[U16](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			expect:        Option[I16]{true, I16(math.MinInt16)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeOption[I16](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			expect:        Option[U32]{true, U32(math.MaxUint32)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeOption[U32](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			expect:        Option[I32]{true, I32(math.MinInt32)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeOption[I32](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			expect:        Option[U64]{true, U64(math.MaxUint64)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeOption[U64](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			expect:        Option[I64]{true, I64(math.MinInt64)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeOption[I64](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			expect:        Option[I128]{true, I128{U64(0), U64(math.MaxInt64 + 1)}},
			bufferLenLeft: 0,
		},
		{
			label:         "Decode Option(true, I128(max))",
			input:         []byte{0x1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f},
			expect:        Option[I128]{true, I128{U64(math.MaxUint64), U64(math.MaxInt64)}},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeOption[I128](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			expect:        Option[U128]{true, U128{math.MaxUint64, math.MaxUint64}},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeOption[U128](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			expect:        Option[Compact]{HasValue: true, Value: Compact(math.MaxUint64)},
			bufferLenLeft: 0,
		},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeOption[Compact](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			label: "Decode Seq[U8]",
			input: []byte{0x1, 0x4, 0x2a},
			expect: Option[Sequence[U8]]{
				HasValue: true, Value: Sequence[U8]{42},
			},
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
			result := DecodeOption[Sequence[U8]](buffer)

			// then:
			assertEqual(t, result, e.expect)
			assertEqual(t, buffer.Len(), e.bufferLenLeft)
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
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// then:
			assertPanic(t, func() {
				DecodeOption[Bool](buffer)
			})
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
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// then:
			assertPanic(t, func() {
				DecodeOption[Bool](buffer)
			})
		})
	}
}

func Test_DecodeOptionPanicsEmptySlice(t *testing.T) {
	var examples = []struct {
		label string
		input []byte
	}{
		{label: "Panic DecodeOption(empty slice)", input: []byte{}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// then:
			assertPanic(t, func() {
				DecodeOption[Bool](buffer)
			})
		})
	}
}

func Test_DecodeOptionPanicsNil(t *testing.T) {
	var examples = []struct {
		label string
		input []byte
	}{
		{label: "Panic DecodeOption(nil)", input: nil},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// then:
			assertPanic(t, func() {
				DecodeOption[Bool](buffer)
			})
		})
	}
}

func Test_DecodeOptionPanicsDifferentType(t *testing.T) {
	var examples = []struct {
		label string
		input []byte
	}{
		{label: "Panic DecodeOption(different type)", input: []byte{0x1}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// then:
			assertPanic(t, func() {
				DecodeOption[testEncodable](buffer)
			})
		})
	}
}

func Test_DecodeOptionPanicsU8MissingBytes(t *testing.T) {
	var examples = []struct {
		label string
		input []byte
	}{
		{label: "Panic DecodeOption(U16 - cannot read bytes, which are not found)", input: []byte{0x1, 0x5}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// then:
			assertPanic(t, func() {
				DecodeOption[U16](buffer)
			})
		})
	}
}
