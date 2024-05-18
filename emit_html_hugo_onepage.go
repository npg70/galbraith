package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tf "github.com/client9/tagfunctions"
)

func renderFuncs() map[string]tf.TagFunc {
	return map[string]tf.TagFunc{
		"ppre":   tf.MakeTagClass("p", "white-space-pre-line"),
		"strike": tf.MakeTag("s"),
		"br": func(args []string, body string) string {
			return "<br>"
		},
		"root": func(args []string, body string) string {
			return body
		},
		"section": tf.MakeTagClass("section", "mb-4"),

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
		"intro":  tf.MakeTagClass("p", ""),
		"nowrap": tf.MakeTagClass("span", "text-nowrap"),
		"csvtable": NewCsvTableHTML(func(tag string, row int, col int) string {
			switch tag {
			case "wrapper":
				return "table-responsive"
			case "table":
				return "table table-borderless table-sm small ms-3"
			case "th", "td":
				return "text-body"
			}
			return ""
		}),
		"date":               tf.MakeTagClass("span", "text-nowrap"),
		"child-list":         tf.MakeTagClass("table", "table-p0 mb-3"),
		"gen":                tf.MakeTag("sup"),
		"children":           tf.MakeTag("div"),
		"child-partner-name": tf.MakeTagClass("span", "text-smallcaps text-nowrap"),
		"child-name":         tf.MakeTagClass("span", "text-smallcaps text-nowrap"),
		"primary-name":       tf.MakeTagClass("span", "fw-bold text-smallcaps text-nowrap"),
		"partner-name":       tf.MakeTagClass("span", "fw-bold text-smallcaps text-nowrap"),
		"lineage-name":       tf.MakeTag("span"),
		"children-intro":     tf.MakeTag("p"),
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
			out.WriteString("<tr><th>Lineage</th><td>")
			out.WriteString(body)
			out.WriteString("</ol></td></tr>\n")
			return out.String()
		},
		"ancestor": func(args []string, body string) string {
			//genNumber := getKey(args, "generation")
			//counter := tf.GetKeyValue(args, "counter")
			mother := tf.GetKeyValue(args, "mother")
			year := tf.GetKeyValue(args, "year")
			if year != "" {
				year = " b. " + year + " "
			}
			return fmt.Sprintf("<div>%s%sm. %s</div>\n", body, year, mother)
		},
		"person": func(args []string, body string) string {
			pid := tf.GetKeyValue(args, "id")
			s := strings.Builder{}
			s.WriteString("<div class='mb-5'>")
			s.WriteString(fmt.Sprintf("<a name=%q></a>\n", pid))
			s.WriteString(body)
			s.WriteString("</div><!-- person -->\n") // person
			return s.String()
		},
		"person-body": func(args []string, body string) string {
			return "<div>" + body + "</div>\n"
		},
		"person-main":      tf.MakeTagClass("div", "print-hack"),
		"person-secondary": tf.MakeTagClass("table", "small"),
		"banner": func(args []string, body string) string {
			return "<div class='mb-4'>" + body + "</div>\n"
		},
		"pagemeta": func(args []string, body string) string {
			return fmt.Sprintf("<tr class='pb-3'><th class='pe-5'>Updated</th><td>%s</td></tr>",
				body)
		},
		"person-bio": tf.MakeTag("div"),
		"headline":   tf.MakeTag("h1"),
		"externals": func(args []string, body string) string {
			if strings.TrimSpace(body) == "" {
				return ""
			}
			return fmt.Sprintf("<tr><th class='pe-5'>External</th><td>%s</td></tr>\n", body)
		},
		"external": func(args []string, body string) string {
			// open external links into new window
			return fmt.Sprintf("<a class='me-3' rel='noreferrer' target='_blank' href=%q>%s</a>\n", args[1], body)
		},
		"tags": func(args []string, body string) string {
			if len(args) == 0 {
				return ""
			}
			s := strings.Builder{}
			s.WriteString("<tr><th class='pe-5'>Tags</th><td>")
			for _, tag := range args[1:] {
				s.WriteString(makeTagButton(nil, tag))
			}
			s.WriteString("</td></tr>\n")
			return s.String()
		},
		"child": func(args []string, body string) string {
			s := strings.Builder{}
			s.WriteString("<tr>")
			s.WriteString("<td class=width-100r>")
			lineageCounter := tf.GetKeyValue(args, "counter")
			if lineageCounter != "" && lineageCounter != "0" {
				s.WriteString(lineageCounter)
			}
			s.WriteString("</td>")
			r, _ := strconv.Atoi(tf.GetKeyValue(args, "birth-order"))
			s.WriteString("<td class='text-end width-200r'>" + toRoman(r) + "." + "</td>")
			s.WriteString("<td class='ps-3'>" + body + "</td>")
			s.WriteString("</tr>\n")
			return s.String()
		},
		"todos": tf.MakeTagClass("ul", "todos"),
		"todo":  tf.MakeTag("li"),
		"notes": tf.MakeTagClass("ul", "notes"),
		"note":  tf.MakeTag("li"),
		"footnotes": func(args []string, body string) string {
			if len(body) == 0 {
				return ""
			}
			s := strings.Builder{}
			s.WriteString("<hr><div>\n")
			s.WriteString("<table class='table-p0 small' >\n")
			s.WriteString(body)
			s.WriteString("</table></div><!-- footnotes -->\n")
			return s.String()
		},
		"footnote": func(args []string, body string) string {
			s := strings.Builder{}
			s.WriteString("<tr class=''>")
			s.WriteString("<td class='width-100r'>")
			s.WriteString("<span>" + args[1] + ".&nbsp;</span>")
			s.WriteString("</td>")
			s.WriteString("<td class=''>")
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
				body = "<blockquote class='fs-100p'>" + body + "</blockquote>"
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
				body = "<blockquote class='fs-100p'>" + body + "</blockquote>"
			}
			person2 := ""
			if len(args) == 4 {
				person2 = args[3]
			}

			return OPRText(args[1], args[2], person2, true) + body
		},
	}
}
