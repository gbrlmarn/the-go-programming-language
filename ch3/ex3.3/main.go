// Exercise 3.3: Color each polygon based on its height, so that the peaks are colored red (#ff0000) and the valleys blue (#0000ff)
package main

import (
    "fmt"
    "math"
)

const (
    width, height = 600, 320            // canvas size in pixels
    cells         = 100                 // number of grid cells
    xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
    xyscale       = width / 2 / xyrange // pixels per x or y unit
    zscale        = height * 0.4        // pixels per z unit
    angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
    fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
        "style='stroke: grey; fill: white; stroke-width: 0.7' "+
        "width='%d' height='%d'>", width, height)
    var clr string
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            ax, ay, z := corner(i+1, j)
            bx, by, _ := corner(i, j)
            cx, cy, _ := corner(i, j+1)
            dx, dy, _ := corner(i+1, j+1)
            corners := []float64{
                ax, ay, bx, by,
                cx, cy, dx, dy,
            }
            if(isFinite(corners)) {
                if z < 0 {
                    clr = "blue"
                } else {
                    clr = "red"
                }
                fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s;'/>\n",
                    ax, ay, bx, by, cx, cy, dx, dy, clr)
            }
        }
    }
    fmt.Println("</svg>")
}
func isFinite(in []float64) bool {
    for _, v := range in {
        if math.IsInf(v, 0) || math.IsNaN(v) {
            return false
        }
    }
    return true
}
func corner(i, j int) (float64, float64, float64) {
    // Find point (x,y) at corner cell (i,j).
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)

    // Compute surface height z
    z := f(x,y)
    
    // Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
    sx := width/2 + (x-y)*cos30*xyscale
    sy := height/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy, z
}

func f(x, y float64) float64 {
    r := math.Hypot(x, y) // distance from (0,0)
    return math.Sin(r) / r
}
