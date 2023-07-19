// Exercise 1.1: Modify the echo program to also print os.Args[0], the name of the command that invoked it.
package main

import (
	"fmt"
	"os"
	"strings"
)

func echo1() {
  var s, sep string;
  for i := 0; i < len(os.Args); i++ {
    s += sep + os.Args[i]
    sep = " "
  }
  fmt.Print("echo1 output:")
  fmt.Println(s);
}

func echo2() {
  s, sep := "", ""
  for _, arg := range os.Args[0:] {
    s += sep + arg    
    sep = " "
  }
  fmt.Print("echo2 output:")
  fmt.Println(s)
}

func echo3() {
  fmt.Print("echo3 output:")
  fmt.Println(strings.Join(os.Args[0:], " "))
}

func main() {
  echo1();
  echo2();
  echo3();
} 
