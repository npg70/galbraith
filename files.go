package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
	"text/template"
)

// in the layout directory
//
//	get a list of all paths
func getDirectories(root string) []string {
	out := []string{}
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			rel, err := filepath.Rel(root, path)
			if err != nil {
				panic("WTF")
			}
			log.Printf("PATH = %s, RELPATH= %s", path, rel)
			out = append(out, rel)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return out
}

type TemplateRouter map[string]*template.Template

func (t TemplateRouter) ExecuteTemplate(wr io.Writer, name string, data any) error {
	dir, file := filepath.Split(name)
	dir = strings.TrimSuffix(dir, string(filepath.Separator))
	// fix some asymmetries
	if dir == "" {
		dir = "."
	}
	//log.Printf("GOT TEMPLATE %q  Dir=%s file=%s", name, dir, file)
	base, ok := t[dir]
	if !ok {
		for k, v := range t {
			log.Printf("GOT dir=%s --> template %s", k, v.Name())
		}
		return fmt.Errorf("Could not file with dir=%q file=%q", dir, file)
	}
	return base.ExecuteTemplate(wr, file, data)
}

func templateMap(root string) (TemplateRouter, error) {
	// init
	out := make(TemplateRouter)

	// nothing!
	fmap := template.FuncMap{}

	//
	dirs := getDirectories(root)
	log.Printf("OUT = %d", len(dirs))
	for _, d := range dirs {
		t := template.New(d).Funcs(fmap)

		log.Printf("IN DIR: %s", d)
		parts := strings.Split(d, string(filepath.Separator))
		for i := 0; i <= len(parts); i++ {
			current := filepath.Join(parts[:i]...)
			templateGlob := filepath.Join(root, current, "*.html")
			log.Printf("Reading current=%q  templates: %q", current, templateGlob)
			if _, err := t.ParseGlob(templateGlob); err != nil {
				return out, err
				// typically empty directory
				//log.Printf("GOT ZERO TEMPLATES")
			}
		}
		out[d] = t
	}
	return out, nil
}
