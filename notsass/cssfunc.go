package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func MakeCssDefine(prefix string) func(string, any) string {
	return func(name string, val any) string {
		return fmt.Sprintf("%s%s:%v", prefix, name, val)
	}
}

func MakeCssVar(prefix string) func(string) string {
	return func(name string) string {
		return fmt.Sprintf("var(%s%v)", prefix, name)
	}
}

var varStart = regexp.MustCompile("var" + regexp.QuoteMeta("(") + " *([^-])")

// Given an embedded "var(foo)", expand to
func addPrefixToVar(namespace string, s string) string {
	prefix := namespaceToPrefix(namespace)
	return varStart.ReplaceAllString(s, "var("+prefix+"$1")
}

// converts a namespace to a css custome var prefix
// "" --> "--"
// "ng" -- "--ng-"
func namespaceToPrefix(namespace string) string {
	prefix := "--"
	if namespace != "" {
		prefix += namespace + "-"
	}
	return prefix
}

func ParseRules(rules string) (map[string]string, error) {
	out := make(map[string]string)
	lines := strings.Split(rules, "\n")
	for _, line := range lines {
		// remove leading and trailing whitespace
		// remove optional ending ";"
		// skip blanks
		// skip single line comments - TODO -- do C-style?
		line = strings.TrimSpace(line)
		line = strings.TrimRight(line, ";")
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		// should be "foo:bar"
		// split into key and value
		key, value, found := strings.Cut(line, ":")
		if !found {
			return nil, fmt.Errorf("On line %q, didn't find a ':'", line)
		}
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		out[key] = value
	}
	return out, nil
}

func MergeThemes(vals map[string]map[string]string, defaults map[string]map[string]string) {
	for theme, _ := range defaults {
		if _, ok := vals[theme]; !ok {
			vals[theme] = make(map[string]string)
		}
		MergeDefaults(vals[theme], defaults[theme])
	}
}

// Mutates input
func MergeDefaults(vals map[string]string, defaults map[string]string) {
	for k, v := range defaults {
		if _, ok := vals[k]; !ok {
			vals[k] = v
		}
	}
}

func MakeSetVars(vars map[string]string) func(string) error {
	return func(rules string) error {
		rulemap, err := ParseRules(rules)
		if err != nil {
			return err
		}
		// OVER-WRITE
		for k, v := range rulemap {
			vars[k] = v
		}
		return nil
	}

}
func MakeDefaultVar(vars map[string]string) func(string) error {
	return func(rules string) error {
		defaults, err := ParseRules(rules)
		if err != nil {
			return err
		}
		// ADD IF MISSING
		MergeDefaults(vars, defaults)
		return nil
	}
}

func DefaultFunc(defaults map[string]map[string]string) func(string, string) (string, error) {
	// TODO change sig to "rules ...string" and process each separately
	return func(name string, rules string) (string, error) {
		if defaults[name] == nil {
			defaults[name] = make(map[string]string)
		}
		orig := defaults[name]
		lines, err := ParseRules(rules)
		if err != nil {
			return "", err
		}
		// add or overwrite existing values
		for k, v := range lines {
			log.Printf("%s %s=%s", name, k, v)
			orig[k] = v
		}
		return "", nil
	}
}
