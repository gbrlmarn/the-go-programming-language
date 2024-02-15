// reverse reverses a slice of ints in place.
package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	reverse(arr[:])
	fmt.Println(arr)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
