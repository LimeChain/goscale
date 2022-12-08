package goscale

import (
	"bytes"
	"math"
	"math/big"
	"testing"
)

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
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
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
			buffer := &bytes.Buffer{}

			testExample.expectation.Encode(buffer)

			result := DecodeCompact(buffer)

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
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assertEqual(t, buffer.Bytes(), testExample.expectation)
		})
	}
}

func Test_DecodeUintCompact(t *testing.T) {
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
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			result := DecodeCompact(buffer)

			assertEqual(t, result, testExample.expectation)
		})
	}
}

func Test_EncodeCompactU128(t *testing.T) {
	var examples = []struct {
		label  string
		input  CompactU128
		expect []byte
	}{
		{label: "Encode CompactU128(0)", input: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(0))}, expect: []byte{0x00}},
		{label: "Encode CompactU128(1)", input: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(1))}, expect: []byte{0x04}},
		{label: "Encode CompactU128(42)", input: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(42))}, expect: []byte{0xa8}},
		{label: "Encode CompactU128(63)", input: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(63))}, expect: []byte{0xfc}},
		{label: "Encode CompactU128(64)", input: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(64))}, expect: []byte{0x01, 0x01}},
		{label: "Encode CompactU128(127)", input: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(127))}, expect: []byte{0xfd, 0x01}},
		{label: "Encode CompactU128(65535)", input: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(65535))}, expect: []byte{0xfe, 0xff, 0x03, 0x00}},
		{label: "Encode CompactU128(16383)", input: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(16383))}, expect: []byte{0xfd, 0xff}},
		{label: "Encode CompactU128(16384)", input: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(16384))}, expect: []byte{0x02, 0x00, 0x01, 0x00}},
		{label: "Encode CompactU128(1073741823)", input: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(1073741823))}, expect: []byte{0xfe, 0xff, 0xff, 0xff}},
		{label: "Encode CompactU128(1073741824)", input: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(1073741824))}, expect: []byte{0x03, 0x00, 0x00, 0x00, 0x40}},
		{label: "Encode CompactU128(100000000000000)", input: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(100000000000000))}, expect: []byte{0x0b, 0x00, 0x40, 0x7a, 0x10, 0xf3, 0x5a}},
		{label: "Encode CompactU128(1<<32 - 1)", input: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(1<<32 - 1))}, expect: []byte{0x03, 0xff, 0xff, 0xff, 0xff}},
		{label: "Encode CompactU128(1 << 32)", input: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(1 << 32))}, expect: []byte{0x07, 0x00, 0x00, 0x00, 0x00, 0x01}},
		{label: "Encode CompactU128(1 << 40)", input: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(1 << 40))}, expect: []byte{0x0b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}},
		{label: "Encode CompactU128(1 << 48)", input: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(1 << 48))}, expect: []byte{0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}},
		{label: "Encode CompactU128(1<<56 - 1)", input: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(1<<56 - 1))}, expect: []byte{0x0f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{label: "Encode CompactU128(1 << 56)", input: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(1 << 56))}, expect: []byte{0x13, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}},
		{label: "Encode CompactU128(math.MaxUint64)", input: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(math.MaxUint64))}, expect: []byte{0x13, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{label: "Encode CompactU128(math.MaxUint64, math.MaxUint64)", input: CompactU128{U128{math.MaxUint64, math.MaxUint64}}, expect: []byte{0x33, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
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

func Test_Decode_CompactU128(t *testing.T) {
	var examples = []struct {
		label  string
		input  []byte
		expect CompactU128
	}{
		{label: "Decode CompactU128(0)", input: []byte{0x00}, expect: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(0))}},
		{label: "Decode CompactU128(1)", input: []byte{0x04}, expect: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(1))}},
		{label: "Decode CompactU128(42)", input: []byte{0xa8}, expect: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(42))}},
		{label: "Decode CompactU128(63)", input: []byte{0xfc}, expect: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(63))}},
		{label: "Decode CompactU128(64)", input: []byte{0x01, 0x01}, expect: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(64))}},
		{label: "Decode CompactU128(127)", input: []byte{0xfd, 0x01}, expect: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(127))}},
		{label: "Decode CompactU128(65535)", input: []byte{0xfe, 0xff, 0x03, 0x00}, expect: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(65535))}},
		{label: "Decode CompactU128(16383)", input: []byte{0xfd, 0xff}, expect: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(16383))}},
		{label: "Decode CompactU128(16384)", input: []byte{0x02, 0x00, 0x01, 0x00}, expect: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(16384))}},
		{label: "Decode CompactU128(1073741823)", input: []byte{0xfe, 0xff, 0xff, 0xff}, expect: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(1073741823))}},
		{label: "Decode CompactU128(1073741824)", input: []byte{0x03, 0x00, 0x00, 0x00, 0x40}, expect: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(1073741824))}},
		{label: "Decode CompactU128(100000000000000)", input: []byte{0x0b, 0x00, 0x40, 0x7a, 0x10, 0xf3, 0x5a}, expect: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(100000000000000))}},
		{label: "Decode CompactU128(1<<32 - 1)", input: []byte{0x03, 0xff, 0xff, 0xff, 0xff}, expect: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(1<<32 - 1))}},
		{label: "Decode CompactU128(1 << 32)", input: []byte{0x07, 0x00, 0x00, 0x00, 0x00, 0x01}, expect: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(1 << 32))}},
		{label: "Decode CompactU128(1 << 40)", input: []byte{0x0b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, expect: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(1 << 40))}},
		{label: "Decode CompactU128(1 << 48)", input: []byte{0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, expect: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(1 << 48))}},
		{label: "Decode CompactU128(1<<56 - 1)", input: []byte{0x0f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, expect: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(1<<56 - 1))}},
		{label: "Decode CompactU128(1 << 56)", input: []byte{0x13, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, expect: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(1 << 56))}},
		{label: "Decode CompactU128(math.MaxUint64)", input: []byte{0x13, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, expect: CompactU128{NewU128FromBigInt(big.NewInt(0).SetUint64(math.MaxUint64))}},
		{label: "Decode CompactU128(math.MaxUint64, math.MaxUint64)", input: []byte{0x33, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, expect: CompactU128{U128{math.MaxUint64, math.MaxUint64}}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			// given:
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			// when:
			result := DecodeCompactU128(buffer)

			// then:
			assertEqual(t, result, e.expect)
		})
	}
}
