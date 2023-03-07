package main

import (
	"flag"
	"fmt"
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
	page := oprindex(db)
	fullpath := filepath.Join("hugo/content", "indexes/opr-birth-index.html")
	log.Printf("Writing %q", fullpath)
	err := os.WriteFile(fullpath, []byte(page), 0666)
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

	page = indexRoots(db, roots)
	fullpath = filepath.Join("hugo/content", "indexes/roots-index.html")
	log.Printf("Writing %q", fullpath)
	err = os.WriteFile(fullpath, []byte(page), 0666)
	if err != nil {
		log.Fatalf("couldn't write %q: %s", fullpath, err)
	}

	for _, rootid := range roots {
		log.Printf("ROOT---------> %s", rootid)
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

			out := emitHTMLHugo(root, singlePage)
			if outdir == "" {
				fmt.Println(out)
				continue
			}

			fullpath := filepath.Join(outdir, uid+suffix)
			if err := os.WriteFile(fullpath, []byte(out), 0666); err != nil {
				log.Fatalf("couldn't write %q: %s", fullpath, err)
			}
		}
	}
}
