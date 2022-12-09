package goscale

import (
	"bytes"
	"testing"
)

type NestedStruct struct {
	AA Bool
	BB Empty
	CC Str
}

type ExampleStruct struct {
	A  Bool
	B  U8
	C  U16
	D  U32
	E  U64
	F  U128
	G  I8
	H  I16
	I  I32
	J  I64
	K  I128
	L0 Compact
	L  CompactU128
	M0 Str
	M  Sequence[I32]
	// N  FixedSequence[U8]   // panic: encoding of SCALE FixedSequence[goscale.U8] field is not implemented
	// O  Dictionary[Str, U8] // panic: encoding of type Dictionary is not implemented
	P Empty

	// Q Option[U8]
	// R Result[U8]
	// S VaryingData
	// T NestedStruct // panic: encoding of type Nested Tuple is not implemented
	// U Sequence[Tuple[NestedStruct]]
}

func Test_EncodeTuple(t *testing.T) {
	buffer := &bytes.Buffer{}
	Tuple[ExampleStruct]{
		Data: ExampleStruct{
			A:  true,
			B:  1,
			C:  2,
			D:  3,
			E:  4,
			F:  U128{},
			G:  -1,
			H:  -2,
			I:  -3,
			J:  -4,
			K:  I128{},
			L0: Compact(1),
			L:  CompactU128{},
			M:  Sequence[I32]{1, 2, 3},
			// N: FixedSequence[U8]{3},
			// O: Dictionary[Str, U8]{},
			P: Empty{},

			// Q: Option[U8]{},
			// R: Result[U8]{},
			// S: VaryingData{},
			// T: NestedStruct{},
			// U: Sequence[Tuple[NestedStruct]]{},
		},
	}.Encode(buffer)
	t.Logf("\n %#x \n", buffer)

	// input := VersionData{
	// 	SpecName:         Str("abc"),
	// 	ImplName:         Str("xyz"),
	// 	AuthoringVersion: U32(1),
	// 	SpecVersion:      U32(2),
	// 	ImplVersion:      U32(3),
	// 	// Apis: Sequence[ApiItem]{
	// 	// 	Values: []ApiItem{
	// 	// 		{Name: FixedSequence[U8]{Values: []U8{6, 7, 8}}, Version: U32(9)},
	// 	// 	},
	// 	// },
	// 	TransactionVersion: U32(4),
	// 	StateVersion:       U8(5),
	// }

	// var testExamples = []struct {
	// 	label       string
	// 	input       VersionData
	// 	expectation []byte
	// }{
	// 	{
	// 		label: "VersionData{}",
	// 		input: input,
	// 		expectation: []byte{
	// 			0x0c, 0x61, 0x62, 0x63,
	// 			0x0c, 0x78, 0x79, 0x7a,
	// 			0x01, 0x00, 0x00, 0x00,
	// 			0x02, 0x00, 0x00, 0x00,
	// 			0x03, 0x00, 0x00, 0x00,
	// 			0x04, 0x00, 0x00, 0x00,
	// 			0x05,
	// 		},
	// 	},
	// }

	// for _, testExample := range testExamples {
	// 	t.Run(testExample.label, func(t *testing.T) {
	// 		buffer := &bytes.Buffer{}

	// 		EncodeTuple(testExample.input, buffer)

	// 		assertEqual(t, buffer.Bytes(), testExample.expectation)
	// 	})
	// }
}
