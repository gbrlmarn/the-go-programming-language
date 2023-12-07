package memo_test

import (
	"testing"

	"gopl.io/ch9/memo1"
	"gopl.io/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func start(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

// NOTE: not concurrency-safe! Test fails
func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}

// go test -v memo_test.go
// === RUN   TestConcurrent
// https://godoc.org, 814.882208ms, 32378 bytes
// https://play.golang.org, 818.017875ms, 30053 bytes
// https://godoc.org, 821.933459ms, 32378 bytes
// https://golang.org, 848.435042ms, 61870 bytes
// https://golang.org, 848.04025ms, 61870 bytes
// https://play.golang.org, 1.011498625s, 30053 bytes
// http://gopl.io, 1.426413875s, 4154 bytes
// http://gopl.io, 1.435844792s, 4154 bytes
// --- PASS: TestConcurrent (1.44s)
// PASS
// ok  	command-line-arguments	1.794s




// go test -run=TestConcurrent -race -v memo_test.go
// <TestConcurrent -race -v memo_test.go
// === RUN   TestConcurrent
// https://godoc.org, 544.075625ms, 32378 bytes
// https://godoc.org, 551.878167ms, 32378 bytes
// https://golang.org, 772.815042ms, 61870 bytes
// https://play.golang.org, 823.586916ms, 30053 bytes
// https://golang.org, 832.936875ms, 61870 bytes
// https://play.golang.org, 868.540792ms, 30053 bytes
// http://gopl.io, 1.099846125s, 4154 bytes
// ==================
// WARNING: DATA RACE
// Write at 0x00c000126c60 by goroutine 12:
//   runtime.mapaccess2_faststr()
//       /usr/local/go/src/runtime/map_faststr.go:108 +0x42c
// ...
// ==================
// ==================
// WARNING: DATA RACE
// ...
//     testing.go:1465: race detected during execution of test
// FAIL
// FAIL	command-line-arguments	1.743s
// FAIL
