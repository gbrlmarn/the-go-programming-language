// Exercise 2.1: Add types, constants, and functions to tempconv for processing temperatures in the Kelvin scale, where zero Kelvin is -273.15°C and a difference
package tempconv

import "fmt"

type Celisius   float64
type Fahrenheit float64
type Kelvin     float64

const (
    AbosoluteZeroC Celisius = -273.15
    FreezingC      Celisius = 0
    BoilingC       Celisius = 100
)

func (c Celisius) String() string { return fmt.Sprintf("%g°C", c) }
func (c Fahrenheit) String() string { return fmt.Sprintf("%g°F", c) }
func (c Kelvin) String() string { return fmt.Sprintf("%g°K", c) }
