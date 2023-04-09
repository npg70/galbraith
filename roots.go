package main

import (
	"fmt"
	"log"
	"path/filepath"
	"sort"
	"strings"

	"github.com/bep/gitmap"
)

// computes roots of the tree and makes indexes for it.

func computeRoots(db Root) []string {
	repo, err := gitmap.Map(".", "")
	if err != nil {
		log.Fatal(err)
	}

	pmap := make(map[string]bool)
	files, err := filepath.Glob("./people/*.sh")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		uid := strings.TrimSuffix(filepath.Base(f), ".sh")
		nextg := db.loadOne(uid)
		if ginfo, ok := repo.Files[f]; ok {
			db[uid].lastUpdate = ginfo.CommitDate
		}

		// if no children this is a leaf node
		if len(nextg) == 0 {
			person, _ := db.Load(uid)
			person.Tags = append([]string{"leaf"}, person.Tags...)
		}
		for _, child := range nextg {
			pmap[child] = true
		}
		if _, ok := pmap[uid]; !ok {
			pmap[uid] = false
		}
	}

	out := []string{}
	for k, v := range pmap {
		if !v {
			person, _ := db.Load(k)
			person.Tags = append([]string{"root"}, person.Tags...)
			out = append(out, k)
		}
	}

	// make consistent
	sort.Strings(out)

	return out
}

// makes a page showing top level roots
func indexRoots(db Root, roots []string, title string) string {
	out := strings.Builder{}
	out.WriteString("---\n")
	out.WriteString("title: Galbraith tag " + title + "\n")
	out.WriteString("---\n")

	for _, r := range roots {
		p := db[r]
		out.WriteString(fmt.Sprintf("<h5><a href=/galbraith/people/%s>%s</a></h5>\n",
			r, WriteTitle(p)))
		if len(p.Tags) > 0 {
			out.WriteString("<div class='ms-3 mb-3'>\n")
			for _, tag := range p.Tags {
				taglink := "/galbraith/tags/" + strings.ReplaceAll(strings.ToLower(tag), " ", "-") + "/"
				if true {
					out.WriteString("<a class='btn btn-sm btn-secondary' href=" + taglink + ">" + TitleCompound(tag) + "</a> ")
				} else {
					out.WriteString("<a href=#><span class='badge bg-secondary'>" + tag + "</span></a> ")
				}
			}
			out.WriteString("</div>\n")
		}
	}
	return out.String()
}

func tagmap(db Root) map[string][]string {
	idx := map[string][]string{}

	// fake tag "all" that has everyone
	all := []string{}

	for uid, p := range db {
		all = append(all, uid)
		for _, tag := range p.Tags {
			tag = strings.ToLower(tag)
			idx[tag] = append(idx[tag], uid)
		}
	}
	idx["all"] = all
	for tag, uids := range idx {
		sort.Strings(uids)
		idx[tag] = uids
	}
	return idx
}
