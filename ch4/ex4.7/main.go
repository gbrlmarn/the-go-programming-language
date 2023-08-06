// Exercise 4.7: Modify reverse to revers the characters of a []byte slice that represents a UTF-8-encoded string, in place. Can you do it without allocating new memory?
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
    sb := []byte("hello mister")
    fmt.Println(string(sb))
    //reverse(sb)
    reverse2(sb)
    fmt.Println(string(sb))
}

// with new memory
func reverse(sb []byte) {
    buf := make([]byte, 0, len(sb)) 
    for i := len(sb); i > 0; {
        _, s := utf8.DecodeLastRune(sb[:i])
        buf = append(buf, sb[i-s:i]...)
        i -= s
    }
    copy(sb, buf)
}

// without new memory
// only works when utf-8 have the same byte length
// at least is faster :D
func reverse2(sb []byte) {
    for i, j := 0, len(sb); i < j; {
        _, si := utf8.DecodeRune(sb[i:])
        _, sj := utf8.DecodeLastRune(sb[:j])
        ci, cj := 0, sj // byte counters
        if si == sj {
            for ci != si && cj != 0 {
                sb[i+ci], sb[j-cj] = sb[j-cj], sb[i+ci]
                ci += 1
                cj -= 1
            }
        }
        i += si
        j -= sj
    }
}
// goos: darwin
// goarch: arm64
// pkg: the-go-programming-language/ch4/ex4.7
// BenchmarkReverse-8    	18488534	        64.89 ns/op
// BenchmarkReverse2-8   	47368342	        25.37 ns/op
// PASS
// ok  	the-go-programming-language/ch4/ex4.7	3.566s
