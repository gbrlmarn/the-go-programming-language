// Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice.
package main

import "fmt"

func main() {
	s := []string{
		"uno", "due",
		"due", "tre",
		"tre", "tre",
		"quatro", "quatro",
	}
	fmt.Printf("Vals: %v\tPtr: %p\n", s, &s[0])
	s = nodup(s)
	fmt.Printf("Vals: %v\tPtr: %p\n", s, &s[0])
}

// nodup removes adjacent duplicates
func nodup(s []string) []string {
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			s = append(s[:i-1], s[i:]...)
			i-- // try again
		}
	}
	return s
}
