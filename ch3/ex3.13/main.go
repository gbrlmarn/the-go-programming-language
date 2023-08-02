// Exercise 3.13: Write const declarations for KB, MB, up through YB as compactly as you can.
package main

import (
	"fmt"
	"strconv"
)

const (
    B = 1 << (10 * iota) 
    KiB
    MiB
    GiB
    TiB
    PiB
    EiB
    ZiB
    YiB
)

func main() {
    fmt.Println(KiB)
    fmt.Println(MiB)
    fmt.Println(GiB)
    fmt.Println(TiB)
    fmt.Println(PiB)
    fmt.Println(EiB)
    fmt.Println(float64(ZiB))
    fmt.Println(float64(YiB))
}
