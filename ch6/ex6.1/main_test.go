package main

import "testing"

func initIntSet(len int) *IntSet {
	s := &IntSet{}
	for x := 0; x < len; x++ {
		s.Add(x)
	}
	return s
}

func TestLen(t *testing.T) {
	len := 100
	s := initIntSet(len)
	if got := s.Len(); got != len {
		t.Fatalf("got %d, want %d", got, len)
	}
}

func TestRemove(t *testing.T) {
	//s := initIntSet(1)
	s := &IntSet{}
	s.Remove(0)
	if s.Has(0) {
		t.Log(s)
		t.Fail()
	}
}

func TestClear(t *testing.T) {
	len := 100
	s := initIntSet(len)
	s.Clear()
	if s.Has(0) || s.Has(100) {
		t.Log(s)
		t.Fail()
	}
}

func TestCopy(t *testing.T) {
	len := 10
	s := initIntSet(len)
	c := s.Copy()
	c.Add(20)
	if !c.Has(20) || s.Has(20) {
		t.Log(s, c)
		t.Fail()
	}
}
