package goscale

import (
	"bytes"
	"io"
	"math"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EncodeCompact(t *testing.T) {
	var examples = []struct {
		label  string
		input  Compact
		expect []byte
	}{
		{label: "Encode Compact(0)", input: Compact{NewU128(big.NewInt(0).SetUint64(0))}, expect: []byte{0x00}},
		{label: "Encode Compact(1)", input: Compact{NewU128(big.NewInt(0).SetUint64(1))}, expect: []byte{0x04}},
		{label: "Encode Compact(42)", input: Compact{NewU128(big.NewInt(0).SetUint64(42))}, expect: []byte{0xa8}},
		{label: "Encode Compact(63)", input: Compact{NewU128(big.NewInt(0).SetUint64(63))}, expect: []byte{0xfc}},
		{label: "Encode Compact(64)", input: Compact{NewU128(big.NewInt(0).SetUint64(64))}, expect: []byte{0x01, 0x01}},
		{label: "Encode Compact(127)", input: Compact{NewU128(big.NewInt(0).SetUint64(127))}, expect: []byte{0xfd, 0x01}},
		{label: "Encode Compact(65535)", input: Compact{NewU128(big.NewInt(0).SetUint64(65535))}, expect: []byte{0xfe, 0xff, 0x03, 0x00}},
		{label: "Encode Compact(16383)", input: Compact{NewU128(big.NewInt(0).SetUint64(16383))}, expect: []byte{0xfd, 0xff}},
		{label: "Encode Compact(16384)", input: Compact{NewU128(big.NewInt(0).SetUint64(16384))}, expect: []byte{0x02, 0x00, 0x01, 0x00}},
		{label: "Encode Compact(1073741823)", input: Compact{NewU128(big.NewInt(0).SetUint64(1073741823))}, expect: []byte{0xfe, 0xff, 0xff, 0xff}},
		{label: "Encode Compact(1073741824)", input: Compact{NewU128(big.NewInt(0).SetUint64(1073741824))}, expect: []byte{0x03, 0x00, 0x00, 0x00, 0x40}},
		{label: "Encode Compact(100000000000000)", input: Compact{NewU128(big.NewInt(0).SetUint64(100000000000000))}, expect: []byte{0x0b, 0x00, 0x40, 0x7a, 0x10, 0xf3, 0x5a}},
		{label: "Encode Compact(1<<32 - 1)", input: Compact{NewU128(big.NewInt(0).SetUint64(1<<32 - 1))}, expect: []byte{0x03, 0xff, 0xff, 0xff, 0xff}},
		{label: "Encode Compact(1 << 32)", input: Compact{NewU128(big.NewInt(0).SetUint64(1 << 32))}, expect: []byte{0x07, 0x00, 0x00, 0x00, 0x00, 0x01}},
		{label: "Encode Compact(1 << 40)", input: Compact{NewU128(big.NewInt(0).SetUint64(1 << 40))}, expect: []byte{0x0b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}},
		{label: "Encode Compact(1 << 48)", input: Compact{NewU128(big.NewInt(0).SetUint64(1 << 48))}, expect: []byte{0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}},
		{label: "Encode Compact(1<<56 - 1)", input: Compact{NewU128(big.NewInt(0).SetUint64(1<<56 - 1))}, expect: []byte{0x0f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{label: "Encode Compact(1 << 56)", input: Compact{NewU128(big.NewInt(0).SetUint64(1 << 56))}, expect: []byte{0x13, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}},
		{label: "Encode Compact(math.MaxUint64)", input: Compact{NewU128(big.NewInt(0).SetUint64(math.MaxUint64))}, expect: []byte{0x13, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{label: "Encode Compact(math.MaxUint64, math.MaxUint64)", input: Compact{U128{math.MaxUint64, math.MaxUint64}}, expect: []byte{0x33, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			err := e.input.Encode(buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, buffer.Bytes())
			assert.Equal(t, e.expect, e.input.Bytes())
		})
	}
}

func Test_DecodeCompact(t *testing.T) {
	var examples = []struct {
		label  string
		input  []byte
		expect Compact
	}{
		{label: "Decode Compact(0)", input: []byte{0x00}, expect: Compact{interface{}(NewU128(0)).(U128)}},
		{label: "Decode Compact(1)", input: []byte{0x04}, expect: Compact{NewU128(big.NewInt(0).SetUint64(1))}},
		{label: "Decode Compact(42)", input: []byte{0xa8}, expect: Compact{NewU128(big.NewInt(0).SetUint64(42))}},
		{label: "Decode Compact(63)", input: []byte{0xfc}, expect: Compact{NewU128(big.NewInt(0).SetUint64(63))}},
		{label: "Decode Compact(64)", input: []byte{0x01, 0x01}, expect: Compact{NewU128(big.NewInt(0).SetUint64(64))}},
		{label: "Decode Compact(127)", input: []byte{0xfd, 0x01}, expect: Compact{NewU128(big.NewInt(0).SetUint64(127))}},
		{label: "Decode Compact(65535)", input: []byte{0xfe, 0xff, 0x03, 0x00}, expect: Compact{NewU128(big.NewInt(0).SetUint64(65535))}},
		{label: "Decode Compact(16383)", input: []byte{0xfd, 0xff}, expect: Compact{NewU128(big.NewInt(0).SetUint64(16383))}},
		{label: "Decode Compact(16384)", input: []byte{0x02, 0x00, 0x01, 0x00}, expect: Compact{NewU128(big.NewInt(0).SetUint64(16384))}},
		{label: "Decode Compact(1073741823)", input: []byte{0xfe, 0xff, 0xff, 0xff}, expect: Compact{NewU128(big.NewInt(0).SetUint64(1073741823))}},
		{label: "Decode Compact(1073741824)", input: []byte{0x03, 0x00, 0x00, 0x00, 0x40}, expect: Compact{NewU128(big.NewInt(0).SetUint64(1073741824))}},
		{label: "Decode Compact(100000000000000)", input: []byte{0x0b, 0x00, 0x40, 0x7a, 0x10, 0xf3, 0x5a}, expect: Compact{NewU128(big.NewInt(0).SetUint64(100000000000000))}},
		{label: "Decode Compact(1<<32 - 1)", input: []byte{0x03, 0xff, 0xff, 0xff, 0xff}, expect: Compact{NewU128(big.NewInt(0).SetUint64(1<<32 - 1))}},
		{label: "Decode Compact(1 << 32)", input: []byte{0x07, 0x00, 0x00, 0x00, 0x00, 0x01}, expect: Compact{NewU128(big.NewInt(0).SetUint64(1 << 32))}},
		{label: "Decode Compact(1 << 40)", input: []byte{0x0b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, expect: Compact{NewU128(big.NewInt(0).SetUint64(1 << 40))}},
		{label: "Decode Compact(1 << 48)", input: []byte{0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, expect: Compact{NewU128(big.NewInt(0).SetUint64(1 << 48))}},
		{label: "Decode Compact(1<<56 - 1)", input: []byte{0x0f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, expect: Compact{NewU128(big.NewInt(0).SetUint64(1<<56 - 1))}},
		{label: "Decode Compact(1 << 56)", input: []byte{0x13, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, expect: Compact{NewU128(big.NewInt(0).SetUint64(1 << 56))}},
		{label: "Decode Compact(math.MaxUint64)", input: []byte{0x13, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, expect: Compact{NewU128(big.NewInt(0).SetUint64(math.MaxUint64))}},
		{label: "Decode Compact(math.MaxUint64, math.MaxUint64)", input: []byte{0x33, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, expect: Compact{U128{math.MaxUint64, math.MaxUint64}}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeCompact[U128](buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, result)
		})
	}
}

func Test_DecodeCompact_StronglyTyped_U128(t *testing.T) {
	var examplesU128 = []struct {
		label  string
		input  []byte
		expect Compact
	}{
		{label: "Decode Compact(42)  Mode 0", input: []byte{0xa8}, expect: Compact{NewU128(42)}},
		{label: "Decode Compact(127) Mode 1", input: []byte{0xfd, 0x01}, expect: Compact{NewU128(127)}},
		{label: "Decode Compact(16384) Mode 2", input: []byte{0x02, 0x00, 0x01, 0x00}, expect: Compact{NewU128(16384)}},
		{label: "Decode Compact(math.MaxUint64, math.MaxUint64) Mode 3", input: []byte{0x33, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, expect: Compact{U128{math.MaxUint64, math.MaxUint64}}},
	}

	for _, e := range examplesU128 {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeCompact[U128](buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, result)
		})
	}
}

func Test_DecodeCompact_StronglyTyped_U64(t *testing.T) {
	var examplesU64 = []struct {
		label  string
		input  []byte
		expect Compact
	}{
		{label: "Decode Compact(42)  Mode 0", input: []byte{0xa8}, expect: Compact{NewU64(42)}},
		{label: "Decode Compact(127) Mode 1", input: []byte{0xfd, 0x01}, expect: Compact{NewU64(127)}},
		{label: "Decode Compact(16384) Mode 2", input: []byte{0x02, 0x00, 0x01, 0x00}, expect: Compact{NewU64(16384)}},
		{label: "Decode Compact(MaxUint64) Mode 3", input: []byte{0x13, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, expect: Compact{NewU64(math.MaxUint64)}},
	}

	for _, e := range examplesU64 {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeCompact[U64](buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, result)
		})
	}
}

func Test_DecodeCompact_StronglyTyped_U32(t *testing.T) {
	var examplesU32 = []struct {
		label  string
		input  []byte
		expect Compact
	}{
		{label: "Decode Compact(42)  Mode 0", input: []byte{0xa8}, expect: Compact{NewU32(42)}},
		{label: "Decode Compact(127) Mode 1", input: []byte{0xfd, 0x01}, expect: Compact{NewU32(127)}},
		{label: "Decode Compact(16384) Mode 2", input: []byte{0x02, 0x00, 0x01, 0x00}, expect: Compact{NewU32(16384)}},
		{label: "Decode Compact(MaxUint64) Mode 3", input: []byte{0x03, 0xff, 0xff, 0xff, 0xff}, expect: Compact{NewU32(1<<32 - 1)}},
	}

	for _, e := range examplesU32 {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeCompact[U32](buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, result)
		})
	}
}

func Test_DecodeCompact_StronglyTyped_U16(t *testing.T) {
	var examplesU16 = []struct {
		label  string
		input  []byte
		expect Compact
	}{
		{label: "Decode Compact(42)  Mode 0", input: []byte{0xa8}, expect: Compact{NewU16(42)}},
		{label: "Decode Compact(127) Mode 1", input: []byte{0xfd, 0x01}, expect: Compact{NewU16(127)}},
		{label: "Decode Compact(16384) Mode 2", input: []byte{0x02, 0x00, 0x01, 0x00}, expect: Compact{NewU16(16384)}},
	}

	for _, e := range examplesU16 {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeCompact[U16](buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, result)
		})
	}
}

func Test_DecodeCompact_StronglyTyped_U8(t *testing.T) {
	var examplesU8 = []struct {
		label  string
		input  []byte
		expect Compact
	}{
		{label: "Decode Compact(42)  Mode 0", input: []byte{0xa8}, expect: Compact{NewU8(42)}},
		{label: "Decode Compact(127) Mode 1", input: []byte{0xfd, 0x01}, expect: Compact{NewU8(127)}},
	}

	for _, e := range examplesU8 {
		t.Run(e.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(e.input)

			result, err := DecodeCompact[U8](buffer)

			assert.NoError(t, err)
			assert.Equal(t, e.expect, result)
		})
	}
}

func Test_DecodeCompact_Empty(t *testing.T) {
	buffer := &bytes.Buffer{}

	result, err := DecodeCompact[U128](buffer)

	assert.Equal(t, io.EOF, err)
	assert.Equal(t, Compact{}, result)
}
