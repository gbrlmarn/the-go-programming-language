// Exercise 7.2: Write a function CountingWriter with the signature below that, given an io.Writer, returns a new Writer that wraps the original, and a pointer to an int64 variable that at any moment contains the number of bytes written to the new Writer.
// func CountingWriter(w io.Writer) (io.Writer, *int64)
package main

import "io"

type ByteCounterWriter struct {
	w io.Writer
	c int64
}

func (cw *ByteCounterWriter) Write(p []byte) (int, error) {
	n, err := cw.w.Write(p)
	cw.c += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &ByteCounterWriter{w, 0}
	return cw, &cw.c
}
