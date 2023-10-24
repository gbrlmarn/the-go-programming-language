// Exercise 7.4: The strings.NewReader function returns a value that satisfies the io.Reader interface (and others) by reading its argument, a string. Implement a simple version of NewReader yourself, and use it to make the HTML parser (§5.2) take input from a string.
package main

import (
	"io"
)

type Reader struct {
	s string
	i int64
}

// Read implements the io.Reader interface.
func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

func NewReader(s string) *Reader {
	return &Reader{s, 0}
}
