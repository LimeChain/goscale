# Implementation of SCALE codec in Go with minimal reflection dependency

The SCALE types in Go are represented by a set of custom-defined types that implement the `Encode` method. There is also a `Decode` function for each type and the type to which data should be decoded is inferred by the context (not self-contained in the encoded data).


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
| `i128`     | `*big.Int` (TODO)         |
| `u128`     | `goscale.U128`            |


## [Length and Compact (Variable Width Integers)](https://github.com/LimeChain/goscale/blob/master/length_compact.go)

| SCALE/Rust      | Go                |
|-----------------|-------------------|
| `Compact<u8>`   | `goscale.Compact` |
| `Compact<u16>`  | `goscale.Compact` |
| `Compact<u32>`  | `goscale.Compact` |
| `Compact<u64>`  | `goscale.Compact` |
| `Compact<u128>` | `*big.Int` (TODO) |

## [VaryingData](https://github.com/LimeChain/goscale/blob/master/varying_data.go)

| SCALE/Rust                   | Go                    |
|------------------------------|-----------------------|
| `Enumeration(tagged-union)`  | `goscale.VaryingData` |

## [Sequence](https://github.com/LimeChain/goscale/blob/master/sequence.go)

| SCALE/Rust | Go                          |
|------------|-----------------------------|
| `bytes`    | `goscale.Sequence[U8]`      |
| `string`   | `goscale.Sequence[U8]`      |
| `[u8; u8]` | `goscale.FixedSequence[U8]` |


## [Empty](https://github.com/LimeChain/goscale/blob/master/empty.go)

| SCALE/Rust         | Go                       |
| ------------------ | ------------------------ |


## [Option](https://github.com/LimeChain/goscale/blob/master/option.go)

| SCALE/Rust         | Go                       |
| ------------------ | ------------------------ |
| `Option<i8>`       | `*int8`                  |
| `Option<u8>`       | `*uint8`                 |
| `Option<i16>`      | `*int16`                 |
| `Option<u16>`      | `*uint16`                |
| `Option<i32>`      | `*int32`                 |
| `Option<u32>`      | `*uint32`                |
| `Option<i64>`      | `*int64`                 |
| `Option<u64>`      | `*uint64`                |
| `Option<bytes>`    | `*[]byte`                |
| `None`             | `nil`                    |


## [Result](https://github.com/LimeChain/goscale/blob/master/result.go)

| SCALE/Rust         | Go                       |
| ------------------ | ------------------------ |


## [Dictionary](https://github.com/LimeChain/goscale/blob/master/dictionary.go)

| SCALE/Rust         | Go                       |
| ------------------ | ------------------------ |


### Run Tests

```sh
  go test -v
```
