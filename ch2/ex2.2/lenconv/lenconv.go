// Package lenconv performs feet and meters length computations.
package lenconv

import "fmt"

type Meter float64
type Foot float64

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (f Foot) String() string { return fmt.Sprintf("%g🦶", f) }  
