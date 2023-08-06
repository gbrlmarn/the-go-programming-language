package main

import (
    "testing"
)

func BenchmarkReverse(b *testing.B) {
    sb := []byte("hello mister")
    for i := 0; i < b.N; i++ {
        reverse(sb)
    }
}

func BenchmarkReverse2(b *testing.B) {
    sb := []byte("hello mister")
    for i := 0; i < b.N; i++ {
        reverse2(sb)
    }
}
