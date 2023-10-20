package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var s IntSet
	s.Add(1)
	s.Add(2)
	if !s.Has(1) || !s.Has(2) {
		t.Fatal(s)
	}
}

func TestString(t *testing.T) {
	var s IntSet
	s.Add(1)
	s.Add(2)
	s.Add(3)
	if "{1 2 3}" != s.String() {
		t.Fatal(s.String())
	}
}
