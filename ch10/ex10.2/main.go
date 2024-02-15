// Exercise 10.2: Define a generic archive file-reading function capable of reading ZIP files(archive/zip) and POSIX tar files (archive/tar). Use a registration mechanism similar to the one described above so that support for each file format can be plugged in using blank imports.
package main

import (
	"log"
	"os"

	"gopl/ch10/ex10.2/arch"
	_ "gopl/ch10/ex10.2/arch/tar"
	_ "gopl/ch10/ex10.2/arch/zip"
)

func main() {
	err := arch.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}
