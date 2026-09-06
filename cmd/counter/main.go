package main

import (
	"fmt"
	"github.com/npg70/galbraith/base32"
)

func main() {
	u32, id := base32.Counter()
	fmt.Println(u32, id)
}
