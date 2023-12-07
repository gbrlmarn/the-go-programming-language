package memo_test

import (
	"testing"

	"gopl.io/ch9/memo3"
	"gopl.io/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func start(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}

// go test -run=TestConcurrent -race -v memo_test.go
// <TestConcurrent -race -v memo_test.go
// === RUN   TestConcurrent
// https://godoc.org, 674.369291ms, 32378 bytes
// https://godoc.org, 704.686208ms, 32378 bytes
// https://golang.org, 739.879083ms, 61870 bytes
// https://golang.org, 761.563333ms, 61870 bytes
// https://play.golang.org, 983.6735ms, 30053 bytes
// https://play.golang.org, 983.6385ms, 30053 bytes
// http://gopl.io, 1.241421667s, 4154 bytes
// http://gopl.io, 1.242644375s, 4154 bytes
// --- PASS: TestConcurrent (1.24s)
// PASS
// ok  	command-line-arguments	2.674s
