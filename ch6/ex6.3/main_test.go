package main

import "testing"

func TestIntersectWith(t *testing.T) {
	s1 := &IntSet{}
	s2 := &IntSet{}
	s1.AddAll(1, 2, 3)
	s2.AddAll(2, 3, 4)
	s1.IntersectWith(s2)
	if s1.Has(1) || s1.Has(4) {
		t.Fatal(s1)
		t.Fail()
	}
}

func TestDifferenceWith(t *testing.T) {
	s1 := &IntSet{}
	s2 := &IntSet{}
	s1.AddAll(1, 2, 3)
	s2.AddAll(2, 3, 4)
	s1.DifferenceWith(s2)
	if s1.Has(2) || s1.Has(3){
		t.Fatal(s1)
		t.Fail()
	}
}

func TestSymetricDifferenceWith(t *testing.T) {
	s1 := &IntSet{}
	s2 := &IntSet{}
	s1.AddAll(1, 2, 3)
	s2.AddAll(2, 3, 4)
	s1.SymetricDifference(s2)
	if s1.Has(2) || s1.Has(3) {
		t.Fatal(s1)
		t.Fail()
	}
}
