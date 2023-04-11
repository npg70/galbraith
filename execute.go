package main

import "strings"

// Execute parses and renders a tag string
func Execute(src string, fmap map[string]TagFunc) string {
	t := Tokenizer{}
	n := t.Parse(strings.NewReader(src))
	return Render(n, fmap)
}
