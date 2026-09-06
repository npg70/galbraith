package base32

const p = 0x2C1B3C6D
const pi = 0x24ea2d65
const mask = 0x3FFFFFFF

func Scramble(n uint32) uint32 {
	if n > mask {
		panic("input too large")
	}
	return (n * p) & mask
}

func Unscramble(n uint32) uint32 {
	if n > mask {
		panic("input too large")
	}
	return (n * pi) & mask
}
