package main

import (
	"sort"
	"strconv"

	"golang.org/x/net/html"

	tf "github.com/client9/tagfunctions"
)

// collects all references and changes them to ordinals
// collects all footnotes and sorts them based on ordinals
func footnoter(n *html.Node) error {

	persons := tf.Selector(n, func(n *html.Node) bool {
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

	refs := tf.Selector(n, func(n *html.Node) bool {
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
		id := tf.GetArg(ref, 0)
		// one source can be used multiple times
		// so only add if it's new.
		if _, ok := ordmap[id]; !ok {
			ordmap[id] = i
			i++
		}
	}

	// get all footnotes --they can be defined anywhere in the document
	footnotes := tf.Selector(n, func(n *html.Node) bool {
		return n.Data == "footnote"
	})
	// remove them from doc.. they'll be added back later
	for _, f := range footnotes {
		f.Parent.RemoveChild(f)
	}

	// sort them according to our map above
	sort.SliceStable(footnotes, func(i, j int) bool {
		a1 := ordmap[tf.GetArg(footnotes[i], 0)]
		a2 := ordmap[tf.GetArg(footnotes[j], 0)]
		return a1 < a2
	})

	// get the container for footnotes.  This is where footnotes will be
	// placed in document
	fnContainerList := tf.Selector(n, func(n *html.Node) bool {
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
	for fnContainer.FirstChild != nil {
		fnContainer.RemoveChild(fnContainer.FirstChild)
	}

	// change the refs to ordinals 1,2,3,4,5...
	for _, ref := range refs {
		if id := ordmap[tf.GetArg(ref, 0)]; id != 0 {
			tf.SetArg(ref, 0, strconv.Itoa(id))
		}
	}
	// change the ID to the ordinals
	for _, fn := range footnotes {
		if id := ordmap[tf.GetArg(fn, 0)]; id != 0 {
			tf.SetArg(fn, 0, strconv.Itoa(id))
		}
	}

	// then addback the footnotes in sorted order
	for _, fn := range footnotes {
		fnContainer.AppendChild(fn)
	}

	return nil
}
