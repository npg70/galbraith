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

// may or may not be needed for CORS + Fonts
func addHeaders(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".woff2") {
			w.Header().Add("Access-Control-Allow-Origin", "http://localhost:1313")
		}
		fs.ServeHTTP(w, r)
	}
}

func serve(outdir string) {
	fmt.Println("Starting")
	http.Handle("/galbraith/", addHeaders(http.StripPrefix("/galbraith/", http.FileServer(http.Dir(outdir)))))
	if err := http.ListenAndServe("localhost:1313", nil); err != nil {
		log.Fatalf("Unable to start server: %s", err)
	}
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
	base, err := CreatePageTemplate(string(baset), nodeExecute())
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
		if err := writePage(string(raw), base, outdir, filepath.Join("sources", outfile)); err != nil {
			log.Fatalf("Source generation failed: %s", err)
		}
	}
	writePage(indexSources(), base, outdir, "sources")

	writePage(indexIndex(), base, outdir, "indexes")

	writePage(oprindex(db, "b"), base, outdir, "indexes/opr-birth-index")

	writePage(oprindex(db, "d"), base, outdir, "indexes/opr-death-index")

	writePage(oprindex(db, "m"), base, outdir, "indexes/opr-marriage-index")

	writePage(spindex(db, "b"), base, outdir, "indexes/statutory-birth-index")

	writePage(spindex(db, "d"), base, outdir, "indexes/statutory-death-index")

	writePage(spindex(db, "m"), base, outdir, "indexes/statutory-marriage-index")

	// TODO: can we simplify and reuse something already?
	tmap := tagmap(db)
	writePage(tagIndex(tmap), base, outdir, "tags")

	pages := tagStart(db)
	for _, tp := range pages {
		tagpath := makeTagFile(tp.path)
		if err := writePage(indexRoots2(db, tp), base, outdir, filepath.Join("tags", tagpath)); err != nil {
			log.Fatalf("WritePage for tags failed")
		}
	}

	// create basic markup to html func
	// write out each person
	count := 0
	for _, rootid := range roots {
		kidsq := []string{rootid}
		db.LoadAll(rootid)
		// one person, one page
		for len(kidsq) > 0 {
			uid := kidsq[0]
			kidsq = kidsq[1:]
			doc, nextg := db.generateOne(uid)
			kidsq = append(kidsq, nextg...)
			count++
			if err := writePage(doc, base, outdir, filepath.Join("people", uid)); err != nil {
				log.Fatalf("Writepage for %s - %s", uid, err)
			}
		}
	}
	if server {
		serve(outdir)
	}
}
