// Exercise 2.1: Add types, constants, and functions to tempconv for processing temperatures in the Kelvin scale, where zero Kelvin is -273.15°C and a difference
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbosoluteZeroC Celsius = -273.15
	FreezingC      Celsius = 0
	BoilingC       Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (c Fahrenheit) String() string { return fmt.Sprintf("%g°F", c) }
func (c Kelvin) String() string     { return fmt.Sprintf("%g°K", c) }
