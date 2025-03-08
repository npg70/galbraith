package main

import (
	"encoding/json"
	"log"
	"os"
)

// This creates a JSON array
// that could be used to make a full text index

type FtIndex struct {
	Id      int    `json:"id"`
	Page    string `json:"page"`
	Context string `json:"content"`
}

func fulltext(db Root) {
	idx := 0
	out := make([]FtIndex, 0, len(db))
	for key, p := range db {
		out = append(out, FtIndex{idx, key, WriteBio(p)})
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
