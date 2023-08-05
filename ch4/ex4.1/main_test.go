package main

import (
	"crypto/sha256"
	"math/rand"
	"testing"
	"time"
)

const (
    lstr = 10
    nstrs = 1000000
)

var tstr1 [nstrs]string
var tstr2 [nstrs]string

func init() {
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < len(tstr1); i++ {
        tstr1[i] = randS(lstr)
    }
    for i := 0; i < len(tstr2); i++ {
        tstr2[i] = randS(lstr)
    }
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randS(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func BenchmarkBitDif1(b *testing.B) {
    for i := 0; i < nstrs; i++ {
        c1 := sha256.Sum256([]byte(tstr1[i]))
        c2 := sha256.Sum256([]byte(tstr2[i]))
        bitDiff1(c1, c2)
    }    
}

func BenchmarkBitDif2(b *testing.B) {
    for i := 0; i < nstrs; i++ {
        c1 := sha256.Sum256([]byte(tstr1[i]))
        c2 := sha256.Sum256([]byte(tstr2[i]))
        bitDiff2(c1, c2)
    }    
}
