package main

import (
	"fmt"
	"strings"

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

func IsHTMLTag(tag string) bool {
	// todo make map
	switch tag {
	case "p", "b", "i", "em", "br", "hr":
		return true
	case "table", "tbody", "th", "tr", "td", "tfoot":
		return true
	case "blockquote", "pre":
		return true
	case "ul", "ol", "li":
		return true
	}
	return false
}
func HTMLTag(n *html.Node, body string) string {
	// hack for now
	out := "<" + n.Data
	for _, a := range n.Attr {
		out += " " + a.Key + "=" + fmt.Sprintf("%q", a.Val)
	}
	out += fmt.Sprintf(">%s</%s>", body, n.Data)
	return out
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
		if IsHTMLTag(n.Data) {
			return HTMLTag(n, body)
		}
		// unknown tag ... pass through
		out := "$" + n.Data
		if len(n.Attr) > 0 {
			out += "["
			htmlargs := []string{}
			for _, arg := range n.Attr {
				if len(arg.Val) == 0 {
					// TODO: better quoting
					if strings.Contains(arg.Key, " ") {
						htmlargs = append(htmlargs, fmt.Sprintf("%q", arg.Key))
						continue
					}
					htmlargs = append(htmlargs, arg.Key)
					continue
				}
				// TODO: better quoting
				if strings.Contains(arg.Val, " ") {
					htmlargs = append(htmlargs, fmt.Sprintf("%s=%q", arg.Key, arg.Val))
					continue
				}
				htmlargs = append(htmlargs, arg.Key+"="+arg.Val)
			}
			out += strings.Join(htmlargs, " ")
			out += "]"
		}
		body = strings.TrimSpace(body)
		if len(body) > 0 {
			out += "{" + body + "}"
		}
		return "<code>" + out + "</code>"
	default:
		panic("unknown node type")
	}
}
