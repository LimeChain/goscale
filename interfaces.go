package goscale

type UnsignedNumeric interface {
	U8 | U16 | U32 | U64
}

type SignedNumeric interface {
	I8 | I16 | I32 | I64
}

type Numeric interface {
	Encodable
	UnsignedNumeric | SignedNumeric
}
