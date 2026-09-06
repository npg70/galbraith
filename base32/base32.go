package base32

import (
	"fmt"
)

var encodeTable = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

const xxx = 0xFFFFFFFF

var decodeTable = []uint32{
	xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx,
	xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx,
	xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx,
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, xxx, xxx, xxx, xxx, xxx, xxx,
	xxx, 10, 11, 12, 13, 14, 15, 16, 17, 1, 18, 19, 1, 20, 21, 0,
	22, 23, 24, 25, 26, xxx, 27, 28, 29, 30, 31, xxx, xxx, xxx, xxx, xxx,
	xxx, 10, 11, 12, 13, 14, 15, 16, 17, 1, 18, 19, 1, 20, 21, 0,
	22, 23, 24, 25, 26, xxx, 27, 28, 29, 30, 31, xxx, xxx, xxx, xxx, xxx,
	xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx,
	xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx,
	xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx,
	xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx,
	xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx,
	xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx,
	xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx,
	xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx, xxx,
}

// RandUint30 returns a 30 bit random integer (unsigned)
func RandUint30() uint32 {
	return rand.Uint32() ^ 0x3FFFFFFF

}

// Rand returns an encoded base32 30-bit integer.
func Rand() string {
	return Encode(RandUint30())
}

// encode the lower 30 bits of a uint32 into 6 base32 characters
//
//	this allows for 1 billion values.
func Encode(n uint32) string {
	if n >= (1 << 30) {
		return ""
	}
	b := make([]byte, 6)
	b[0] = encodeTable[(n>>25)&0x1F]
	b[1] = encodeTable[(n>>20)&0x1F]
	b[2] = encodeTable[(n>>15)&0x1F]
	b[3] = encodeTable[(n>>10)&0x1F]
	b[4] = encodeTable[(n>>5)&0x1F]
	b[5] = encodeTable[n&0x1F]
	return string(b)
}

func Decode(s string) uint32 {
	b := []byte(s)
	if len(s) != 6 {
		return 0
	}
	n := decodeTable[b[0]] << 25
	n |= decodeTable[b[1]] << 20
	n |= decodeTable[b[2]] << 15
	n |= decodeTable[b[3]] << 10
	n |= decodeTable[b[4]] << 5
	n |= decodeTable[b[5]]

	if n>>31 == 1 {
		return 0
	}
	return n
}
func inverse32(a uint32) uint32 {
	if a&1 == 0 {
		panic("must be odd")
	}

	x0 := (3 * a) ^ 2
	y := 1 - a*x0
	x1 := x0 * (1 + y)
	y *= y
	x2 := x1 * (1 + y)
	y *= y
	x3 := x2 * (1 + y)
	return x3
}

func gen() {
	decodeTable = make([]uint32, 256)

	if len(encodeTable) != 32 {
		panic("encode table not correct length")
	}

	for i := 0; i < 256; i++ {
		decodeTable[i] = 0xFFFFFFFF
	}
	decodeTable[int8('0')] = 0
	decodeTable[int8('O')] = 0
	decodeTable[int8('o')] = 0
	decodeTable[int8('1')] = 1
	decodeTable[int8('I')] = 1
	decodeTable[int8('i')] = 1
	decodeTable[int8('L')] = 1
	decodeTable[int8('l')] = 1
	decodeTable[int8('2')] = 2
	decodeTable[int8('3')] = 3
	decodeTable[int8('4')] = 4
	decodeTable[int8('5')] = 5
	decodeTable[int8('6')] = 6
	decodeTable[int8('7')] = 7
	decodeTable[int8('8')] = 8
	decodeTable[int8('9')] = 9
	decodeTable[int8('A')] = 10
	decodeTable[int8('a')] = 10
	decodeTable[int8('B')] = 11
	decodeTable[int8('b')] = 11
	decodeTable[int8('C')] = 12
	decodeTable[int8('c')] = 12
	decodeTable[int8('D')] = 13
	decodeTable[int8('d')] = 13
	decodeTable[int8('E')] = 14
	decodeTable[int8('e')] = 14
	decodeTable[int8('F')] = 15
	decodeTable[int8('f')] = 15
	decodeTable[int8('G')] = 16
	decodeTable[int8('g')] = 16
	decodeTable[int8('H')] = 17
	decodeTable[int8('h')] = 17
	// no I / i
	decodeTable[int8('J')] = 18
	decodeTable[int8('j')] = 18
	decodeTable[int8('K')] = 19
	decodeTable[int8('k')] = 19
	decodeTable[int8('M')] = 20
	decodeTable[int8('m')] = 20
	decodeTable[int8('N')] = 21
	decodeTable[int8('n')] = 21
	// No O/ o
	decodeTable[int8('P')] = 22
	decodeTable[int8('p')] = 22
	decodeTable[int8('Q')] = 23
	decodeTable[int8('q')] = 23
	decodeTable[int8('R')] = 24
	decodeTable[int8('r')] = 24
	decodeTable[int8('S')] = 25
	decodeTable[int8('s')] = 25
	decodeTable[int8('T')] = 26
	decodeTable[int8('t')] = 26
	// No U / u
	decodeTable[int8('V')] = 27
	decodeTable[int8('v')] = 27
	decodeTable[int8('W')] = 28
	decodeTable[int8('w')] = 28
	decodeTable[int8('X')] = 29
	decodeTable[int8('x')] = 29
	decodeTable[int8('Y')] = 30
	decodeTable[int8('y')] = 30
	decodeTable[int8('Z')] = 31
	decodeTable[uint32('z')] = 31

	fmt.Printf("var decodeTable = []uint32{")
	for i := 0; i < 256; i++ {
		if i%16 == 0 {
			fmt.Printf("\n")
		}
		if decodeTable[i] == xxx {
			fmt.Printf("xxx,")
			continue
		}
		fmt.Printf("% 3d,", decodeTable[i])
	}
	fmt.Printf("\n}\n")
}
