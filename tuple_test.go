package goscale

import (
	"bytes"
	"testing"
)

func Test_EncodeTuple(t *testing.T) {
	type ApiItem struct {
		Name    [8]byte
		Version uint32
	}

	type VersionData struct {
		SpecName           []byte
		ImplName           []byte
		AuthoringVersion   uint32
		SpecVersion        uint32
		ImplVersion        uint32
		Apis               ApiItem
		TransactionVersion uint32
		StateVersion       uint8
	}

	input := VersionData{
		SpecName:           []byte("node"),
		ImplName:           []byte("gosemble-node"),
		AuthoringVersion:   uint32(5),
		SpecVersion:        uint32(4),
		ImplVersion:        uint32(3),
		Apis:               ApiItem{},
		TransactionVersion: uint32(2),
		StateVersion:       uint8(1),
	}

	var testExamples = []struct {
		label       string
		input       VersionData
		expectation []byte
	}{
		{label: "VersionData{}", input: input, expectation: []byte{}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := Encoder{Writer: &buffer}
			enc.EncodeTuple(testExample.input)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
		})
	}
}
