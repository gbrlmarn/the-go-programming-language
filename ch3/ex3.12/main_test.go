package main

import (
	"testing"
)

func TestAnagram(t *testing.T) {
	tests := []struct {
		s1, s2 string
		want   bool
	}{
		{"asdf", "fdsa", true},
		{"asdf", "fdsb", false},
	}
	for _, test := range tests {
		got := anagram(test.s1, test.s2)
		if got != test.want {
			t.Errorf("anagram(%q, %q), got %v, want %v", test.s1, test.s2, got, test.want)
		}
	}
}
