package main

import (
	"encoding/json"
	"log"
	"maps"
	"net/url"
	"os"
	"slices"
	"sort"
	"strings"
)

// This creates a JSON array
// that could be used to make a full text index

type FtIndex struct {
	Id      int      `json:"id"`      // sequential int
	Doc     string   `json:"doc"`     // the doc to index
	Tags    []string `json:"tags"`    // tags.. tbd
	Content string   `json:"content"` // content to display
}

func uniqueStrings(input []string) []string {
	// Create a map to track unique strings.
	uniqueMap := make(map[string]struct{})
	// Iterate over the input slice and add unique strings to the map.
	for _, str := range input {
		uniqueMap[str] = struct{}{}
	}
	// Extract the keys from the map, which are the unique strings, and collect them into a slice.
	out := slices.Collect(maps.Keys(uniqueMap))
	if out == nil {
		// we need a real slice for flexsearch
		return []string{}
	}

	// sort for various reasons
	sort.Strings(out)
	return out
}

// #argyll:campbeltown ==> #argyll #argyll_campbeltown
func tagsCompound(tags []string) []string {
	out := []string{}
	for _, e := range tags {
		// takes 'compound tags' in the form of #foo:bar:ding
		// and makes #foo #foo_bar #foo_bar_ding
		parts := strings.Split(strings.ToLower(e), ":")
		for i := 0; i < len(parts); i++ {
			newtag := strings.Join(parts[0:i+1], "_")
			out = append(out, newtag)
		}
	}
	return out
}

// hand tags "#saddell and skipness" ==> #saddell #skipness
func tagsAnd(tags []string) []string {
	out := []string{}
	for _, e := range tags {
		parts := strings.Split(e, " and ")
		out = append(out, parts...)
	}
	return out
}

// for now just remove them
func tagsSpaces(tags []string) []string {
	for i, e := range tags {
		tags[i] = strings.ReplaceAll(e, " ", "")
	}
	return tags
}
func tagsNormalize(tags []string) []string {
	return uniqueStrings(tagsSpaces(tagsAnd(tagsCompound(tags))))
}

// given a list of tags, return an encoded string for
// creating a raw URL
func tagsQueryString(tags []string) string {
	out := []string{}
	for _, t := range tags {
		out = append(out, "#"+t)
	}
	return url.QueryEscape(strings.Join(out, " "))
}

// add a tag for major family names
func tagAddFamilyName(db Root) {
	for _, p := range db {
		name := strings.ToLower(p.FullName())
		switch {
		case strings.Contains(name, "galbreath") || strings.Contains(name, "galbraith"):
			p.Tags = append(p.Tags, "galbraith")
		case strings.Contains(name, "mcnair") || strings.Contains(name, "macnair"):
			p.Tags = append(p.Tags, "mcnair")
		case strings.Contains(name, "harvey") || strings.Contains(name, "harvie"):
			p.Tags = append(p.Tags, "harvey")
		case strings.Contains(name, "fleeming") || strings.Contains(name, "fleming") || strings.Contains(name, "flemming"):
			p.Tags = append(p.Tags, "fleeming")
		case strings.Contains(name, "mitchel"):
			p.Tags = append(p.Tags, "mitchell")
		case strings.Contains(name, "greenlee"):
			p.Tags = append(p.Tags, "greenlees")
		case strings.Contains(name, "armour"):
			p.Tags = append(p.Tags, "armour")
		}
	}
}

func fulltext(db Root) {
	// get all keys, but put them into determinant order
	dbkeys := slices.Collect(maps.Keys(db))
	slices.Sort(dbkeys)

	idx := 0
	out := make([]FtIndex, 0, len(dbkeys))
	for _, key := range dbkeys {
		p := db[key]
		p.Tags = tagsNormalize(p.Tags)
		out = append(out, FtIndex{
			idx,
			WriteFulltextDoc(p),
			p.Tags,
			WriteBio(p, 2)})
		idx += 1
	}

	dump, err := json.MarshalIndent(out, "  ", "  ")
	if err != nil {
		log.Fatalf("json dump failed: %s", err)
	}
	if err := os.WriteFile("./public/fulltext.json", dump, 0644); err != nil {
		log.Fatalf("write of dump failed: %s", err)
	}
}
