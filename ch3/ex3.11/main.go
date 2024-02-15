// Exercise 3.11: Enhance comma so that it deals correctly with floating-point numbers and an optional sign.
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Println(comma(os.Args[1]))
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var b bytes.Buffer
	if s[0] == '+' || s[0] == '-' {
		b.WriteByte(s[0])
		s = s[1:]
	}
	fpidx, afterfp := afterfp(s)
	s = s[:fpidx]

	pre := len(s) % 3
	if pre == 0 {
		pre = 3
	}
	b.WriteString(s[0:pre])
	for i := pre; i < len(s); i += 3 {
		b.WriteByte(',')
		b.WriteString(s[i : i+3])
	}
	b.WriteString(afterfp)
	return b.String()
}

func afterfp(s string) (int, string) {
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			return i, s[i:]
		}
	}
	return len(s), ""
}
