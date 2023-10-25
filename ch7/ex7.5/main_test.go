package main

import (
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	s := "hello mister"
	l := LimitReader(strings.NewReader(s), 5)

	buf := make([]byte, 6)
	n, err := l.Read(buf)
	if n != 5 {
		t.Fatal(n, err)
	}
}
