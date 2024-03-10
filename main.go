package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	// flag stuff
}

func writeIndexPage() {
	/* set up base template */
	baset, err := os.ReadFile("baseof.html")
	if err != nil {
		log.Fatalf("can't find template: %s", err)
	}
	base, err := CreatePageTemplate(string(baset), RenderFunc(renderFuncs()))
	if err != nil {
		log.Fatalf("Unable to create template: %s", err)
	}

	p := Tokenizer{}
	root := p.Parse(strings.NewReader(indexIndex()))
	tout, err := RenderPage(base, root)
	if err != nil {
		log.Fatalf("Template failed: %s", err)
	}
	fullpath := "hugo/static/indexes"
	if err := os.MkdirAll(fullpath, 0750); err != nil {
		log.Fatalf("Unable to make directory %s", err)
	}
	fullpath = filepath.Join(fullpath, "index.html")
	log.Printf("Writing %q", fullpath)
	err = os.WriteFile(fullpath, []byte(tout), 0666)
	if err != nil {
		log.Fatalf("couldn't write %q: %s", fullpath, err)
	}
}

func main() {
	outdir := ""
	rootsOnly := false
	flag.StringVar(&outdir, "out", "", "out directory")
	flag.BoolVar(&rootsOnly, "root", false, "run and display roots")
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
		outfile := strings.TrimSuffix(filepath.Base(s), suffix) + ".html"
		outpath := filepath.Join("hugo", "content", "sources", outfile)

		raw, err := os.ReadFile(s)
		page := Execute(string(raw), renderFuncs())
		log.Printf("Writing %q", outpath)
		err = os.WriteFile(outpath, []byte(page), 0666)
		if err != nil {
			log.Fatalf("couldn't write %q: %s", outpath, err)
		}
	}

	writeIndexPage()

	//
	// Write OPR Birth Index
	//
	page := Execute(oprindex(db, "b"), renderFuncs())
	fullpath := filepath.Join("hugo/content", "indexes/opr-birth-index.html")
	log.Printf("Writing %q", fullpath)
	err = os.WriteFile(fullpath, []byte(page), 0666)
	if err != nil {
		log.Fatalf("couldn't write %q: %s", fullpath, err)
	}

	// Write OPR Death Index
	//
	page = Execute(oprindex(db, "d"), renderFuncs())
	fullpath = filepath.Join("hugo/content", "indexes/opr-death-index.html")
	log.Printf("Writing %q", fullpath)
	err = os.WriteFile(fullpath, []byte(page), 0666)
	if err != nil {
		log.Fatalf("couldn't write %q: %s", fullpath, err)
	}
	// Write OPR Marriage Index
	//
	page = Execute(oprindex(db, "m"), renderFuncs())
	fullpath = filepath.Join("hugo/content", "indexes/opr-marriage-index.html")
	log.Printf("Writing %q", fullpath)
	err = os.WriteFile(fullpath, []byte(page), 0666)
	if err != nil {
		log.Fatalf("couldn't write %q: %s", fullpath, err)
	}
	// Write SP Birth Index
	page = Execute(spindex(db, "b"), renderFuncs())
	fullpath = filepath.Join("hugo/content", "indexes/statutory-birth-index.html")
	log.Printf("Writing %q", fullpath)
	err = os.WriteFile(fullpath, []byte(page), 0666)
	if err != nil {
		log.Fatalf("couldn't write %q: %s", fullpath, err)
	}
	// Write SP Death Index
	page = Execute(spindex(db, "d"), renderFuncs())
	fullpath = filepath.Join("hugo/content", "indexes/statutory-death-index.html")
	log.Printf("Writing %q", fullpath)
	err = os.WriteFile(fullpath, []byte(page), 0666)
	if err != nil {
		log.Fatalf("couldn't write %q: %s", fullpath, err)
	}

	// Write SP Marriage Index
	page = Execute(spindex(db, "m"), renderFuncs())
	fullpath = filepath.Join("hugo/content", "indexes/statutory-marriage-index.html")
	log.Printf("Writing %q", fullpath)
	err = os.WriteFile(fullpath, []byte(page), 0666)
	if err != nil {
		log.Fatalf("couldn't write %q: %s", fullpath, err)
	}

	tmap := tagmap(db)

	// make descriptive home page for tags and labels
	p := Tokenizer{}
	root := p.Parse(strings.NewReader(tagIndex(tmap)))
	tout, err := RenderPage(base, root)
	if err != nil {
		log.Fatalf("Template failed: %s", err)
	}
	fullpath = "hugo/static/tags"
	if err := os.MkdirAll(fullpath, 0750); err != nil {
		log.Fatalf("Unable to make directory %s", err)
	}
	fullpath = filepath.Join(fullpath, "index.html")
	log.Printf("Writing %q", fullpath)
	err = os.WriteFile(fullpath, []byte(tout), 0666)
	if err != nil {
		log.Fatalf("couldn't write %q: %s", fullpath, err)
	}

	pages := tagStart(db)
	for _, tp := range pages {
		page := p.Parse(strings.NewReader(indexRoots2(db, tp)))
		tout, err := RenderPage(base, page)
		if err != nil {
			log.Fatalf("Template failed: %s", err)
		}
		tagpath := makeTagFile(tp.path)
		fullpath = filepath.Join("hugo/static", "tags", tagpath)
		if err := os.MkdirAll(fullpath, 0750); err != nil {
			log.Fatalf("Unable to make directory %s", err)
		}
		fullpath = filepath.Join(fullpath, "index.html")
		log.Printf("Writing %q", fullpath)
		if err := os.WriteFile(fullpath, []byte(tout), 0666); err != nil {
			log.Fatalf("couldn't write %q: %s", fullpath, err)
		}
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
}
