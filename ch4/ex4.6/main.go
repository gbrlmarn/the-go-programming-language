// Exercise 4.6: Write an in-place function that squashes each run of adjacent Unicode spaces (see unicode.IsSpace) in a UTF-8 encoded []byte slice into a single ASCII space.
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	bs := []byte("abc\r\t  \n\rdef")
	fmt.Printf("%q\n", bs)
	fmt.Printf("%q\n", string(nodupsp(bs)))
}

func nodupsp(bs []byte) []byte {
	out := bs[:0]
	add := false
	for i := 0; i < len(bs); {
		r, s := utf8.DecodeRune(bs[i:])
		if unicode.IsSpace(r) && add {
			out = append(out, ' ')
			add = false
		} else if unicode.IsSpace(r) {
			add = false
		} else {
			out = append(out, bs[i:i+s]...)
			add = true
		}
		i += s
	}
	return out
}
