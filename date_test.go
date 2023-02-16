package main

import "testing"

func TestDate(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{"Jul 1729", "Jul 1729"},
		{"July 2020", "Jul 2020"},
		{"3 July 2020", "3 Jul 2020"},
		{"03.Jul.2020", "3 Jul 2020"},
		{"3-JUL-2020", "3 Jul 2020"},
	}
	for _, tc := range tests {
		d, _ := ParseDate(tc.input)
		got := d.DayMonYear()
		if got != tc.want {
			t.Fatalf("expected: %v, got %v", tc.want, got)
		}
	}
}
