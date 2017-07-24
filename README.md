# flex - flexbox layout implementation in Go

Go implementation of flexbox CSS layout algorithm.

A pure GO port of [Facebook's Yoga](https://github.com/facebook/yoga).

For usage examples look at `_test.go` files.

Logic is currently synced up to  https://github.com/facebook/yoga/commit/c9384762ee367f890a3de57ff3270d8f9c445866

## How the port was made

* manually ported [C code](https://github.com/facebook/yoga/tree/master/yoga) to Go, line-by-line
* manually ported [tests](https://github.com/facebook/yoga/tree/master/tests) to Go
* tweak the API from C style to be more Go like. The structure and logic still is very close to C code (this makes porting future C changes easy)

## Status

The port is finished. The code works and passess all Yoga tests.

The API is awkward by Go standards but it's the best I could do given that I want to stay close to C version.
