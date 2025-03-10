package main

import (
	"encoding/json"
	"log"
	"maps"
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

// #argyll:campbeltown ==> #argyll #campbeltown
func tagsCompound(tags []string) []string {
	out := []string{}
	for _, e := range tags {
		// takes 'compound tags' in the form of #foo:bar:ding
		// and makes #foo #bar #ding
		parts := strings.Split(strings.ToLower(e), ":")
		out = append(out, parts...)
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

func tagsNormalize(tags []string) []string {
	return uniqueStrings(tagsAnd(tagsCompound(tags)))
}

func fulltext(db Root) {
	idx := 0
	out := make([]FtIndex, 0, len(db))
	for _, p := range db {
		out = append(out, FtIndex{
			idx,
			WriteFulltextDoc(p),
			tagsNormalize(p.Tags),
			WriteBio(p, 2)})
		idx += 1
	}

	dump, err := json.MarshalIndent(out, "  ", "  ")
	if err != nil {
		log.Fatalf("json dump failed: %s", err)
	}
	if err := os.WriteFile("fulltext.json", dump, 0644); err != nil {
		log.Fatalf("write of dump failed: %s", err)
	}
}
