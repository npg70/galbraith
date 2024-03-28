package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func renderFuncs() map[string]TagFunc {
	return map[string]TagFunc{
		"root": func(args []string, body string) string {
			return body
		},
		"section": makeTagClass("section", "mb-4"),

		"front": func(args []string, body string) string {
			return ""
		},
		"elink": func(args []string, body string) string {
			return fmt.Sprintf("<a rel='noreferrer' target='_blank' href=%q>%s</a>", args[1], body)
		},
		"tag-link": func(args []string, body string) string {
			return makeTagButton(nil, body)
		},
		"title": func(args []string, body string) string {
			return "<h1>" + body + "</h1>\n"
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
		"source-link": func(args []string, body string) string {
			return "<a href=/galbraith/sources/" + args[1] + ">" + body + "</a>"
		},
		"child-link": func(args []string, body string) string {
			return "<a class='text-smallcaps' href=/galbraith/people/" + args[1] + ">" + body + "</a>"
		},
		"child-link-plain": func(args []string, body string) string {
			return "<a class=text-body href=/galbraith/people/" + args[1] + ">" + body + "</a>"
		},
		"lineage": func(args []string, body string) string {
			out := strings.Builder{}
			out.WriteString("<tr><th class=col-1>Lineage</th><td>")
			out.WriteString(body)
			out.WriteString("</ol></td></tr>\n")
			return out.String()
		},
		"ancestor": func(args []string, body string) string {
			//genNumber := getKey(args, "generation")
			//counter := getKeyValue(args, "counter")
			mother := getKeyValue(args, "mother")
			year := getKeyValue(args, "year")
			if year != "" {
				year = " b. " + year + " "
			}
			/*
				return fmt.Sprintf("<li value=%s class=''>%s %s</li>\n",
					counter, body, mother)
			*/
			return fmt.Sprintf("<div>%s%sm. %s</div>\n", body, year, mother)
		},
		"person": func(args []string, body string) string {
			pid := getKeyValue(args, "id")
			s := strings.Builder{}
			s.WriteString("<div class='mb-5'>")
			s.WriteString(fmt.Sprintf("<a name=%q></a>\n", pid))
			s.WriteString(body)
			s.WriteString("</div><!-- person -->\n") // person
			return s.String()
		},
		"person-body": func(args []string, body string) string {
			return "<div class='row'><div class='col-12'>" + body + "</div></div>\n"
		},
		"person-main":      makeTagClass("div", "print-hack"),
		"person-secondary": makeTagClass("table", "table table-borderless .table-sm small"),
		"banner": func(args []string, body string) string {
			return "<div class='mb-4'>" + body + "</div>\n"
		},
		"pagemeta": func(args []string, body string) string {
			return fmt.Sprintf("<tr><th class='col-1'>Updated</th><td>%s</td></tr>",
				body)
		},
		"person-bio": makeTag("div"),
		"headline":   makeTag("h1"),
		"externals": func(args []string, body string) string {
			if strings.TrimSpace(body) == "" {
				return ""
			}
			return fmt.Sprintf("<tr><th class=col-1>External</th><td><ul class='ps-0 mb-0' >%s</ul></td></tr>\n", body)
		},
		"external": func(args []string, body string) string {
			// open external links into new window
			return fmt.Sprintf("<li class='d-inline-block me-3'><a rel='noreferrer' target='_blank' href=%q>%s</a></li>\n", args[1], body)
		},
		"tags": func(args []string, body string) string {
			if len(args) == 0 {
				return ""
			}
			s := strings.Builder{}
			s.WriteString("<tr><th class=col-1>Tags</th><td>")
			for _, tag := range args[1:] {
				s.WriteString(makeTagButton(nil, tag))
			}
			s.WriteString("</td></tr>\n")
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
			s.WriteString("<div class='small'>\n")
			s.WriteString("<hr>\n<table>\n")
			s.WriteString(body)
			s.WriteString("</table></div><!-- footnotes -->\n")
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
