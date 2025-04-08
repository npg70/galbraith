package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
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
	Themes    MapAny
	Root      MapAny
	Config    MapAny
	Arg       any
}

func (g *Globals) WithArg(v any) *Globals {
	g.Arg = v
	return g
}

func main() {
	flag.Parse()
	if inputFile == "" {
		flag.PrintDefaults()
		return
	}

	g := Globals{}

	themes := make(MapAny)
	root := make(MapAny)
	config := make(MapAny)
	g.Namespace = "ng"
	g.Prefix = namespaceToPrefix(g.Namespace)
	g.Config = config
	g.Themes = themes
	g.Root = root

	funcMap := template.FuncMap{
		"rgb":        hexToRGB,
		"def":        MakeCssDefine(namespaceToPrefix(g.Namespace)),
		"var":        MakeCssVar(namespaceToPrefix(g.Namespace)),
		"Theme":      OptionThemeFunc(themes),
		"Config":     OptionFunc(config),
		"Root":       OptionFunc(root),
		"Properties": ParseRules,
	}

	tmpl := template.New(inputFile).Funcs(funcMap).Funcs(FuncMapMath)

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

	headTemplate := tmpl.Lookup("head.css")
	err = headTemplate.Execute(&head, &g)
	if err != nil {
		log.Fatalf("head execution: %s", err)
	}

	fileout := head.String() + "\n" + body.String()

	// convert "--$" ==> "--" or "--foo-"
	fileout = addPrefixToDollar(g.Namespace, fileout)

	// strip out repeated newlines (from golang templates)
	var newlinesRun = regexp.MustCompile("\n\n+")
	fileout = newlinesRun.ReplaceAllString(fileout, "\n")

	// replace tabs with spaces
	fileout = strings.ReplaceAll(fileout, "\t", " ")
	// more than 2 spaces get turned into 1
	var spaceRun = regexp.MustCompile("  +")
	fileout = spaceRun.ReplaceAllString(fileout, " ")
	fmt.Println(strings.TrimSpace(fileout))
}
