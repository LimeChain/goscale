# Implementation of SCALE codec in Go with minimal reflection dependency

Not all SCALE primitive types have a corresponding Go type, thus a few custom defined types are introduced to implement the missing SCALE types.

**Run Tests**

```sh
  go test -v
```


## Boolean

* [x] [Done](https://github.com/LimeChain/goscale/blob/master/boolean.go)

| SCALE      | Go                        |
|------------|---------------------------|
| `bool`     | `bool`                    |


## Fixed Length Integers

* [x] [Done](https://github.com/LimeChain/goscale/blob/master/fixed_length.go)

| SCALE      | Go                        |
|------------|---------------------------|
| `i8`       | `int8`                    |
| `u8`       | `uint8`                   |
| `i16`      | `int16`                   |
| `u16`      | `uint16`                  |
| `i32`      | `int32`                   |
| `u32`      | `uint32`                  |
| `i64`      | `int64`                   |
| `u64`      | `uint64`                  |
| `i128`     | `*big.Int`                |
| `u128`     | `*goscale.Uint128`        |


## Length and Compact (Variable Width Integers)

* [x] [Done](https://github.com/LimeChain/goscale/blob/master/length_compact.go)

| SCALE           | Go         |
|-----------------|------------|
| `Compact<u8>`   | `uint`     |
| `Compact<u16>`  | `uint`     |
| `Compact<u32>`  | `uint`     |
| `Compact<u64>`  | `uint`     |
| `Compact<u128>` | `*big.Int` |


## Sequence

* [x] [Done](https://github.com/LimeChain/goscale/blob/master/sequence.go)

| SCALE      | Go                        |
|------------|---------------------------|
| `bytes`    | `[]byte`                  |
| `string`   | `string`                  |


## Tuple

Go structs are encoded as SCALE Tuple, where each struct field is encoded in a sequence containing all the fields.
To decode SCALE encoded structs, it is required to have prior knowledge of the destination data type, since it is not self-contained in the SCALE encoded data itself.

* [x] [Done](https://github.com/LimeChain/goscale/blob/master/tuple.go)

| SCALE      | Go                        |
|------------|---------------------------|
| `struct`   | `struct`                  |

### Struct Tags

Struct tags allow modification of the struct field ordering, that preserve the encoding/decoding order.

* [ ] Done


## Option & Result

For all `Option<T>` a pointer to the underlying type is used.

* [x] [Done](https://github.com/LimeChain/goscale/blob/master/option.go)

| SCALE              | Go                       |
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
