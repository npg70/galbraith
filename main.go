package main

import (
	"flag"
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	// flag stuff
}

func serve(outdir string) {
	fmt.Println("Starting")
	http.Handle("/galbraith/", http.StripPrefix("/galbraith/", http.FileServer(http.Dir(outdir))))
	http.ListenAndServe("localhost:1313", nil)
}

func main() {
	outdir := ""
	rootsOnly := false
	server := false
	flag.StringVar(&outdir, "out", "", "out directory")
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

	/* set up base template */
	baset, err := os.ReadFile("baseof.html")
	if err != nil {
		log.Fatalf("can't find template: %s", err)
	}
	base, err := CreatePageTemplate(string(baset), RenderFunc(renderFuncs()))
	if err != nil {
		log.Fatalf("Unable to create template: %s", err)
	}

	// for each file in sources
	// read it in,
	suffix := ".sh"
	sources, err := filepath.Glob("sources/*" + suffix)
	if err != nil {
		log.Fatalf("Glob failed")
	}
	for _, s := range sources {
		raw, err := os.ReadFile(s)
		if err != nil {
			log.Fatalf("Unable to read %s", s)
		}
		outfile := strings.TrimSuffix(filepath.Base(s), suffix)
		writePage(string(raw), nil, base, outdir, filepath.Join("sources", outfile))
	}

	writePage(indexSources(), paragrapher, base, outdir, "sources")
	writePage(indexIndex(), nil, base, outdir, "indexes")

	writePage(oprindex(db, "b"), nil, base, outdir, "indexes/opr-birth-index")

	writePage(oprindex(db, "d"), nil, base, outdir, "indexes/opr-death-index")

	writePage(oprindex(db, "m"), nil, base, outdir, "indexes/opr-marriage-index")

	writePage(spindex(db, "b"), nil, base, outdir, "indexes/statutory-birth-index")

	writePage(spindex(db, "d"), nil, base, outdir, "indexes/statutory-death-index")

	writePage(spindex(db, "m"), nil, base, outdir, "indexes/statutory-marriage-index")

	// TODO: can we simplify and reuse something already?
	tmap := tagmap(db)
	writePage(tagIndex(tmap), nil, base, outdir, "tags")

	pages := tagStart(db)
	for _, tp := range pages {
		tagpath := makeTagFile(tp.path)
		writePage(indexRoots2(db, tp), nil, base, outdir, filepath.Join("tags", tagpath))
	}

	// write out each person
	for _, rootid := range roots {
		kidsq := []string{rootid}
		db.LoadAll(rootid)
		// one person, one page
		for len(kidsq) > 0 {
			uid := kidsq[0]
			kidsq = kidsq[1:]
			doc, nextg := db.generateOne(uid)
			kidsq = append(kidsq, nextg...)
			writePage(doc, peopleTree,
				base, outdir, filepath.Join("people", uid))
		}
	}

	if server {
		mime.AddExtensionType(".woff2", "font/woff2")
		serve(outdir)
	}
}
