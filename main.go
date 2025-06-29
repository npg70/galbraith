package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/client9/ssg"
)

func init() {
	// flag stuff
}

func getLines(db Root, roots []string) [][]*Person {
	paths := [][]*Person{}
	for _, rootid := range roots {
		person := db[rootid]
		if person.Ydna != "" {
			paths = append(paths, []*Person{person})
		}
	}
	output := getLine(paths)
	return output
}

func getLine(paths [][]*Person) [][]*Person {
	output := [][]*Person{}
	for len(paths) > 0 {
		next := [][]*Person{}
		for _, p := range paths {
			// get last person in path
			person := p[len(p)-1]
			endofpath := true
			for _, kid := range person.AllChildren() {
				if kid.Ydna != "" {
					// copy existing path
					// add this kid
					// and to next queue
					// mark that this path keeps going
					line := make([]*Person, len(p))
					copy(line, p)
					line = append(p, kid)
					next = append(next, line)
					endofpath = false
				}
			}
			if endofpath {
				output = append(output, p)
			}
		}
		paths = next
	}
	return output
}

func main() {
	sc := ssg.SiteConfig{}
	outputDir := ""
	rootsOnly := false
	flag.StringVar(&outputDir, "out", "public", "out directory")
	flag.BoolVar(&rootsOnly, "root", false, "run and display roots")
	flag.Parse()

	db := make(Root)
	roots := computeRoots(db)
	if rootsOnly {
		return
	}

	paths := [][]*Person{}
	for _, rootid := range roots {
		person := db[rootid]
		if person.Ydna != "" {
			paths = append(paths, []*Person{person})
			log.Printf("YNDA ROOT: %s", person.FullName())
		}
	}
	paths = getLines(db, roots)

	for i, path := range paths {
		for j, p := range path {
			fmt.Printf("LINE %d - %d: %s\n", i+1, j+1, WriteBio(p, 1))
		}
		fmt.Printf("\n")
	}

	// add meta tag.. if todos exist, add todo tag
	todotag(db)

	tagAddFamilyName(db)

	// make full text index, do after all tags are created.
	fulltext(db)

	pages := []ssg.ContentSourceConfig{}
	pages = append(pages, oprPageIndex(db, "indexes/opr-page-index/index.html"))

	pages = append(pages, oprindex(db, "b", "indexes/opr-birth-index/index.html"))
	pages = append(pages, oprindex(db, "d", "indexes/opr-death-index/index.html"))
	pages = append(pages, oprindex(db, "m", "indexes/opr-marriage-index/index.html"))
	pages = append(pages, spindex(db, "b", "indexes/statutory-birth-index/index.html"))
	pages = append(pages, spindex(db, "d", "indexes/statutory-death-index/index.html"))
	pages = append(pages, spindex(db, "m", "indexes/statutory-marriage-index/index.html"))

	// TAGS ------------------
	tmap := tagmap(db)
	pages = append(pages, tagIndex(tmap, "tags/index.html"))
	/*  replaced with fulltext search
	tpages := tagStart(db)
	for _, tp := range tpages {
		tagpath := makeTagFile(tp.path)
		pages = append(pages, indexRoots2(db, tp, filepath.Join("tags", tagpath, "index.html")))
	}
	*/

	// PEOPLE
	count := 0
	for _, rootid := range roots {
		kidsq := []string{rootid}
		db.LoadAll(rootid)
		// one person, one page
		for len(kidsq) > 0 {
			uid := kidsq[0]
			kidsq = kidsq[1:]
			page, nextg := db.generateOne(uid, filepath.Join("people", uid, "index.html"))
			pages = append(pages, page)
			kidsq = append(kidsq, nextg...)
			count++
		}
	}

	start := time.Now()

	if err := Main2(sc, outputDir, &pages); err != nil {
		log.Fatalf("Main2 failed: %v", err)
	}

	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("%d pages in %s", len(pages), elapsed)
}
