package format_test

import (
	"fmt"
	"gopl/ch12/examples/format"
	"testing"
	"time"
)

func Test(t *testing.T) {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(format.Any(x))
	fmt.Println(format.Any(d))
	fmt.Println(format.Any([]int64{x}))
	fmt.Println(format.Any([]time.Duration{d}))
}

// 1
// 1
// []int64 0xc000012180
// []time.Duration 0xc000012188
// PASS
// ok  	gopl/ch12/examples/format	0.002s
