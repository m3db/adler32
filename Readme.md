# adler32

Port of adler32 checksum function as described here: https://www.ietf.org/rfc/rfc1950.txt to Go.

## Example:

```go
 adler32.Checksum([]byte("Hello World"))
```

## Tests

```bash
$ go test
PASS
ok      github.com/sent-hil/adler32     2.429s

$ go test -bench=.
# This library is slightly faster than the one in standard library.
$ go test -bench=.
BenchmarkThis-4            10000            230169 ns/op
BenchmarkStdLib-4          10000            190834 ns/op
PASS
ok      github.com/sent-hil/adler32     6.554s
```
