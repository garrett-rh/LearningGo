package main

import (
	"testing"
)

/*
This is an interface
If the object has a .Area() method, then this interface can be used with it
*/

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	// Defines areaTests as an array of structs
	// Contains a shape & what the expected value is
	// Known as an anonymous struct
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		// Can do this the commented out way or the way below
		// Type hints of what each variable means
		{name: "Rectangle", shape: Rectangle{width: 12, height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{base: 12, height: 6}, hasArea: 36.0},
		// {Rectangle{12, 6}, 72.0},
		// {Circle{10}, 314.1592653589793},
		// {Triangle{12, 6}, 36.0},
	}

	// _ is the index
	// tt is the areaTests object
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.shape.Area()
			if got != tt.hasArea {
				// The %#v increases the readability of the error
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
