package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/net/html"

	tf "github.com/client9/tagfunctions"
)

func nodeExecute() func(*html.Node) (string, error) {
	tmp := renderFuncs()
	tagexec := tf.ExecuteFunc(tmp)

	toHTML := tf.RenderStringFunc(tf.RenderHTML)

	p1 := tf.Paragrapher{}
	p1.Tag = "p"

	p2 := tf.Paragrapher{}
	p2.Tag = "blockquote"

	return func(n *html.Node) (string, error) {
		if err := footnoter(n); err != nil {
			return "", fmt.Errorf("Footnoter failed: %s", err)
		}

		if err := p1.Execute(n); err != nil {
			return "", fmt.Errorf("Paragrapher for $p failed: %s", err)
		}

		if err := p2.Execute(n); err != nil {
			return "", fmt.Errorf("Paragrapher for $blockquote failed: %s", err)
		}
		if err := tagexec(n); err != nil {
			return "", fmt.Errorf("TagFunc failed: %s", err)
		}

		return toHTML(n)
	}
}

// writePage
//
//	tfsrc -- the tagfunction source code
//	treefn --
//	base -- golang html template, "the shell" of the page
//	outdir, outfile - duh   should be redone
func writePage(tfsrc string,
	base *template.Template,
	outdir string, outfile string) error {

	// this makes the tree from datafile
	p := tf.Tokenizer{}
	root := p.Parse(strings.NewReader(tfsrc))

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
	err = os.WriteFile(fullpath, []byte(tout), 0666)
	if err != nil {
		return fmt.Errorf("Error writing %q: %s", fullpath, err)
	}
	return nil
}
