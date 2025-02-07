package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tf "github.com/client9/tagfunctions"
	"golang.org/x/net/html"
)

func renderFuncs() map[string]tf.NodeFunc {
	m := map[string]tf.NodeFunc{
		"ent": tf.Entity,
		"banner": func(n *html.Node) error {
			tf.TransformElement(n, "div", "class", "mb-4")
			return nil
		},
		"ppre":   tf.MakeTagClass("p", "white-space-pre-line"),
		"strike": tf.MakeTag("s"),
		"root": func(n *html.Node) error {
			return nil
		},
		"section": tf.MakeTagClass("section", "mb-4"),

		"front": func(n *html.Node) error {
			// TODO: interesting case
			// do we delete it?
			tf.RemoveChildren(n)
			// turn into something benign
			tf.TransformElement(n, "div")
			return nil
		},
		"elink": func(n *html.Node) error {
			href := tf.GetArg(n, 0)
			tf.TransformElement(n, "a", "rel", "noreferrer", "target", "_blank", "href", href)
			return nil
		},
		"journal-link": func(n *html.Node) error {
			tf.TransformElement(n, "a", "href", "/galbraith/journal/"+tf.GetArg(n, 0))
			return nil
		},
		"tag-link": func(n *html.Node) error {
			path := strings.Split(tf.GetArg(n, 0), "/")
			tf.Replace(n, makeTagButton(path, n.FirstChild.Data))
			return nil
		},
		"title":  tf.MakeTag("h1"),
		"intro":  tf.MakeTagClass("p", ""),
		"nowrap": tf.MakeTagClass("span", "text-nowrap"),
		"csvtable": tf.NewCsvTableHTML(func(tag string, row int, col int) string {
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
		"primary-number": func(n *html.Node) error {
			// assuming FirstChild exists and is a text node
			// add a dot to end of text
			n.FirstChild.Data += "."
			tf.TransformElement(n, "span", "class", "fw-bold pe-3")
			return nil
		},
		"source-link": func(n *html.Node) error {
			tf.TransformElement(n, "a", "href", "/galbraith/sources/"+tf.GetArg(n, 0))
			return nil
		},
		"child-link": func(n *html.Node) error {
			tf.TransformElement(n, "a", "class", "text-smallcaps", "href", "/galbraith/people/"+tf.GetArg(n, 0))
			return nil
		},
		"child-link-plain": func(n *html.Node) error {
			path := tf.GetArg(n, 0)
			tf.TransformElement(n, "a", "class", "text-body", "href", "/galbraith/people/"+path)
			return nil
		},
		"lineage": func(n *html.Node) error {
			// <tr><th>Lineage</th><td>inner</td></tr>
			newNode :=
				tf.Append(
					tf.NewElement("tr"),
					tf.Append(tf.NewElement("th"), tf.NewText("Lineage")),
					tf.Reparent(tf.NewElement("td"), n))

			tf.Replace(n, newNode)

			return nil
		},
		"ancestor": func(n *html.Node) error {
			//genNumber := getKey(args, "generation")
			//counter := tf.GetKeyValue(args, "counter")
			mother := tf.GetAttr(n, "mother")
			year := tf.GetAttr(n, "year")
			if year != "" {
				year = " b. " + year + " "
			}
			tf.TransformElement(n, "div")
			n.AppendChild(tf.NewText(fmt.Sprintf("%sm. %s", year, mother)))
			return nil
		},
		"person": func(n *html.Node) error {
			pid := tf.GetAttr(n, "id")
			tf.TransformElement(n, "div", "class", "mb-5")
			n.InsertBefore(tf.NewElement("a", "name", pid), n.FirstChild)
			return nil
		},
		"person-body":      tf.MakeTag("div"),
		"person-main":      tf.MakeTagClass("div", "print-hack"),
		"person-secondary": tf.MakeTagClass("table", "small"),
		"pagemeta": func(n *html.Node) error {
			td := tf.Reparent(tf.NewElement("td"), n)
			tf.Append(tf.TransformElement(n, "tr", "class", "pb-3"),
				tf.Append(tf.NewElement("th", "class", "pe-5"), tf.NewText("External")),
				td)
			return nil
		},
		"person-bio": tf.MakeTag("div"),
		"headline":   tf.MakeTag("h1"),
		"externals": func(n *html.Node) error {
			body := tf.TextContent(n)
			if strings.TrimSpace(body) == "" {
				// TODO
				//n.Parent.RemoveChild(n)
				return nil
			}
			tf.Append(tf.TransformElement(n, "tr"),
				tf.Append(tf.NewElement("th", "class", "pe-5"), tf.NewText("External")),
				tf.Reparent(tf.NewElement("td"), n))
			return nil
		},
		"external": func(n *html.Node) error {
			arg := tf.GetArg(n, 0)
			tf.TransformElement(n, "a", "class", "me-3", "rel", "noreferrer", "target", "_blank", "href", arg)
			return nil
		},
		"tags": func(n *html.Node) error {
			args := tf.ToArgs(n)

			// no tags, just do nothing.
			if len(args) == 0 {
				// TODO
				//n.Parent.RemoveChild(n)
				return nil
			}

			// make a TD element containing the buttons
			td := tf.NewElement("td")
			for _, tag := range args {
				td.AppendChild(makeTagButton(nil, tag))
			}

			// make a row and replace
			tf.Replace(n, tf.Append(tf.NewElement("tr"),
				tf.Append(tf.NewElement("th", "class", "pe-5"), tf.NewText("Tags")),
				td))

			return nil
		},
		"child": func(n *html.Node) error {
			lineageCounter := tf.GetAttr(n, "counter")
			r, _ := strconv.Atoi(tf.GetAttr(n, "birth-order"))

			td := tf.NewElement("td", "class", "width-100r")
			if lineageCounter != "" && lineageCounter != "0" {
				tf.Append(td, tf.NewText(lineageCounter))
			}

			tf.Append(
				tf.TransformElement(n, "tr"),
				td,
				tf.Append(tf.NewElement("td", "class", "text-end width-200r"), tf.NewText(toRoman(r))),
				tf.Reparent(tf.NewElement("td", "class", "ps-3"), n),
			)
			return nil
		},
		"todos": tf.MakeTagClass("ul", "todos"),
		"todo":  tf.MakeTag("li"),
		"notes": tf.MakeTagClass("ul", "notes"),
		"note":  tf.MakeTag("li"),
		"footnotes": func(n *html.Node) error {
			body := tf.TextContent(n)
			if len(body) == 0 {
				// TODO
				//n.Parent.RemoveChild(n)
				return nil
			}
			n.Parent.InsertBefore(tf.NewElement("hr"), n)
			tf.Append(tf.TransformElement(n, "div"),
				tf.Reparent(tf.NewElement("table", "class", "table-p0 small"), n))
			return nil
		},
		"footnote": func(n *html.Node) error {
			ref := tf.GetArg(n, 0)
			td := tf.Reparent(tf.NewElement("td"), n)
			tf.Append(tf.TransformElement(n, "tr"),
				tf.Append(tf.NewElement("td", "class", "width-100r"), tf.Append(tf.NewElement("span"), (tf.NewText(ref)))),
				td)
			return nil
			/*
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
			*/
		},
		"ref": func(n *html.Node) error {
			ref := tf.GetArg(n, 0)
			if ref == "" {
				log.Printf("got empty ref")
				return nil
			}
			tf.Append(tf.TransformElement(n, "sup", "class", "footnote-ref"), tf.NewText("["+ref+"]"))
			return nil
		},
		"sp-ref": func(n *html.Node) error {
			// sp-ref: arg0 = ID, arg 1 = name, arg2 is optional name
			args := tf.ToArgs(n)
			if len(args) == 0 || len(args[0]) == 0 {
				return fmt.Errorf("%s: expected at least 2 args, got %v", n.Data, args)
			}
			person2 := ""
			refid := args[0]
			switch refid[0] {
			case 'b':
				if len(args) != 2 {
					return fmt.Errorf("%s: birth record expects 2 args, got %v", n.Data, args)
				}
			case 'd':
				if len(args) != 2 {
					return fmt.Errorf("%s: death record expects 2 args, got %v", n.Data, args)
				}
			case 'm':
				if len(args) != 3 {
					return fmt.Errorf("%s: marriage record expects 3 args, got %v", n.Data, args)
				}
				person2 = args[2]
			default:
				return fmt.Errorf("%s: unknown ID reference: %v", args)
			}

			person1 := args[1]
			tf.TransformElement(n, "div")
			return SPText(n, refid, "", person1, person2)
		},
		"opr-page": func(n *html.Node) error {

			args := tf.ToArgs(n)
			if len(args) != 1 || len(args[0]) == 0 {
				return fmt.Errorf("%s: expected at least 1 args, got %v", n.Data, args)
			}
			refid := args[0]

			tf.TransformElement(n, "div")
			return OPRPage(n, refid)
		},
		"sp-ref-link": func(n *html.Node) error {
			// sp-ref: arg0 = ID, arg1 = imgid, arg 2 = name, arg3 is optional name
			args := tf.ToArgs(n)
			if len(args) == 0 || len(args[0]) == 0 {
				return fmt.Errorf("%s: expected at least 3 args, got %v", n.Data, args)
			}
			person2 := ""
			refid := args[0]
			switch refid[0] {
			case 'b':
				if len(args) != 3 {
					return fmt.Errorf("%s: birth record expects 3 args, got %v", n.Data, args)
				}
			case 'd':
				if len(args) != 3 {
					return fmt.Errorf("%s: death record expects 3 args, got %v", n.Data, args)
				}
			case 'm':
				if len(args) != 4 {
					return fmt.Errorf("%s: marriage record expects 4 args, got %v", n.Data, args)
				}
				person2 = args[3]
			default:
				return fmt.Errorf("%s: unknown ID reference: %v", args)
			}
			imgid := args[1]
			person1 := args[2]
			tf.TransformElement(n, "div")
			return SPText(n, refid, imgid, person1, person2)
		},
		"opr-ref": func(n *html.Node) error {
			// sp-ref: arg0 = ID, arg1 = imgid, arg 2 = name, arg3 is optional name
			args := tf.ToArgs(n)
			if len(args) == 0 || len(args[0]) == 0 {
				return fmt.Errorf("%s: expected at least 3 args, got %v", n.Data, args)
			}
			person2 := ""
			refid := args[0]
			switch refid[0] {
			case 'b':
				if len(args) != 2 {
					return fmt.Errorf("%s: birth record expects 2 args, got %v", n.Data, args)
				}
			case 'd':
				if len(args) != 2 {
					return fmt.Errorf("%s: death record expects 2 args, got %v", n.Data, args)
				}
			case 'm':
				if len(args) != 3 {
					return fmt.Errorf("%s: marriage record expects 3 args, got %v", n.Data, args)
				}
				person2 = args[2]
			}
			person1 := args[1]
			tf.TransformElement(n, "div")
			return OPRText(n, refid, person1, person2, false)
		},
		"opr-ref-link": func(n *html.Node) error {
			// sp-ref: arg0 = ID, arg1 = imgid, arg 2 = name, arg3 is optional name
			args := tf.ToArgs(n)
			if len(args) == 0 || len(args[0]) == 0 {
				return fmt.Errorf("%s: expected at least 3 args, got %v", n.Data, args)
			}
			person2 := ""
			refid := args[0]
			switch refid[0] {
			case 'b':
				if len(args) != 2 {
					return fmt.Errorf("%s: birth record expects 2 args, got %v", n.Data, args)
				}
			case 'd':
				if len(args) != 2 {
					return fmt.Errorf("%s: death record expects 2 args, got %v", n.Data, args)
				}
			case 'm':
				if len(args) != 3 {
					return fmt.Errorf("%s: marriage record expects 3 args, got %v", n.Data, args)
				}
				person2 = args[2]
			}
			person1 := args[1]
			tf.TransformElement(n, "div")
			// move all children into a blockquote
			bq := tf.Reparent(tf.NewElement("blockquote", "class", "fs-100p"), n)
			tf.Append(n, bq)

			return OPRText(n, refid, person1, person2, true)
		},
	}

	return m
}
