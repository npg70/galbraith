package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	tf "github.com/client9/tagfunctions"
	"github.com/npg70/ssg"
	"gopkg.in/yaml.v3"
)

func Main2(config siteConfig, pages *[]ssg.ContentSource) error {
	tmpl, _ := templateMap("layouts")

	if err := LoadContent(config, pages); err != nil {
		return fmt.Errorf("Load content failed: %w", err)
	}
	if err := ssg.Execute(config, tmpl, *pages); err != nil {
		return fmt.Errorf("ssg.Execute failed: %w", err)
	}

	return nil
}

func LoadContent(config ssg.SiteConfig, out *[]ssg.ContentSource) error {
	cs := ssg.ContentSplitter{}
	cs.Register(ssg.HeadYaml)

	// for each file in sources
	// read it in,
	suffix := ".sh"
	contentDir := "content"
	err := filepath.WalkDir(contentDir, func(path string, d fs.DirEntry, err error) error {
		//log.Printf("WalkDir: got %s", path)
		// not sure how this works
		if err != nil {
			return err
		}

		// dont look at linux/mac dot dirs
		if d.IsDir() {
			if strings.HasPrefix(d.Name(), ".") {
				return filepath.SkipDir
			}
			return nil
		}

		if !strings.HasSuffix(path, suffix) {
			return nil
		}

		log.Printf("reading page file %s", path)
		raw, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading page file %s failed: %w", path, err)
		}

		htype, head, body := cs.Split(string(raw))
		if htype != "yaml" {
			// TBD on right behavior
			log.Fatalf("Expected YAML sample: got %q", htype)
		}

		page := make(ssg.ContentSourceConfig)
		if err := yaml.Unmarshal([]byte(head), &page); err != nil {
			log.Fatalf("Unable to un-yaml: %v", err)
		}
		if _, ok := page["TemplateName"]; !ok {
			page["TemplateName"] = "baseof.html"
		}
		// have: content/foo/bar/page.sh
		// want: foo/bar/page/index.html
		if _, ok := page["OutputFile"]; !ok {
			s := path[len(contentDir)+1:]
			s = strings.TrimSuffix(s, filepath.Ext(s))
			if d.Name() == "index.sh" {
				s += ".html"
			} else {
				s = filepath.Join(s, "index.html")
			}
			log.Printf("Setting outfile to %s", s)
			page["OutputFile"] = s
		}
		page["Content"] = body

		*out = append(*out, page)
		return nil
	})

	return err
}

func render(content string) (string, error) {

	// Parse raw into nodes
	p := tf.Tokenizer{}
	n := p.Parse(strings.NewReader(content))

	// CSV Tables has a non-standard parsing
	// do them first, so other stuff below works
	formatter := func(tag string, row int, col int) string {
		switch tag {
		case "wrapper":
			return "table-responsive"
		case "table":
			return "table table-borderless table-sm small ms-3"
		case "th", "td":
			return "text-body"
		}
		return ""
	}
	if err := tf.CsvTable(n, formatter); err != nil {
		return "", fmt.Errorf("CSV Tabler failed: %s", err)
	}
	// Add footnotes
	if err := footnoter(n); err != nil {
		return "", fmt.Errorf("Footnoter failed: %s", err)
	}

	// Convert custom nodes to HTML nodes
	tmp := renderFuncs()
	tagexec := tf.ExecuteFunc(tmp)
	if err := tagexec(n); err != nil {
		return "", fmt.Errorf("TagFunc failed: %w", err)
	}

	// Auto-split paragraphs
	p1 := tf.Paragrapher{}
	p1.Tag = "p"
	if err := p1.Execute(n); err != nil {
		return "", fmt.Errorf("Paragrapher for $p failed: %w", err)
	}

	// Auto-split blockquotes
	p2 := tf.Paragrapher{}
	p2.Tag = "blockquote"
	if err := p2.Execute(n); err != nil {
		return "", fmt.Errorf("Paragrapher for $blockquote failed: %w", err)
	}

	// Turn into HTML
	toHTML := tf.RenderStringFunc(tf.RenderHTML)
	return toHTML(n)
}
