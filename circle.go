package main

import "math"

// contains 1 feild name , radius
type Circle struct {
	radius float64
}

// Area of a circle
func (c Circle) Area() float64 {
	//formula to find the area of a circle
	area := math.Pi * math.Pow(c.radius, 2)
	return area
}

// Perimeter of a circle
func (c Circle) Parimeter() float64 {
	//formula to find the perimeter of a circle
	perimeter := 2 * math.Pi * c.radius
	return perimeter
}
