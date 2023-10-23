package main

import "testing"

func TestWordCounter(t *testing.T) {
	var c WordCounter
	c.Write([]byte("hello mister"))
	if c != 2 {
		t.Fatal(c)
	}
	c.Write([]byte("hello again"))
	if c != 4 {
		t.Fatal(c)
	}
}

func TestLineCounter(t *testing.T) {
	var c LineCounter
	c.Write([]byte("First line"))
	if c != 1 {
		t.Fatal(c)
	}
	c.Write([]byte("Second line"))
	if c != 2 {
		t.Fatal(c)
	}
}

