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
