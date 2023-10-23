package main

import (
	"io"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	w, c := CountingWriter(io.Discard)
	w.Write([]byte("hello"))
	if *c != 5 {
		t.Fatal(*c, w)
	}
	w.Write([]byte("again"))
	if *c != 10 {
		t.Fatal(*c, w)
	}
}
