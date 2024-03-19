package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/net/html"
)

type TreeFunc func(*html.Node) error

// some post-processing of people data
func peopleTree(root *html.Node) error {
	if err := footnoter(root); err != nil {
		return fmt.Errorf("Footnoter failed: %s", err)
	}
	if err := paragrapher(root); err != nil {
		return fmt.Errorf("Paragrapher failed: %s", err)
	}
	return nil
}

func writePage(tpl string, treefn TreeFunc,
	base *template.Template, outdir string, outfile string) error {

	// this makes the tree from datafile
	p := Tokenizer{}
	root := p.Parse(strings.NewReader(tpl))

	// any tree editing
	if treefn != nil {
		if err := treefn(root); err != nil {
			return err
		}
	}

	// insert into template, generate output
	tout, err := RenderPage(base, root)
	if err != nil {
		return fmt.Errorf("Template failed: %s", err)
	}

	// write output
	fullpath := filepath.Join(outdir, outfile)
	if err := os.MkdirAll(fullpath, 0750); err != nil {
		return err
	}
	fullpath = filepath.Join(fullpath, "index.html")
	log.Printf("Writing %q", fullpath)
	err = os.WriteFile(fullpath, []byte(tout), 0666)
	if err != nil {
		return err
	}
	return nil
}
