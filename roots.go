package main

import (
	"bytes"
	"fmt"
	"log"
	"maps"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/bep/gitmap"
	"github.com/client9/ssg"
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
		if v {
			continue
		}
		person, _ := db.Load(k)

		// right now, almost no one is a root and illegitimate.
		// this may change but right now it's just noise
		if person.HasTag("illegitimate") {
			continue
		}

		person.Tags = append([]string{"root"}, person.Tags...)
		out = append(out, k)
	}

	// make consistent
	sort.Strings(out)

	return out
}

// home page for tags
func tagIndex(tmap map[string][]string, outputFile string) ssg.ContentSourceConfig {
	out := bytes.Buffer{}
	out.WriteString(`
$table{
	$tr{
		$th[colspan=2]{Major Families}
	}
	$tr{
		$td{$tag-link{colville}
		$td{...}
	}
	$tr{
		$td{$tag-link{fleeming}
		$td{...}
	}
	$tr{
		$td{$tag-link{galbraith}
		$td{...}
	}
	$tr{
		$td{$tag-link{greenlees}
		$td{...}
	}
	$tr{
		$td{$tag-link{harvey}
		$td{...}
	}
	$tr{
		$td{$tag-link{langwill}
		$td{...}
	}
	$tr{
		$td{$tag-link{mcnair}
		$td{...}
	}
	$tr{
		$td{$tag-link{mitchell}
		$td{...}
	}
	$tr{
		$td{$tag-link{ralston}
		$td{...}
	}
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
		$td{$tag-link{nochildren}}
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
		$td{$tag-link{daughteredout}}
		$td{No direct male decendants (open to better word or terminology here).}
	}
	$tr{
		$td{$tag-link{endofline}}
		$td{No children, or all children had no children}
	}
	$tr{
		$td{$tag-link{deadend}}
		$td{Families with no records after a time. Either immigranted or died-out without being recorded.}
	}
	$tr{
		$th[colspan=2]{All Tags}
	}
	$tr{
		$td{ }
	`)

	// dump everything else
	tags := []string{}
	special := ""
	for tname := range tmap {
		tname = strings.ToLower(tname)
		if strings.Contains(special, tname) {
			continue
		}
		tags = append(tags, tname)
	}
	sort.Strings(tags)
	out.WriteString("$td{")
	for _, tag := range tags {
		out.WriteString("$tag-link{" + tag + "} ")
	}
	out.WriteString("}\n") // td
	out.WriteString("}\n") // row
	out.WriteString("}\n") // table

	page := make(ssg.ContentSourceConfig)
	page["OutputFile"] = outputFile
	page["TemplateName"] = "tags/baseof.html"
	page["title"] = "People Tags"
	page["Content"] = out.Bytes()
	return page
}

func indexRoots2(db Root, tpage tagpage, outputFile string) ssg.ContentSourceConfig {
	fulltag := strings.Join(tpage.path, " / ")

	out := bytes.Buffer{}
	out.WriteString("$h1{" + fulltag + "}")
	for _, kid := range tpage.tags {
		out.WriteString(fmt.Sprintf("$tag-link[%s]{%s}", strings.Join(tpage.path, "/"), kid))
	}
	out.WriteString("$hr{}\n")
	for _, r := range tpage.uids {
		p := db[r]
		out.WriteString(fmt.Sprintf("$h5{$a[href=/galbraith/people/%s]{%s}}\n",
			r, WriteTitle(p)))
		if len(p.Tags) > 0 {
			out.WriteString("$div[class='ms-3 mb-3']{")
			for _, tag := range p.Tags {
				out.WriteString(fmt.Sprintf("$tag-link[%s]{%s}", strings.Join(tpage.path, "/"), tag))
			}
			out.WriteString("}\n")
		}
		if len(p.Intro) > 0 {
			out.WriteString("$intro{" + p.Intro + "}\n")
		}
	}
	page := make(ssg.ContentSourceConfig)
	page["OutputFile"] = outputFile
	page["TemplateName"] = "tags/baseof.html"
	page["title"] = "People Tag " + fulltag
	page["Content"] = out.Bytes()
	return page
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
	children := slices.Sorted(maps.Keys(tmap))
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
