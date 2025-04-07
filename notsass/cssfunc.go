package main

import (
	"fmt"
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

func ParseRules(rules string) (map[string]any, error) {
	out := make(map[string]any)
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

// Options
//
//	returns the backing map[string]any allowing full use of
//	golang template functions (iteration, index)
//
// Options val:string
//
//	Evaluates val as a "properties file", over-writing any values
//
// Options (map[string]any)
//
//	copies the map, over-writing any values
//
// Options key1:string val1:any ...
//
// Options "theme" | Options
func OptionFunc(omap MapAny) func(...any) (any, error) {
	return func(args ...any) (any, error) {
		return omap.fromArgs(args...)
	}
}

// OptionTheme is a map of "themes" to MapAnys
// A theme is just a collection of related items.
//
// If a theme does not exist, a new MapAny is created.
func OptionThemeFunc(omap MapAny) func(...any) (any, error) {
	return func(args ...any) (any, error) {
		if len(args) == 0 {
			return omap, nil
		}
		theme, ok := args[0].(string)
		if !ok {
			return nil, fmt.Errorf("expected string name for theme, got %v", args[0])
		}
		val, ok := omap[theme]
		if !ok {
			val = make(MapAny)
			omap[theme] = val
		}
		return val.(MapAny).fromArgs(args[1:]...)
	}
}

type MapAny map[string]any

func (m MapAny) fromArgs(args ...any) (any, error) {
	if len(args) == 0 {
		return m, nil
	}
	if len(args) == 1 {
		switch v := args[0].(type) {
		case string:
			dict, err := ParseRules(v)
			if err != nil {
				return "", err
			}
			for k, prop := range dict {
				m[k] = prop
			}
			return "", nil
		case MapAny:
			for k, prop := range v {
				m[k] = prop
			}
			return "", nil
		case map[string]any:
			for k, prop := range v {
				m[k] = prop
			}
			return "", nil
		default:
			return "", fmt.Errorf("Expected a string or dict, got %v", v)
		}
	}
	if len(args)%2 != 0 {
		return "", fmt.Errorf("expected 0,1 or even number of args, got %d", len(args))
	}
	for i := 0; i < len(args); i += 2 {
		key, ok := args[i].(string)
		if !ok {
			return "", fmt.Errorf("expected string for key %v at pos %d", args[i], i)
		}
		m[key] = args[i+1]
	}
	return "", nil
}
