// Exercise 9.6: Measure how the performance of a computer-bound parallel program(see Exercise 8.5) varies with GOMAXPROCS. What is the optimal value on your computer? How many CPUs does your computer have?
package main

import (
	"fmt"
	"image"
	"image/color"
	"math/cmplx"
	"runtime"
	"sync"
	"time"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	for i := 1; i <= runtime.NumCPU(); i++ {
		cdraw(i)
	}
}

func cdraw(nprocs int) {
	start := time.Now()
	workers := runtime.GOMAXPROCS(nprocs)
	var wg sync.WaitGroup // number of working goroutines
    img := image.NewRGBA(image.Rect(0, 0, width, height))
	for w := 0; w < workers; w++ {
		wg.Add(1)
		// worker
		go func() {
			defer wg.Done()
			for py := 0; py < height; py++ {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					// Image point (px, py) represents complex value z.
					img.Set(px, py, mandlebrot(z))
				}
			}
		}()
	}
	
	// closer
	go func() {
		wg.Wait()
	}()

	//f, _ := os.Create("mandlebrot.png")
	//png.Encode(f, img) // NOTE: ignoring errors
	//fmt.Printf("encoded in: %v\n", time.Since(start))
	fmt.Printf("encoded in %v\t with %v CPUs\n", time.Since(start), nprocs)
}

func mandlebrot(z complex128) color.Color {
    const iterations = 200
    const contrast = 15

    var v complex128
    for n := uint8(0); n < iterations; n++ {
        v = v*v + z
        if cmplx.Abs(v) > 2 {
            return color.Gray{255 - contrast*n}
       }
    }
    return color.Black
}


// go run main.go 
// encoded in 135.916µs	 with 1 CPUs
// encoded in 111.583µs	 with 2 CPUs
// encoded in 276.708µs	 with 3 CPUs
// encoded in 134.875µs	 with 4 CPUs
// encoded in 66.75µs	 with 5 CPUs
// encoded in 497.958µs	 with 6 CPUs
// encoded in 208.416µs	 with 7 CPUs
// encoded in 93.459µs	 with 8 CPUs
