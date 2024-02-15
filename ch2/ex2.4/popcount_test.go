package popcount_test

import (
	"testing"

	"gopl/ch2/ex2.4"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x123456789ABCDEF)
	}
}

// pkg: gopl/ch2/ex2.4
// cpu: Intel(R) Core(TM) i5-7300U CPU @ 2.60GHz
// BenchmarkPopCount-4   	57588987	        20.48 ns/op
