package main

import (
	"testing"
)

func TestValues(t *testing.T) {
	tree := &tree{0, nil, nil}
	add(tree, 1)
	add(tree, 2)

	want := "012"
	got := tree.String()
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
