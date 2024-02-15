// Modify dup2 to print the names of all files in which each duplicated line occurs.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex1.4:%v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for filename, lines := range counts {
		for line, n := range lines {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", filename, n, line)
			}
		}
	}
}

func countLines(f *os.File, count map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if count[f.Name()] == nil {
			count[f.Name()] = make(map[string]int)
		}
		count[f.Name()][input.Text()]++
	}
}
