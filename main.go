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

	//
	// Write OPR Birth Index
	//
	page := oprindex(db, "b")
	fullpath := filepath.Join("hugo/content", "indexes/opr-birth-index.html")
	log.Printf("Writing %q", fullpath)
	err := os.WriteFile(fullpath, []byte(page), 0666)
	if err != nil {
		log.Fatalf("couldn't write %q: %s", fullpath, err)
	}

	// Write OPR Death Index
	//
	page = oprindex(db, "d")
	fullpath = filepath.Join("hugo/content", "indexes/opr-death-index.html")
	log.Printf("Writing %q", fullpath)
	err = os.WriteFile(fullpath, []byte(page), 0666)
	if err != nil {
		log.Fatalf("couldn't write %q: %s", fullpath, err)
	}
	// Write OPR Marriage Index
	//
	page = oprindex(db, "m")
	fullpath = filepath.Join("hugo/content", "indexes/opr-marriage-index.html")
	log.Printf("Writing %q", fullpath)
	err = os.WriteFile(fullpath, []byte(page), 0666)
	if err != nil {
		log.Fatalf("couldn't write %q: %s", fullpath, err)
	}
	// Write SP Birth Index
	page = spindex(db, "b")
	fullpath = filepath.Join("hugo/content", "indexes/statutory-birth-index.html")
	log.Printf("Writing %q", fullpath)
	err = os.WriteFile(fullpath, []byte(page), 0666)
	if err != nil {
		log.Fatalf("couldn't write %q: %s", fullpath, err)
	}
	// Write SP Death Index
	page = spindex(db, "d")
	fullpath = filepath.Join("hugo/content", "indexes/statutory-death-index.html")
	log.Printf("Writing %q", fullpath)
	err = os.WriteFile(fullpath, []byte(page), 0666)
	if err != nil {
		log.Fatalf("couldn't write %q: %s", fullpath, err)
	}

	// Write SP Marriage Index
	page = spindex(db, "m")
	fullpath = filepath.Join("hugo/content", "indexes/statutory-marriage-index.html")
	log.Printf("Writing %q", fullpath)
	err = os.WriteFile(fullpath, []byte(page), 0666)
	if err != nil {
		log.Fatalf("couldn't write %q: %s", fullpath, err)
	}

	page = indexRoots(db, roots, "Roots")
	tmap := tagmap(db)
	for tag, uids := range tmap {
		page := indexRoots(db, uids, tag)

		tag = strings.ToLower(tag)
		tag = strings.ReplaceAll(tag, " ", "-")

		// index roots
		fullpath = filepath.Join("hugo/content", "tags/"+tag+".html")
		log.Printf("Writing %q", fullpath)
		err = os.WriteFile(fullpath, []byte(page), 0666)
		if err != nil {
			log.Fatalf("couldn't write %q: %s", fullpath, err)
		}
	}

	baset, err := os.ReadFile("baseof.html")
	if err != nil {
		log.Fatalf("can't find template: %s", err)
	}
	base, err := CreatePageTemplate(string(baset), RenderFunc(renderFuncs()))
	if err != nil {
		log.Fatalf("Unable to create template: %s", err)
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
				log.Fatalf("UMable to make directory %s", err)
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
