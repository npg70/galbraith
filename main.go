package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func computeRoots() []string {
	db := make(Root)
	pmap := make(map[string]bool)
	files, err := filepath.Glob("./people/*.sh")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		uid := strings.TrimSuffix(filepath.Base(f), ".sh")
		_, nextg := db.generateOne(uid)
		for _, child := range nextg {
			pmap[child] = true
		}
		if _, ok := pmap[uid]; !ok {
			pmap[uid] = false
		}
	}

	out := []string{}
	for k, v := range pmap {
		if !v {
			out = append(out, k)
		}
	}
	return out
}
func init() {
	// flag stuff
}
func main() {
	outdir := ""
	rootsOnly := false
	flag.StringVar(&outdir, "out", "", "out directory")
	flag.BoolVar(&rootsOnly, "root", false, "run and display roots")
	flag.Parse()

	// soon to be obsolete
	log.Printf("Loading OPR Database")
	oprdb, err := LoadOPRBaptisms()
	if err != nil {
		log.Fatalf("can't load opr baptism", err)
	}

	db := make(Root)
	roots := computeRoots()
	if rootsOnly {
		return
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

			out := emitHTMLHugo(root, oprdb, singlePage)
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

			out := emitHTMLHugo(root, oprdb, singlePage)
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
