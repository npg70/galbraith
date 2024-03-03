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

// content for index page
func indexIndex() string {
	out := strings.Builder{}
	out.WriteString(`
<ul>
<li><a href="/galbraith/indexes/opr-birth-index/">OPR Birth Index</a></li>
<li><a href="/galbraith/indexes/opr-death-index/">OPR Death Index</a></li>
<li><a href="/galbraith/indexes/opr-marriage-index/">OPR Marriage Index</a></li>
</ul>
<ul>
<li><a href="/galbraith/indexes/statutory-birth-index/">Statutory Birth Index</a></li>
<li><a href="/galbraith/indexes/statutory-death-index/">Statutory Death Index</a></li>
<li><a href="/galbraith/indexes/statutory-marriage-index/">Statutory Marriage Index</a></li>
</ul>
	`)
	return out.String()
}

func tagIndex(tmap map[string][]string) string {
	out := strings.Builder{}
	out.WriteString(`
$table{
	$tr{
		$th[colspan=2]{Tree Position}
	}
	$tr{
		$td{$tag-link{all}}
		$td{All people with dediciated pages}
	}
	$tr{
		$td{$tag-link{root}}
		$td{Roots, or start of a line}
	}
	$tr{
		$td{$tag-link{leaf}}
		$td{Leaves, or most recent persons with any children followed.}
	}
	$tr{
		$th[colspan=2]{Life Events}
	}
	$tr{
		$td{$tag-link{immigrant}}
		$td{People who left Scotland or UK}
	}
	$tr{
		$td{$tag-link{veteran}}
		$td{People who served in the military}
	}
	$tr{
		$td{$tag-link{no children}}
		$td{Person or couple has no children}
	}
	$tr{
	    $td{$tag-link{illegitimate}}
		$td{Typically a NPE with unknown or non-Galbraith father}
	}
	$tr{	
		$th[colspan=2]{Research}
	}
	$tr{
		$td{$tag-link{todo}}
		$td{Page has research to do, or need typographical or copy edits.}
	}
	$tr{
		$td{$tag-link{daughtered out}}
		$td{No direct male decendants (open to better word or terminology here).}
	}
	$tr{
		$td{$tag-link{end of line}}
		$td{No children, or all children had no children}
	}
	$tr{
		$td{$tag-link{dead end}}
		$td{Families with no records after a time. Either immigranted or died-out without being recorded.}
	}
	$th[colspan=2]{Places}
	`)

	// now figure out place name tags
	tags := []string{}
	special := "|all|daughtered out|dead end|illegitimate|immigrant|leaf|no children|root|todo|veteran|end of line|"
	for tname := range tmap {
		tname = strings.ToLower(tname)
		if strings.Contains(special, tname) {
			continue
		}
		tags = append(tags, tname)
	}
	sort.Strings(tags)
	for _, tag := range tags {
		out.WriteString("$tr{$td{$tag-link{" + tag + "}}$td{}}\n")
	}
	out.WriteString("}\n") // table end
	return out.String()
}

// makes a page showing pages that have a tag
func indexRoots(db Root, roots []string, title string) string {
	out := strings.Builder{}
	for _, r := range roots {
		p := db[r]
		out.WriteString(fmt.Sprintf("<h5><a href=/galbraith/people/%s>%s</a></h5>\n",
			r, WriteTitle(p)))
		if len(p.Tags) > 0 {
			out.WriteString("<div class='ms-3 mb-3'>\n")
			for _, tag := range p.Tags {
				taglink := "/galbraith/tags/" + strings.ReplaceAll(strings.ToLower(tag), " ", "-") + "/"
				out.WriteString("<a class='btn btn-sm btn-secondary' href=" + taglink + ">" + TitleTag(tag) + "</a> ")
			}
			out.WriteString("</div>\n")
		}
		if len(p.Intro) > 0 {
			out.WriteString("$intro{" + p.Intro + "}\n")
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
