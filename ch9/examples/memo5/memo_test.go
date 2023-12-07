package memo_test

import (
	"testing"

	"gopl.io/ch9/memo5"
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
// https://godoc.org, 663.045458ms, 32378 bytes
// https://godoc.org, 663.075667ms, 32378 bytes
// https://golang.org, 764.354333ms, 61870 bytes
// https://golang.org, 764.338666ms, 61870 bytes
// https://play.golang.org, 917.491916ms, 30053 bytes
// https://play.golang.org, 917.622625ms, 30053 bytes
// http://gopl.io, 1.315222375s, 4154 bytes
// http://gopl.io, 1.315345083s, 4154 bytes
// --- PASS: TestConcurrent (1.32s)
// PASS
// ok  	command-line-arguments	2.709s
