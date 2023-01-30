package main

import "math"

// function to find the area and perimeter of a square
// thus it will return two values
func Square(x float64) (float64, float64) {
	//area of the square
	Area := math.Pow(x, 2)
	//perimeter of the square
	Perimeter := 4 * x

	return Area, Perimeter
}
