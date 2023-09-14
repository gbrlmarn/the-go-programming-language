// Exercise 5.9: Write a function expand(s string, f func(string) string) string that replaces each substring "$foo" withing s by the text returned by f("foo").
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(expand(input, caps))
}

func expand(s string, f func(string) string) string {
	var res []string
	words := strings.Split(s, " ")
	for _, w := range words {
		if strings.HasPrefix(w, "$") {
			res = append(res, f(w[1:]))
		} else {
			res = append(res, w)
		}
	}
	return strings.Join(res, " ")
}

func caps(s string) string {
	return strings.ToUpper(s)
}
