package goscale

import (
	"bytes"
	"io"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EncodeU32(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       U32
		expectation []byte
	}{
		{label: "uint32(127)", input: U32(127), expectation: []byte{0x7f, 0x00, 0x00, 0x00}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			err := testExample.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, testExample.expectation, buffer.Bytes())
			assert.Equal(t, testExample.expectation, testExample.input.Bytes())
		})
	}
}

func Test_DecodeU32(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation U32
	}{
		{label: "(0x7f000000)", input: []byte{0x7f, 0x00, 0x00, 0x00}, expectation: U32(127)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			result, err := DecodeU32(buffer)

			assert.NoError(t, err)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_DecodeU32_Empty(t *testing.T) {
	buffer := &bytes.Buffer{}

	result, err := DecodeU32(buffer)

	assert.Equal(t, io.EOF, err)
	assert.Equal(t, U32(0), result)
}

func Test_U32_ToBigInt(t *testing.T) {
	n := U32(127)
	nBigInt := n.ToBigInt()
	expect, ok := new(big.Int).SetString("127", 10)
	assert.True(t, ok)
	assert.Equal(t, expect, nBigInt)
}

func Test_U32_Interface(t *testing.T) {
	n := U32(127)
	assert.Equal(t, n.Interface(), Numeric(n))
}
