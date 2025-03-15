package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	tf "github.com/client9/tagfunctions"
	"github.com/npg70/ssg"
	"golang.org/x/net/html"
)

// if person has any todos, make a tag
func todotag(db Root) {
	for _, p := range db {
		if len(p.Todos) > 0 {
			p.Tags = append(p.Tags, "todo")
		}
		for _, td := range p.Todos {
			if td.Name != "" {
				p.Tags = append(p.Tags, "todo_"+td.Name)
			}
		}
	}
}

func oprPageIndex(db Root, outputFile string) ssg.ContentSourceConfig {
	tok := tf.Tokenizer{}
	pages := map[string]string{}
	for _, p := range db {
		for _, fnote := range p.Footnotes {

			// parse the markup
			root := tok.Parse(strings.NewReader(fnote))

			// get the "opr-ref-link"
			footnotes := tf.Selector(root, func(n *html.Node) bool {
				return n.Data == "opr-ref-link"
			})

			//
			for _, fnote := range footnotes {
				// format ID Name Spouse?
				// ID is form of TYPE-YEAR-000-000-0000-000
				args := tf.ToArgs(fnote)
				val := args[0]
				if len(val) != 24 {
					log.Fatalf("Wrong size: %q", val)
				}
				pages[val[7:]] = val
			}
		}
	}

	// get keys, sort, print
	keys := make([]string, 0, len(pages))
	for k := range pages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	out := strings.Builder{}
	out.WriteString("$ul{\n")
	for _, val := range keys {
		out.WriteString("$li{$opr-page[" + pages[val] + "]}\n")
	}
	out.WriteString("}\n")
	page := make(ssg.ContentSourceConfig)
	page["OutputFile"] = outputFile
	page["TemplateName"] = "baseof.html"
	page["title"] = "OPR Page Index"
	page["Content"] = out.String()
	return page
}

// input, output is content
func oprindex(db Root, rtype string, outputFile string) ssg.ContentSourceConfig {

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

	tok := tf.Tokenizer{}
	for id, p := range db {
		for _, fnote := range p.Footnotes {
			root := tok.Parse(strings.NewReader(fnote))

			footnotes := tf.Selector(root, func(n *html.Node) bool {
				return n.Data == "opr-ref" || n.Data == "opr-ref-link"
			})
			for _, fnote := range footnotes {
				// format ID Name Spouse?
				args := tf.ToArgs(fnote)
				parts := strings.Split(args[0], "-")
				if parts[0] != rtype {
					continue
				}
				oprbirth := pmap[parts[2]]
				spouse := ""
				if rtype == "m" {
					spouse = args[len(args)-1]
				}
				if len(args) < 2 {
					log.Printf("Wrong arg length: %v", args)
				}
				pp := pair{args[0], id, args[1], spouse}
				pmap[parts[2]] = append(oprbirth, pp)
			}
		}
	}

	// sort parish numbers
	plist := []string{}
	for k := range pmap {
		plist = append(plist, k)
	}
	sort.Strings(plist)

	// generate content
	out := strings.Builder{}
	out.WriteString("$table{\n")
	for _, pnum := range plist {
		out.WriteString("$tr{\n")
		out.WriteString("  $th[colspan=4 class='pt-3 pb-2 fs-5']{")
		out.WriteString(fmt.Sprintf("Parish %s $ent[mdash] %s", pnum, ParishName(pnum, "")))
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

	page := make(ssg.ContentSourceConfig)
	page["OutputFile"] = outputFile
	page["TemplateName"] = "baseof.html"
	page["title"] = "OPR " + word + " Index"
	page["Content"] = out.String()
	return page
}

func spindex(db Root, rtype string, outputFile string) ssg.ContentSourceConfig {
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

	tok := tf.Tokenizer{}
	for id, p := range db {
		for _, fnote := range p.Footnotes {
			root := tok.Parse(strings.NewReader(fnote))

			footnotes := tf.Selector(root, func(n *html.Node) bool {
				return n.Data == "sp-ref" || n.Data == "sp-ref-link"
			})
			for _, fnote := range footnotes {
				args := tf.ToArgs(fnote)
				parts := strings.Split(args[0], "-")
				if parts[0] != rtype {
					continue
				}
				oprbirth := pmap[parts[2]]

				sv := args[1]
				if fnote.Data == "sp-ref-link" {
					if len(args) < 3 {
						log.Fatalf("OOPS: %v", args)
					}
					sv = args[2]
				}
				spouse := ""
				if parts[0] == "m" {
					// last arg
					spouse = args[len(args)-1]
				}
				pmap[parts[2]] = append(oprbirth, pair{args[0], id, sv, spouse})
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
		out.WriteString(fmt.Sprintf("Parish %s $ent[mdash] %s", pnum, ParishName(pnum, "")))
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

	page := make(ssg.ContentSourceConfig)
	page["OutputFile"] = outputFile
	page["TemplateName"] = "baseof.html"
	page["title"] = "Statutory " + word + " Index"
	page["Content"] = out.String()
	return page
}
