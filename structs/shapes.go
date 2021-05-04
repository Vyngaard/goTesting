package main

import "math"

//this shape interface is used in our testing calls to refer to each of our shapes
type Shape interface {
	Area() float64
}

//we use distinct structs to define each shape with its measurements
type Rectangle struct {
	Width  float64
	Height float64
}

//then each struct has an Area method that we can then use to determine the area of the shape
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

/* func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
} */
