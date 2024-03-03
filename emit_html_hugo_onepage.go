package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func renderFuncs() map[string]TagFunc {
	return map[string]TagFunc{
		"root": func(args []string, body string) string {
			return body
		},
		"front": func(args []string, body string) string {
			return ""
		},
		"elink": func(args []string, body string) string {
			return fmt.Sprintf("<a rel='noreferrer' target='_blank' href=%q>%s</a>", args[1], body)
		},
		"tag-link": func(args []string, body string) string {
			tag := strings.ToLower(body)
			taglink := "/galbraith/tags/" + strings.ReplaceAll(tag, " ", "-") + "/"
			return "<a class='btn btn-sm btn-secondary' href=" + taglink + ">" + TitleTag(tag) + "</a>\n"
		},
		"title": func(args []string, body string) string {
			return "title: " + body + "\n"
		},
		"intro":              makeTagClass("p", ""),
		"nowrap":             makeTagClass("span", "text-nowrap"),
		"csvtable":           CsvTableHTML,
		"date":               makeTagClass("span", "text-nowrap"),
		"child-list":         makeTagClass("table", "mb-3"),
		"gen":                makeTag("sup"),
		"children":           makeTag("div"),
		"child-partner-name": makeTagClass("span", "text-smallcaps text-nowrap"),
		"child-name":         makeTagClass("span", "text-smallcaps text-nowrap"),
		"primary-name":       makeTagClass("span", "fw-bold text-smallcaps text-nowrap"),
		"partner-name":       makeTagClass("span", "fw-bold text-smallcaps text-nowrap"),
		"lineage-name":       makeTag("span"),
		"children-intro":     makeTag("p"),
		"primary-number": func(args []string, body string) string {
			return "<span class='fw-bold pe-3'>" + body + ". </span>"
		},
		"child-link": func(args []string, body string) string {
			return "<a href=/galbraith/people/" + args[1] + ">" + body + "</a>"
		},
		"lineage": func(args []string, body string) string {
			out := strings.Builder{}
			out.WriteString("<div class='mb-3'>")
			out.WriteString("<h6>Lineage</h6>")
			out.WriteString("<ol class='m-0'>")
			out.WriteString(body)
			out.WriteString("</ol>")
			out.WriteString("</div><!-- lineage -->\n")
			return out.String()
		},
		"ancestor": func(args []string, body string) string {
			//genNumber := getKey(args, "generation")
			counter := getKeyValue(args, "counter")
			mother := getKeyValue(args, "mother")
			return fmt.Sprintf("<li value=%s>%s<br>%s</li>\n",
				counter, body, mother)
		},
		"person": func(args []string, body string) string {
			pid := getKeyValue(args, "id")
			s := strings.Builder{}
			s.WriteString("<div class='container-fluid mx-auto mb-5'>")
			s.WriteString(fmt.Sprintf("<a name=%q></a>\n", pid))
			s.WriteString(body)
			s.WriteString("</div><!-- person -->\n") // person
			return s.String()
		},
		"person-body":      makeTagClass("div", "row"),
		"person-main":      makeTagClass("div", "col-12 col-md-9 col-order-first print-hack"),
		"person-secondary": makeTagClass("div", "person-secondary col-12 col-md-3 order-last small"),
		"banner": func(args []string, body string) string {
			return "<div class='row'><div class='col-12 mb-4'>" + body + "</div></div>\n"
		},
		"pagemeta": func(args []string, body string) string {
			return fmt.Sprintf("<div class=%s><h6>Meta</h6>%s</div><!-- %s -->\n",
				args[0], body, args[0])
		},
		"person-bio": makeTag("div"),
		"headline":   makeTag("h1"),
		"externals": func(args []string, body string) string {
			if strings.TrimSpace(body) == "" {
				return ""
			}
			return fmt.Sprintf("<div class=%s><h6>External</h6><ul class='list-unstyled'>%s</ul></div><!-- %s -->\n", args[0], body, args[0])
		},
		"external": func(args []string, body string) string {
			// open external links into new window
			return fmt.Sprintf("<li><a rel='noreferrer' target='_blank' href=%q>%s</a></li>\n", args[1], body)
		},
		"tags": func(args []string, body string) string {
			if len(args) == 0 {
				return ""
			}
			s := strings.Builder{}
			s.WriteString("<div class=''>\n")
			for _, tag := range args[1:] {
				taglink := "/galbraith/tags/" + strings.ReplaceAll(strings.ToLower(tag), " ", "-") + "/"
				s.WriteString("<a class='btn btn-sm btn-secondary text-decoration-none' href=" + taglink + ">" + TitleTag(tag) + "</a> ")
			}
			s.WriteString("</div>\n")
			return s.String()
		},
		"child": func(args []string, body string) string {
			s := strings.Builder{}
			s.WriteString("<tr>")
			s.WriteString("<td class=child-counter>")
			lineageCounter := getKeyValue(args, "counter")
			if lineageCounter != "" && lineageCounter != "0" {
				s.WriteString(lineageCounter)
			}
			s.WriteString("</td>")
			r, _ := strconv.Atoi(getKeyValue(args, "birth-order"))
			s.WriteString("<td class='text-end roman'>" + toRoman(r) + "." + "</td>")
			s.WriteString("<td class='ps-3'>" + body + "</td>")
			s.WriteString("</tr>\n")
			return s.String()
		},
		"todos": makeTagClass("ul", "todos"),
		"todo":  makeTag("li"),
		"notes": makeTagClass("ul", "notes"),
		"note":  makeTag("li"),
		"footnotes": func(args []string, body string) string {
			if len(body) == 0 {
				return ""
			}
			s := strings.Builder{}
			s.WriteString("<div class='row'><div class='col-12 small'>\n")
			s.WriteString("<hr>\n<table>\n")
			s.WriteString(body)
			s.WriteString("</table></div></div><!-- footnotes -->\n")
			return s.String()
		},
		"footnote": func(args []string, body string) string {
			s := strings.Builder{}
			s.WriteString("<tr class=break-inside-avoid>")
			s.WriteString("<td class=footnote-counter>")
			s.WriteString("<span><sup>" + args[1] + "</sup>&nbsp;</span>")
			s.WriteString("</td>")
			s.WriteString("<td>")
			s.WriteString(strings.TrimSpace(body))
			s.WriteString("</td>")
			s.WriteString("</tr>\n")
			return s.String()
		},
		"ref": func(args []string, body string) string {
			return "<sup class=footnote-ref>[" + args[1] + "]</sup>"
		},
		"sp-ref": func(args []string, body string) string {
			if len(args) < 3 {
				log.Fatalf("%s: expected at least 3 args, got %v", args[0], args[1:])
			}
			if strings.HasPrefix(args[1], "b") || strings.HasPrefix(args[1], "d") {
				if len(args) != 3 {
					log.Fatalf("%s: Got Birth or Death record, but not 3 args, %v", args[0], args[1:])
				}
			}
			person2 := ""
			if len(args) == 4 {
				person2 = args[3]
			}
			return SPText(args[1], "", args[2], person2)
		},
		"sp-ref-link": func(args []string, body string) string {
			if len(args) < 4 {
				log.Fatalf("Expected 4 or 5 args, got %v", args)
			}
			body = strings.TrimSpace(body)
			if len(body) > 0 {
				body = "<blockquote>" + body + "</blockquote>"
			}
			person2 := ""
			if len(args) == 5 {
				person2 = args[4]
			}
			return SPText(args[1], args[2], args[3], person2) + body
		},

		"opr-ref": func(args []string, body string) string {
			if len(args) < 3 {
				log.Fatalf("%s: expected at least 3 args, got %v", args[0], args[1:])
			}
			if strings.HasPrefix(args[1], "b") || strings.HasPrefix(args[1], "d") {
				if len(args) != 3 {
					log.Fatalf("%s: Got Birth or Death record, but not 3 args, %v", args[0], args[1:])
				}
			}
			if len(args) < 3 || len(args) > 4 {
				log.Fatalf("%s: expected 3 or 4 args got %v", args[0], args[1:])
			}
			person2 := ""
			if len(args) == 4 {
				person2 = args[3]
			}
			return OPRText(args[1], args[2], person2, false)
		},
		"opr-ref-link": func(args []string, body string) string {
			if len(args) < 3 || len(args) > 4 {
				log.Fatalf("%s: expected 3 or 4 args got %v", args[0], args[1:])
			}
			body = strings.TrimSpace(body)
			if len(body) > 0 {
				body = "<blockquote>" + body + "</blockquote>"
			}
			person2 := ""
			if len(args) == 4 {
				person2 = args[3]
			}

			return OPRText(args[1], args[2], person2, true) + body
		},
	}
}

func emitHTMLHugo(n *html.Node, singlePage bool) string {
	front := ""
	if singlePage {
		front += "---\n"
		front += "title: test\n"
		front += "---\n"
	}

	out := Render(n, renderFuncs())
	return front + out
}
