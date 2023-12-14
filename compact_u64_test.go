package goscale

import (
	"bytes"
	"math"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EncodeCompactU64(t *testing.T) {
	var examples = []struct {
		label  string
		input  CompactU64
		expect []byte
	}{
		{label: "Encode CompactU64(0)", input: CompactU64(big.NewInt(0).SetUint64(0).Uint64()), expect: []byte{0x00}},
		{label: "Encode CompactU64(1)", input: CompactU64(big.NewInt(0).SetUint64(1).Uint64()), expect: []byte{0x04}},
		{label: "Encode CompactU64(42)", input: CompactU64(big.NewInt(0).SetUint64(42).Uint64()), expect: []byte{0xa8}},
		{label: "Encode CompactU64(63)", input: CompactU64(big.NewInt(0).SetUint64(63).Uint64()), expect: []byte{0xfc}},
		{label: "Encode CompactU64(64)", input: CompactU64(big.NewInt(0).SetUint64(64).Uint64()), expect: []byte{0x01, 0x01}},
		{label: "Encode CompactU64(127)", input: CompactU64(big.NewInt(0).SetUint64(127).Uint64()), expect: []byte{0xfd, 0x01}},
		{label: "Encode CompactU64(65535)", input: CompactU64(big.NewInt(0).SetUint64(65535).Uint64()), expect: []byte{0xfe, 0xff, 0x03, 0x00}},
		{label: "Encode CompactU64(16383)", input: CompactU64(big.NewInt(0).SetUint64(16383).Uint64()), expect: []byte{0xfd, 0xff}},
		{label: "Encode CompactU64(16384)", input: CompactU64(big.NewInt(0).SetUint64(16384).Uint64()), expect: []byte{0x02, 0x00, 0x01, 0x00}},
		{label: "Encode CompactU64(1073741823)", input: CompactU64(big.NewInt(0).SetUint64(1073741823).Uint64()), expect: []byte{0xfe, 0xff, 0xff, 0xff}},
		{label: "Encode CompactU64(1073741824)", input: CompactU64(big.NewInt(0).SetUint64(1073741824).Uint64()), expect: []byte{0x03, 0x00, 0x00, 0x00, 0x40}},
		{label: "Encode CompactU64(100000000000000)", input: CompactU64(big.NewInt(0).SetUint64(100000000000000).Uint64()), expect: []byte{0x0b, 0x00, 0x40, 0x7a, 0x10, 0xf3, 0x5a}},
		{label: "Encode CompactU64(1<<32 - 1)", input: CompactU64(big.NewInt(0).SetUint64(1<<32 - 1).Uint64()), expect: []byte{0x03, 0xff, 0xff, 0xff, 0xff}},
		{label: "Encode CompactU64(1 << 32)", input: CompactU64(big.NewInt(0).SetUint64(1 << 32).Uint64()), expect: []byte{0x07, 0x00, 0x00, 0x00, 0x00, 0x01}},
		{label: "Encode CompactU64(1 << 40)", input: CompactU64(big.NewInt(0).SetUint64(1 << 40).Uint64()), expect: []byte{0x0b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}},
		{label: "Encode CompactU64(1 << 48)", input: CompactU64(big.NewInt(0).SetUint64(1 << 48).Uint64()), expect: []byte{0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}},
		{label: "Encode CompactU64(1<<56 - 1)", input: CompactU64(big.NewInt(0).SetUint64(1<<56 - 1).Uint64()), expect: []byte{0x0f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{label: "Encode CompactU64(1 << 56)", input: CompactU64(big.NewInt(0).SetUint64(1 << 56).Uint64()), expect: []byte{0x13, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}},
		{label: "Encode CompactU64(math.MaxUint64)", input: CompactU64(big.NewInt(0).SetUint64(math.MaxUint64).Uint64()), expect: []byte{0x13, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}}}

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
