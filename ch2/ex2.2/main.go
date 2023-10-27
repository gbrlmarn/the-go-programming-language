// Exercise 2.2: Write a general-purpose unit-conversion analogous to cf that reads numbers from its command-line arguments or from the standard input if there are no arguments, and converts each number into units like temperature in Celsius and Fahrenheit, length in feet and meters, weight in pounds and kilograms, and the like.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"the-go-programming-language/ch2/ex2.1"
    "the-go-programming-language/ch2/ex2.2/wconv"
    "the-go-programming-language/ch2/ex2.2/lenconv"
)

func main() {
    var vals []string
    if len(os.Args) > 1 {
        vals = os.Args[1:]
    } else {
        vals = readStdio() 
    }
    for _, arg := range vals {
        v, err := strconv.ParseFloat(arg, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "ex2.2: %v\n", err)
            os.Exit(1)
        }
        printTemp(v)
        printLen(v)
        printWght(v)
    }
}

func printTemp(v float64) {
    c := tempconv.Celsius(v)
    f := tempconv.Fahrenheit(v)
    fmt.Printf("%s = %s, %s = %s\n",
        c, tempconv.CToF(c), f, tempconv.FToC(f))
}
func printLen(v float64) {
    m := lenconv.Meter(v)
    f := lenconv.Foot(v)
    fmt.Printf("%s = %s, %s = %s\n",
        m, lenconv.MToF(m), f, lenconv.FToM(f))
}
func printWght(v float64) {
    k := wconv.Kilogram(v)
    p := wconv.Pound(v)
    fmt.Printf("%s = %s, %s = %s\n",
        k, wconv.KToP(k), p, wconv.PToK(p))
}

func readStdio() []string {
    r := bufio.NewReader(os.Stdin)
    v, err := r.ReadString('\n')
    if err != nil {
        fmt.Fprintf(os.Stderr, "ex2.2: %v\n", err)
        os.Exit(1)
    }
    v = strings.Trim(v, "\n")
    vals := strings.Split(v, " ")
    return vals
}

