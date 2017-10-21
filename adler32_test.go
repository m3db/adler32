package adler32

import (
	"fmt"
	"hash/adler32"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func TestChecksum(t *testing.T) {
	for i := 0; i < 10000; i++ {
		randBytes := generateRandomByte(i)
		if Checksum(randBytes) != adler32.Checksum(randBytes) {
			fmt.Println(len(randBytes))
			t.FailNow()
		}
	}
}

func BenchmarkThis(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Checksum(generateRandomByte(n))
	}
}

func BenchmarkStdLib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		adler32.Checksum(generateRandomByte(n))
	}
}

func generateRandomByte(n int) []byte {
	b := make([]byte, n)
	for i := 0; i < len(b); i++ {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}
