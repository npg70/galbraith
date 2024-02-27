package main

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TitleCompound(s string) string {
	parts := strings.Split(s, ",")
	for i, p := range parts {
		parts[i] = TitlePhrase(p)
	}
	return strings.Join(parts, ", ")
}

// TitleCase cases a phrase or name
// ignoring small English works typically not capitalized

var keepLower = map[string]bool{
	"a":   true,
	"an":  true,
	"and": true,
	"the": true,
	"by":  true,
	"to":  true,
}
var allUpper = map[string]bool{
	"usa": true,
	"uk":  true,
}

func TitlePhrase(s string) string {
	parts := strings.Fields(strings.ToLower(s))

	for i, p := range parts {
		switch {
		case keepLower[p]:
			continue
		case allUpper[p]:
			parts[i] = strings.ToUpper(p)
		default:
			parts[i] = Title(p)
		}
	}
	return strings.Join(parts, " ")
}

func Title(s string) string {
	caser := cases.Title(language.English)
	return caser.String(s)
}

func ChildPlural(num int) string {
	if num == 1 {
		return "Child"
	}
	return "Children"
}

func Ordinal(i int) (ordinal string) {
	switch i {
	case 1:
		ordinal = "first"
	case 2:
		ordinal = "second"
	case 3:
		ordinal = "third"
	case 4:
		ordinal = "fourth"
	default:
		ordinal = "next"
	}
	return ordinal
}

func AllMostSome(s []string) (string, string) {
	if len(s) < 2 {
		return "", ""
	}

	// normalize cases

	for i, p := range s {
		s[i] = TitleCompound(p)
	}

	count := 0
	first := s[0]
	for _, n := range s {
		if n == first {
			count++
		}
	}
	if count == len(s) {
		return first, "all"
	}

	// that is count >= (2/3) *len(s)
	if count*3 >= len(s)*2 {
		return first, "except"
	}
	return "", ""
}
