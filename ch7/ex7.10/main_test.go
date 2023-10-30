package main

import (
	"sort"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	if IsPalindrome(sort.IntSlice(ints)) {
		t.Fatal(ints)
	}

	strs := []string{"hei", "you", "you", "hei"}
	if !IsPalindrome(sort.StringSlice(strs)) {
		t.Fatal(strs)
	}
}
