package main

import (
	"testing"
)

func TestNewRead(t *testing.T) {
	s := "hello mister"
	r := NewReader(s)
	
	buf := make([]byte, 6) // for 'hello '
	n, err := r.Read(buf)
	if n != 6 {
		t.Fatal(r, err, n)
	}
	
	buf = make([]byte, 6) // just 'mister'
	n, err = r.Read(buf)
	if n != 6 {
		t.Fatal(r, err, n)
	}
}
