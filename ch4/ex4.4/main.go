// Exercise 4.4: Write a version of rotate that operates in a single pass.
package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(a)
	rotate(a[:], 2)
	fmt.Println(a)
	rotate(a[:], 4)
	fmt.Println(a)
	rotate2(a[:], 2)
	fmt.Println(a)
	rotate2(a[:], 4)
	fmt.Println(a)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// rotate at position p
func rotate(s []int, p int) {
	reverse(s[:p])
	reverse(s[p:])
	reverse(s)
}
func rotate2(s []int, p int) {
	copy(s, append(s[p:], s[:p]...))
}

// goos: linux
// goarch: amd64
// pkg: gopl/ch4/ex4.4
// cpu: Intel(R) Core(TM) i5-7300U CPU @ 2.60GHz
// BenchmarkRotate-4    	93279985	        11.50 ns/op
// BenchmarkRotate2-4   	26106001	        42.14 ns/op
// PASS
// ok  	gopl/ch4/ex4.4	2.237s
