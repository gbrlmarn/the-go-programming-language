// Exercise 4.3: Rewrite reverse to use an array pointer instead of a slice
package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	reverse(&arr)
	fmt.Println(arr)
}

func reverse(sp *[5]int) {
	for i, j := 0, len(sp)-1; i < j; i, j = i+1, j-1 {
		sp[i], sp[j] = sp[j], sp[i]
	}
}
func reverse2(sp *[5]int) {
	for i, j := 0, len(sp)-1; i < j; i, j = i+1, j-1 {
		// C style deref...
		(*sp)[i], (*sp)[j] = (*sp)[j], (*sp)[i]
	}
}
