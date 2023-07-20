// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

type Celisius float64
type Fahrenheit float64

const (
    AbosoluteZeroC Celisius = -273.15
    FreezingC      Celisius = 0
    BoilingC       Celisius = 100
)

func (c Celisius) String() string { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
