package goscale

type Numeric interface {
	Encodable
	U8 | I8 | U16 | I16 | U32 | I32 | U64 | I64 // TODO | U128 | I128
}
