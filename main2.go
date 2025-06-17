package main

import (
	"bytes"
	"fmt"
	"io"

	tf "github.com/client9/tagfunctions"
	"github.com/client9/ssg"
	"gopkg.in/yaml.v3"
)

// split input source into metadata and content
func SplitYaml(src []byte) ([]byte,[]byte) {
	return ssg.Splitter(ssg.HeadYaml, src)
}

// parse metadata as Yaml
func ParseYaml(s []byte) (ssg.ContentSourceConfig, error) {
	meta := ssg.ContentSourceConfig{}
	if err := yaml.Unmarshal(s, &meta); err != nil {
		return nil, fmt.Errorf("Unable to un-yaml: %v", err)
	}
	return meta, nil
}

func Main2(config ssg.SiteConfig, pages *[]ssg.ContentSourceConfig) error {
	// create page assembly templates
	pageTemplate, err := ssg.NewPageRender(config.TemplateDir, nil)
	if err != nil {
		return fmt.Errorf("Page Template failed: %v", err)
	}

	conf := ssg.SiteConfig{
		InputExt: ".sh",
		OutputExt: ".html",
		IndexSource : "index.sh",
	 	IndexDest: "index.html",
		Split: SplitYaml,
		Metaparser: ParseYaml,
		Pipeline: []ssg.Renderer{
			TagRender,
			pageTemplate,
		},
	}

	// do it
	err = ssg.Main2(conf, pages)
	if err != nil {
		return fmt.Errorf("Main failed: %s", err)
	}
	return nil
}

func TagRender(wr io.Writer, src []byte, data any) error {
	out, err := render(src)
	if err != nil {
		return err
	}
	_, err = wr.Write(out)
	return err
}

func render(content []byte) ([]byte, error) {

	// Parse raw into nodes
	p := tf.Tokenizer{}
	n := p.Parse(bytes.NewReader(content))

	// CSV Tables has a non-standard parsing
	// do them first, so other stuff below works
	formatter := func(tag string, row int, col int) string {
		switch tag {
		case "wrapper":
			return "table-responsive"
		case "table":
			return "table table-borderless table-sm small ms-3"
		}
		return ""
	}
	if err := tf.CsvTable(n, formatter); err != nil {
		return nil, fmt.Errorf("CSV Tabler failed: %s", err)
	}
	// Add footnotes
	if err := footnoter(n); err != nil {
		return nil, fmt.Errorf("Footnoter failed: %s", err)
	}

	// Convert custom nodes to HTML nodes
	tmp := renderFuncs()
	tagexec := tf.ExecuteFunc(tmp)
	if err := tagexec(n); err != nil {
		return nil, fmt.Errorf("TagFunc failed: %w", err)
	}

	// Auto-split paragraphs
	p1 := tf.Paragrapher{}
	p1.Tag = "p"
	if err := p1.Execute(n); err != nil {
		return nil, fmt.Errorf("Paragrapher for $p failed: %w", err)
	}

	// Auto-split blockquotes
	p2 := tf.Paragrapher{}
	p2.Tag = "blockquote"
	if err := p2.Execute(n); err != nil {
		return nil, fmt.Errorf("Paragrapher for $blockquote failed: %w", err)
	}

	// Turn into HTML
	toHTML := tf.RenderStringFunc(tf.RenderHTML)
	return toHTML(n)
}
