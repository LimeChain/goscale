package goscale

import (
	"bytes"
	"math"
	"testing"
)

// TODO: {label: "int(-127)", input: int(-127), expectation: []byte{0x13, 0x81, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},

func Test_EncodeUintCompact32(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Compact
		expectation []byte
	}{
		{label: "uint64(0)", input: Compact(0), expectation: []byte{0x00}},
		{label: "uint64(1)", input: Compact(1), expectation: []byte{0x04}},
		{label: "uint64(42)", input: Compact(42), expectation: []byte{0xa8}},
		{label: "uint64(63)", input: Compact(63), expectation: []byte{0xfc}},
		{label: "uint64(64)", input: Compact(64), expectation: []byte{0x01, 0x01}},
		{label: "uint64(69)", input: Compact(69), expectation: []byte{0x15, 0x01}},
		{label: "uint64(127)", input: Compact(127), expectation: []byte{0xfd, 0x01}},
		{label: "uint64(65535)", input: Compact(65535), expectation: []byte{0xfe, 0xff, 0x03, 0x00}},
		{label: "uint64(16383)", input: Compact(16383), expectation: []byte{0xfd, 0xff}},
		{label: "uint64(16384)", input: Compact(16384), expectation: []byte{0x02, 0x00, 0x01, 0x00}},
		{label: "uint64(1073741823)", input: Compact(1073741823), expectation: []byte{0xfe, 0xff, 0xff, 0xff}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := &Encoder{Writer: &buffer}
			testExample.input.Encode(enc)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_DecodeUintCompact32(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation Compact
	}{
		{label: "(0x00)", input: []byte{0x00}, expectation: Compact(0)},
		{label: "(0x04)", input: []byte{0x04}, expectation: Compact(1)},
		{label: "(0xa8)", input: []byte{0xa8}, expectation: Compact(42)},
		{label: "(0xfc)", input: []byte{0xfc}, expectation: Compact(63)},
		{label: "(0x0101)", input: []byte{0x01, 0x01}, expectation: Compact(64)},
		{label: "(0x1501)", input: []byte{0x15, 0x01}, expectation: Compact(69)},
		{label: "(0xfd01)", input: []byte{0xfd, 0x01}, expectation: Compact(127)},
		{label: "(0xfeff0300)", input: []byte{0xfe, 0xff, 0x03, 0x00}, expectation: Compact(65535)},
		{label: "(0xfdff)", input: []byte{0xfd, 0xff}, expectation: Compact(16383)},
		{label: "(0x02000100)", input: []byte{0x02, 0x00, 0x01, 0x00}, expectation: Compact(16384)},
		{label: "(0xfeffffff)", input: []byte{0xfe, 0xff, 0xff, 0xff}, expectation: Compact(1073741823)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := &Encoder{Writer: &buffer}
			testExample.expectation.Encode(enc)

			dec := Decoder{Reader: &buffer}
			result := dec.DecodeCompact()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeUintCompact(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       Compact
		expectation []byte
	}{
		{label: "uint64(1073741824)", input: Compact(1073741824), expectation: []byte{0x03, 0x00, 0x00, 0x00, 0x40}},
		{label: "uint64(100000000000000)", input: Compact(100000000000000), expectation: []byte{0x0b, 0x00, 0x40, 0x7a, 0x10, 0xf3, 0x5a}},
		{label: "uint64(1<<32 - 1)", input: Compact(1<<32 - 1), expectation: []byte{0x03, 0xff, 0xff, 0xff, 0xff}},
		{label: "uint64(1 << 32)", input: Compact(1 << 32), expectation: []byte{0x07, 0x00, 0x00, 0x00, 0x00, 0x01}},
		{label: "uint64(1 << 40)", input: Compact(1 << 40), expectation: []byte{0x0b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}},
		{label: "uint64(1 << 48)", input: Compact(1 << 48), expectation: []byte{0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}},
		{label: "uint64(1<<56 - 1)", input: Compact(1<<56 - 1), expectation: []byte{0x0f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{label: "uint64(1 << 56)", input: Compact(1 << 56), expectation: []byte{0x13, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}},
		{label: "math.MaxUint64", input: Compact(math.MaxUint64), expectation: []byte{0x13, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := &Encoder{Writer: &buffer}
			testExample.input.Encode(enc)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_DecodeUintCompact(t *testing.T) {
	t.Skip()
	var testExamples = []struct {
		label       string
		input       []byte
		expectation Compact
	}{
		{label: "(0x0300000040)", input: []byte{0x03, 0x00, 0x00, 0x00, 0x40}, expectation: Compact(1073741824)},
		{label: "(0x0b00407a10f35a)", input: []byte{0x0b, 0x00, 0x40, 0x7a, 0x10, 0xf3, 0x5a}, expectation: Compact(100000000000000)},
		{label: "(0x03ffffffff)", input: []byte{0x03, 0xff, 0xff, 0xff, 0xff}, expectation: Compact(1<<32 - 1)},
		{label: "(0x070000000001)", input: []byte{0x07, 0x00, 0x00, 0x00, 0x00, 0x01}, expectation: Compact(1 << 32)},
		{label: "(0x0b000000000001)", input: []byte{0x0b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, expectation: Compact(1 << 40)},
		{label: "(0x0f00000000000001)", input: []byte{0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, expectation: Compact(1 << 48)},
		{label: "(0x0fffffffffffffff)", input: []byte{0x0f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, expectation: Compact(1<<56 - 1)},
		{label: "(0x130000000000000001)", input: []byte{0x13, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, expectation: Compact(1 << 56)},
		{label: "(0x13ffffffffffffffff)", input: []byte{0x13, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, expectation: Compact(math.MaxUint64)},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := bytes.Buffer{}

			enc := &Encoder{Writer: &buffer}
			testExample.expectation.Encode(enc)

			result := buffer.Bytes()

			assertEqual(t, result, testExample.expectation)
		})
	}
}
