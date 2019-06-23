package colour

import "github.com/lukeshiner/raytrace/comparison"

// Colour holds colour data.
type Colour struct {
	Red, Green, Blue float64
}

// Equal returns true if c and other are equal, otherwise false.
func (c Colour) Equal(other Colour) bool {
	if comparison.EpsilonEqual(c.Red, other.Red) != true {
		return false
	}
	if comparison.EpsilonEqual(c.Green, other.Green) != true {
		return false
	}
	if comparison.EpsilonEqual(c.Blue, other.Blue) != true {
		return false
	}
	return true
}

// Add returns the addition of another colour with c.
func (c Colour) Add(other Colour) Colour {
	return Colour{c.Red + other.Red, c.Green + other.Green, c.Blue + other.Blue}
}

// Sub returns colour produced by subtracting other from c.
func (c Colour) Sub(other Colour) Colour {
	return Colour{c.Red - other.Red, c.Green - other.Green, c.Blue - other.Blue}
}

// ScalarMult returns the colour produced by multiplying c by scalar.
func (c Colour) ScalarMult(scalar float64) Colour {
	return Colour{c.Red * scalar, c.Green * scalar, c.Blue * scalar}
}

// Mult returns the multiplying c and other.
func (c Colour) Mult(other Colour) Colour {
	return Colour{c.Red * other.Red, c.Green * other.Green, c.Blue * other.Blue}
}
