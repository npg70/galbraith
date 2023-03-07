package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"golang.org/x/net/html"
)

func oprindex(db Root) string {
	type pair struct {
		oprid string // b-1807-607-000-1234-1245
		pid   string // galbreath-foo-year-spouse
		sv    string // name of person on record
	}
	pmap := map[string][]pair{}

	tok := Tokenizer{}
	for id, p := range db {
		for _, fnote := range p.Footnotes {
			root := tok.Parse(strings.NewReader(fnote))

			footnotes := Selector(root, func(n *html.Node) bool {
				return n.Data == "opr-ref" || n.Data == "opr-ref-link"
			})
			for _, fnote := range footnotes {
				args := toArgv(fnote)
				// recall: args[0] is name of function
				parts := strings.Split(args[1], "-")
				if parts[0] != "b" {
					continue
				}
				oprbirth := pmap[parts[2]]
				pmap[parts[2]] = append(oprbirth, pair{args[1], id, args[2]})
			}
		}
	}

	// sort parish numbers
	plist := []string{}
	for k, _ := range pmap {
		plist = append(plist, k)
	}
	sort.Strings(plist)

	out := strings.Builder{}
	out.WriteString("---\n")
	out.WriteString("title: OPR Baptism Index\n")
	out.WriteString("---\n")

	out.WriteString("<table class=base>\n")
	for _, pnum := range plist {
		out.WriteString("<tr><th colspan=4>")
		out.WriteString(fmt.Sprintf("Parish %s &mdash; %s", pnum, ParishName(pnum, "")))
		out.WriteString("</th></tr>\n")

		oprbirth := pmap[pnum]
		sort.Slice(oprbirth, func(i, j int) bool {
			return oprbirth[i].oprid < oprbirth[j].oprid
		})
		for i, item := range oprbirth {
			person := WriteTitle(db[item.pid])

			plink := fmt.Sprintf("<a href=/galbraith/people/%s/>%s</a>", item.pid, person)
			parts := strings.Split(item.oprid, "-")
			out.WriteString("<tr>")
			out.WriteString(fmt.Sprintf("<td>%d</td><td class=nowrap>%s</td><td class=nowrap>%s</td><td>%s</td>",
				i+1, oprref(parts), item.sv, plink))
			out.WriteString("</tr>\n")
		}
	}
	out.WriteString("</table>\n")
	return out.String()
}
func spindex(db Root, rtype string) string {
	word := ""
	switch rtype {
	case "b":
		word = "Birth"
	case "d":
		word = "Death"
	default:
		log.Fatalf("unknown type %q", rtype)
	}

	type pair struct {
		oprid string // b-1807-507-00-1234
		pid   string // galbreath-foo-year-spouse
		sv    string // name of person on record
	}
	pmap := map[string][]pair{}

	tok := Tokenizer{}
	for id, p := range db {
		for _, fnote := range p.Footnotes {
			root := tok.Parse(strings.NewReader(fnote))

			footnotes := Selector(root, func(n *html.Node) bool {
				return n.Data == "sp-ref" || n.Data == "sp-ref-link"
			})
			for _, fnote := range footnotes {
				args := toArgv(fnote)
				// recall: args[0] is name of function
				parts := strings.Split(args[1], "-")
				if parts[0] != rtype {
					continue
				}
				oprbirth := pmap[parts[2]]

				sv := args[2]
				if fnote.Data == "sp-ref-link" {
					sv = args[3]
				}
				pmap[parts[2]] = append(oprbirth, pair{args[1], id, sv})
			}
		}
	}

	// sort parish numbers
	plist := []string{}
	for k, _ := range pmap {
		plist = append(plist, k)
	}
	sort.Strings(plist)

	out := strings.Builder{}
	out.WriteString("---\n")
	out.WriteString(fmt.Sprintf("title: Statutory %s Index\n", word))
	out.WriteString("---\n")

	out.WriteString("<table class=base>\n")
	for _, pnum := range plist {
		out.WriteString("<tr><th colspan=4>")
		out.WriteString(fmt.Sprintf("Parish %s &mdash; %s\n", pnum, ParishName(pnum, "")))
		out.WriteString("</th></tr>\n")

		oprbirth := pmap[pnum]
		sort.Slice(oprbirth, func(i, j int) bool {
			return oprbirth[i].oprid < oprbirth[j].oprid
		})
		for i, item := range oprbirth {
			person := WriteTitle(db[item.pid])

			plink := fmt.Sprintf("<a href=/galbraith/people/%s/>%s</a>", item.pid, person)
			parts := strings.Split(item.oprid, "-")
			out.WriteString("<tr>")
			out.WriteString(fmt.Sprintf("<td>%d</td><td class=nowrap>%s</td class=nowrap><td>%s</td><td>%s</td>",
				i+1, statref(parts), item.sv, plink))
			out.WriteString("</tr>\n")
		}
	}
	out.WriteString("</table>\n")
	return out.String()
}
