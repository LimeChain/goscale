package goscale

// import "math"

// // func (a U8) SaturatingDiv(b U8) U8 {
// // 	if b == 0 {
// // 		return U8(math.MaxUint8)
// // 	}

// // 	return a / b
// // }

// // func (a U16) SaturatingDiv(b U16) U16 {
// // 	if b == 0 {
// // 		return U16(math.MaxUint16)
// // 	}

// // 	return a / b
// // }

// // func (a U32) SaturatingDiv(b U32) U32 {
// // 	if b == 0 {
// // 		return U32(math.MaxUint32)
// // 	}

// // 	return a / b
// // }

// // func (a U64) SaturatingDiv(b U64) U64 {
// // 	if b == 0 {
// // 		return U64(math.MaxUint64)
// // 	}

// // 	return a / b
// // }

func (a U64) CheckedAdd(b U64) (U64, error) {
	c := a + b

	if c < a {
		return 0, ErrOverflow
	}

	return c, nil
}
