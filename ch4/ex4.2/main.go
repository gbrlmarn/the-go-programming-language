// Exercise 4.2: Write a program that prints the SHA256 hash of its standard input by default but supports a command-line flag to print the SHA384 or SHA512 hash instead
package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	s512 := flag.Bool("s512", false, "use sha512 alg")
	flag.Parse()
	r := bufio.NewReader(os.Stdin)
	in, err := r.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "ReadString: %v\n", err)
	}
	if *s512 {
		fmt.Printf("%x\n", sha512.Sum512([]byte(in)))
	} else {
		fmt.Printf("%x\n", sha256.Sum256([]byte(in)))
		fmt.Printf("Use -s512 flag for better encryption\n")
	}
}
