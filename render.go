package main

import (
	"fmt"

	"golang.org/x/net/html"
)

type TagFunc func([]string, string) string

func makeTag(tag string) TagFunc {
	return func(args []string, body string) string {
		return fmt.Sprintf("<%s>%s</%s>", tag, body, tag)
	}
}

func makeTagClass(tag string) TagFunc {
	return func(args []string, body string) string {
		return fmt.Sprintf("<%s class=%s>%s</%s>", tag, args[0], body, tag)
	}
}

// pass through div
func makeDivClass(cnames string) TagFunc {
	return func(args []string, body string) string {
		return fmt.Sprintf("<div class=%q>%s</div><!-- %s -->\n", cnames, body, args[0])
	}
}

func Render(n *html.Node, fmap map[string]TagFunc) string {
	switch n.Type {
	case html.TextNode:
		return n.Data
	case html.ElementNode:
		body := ""
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			body += Render(c, fmap)
		}
		if fn, ok := fmap[n.Data]; ok {
			out := fn(toArgv(n), body)
			return out
		}
		// unknown tag ... pass through
		out := "$" + n.Data
		if len(n.Attr) > 0 {
			out += "["
			for _, arg := range n.Attr {
				out += arg.Key
				if len(arg.Val) > 0 {
					out += "="
					out += arg.Val
					out += " "
				}
			}
			out += "]"
		}
		if len(body) > 0 {
			out += "{" + body + "}"
		}
		return "<code>" + out + "</code>"
	default:
		panic("unknown node type")
	}
}
