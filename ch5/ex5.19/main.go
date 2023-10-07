// Exercise 5.19: Use panic and recover to write a function that contains no return statement yet returns a non-zero value.
package main

import "fmt"

func main() {
	fmt.Println(retNonZero())
}

func retNonZero() (res int) {
	defer func() {
		res = 3 // lucky number
		recover()
	}()
	panic("We are doomed...")
}
