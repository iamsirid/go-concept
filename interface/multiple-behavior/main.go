package main

import (
	"fmt"
	"math"
)

type Sizer interface {
	Area() float64
}

// Shaper interface embeds Sizer and Stringer interfaces
type Shaper interface {
	Sizer
	fmt.Stringer
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c *Circle) String() string {
	return fmt.Sprintf("Circle with radius %.2f", c.Radius)
}

type Square struct {
	Width  float64
	Height float64
}

func (s *Square) Area() float64 {
	return s.Width * s.Height
}

func (s *Square) String() string {
	return fmt.Sprintf("Square with width %.2f and height %.2f", s.Width, s.Height)
}

func main() {
	c := &Circle{Radius: 5}
	PrintArea(c)
	s := &Square{Width: 5, Height: 5}
	PrintArea(s)

	l := Less(c, s)
	fmt.Printf("%+v is the smallest\n", l)

}

func Less(s1, s2 Sizer) Sizer {
	if s1.Area() < s2.Area() {
		return s1
	}
	return s2
}

func PrintArea(s Shaper) {
	fmt.Printf("Area of %s is %.2f\n", s.String(), s.Area())
}
