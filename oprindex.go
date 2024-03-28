package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

// if person has any todos, make a tag
func todotag(db Root) {
	for _, p := range db {
		if len(p.Todos) > 0 {
			p.Tags = append(p.Tags, "TODO")
		}
	}
}

func oprindex(db Root, rtype string) string {
	word := ""
	switch rtype {
	case "b":
		word = "Baptism"
	case "d":
		word = "Death"
	case "m":
		word = "Marriage"
	default:
		log.Fatalf("unknown type %q", rtype)
	}
	type pair struct {
		oprid  string // b-1807-607-000-1234-1245
		pid    string // galbreath-foo-year-spouse
		sv     string // name of person on record
		spouse string // name of spouse, for marriage records
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
				if parts[0] != rtype {
					continue
				}
				oprbirth := pmap[parts[2]]
				spouse := ""
				if rtype == "m" {
					spouse = args[len(args)-1]
				}
				pmap[parts[2]] = append(oprbirth, pair{args[1], id, args[2], spouse})
			}
		}
	}

	// sort parish numbers
	plist := []string{}
	for k := range pmap {
		plist = append(plist, k)
	}
	sort.Strings(plist)

	out := strings.Builder{}

	// TODO: wrap in meta tag?
	out.WriteString("$title{OPR " + word + " Index}\n")
	out.WriteString("$table{\n")
	for _, pnum := range plist {
		out.WriteString("$tr{\n")
		out.WriteString("  $th[colspan=4 class='pt-3 pb-2 fs-5']{")
		out.WriteString(fmt.Sprintf("Parish %s &mdash; %s", pnum, ParishName(pnum, "")))
		out.WriteString("}\n")
		out.WriteString("}\n")

		oprbirth := pmap[pnum]
		sort.Slice(oprbirth, func(i, j int) bool {
			return oprbirth[i].oprid < oprbirth[j].oprid
		})
		for i, item := range oprbirth {
			person := WriteTitle(db[item.pid])
			parts := strings.Split(item.oprid, "-")
			name := "$nowrap{" + item.sv + "}"
			if rtype == "m" {
				name += " / " + "$nowrap{" + item.spouse + "}"
			}
			out.WriteString("$tr{\n")
			out.WriteString("  $td{" + strconv.Itoa(i+1) + "}\n")
			out.WriteString("  $td{$nowrap{" + parts[1] + " " + oprref(parts) + "}}\n")
			out.WriteString("  $td{" + name + "}\n")
			out.WriteString("  $td{$nowrap{$child-link-plain[" + item.pid + "]{" + person + "}}}\n")
			out.WriteString("}\n")
		}
	}
	out.WriteString("}\n") /* end of table tag */
	return out.String()
}
func spindex(db Root, rtype string) string {
	word := ""
	switch rtype {
	case "b":
		word = "Birth"
	case "d":
		word = "Death"
	case "m":
		word = "Marriage"
	default:
		log.Fatalf("unknown type %q", rtype)
	}

	type pair struct {
		oprid  string // b-1807-507-00-1234
		pid    string // galbreath-foo-year-spouse
		sv     string // name of person on record
		spouse string
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
					if len(args) < 4 {
						log.Fatalf("OOPS: %v", args)
					}
					sv = args[3]
				}
				spouse := ""
				if parts[0] == "m" {
					// last arg
					spouse = args[len(args)-1]
				}
				pmap[parts[2]] = append(oprbirth, pair{args[1], id, sv, spouse})
			}
		}
	}

	// sort parish numbers
	plist := []string{}
	for k := range pmap {
		plist = append(plist, k)
	}
	sort.Strings(plist)

	out := strings.Builder{}
	out.WriteString("$title{Statutory " + word + " Index}\n")

	out.WriteString("$table{\n")
	for _, pnum := range plist {
		out.WriteString("$tr{\n")
		out.WriteString("  $th[colspan=4 class='pt-3 pb-2 fs-5']{")
		out.WriteString(fmt.Sprintf("Parish %s &mdash; %s", pnum, ParishName(pnum, "")))
		out.WriteString("}\n")
		out.WriteString("}\n")

		oprbirth := pmap[pnum]
		sort.Slice(oprbirth, func(i, j int) bool {
			return oprbirth[i].oprid < oprbirth[j].oprid
		})
		for i, item := range oprbirth {
			person := WriteTitle(db[item.pid])

			name := "$nowrap{" + item.sv + "}"
			if rtype == "m" {
				name += " / " + "$nowrap{" + item.spouse + "}"
			}
			parts := strings.Split(item.oprid, "-")
			out.WriteString("$tr{\n")
			out.WriteString("  $td{" + strconv.Itoa(i+1) + "}\n")
			out.WriteString("  $td{$nowrap{" + statref(parts) + "}}\n")
			out.WriteString("  $td{" + name + "}\n")
			out.WriteString("  $td{$nowrap{$child-link-plain[" + item.pid + "]{" + person + "}}}\n")
			out.WriteString("}\n")
		}
	}
	out.WriteString("}\n") // table
	return out.String()
}
