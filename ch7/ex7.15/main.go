// Exercise 7.15: Write a program that reads a single expression from the standard input, prompts the user to provide values for any variables, then evaluates the expression in the resulting environment. Handle all error gracefully.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"gopl/ch7/examples/eval"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	env := make(eval.Env)
	for {
		fmt.Print("=> ")
		s, _ := in.ReadString('\n')
		expr, err := eval.Parse(s)
		if err != nil {
			log.Fatal(err)
		}
		vars := make(map[eval.Var]bool)
		if err := expr.Check(vars); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v\n", expr.Eval(env))
	}
}
