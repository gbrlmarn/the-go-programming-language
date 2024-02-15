// Exercise 7.1: Using the ideas from ByteCounter, implement counters for words and for lines. You will find bufio.ScanWords useful.
package main

import (
	"bufio"
	"strings"
)

type ByteCounter int
type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(strings.NewReader(string(p[:])))
	s.Split(bufio.ScanWords)
	count := 0
	for s.Scan() {
		count += 1
	}
	*c += WordCounter(count)
	return count, nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(strings.NewReader(string(p[:])))
	s.Split(bufio.ScanLines)
	count := 0
	for s.Scan() {
		count += 1
	}
	*c += LineCounter(count)
	return count, nil
}
