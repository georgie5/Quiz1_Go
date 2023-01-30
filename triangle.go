package main

import (
	"math"
)

// parts of a triangle
type Triangle struct {
	base   float64
	height float64
}

// Area of a triangle
func (t Triangle) Area() float64 {
	area := 0.5 * t.base * t.height // formula to find the area of a triangle
	return area
}

// Parimiter of a triangle
func (t Triangle) Parimeter() float64 {
	//In order to find the parimeter of the triangle the hypotenuse must be found first
	hypotenuse := math.Sqrt(math.Pow(t.base, 2) + math.Pow(t.height, 2))
	//formula to find parimeter
	parimeter := t.base + t.height + hypotenuse
	return parimeter
}
