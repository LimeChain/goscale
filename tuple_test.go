package goscale

import (
	"bytes"
	"math"
	"math/big"
	"testing"
)

type TupleBool struct {
	Tuple
	A0 Bool
	A1 Bool
}

func Test_EncodeTupleBool(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       TupleBool
		expectation []byte
	}{
		{
			label: "TupleBool",
			input: TupleBool{A0: true, A1: false},
			expectation: []byte{
				0x01, // A0
				0x00, // A1
			},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			EncodeTuple(testExample.input, buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

type TupleU8I8 struct {
	Tuple
	B0 U8
	B1 I8
}

func Test_EncodeTupleU8I8(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       TupleU8I8
		expectation []byte
	}{
		{
			label: "TupleU8I8",
			input: TupleU8I8{B0: 255, B1: -128},
			expectation: []byte{
				0xff, // B0
				0x80, // B1
			},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			EncodeTuple(testExample.input, buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

type TupleU16I16 struct {
	Tuple
	C0 U16
	C1 I16
}

func Test_EncodeTupleU16I16(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       TupleU16I16
		expectation []byte
	}{
		{
			label: "TupleU16I16",
			input: TupleU16I16{C0: 65535, C1: -128},
			expectation: []byte{
				0xff, 0xff, // C0
				0x80, 0xff, // C1
			},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			EncodeTuple(testExample.input, buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

type TupleU32I32 struct {
	Tuple
	D0 U32
	D1 I32
}

func Test_EncodeTupleU32I32(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       TupleU32I32
		expectation []byte
	}{
		{
			label: "TupleU32I32",
			input: TupleU32I32{D0: 4294967295, D1: 16777215},
			expectation: []byte{
				0xff, 0xff, 0xff, 0xff, // D0
				0xff, 0xff, 0xff, 0x00, // D1
			},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			EncodeTuple(testExample.input, buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

type TupleU64I64 struct {
	Tuple
	E0 U64
	E1 I64
}

func Test_EncodeTupleU64I64(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       TupleU64I64
		expectation []byte
	}{
		{
			label: "TupleU64I64",
			input: TupleU64I64{E0: 18446744073709551615, E1: -128},
			expectation: []byte{
				0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // E0
				0x80, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // E1
			},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			EncodeTuple(testExample.input, buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

type TupleU128I128 struct {
	Tuple
	F0 U128
	F1 I128
}

func Test_EncodeTupleU128I128(t *testing.T) {
	u, _ := new(big.Int).SetString("340282366920938463463374607431768211356", 10)
	i, _ := new(big.Int).SetString("-170141183460469231731687303715884105728", 10)

	var testExamples = []struct {
		label       string
		input       TupleU128I128
		expectation []byte
	}{
		{
			label: "TupleU128I128",
			input: TupleU128I128{F0: NewU128FromBigInt(u), F1: NewI128FromBigInt(*i)},
			expectation: []byte{
				0x9c, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // F0
				0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x80, // F1
			},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			EncodeTuple(testExample.input, buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

type TupleCompact struct {
	Tuple
	G0 Compact
	G1 CompactU128
}

func Test_EncodeTupleCompact(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       TupleCompact
		expectation []byte
	}{
		{
			label: "TupleCompact",
			input: TupleCompact{
				G0: Compact(1073741824),
				G1: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(1073741823))},
			},
			expectation: []byte{
				0x03, 0x00, 0x00, 0x00, 0x40, // G0
				0xfe, 0xff, 0xff, 0xff, // G1
			},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			EncodeTuple(testExample.input, buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

type TupleStr struct {
	Tuple
	H0 Str
	H1 Str
}

func Test_EncodeTupleStr(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       TupleStr
		expectation []byte
	}{
		{
			label: "TupleStr",
			input: TupleStr{
				H0: Str("abc"),
				H1: Str("xyz"),
			},
			expectation: []byte{
				0x0c, 0x61, 0x62, 0x63, // H0
				0x0c, 0x78, 0x79, 0x7a, // H1
			},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			EncodeTuple(testExample.input, buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

type TupleSequence struct {
	Tuple
	I0  Sequence[Bool]
	I1  Sequence[U8]
	I2  Sequence[I8]
	I3  Sequence[U16]
	I4  Sequence[I16]
	I5  Sequence[U32]
	I6  Sequence[I32]
	I7  Sequence[U64]
	I8  Sequence[I64]
	I9  Sequence[U128]
	I10 Sequence[I128]
	I11 Sequence[Compact]
	I12 Sequence[CompactU128]
	I13 Sequence[Str]
}

func Test_EncodeTupleSequence(t *testing.T) {
	t.Skip()
	var testExamples = []struct {
		label       string
		input       TupleSequence
		expectation []byte
	}{
		{
			label:       "TupleSequence",
			input:       TupleSequence{},
			expectation: []byte{},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			EncodeTuple(testExample.input, buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

// J FixedSequence[T] // not implemented

// K Dictionary[K, V] // not implemented

type TupleVaryingData struct {
	Tuple
	L0 VaryingData
	L1 VaryingData
}

func Test_EncodeTupleVaryingData(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       TupleVaryingData
		expectation []byte
	}{
		{
			label: "TupleVaryingData",
			input: TupleVaryingData{
				L0: NewVaryingData(U8(42), Bool(true)),
				L1: NewVaryingData(U128{math.MaxUint64, math.MaxUint64}, Empty{}),
			},
			expectation: []byte{
				0x00, 0x2a, 0x01, 0x01, // L0
				0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1, // L1
			},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			EncodeTuple(testExample.input, buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

type TupleOption struct {
	Tuple
	M0 Option[U8]
	M1 Option[Bool]
	M2 Option[Str]
	// M3 Option[TupleNested]
}

func Test_EncodeTupleOption(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       TupleOption
		expectation []byte
	}{
		{
			label: "TupleOption",
			input: TupleOption{
				M0: Option[U8]{true, 3},
				M1: Option[Bool]{true, false},
				M2: Option[Str]{true, "abc"},
			},
			expectation: []byte{
				0x01, 0x03, // M0
				0x01, 0x00, // M1
				0x01, 0x0c, 0x61, 0x62, 0x63, // M2
			},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			EncodeTuple(testExample.input, buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

type TupleResult struct {
	Tuple
	N0 Result[U8]
	N1 Result[Bool]
	N2 Result[Str]
	// N3 Result[TupleNested]
}

func Test_EncodeTupleResult(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       TupleResult
		expectation []byte
	}{
		{
			label: "TupleResult",
			input: TupleResult{
				N0: Result[U8]{true, 3},
				N1: Result[Bool]{true, false},
				N2: Result[Str]{true, "abc"},
				// N3: Result[TupleNested]{},
			},
			expectation: []byte{
				0x1, 0x03, // Result[U8]
				0x01, 0x00, // Result[Bool]
				0x01, 0x0c, 0x61, 0x62, 0x63, // Result[Str]
			},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			EncodeTuple(testExample.input, buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

type TupleEmpty struct {
	Tuple
	O0 Empty
	O1 Empty
}

func Test_EncodeTupleEmpty(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       TupleEmpty
		expectation int
	}{
		{
			label: "TupleEmpty",
			input: TupleEmpty{
				O0: Empty{},
				O1: Empty{},
			},
			expectation: 0,
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			EncodeTuple(testExample.input, buffer)

			assertEqual(t, len(buffer.Bytes()), testExample.expectation)
		})
	}
}

type TupleNested struct {
	Tuple
	P0 TupleBool
	P1 TupleU8I8
	P3 TupleStr
}

func Test_EncodeTupleNested(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       TupleNested
		expectation []byte
	}{
		{
			label: "TupleNested",
			input: TupleNested{
				P0: TupleBool{A0: true, A1: false},
				P1: TupleU8I8{B0: 1, B1: 2},
				P3: TupleStr{H0: Str("abc"), H1: Str("xyz")},
			},
			expectation: []byte{
				0x01, 0x00,
				0x01, 0x02,
				0x0c, 0x61, 0x62, 0x63,
				0x0c, 0x78, 0x79, 0x7a,
			},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			EncodeTuple(testExample.input, buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)

			// t.Logf("\n\nRESULT: %#x\n\n", buffer.Bytes())
		})
	}
}

// *T
// func() T
// chan T
// unsafe.Pointer
// uintptr
// complex128
// complex64
// float64
// float32
