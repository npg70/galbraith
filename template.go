package main

import (
	"strings"
	"text/template"

	"golang.org/x/net/html"

	tf "github.com/client9/tagfunctions"
)

func CreatePageTemplate(base string, renderfn func(*html.Node) string) (*template.Template, error) {
	// TODO - load file or string into template
	fmap := template.FuncMap{
		"render": renderfn,
		"select": tf.Select,
	}
	t, err := template.New("base").Funcs(fmap).Parse(base)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func RenderPage(t *template.Template, n *html.Node) (string, error) {
	out := &strings.Builder{}
	err := t.Execute(out, n)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
