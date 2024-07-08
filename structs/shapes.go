package main

import "math"

// Shape is implemented by anything that can tell us its Area.
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

// Area returns the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of a rectangle.
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

type Circle struct {
	Radius float64
}

// The key difference this is a function
// func Area(circle Circle) float64 {
// 	return math.Pi * circle.Radius * circle.Radius
// }

// This is how Go defines a method on Circle... defines it as a reciever
func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

// Triangle represents the dimensions of a triangle.
type Triangle struct {
	Base   float64
	Height float64
}

// Area returns the area of the triangle.
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
