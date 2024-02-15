// Exercise 4.8: Modify charcount to count letters, digits, and so on in their Unicode categories, using functions like unicode.IsLetter
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // counts of lengths of UTF-8 encodings
	utfcat := make(map[string]int)  // counts of categories of UTF-8 encodings
	invalid := 0                    // count invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, err
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex4.8: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		for name, table := range unicode.Categories {
			if unicode.Is(table, r) {
				utfcat[name]++
			}
		}
	}
	fmt.Print("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\ncat\tcount\n")
	for cat, n := range utfcat {
		fmt.Printf("%q\t%d\n", cat, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
