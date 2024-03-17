package main

import (
	"flag"
	"fmt"
	"log"
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
		writePage(outdir, filepath.Join("sources", outfile),
			base, string(raw))
	}

	writePage(outdir, "indexes", base, indexIndex())

	writePage(outdir, "indexes/opr-birth-index", base,
		oprindex(db, "b"))

	writePage(outdir, "indexes/opr-death-index", base,
		oprindex(db, "d"))

	writePage(outdir, "indexes/opr-marriage-index", base,
		oprindex(db, "m"))

	writePage(outdir, "indexes/statutory-birth-index", base,
		spindex(db, "b"))

	writePage(outdir, "indexes/statutory-death-index", base,
		spindex(db, "d"))

	writePage(outdir, "indexes/statutory-marriage-index", base,
		spindex(db, "m"))

	tmap := tagmap(db)

	writePage(outdir, "tags", base, tagIndex(tmap))

	pages := tagStart(db)
	for _, tp := range pages {
		tagpath := makeTagFile(tp.path)
		writePage(outdir, filepath.Join("tags", tagpath), base,
			indexRoots2(db, tp))
	}

	for _, rootid := range roots {
		kidsq := []string{rootid}
		db.LoadAll(rootid)
		// one person, one page
		for len(kidsq) > 0 {
			uid := kidsq[0]
			kidsq = kidsq[1:]
			doc, nextg := db.generateOne(uid)
			kidsq = append(kidsq, nextg...)
			log.Printf("Processing %s", uid)
			p := Tokenizer{}
			root := p.Parse(strings.NewReader(doc))
			if err := footnoter(root); err != nil {
				log.Fatalf("Footnoter failed: %s", err)
			}
			if err := paragrapher(root); err != nil {
				log.Fatalf("Paragrapher failed: %s", err)
			}
			tout, err := RenderPage(base, root)
			if err != nil {
				log.Fatalf("Template failed: %s", err)
			}
			fullpath := filepath.Join(outdir, "people", uid)
			if err := os.MkdirAll(fullpath, 0750); err != nil {
				log.Fatalf("Unable to make directory %s", err)
			}
			fullpath = filepath.Join(fullpath, "index.html")
			if err := os.WriteFile(fullpath, []byte(tout), 0666); err != nil {
				log.Fatalf("couldn't write %q: %s", fullpath, err)
			}
		}
	}

	if server {
		serve(outdir)
	}
}
