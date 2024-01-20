// Exercise 11.6: Write benchmarks to comapare the PopCount implementation in Section 2.6.2 with your solution to Exercise 2.4 and Exercise 2.5. At what point does the table-based approach break even?
package main

import (
	ex24 "gopl/ch2/ex2.4"
	ex25 "gopl/ch2/ex2.5"
	"gopl/ch2/examples/popcount"
	"testing"
)

func benchmark(b *testing.B, n int, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			f(uint64(j))
		}
	}
}

func BenchmarkPopCountTable1(b *testing.B)      { benchmark(b, 1, popcount.PopCount) }
func BenchmarkPopCountTable10(b *testing.B)     { benchmark(b, 10, popcount.PopCount) }
func BenchmarkPopCountTable100(b *testing.B)    { benchmark(b, 100, popcount.PopCount) }
func BenchmarkPopCountTable1000(b *testing.B)   { benchmark(b, 1000, popcount.PopCount) }
func BenchmarkPopCountTable10000(b *testing.B)  { benchmark(b, 10000, popcount.PopCount) }
func BenchmarkPopCountTable100000(b *testing.B) { benchmark(b, 100000, popcount.PopCount) }

func BenchmarkPopCountShift1(b *testing.B)      { benchmark(b, 1, ex24.PopCount) }
func BenchmarkPopCountShift10(b *testing.B)     { benchmark(b, 10, ex24.PopCount) }
func BenchmarkPopCountShift100(b *testing.B)    { benchmark(b, 100, ex24.PopCount) }
func BenchmarkPopCountShift1000(b *testing.B)   { benchmark(b, 1000, ex24.PopCount) }
func BenchmarkPopCountShift10000(b *testing.B)  { benchmark(b, 10000, ex24.PopCount) }
func BenchmarkPopCountShift100000(b *testing.B) { benchmark(b, 100000, ex24.PopCount) }

func BenchmarkPopCountClear1(b *testing.B)      { benchmark(b, 1, ex25.PopCount) }
func BenchmarkPopCountClear10(b *testing.B)     { benchmark(b, 10, ex25.PopCount) }
func BenchmarkPopCountClear100(b *testing.B)    { benchmark(b, 100, ex25.PopCount) }
func BenchmarkPopCountClear1000(b *testing.B)   { benchmark(b, 1000, ex25.PopCount) }
func BenchmarkPopCountClear10000(b *testing.B)  { benchmark(b, 10000, ex25.PopCount) }
func BenchmarkPopCountClear100000(b *testing.B) { benchmark(b, 100000, ex25.PopCount) }

// goos: linux
// goarch: amd64
// pkg: gopl/ch11/ex11.6
// cpu: Intel(R) Core(TM) i5-9400F CPU @ 2.90GHz
// BenchmarkPopCountTable1-6        	307830816	         3.896 ns/op
// BenchmarkPopCountTable10-6       	34422615	        34.31 ns/op
// BenchmarkPopCountTable100-6      	 3465213	       345.6 ns/op
// BenchmarkPopCountTable1000-6     	  355166	      3390 ns/op
// BenchmarkPopCountTable10000-6    	   35553	     33847 ns/op
// BenchmarkPopCountTable100000-6   	    3559	    337810 ns/op
// BenchmarkPopCountShift1-6        	27355171	        43.67 ns/op
// BenchmarkPopCountShift10-6       	 2743077	       442.5 ns/op
// BenchmarkPopCountShift100-6      	  278389	      4313 ns/op
// BenchmarkPopCountShift1000-6     	   27835	     43163 ns/op
// BenchmarkPopCountShift10000-6    	    2786	    431733 ns/op
// BenchmarkPopCountShift100000-6   	     277	   4311607 ns/op
// BenchmarkPopCountClear1-6        	417568454	         2.873 ns/op
// BenchmarkPopCountClear10-6       	38471385	        31.17 ns/op
// BenchmarkPopCountClear100-6      	 2628230	       454.9 ns/op
// BenchmarkPopCountClear1000-6     	  271195	      4414 ns/op
// BenchmarkPopCountClear10000-6    	   25308	     47395 ns/op
// BenchmarkPopCountClear100000-6   	    2277	    533602 ns/op
// PASS
// ok  	gopl/ch11/ex11.6	25.619s
