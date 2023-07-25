// Exercise 3.2: Experiment with visualizations of other funtions form the math package. Can you produce an egg box, moguls, or a saddle?
package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

const (
    width, height = 600, 400            // canvas size
    cells         = 100                 // number of grid cells
    xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
    xyscale       = width / 2 / xyrange // pixels per x or y unit
    zscale        = height * 0.4        // pixels per z unit
    angle         = math.Pi / 6         // angle of x, y axes (=30°)
    eggscale      = 0.2                 // scaler for egg function
    acurv, bcurv  = 15, 30              // saddle curvature levels
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type hFct func(x, y float64) float64

func main() {
    help := "ex3.2 egg|saddle\n"
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, help)
        os.Exit(1)
    }
    var f hFct; 
    if os.Args[1] == "egg" {
        f = egg
    } else if os.Args[1] == "saddle" {
        f = saddle
    } else {
        fmt.Fprintf(os.Stderr, help)
        os.Exit(1)
    }
    svg(os.Stdout, f)
}

func svg(w io.Writer, f hFct) {
    fmt.Fprintf(w, "<svg xmlns='http://w3.org/2000/svg' "+
        "style='stroke: grey; fill: white; sroke-width: 0.7' "+
    "width='%d' heigth='%d'>", width, height)
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            ax, ay := corner(i+1, j, f)
            bx, by := corner(i, j, f)
            cx, cy := corner(i, j+1, f)
            dx, dy := corner(i+1, j+1, f)
            fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy)
        }
    }
}

func corner(i, j int, f hFct) (float64, float64) {
    // Find point (x,y) at corner of cell (i,j).
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)

    // Compute surface heigh z
    z := f(x, y)

    // Project(x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
    sx := width/2 + (x-y)*cos30*xyscale
    sy := height/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy
}
func saddle(x, y float64) float64 {
    return (math.Pow(y,2) / math.Pow(bcurv,2) - math.Pow(x,2) / math.Pow(acurv,2))
}
func egg(x, y float64) float64 {
    return eggscale*(math.Sin(x)*math.Sin(y))
}
