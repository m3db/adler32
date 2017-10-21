# adler32

Implements adler32 checksum function as described here: https://www.ietf.org/rfc/rfc1950.txt.

## Example:

      adler32.Checksum([]byte("Hello World"))

## Tests

      go test
      go test -bench=.

      # This library is slightly faster than the one in standard library.
      $ go test -bench=.
      BenchmarkThis-4            10000            230169 ns/op
      BenchmarkStdLib-4          10000            190834 ns/op
      PASS
      ok      github.com/sent-hil/adler32     6.554s
