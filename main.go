package main

import "log"

func main() {
	//assign a variable to type triangle which the base would be 7 and height would be 9
	triange1 := Triangle{
		base:   4,
		height: 5,
	}

	circle1 := Circle{
		radius: 6,
	}

	log.Println("Area of triangle1: ", triange1.Area())           //prints the area of the triangle
	log.Println("Parimeter of triangle1: ", triange1.Parimeter()) //prints the parimeter of the triangle

	log.Println("Area of circle: ", circle1.Area())           //prints the area of circle
	log.Println("Perimeter of circle: ", circle1.Parimeter()) //prints perimeter of circle

	Asquare, Psquare := Square(8)
	log.Println("Area of square: ", Asquare, ",Perimeter of Square: ", Psquare)
}
