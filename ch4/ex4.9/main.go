// Exercise 4.9: Write a program wordfreq to report the frequency of each word in an input text file. Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into words instead of lines
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    wfreq := make(map[string]int) // words freq
    input := bufio.NewScanner(os.Stdin)
    input.Split(bufio.ScanWords) 
    for input.Scan() {
        word := input.Text() 
        wfreq[word]++
    }
    fmt.Print("word\tfreq\n")
    for w, f := range wfreq {
        fmt.Printf("%s\t%d\n", w, f)
    }
}
