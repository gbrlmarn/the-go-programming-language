package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestCharCount(t *testing.T) {
	tests := []struct {
		in      string
		counts  map[rune]int
		utflen  []int
		invalid int
	}{
		{
			in: "Hello mister",
			counts: map[rune]int{
				'H': 1,
				'e': 2,
				'l': 2,
				'o': 1,
                ' ': 1,
				'm': 1,
				'i': 1,
				's': 1,
				't': 1,
				'r': 1,
			},
            utflen: []int{0, 12, 0, 0, 0},
            invalid: 0,
		},
	}
	for _, test := range tests {
        counts, utflen, invalid := charcount(strings.NewReader(test.in))
        if !reflect.DeepEqual(counts, test.counts) {
            t.Errorf("charcount(%s) = %v", test.in, counts)
        }
        if !reflect.DeepEqual(utflen, test.utflen) {
            t.Errorf("want:%v\ngot:%v", test.utflen, utflen)
        }
        if invalid != test.invalid {
            t.Errorf("want:%d\ngot:%d", test.utflen, utflen)
        }
	}
}
