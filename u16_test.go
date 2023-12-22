package goscale

import (
	"bytes"
	"io"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EncodeU16(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       U16
		expectation []byte
	}{
		{label: "uint16(127)", input: U16(127), expectation: []byte{0x7f, 0x00}},
		{label: "uint16(42)", input: U16(42), expectation: []byte{0x2a, 0x00}},
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

func Test_DecodeU16(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation U16
	}{
		{label: "(0x2a00)", input: []byte{0x2a, 0x00}, expectation: U16(42)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			result, err := DecodeU16(buffer)

			assert.NoError(t, err)
			assert.Equal(t, testExample.expectation, result)
		})
	}
}

func Test_U16_Interface(t *testing.T) {
	n := U16(127)
	assert.Equal(t, n.Interface(), Numeric(n))
}

func Test_DecodeU16_Empty(t *testing.T) {
	buffer := &bytes.Buffer{}

	result, err := DecodeU16(buffer)

	assert.Equal(t, io.EOF, err)
	assert.Equal(t, U16(0), result)
}

func Test_U16_ToBigInt(t *testing.T) {
	n := U16(127)
	nBigInt := n.ToBigInt()
	expect, ok := new(big.Int).SetString("127", 10)
	assert.True(t, ok)
	assert.Equal(t, expect, nBigInt)
}
