package goscale

// import (
// 	"bytes"
// 	"testing"
// )

// func Test_EncodeTuple(t *testing.T) {
// 	t.Skip()
// 	type ApiItem struct {
// 		Name    []byte // [8]byte
// 		Version uint32
// 	}

// 	type VersionData struct {
// 		SpecName           []byte
// 		ImplName           []byte
// 		AuthoringVersion   uint32
// 		SpecVersion        uint32
// 		ImplVersion        uint32
// 		Apis               []ApiItem
// 		TransactionVersion uint32
// 		StateVersion       uint8
// 	}

// 	input := VersionData{
// 		SpecName:           []byte("polkadot"),
// 		ImplName:           []byte("parity-polkadot"),
// 		AuthoringVersion:   uint32(0),
// 		SpecVersion:        uint32(9160),
// 		ImplVersion:        uint32(0),
// 		Apis:               []ApiItem{{Name: []byte{3, 3, 3, 3, 3, 3, 3, 3}, Version: 5}},
// 		TransactionVersion: uint32(0),
// 		StateVersion:       uint8(0),
// 	}

// 	var testExamples = []struct {
// 		label       string
// 		input       VersionData
// 		expectation []byte
// 	}{
// 		{
// 			label: "VersionData{}",
// 			input: input,
// 			expectation: []byte{
// 				0x20, 0x70, 0x6f, 0x6c, 0x6b, 0x61, 0x64, 0x6f, 0x74, 0x3c, 0x70, 0x61, 0x72, 0x69, 0x74, 0x79, 0x2d, 0x70, 0x6f, 0x6c, 0x6b, 0x61, 0x64, 0x6f, 0x74, 0x00, 0x00, 0x00, 0x00, 0xc8, 0x23, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x20, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x03, 0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
// 			},
// 		},
// 	}

// 	for _, testExample := range testExamples {
// 		t.Run(testExample.label, func(t *testing.T) {
// 			buffer := bytes.Buffer{}

// 			enc := Encoder{Writer: &buffer}
// 			enc.EncodeTuple(testExample.input)

// 			result := buffer.Bytes()

// 			assertEqual(t, result, testExample.expectation)
// 		})
// 	}
// }
