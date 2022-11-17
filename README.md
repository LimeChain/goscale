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