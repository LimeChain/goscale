package goscale

import (
	"math/bits"
)

func (a U8) TrailingZeros() int {
	return bits.TrailingZeros(uint(a))
}

func (a U16) TrailingZeros() int {
	return bits.TrailingZeros(uint(a))
}

func (a U32) TrailingZeros() int {
	return bits.TrailingZeros(uint(a))
}

func (a U64) TrailingZeros() int {
	return bits.TrailingZeros(uint(a))
}
