// Bzipper reads input, bzip2-compresses it, and writes it out.
package main

import (
	"gopl/ch13/ex13.4/bzip"
	"io"
	"log"
	"os"
)

func main() {
	w, err := bzip.NewWriter(os.Stdout)
	if err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("bzipper: close: %v\n", err)
	}
}
