// Exercise 3.12: Write a function that reports whether two string are anagrams of each other, that is, they contain the same letters in a different order.
package main

import (
	"fmt"
	"os"
)

func main() {
    if len(os.Args) == 2 {
        fmt.Println(anagram(os.Args[1], os.Args[2]))
    } else {
        fmt.Println("Please enter 2 strings as args")
    }
}

func anagram(s1, s2 string) bool {
    m := make(map[rune]int)
    for _, v := range s1 {
        m[v]++
    }
    for _, v := range s2 {
        m[v]--
    }
    for _, v := range s1 {
        if m[v] != 0 {
            return false
        }
    }
    return true
}
