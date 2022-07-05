package main

import "math"

func main() {
	Prinhello()

	circle := NewCircle(5.0)

	calculateArea(&circle)
}

func Prinhello() {
	println("Hello World")
}

// Task: write a struct for a circle, you can set the radius when creating it, provide a function which returns the circle's area

type Circle struct {
	r float64
}

func NewCircle(radius float64) Circle {
	return Circle{r: radius}
}

func calculateArea(v *Circle) {
	var area = math.Pi * math.Pow(v.r, 2)

	println("this is the area of the cirle:", area)
}

// A = pr2
