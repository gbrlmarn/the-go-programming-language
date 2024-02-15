// Exercise 2.3: Rewrite PopCount to use a loop instead of a single expnsion. Compare the performance of the two versions. (Section 11.4 shows how to comapare the performance of different implementations systematically)
package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	n := 0
	for i := 0; i < 8; i++ {
		n += int(pc[byte(x>>(i*8))])
	}
	return n
}
