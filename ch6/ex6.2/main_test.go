package main

import "testing"

func TestAddAll(t *testing.T) {
	s := &IntSet{}
	s.AddAll(1, 2, 3)
	if !s.Has(1) || !s.Has(2) || !s.Has(3) {
		t.Fatal(s)
		t.Fail()
	}
}
