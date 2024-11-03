package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/exp/maps"

	"github.com/bep/gitmap"
)

// computes roots of the tree and makes indexes for it.

func computeRoots(db Root) []string {
	opt := gitmap.Options{}
	repo, err := gitmap.Map(opt)
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
$ul{
$li{$a[href=/galbraith/indexes/opr-birth-index/ ]{OPR Birth Index}}
$li{$a[href=/galbraith/indexes/opr-death-index/ ]{OPR Death Index}}
$li{$a[href=/galbraith/indexes/opr-marriage-index/ ]{OPR Marriage Index}}
}
$ul{
$li{$a[href=/galbraith/indexes/statutory-birth-index/ ]{Statutory Birth Index}}
$li{$a[href=/galbraith/indexes/statutory-death-index/ ]{Statutory Death Index}}
$li{$a[href=/galbraith/indexes/statutory-marriage-index/ ]{Statutory Marriage Index}}
}
	`)
	return out.String()
}

func indexSources() string {
	b, err := os.ReadFile("sources/_index.sh")
	if err != nil {
		log.Fatalf("Can't read roots/index.sh: %v", err)
	}
	return string(b)
}

// home page for tags
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

func indexRoots2(db Root, page tagpage) string {
	out := strings.Builder{}

	out.WriteString("$h1{")
	out.WriteString(strings.Join(page.path, " / "))
	out.WriteString("}")
	for _, kid := range page.tags {
		out.WriteString(fmt.Sprintf("$tag-link[%s]{%s}", strings.Join(page.path, "/"), kid))
	}
	out.WriteString("$hr{}\n")
	for _, r := range page.uids {
		p := db[r]
		out.WriteString(fmt.Sprintf("$h5{$a[href=/galbraith/people/%s]{%s}}\n",
			r, WriteTitle(p)))
		if len(p.Tags) > 0 {
			out.WriteString("$div[class='ms-3 mb-3']{")
			for _, tag := range p.Tags {
				out.WriteString(fmt.Sprintf("$tag-link[%s]{%s}", strings.Join(page.path, "/"), tag))
			}
			out.WriteString("}\n")
		}
		if len(p.Intro) > 0 {
			out.WriteString("$intro{" + p.Intro + "}\n")
		}
	}
	return out.String()
}

type tagpage struct {
	path []string
	uids []string
	tags []string
}

// Current Path
// All matching childen
// Tags for next path

func tagmap2(db Root, page tagpage, depth int) []tagpage {
	if depth == 4 {
		return nil
	}
	// this is a map of TAG --> USERID
	tmap := make(map[string][]string)
	for _, uid := range page.uids {
		// for each tag the user has
		for _, tag := range db[uid].Tags {

			// split into parts
			tag = strings.ToLower(tag)
			tp := strings.Split(tag, ":")

			// foo, foo:bar, foo:bar:ding
			for i := 1; i <= len(tp); i++ {
				tg := strings.Join(tp[0:i], ":")
				// log.Printf("trying: path=%q, tag=%q --> %v", page.path, tg, inLabelPath(page.path, tg))
				if !inLabelPath(page.path, tg) {
					tmap[tg] = append(tmap[tg], uid)
				}
			}
		}
	}
	children := maps.Keys(tmap)
	sort.Strings(children)
	page.tags = children
	result := []tagpage{page}

	// we now have a set of all tags used with given UIDs
	for tag, newuids := range tmap {
		next := tagpage{
			path: append(page.path, tag),
			uids: newuids,
		}
		result = append(result, tagmap2(db, next, depth+1)...)
	}
	return result
}

func tagStart(db Root) []tagpage {

	// all is all people
	all := []string{}
	for uid := range db {
		all = append(all, uid)
	}
	sort.Strings(all)

	next := tagpage{
		path: []string{"all"},
		uids: all,
	}
	return tagmap2(db, next, 0)
}

func tagmap(db Root) map[string][]string {

	// mapping of tag ---> list of UIDs
	idx := map[string][]string{}

	// fake tag "all" that has everyone
	all := []string{}

	for uid, p := range db {

		// all gets everything
		all = append(all, uid)

		// append the UID to each tag
		for _, tag := range p.Tags {
			tag = strings.ToLower(tag)
			idx[tag] = append(idx[tag], uid)
		}
	}
	idx["all"] = all

	// sort the list of UIDs so it's stable
	for tag, uids := range idx {
		sort.Strings(uids)
		idx[tag] = uids
	}

	return idx
}
