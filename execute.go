package main

import (
	"strings"

	tf "github.com/client9/tagfunctions"
)

// Execute parses and renders a tag string
func Execute(src string, fmap map[string]tf.TagFunc) string {
	t := tf.Tokenizer{}
	n := t.Parse(strings.NewReader(src))
	return tf.Render(n, fmap)
}
