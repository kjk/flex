# flex - flexbox layout implementation in Go

Go implementation of flexbox, port of [Facebook's Yoga](https://github.com/facebook/yoga) to Go.

Logic is currently synced up to  https://github.com/facebook/yoga/commit/c9384762ee367f890a3de57ff3270d8f9c445866

## How the port was made

* manually ported [C code](https://github.com/facebook/yoga/tree/master/yoga) to Go, line-by-line
* manually ported [tests](https://github.com/facebook/yoga/tree/master/tests) to Go
* tweak the API from C style to be more Go like. The structure and logic still is very close to C code (this makes porting C changes easier)

## Status

Unstable.

The code works and passes all the tests but I'm still tweaking the API to be more Go-like
