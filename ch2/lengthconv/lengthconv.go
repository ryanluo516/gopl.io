// Exercise 2.2
package lengthconv

import "fmt"

type Feet float64
type Meters float64

func (f Feet) String() string   { return fmt.Sprintf("%g feet", f) }
func (m Meters) String() string { return fmt.Sprintf("%g meters", m) }

const (
	footInMeters Meters = 0.3048
	meterInFeet  Feet   = 3.28084
)

// FToM converts a feet length to meters.
func FToM(f Feet) Meters {
	return Meters(f) * footInMeters
}

// FToM converts a meters length to feet.
func MToF(m Meters) Feet {
	return Feet(m) * meterInFeet
}
