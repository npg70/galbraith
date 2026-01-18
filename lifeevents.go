package main

import (
	"fmt"
	"io"
	"log"

	"github.com/client9/cmdconfig"
)

type Event struct {
	Type     string
	Date     Date
	Location string
	Name     string
	Ref      string
	Note     string
}

func (e *Event) UnmarshalText(text []byte) error {
	scan := cmdconfig.NewScanner(text)
	for {
		args, _, err := scan.Next()
		if args == nil && err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if len(args) != 2 {
			return fmt.Errorf("inside Event.%s expected two args, got %v", args[0], args)
		}
		log.Printf("EVENT: args %q", args)
		switch args[0] {
		case "date":
			e.Date.UnmarshalText([]byte(args[1]))
		case "location":
			e.Location = args[1]
		case "ref":
			e.Ref = args[1]
		case "note":
			e.Note = args[1]
		case "Name":
			e.Note = args[1]
		default:
			return fmt.Errorf("unknown Event command %v", args)
		}
	}
}

func (e Event) DatePlace() string {
	return fmt.Sprintf("%s at %s", e.Date, e.Location)
}
