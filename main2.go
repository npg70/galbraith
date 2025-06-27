package main

import (
	"fmt"
	"io"

	"github.com/client9/ssg"
	tf "github.com/client9/tagfunctions"
)

// parse metadata
func ParseMeta(s []byte) (ssg.ContentSourceConfig, error) {
	meta := ssg.ContentSourceConfig{}
	if err := ssg.EmailUnmarshal(s, meta); err != nil {
		return nil, fmt.Errorf("unable to parse metadata: %v", err)
	}
	return meta, nil
}

func Main2(config ssg.SiteConfig, outputDir string, pages *[]ssg.ContentSourceConfig) error {

	conf := ssg.SiteConfig{
		InputExt:    ".sh",
		OutputExt:   ".html",
		IndexSource: "index.sh",
		IndexDest:   "index.html",
		MetaSplit:    ssg.MetaSplitEmail,
		MetaParser:   ParseMeta,
		Pipeline: []ssg.Renderer{
			TagRender,
			ssg.Must(ssg.NewPageRender("layouts", nil)),
			ssg.WriteOutput(outputDir),
		},
	}

	if err := ssg.Main2(conf, pages); err != nil {
		return fmt.Errorf("main failed: %s", err)
	}
	return nil
}

func TagRender(wr io.Writer, src io.Reader, data any) error {
	// Parse raw into nodes
	p := tf.Tokenizer{}
	n := p.Parse(src)

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
		return fmt.Errorf("csv Tabler failed: %s", err)
	}
	// Add footnotes
	if err := footnoter(n); err != nil {
		return fmt.Errorf("footnoter failed: %s", err)
	}

	// Convert custom nodes to HTML nodes
	tmp := renderFuncs()
	tagexec := tf.ExecuteFunc(tmp)
	if err := tagexec(n); err != nil {
		return fmt.Errorf("tagfunc failed: %w", err)
	}

	// Auto-split paragraphs
	p1 := tf.Paragrapher{}
	p1.Tag = "p"
	if err := p1.Execute(n); err != nil {
		return fmt.Errorf("paragrapher for $p failed: %w", err)
	}

	// Auto-split blockquotes
	p2 := tf.Paragrapher{}
	p2.Tag = "blockquote"
	if err := p2.Execute(n); err != nil {
		return fmt.Errorf("paragrapher for $blockquote failed: %w", err)
	}

	return tf.RenderHTML(wr, n)
}
