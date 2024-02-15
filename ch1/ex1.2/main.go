// Exercise 1.2: Modify the echo program to print the index and value of each of its arguments one per line

package main

import (
	"fmt"
	"os"
	"strings"
)

func echo1() {
	fmt.Println("echo1:")
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%d: %s\n", i, os.Args[i])
	}
}

func echo2() {
	fmt.Println("echo2:")
	for idx, arg := range os.Args[1:] {
		fmt.Printf("%d: %s\n", idx, arg)
	}
}

func main() {
	echo1()
	echo2()
}
