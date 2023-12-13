package main

import "testing"

func BenchmarkPipe(b *testing.B) {
	in, out := pipe(100000)
	for i := 0; i < b.N; i++ {
		in <- 1
		<-out
	}
	close(in)
}

// go test -bench=.
// goos: linux
// goarch: amd64
// pkg: the-go-programming-language/ch9/ex9.4
// cpu: Intel(R) Core(TM) i5-7300U CPU @ 2.60GHz
// BenchmarkPipe-4   	      32	  34631122 ns/op
// PASS
// ok  	the-go-programming-language/ch9/ex9.4	2.710s
