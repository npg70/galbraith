package main

import (
	"flag"
	"log"
	"path/filepath"
	"time"

	tf "github.com/client9/tagfunctions"
	"github.com/npg70/ssg"
)

func init() {
	// flag stuff
}

type siteConfig struct {
	outputDir string
}

func (sconfig siteConfig) OutputDir() string {
	return sconfig.outputDir
}

func main() {
	sc := siteConfig{}
	rootsOnly := false
	server := false
	flag.StringVar(&sc.outputDir, "out", "", "out directory")
	flag.BoolVar(&rootsOnly, "root", false, "run and display roots")
	flag.BoolVar(&server, "serve", false, "run webserver")
	flag.Parse()

	db := make(Root)
	roots := computeRoots(db)
	if rootsOnly {
		return
	}
	// add meta tag.. if todos exist, add todo tag
	todotag(db)

	pages := []ssg.ContentSource{}
	pages = append(pages, oprindex(db, "b", "indexes/opr-birth-index/index.html"))
	pages = append(pages, oprindex(db, "d", "indexes/opr-death-index/index.html"))
	pages = append(pages, oprindex(db, "m", "indexes/opr-marriage-index/index.html"))
	pages = append(pages, spindex(db, "b", "indexes/statutory-birth-index/index.html"))
	pages = append(pages, spindex(db, "d", "indexes/statutory-death-index/index.html"))
	pages = append(pages, spindex(db, "m", "indexes/statutory-marriage-index/index.html"))

	// TAGS ------------------
	tmap := tagmap(db)
	pages = append(pages, tagIndex(tmap, "tags/index.html"))
	tpages := tagStart(db)
	for _, tp := range tpages {
		tagpath := makeTagFile(tp.path)
		pages = append(pages, indexRoots2(db, tp, filepath.Join("tags", tagpath, "index.html")))
	}

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

	if err := Main2(sc, &pages); err != nil {
		log.Fatalf("Main2 failed: %v", err)
	}

	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("%d pages in %s", len(pages), elapsed)
	if server {
		tf.Serve(sc.OutputDir(), "galbraith")
	}
}
