// Exercise 4.1: Write a function that counts the number of bits that are different in two SHA256 hashes. (See PopCount from Section 2.6.2.)
package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	var c1, c2 [32]byte
	if len(os.Args) == 3 {
		c1 = sha256.Sum256([]byte(os.Args[1]))
		c2 = sha256.Sum256([]byte(os.Args[2]))
	} else {
		c1 = sha256.Sum256([]byte("x"))
		c2 = sha256.Sum256([]byte("X"))
	}
	fmt.Printf("Number of different bits: %d\n", bitDiff1(c1, c2))
	fmt.Printf("Number of different bits: %d\n", bitDiff2(c1, c2))
}

func bitDiff1(c1, c2 [sha256.Size]byte) int {
	var c int
	for i := 0; i < sha256.Size; i++ {
		for j := 0; j < 8; j++ {
			if ((c1[i] >> j) & 1) != ((c2[i] >> j) & 1) {
				c++
			}
		}
	}
	return c
}
func bitDiff2(c1, c2 [sha256.Size]byte) int {
	var c int
	for i := 0; i < sha256.Size; i++ {
		c += int(pc[c1[i]^c2[i]])
	}
	return c
}

// goos: darwin
// goarch: arm64
// pkg: gopl/ch4/ex4.1
// BenchmarkBitDif1-8   	1000000000	         0.2140 ns/op
// BenchmarkBitDif2-8   	1000000000	         0.1061 ns/op
// PASS
// ok  	gopl/ch4/ex4.1	4.669s
// The second one is better :D
