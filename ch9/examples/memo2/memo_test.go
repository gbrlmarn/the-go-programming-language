package memo_test

import (
	"testing"

	"gopl.io/ch9/memo2"
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
// https://godoc.org, 490.768875ms, 32378 bytes
// https://golang.org, 1.106616667s, 61870 bytes
// http://gopl.io, 3.692236333s, 4154 bytes
// https://play.golang.org, 4.615697125s, 30053 bytes
// https://godoc.org, 4.615665875s, 32378 bytes
// https://play.golang.org, 4.6155315s, 30053 bytes
// http://gopl.io, 4.6155455s, 4154 bytes
// https://golang.org, 4.6156815s, 61870 bytes
// --- PASS: TestConcurrent (4.62s)
// PASS
// ok  	command-line-arguments	5.972s
