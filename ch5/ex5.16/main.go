// Exercise 5.16: Write a variadic version of strings.Join.
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Join([]string{"Hello", "Mister"}, " "))
	fmt.Println(join(" ", "Hello", "Mister"))
}

func join(sep string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	res := ""
	for _, str := range strs {
		res += str
		res += sep
	}
	return res
}
