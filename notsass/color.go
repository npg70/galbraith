package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func hexToRGB(hexColor string) (string, error) {
	hexColor = strings.TrimPrefix(hexColor, "#")

	// turn #123 --> #112233
	if len(hexColor) == 3 {
		hexColor = string([]byte{hexColor[0], hexColor[0], hexColor[1], hexColor[1], hexColor[2], hexColor[2]})
	} else if len(hexColor) != 6 {
		return "", fmt.Errorf("invalid hex format")
	}

	bytes, err := hex.DecodeString(hexColor)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d,%d,%d", int(bytes[0]), int(bytes[1]), int(bytes[2])), nil
}
