// Exercise 2.5: The expression x&(x-1) clears the rightmost non-zero bit of x. Write a version of PopCount that counts bits by using this fact, and assess its performance.
package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1)
		n++
	}
	return n
}

// goos: linux
// goarch: amd64
// pkg: gopl/ch2/ex2.5
// cpu: Intel(R) Core(TM) i5-7300U CPU @ 2.60GHz
// BenchmarkPopCount-4   	60004566	        20.07 ns/op
// PASS
// ok  	gopl/ch2/ex2.5	1.228s
