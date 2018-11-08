package area

import (
	"math"
)

// Shape interface defines shape behaviour
type Shape interface {
	Area() float64
}

// Rectangle rdefines a rectangle
type Rectangle struct {
	Width, Height float64
}

// Area calculates area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Circle defines a circle
type Circle struct {
	Radius float64
}

// Area calculates area of a rectangle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle defines a triangle
type Triangle struct {
	Base, Altitude float64
}

// Area calculates area of a rectangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Altitude
}
