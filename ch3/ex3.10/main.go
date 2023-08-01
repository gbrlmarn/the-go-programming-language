// Exercie 3.10: Write a non-recursive version of comma, using bytes.Buffer instead of string concatenation
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
    pre := len(s) % 3 

    if pre == 0 {
        pre = 3
    }
    b.WriteString(s[0:pre])
    for i := pre; i < len(s); i += 3 {
        b.WriteByte(',')
        b.WriteString(s[i : i+3])
    }
    return b.String()
}
