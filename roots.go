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
			out = append(out, k)
		}
	}

	// make consistent
	sort.Strings(out)

	return out
}

// makes a page showing top level roots
func indexRoots(db Root, roots []string) string {
	out := strings.Builder{}
	out.WriteString("---\n")
	out.WriteString("title: Galbraith Progenitors\n")
	out.WriteString("---\n")

	for _, r := range roots {
		p := db[r]
		out.WriteString(fmt.Sprintf("<h5><a href={{< relref %s >}}>%s</a></h5>\n",
			r, WriteTitle(p)))
		if len(p.Tags) > 0 {
			out.WriteString("<div class='ms-3 mb-3'>\n")
			for _, tag := range p.Tags {
				out.WriteString("<span class='badge bg-secondary'>" + tag + "</span>\n")
			}
			out.WriteString("</div>\n")
		}
	}
	return out.String()
}
