package base32

import (
	"log"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
)

var tCounter atomic.Uint32

func TransientCounter() uint32 {
	out := (&tCounter).Add(1)
	if out == 0x0 {
		panic("transient counter overflowed")
	}
	return out
}

var counterFile = "counter.txt"

func Counter() (uint32, string) {

	data, err := os.ReadFile(counterFile)
	if err != nil {
		// Log the error and terminate the program if file reading fails
		log.Fatal(err)
	}
	s := strings.TrimSpace(string(data))
	u64, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	u64 += 1

	u32 := uint32(u64)
	id := Encode(Scramble(u32))
	if Unscramble(Decode(id)) != u32 {
		panic("Round trip failed")
	}

	out := strconv.FormatUint(u64, 10) + "\n"
	if err := os.WriteFile(counterFile, []byte(out), 0644); err != nil {
		log.Fatal(err)
	}

	return u32, id
}
