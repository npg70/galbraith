package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"text/template"
)

var inputFile string

func init() {
	flag.StringVar(&inputFile, "input", "", "input css template file name")
}

type Globals struct {
	Namespace string
	Prefix    string
	Resources map[string]map[string]string
	Params    map[string]string
}

func main() {
	flag.Parse()
	if inputFile == "" {
		flag.PrintDefaults()
		return
	}

	g := Globals{}

	defaults := make(map[string]map[string]string)
	properties := make(map[string]map[string]string)

	g.Namespace = "ng"
	g.Prefix = namespaceToPrefix(g.Namespace)
	vars := make(map[string]map[string]string)
	vars["root"] = make(map[string]string)

	funcMap := template.FuncMap{
		"rgb":      hexToRGB,
		"def":      MakeCssDefine(namespaceToPrefix(g.Namespace)),
		"var":      MakeCssVar(namespaceToPrefix(g.Namespace)),
		"default":  DefaultFunc(defaults),
		"property": DefaultFunc(properties),
	}

	tmpl := template.New(inputFile).Funcs(funcMap)

	tmpl, err := tmpl.ParseGlob("*.css")
	if err != nil {
		log.Fatalf("parseglob: %s", err)
	}

	head := strings.Builder{}
	body := strings.Builder{}
	// Run the template to verify the output.
	err = tmpl.Execute(&body, &g)
	if err != nil {
		log.Fatalf("body execution: %s", err)
	}

	// at this point we have loaded all the defaults
	// and all the overrides.  Merge them.
	MergeThemes(properties, defaults)
	g.Resources = properties

	headTemplate := tmpl.Lookup("head.css")
	err = headTemplate.Execute(&head, &g)
	if err != nil {
		log.Fatalf("head execution: %s", err)
	}

	// convert raw `var(foo)` into `var(--foo)` or `var(--prefix-foo)`
	fileout := addPrefixToVar(g.Namespace,
		head.String()+body.String())

	fmt.Println(fileout)
}
