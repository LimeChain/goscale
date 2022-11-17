package goscale

import (
	"bytes"
	"testing"
)

type ApiItem struct {
	Name    Sequence[U8]
	Version U32
}

func (api ApiItem) Encode(buffer *bytes.Buffer) {
	EncodeTuple(api, buffer)
}

type VersionData struct {
	SpecName         Sequence[U8]
	ImplName         Sequence[U8]
	AuthoringVersion Compact
	SpecVersion      U32
	ImplVersion      U32
	// Apis               Sequence[ApiItem]
	TransactionVersion U32
	StateVersion       U8
}

func Test_EncodeTuple(t *testing.T) {
	input := VersionData{
		SpecName:         Sequence[U8]{Values: StringToSliceU8("polkadot")},
		ImplName:         Sequence[U8]{Values: StringToSliceU8("parity-polkadot")},
		AuthoringVersion: Compact(1),
		SpecVersion:      U32(2),
		ImplVersion:      U32(3),
		// Apis: Sequence[ApiItem]{
		// 	Values: []ApiItem{
		// 		{Name: Sequence[U8]{Values: []U8{6, 7, 8}}, Version: U32(9)},
		// 	},
		// },
		TransactionVersion: U32(4),
		StateVersion:       U8(5),
	}

	var testExamples = []struct {
		label       string
		input       VersionData
		expectation []byte
	}{
		{
			label:       "VersionData{}",
			input:       input,
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
