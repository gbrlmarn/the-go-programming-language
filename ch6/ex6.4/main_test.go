package main

import
	"testing"

func TestElems(t *testing.T) {
	s := &IntSet{}
	s.Add(1)
	s.Add(2)
	s.Add(3)
	sElems := s.Elems()
	for i, el := range sElems {
		if [3]int{1, 2, 3}[i] != el {
			t.Fatal(sElems)
		}
	}
}
