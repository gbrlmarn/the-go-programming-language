// Exercise 5.15: Write variadic functions max and min, analogous to sum. What should these functions do when called with no arguments? Write variants that require at least one argument.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(min())
	fmt.Println(max())
	fmt.Println(min(1))
	fmt.Println(max(1))
	fmt.Println(min(1, 2, 3))
	fmt.Println(max(1, 2, 3))

	fmt.Println()

	// fmt.Println(min2()) // will not compile
	// fmt.Println(max2()) // will not compile
	fmt.Println(min2(1))
	fmt.Println(max2(1))
	fmt.Println(min2(1, 2, 3))
	fmt.Println(max2(1, 2, 3))
}

func min(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	min := vals[0]
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func max(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

// Variants that require at least one arg
func min2(first int, others ...int) int {
	min := first
	for _, val := range others {
		if val < min {
			min = val
		}
	}
	return min
}

func max2(first int, others ...int) int {
	max := first
	for _, val := range others {
		if val > max {
			max = val
		}
	}
	return max
}
