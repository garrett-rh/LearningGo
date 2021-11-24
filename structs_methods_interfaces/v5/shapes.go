package main

import (
	"math"
)

type Rectangle struct {
	width  float64
	height float64
}

/*
This is an object method
Syntax:
	func (receiverName ReceiverType) MethodName(args) {...}
Convention is to set the receiver variable as the first letter of the type
*/
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}
