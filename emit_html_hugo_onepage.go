package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// paragrapher -- creates paragraphs tags based on \n\n

func paragrapher(n *html.Node) error {

	blocks := Selector(n, func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == "blockquote"
	})
	for _, block := range blocks {
		p := &html.Node{
			Type:     html.ElementNode,
			DataAtom: atom.P,
			Data:     "p",
		}
		bq := &html.Node{
			Type:     html.ElementNode,
			DataAtom: atom.Blockquote,
			Data:     "blockquote",
		}
		current := block.FirstChild
		for current != nil {
			// TODO: if it's a block element, close of paragraph tag if any
			// and add it.
			if current.Type == html.ElementNode {
				next := current.NextSibling
				block.RemoveChild(current)
				bq.AppendChild(current)
				current = next
				continue
			}
			if current.Type == html.TextNode {
				ps := strings.Split(current.Data, "\n\n")
				p.AppendChild(&html.Node{
					Type: html.TextNode,
					Data: ps[0],
				})
				for _, txt := range ps[1:] {
					bq.AppendChild(p)

					p = &html.Node{
						Type:     html.ElementNode,
						DataAtom: atom.P,
						Data:     "p",
					}
					p.AppendChild(&html.Node{
						Type: html.TextNode,
						Data: txt,
					})
				}

				next := current.NextSibling
				block.RemoveChild(current)
				current = next
				continue
			}
		}
		if p.FirstChild != nil {
			bq.AppendChild(p)
		}
		// at this point we have a new blockquote element all
		// filled out
		parent := block.Parent
		parent.InsertBefore(bq, block)
		parent.RemoveChild(block)
	}

	return nil
}

// collects all references and changes them to ordinals
// collects all footnotes and sorts them based on ordinals
func footnoter(n *html.Node) error {

	persons := Selector(n, func(n *html.Node) bool {
		return n.Data == "person"
	})
	for _, p := range persons {
		if err := footnoter_person(p); err != nil {
			return err
		}
	}
	return nil
}

func footnoter_person(n *html.Node) error {
	ordmap := map[string]int{}

	refs := Selector(n, func(n *html.Node) bool {
		return n.Data == "ref"
	})

	// $ref[xxx]. --> .$ref[xxx]
	for _, ref := range refs {
		if ref.NextSibling != nil {
			next := ref.NextSibling
			if next.Type == html.TextNode {
				text := next.Data
				if len(text) > 0 {
					first := text[0]
					if first == '.' || first == ',' || first == ';' || first == ':' {
						next.Data = next.Data[1:]
						if ref.PrevSibling != nil && ref.PrevSibling.Type == html.TextNode {
							ref.PrevSibling.Data += string(first)
						} else {
							newNode := &html.Node{
								Type: html.TextNode,
								Data: string(first),
							}
							ref.Parent.InsertBefore(newNode, ref)
						}
					}
				}
			}
		}
	}
	i := 1
	for _, ref := range refs {
		id := getArg(ref, 1)
		// one source can be used multiple times
		// so only add if it's new.
		if _, ok := ordmap[id]; !ok {
			ordmap[id] = i
			i++
		}
	}

	// get all footnotes --they can be defined anywhere in the document
	footnotes := Selector(n, func(n *html.Node) bool {
		return n.Data == "footnote"
	})
	// remove them from doc.. they'll be added back later
	for _, f := range footnotes {
		f.Parent.RemoveChild(f)
	}

	// sort them according to our map above
	sort.SliceStable(footnotes, func(i, j int) bool {
		a1 := ordmap[getArg(footnotes[i], 1)]
		a2 := ordmap[getArg(footnotes[j], 1)]
		return a1 < a2
	})

	// get the container for footnotes.  This is where footnotes will be
	// placed in document
	fnContainerList := Selector(n, func(n *html.Node) bool {
		return n.Data == "footnotes"
	})

	if len(refs) == 0 && len(footnotes) == 0 {
		return nil
	}

	if len(fnContainerList) != 1 {
		return nil

		//panic(fmt.Sprintf("refs: %d  footnotes: %d - Expected to find one footnotes container",
		//		len(refs), len(footnotes)))
	}
	fnContainer := fnContainerList[0]

	// delete everything in fnContainer, as it might contain text nodes
	// or unsorted foot notes
	for c := fnContainer.FirstChild; c != nil; c = c.NextSibling {
		fnContainer.RemoveChild(c)
	}

	// change the refs to ordinals 1,2,3,4,5...
	for _, ref := range refs {
		if id := ordmap[getArg(ref, 1)]; id != 0 {
			setArg(ref, 1, strconv.Itoa(id))
		}
	}
	// change the ID to the ordinals
	for _, fn := range footnotes {
		if id := ordmap[getArg(fn, 1)]; id != 0 {
			setArg(fn, 1, strconv.Itoa(id))
		}
	}

	// then addback the footnotes in sorted order
	for _, fn := range footnotes {
		fnContainer.AppendChild(fn)
	}

	return nil
}

func emitHTMLHugo(n *html.Node, singlePage bool) string {
	front := ""
	if singlePage {
		front += "---\n"
		front += "title: test\n"
		front += "---\n"
	}

	fmap := map[string]TagFunc{
		"csvtable": CsvTableHTML,
		"root": func(args []string, body string) string {
			return body
		},
		"front": func(args []string, body string) string {
			out := ""
			if !singlePage {
				out += "---\n"
				out += body
				out += "---\n"
			}
			return out
		},
		"title": func(args []string, body string) string {
			return "title: " + body + "\n"
		},
		"date":       makeTagClass("span"),
		"child-list": makeTagClass("table"),
		"gen": func(args []string, body string) string {
			return "<sup>" + body + "</sup>"
		},
		"children":           makeTagClass("div"),
		"child-partner-name": makeTagClass("span"),
		"partner-name":       makeTagClass("span"),
		"child-name":         makeTagClass("span"),
		"primary-number": func(args []string, body string) string {
			return fmt.Sprintf("<span class=%s>%s. </span>", args[0], body)
		},
		"primary-name":   makeTagClass("span"),
		"lineage-name":   makeTagClass("span"),
		"children-intro": makeTagClass("p"),
		"child-link": func(args []string, body string) string {
			href := args[1] //  getKeyValue(args, "href")
			out := "<a class=child-link href='"
			if singlePage {
				// link within page
				out += "#" + href
			} else {
				// link to new web page
				out += "{{< relref " + href + ">}}"
			}
			out += "'>"
			out += body
			out += "</a>"
			return out
		},
		"lineage": func(args []string, body string) string {
			out := strings.Builder{}
			out.WriteString("<div>")
			out.WriteString("<h6>Lineage</h6>")
			out.WriteString("<ol class=lineage>")
			out.WriteString(body)
			out.WriteString("</ol>")
			out.WriteString("</div><!-- lineage -->\n")
			return out.String()
		},
		"ancestor": func(args []string, body string) string {
			//genNumber := getKey(args, "generation")
			counter := getKeyValue(args, "counter")
			mother := getKeyValue(args, "mother")
			return fmt.Sprintf("<li value=%s class=%s>%s<br>%s</li>\n",
				counter, args[0], body, mother)
		},
		"person": func(args []string, body string) string {
			pid := getKeyValue(args, "id")
			s := strings.Builder{}
			s.WriteString("<div class='container-fluid person mx-auto'>")
			s.WriteString(fmt.Sprintf("<a name=%q></a>\n", pid))
			s.WriteString(body)
			s.WriteString("</div><!-- person -->\n") // person
			return s.String()
		},
		"person-body": func(args []string, body string) string {
			return "<div class='row'>\n" + body + "</div><!-- row -->\n"
		},
		"person-main":      makeDivClass("col-12 col-md-9 col-order-first print-hack"),
		"person-secondary": makeDivClass("person-secondary col-12 col-md-3 order-last"),
		"banner": func(args []string, body string) string {
			return "<div class='row'><div class='col-12 banner'>" + body + "</div></div>\n"
		},
		"pagemeta": func(args []string, body string) string {
			return fmt.Sprintf("<div class=%s><h6>Meta</h6>%s</div><!-- %s -->\n",
				args[0], body, args[0])
		},
		"person-bio": makeTagClass("div"),

		"headline": makeTag("h1"),
		"externals": func(args []string, body string) string {
			if strings.TrimSpace(body) == "" {
				return ""
			}
			return fmt.Sprintf("<div class=%s><h6>External</h6><ul class=externals>%s</ul></div><!-- %s -->\n", args[0], body, args[0])
		},
		"external": func(args []string, body string) string {
			return fmt.Sprintf("<li><a href=%q>%s</a></li>\n", args[1], body)
		},
		"tags": func(args []string, body string) string {
			if len(args) == 0 {
				return ""
			}
			s := strings.Builder{}
			s.WriteString("<div class=''>\n")
			for _, tag := range args[1:] {
				taglink := "/galbraith/tags/" + strings.ReplaceAll(strings.ToLower(tag), " ", "-") + "/"
				s.WriteString("<a class='btn btn-sm btn-secondary text-decoration-none' href=" + taglink + ">" + TitleCompound(tag) + "</a> ")
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
			s.WriteString("<td class=roman>" + toRoman(r) + "." + "</td>")
			s.WriteString("<td class=child-entry>")
			s.WriteString(body)
			s.WriteString("</td>")
			s.WriteString("</tr>\n")
			return s.String()
		},
		"notes": makeTagClass("ul"),
		"note":  makeTag("li"),
		"footnotes": func(args []string, body string) string {
			if len(body) == 0 {
				return ""
			}
			s := strings.Builder{}
			s.WriteString("<div class='row'><div class='col-12 footnotes'>\n")
			s.WriteString("<hr>\n<table>\n")
			s.WriteString(body)
			s.WriteString("</table></div></div><!-- footnotes -->\n")
			return s.String()
		},
		"footnote": func(args []string, body string) string {
			s := strings.Builder{}
			s.WriteString("<tr>")
			s.WriteString("<td class=footnote-counter>")
			s.WriteString("<span><sup>" + args[1] + "</sup>&nbsp;</span>")
			s.WriteString("</td>")
			s.WriteString("<td class=footnote-body>")
			s.WriteString(strings.TrimSpace(body))
			s.WriteString("</td>")
			s.WriteString("</tr>\n")
			return s.String()
		},
		"ref": func(args []string, body string) string {
			return "<sup class=footnote-ref>[" + args[1] + "]</sup>"
		},
		"sp-ref": func(args []string, body string) string {
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
			person2 := ""
			if len(args) == 5 {
				person2 = args[4]
			}
			return SPText(args[1], args[2], args[3], person2) + "<blockquote>" + strings.TrimSpace(body) + "</blockquote>"
		},

		"opr-ref": func(args []string, body string) string {
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
			person2 := ""
			if len(args) == 4 {
				person2 = args[3]
			}

			return OPRText(args[1], args[2], person2, true) + "<blockquote>" + strings.TrimSpace(body) + "</blockquote>\n"
		},
	}

	out := Render(n, fmap)
	return front + out
}
