package adler32

var prime uint32 = 65521

func Checksum(buf []byte) uint32 {
	var s1 uint32 = 1 & 0xffff
	var s2 uint32 = (1 >> 16) & 0xffff

	for n := 0; n < len(buf); n++ {
		s1 = (s1 + uint32(buf[n])) % prime
		s2 = (s2 + s1) % prime
	}

	return ((s2 << 16) | s1)
}
