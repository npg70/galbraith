package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type test struct {
		input string
		args  []string
		body  string
	}

	tests := []test{
		{
			input: "name John Brown",
			args:  []string{"name", "John", "Brown"},
		},
		{
			input: "name 'John' \"Brown\"",
			args:  []string{"name", "John", "Brown"},
		},
		{
			input: "name John Brown\n\n\n",
			args:  []string{"name", "John", "Brown"},
		},
		{
			input: "   name   John   Brown   \n\n\n",
			args:  []string{"name", "John", "Brown"},
		},
		{
			// check to make sure flags are ok
			input: "name -first John -last Brown",
			args:  []string{"name", "-first", "John", "-last", "Brown"},
		},
		{
			// check to make sure key=val are ok
			input: "name first=John last=Brown",
			args:  []string{"name", "first=John", "last=Brown"},
		},
		{
			// check arg quotes
			input: `func first 'second' "third"`,
			args:  []string{"func", "first", "second", "third"},
		},
		{
			// check to make sure flags are ok
			input: "name -first 'john'",
			args:  []string{"name", "-first", "john"},
		},
		{
			// check arg quotes
			input: `name first="mary ann" last="brown"`,
			args:  []string{"name", "first=mary ann", "last=brown"},
		},
		{
			// check arg quotes
			input: `name -first "mary ann"`,
			args:  []string{"name", "-first", "mary ann"},
		},
		{
			// check arg quotes
			input: `name -first 'mary ann'   `,
			args:  []string{"name", "-first", "mary ann"},
		},
		{
			// check arg quotes
			input: `name first='mary ann' last='brown'`,
			args:  []string{"name", "first=mary ann", "last=brown"},
		},
		{
			// check arg quotes
			input: `name "first"='mary ann' "last"='brown'`,
			args:  []string{"name", "first=mary ann", "last=brown"},
		},
		{
			input: "name John Brown {\napple {banana}}",
			args:  []string{"name", "John", "Brown"},
			body:  "\napple {banana}",
		},
	}

	for i, tc := range tests {
		s := NewScanner([]byte(tc.input))
		args, body, err := s.Next()
		if err != nil {
			t.Fatalf("case %d, got error %v", i, err)
		}
		if !reflect.DeepEqual(args, tc.args) {
			t.Fatalf("case %d, expected args %v got %v", i, tc.args, args)
		}
		if tc.body != body {
			t.Fatalf("case %d, expected body %q got %q", i, tc.body, body)
		}
	}
}
