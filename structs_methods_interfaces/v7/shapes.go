package main

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base   float64
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

func (t Triangle) Area() float64 {
	return .5 * t.base * t.height
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}
