package main

import (
	"encoding/json"
	"log"
	"os"
)

// This creates a JSON array
// that could be used to make a full text index

type FtIndex struct {
	Id      int      `json:"id"`      // sequential int
	Doc     string   `json:"doc"`     // the doc to index
	Tags    []string `json:"tags"`    // tags.. tbd
	Content string   `json:"content"` // content to display
}

func fulltext(db Root) {
	idx := 0
	out := make([]FtIndex, 0, len(db))
	for _, p := range db {
		if p.Tags == nil {
			// flexsearch blows up on nil
			p.Tags = []string{}
		}
		out = append(out, FtIndex{idx, WriteFulltextDoc(p), p.Tags, WriteBio(p, 2)})
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
