package main

import (
	"testing"
)

// is
func TestLabelSubset(t *testing.T) {

	type test struct {
		a    string
		b    string
		want bool
	}

	tests := []test{
		{"", "", true},
		{"foo", "foo", true},
		{"", "foo", false},
		//{"foo", "", TBD},
		{"foo", "bar", false},
		{"foo", "bar:junk", false},
		{"foo:a", "bar:junk", false},
		{"foo:a", "foo", true},
		{"foo", "foo:b", false},
		{"foo:a", "foo:b", true},
		{"foo:bar:bing", "foo:bar", true},
	}

	for i, tc := range tests {
		got := labelSubset(tc.a, tc.b)
		if got != tc.want {
			t.Fatalf("case %d: %q vs. %q expected: %v, got %v", i, tc.a, tc.b, tc.want, got)
		}
	}
}
func TestLabelPath(t *testing.T) {

	type test struct {
		a    []string
		b    string
		want bool
	}

	tests := []test{
		{[]string{"foo"}, "bar", false},
		{[]string{"foo"}, "foo", true},
		{[]string{"foo", "bar"}, "bar", true},
		{[]string{"all", "foo:bar"}, "other", false},
	}
	for i, tc := range tests {
		got := inLabelPath(tc.a, tc.b)
		if got != tc.want {
			t.Fatalf("case %d: %q vs. %q expected: %v, got %v", i, tc.a, tc.b, tc.want, got)
		}
	}
}
