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

	page = indexRoots(db, roots, "Roots")
	tmap := tagmap(db)

	// make tag index
	p := Tokenizer{}
	root := p.Parse(strings.NewReader(tagIndex(tmap)))
	tout, err := RenderPage(base, root)
	if err != nil {
		log.Fatalf("Template failed: %s", err)
	}
	fullpath = "hugo/layouts/tags/list.html"
	log.Printf("Writing %q", fullpath)
	err = os.WriteFile(fullpath, []byte(tout), 0666)
	if err != nil {
		log.Fatalf("couldn't write %q: %s", fullpath, err)
	}

	// make a page per tag
	for tag, uids := range tmap {
		page := p.Parse(strings.NewReader(indexRoots(db, uids, tag)))
		tout, err := RenderPage(base, page)
		if err != nil {
			log.Fatalf("Template failed: %s", err)
		}
		tag = strings.ReplaceAll(strings.ToLower(tag), " ", "-")
		fullpath = filepath.Join("hugo/content", "tags", tag)
		if err := os.MkdirAll(fullpath, 0750); err != nil {
			log.Fatalf("Unable to make directory %s", err)
		}
		fullpath = filepath.Join(fullpath, "index.html")
		if err := os.WriteFile(fullpath, []byte(tout), 0666); err != nil {
			log.Fatalf("couldn't write %q: %s", fullpath, err)
		}
		// index roots
		log.Printf("Writing %q", fullpath)
		err = os.WriteFile(fullpath, []byte(tout), 0666)
		if err != nil {
			log.Fatalf("couldn't write %q: %s", fullpath, err)
		}
	}

	for _, rootid := range roots {
		kidsq := []string{rootid}
		db.LoadAll(rootid)
		suffix := ".html"
		singlePage := false

		doc := ""

		if singlePage {
			for len(kidsq) > 0 {
				uid := kidsq[0]
				kidsq = kidsq[1:]
				intermediate, nextg := db.generateOne(uid)
				kidsq = append(kidsq, nextg...)
				if len(kidsq) == 0 {
					break
				}
				doc += intermediate
			}

			p := Tokenizer{}
			root := p.Parse(strings.NewReader(doc))

			if err := footnoter(root); err != nil {
				log.Fatalf("Footnoter failed: %s", err)
			}

			if err := paragrapher(root); err != nil {
				log.Fatalf("Paragrapher failed: %s", err)
			}
			out := emitHTMLHugo(root, singlePage)
			if outdir != "" {
				fullpath := filepath.Join(outdir, "test"+suffix)
				log.Printf("Writing %q", fullpath)
				err := os.WriteFile(fullpath, []byte(out), 0666)
				if err != nil {
					log.Fatalf("couldn't write %q: %s", fullpath, err)
				}
			}
			return
		}

		// one person, one page
		for len(kidsq) > 0 {
			uid := kidsq[0]
			kidsq = kidsq[1:]
			doc, nextg := db.generateOne(uid)
			kidsq = append(kidsq, nextg...)

			//		fmt.Println(doc)
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

			//-----------
			fullpath := filepath.Join(outdir, uid)
			if err := os.MkdirAll(fullpath, 0750); err != nil {
				log.Fatalf("Unable to make directory %s", err)
			}
			fullpath = filepath.Join(fullpath, "index.html")
			if err := os.WriteFile(fullpath, []byte(tout), 0666); err != nil {
				log.Fatalf("couldn't write %q: %s", fullpath, err)
			}
			// -----------
			/*
				out := emitHTMLHugo(root, singlePage)
				if outdir == "" {
					fmt.Println(out)
					continue
				}

				fullpath = filepath.Join(outdir, uid+suffix)
				if err := os.WriteFile(fullpath, []byte(out), 0666); err != nil {
					log.Fatalf("couldn't write %q: %s", fullpath, err)
				}
			*/
		}
	}
}
