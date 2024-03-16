package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func writePage(rootDir string, outdir string, base *template.Template, tpl string) error {
	p := Tokenizer{}
	root := p.Parse(strings.NewReader(tpl))
	tout, err := RenderPage(base, root)
	if err != nil {
		log.Fatalf("Template failed: %s", err)
	}
	fullpath := filepath.Join(rootDir, outdir)
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
