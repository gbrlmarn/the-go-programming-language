package main

import (
	"fmt"
	"testing"
)

func BenchmarkComma(b *testing.B) {
	for i := 0; i < 99999; i++ {
		comma(fmt.Sprint(i))
	}
}
