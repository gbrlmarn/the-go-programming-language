// Copyright © 2023 Gabriel Marin

package tempconv

import (
	"fmt"
	"testing"
)

func TestConv(t *testing.T) {
	fmt.Printf("%g°C to %g°K\n", BoilingC, CToK(BoilingC))
	fmt.Printf("%g°K to %g°C\n", CToK(BoilingC), BoilingC)
	fmt.Printf("%g°F to %g°K\n", CToF(BoilingC), CToK(BoilingC))
	fmt.Printf("%g°K to %g°F\n", CToK(BoilingC), CToF(BoilingC))
}
