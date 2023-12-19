package goscale

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type encodableType struct{}

func (e encodableType) Encode(buffer *bytes.Buffer) error {
	return errors.New("error")
}

func (e encodableType) Bytes() []byte {
	return []byte{}
}

func Test_EncodeEach(t *testing.T) {
	buffer := &bytes.Buffer{}

	err := EncodeEach(buffer, Bool(true), U8(127), I16(-128))

	assert.NoError(t, err)
	assert.Equal(t, []byte{0x01, 0x7f, 0x80, 0xff}, buffer.Bytes())
}

func Test_EncodeEach_Error(t *testing.T) {
	buffer := &bytes.Buffer{}

	err := EncodeEach(buffer, Bool(true), encodableType{}, U8(127))

	assert.Error(t, err)
	assert.Equal(t, []byte{0x01}, buffer.Bytes())
}

func Test_ToCompact(t *testing.T) {
	var examples = []struct {
		label  string
		input  interface{}
		expect Compact[BigNumbers]
	}{
		{label: "ToCompact(int)", input: 1, expect: Compact[BigNumbers]{NewU128(1)}},
		{label: "ToCompact(uint)", input: uint(2), expect: Compact[BigNumbers]{NewU128(2)}},
		{label: "ToCompact(int8)", input: int8(3), expect: Compact[BigNumbers]{NewU128(3)}},
		{label: "ToCompact(I8)", input: I8(4), expect: Compact[BigNumbers]{NewU128(4)}},
		{label: "ToCompact(uint8)", input: uint8(5), expect: Compact[BigNumbers]{NewU128(5)}},
		{label: "ToCompact(U8)", input: U8(6), expect: Compact[BigNumbers]{NewU128(6)}},
		{label: "ToCompact(int16)", input: int16(7), expect: Compact[BigNumbers]{NewU128(7)}},
		{label: "ToCompact(I16)", input: I16(8), expect: Compact[BigNumbers]{NewU128(8)}},
		{label: "ToCompact(uint16)", input: uint16(9), expect: Compact[BigNumbers]{NewU128(9)}},
		{label: "ToCompact(U16)", input: U16(10), expect: Compact[BigNumbers]{NewU128(10)}},
		{label: "ToCompact(int32)", input: int32(11), expect: Compact[BigNumbers]{NewU128(11)}},
		{label: "ToCompact(I32)", input: I32(12), expect: Compact[BigNumbers]{NewU128(12)}},
		{label: "ToCompact(uint32)", input: uint32(13), expect: Compact[BigNumbers]{NewU128(13)}},
		{label: "ToCompact(U32)", input: U32(14), expect: Compact[BigNumbers]{NewU128(14)}},
		{label: "ToCompact(int64)", input: int64(15), expect: Compact[BigNumbers]{NewU128(15)}},
		{label: "ToCompact(I64)", input: I64(16), expect: Compact[BigNumbers]{NewU128(16)}},
		{label: "ToCompact(uint64)", input: uint64(17), expect: Compact[BigNumbers]{NewU128(17)}},
		{label: "ToCompact(U64)", input: U64(18), expect: Compact[BigNumbers]{NewU128(18)}},
		{label: "ToCompact(I128)", input: NewI128(19), expect: Compact[BigNumbers]{NewU128(19)}},
		{label: "ToCompact(U128)", input: NewU128(20), expect: Compact[BigNumbers]{NewU128(20)}},
	}

	for _, e := range examples {
		t.Run(e.label, func(t *testing.T) {
			result := ToCompact(e.input)

			assert.Equal(t, e.expect, result)
		})
	}
}

func Test_ToCompact_Panics(t *testing.T) {
	assert.PanicsWithValue(t,
		"invalid numeric type in ToCompact()",
		func() {
			ToCompact([]byte{})
		},
	)
}
