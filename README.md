# SCALE codec implementation in Go compatible with a Tinygo based toolchain

[![codecov](https://codecov.io/gh/LimeChain/goscale/branch/master/graph/badge.svg?token=3OQPC6VNMB)](https://codecov.io/gh/LimeChain/goscale)

The SCALE types in Go are represented by a set of custom-defined types that implement the `Encodable` interface. 
Each type also has a corresponding decode function using the convention `Decode<TypeName>`. Note that the type to which data should be decoded is inferred from the context and is not self-contained in the SCALE-encoded data.

One exception is the `Tuple` type. It doesn't have methods attached. Instead, there are `EncodeTuple` and `DecodeTuple` functions that can be invoked with any custom struct that embeds the `Tuple` interface.

Some quirks deserve mention. For example, the `FixedSequence` type, which has the same representation as the `Sequence` type, facilitates the encoding of arrays. As arrays are fixed-size sequences, they cannot be encoded as the `Sequence` type. Note that there are no type checks on the size.

The use of custom-defined types and generics reduces reliance on reflection, which isn't fully supported by TinyGo.

---

## [Boolean](https://github.com/LimeChain/goscale/blob/master/boolean.go)

| SCALE/Rust | Go                        |
|------------|---------------------------|
| `bool`     | `goscale.Bool`            |


## [Fixed Length Integers](https://github.com/LimeChain/goscale/blob/master/fixed_length.go)

| SCALE/Rust | Go                        |
|------------|---------------------------|
| `i8`       | `goscale.I8`              |
| `u8`       | `goscale.U8`              |
| `i16`      | `goscale.I16`             |
| `u16`      | `goscale.U16`             |
| `i32`      | `goscale.I32`             |
| `u32`      | `goscale.U32`             |
| `i64`      | `goscale.I64`             |
| `u64`      | `goscale.U64`             |
| `i128`     | `goscale.I128`            |
| `u128`     | `goscale.U128`            |


## [Length and Compact (Variable Width Integers)](https://github.com/LimeChain/goscale/blob/master/length_compact.go)

| SCALE/Rust      | Go                |
|-----------------|-------------------|
| `Compact<u8>`   | `goscale.Compact` |
| `Compact<u16>`  | `goscale.Compact` |
| `Compact<u32>`  | `goscale.Compact` |
| `Compact<u64>`  | `goscale.Compact` |
| `Compact<u128>` | `goscale.Compact` |


## [Sequence](https://github.com/LimeChain/goscale/blob/master/sequence.go)

| SCALE/Rust | Go                          |
|------------|-----------------------------|
| `bytes`    | `goscale.Sequence[U8]`      |
| `[u8; u8]` | `goscale.FixedSequence[U8]` |
| `string`   | `goscale.Str`               |


## [Dictionary](https://github.com/LimeChain/goscale/blob/master/dictionary.go)

| SCALE/Rust | Go                         |
|------------|----------------------------|
|            | `goscale.Dictionary[K, V]` |


## [Empty](https://github.com/LimeChain/goscale/blob/master/empty.go)

| SCALE/Rust | Go              |
|------------|-----------------|
|            | `goscale.Empty` |


## [VaryingData](https://github.com/LimeChain/goscale/blob/master/varying_data.go)

| SCALE/Rust                   | Go                    |
|------------------------------|-----------------------|
| `Enumeration(tagged-union)`  | `goscale.VaryingData` |
|                              |                       |      


## [Option](https://github.com/LimeChain/goscale/blob/master/option.go)

| SCALE/Rust      | Go                     |
|-----------------|------------------------|
| `Option<bool>`  | `Option[goscale.Bool]` |
| `Option<i8>`    | `Option[goscale.I8]`   |
| `Option<u8>`    | `Option[goscale.U8]`   |
| `Option<i16>`   | `Option[goscale.I16]`  |
| `Option<u16>`   | `Option[goscale.U16]`  |
| `Option<i32>`   | `Option[goscale.I32]`  |
| `Option<u32>`   | `Option[goscale.U32]`  |
| `Option<i64>`   | `Option[goscale.I64]`  |
| `Option<u64>`   | `Option[goscale.U64]`  |
| `Option<i128>`  | `Option[goscale.I128]` |
| `Option<u128>`  | `Option[goscale.U128]` |
| `Option<bytes>` | `Option[Sequence[U8]]` |
| `OptionBool`    | `OptionBool`           |
| `None`          | `nil`                  |


## [Result](https://github.com/LimeChain/goscale/blob/master/result.go)

| SCALE/Rust                     | Go                                |
|--------------------------------|-----------------------------------|
| Result<EncodeLike, EncodeLike> | goscale.Result[goscale.Encodable] |


## [Tuple](https://github.com/LimeChain/goscale/blob/master/tuple.go)

Go structs are encoded as SCALE Tuple, where each struct field is encoded in a sequence containing all the fields.
To decode SCALE encoded structs, it is required to have prior knowledge of the destination data type.

| SCALE      | Go                        |
|------------|---------------------------|
| `struct`   | `goscale.Tuple`           |


### Run Tests

```sh
go test -v
```