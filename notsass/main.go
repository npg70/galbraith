package main

import (
	"flag"
	"log"
	"os"
	"text/template"
)

var inputFile string

func init() {
	flag.StringVar(&inputFile, "input", "", "input css template file name")
}

func main() {
	flag.Parse()
	if inputFile == "" {
		flag.PrintDefaults()
		return
	}
	funcMap := template.FuncMap{
		"rgb": hexToRGB,
	}

	tmpl := template.New(inputFile).Funcs(funcMap)

	tmpl, err := tmpl.ParseGlob("*.css")
	if err != nil {
		log.Fatalf("parseglob: %s", err)
	}

	// Run the template to verify the output.
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}
}
