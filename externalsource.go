package main

import (
	"fmt"
	"io"
	"strings"
)

type SourceLink struct {
	order    int // for ordering
	RecordID string
	URL      string
	Name     string // for display
}

type ExternalSource map[string]SourceLink

func (e ExternalSource) List() []SourceLink {
	out := make([]SourceLink, len(e))
	i := 0
	for _, val := range e {
		out[i] = val
		i++
	}
	//TODO Sort by ordering
	return out
}

func (e ExternalSource) UnmarshalText(text []byte) error {
	scan := NewScanner(text)
	for {
		args, _, err := scan.Next()
		if args == nil && err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if len(args) != 2 {
			return fmt.Errorf("expected two args, got %v", args)
		}
		id := strings.ToLower(args[0])
		switch id {
		case "familysearch":
			e[id] = SourceLink{
				order:    10,
				RecordID: args[1],
				URL:      "https://www.familysearch.org/tree/person/details/" + args[1],
				Name:     "FamilySearch",
			}
		case "findagrave":
			e[id] = SourceLink{
				order:    20,
				RecordID: args[1],
				URL:      "https://www.findagrave.com/memorial/" + args[1],
				Name:     "Find a Grave",
			}
		case "billiongraves":
			e[id] = SourceLink{
				order:    50,
				RecordID: args[1],
				URL:      "https://billiongraves.com/grave/" + args[1],
				Name:     "Billion Graves",
			}
		case "wikitree":
			e[id] = SourceLink{
				order:    40,
				RecordID: args[1],
				URL:      "https://www.wikitree.com/wiki/" + args[1],
				Name:     "WikiTree",
			}
		case "geni":
			e[id] = SourceLink{
				order:    60,
				RecordID: args[1],
				URL:      "https://www.geni.com/people/" + args[1],
				Name:     "Geni",
			}
		case "cgs":
			// Dana Love's quasi-official Clan Galbraith Society tree
			e[id] = SourceLink{
				order:    30,
				RecordID: args[1],
				URL:      "https://www.ancestry.com/family-tree/person/tree/112802735/person/" + strings.Trim(args[1], "/") + "/",
				Name:     "CGS",
			}
		case "ancestry":
			e[id] = SourceLink{
				order:    30,
				RecordID: args[1],
				URL:      "https://www.ancestry.com/family-tree/person/" + strings.Trim(args[1], "/") + "/",
				Name:     "Ancestry",
			}
		default:
			return fmt.Errorf("Unknown external source: %q", args[0])
		}
	}
}
