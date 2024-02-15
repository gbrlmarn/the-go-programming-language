// basename removes directory components and a .sufix.
// e.g. a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		fmt.Println(basename(os.Args[1]))
	}
}
func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
