// Exercise 11.5: Extend TextSplit to use a table of inputs and expected outputs.

package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		s    string
		sep  string
		want int
	}{
		{"a:b:c", ":", 3},
		{"x,y,z", ",", 3},
		{"Hello mister", " ", 2},
		{"variable_name_with_underscore", "_", 4},
        {"Separe line\nby line", "\n", 2},
	}
	for _, test := range tests {
		got := len(strings.Split(test.s, test.sep))
		if got != test.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d", test.s, test.sep, got, test.want)
		}
	}
}
