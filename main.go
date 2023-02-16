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
	flag.StringVar(&outdir, "out", "", "out directory")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatalf("Expected UID as first arg")
	}

	// root is the full path and suffix
	root := args[0]
	// remove any directory, and remove trailing suffix
	rootid := strings.TrimSuffix(filepath.Base(root), ".sh")

	db := make(Root)
	kidsq := []string{rootid}
	db.LoadAll(rootid)
	suffix := ".html"
	singlePage := false

	log.Printf("Loading OPR Database")
	oprdb, err := LoadOPRBaptisms()
	if err != nil {
		log.Fatalf("can't load opr baptism", err)
	}
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
		log.Printf("Writing %q", fullpath)
		if err := os.WriteFile(fullpath, []byte(out), 0666); err != nil {
			log.Fatalf("couldn't write %q: %s", fullpath, err)
		}
	}
}
